package web

import (
	"custompbx/webStruct"
	"testing"
)

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
