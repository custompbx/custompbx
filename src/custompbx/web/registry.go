package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"fmt"
	"strconv"
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

func (r *handlerRegistry) EventNames() []string {
	r.mx.RLock()
	defer r.mx.RUnlock()
	names := make([]string, 0, len(r.events))
	for name := range r.events {
		names = append(names, name)
	}
	return names
}

func mustRegister(r *handlerRegistry, name string, handler eventHandler, groups accessGroups) {
	if err := r.Register(name, handler, groups); err != nil {
		panic(err)
	}
}

func mustRegisterAdmin(r *handlerRegistry, name string, handler eventHandler, overrides map[string]eventHandler) {
	mustRegister(r, name, registeredHandler(name, handler, overrides), adminOnly)
}

func mustRegisterAdminOrManager(r *handlerRegistry, name string, handler eventHandler, overrides map[string]eventHandler) {
	mustRegister(r, name, registeredHandler(name, handler, overrides), adminOrManager)
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
	registerCoreDirectoryDomainEvents(r, overrides)
	registerCoreModuleEvents(r, overrides)
	registerCoreACLEvents(r, overrides)
	registerCoreSofiaEvents(r, overrides)
	registerCoreCDREvents(r, overrides)
	registerCoreLCREvents(r, overrides)
	registerCoreVertoConfigEvents(r, overrides)
	registerCoreSimpleModuleSettingEvents(r, overrides)
	registerCorePostSwitchEvents(r, overrides)
	registerCoreDirectoryConfigEvents(r, overrides)
	registerCoreFifoEvents(r, overrides)
	registerCoreTelephonyModuleEvents(r, overrides)
	registerCoreConferenceEvents(r, overrides)
	registerCoreRemainingConfigEvents(r, overrides)
	registerCoreConversationEvents(r, overrides)
	registerCoreDialplanEvents(r, overrides)
	registerCoreAuthTokenEvents(r, overrides)
	registerCoreDebugEvents(r, overrides)
	registerCoreSubscriptionEvents(r)
	return r
}

func registerCoreSystemEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventGetSettings, registeredHandler(eventGetSettings, checkSettings, overrides), adminOnly)
	mustRegister(r, eventSetSettings, registeredHandler(eventSetSettings, updateSettings, overrides), adminOnly)
	mustRegister(r, webStruct.GetDashboard, registeredHandler(webStruct.GetDashboard, getDashboardData, overrides), adminOnly)
	mustRegister(r, eventGetInstances, registeredHandler(eventGetInstances, GetInstances, overrides), adminOnly)
	mustRegister(r, eventUpdateInstanceDescription, registeredHandler(eventUpdateInstanceDescription, UpdateInstanceDescription, overrides), adminOnly)
	mustRegister(r, eventGetWebSettings, registeredHandler(eventGetWebSettings, GetWebSettings, overrides), adminOnly)
	mustRegister(r, eventSaveWebSettings, registeredHandler(eventSaveWebSettings, SaveWebSettings, overrides), adminOnly)
	mustRegister(r, eventGetCDR, registeredHandler(eventGetCDR, getCDR, overrides), adminOnly)
	mustRegister(r, eventGetHEP, registeredHandler(eventGetHEP, getHEP, overrides), adminOnly)
	mustRegister(r, eventGetHEPDetails, registeredHandler(eventGetHEPDetails, GetHEPDetails, overrides), adminOnly)
	mustRegister(r, eventGetLogs, registeredHandler(eventGetLogs, GetLogs, overrides), adminOnly)
	mustRegister(r, eventGetPhoneCreds, registeredHandler(eventGetPhoneCreds, getPhoneCreds, overrides), adminOnly)
	mustRegister(r, eventSendFSCLICommand, registeredHandler(eventSendFSCLICommand, runCLICommand, overrides), adminOnly)
	mustRegister(r, eventRealFSCLIConnect, registeredHandler(eventRealFSCLIConnect, RealFSCLIConnect, overrides), adminOnly)
	mustRegister(r, eventRealFSCLICommand, registeredHandler(eventRealFSCLICommand, RealFSCLICommand, overrides), adminOnly)
}

func registerCoreWebUserEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventSettingsUsersGet, registeredHandler(eventSettingsUsersGet, getWebUsers, overrides), adminOnly)
	mustRegister(r, eventGetWebUsersByDirectory, registeredHandler(eventGetWebUsersByDirectory, GetWebUsersByDirectory, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersAdd, registeredHandler(eventSettingsUsersAdd, addWebUsers, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersRename, registeredHandler(eventSettingsUsersRename, renameWebUsers, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersDelete, registeredHandler(eventSettingsUsersDelete, deleteWebUsers, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersSwitch, registeredHandler(eventSettingsUsersSwitch, switchWebUser, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdatePass, registeredHandler(eventSettingsUsersUpdatePass, updateWebUsersPassword, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateLang, registeredHandler(eventSettingsUsersUpdateLang, updateWebUsersLang, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateSip, registeredHandler(eventSettingsUsersUpdateSip, updateWebUsersSipUser, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateWS, registeredHandler(eventSettingsUsersUpdateWS, updateWebUsersWs, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateVerto, registeredHandler(eventSettingsUsersUpdateVerto, updateWebUsersVertoWs, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateRTC, registeredHandler(eventSettingsUsersUpdateRTC, UpdateWebUserWebRTCLib, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateStun, registeredHandler(eventSettingsUsersUpdateStun, updateWebUsersStun, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersUpdateAvatar, registeredHandler(eventSettingsUsersUpdateAvatar, updateWebUsersAvatar, overrides), adminOnly)
	mustRegister(r, eventSettingsUsersClearAvatar, registeredHandler(eventSettingsUsersClearAvatar, clearWebUsersAvatar, overrides), adminOnly)
}

func registerCoreDirectoryTemplateEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventUpdateWebUserGroup, UpdateWebUserGroup, overrides)
	mustRegisterAdmin(r, eventGetWebDirUserTemplates, GetWebDirectoryUsersTemplates, overrides)
	mustRegisterAdmin(r, eventAddWebDirUserTemplate, AddWebDirectoryUsersTemplate, overrides)
	mustRegisterAdmin(r, eventDelWebDirUserTemplate, DelWebDirectoryUsersTemplate, overrides)
	mustRegisterAdmin(r, eventUpdateWebDirUserTemplate, UpdateWebDirectoryUsersTemplate, overrides)
	mustRegisterAdmin(r, eventSwitchWebDirUserTemplate, SwitchWebDirectoryUsersTemplate, overrides)
	mustRegisterAdmin(r, eventGetWebDirUserTplParams, GetWebDirectoryUsersTemplateParameters, overrides)
	mustRegisterAdmin(r, eventAddWebDirUserTplParam, AddWebDirectoryUsersTemplateParameter, overrides)
	mustRegisterAdmin(r, eventDelWebDirUserTplParam, DelWebDirectoryUsersTemplateParameter, overrides)
	mustRegisterAdmin(r, eventSwitchWebDirUserTplParam, SwitchWebDirectoryUsersTemplateParameter, overrides)
	mustRegisterAdmin(r, eventUpdateWebDirUserTplParam, UpdateWebDirectoryUsersTemplateParameter, overrides)
	mustRegisterAdmin(r, eventGetWebDirUserTplVars, GetWebDirectoryUsersTemplateVariables, overrides)
	mustRegisterAdmin(r, eventAddWebDirUserTplVar, AddWebDirectoryUsersTemplateVariable, overrides)
	mustRegisterAdmin(r, eventDelWebDirUserTplVar, DelWebDirectoryUsersTemplateVariable, overrides)
	mustRegisterAdmin(r, eventSwitchWebDirUserTplVar, SwitchWebDirectoryUsersTemplateVariable, overrides)
	mustRegisterAdmin(r, eventUpdateWebDirUserTplVar, UpdateWebDirectoryUsersTemplateVariable, overrides)
	mustRegisterAdminOrManager(r, eventGetWebDirUserTplList, GetWebDirectoryUsersTemplatesList, overrides)
	mustRegisterAdminOrManager(r, eventGetWebDirUserTplForm, GetWebDirectoryUsersTemplateForm, overrides)
	mustRegisterAdminOrManager(r, eventCreateWebDirUserByTpl, CreateWebDirectoryUsersByTemplate, overrides)
}

func registerCoreDirectoryDomainEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventDirDomainsGet, getDirectoryDomains, overrides)
	mustRegisterAdmin(r, eventDirImport, importDirectory, overrides)
	mustRegisterAdmin(r, eventDirDomainImportXML, importDirectoryXMLDomain, overrides)
	mustRegisterAdmin(r, eventDirDomainAdd, addDirectoryDomain, overrides)
	mustRegisterAdmin(r, eventDirDomainRename, renameDirectoryDomain, overrides)
	mustRegisterAdmin(r, eventDirDomainSwitch, switchDirectoryDomain, overrides)
	mustRegisterAdmin(r, eventDirDomainDelete, deleteDirectoryDomain, overrides)
	mustRegisterAdmin(r, eventDirDomainDetails, getDirectoryDomainDetails, overrides)
	mustRegisterAdmin(r, eventDirDomainAddParam, addDirectoryDomainParameter, overrides)
	mustRegisterAdmin(r, eventDirDomainUpdateParam, updateDirectoryDomainParameter, overrides)
	mustRegisterAdmin(r, eventDirDomainSwitchParam, switchDirectoryDomainParameter, overrides)
	mustRegisterAdmin(r, eventDirDomainDeleteParam, deleteDirectoryDomainParameter, overrides)
	mustRegisterAdmin(r, eventDirDomainAddVar, addDirectoryDomainVariable, overrides)
	mustRegisterAdmin(r, eventDirDomainUpdateVar, updateDirectoryDomainVariable, overrides)
	mustRegisterAdmin(r, eventDirDomainSwitchVar, switchDirectoryDomainVariable, overrides)
	mustRegisterAdmin(r, eventDirDomainDeleteVar, deleteDirectoryDomainVariable, overrides)
	mustRegisterAdmin(r, webStruct.GetDirectoryUser, getDirectoryUsers, overrides)
	mustRegisterAdmin(r, eventDirUserDetails, getDirectoryUserDetails, overrides)
	mustRegisterAdmin(r, eventDirUserAddParam, addDirectoryUserParameter, overrides)
	mustRegisterAdmin(r, eventDirUserAddVar, addDirectoryUserVariable, overrides)
	mustRegisterAdmin(r, eventDirUserDeleteParam, deleteDirectoryUserParameter, overrides)
	mustRegisterAdmin(r, eventDirUserDeleteVar, deleteDirectoryUserVariable, overrides)
	mustRegisterAdmin(r, eventDirUserUpdateParam, updateDirectoryUserParameter, overrides)
	mustRegisterAdmin(r, eventDirUserUpdateVar, updateDirectoryUserVariable, overrides)
	mustRegisterAdmin(r, eventDirUserUpdateCache, updateDirectoryUserCache, overrides)
	mustRegisterAdmin(r, eventDirUserUpdateCidr, updateDirectoryUserCidr, overrides)
	mustRegisterAdmin(r, eventDirUserUpdateNumberAlias, updateDirectoryUserNumberAlias, overrides)
	mustRegisterAdmin(r, eventDirUserAdd, addNewUser, overrides)
	mustRegisterAdmin(r, eventDirUserImportXML, importDirectoryXMLUser, overrides)
	mustRegisterAdmin(r, eventDirUserDelete, deleteDirectoryUser, overrides)
	mustRegisterAdmin(r, eventDirUserUpdateName, updateDirectoryUserName, overrides)
	mustRegisterAdmin(r, eventDirUserSwitch, switchDirectoryUser, overrides)
	mustRegisterAdmin(r, eventDirUserSwitchParam, switchDirectoryUserParameter, overrides)
	mustRegisterAdmin(r, eventDirUserSwitchVar, switchDirectoryUserVariable, overrides)
	mustRegisterAdmin(r, eventDirGroupsGet, getDirectoryGroups, overrides)
	mustRegisterAdmin(r, eventDirGroupUsersGet, getDirectoryGroupUsers, overrides)
	mustRegisterAdmin(r, eventDirGroupAdd, addDirectoryGroup, overrides)
	mustRegisterAdmin(r, eventDirGroupDelete, deleteDirectoryGroup, overrides)
	mustRegisterAdmin(r, eventDirGroupUpdateName, updateDirectoryGroupName, overrides)
	mustRegisterAdmin(r, eventDirGroupUserAdd, addDirectoryGroupUser, overrides)
	mustRegisterAdmin(r, eventDirGroupUserDelete, deleteDirectoryGroupUser, overrides)
	mustRegisterAdmin(r, eventDirUserGatewaysGet, getDirectoryUserGateways, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayDetails, getDirectoryUserGatewayDetails, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayAddParam, addDirectoryUserGatewayParameter, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayDeleteParam, deleteDirectoryUserGatewayParameter, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayUpdateParam, updateDirectoryUserGatewayParameter, overrides)
	mustRegisterAdmin(r, eventDirUserGatewaySwitchParam, switchDirectoryUserGatewayParameter, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayAddVar, addDirectoryUserGatewayVariable, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayUpdateVar, updateDirectoryUserGatewayVariable, overrides)
	mustRegisterAdmin(r, eventDirUserGatewaySwitchVar, switchDirectoryUserGatewayVariable, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayDeleteVar, deleteDirectoryUserGatewayVariable, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayAdd, addDirectoryUserGateway, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayDelete, deleteDirectoryUserGateway, overrides)
	mustRegisterAdmin(r, eventDirUserGatewayUpdateName, updateDirectoryUserGatewayName, overrides)
}

func getDirectoryDomains(data *webStruct.MessageData) webStruct.UserResponse {
	data.Id = cache.GetCurrentInstanceId()
	return getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomain{}, adminOnly())
}

func importDirectoryXMLDomain(data *webStruct.MessageData) webStruct.UserResponse {
	resp := ImportXMLDomain(data)
	if resp.Error != "" {
		return resp
	}
	return getDirectoryDomains(data)
}

func addDirectoryDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomain{Name: data.Name, Enabled: true, Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}}, adminOnly())
}

func renameDirectoryDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomain{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func switchDirectoryDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomain{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteDirectoryDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomain{Id: data.Id}, adminOnly())
}

func getDirectoryDomainDetails(data *webStruct.MessageData) webStruct.UserResponse {
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainParameter{}, adminOnly())
	resp2 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainVariable{}, adminOnly())
	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		S   interface{} `json:"parameters"`
		Sch interface{} `json:"variables"`
	}{S: resp1.Data, Sch: resp2.Data}}
}

func addDirectoryDomainParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainParameter{Name: data.Name, Value: data.Value, Enabled: true, Parent: &altStruct.DirectoryDomain{Id: data.Id}}, adminOnly())
}

func updateDirectoryDomainParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainParameter{Id: data.Id, Name: data.Name, Value: data.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchDirectoryDomainParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainParameter{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteDirectoryDomainParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainParameter{Id: data.Id}, adminOnly())
}

func addDirectoryDomainVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainVariable{Name: data.Name, Value: data.Value, Enabled: true, Parent: &altStruct.DirectoryDomain{Id: data.Id}}, adminOnly())
}

func updateDirectoryDomainVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainVariable{Id: data.Id, Name: data.Name, Value: data.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchDirectoryDomainVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainVariable{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteDirectoryDomainVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainVariable{Id: data.Id}, adminOnly())
}

func getDirectoryUsers(data *webStruct.MessageData) webStruct.UserResponse {
	data.Id = cache.GetCurrentInstanceId()
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomain{}, adminOnly())
	domains, ok := resp1.Data.(map[int64]interface{})
	if !ok {
		return webStruct.UserResponse{Error: "domains not found", MessageType: data.Event}
	}
	var ids []int64
	for _, d := range domains {
		domain, ok := d.(altStruct.DirectoryDomain)
		if !ok || domain.Id == 0 {
			continue
		}
		ids = append(ids, domain.Id)
	}
	data.IntSlice = ids
	resp2 := getUserForConfig(data, getDirectoryByParents, &altStruct.DirectoryDomainUser{}, adminOnly())
	users, ok := resp2.Data.(map[int64]interface{})
	if !ok {
		return webStruct.UserResponse{Error: "users not found", MessageType: data.Event}
	}
	directoryCache := cache.GetDirectoryCache()
	for k, u := range users {
		user, ok := u.(altStruct.DirectoryDomainUser)
		if !ok || user.Id == 0 {
			continue
		}
		directoryCache.UserCache.ApplyToUser(&user)
		users[k] = user
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		S   interface{} `json:"domains"`
		Sch interface{} `json:"directory_users"`
	}{S: resp1.Data, Sch: users}}
}

