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

func registeredHandler(name string, fallback eventHandler, overrides map[string]eventHandler) eventHandler {
	if overrides == nil {
		return fallback
	}
	if handler, ok := overrides[name]; ok {
		return handler
	}
	return fallback
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

func updateSettings(data *webStruct.MessageData) webStruct.UserResponse {
	if err := data.Payload.Web.NormalizeAndValidateOrigins(); err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: "settings"}
	}
	data.Payload.Web.Route = normalizeRoute(data.Payload.Web.Route)
	data.Payload.XMLCurl.Route = normalizeRoute(data.Payload.XMLCurl.Route)
	return setSettings(data)
}

func buildCoreEvents(overrides map[string]eventHandler) *handlerRegistry {
	r := newHandlerRegistry()
	registerCoreSystemEvents(r, overrides)
	registerCoreWebUserEvents(r, overrides)
	registerCoreDirectoryTemplateEvents(r, overrides)
	registerCoreConversationEvents(r, overrides)
	registerCoreAuthTokenEvents(r, overrides)
	registerCoreDebugEvents(r, overrides)
	registerCoreSubscriptionEvents(r)
	return r
}

func registerCoreSystemEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventGetSettings, registeredHandler(eventGetSettings, checkSettings, overrides), onlyAdminGroup)
	mustRegister(r, eventSetSettings, registeredHandler(eventSetSettings, updateSettings, overrides), onlyAdminGroup)
	mustRegister(r, webStruct.GetDashboard, registeredHandler(webStruct.GetDashboard, getDashboardData, overrides), onlyAdminGroup)
	mustRegister(r, eventGetInstances, registeredHandler(eventGetInstances, GetInstances, overrides), onlyAdminGroup)
	mustRegister(r, eventUpdateInstanceDescription, registeredHandler(eventUpdateInstanceDescription, UpdateInstanceDescription, overrides), onlyAdminGroup)
	mustRegister(r, eventGetWebSettings, registeredHandler(eventGetWebSettings, GetWebSettings, overrides), onlyAdminGroup)
	mustRegister(r, eventSaveWebSettings, registeredHandler(eventSaveWebSettings, SaveWebSettings, overrides), onlyAdminGroup)
	mustRegister(r, eventGetCDR, registeredHandler(eventGetCDR, getCDR, overrides), onlyAdminGroup)
	mustRegister(r, eventGetHEP, registeredHandler(eventGetHEP, getHEP, overrides), onlyAdminGroup)
	mustRegister(r, eventGetHEPDetails, registeredHandler(eventGetHEPDetails, GetHEPDetails, overrides), onlyAdminGroup)
	mustRegister(r, eventGetLogs, registeredHandler(eventGetLogs, GetLogs, overrides), onlyAdminGroup)
	mustRegister(r, eventGetPhoneCreds, registeredHandler(eventGetPhoneCreds, getPhoneCreds, overrides), onlyAdminGroup)
	mustRegister(r, eventSendFSCLICommand, registeredHandler(eventSendFSCLICommand, runCLICommand, overrides), onlyAdminGroup)
	mustRegister(r, eventRealFSCLIConnect, registeredHandler(eventRealFSCLIConnect, RealFSCLIConnect, overrides), onlyAdminGroup)
	mustRegister(r, eventRealFSCLICommand, registeredHandler(eventRealFSCLICommand, RealFSCLICommand, overrides), onlyAdminGroup)
}

func registerCoreWebUserEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventSettingsUsersGet, registeredHandler(eventSettingsUsersGet, getWebUsers, overrides), onlyAdminGroup)
	mustRegister(r, eventGetWebUsersByDirectory, registeredHandler(eventGetWebUsersByDirectory, GetWebUsersByDirectory, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersAdd, registeredHandler(eventSettingsUsersAdd, addWebUsers, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersRename, registeredHandler(eventSettingsUsersRename, renameWebUsers, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersDelete, registeredHandler(eventSettingsUsersDelete, deleteWebUsers, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersSwitch, registeredHandler(eventSettingsUsersSwitch, switchWebUser, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdatePass, registeredHandler(eventSettingsUsersUpdatePass, updateWebUsersPassword, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateLang, registeredHandler(eventSettingsUsersUpdateLang, updateWebUsersLang, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateSip, registeredHandler(eventSettingsUsersUpdateSip, updateWebUsersSipUser, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateWS, registeredHandler(eventSettingsUsersUpdateWS, updateWebUsersWs, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateVerto, registeredHandler(eventSettingsUsersUpdateVerto, updateWebUsersVertoWs, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateRTC, registeredHandler(eventSettingsUsersUpdateRTC, UpdateWebUserWebRTCLib, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateStun, registeredHandler(eventSettingsUsersUpdateStun, updateWebUsersStun, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersUpdateAvatar, registeredHandler(eventSettingsUsersUpdateAvatar, updateWebUsersAvatar, overrides), onlyAdminGroup)
	mustRegister(r, eventSettingsUsersClearAvatar, registeredHandler(eventSettingsUsersClearAvatar, clearWebUsersAvatar, overrides), onlyAdminGroup)
}

func registerCoreDirectoryTemplateEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventUpdateWebUserGroup, registeredHandler(eventUpdateWebUserGroup, UpdateWebUserGroup, overrides), onlyAdminGroup)
	mustRegister(r, eventGetWebDirUserTemplates, registeredHandler(eventGetWebDirUserTemplates, GetWebDirectoryUsersTemplates, overrides), onlyAdminGroup)
	mustRegister(r, eventAddWebDirUserTemplate, registeredHandler(eventAddWebDirUserTemplate, AddWebDirectoryUsersTemplate, overrides), onlyAdminGroup)
	mustRegister(r, eventDelWebDirUserTemplate, registeredHandler(eventDelWebDirUserTemplate, DelWebDirectoryUsersTemplate, overrides), onlyAdminGroup)
	mustRegister(r, eventUpdateWebDirUserTemplate, registeredHandler(eventUpdateWebDirUserTemplate, UpdateWebDirectoryUsersTemplate, overrides), onlyAdminGroup)
	mustRegister(r, eventSwitchWebDirUserTemplate, registeredHandler(eventSwitchWebDirUserTemplate, SwitchWebDirectoryUsersTemplate, overrides), onlyAdminGroup)
	mustRegister(r, eventGetWebDirUserTplParams, registeredHandler(eventGetWebDirUserTplParams, GetWebDirectoryUsersTemplateParameters, overrides), onlyAdminGroup)
	mustRegister(r, eventAddWebDirUserTplParam, registeredHandler(eventAddWebDirUserTplParam, AddWebDirectoryUsersTemplateParameter, overrides), onlyAdminGroup)
	mustRegister(r, eventDelWebDirUserTplParam, registeredHandler(eventDelWebDirUserTplParam, DelWebDirectoryUsersTemplateParameter, overrides), onlyAdminGroup)
	mustRegister(r, eventSwitchWebDirUserTplParam, registeredHandler(eventSwitchWebDirUserTplParam, SwitchWebDirectoryUsersTemplateParameter, overrides), onlyAdminGroup)
	mustRegister(r, eventUpdateWebDirUserTplParam, registeredHandler(eventUpdateWebDirUserTplParam, UpdateWebDirectoryUsersTemplateParameter, overrides), onlyAdminGroup)
	mustRegister(r, eventGetWebDirUserTplVars, registeredHandler(eventGetWebDirUserTplVars, GetWebDirectoryUsersTemplateVariables, overrides), onlyAdminGroup)
	mustRegister(r, eventAddWebDirUserTplVar, registeredHandler(eventAddWebDirUserTplVar, AddWebDirectoryUsersTemplateVariable, overrides), onlyAdminGroup)
	mustRegister(r, eventDelWebDirUserTplVar, registeredHandler(eventDelWebDirUserTplVar, DelWebDirectoryUsersTemplateVariable, overrides), onlyAdminGroup)
	mustRegister(r, eventSwitchWebDirUserTplVar, registeredHandler(eventSwitchWebDirUserTplVar, SwitchWebDirectoryUsersTemplateVariable, overrides), onlyAdminGroup)
	mustRegister(r, eventUpdateWebDirUserTplVar, registeredHandler(eventUpdateWebDirUserTplVar, UpdateWebDirectoryUsersTemplateVariable, overrides), onlyAdminGroup)
	mustRegister(r, eventGetWebDirUserTplList, registeredHandler(eventGetWebDirUserTplList, GetWebDirectoryUsersTemplatesList, overrides), onlyAdminAndManagerGroup)
	mustRegister(r, eventGetWebDirUserTplForm, registeredHandler(eventGetWebDirUserTplForm, GetWebDirectoryUsersTemplateForm, overrides), onlyAdminAndManagerGroup)
	mustRegister(r, eventCreateWebDirUserByTpl, registeredHandler(eventCreateWebDirUserByTpl, CreateWebDirectoryUsersByTemplate, overrides), onlyAdminAndManagerGroup)
}

func registerCoreConversationEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventGetConvPrivateMessages, registeredHandler(eventGetConvPrivateMessages, GetConversationPrivateMessages, overrides), onlyAdminGroup)
	mustRegister(r, eventGetConvPrivateCalls, registeredHandler(eventGetConvPrivateCalls, GetConversationPrivateCalls, overrides), onlyAdminGroup)
	mustRegister(r, eventGetConvRoomMessages, registeredHandler(eventGetConvRoomMessages, GetConversationRoomMessages, overrides), onlyAdminGroup)
	mustRegister(r, eventSendConvPrivateMessage, registeredHandler(eventSendConvPrivateMessage, SendConversationPrivateMessage, overrides), onlyAdminGroup)
	mustRegister(r, eventSendConvPrivateCall, registeredHandler(eventSendConvPrivateCall, SendConversationPrivateCall, overrides), onlyAdminGroup)
	mustRegister(r, eventSendConvRoomMessage, registeredHandler(eventSendConvRoomMessage, SendConversationRoomMessage, overrides), onlyAdminGroup)
}

func registerCoreAuthTokenEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventRelogin, registeredHandler(eventRelogin, checkRelogin, overrides), onlyAdminGroup)
	mustRegister(r, eventLogOut, registeredHandler(eventLogOut, logoutAndClearSubscriptions, overrides), onlyAdminGroup)
	mustRegister(r, "AddUserToken", registeredHandler("AddUserToken", createAPIToken, overrides), onlyAdminGroup)
	mustRegister(r, "GetUserTokens", registeredHandler("GetUserTokens", GetUserTokens, overrides), onlyAdminGroup)
	mustRegister(r, "UserGetOwnTokens", registeredHandler("UserGetOwnTokens", UserGetOwnTokens, overrides), onlyAdminGroup)
	mustRegister(r, "RemoveUserToken", registeredHandler("RemoveUserToken", RemoveUserToken, overrides), onlyAdminGroup)
}

func registerCoreDebugEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, webStruct.DialplanDebug, registeredHandler(webStruct.DialplanDebug, getDialplanDebug, overrides), onlyAdminGroup)
	mustRegister(r, webStruct.SubscribeHepPackages, registeredHandler(webStruct.SubscribeHepPackages, getDialplanDebug, overrides), onlyAdminGroup)
	mustRegister(r, eventSwitchDialplanDebug, registeredHandler(eventSwitchDialplanDebug, switchDialplanDebug, overrides), onlyAdminGroup)
}

func registerCoreSubscriptionEvents(r *handlerRegistry) {
	mustRegisterContext(r, eventSubscriptionList, replaceSubscriptions, onlyAdminGroup)
	mustRegisterContext(r, eventPersistentSubscription, addPersistentSubscriptions, onlyAdminGroup)
	mustRegisterContext(r, webStruct.Unsubscribe, unsubscribe, onlyAdminGroup)
}

var coreEvents = buildCoreEvents(nil)

func mustRegisterContext(r *handlerRegistry, name string, handler contextEventHandler, groups accessGroups) {
	if err := r.RegisterWithContext(name, handler, groups); err != nil {
		panic(err)
	}
}

func normalizeRoute(route string) string {
	if route == "" || route[0] == '/' {
		return route
	}
	return "/" + route
}
