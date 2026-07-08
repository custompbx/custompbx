package web

import (
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"testing"
)

func TestHandlerRegistryRejectsDuplicates(t *testing.T) {
	r := newHandlerRegistry()
	h := func(data *webStruct.MessageData) webStruct.UserResponse {
		return webStruct.UserResponse{MessageType: data.Event}
	}
	groups := func() []int { return []int{1} }
	if err := r.Register("event", h, groups); err != nil {
		t.Fatal(err)
	}
	if err := r.Register("event", h, groups); err == nil {
		t.Fatal("duplicate registration accepted")
	}
}

func TestHandlerRegistryDispatch(t *testing.T) {
	r := newHandlerRegistry()
	h := func(data *webStruct.MessageData) webStruct.UserResponse {
		return webStruct.UserResponse{MessageType: data.Event}
	}
	groups := func() []int { return []int{mainStruct.GetAdminId()} }
	if err := r.Register("event", h, groups); err != nil {
		t.Fatal(err)
	}

	data := &webStruct.MessageData{
		Event: "event",
		Context: &webStruct.WsContext{
			User: &mainStruct.WebUser{Id: 1, Login: "admin", GroupId: mainStruct.GetAdminId()},
		},
	}
	resp, ok := r.Dispatch(data, data.Context)
	if !ok {
		t.Fatal("registered event was not dispatched")
	}
	if resp.MessageType != "event" {
		t.Fatalf("message type = %q, want event", resp.MessageType)
	}
}

func TestHandlerRegistryDispatchRejectsWrongGroup(t *testing.T) {
	r := newHandlerRegistry()
	if err := r.Register("event", func(data *webStruct.MessageData) webStruct.UserResponse {
		t.Fatal("handler should not run for unauthorized group")
		return webStruct.UserResponse{}
	}, func() []int { return []int{mainStruct.GetAdminId()} }); err != nil {
		t.Fatal(err)
	}

	data := &webStruct.MessageData{
		Event: "event",
		Context: &webStruct.WsContext{
			User: &mainStruct.WebUser{Id: 2, Login: "user", GroupId: mainStruct.GetUserId()},
		},
	}
	resp, ok := r.Dispatch(data, data.Context)
	if !ok {
		t.Fatal("registered event was not dispatched")
	}
	if resp.MessageType != "no_access" {
		t.Fatalf("message type = %q, want no_access", resp.MessageType)
	}
}

func TestHandlerRegistryUnknownEventFallsBack(t *testing.T) {
	r := newHandlerRegistry()
	if _, ok := r.Dispatch(&webStruct.MessageData{Event: "unknown"}, nil); ok {
		t.Fatal("unknown event was handled by registry")
	}
}

func TestCoreRegistryIncludesMigratedEvents(t *testing.T) {
	events := []string{
		eventRelogin,
		eventLogOut,
		eventSubscriptionList,
		eventPersistentSubscription,
		webStruct.Unsubscribe,
		webStruct.DialplanDebug,
		webStruct.SubscribeHepPackages,
		eventSwitchDialplanDebug,
	}
	for _, event := range events {
		if !coreEvents.Has(event) {
			t.Fatalf("%s event is not registered", event)
		}
	}
}