func getDirectoryUserDetails(data *webStruct.MessageData) webStruct.UserResponse {
	userMsg := getUserForConfig(data, getDirectoryById, &altStruct.DirectoryDomainUser{}, adminOnly())
	user, ok := userMsg.Data.(altStruct.DirectoryDomainUser)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUserParameter{}, adminOnly())
	resp2 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUserVariable{}, adminOnly())
	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		A interface{} `json:"parameters"`
		B interface{} `json:"variables"`
		C interface{} `json:"user"`
	}{A: resp1.Data, B: resp2.Data, C: user}}
}

func addDirectoryUserParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainUserParameter{Name: data.Name, Value: data.Value, Enabled: true, Parent: &altStruct.DirectoryDomainUser{Id: data.Id}}, adminOnly())
}

func addDirectoryUserVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainUserVariable{Name: data.Name, Value: data.Value, Enabled: true, Parent: &altStruct.DirectoryDomainUser{Id: data.Id}}, adminOnly())
}

func deleteDirectoryUserParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainUserParameter{Id: data.Id}, adminOnly())
}

func deleteDirectoryUserVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainUserVariable{Id: data.Id}, adminOnly())
}

func updateDirectoryUserParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserParameter{Id: data.Id, Name: data.Name, Value: data.Value}, []string{"Name", "Value"}}, adminOnly())
}

func updateDirectoryUserVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserVariable{Id: data.Id, Name: data.Name, Value: data.Value}, []string{"Name", "Value"}}, adminOnly())
}

func updateDirectoryUserCache(data *webStruct.MessageData) webStruct.UserResponse {
	cacheValue, err := strconv.ParseUint(data.Value, 10, 32)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUser{Id: data.Id, Cache: uint(cacheValue)}, []string{"Cache"}}, adminOnly())
}

func updateDirectoryUserCidr(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUser{Id: data.Id, Cidr: data.Value}, []string{"Cidr"}}, adminOnly())
}

func updateDirectoryUserNumberAlias(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUser{Id: data.Id, NumberAlias: data.Value}, []string{"NumberAlias"}}, adminOnly())
}

func importDirectoryXMLUser(data *webStruct.MessageData) webStruct.UserResponse {
	resp := ImportXMLDomainUser(data)
	if resp.Error != "" {
		return resp
	}
	return getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUser{}, adminOnly())
}

func deleteDirectoryUser(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainUser{Id: data.Id}, adminOnly())
}

func updateDirectoryUserName(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUser{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func switchDirectoryUser(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUser{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func switchDirectoryUserParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserParameter{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func switchDirectoryUserVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserVariable{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func getDirectoryGroups(data *webStruct.MessageData) webStruct.UserResponse {
	data.Id = cache.GetCurrentInstanceId()
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomain{}, adminOnly())
	domains, ok := resp1.Data.(map[int64]interface{})
	if !ok {
		return webStruct.UserResponse{Error: "domains not found", MessageType: data.Event}
	}
	var ids []int64
	for _, d := range domains {
		domain, ok := d.(altStruct.DirectoryDomain)
		if !ok || domain.Id == 0 {
			continue
		}
		ids = append(ids, domain.Id)
	}
	data.IntSlice = ids
	resp2 := getUserForConfig(data, getDirectoryByParents, &altStruct.DirectoryDomainGroup{}, adminOnly())
	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		S   interface{} `json:"domains"`
		Sch interface{} `json:"list"`
	}{S: resp1.Data, Sch: resp2.Data}}
}

func getDirectoryGroupUsers(data *webStruct.MessageData) webStruct.UserResponse {
	resp0 := getUserForConfig(data, getDirectoryById, &altStruct.DirectoryDomainGroup{}, adminOnly())
	group, ok := resp0.Data.(altStruct.DirectoryDomainGroup)
	if !ok || group.Id == 0 || group.Parent.Id == 0 {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainGroupUser{}, adminOnly())
	data.Id = group.Parent.Id
	resp2 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUser{}, adminOnly())
	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		S   interface{} `json:"group_users"`
		Sch interface{} `json:"users"`
	}{S: resp1.Data, Sch: resp2.Data}}
}

func addDirectoryGroup(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainGroup{Name: data.Name, Enabled: true, Parent: &altStruct.DirectoryDomain{Id: data.Id}}, adminOnly())
}

func deleteDirectoryGroup(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainGroup{Id: data.Id}, adminOnly())
}

func updateDirectoryGroupName(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainGroup{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func addDirectoryGroupUser(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainGroupUser{UserId: &altStruct.DirectoryDomainUser{Id: data.IdInt}, Enabled: true, Parent: &altStruct.DirectoryDomainGroup{Id: data.Id}}, adminOnly())
}

func deleteDirectoryGroupUser(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainGroupUser{Id: data.Id}, adminOnly())
}

