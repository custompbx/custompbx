package web

import (
	"custompbx/webStruct"
	"fmt"
	"sync"
)

type eventHandler func(*webStruct.MessageData) webStruct.UserResponse
type contextEventHandler func(*webStruct.MessageData, *webStruct.WsContext) webStruct.UserResponse
type accessGroups func() []int
type registeredEvent struct {
	handler contextEventHandler
	groups  accessGroups
}

type handlerRegistry struct {
	mx     sync.RWMutex
	events map[string]registeredEvent
}

func newHandlerRegistry() *handlerRegistry {
	return &handlerRegistry{events: make(map[string]registeredEvent)}
}

func (r *handlerRegistry) Register(name string, handler eventHandler, groups accessGroups) error {
	if handler == nil {
		return fmt.Errorf("invalid event registration %q", name)
	}
	return r.RegisterWithContext(name, func(data *webStruct.MessageData, _ *webStruct.WsContext) webStruct.UserResponse {
		return handler(data)
	}, groups)
}

func (r *handlerRegistry) RegisterWithContext(name string, handler contextEventHandler, groups accessGroups) error {
	if name == "" || handler == nil || groups == nil {
		return fmt.Errorf("invalid event registration %q", name)
	}
	r.mx.Lock()
	defer r.mx.Unlock()
	if _, exists := r.events[name]; exists {
		return fmt.Errorf("event %q already registered", name)
	}
	r.events[name] = registeredEvent{handler: handler, groups: groups}
	return nil
}

func (r *handlerRegistry) Dispatch(data *webStruct.MessageData, wsContext *webStruct.WsContext) (webStruct.UserResponse, bool) {
	r.mx.RLock()
	event, ok := r.events[data.Event]
	r.mx.RUnlock()
	if !ok {
		return webStruct.UserResponse{}, false
	}
	if resp := checkAccessGroup(data, event.groups()); resp != nil {
		return *resp, true
	}
	return event.handler(data, wsContext), true
}

func (r *handlerRegistry) Has(name string) bool {
	r.mx.RLock()
	defer r.mx.RUnlock()
	_, ok := r.events[name]
	return ok
}

func mustRegister(r *handlerRegistry, name string, handler eventHandler, groups accessGroups) {
	if err := r.Register(name, handler, groups); err != nil {
		panic(err)
	}
}

func logoutAndClearSubscriptions(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Context != nil && data.Context.Subscriptions != nil {
		data.Context.Subscriptions.Clear()
	}
	return loginOut(data)
}

func replaceSubscriptions(data *webStruct.MessageData, wsContext *webStruct.WsContext) webStruct.UserResponse {
	resp := webStruct.UserResponse{MessageType: eventSubscriptionList}
	wsContext.Subscriptions.Clear()
	if len(data.ArrVal) > 10 || len(data.ArrVal) == 0 {
		resp.Error = "can't subscribe!"
		return resp
	}
	for _, name := range data.ArrVal {
		wsContext.Subscriptions.Set(name)
	}
	return resp
}

func addPersistentSubscriptions(data *webStruct.MessageData, wsContext *webStruct.WsContext) webStruct.UserResponse {
	resp := webStruct.UserResponse{MessageType: eventPersistentSubscription}
	if len(data.ArrVal) > 10 || len(data.ArrVal) == 0 {
		resp.Error = "can't subscribe!"
		return resp
	}
	for _, name := range data.ArrVal {
		wsContext.Subscriptions.SetPersistent(name)
	}
	return resp
}

func unsubscribe(data *webStruct.MessageData, wsContext *webStruct.WsContext) webStruct.UserResponse {
	if data.Name != "" {
		wsContext.Subscriptions.Del(data.Name)
	} else {
		wsContext.Subscriptions.Clear()
	}
	return webStruct.UserResponse{MessageType: eventSubscriptionList}
}

var coreEvents = func() *handlerRegistry {
	r := newHandlerRegistry()
	mustRegister(r, eventRelogin, checkRelogin, onlyAdminGroup)
	mustRegister(r, eventLogOut, logoutAndClearSubscriptions, onlyAdminGroup)
	mustRegister(r, "AddUserToken", createAPIToken, onlyAdminGroup)
	mustRegister(r, "GetUserTokens", GetUserTokens, onlyAdminGroup)
	mustRegister(r, "UserGetOwnTokens", UserGetOwnTokens, onlyAdminGroup)
	mustRegister(r, "RemoveUserToken", RemoveUserToken, onlyAdminGroup)
	mustRegister(r, webStruct.DialplanDebug, getDialplanDebug, onlyAdminGroup)
	mustRegister(r, webStruct.SubscribeHepPackages, getDialplanDebug, onlyAdminGroup)
	mustRegister(r, eventSwitchDialplanDebug, switchDialplanDebug, onlyAdminGroup)
	mustRegisterContext(r, eventSubscriptionList, replaceSubscriptions, onlyAdminGroup)
	mustRegisterContext(r, eventPersistentSubscription, addPersistentSubscriptions, onlyAdminGroup)
	mustRegisterContext(r, webStruct.Unsubscribe, unsubscribe, onlyAdminGroup)
	return r
}()

func mustRegisterContext(r *handlerRegistry, name string, handler contextEventHandler, groups accessGroups) {
	if err := r.RegisterWithContext(name, handler, groups); err != nil {
		panic(err)
	}
}
