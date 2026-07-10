package webStruct

import (
	"custompbx/cfg"
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
	fillOutboundQueue(t, context)
	hub.Broadcast(UserResponse{MessageType: "event"})
	if got := hub.Metrics(); got.Active != 0 || got.QueueOverflows != 1 {
		t.Fatalf("unexpected metrics: %+v", got)
	}
	if err := context.Close(); err != nil {
		t.Fatalf("second close failed: %v", err)
	}
}

func TestHubBroadcastDropsSlowClientAndContinues(t *testing.T) {
	hub := NewWsHub()
	slow, _ := testWebSocketContext(t)
	fast, _ := testWebSocketContext(t)
	unsubscribed, _ := testWebSocketContext(t)
	slow.Subscriptions.Set("event")
	fast.Subscriptions.Set("event")
	hub.Register(slow)
	hub.Register(fast)
	hub.Register(unsubscribed)
	fillOutboundQueue(t, slow)

	done := make(chan struct{})
	go func() {
		hub.Broadcast(UserResponse{MessageType: "event"})
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("broadcast blocked on slow client")
	}

	if got := hub.Metrics(); got.Active != 2 || got.QueueOverflows != 1 || got.Broadcasts != 1 {
		t.Fatalf("unexpected metrics: %+v", got)
	}
	if len(fast.send) != 1 {
		t.Fatalf("fast client queue length = %d, want 1", len(fast.send))
	}
	if len(unsubscribed.send) != 0 {
		t.Fatalf("unsubscribed client queue length = %d, want 0", len(unsubscribed.send))
	}
	_ = fast.Close()
	_ = unsubscribed.Close()
}

func TestHubBroadcastConnectionBypassesSubscriptionFilter(t *testing.T) {
	hub := NewWsHub()
	context, _ := testWebSocketContext(t)
	hub.Register(context)

	hub.Broadcast(UserResponse{MessageType: BroadcastConnection})

	if len(context.send) != 1 {
		t.Fatalf("queue length = %d, want 1", len(context.send))
	}
	_ = context.Close()
}

func TestWsContextCloseWithNilWebSocketIsSafeAndIdempotent(t *testing.T) {
	context := CreateWsContext(nil)
	closed := 0
	context.onClose = func(*WsContext) { closed++ }

	if err := context.CloseWithReason("unit test"); err != nil {
		t.Fatalf("close returned error: %v", err)
	}
	if err := context.CloseWithReason("second close"); err != nil {
		t.Fatalf("second close returned error: %v", err)
	}
	if closed != 1 {
		t.Fatalf("onClose called %d times, want 1", closed)
	}
	if context.Enqueue(&UserResponse{MessageType: "after-close"}) {
		t.Fatal("enqueue succeeded after close")
	}
}

func TestWaitersWithNilWebSocketCloseSafely(t *testing.T) {
	context := CreateWsContext(nil)
	closed := 0
	context.onClose = func(*WsContext) { closed++ }

	context.SendWaiter()
	context.ReadWaiter(func(*Message, *WsContext) {
		t.Fatal("handler should not run without websocket")
	})

	if closed != 1 {
		t.Fatalf("close callback called %d times, want 1", closed)
	}
	if context.Enqueue(&UserResponse{MessageType: "after-close"}) {
		t.Fatal("enqueue succeeded after nil websocket close")
	}
}

func TestHubRegisterDuplicateConnectionIsIdempotent(t *testing.T) {
	hub := NewWsHub()
	context, _ := testWebSocketContext(t)

	hub.Register(context)
	hub.Register(context)

	if got := hub.Metrics(); got.Active != 1 {
		t.Fatalf("active connections = %d, want 1", got.Active)
	}
	_ = context.Close()
	if got := hub.Metrics(); got.Active != 0 {
		t.Fatalf("active connections after close = %d, want 0", got.Active)
	}
}