func getDirectoryUserGateways(data *webStruct.MessageData) webStruct.UserResponse {
	data.Id = cache.GetCurrentInstanceId()
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomain{}, adminOnly())
	data.Id = 0
	domains, ok := resp1.Data.(map[int64]interface{})
	if !ok {
		return webStruct.UserResponse{Error: "domains not found", MessageType: data.Event}
	}
	for _, d := range domains {
		domain, ok := d.(altStruct.DirectoryDomain)
		if !ok || domain.Id == 0 {
			continue
		}
		data.Id = domain.Id
		break
	}
	resp2 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUser{}, adminOnly())
	users, ok := resp2.Data.(map[int64]interface{})
	if !ok {
		return webStruct.UserResponse{Error: "users not found", MessageType: data.Event}
	}
	var ids []int64
	for _, u := range users {
		user, ok := u.(altStruct.DirectoryDomainUser)
		if !ok || user.Id == 0 {
			continue
		}
		ids = append(ids, user.Id)
	}
	data.IntSlice = ids
	resp3 := getUserForConfig(data, getDirectoryByParents, &altStruct.DirectoryDomainUserGateway{}, adminOnly())
	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		S   interface{} `json:"domains"`
		Sch interface{} `json:"directory_users"`
		U   interface{} `json:"user_gateways"`
	}{S: resp1.Data, Sch: resp2.Data, U: resp3.Data}}
}

func getDirectoryUserGatewayDetails(data *webStruct.MessageData) webStruct.UserResponse {
	resp1 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUserGatewayParameter{}, adminOnly())
	resp2 := getUserForConfig(data, getDirectoryByParent, &altStruct.DirectoryDomainUserGatewayVariable{}, adminOnly())
	return webStruct.UserResponse{MessageType: data.Event, Data: struct {
		S   interface{} `json:"parameters"`
		Sch interface{} `json:"variables"`
	}{S: resp1.Data, Sch: resp2.Data}}
}

func addDirectoryUserGatewayParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainUserGatewayParameter{Name: data.Name, Value: data.Value, Enabled: true, Parent: &altStruct.DirectoryDomainUserGateway{Id: data.Id}}, adminOnly())
}

func deleteDirectoryUserGatewayParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainUserGatewayParameter{Id: data.Id}, adminOnly())
}

func updateDirectoryUserGatewayParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserGatewayParameter{Id: data.Id, Name: data.Name, Value: data.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchDirectoryUserGatewayParameter(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserGatewayParameter{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func addDirectoryUserGatewayVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainUserGatewayVariable{Name: data.Name, Value: data.Value, Direction: data.Direction, Enabled: true, Parent: &altStruct.DirectoryDomainUserGateway{Id: data.Id}}, adminOnly())
}

func updateDirectoryUserGatewayVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserGatewayVariable{Id: data.Id, Name: data.Name, Value: data.Value, Direction: data.Direction}, []string{"Name", "Value", "Direction"}}, adminOnly())
}

func switchDirectoryUserGatewayVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserGatewayVariable{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteDirectoryUserGatewayVariable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainUserGatewayVariable{Id: data.Id}, adminOnly())
}

func addDirectoryUserGateway(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.DirectoryDomainUserGateway{Name: data.Name, Enabled: true, Parent: &altStruct.DirectoryDomainUser{Id: data.Id}}, adminOnly())
}

func deleteDirectoryUserGateway(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.DirectoryDomainUserGateway{Id: data.Id}, adminOnly())
}

func updateDirectoryUserGatewayName(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.DirectoryDomainUserGateway{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func registerCoreModuleEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, webStruct.GetModules, GetConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleReload, reloadConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleUnload, unloadConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleLoad, loadConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleSwitch, switchConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleFromScratch, fromScratchConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleImport, importConfModules, overrides)
	mustRegisterAdmin(r, eventConfigModulesImportAll, importConfAllModules, overrides)
	mustRegisterAdmin(r, eventConfigModuleTruncate, TruncateModuleConfig, overrides)
	mustRegisterAdmin(r, eventConfigModuleImportXML, ImportXMLModuleConfig, overrides)
	mustRegisterAdmin(r, eventConfigModuleAutoload, autoloadModule, overrides)
}

