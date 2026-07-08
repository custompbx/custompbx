package web

import (
	"custompbx/daemonCache"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestMessageHandlerRejectsInvalidMessage(t *testing.T) {
	context, client := testMessageHandlerContext(t)
	go context.SendWaiter()

	messageHandler(&webStruct.Message{Event: "bad"}, context)

	response := readWSResponse(t, client)
	if response.MessageType != "none" || response.Error != "invalid message" {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestMessageHandlerReturnsDaemonStatusWhenDatabaseDown(t *testing.T) {
	oldState := daemonCache.State
	daemonCache.State = &mainStruct.DaemonState{DatabaseConnection: false, ESLConnection: true}
	t.Cleanup(func() { daemonCache.State = oldState })

	context, client := testMessageHandlerContext(t)
	go context.SendWaiter()

	messageHandler(&webStruct.Message{Event: "anything", Data: &webStruct.MessageData{}}, context)

	response := readWSResponse(t, client)
	if response.MessageType != webStruct.BroadcastConnection {
		t.Fatalf("message type = %q, want %q", response.MessageType, webStruct.BroadcastConnection)
	}
	if response.Daemon == nil || response.Daemon.DatabaseConnection {
		t.Fatalf("unexpected daemon state: %+v", response.Daemon)
	}
}

func testMessageHandlerContext(t *testing.T) (*webStruct.WsContext, *websocket.Conn) {
	t.Helper()
	contexts := make(chan *webStruct.WsContext, 1)
	done := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := (&websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}).Upgrade(w, r, nil)
		if err != nil {
			return
		}
		context := webStruct.CreateWsContext(ws)
		contexts <- context
		<-done
	}))
	t.Cleanup(func() { close(done) })
	t.Cleanup(server.Close)

	client, _, err := websocket.DefaultDialer.Dial("ws"+strings.TrimPrefix(server.URL, "http"), nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = client.Close() })

	select {
	case context := <-contexts:
		t.Cleanup(func() { _ = context.Close() })
		return context, client
	case <-time.After(time.Second):
		t.Fatal("websocket context was not created")
		return nil, nil
	}
}

func readWSResponse(t *testing.T, client *websocket.Conn) webStruct.UserResponse {
	t.Helper()
	if err := client.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	_, raw, err := client.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	var response webStruct.UserResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatal(err)
	}
	return response
}
