package webStruct

import (
	"custompbx/mainStruct"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestSubscriptionsClearPreservesOnlyPersistent(t *testing.T) {
	s := newSubscriptions()
	s.Set("temporary")
	s.SetPersistent("persistent")
	s.Clear()
	if s.Get("temporary") {
		t.Fatal("temporary subscription survived clear")
	}
	if !s.Get("persistent") {
		t.Fatal("persistent subscription was cleared")
	}
	s.Del("persistent")
	if s.Get("persistent") {
		t.Fatal("deleted subscription remains")
	}
}

func testWebSocketContext(t *testing.T) (*WsContext, *websocket.Conn) {
	t.Helper()
	contexts := make(chan *WsContext, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := (&websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}).Upgrade(w, r, nil)
		if err != nil {
			return
		}
		context := CreateWsContext(ws)
		contexts <- context
		<-context.done
	}))
	t.Cleanup(server.Close)
	client, _, err := websocket.DefaultDialer.Dial("ws"+strings.TrimPrefix(server.URL, "http"), nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return <-contexts, client
}

func TestHubDisconnectsQueueOverflowAndClosesOnce(t *testing.T) {
	context, _ := testWebSocketContext(t)
	hub := NewWsHub()
	hub.Register(context)
	context.Subscriptions.Set("event")
	for i := 0; i < cap(context.send); i++ {
		if !context.Enqueue(&UserResponse{MessageType: "queued"}) {
			t.Fatal("queue filled early")
		}
	}
	hub.Broadcast(UserResponse{MessageType: "event"})
	if got := hub.Metrics(); got.Active != 0 || got.QueueOverflows != 1 {
		t.Fatalf("unexpected metrics: %+v", got)
	}
	if err := context.Close(); err != nil {
		t.Fatalf("second close failed: %v", err)
	}
}

func TestHubRegisterRemoveConcurrent(t *testing.T) {
	hub := NewWsHub()
	var wg sync.WaitGroup
	for i := 0; i < 25; i++ {
		context, _ := testWebSocketContext(t)
		wg.Add(1)
		go func(c *WsContext) {
			defer wg.Done()
			hub.Register(c)
			_ = c.Close()
		}(context)
	}
	wg.Wait()
	if got := hub.Metrics(); got.Active != 0 {
		t.Fatalf("active connections = %d", got.Active)
	}
}

func TestHubShutdownClosesConnectionsAndStopsBroadcasts(t *testing.T) {
	hub := NewWsHub()
	context, _ := testWebSocketContext(t)
	hub.Register(context)
	context.Subscriptions.Set("event")
	hub.Shutdown()

	got := hub.Metrics()
	if got.Active != 0 || !got.ShuttingDown {
		t.Fatalf("unexpected metrics after shutdown: %+v", got)
	}
	hub.Broadcast(UserResponse{MessageType: "event"})
	if after := hub.Metrics(); after.Broadcasts != got.Broadcasts {
		t.Fatalf("broadcast counted after shutdown: before=%+v after=%+v", got, after)
	}

	late, _ := testWebSocketContext(t)
	hub.Register(late)
	if after := hub.Metrics(); after.Active != 0 {
		t.Fatalf("late register survived shutdown: %+v", after)
	}
}

func TestHubUnicastUsesStableUserID(t *testing.T) {
	hub := NewWsHub()
	first, _ := testWebSocketContext(t)
	second, _ := testWebSocketContext(t)
	first.SetUser(&mainStruct.WebUser{Id: 101})
	second.SetUser(&mainStruct.WebUser{Id: 202})
	hub.Register(first)
	hub.Register(second)

	sent := hub.Unicast(UserResponse{MessageType: "event"}, []*mainStruct.WebUser{{Id: 202}})
	if len(sent) != 1 || sent[0] != 202 {
		t.Fatalf("sent users = %v", sent)
	}
	if len(first.send) != 0 || len(second.send) != 1 {
		t.Fatalf("unexpected queue lengths: first=%d second=%d", len(first.send), len(second.send))
	}
	_ = first.Close()
	_ = second.Close()
}

func TestReadWaiterProcessesMessagesInOrder(t *testing.T) {
	context, client := testWebSocketContext(t)
	var mx sync.Mutex
	var events []string
	done := make(chan struct{})
	go context.ReadWaiter(func(message *Message, _ *WsContext) {
		mx.Lock()
		events = append(events, message.Event)
		count := len(events)
		mx.Unlock()
		if count == 2 {
			close(done)
		}
	})
	for _, event := range []string{"first", "second"} {
		if err := client.WriteJSON(Message{Event: event, Data: &MessageData{}}); err != nil {
			t.Fatal(err)
		}
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("messages not handled")
	}
	mx.Lock()
	defer mx.Unlock()
	if len(events) != 2 || events[0] != "first" || events[1] != "second" {
		t.Fatalf("events = %v", events)
	}
	_ = context.Close()
}

func TestReadWaiterRejectsMalformedMessagesWithoutCallingHandler(t *testing.T) {
	context, client := testWebSocketContext(t)
	go context.SendWaiter()
	handled := make(chan struct{}, 1)
	go context.ReadWaiter(func(message *Message, _ *WsContext) {
		handled <- struct{}{}
	})

	if err := client.WriteMessage(websocket.TextMessage, []byte(`{"event":"bad","data":null}`)); err != nil {
		t.Fatal(err)
	}
	_, raw, err := client.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	var response UserResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatal(err)
	}
	if response.MessageType != "none" || response.Error == "" {
		t.Fatalf("unexpected response: %+v", response)
	}
	select {
	case <-handled:
		t.Fatal("handler was called for malformed message")
	case <-time.After(100 * time.Millisecond):
	}
	_ = context.Close()
}

func TestSubscriptionsConcurrentAccess(t *testing.T) {
	s := newSubscriptions()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); s.Set("event"); _ = s.Get("event"); s.Clear() }()
	}
	wg.Wait()
}