func autoloadModule(data *webStruct.MessageData) webStruct.UserResponse {
	res, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigurationsList{Id: data.Id})
	if err != nil || res == nil {
		return webStruct.UserResponse{Error: "module not found", MessageType: data.Event}
	}
	module, ok := res.(altStruct.ConfigurationsList)
	if !ok {
		return webStruct.UserResponse{Error: "module not found", MessageType: data.Event}
	}
	module.Module = mainStruct.GetModuleNameByConfName(module.Name)
	result, err := intermediateDB.GetByValue(&altStruct.ConfigPostLoadModule{Name: module.Module, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadModule{}))}, map[string]bool{"Parent": true, "Name": true})
	if err != nil || module.Module == "" {
		return webStruct.UserResponse{Error: "module not found", MessageType: data.Event}
	}
	if len(result) == 0 {
		return getUserForConfig(data, setConfig, &altStruct.ConfigPostLoadModule{Name: module.Module, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadModule{}))}, adminOnly())
	}
	postloadMod, ok := result[0].(altStruct.ConfigPostLoadModule)
	if !ok {
		return webStruct.UserResponse{Error: "module not found", MessageType: data.Event}
	}
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigPostLoadModule{Id: postloadMod.Id, Enabled: !postloadMod.Enabled}, []string{"Enabled"}}, adminOnly())
}

func registerCoreConversationEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventGetConvPrivateMessages, registeredHandler(eventGetConvPrivateMessages, GetConversationPrivateMessages, overrides), adminOnly)
	mustRegister(r, eventGetConvPrivateCalls, registeredHandler(eventGetConvPrivateCalls, GetConversationPrivateCalls, overrides), adminOnly)
	mustRegister(r, eventGetConvRoomMessages, registeredHandler(eventGetConvRoomMessages, GetConversationRoomMessages, overrides), adminOnly)
	mustRegister(r, eventSendConvPrivateMessage, registeredHandler(eventSendConvPrivateMessage, SendConversationPrivateMessage, overrides), adminOnly)
	mustRegister(r, eventSendConvPrivateCall, registeredHandler(eventSendConvPrivateCall, SendConversationPrivateCall, overrides), adminOnly)
	mustRegister(r, eventSendConvRoomMessage, registeredHandler(eventSendConvRoomMessage, SendConversationRoomMessage, overrides), adminOnly)
}

func registerCoreDialplanEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventSwitchDialplanNoProceed, registeredHandler(eventSwitchDialplanNoProceed, SwitchDialplanNoProceed, overrides), adminOnly)
	mustRegister(r, eventDialplanGetSettings, registeredHandler(eventDialplanGetSettings, DialplanGetSettings, overrides), adminOnly)
	mustRegister(r, eventDialplanGetContexts, registeredHandler(eventDialplanGetContexts, getDialplanContexts, overrides), adminOnly)
	mustRegister(r, eventDialplanImport, registeredHandler(eventDialplanImport, importDialplan, overrides), adminOnly)
	mustRegister(r, eventDialplanGetExtensions, registeredHandler(eventDialplanGetExtensions, getDialplanExtensions, overrides), adminOnly)
	mustRegister(r, eventDialplanGetConditions, registeredHandler(eventDialplanGetConditions, getDialplanConditions, overrides), adminOnly)
	mustRegister(r, eventDialplanGetExtDetails, registeredHandler(eventDialplanGetExtDetails, getDialplanExtenDetails, overrides), adminOnly)
	mustRegister(r, eventDialplanMoveExtension, registeredHandler(eventDialplanMoveExtension, moveDialplanExtension, overrides), adminOnly)
	mustRegister(r, eventDialplanMoveCondition, registeredHandler(eventDialplanMoveCondition, moveDialplanCondition, overrides), adminOnly)
	mustRegister(r, eventDialplanMoveAction, registeredHandler(eventDialplanMoveAction, moveDialplanAction, overrides), adminOnly)
	mustRegister(r, eventDialplanMoveAntiaction, registeredHandler(eventDialplanMoveAntiaction, moveDialplanAntiAction, overrides), adminOnly)
	mustRegister(r, eventDialplanAddRegex, registeredHandler(eventDialplanAddRegex, addRegex, overrides), adminOnly)
	mustRegister(r, eventDialplanAddAction, registeredHandler(eventDialplanAddAction, addAction, overrides), adminOnly)
	mustRegister(r, eventDialplanAddAntiaction, registeredHandler(eventDialplanAddAntiaction, addAntiAction, overrides), adminOnly)
	mustRegister(r, eventDialplanDeleteRegex, registeredHandler(eventDialplanDeleteRegex, delRegex, overrides), adminOnly)
	mustRegister(r, eventDialplanDeleteAction, registeredHandler(eventDialplanDeleteAction, delAction, overrides), adminOnly)
	mustRegister(r, eventDialplanDeleteAntiaction, registeredHandler(eventDialplanDeleteAntiaction, delAntiAction, overrides), adminOnly)
	mustRegister(r, eventDialplanUpdateRegex, registeredHandler(eventDialplanUpdateRegex, updateRegex, overrides), adminOnly)
	mustRegister(r, eventDialplanUpdateAction, registeredHandler(eventDialplanUpdateAction, updateAction, overrides), adminOnly)
	mustRegister(r, eventDialplanUpdateAntiaction, registeredHandler(eventDialplanUpdateAntiaction, updateAntiAction, overrides), adminOnly)
	mustRegister(r, eventDialplanSwitchRegex, registeredHandler(eventDialplanSwitchRegex, switchRegex, overrides), adminOnly)
	mustRegister(r, eventDialplanSwitchAction, registeredHandler(eventDialplanSwitchAction, switchAction, overrides), adminOnly)
	mustRegister(r, eventDialplanSwitchAntiaction, registeredHandler(eventDialplanSwitchAntiaction, switchAntiAction, overrides), adminOnly)
	mustRegister(r, eventDialplanAddContext, registeredHandler(eventDialplanAddContext, addContext, overrides), adminOnly)
	mustRegister(r, eventDialplanAddExtension, registeredHandler(eventDialplanAddExtension, addExtension, overrides), adminOnly)
	mustRegister(r, eventDialplanAddCondition, registeredHandler(eventDialplanAddCondition, addCondition, overrides), adminOnly)
	mustRegister(r, eventDialplanRenameContext, registeredHandler(eventDialplanRenameContext, renameContext, overrides), adminOnly)
	mustRegister(r, eventDialplanRenameExtension, registeredHandler(eventDialplanRenameExtension, renameExtension, overrides), adminOnly)
	mustRegister(r, eventDialplanDeleteContext, registeredHandler(eventDialplanDeleteContext, deleteContext, overrides), adminOnly)
	mustRegister(r, eventDialplanDeleteExtension, registeredHandler(eventDialplanDeleteExtension, deleteExtension, overrides), adminOnly)
	mustRegister(r, eventDialplanSwitchExtContinue, registeredHandler(eventDialplanSwitchExtContinue, switchExtensionContinue, overrides), adminOnly)
	mustRegister(r, eventDialplanUpdateCondition, registeredHandler(eventDialplanUpdateCondition, updateCondition, overrides), adminOnly)
	mustRegister(r, eventDialplanSwitchCondition, registeredHandler(eventDialplanSwitchCondition, switchCondition, overrides), adminOnly)
	mustRegister(r, eventDialplanDeleteCondition, registeredHandler(eventDialplanDeleteCondition, deleteCondition, overrides), adminOnly)
}

func registerCoreAuthTokenEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, eventRelogin, registeredHandler(eventRelogin, checkRelogin, overrides), adminOnly)
	mustRegister(r, eventLogOut, registeredHandler(eventLogOut, logoutAndClearSubscriptions, overrides), adminOnly)
	mustRegister(r, "AddUserToken", registeredHandler("AddUserToken", createAPIToken, overrides), adminOnly)
	mustRegister(r, "GetUserTokens", registeredHandler("GetUserTokens", GetUserTokens, overrides), adminOnly)
	mustRegister(r, "UserGetOwnTokens", registeredHandler("UserGetOwnTokens", UserGetOwnTokens, overrides), adminOnly)
	mustRegister(r, "RemoveUserToken", registeredHandler("RemoveUserToken", RemoveUserToken, overrides), adminOnly)
}

func registerCoreDebugEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegister(r, webStruct.DialplanDebug, registeredHandler(webStruct.DialplanDebug, getDialplanDebug, overrides), adminOnly)
	mustRegister(r, webStruct.SubscribeHepPackages, registeredHandler(webStruct.SubscribeHepPackages, getDialplanDebug, overrides), adminOnly)
	mustRegister(r, eventSwitchDialplanDebug, registeredHandler(eventSwitchDialplanDebug, switchDialplanDebug, overrides), adminOnly)
}

func registerCoreSubscriptionEvents(r *handlerRegistry) {
	mustRegisterContext(r, eventSubscriptionList, replaceSubscriptions, adminOnly)
	mustRegisterContext(r, eventPersistentSubscription, addPersistentSubscriptions, adminOnly)
	mustRegisterContext(r, webStruct.Unsubscribe, unsubscribe, adminOnly)
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
