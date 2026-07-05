package web

import (
	"custompbx/webStruct"
	"fmt"
	"sync"
)

type eventHandler func(*webStruct.MessageData) webStruct.UserResponse
type accessGroups func() []int
type registeredEvent struct {
	handler eventHandler
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

func (r *handlerRegistry) Dispatch(data *webStruct.MessageData) (webStruct.UserResponse, bool) {
	r.mx.RLock()
	event, ok := r.events[data.Event]
	r.mx.RUnlock()
	if !ok {
		return webStruct.UserResponse{}, false
	}
	return getUser(data, event.handler, event.groups()), true
}

func mustRegister(r *handlerRegistry, name string, handler eventHandler, groups accessGroups) {
	if err := r.Register(name, handler, groups); err != nil {
		panic(err)
	}
}

var coreEvents = func() *handlerRegistry {
	r := newHandlerRegistry()
	mustRegister(r, "AddUserToken", createAPIToken, onlyAdminGroup)
	mustRegister(r, "GetUserTokens", GetUserTokens, onlyAdminGroup)
	mustRegister(r, "UserGetOwnTokens", UserGetOwnTokens, onlyAdminGroup)
	mustRegister(r, "RemoveUserToken", RemoveUserToken, onlyAdminGroup)
	return r
}()

func normalizePagination(limit, page int) (int, int) {
	if limit <= 0 || limit > 5000 {
		limit = 250
	}
	if page < 0 {
		page = 0
	}
	return limit, page * limit
}