func TestHubRegisterDuplicateIDDoesNotReplaceExistingConnection(t *testing.T) {
	hub := NewWsHub()
	original, _ := testWebSocketContext(t)
	duplicate, _ := testWebSocketContext(t)
	duplicate.ID = original.ID

	hub.Register(original)
	hub.Register(duplicate)

	if got := hub.Metrics(); got.Active != 1 {
		t.Fatalf("active connections = %d, want 1", got.Active)
	}
	if duplicate.Enqueue(&UserResponse{MessageType: "closed"}) {
		t.Fatal("duplicate connection was not closed")
	}
	original.Subscriptions.Set("event")
	hub.Broadcast(UserResponse{MessageType: "event"})
	if len(original.send) != 1 {
		t.Fatalf("original queue length = %d, want 1", len(original.send))
	}
	_ = original.Close()
}

func TestHubUnregisterIgnoresUnregisteredConnection(t *testing.T) {
	hub := NewWsHub()
	registered, _ := testWebSocketContext(t)
	unregistered, _ := testWebSocketContext(t)
	registered.Subscriptions.Set("event")
	hub.Register(registered)

	hub.unregister(unregistered)
	hub.Broadcast(UserResponse{MessageType: "event"})

	if got := hub.Metrics(); got.Active != 1 {
		t.Fatalf("active connections = %d, want 1", got.Active)
	}
	if len(registered.send) != 1 {
		t.Fatalf("registered queue length = %d, want 1", len(registered.send))
	}
	_ = registered.Close()
	_ = unregistered.Close()
}

func TestCreateWsContextUsesConfiguredQueueSize(t *testing.T) {
	oldWeb := cfg.CustomPbx.Web
	t.Cleanup(func() { cfg.CustomPbx.Web = oldWeb })
	cfg.CustomPbx.Web.WebSocketQueueSize = 3

	context := CreateWsContext(nil)
	defer context.Close()

	if got := cap(context.send); got != 3 {
		t.Fatalf("queue capacity = %d, want 3", got)
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
	hub.Shutdown()
	if after := hub.Metrics(); after.Active != 0 || !after.ShuttingDown || after.Broadcasts != got.Broadcasts {
		t.Fatalf("shutdown was not idempotent: before=%+v after=%+v", got, after)
	}

	late, _ := testWebSocketContext(t)
	hub.Register(late)
	if after := hub.Metrics(); after.Active != 0 {
		t.Fatalf("late register survived shutdown: %+v", after)
	}
	if late.Enqueue(&UserResponse{MessageType: "after-shutdown"}) {
		t.Fatal("late connection accepted outbound message after shutdown")
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

func TestHubUnicastDropsFullMatchingConnectionAndContinues(t *testing.T) {
	hub := NewWsHub()
	full, _ := testWebSocketContext(t)
	ready, _ := testWebSocketContext(t)
	other, _ := testWebSocketContext(t)
	full.SetUser(&mainStruct.WebUser{Id: 303})
	ready.SetUser(&mainStruct.WebUser{Id: 303})
	other.SetUser(&mainStruct.WebUser{Id: 404})
	hub.Register(full)
	hub.Register(ready)
	hub.Register(other)
	fillOutboundQueue(t, full)

	sent := hub.Unicast(UserResponse{MessageType: "event"}, []*mainStruct.WebUser{nil, {Id: 303}})

	if len(sent) != 1 || sent[0] != 303 {
		t.Fatalf("sent users = %v, want [303]", sent)
	}
	if got := hub.Metrics(); got.Active != 2 || got.QueueOverflows != 1 {
		t.Fatalf("unexpected metrics: %+v", got)
	}
	if len(ready.send) != 1 {
		t.Fatalf("ready client queue length = %d, want 1", len(ready.send))
	}
	if len(other.send) != 0 {
		t.Fatalf("other client queue length = %d, want 0", len(other.send))
	}
	_ = ready.Close()
	_ = other.Close()
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

func fillOutboundQueue(t *testing.T, context *WsContext) {
	t.Helper()
	for i := 0; i < cap(context.send); i++ {
		if !context.Enqueue(&UserResponse{MessageType: "queued"}) {
			t.Fatal("queue filled early")
		}
	}
}