func TestSubscriptionRegistryHandlers(t *testing.T) {
	t.Run("replace temporary subscriptions preserves persistent", func(t *testing.T) {
		ctx := adminContext()
		ctx.Subscriptions.Set("old")
		ctx.Subscriptions.SetPersistent("persistent")
		data := messageData(ctx, eventSubscriptionList)
		data.ArrVal = []string{"calls"}

		resp, ok := coreEvents.Dispatch(data, ctx)

		if !ok || resp.MessageType != eventSubscriptionList || resp.Error != "" {
			t.Fatalf("resp=%+v ok=%t", resp, ok)
		}
		if ctx.Subscriptions.Get("old") {
			t.Fatal("old temporary subscription survived replace")
		}
		if !ctx.Subscriptions.Get("persistent") {
			t.Fatal("persistent subscription was not preserved")
		}
		if !ctx.Subscriptions.Get("calls") {
			t.Fatal("new subscription was not added")
		}
	})

	t.Run("replace rejects empty subscription list", func(t *testing.T) {
		ctx := adminContext()
		ctx.Subscriptions.SetPersistent("persistent")
		data := messageData(ctx, eventSubscriptionList)

		resp, ok := coreEvents.Dispatch(data, ctx)

		if !ok || resp.MessageType != eventSubscriptionList || resp.Error != "can't subscribe!" {
			t.Fatalf("resp=%+v ok=%t", resp, ok)
		}
		if !ctx.Subscriptions.Get("persistent") {
			t.Fatal("persistent subscription was not preserved on rejected replace")
		}
	})

	t.Run("persistent subscription survives clear", func(t *testing.T) {
		ctx := adminContext()
		data := messageData(ctx, eventPersistentSubscription)
		data.ArrVal = []string{"persistent"}

		resp, ok := coreEvents.Dispatch(data, ctx)

		if !ok || resp.MessageType != eventPersistentSubscription || resp.Error != "" {
			t.Fatalf("resp=%+v ok=%t", resp, ok)
		}
		ctx.Subscriptions.Clear()
		if !ctx.Subscriptions.Get("persistent") {
			t.Fatal("persistent subscription did not survive clear")
		}
	})

	t.Run("unsubscribe deletes one subscription", func(t *testing.T) {
		ctx := adminContext()
		ctx.Subscriptions.SetPersistent("persistent")
		data := messageData(ctx, webStruct.Unsubscribe)
		data.Name = "persistent"

		resp, ok := coreEvents.Dispatch(data, ctx)

		if !ok || resp.MessageType != eventSubscriptionList || resp.Error != "" {
			t.Fatalf("resp=%+v ok=%t", resp, ok)
		}
		if ctx.Subscriptions.Get("persistent") {
			t.Fatal("unsubscribe did not delete persistent subscription")
		}
	})

	t.Run("unsubscribe without name clears temporary only", func(t *testing.T) {
		ctx := adminContext()
		ctx.Subscriptions.Set("temporary")
		ctx.Subscriptions.SetPersistent("persistent")
		data := messageData(ctx, webStruct.Unsubscribe)

		resp, ok := coreEvents.Dispatch(data, ctx)

		if !ok || resp.MessageType != eventSubscriptionList || resp.Error != "" {
			t.Fatalf("resp=%+v ok=%t", resp, ok)
		}
		if ctx.Subscriptions.Get("temporary") {
			t.Fatal("temporary subscription survived clear")
		}
		if !ctx.Subscriptions.Get("persistent") {
			t.Fatal("persistent subscription was not preserved")
		}
	})

	t.Run("wrong group rejected", func(t *testing.T) {
		ctx := userContext()
		data := messageData(ctx, eventSubscriptionList)
		data.ArrVal = []string{"calls"}

		resp, ok := coreEvents.Dispatch(data, ctx)

		if !ok || resp.MessageType != "no_access" {
			t.Fatalf("resp=%+v ok=%t", resp, ok)
		}
		if ctx.Subscriptions.Get("calls") {
			t.Fatal("unauthorized subscription was added")
		}
	})
}

func adminContext() *webStruct.WsContext {
	ctx := webStruct.CreateWsContext(nil)
	ctx.SetUser(&mainStruct.WebUser{Id: 1, Login: "admin", GroupId: mainStruct.GetAdminId()})
	return ctx
}

func userContext() *webStruct.WsContext {
	ctx := webStruct.CreateWsContext(nil)
	ctx.SetUser(&mainStruct.WebUser{Id: 2, Login: "user", GroupId: mainStruct.GetUserId()})
	return ctx
}

func messageData(ctx *webStruct.WsContext, event string) *webStruct.MessageData {
	return &webStruct.MessageData{Event: event, Context: ctx}
}

func TestNormalizePagination(t *testing.T) {
	limit, offset := normalizePagination(0, -1)
	if limit != 250 || offset != 0 {
		t.Fatalf("got %d, %d", limit, offset)
	}
	limit, offset = normalizePagination(100, 3)
	if limit != 100 || offset != 300 {
		t.Fatalf("got %d, %d", limit, offset)
	}
}
