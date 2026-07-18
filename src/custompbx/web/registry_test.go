package web

import (
	"custompbx/cfg"
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

func TestCoreRegistryEventInventoryHasNoDuplicates(t *testing.T) {
	names := coreEvents.EventNames()
	seen := map[string]bool{}
	for _, name := range names {
		if seen[name] {
			t.Fatalf("duplicate registered event %q", name)
		}
		seen[name] = true
	}
	if len(names) == 0 {
		t.Fatal("core registry is empty")
	}
}

func TestCoreRegistryIncludesMigratedEvents(t *testing.T) {
	events := []string{
		eventGetSettings,
		eventSetSettings,
		webStruct.GetDashboard,
		eventGetInstances,
		eventUpdateInstanceDescription,
		eventGetWebSettings,
		eventSaveWebSettings,
		eventGetCDR,
		eventGetHEP,
		eventGetHEPDetails,
		eventGetLogs,
		eventGetPhoneCreds,
		eventSendFSCLICommand,
		eventRealFSCLIConnect,
		eventRealFSCLICommand,
		eventSettingsUsersGet,
		eventGetWebUsersByDirectory,
		eventSettingsUsersAdd,
		eventSettingsUsersRename,
		eventSettingsUsersDelete,
		eventSettingsUsersSwitch,
		eventSettingsUsersUpdatePass,
		eventSettingsUsersUpdateLang,
		eventSettingsUsersUpdateSip,
		eventSettingsUsersUpdateWS,
		eventSettingsUsersUpdateVerto,
		eventSettingsUsersUpdateRTC,
		eventSettingsUsersUpdateStun,
		eventSettingsUsersUpdateAvatar,
		eventSettingsUsersClearAvatar,
		eventUpdateWebUserGroup,
		eventGetWebDirUserTemplates,
		eventAddWebDirUserTemplate,
		eventDelWebDirUserTemplate,
		eventUpdateWebDirUserTemplate,
		eventSwitchWebDirUserTemplate,
		eventGetWebDirUserTplParams,
		eventAddWebDirUserTplParam,
		eventDelWebDirUserTplParam,
		eventSwitchWebDirUserTplParam,
		eventUpdateWebDirUserTplParam,
		eventGetWebDirUserTplVars,
		eventAddWebDirUserTplVar,
		eventDelWebDirUserTplVar,
		eventSwitchWebDirUserTplVar,
		eventUpdateWebDirUserTplVar,
		eventGetWebDirUserTplList,
		eventGetWebDirUserTplForm,
		eventCreateWebDirUserByTpl,
		eventDirDomainsGet,
		eventDirImport,
		eventDirDomainImportXML,
		eventDirDomainAdd,
		eventDirDomainRename,
		eventDirDomainSwitch,
		eventDirDomainDelete,
		eventDirDomainDetails,
		eventDirDomainAddParam,
		eventDirDomainUpdateParam,
		eventDirDomainSwitchParam,
		eventDirDomainDeleteParam,
		eventDirDomainAddVar,
		eventDirDomainUpdateVar,
		eventDirDomainSwitchVar,
		eventDirDomainDeleteVar,
		webStruct.GetDirectoryUser,
		eventDirUserDetails,
		eventDirUserAddParam,
		eventDirUserAddVar,
		eventDirUserDeleteParam,
		eventDirUserDeleteVar,
		eventDirUserUpdateParam,
		eventDirUserUpdateVar,
		eventDirUserUpdateCache,
		eventDirUserUpdateCidr,
		eventDirUserUpdateNumberAlias,
		eventDirUserAdd,
		eventDirUserImportXML,
		eventDirUserDelete,
		eventDirUserUpdateName,
		eventDirUserSwitch,
		eventDirUserSwitchParam,
		eventDirUserSwitchVar,
		eventDirGroupsGet,
		eventDirGroupUsersGet,
		eventDirGroupAdd,
		eventDirGroupDelete,
		eventDirGroupUpdateName,
		eventDirGroupUserAdd,
		eventDirGroupUserDelete,
		eventDirUserGatewaysGet,
		eventDirUserGatewayDetails,
		eventDirUserGatewayAddParam,
		eventDirUserGatewayDeleteParam,
		eventDirUserGatewayUpdateParam,
		eventDirUserGatewaySwitchParam,
		eventDirUserGatewayAddVar,
		eventDirUserGatewayUpdateVar,
		eventDirUserGatewaySwitchVar,
		eventDirUserGatewayDeleteVar,
		eventDirUserGatewayAdd,
		eventDirUserGatewayDelete,
		eventDirUserGatewayUpdateName,
		webStruct.GetModules,
		eventConfigModuleReload,
		eventConfigModuleUnload,
		eventConfigModuleLoad,
		eventConfigModuleSwitch,
		eventConfigModuleFromScratch,
		eventConfigModuleImport,
		eventConfigModulesImportAll,
		eventConfigModuleTruncate,
		eventConfigModuleImportXML,
		eventConfigModuleAutoload,
		eventGetConvPrivateMessages,
		eventGetConvPrivateCalls,
		eventGetConvRoomMessages,
		eventSendConvPrivateMessage,
		eventSendConvPrivateCall,
		eventSendConvRoomMessage,
		eventSwitchDialplanNoProceed,
		eventDialplanGetSettings,
		eventDialplanGetContexts,
		eventDialplanImport,
		eventDialplanGetExtensions,
		eventDialplanGetConditions,
		eventDialplanGetExtDetails,
		eventDialplanMoveExtension,
		eventDialplanMoveCondition,
		eventDialplanMoveAction,
		eventDialplanMoveAntiaction,
		eventDialplanAddRegex,
		eventDialplanAddAction,
		eventDialplanAddAntiaction,
		eventDialplanDeleteRegex,
		eventDialplanDeleteAction,
		eventDialplanDeleteAntiaction,
		eventDialplanUpdateRegex,
		eventDialplanUpdateAction,
		eventDialplanUpdateAntiaction,
		eventDialplanSwitchRegex,
		eventDialplanSwitchAction,
		eventDialplanSwitchAntiaction,
		eventDialplanAddContext,
		eventDialplanAddExtension,
		eventDialplanAddCondition,
		eventDialplanRenameContext,
		eventDialplanRenameExtension,
		eventDialplanDeleteContext,
		eventDialplanDeleteExtension,
		eventDialplanSwitchExtContinue,
		eventDialplanUpdateCondition,
		eventDialplanSwitchCondition,
		eventDialplanDeleteCondition,
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

func TestCoreRegistryIncludesWebUserSettingsFamily(t *testing.T) {
	adminEvents := []string{
		eventSettingsUsersGet,
		eventGetWebUsersByDirectory,
		eventSettingsUsersAdd,
		eventSettingsUsersRename,
		eventSettingsUsersDelete,
		eventSettingsUsersSwitch,
		eventSettingsUsersUpdateLang,
		eventSettingsUsersUpdateSip,
	}
	assertAdminOnlyEventsDispatch(t, adminEvents)

	selfServiceEvents := []string{
		eventSettingsUsersUpdatePass,
		eventSettingsUsersUpdateWS,
		eventSettingsUsersUpdateVerto,
		eventSettingsUsersUpdateRTC,
		eventSettingsUsersUpdateStun,
		eventSettingsUsersUpdateAvatar,
		eventSettingsUsersClearAvatar,
	}
	for _, event := range selfServiceEvents {
		calls := 0
		r := buildCoreEvents(map[string]eventHandler{
			event: func(data *webStruct.MessageData) webStruct.UserResponse {
				calls++
				return webStruct.UserResponse{MessageType: data.Event}
			},
		})
		ctx := userContext()
		response, ok := r.Dispatch(messageData(ctx, event), ctx)
		if !ok || calls != 1 || response.MessageType != event {
			t.Fatalf("%s was not available to an authenticated user: ok=%v calls=%d response=%+v", event, ok, calls, response)
		}
	}
}

func TestCoreRegistryHandlerOverrides(t *testing.T) {
	calls := 0
	r := buildCoreEvents(map[string]eventHandler{
		eventUpdateWebUserGroup: func(data *webStruct.MessageData) webStruct.UserResponse {
			calls++
			return webStruct.UserResponse{MessageType: data.Event}
		},
	})
	ctx := adminContext()
	data := messageData(ctx, eventUpdateWebUserGroup)

	resp, ok := r.Dispatch(data, ctx)

	if !ok {
		t.Fatal("event was not dispatched")
	}
	if calls != 1 {
		t.Fatalf("handler calls = %d, want 1", calls)
	}
	if resp.MessageType != eventUpdateWebUserGroup {
		t.Fatalf("message type = %q, want %q", resp.MessageType, eventUpdateWebUserGroup)
	}
}

func TestCoreRegistryOverrideKeepsAccessCheck(t *testing.T) {
	r := buildCoreEvents(map[string]eventHandler{
		eventGetWebDirUserTemplates: func(data *webStruct.MessageData) webStruct.UserResponse {
			t.Fatal("handler should not run for unauthorized group")
			return webStruct.UserResponse{}
		},
	})
	ctx := userContext()
	data := messageData(ctx, eventGetWebDirUserTemplates)

	resp, ok := r.Dispatch(data, ctx)

	if !ok {
		t.Fatal("event was not dispatched")
	}
	if resp.MessageType != "no_access" {
		t.Fatalf("message type = %q, want no_access", resp.MessageType)
	}
}

func TestCoreRegistryIncludesWebDirectoryTemplateFamily(t *testing.T) {
	events := []string{
		eventGetWebDirUserTemplates,
		eventAddWebDirUserTemplate,
		eventDelWebDirUserTemplate,
		eventUpdateWebDirUserTemplate,
		eventSwitchWebDirUserTemplate,
		eventGetWebDirUserTplParams,
		eventAddWebDirUserTplParam,
		eventDelWebDirUserTplParam,
		eventSwitchWebDirUserTplParam,
		eventUpdateWebDirUserTplParam,
		eventGetWebDirUserTplVars,
		eventAddWebDirUserTplVar,
		eventDelWebDirUserTplVar,
		eventSwitchWebDirUserTplVar,
		eventUpdateWebDirUserTplVar,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryManagerDirectoryTemplateEvents(t *testing.T) {
	events := []string{
		eventGetWebDirUserTplList,
		eventGetWebDirUserTplForm,
		eventCreateWebDirUserByTpl,
	}
	for _, event := range events {
		event := event
		calls := 0
		r := buildCoreEvents(map[string]eventHandler{
			event: func(data *webStruct.MessageData) webStruct.UserResponse {
				calls++
				return webStruct.UserResponse{MessageType: data.Event}
			},
		})
		ctx := managerContext()
		data := messageData(ctx, event)

		resp, ok := r.Dispatch(data, ctx)

		if !ok {
			t.Fatalf("%s event was not dispatched", event)
		}
		if calls != 1 {
			t.Fatalf("%s handler calls = %d, want 1", event, calls)
		}
		if resp.MessageType != event {
			t.Fatalf("%s response message type = %q", event, resp.MessageType)
		}
	}
}

func TestCoreRegistryManagerDirectoryTemplateEventsRejectUsers(t *testing.T) {
	events := []string{
		eventGetWebDirUserTplList,
		eventGetWebDirUserTplForm,
		eventCreateWebDirUserByTpl,
	}
	for _, event := range events {
		event := event
		r := buildCoreEvents(map[string]eventHandler{
			event: func(data *webStruct.MessageData) webStruct.UserResponse {
				t.Fatalf("%s handler should not run for unauthorized group", event)
				return webStruct.UserResponse{}
			},
		})
		ctx := userContext()
		data := messageData(ctx, event)

		resp, ok := r.Dispatch(data, ctx)

		if !ok {
			t.Fatalf("%s event was not dispatched", event)
		}
		if resp.MessageType != "no_access" {
			t.Fatalf("%s message type = %q, want no_access", event, resp.MessageType)
		}
	}
}

func TestCoreRegistryIncludesDirectoryDomainConfigFamily(t *testing.T) {
	events := []string{
		eventDirDomainsGet,
		eventDirImport,
		eventDirDomainImportXML,
		eventDirDomainAdd,
		eventDirDomainRename,
		eventDirDomainSwitch,
		eventDirDomainDelete,
		eventDirDomainDetails,
		eventDirDomainAddParam,
		eventDirDomainUpdateParam,
		eventDirDomainSwitchParam,
		eventDirDomainDeleteParam,
		eventDirDomainAddVar,
		eventDirDomainUpdateVar,
		eventDirDomainSwitchVar,
		eventDirDomainDeleteVar,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesDirectoryUserFamily(t *testing.T) {
	events := []string{
		webStruct.GetDirectoryUser,
		eventDirUserDetails,
		eventDirUserAddParam,
		eventDirUserAddVar,
		eventDirUserDeleteParam,
		eventDirUserDeleteVar,
		eventDirUserUpdateParam,
		eventDirUserUpdateVar,
		eventDirUserUpdateCache,
		eventDirUserUpdateCidr,
		eventDirUserUpdateNumberAlias,
		eventDirUserAdd,
		eventDirUserImportXML,
		eventDirUserDelete,
		eventDirUserUpdateName,
		eventDirUserSwitch,
		eventDirUserSwitchParam,
		eventDirUserSwitchVar,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesDirectoryGroupFamily(t *testing.T) {
	events := []string{
		eventDirGroupsGet,
		eventDirGroupUsersGet,
		eventDirGroupAdd,
		eventDirGroupDelete,
		eventDirGroupUpdateName,
		eventDirGroupUserAdd,
		eventDirGroupUserDelete,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesDirectoryUserGatewayFamily(t *testing.T) {
	events := []string{
		eventDirUserGatewaysGet,
		eventDirUserGatewayDetails,
		eventDirUserGatewayAddParam,
		eventDirUserGatewayDeleteParam,
		eventDirUserGatewayUpdateParam,
		eventDirUserGatewaySwitchParam,
		eventDirUserGatewayAddVar,
		eventDirUserGatewayUpdateVar,
		eventDirUserGatewaySwitchVar,
		eventDirUserGatewayDeleteVar,
		eventDirUserGatewayAdd,
		eventDirUserGatewayDelete,
		eventDirUserGatewayUpdateName,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesModuleFamily(t *testing.T) {
	events := []string{
		webStruct.GetModules,
		eventConfigModuleReload,
		eventConfigModuleUnload,
		eventConfigModuleLoad,
		eventConfigModuleSwitch,
		eventConfigModuleFromScratch,
		eventConfigModuleImport,
		eventConfigModulesImportAll,
		eventConfigModuleTruncate,
		eventConfigModuleImportXML,
		eventConfigModuleAutoload,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesACLFamily(t *testing.T) {
	events := []string{
		eventACLListsGet,
		eventACLListAdd,
		eventACLListUpdateDefault,
		eventACLListUpdate,
		eventACLListDelete,
		eventACLListConfigUpdateDefault,
		eventACLNodesGet,
		eventACLNodeAdd,
		eventACLNodeDelete,
		eventACLNodeUpdate,
		eventACLNodeSwitch,
		eventACLNodeMove,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesSofiaGlobalFamily(t *testing.T) {
	events := []string{
		eventSofiaGlobalSettingsGet,
		eventSofiaGlobalSettingUpdate,
		eventSofiaGlobalSettingSwitch,
		eventSofiaGlobalSettingAdd,
		eventSofiaGlobalSettingDelete,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesSofiaProfileFamily(t *testing.T) {
	events := []string{
		webStruct.GetSofiaProfiles,
		eventSofiaProfileParamsGet,
		eventSofiaProfileParamAdd,
		eventSofiaProfileParamDelete,
		eventSofiaProfileParamSwitch,
		eventSofiaProfileParamUpdate,
		eventSofiaProfileDomainsGet,
		eventSofiaProfileDomainAdd,
		eventSofiaProfileDomainDelete,
		eventSofiaProfileDomainSwitch,
		eventSofiaProfileDomainUpdate,
		eventSofiaProfileAliasesGet,
		eventSofiaProfileAliasAdd,
		eventSofiaProfileAliasDelete,
		eventSofiaProfileAliasSwitch,
		eventSofiaProfileAliasUpdate,
		eventSofiaProfileAdd,
		eventSofiaProfileRename,
		eventSofiaProfileDelete,
		eventSofiaProfileCommand,
		eventSofiaProfileSwitch,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesSofiaGatewayFamily(t *testing.T) {
	events := []string{
		eventSofiaProfileGatewaysGet,
		eventSofiaGatewayVarsGet,
		eventSofiaGatewayParamsGet,
		eventSofiaGatewayParamAdd,
		eventSofiaGatewayParamUpdate,
		eventSofiaGatewayParamSwitch,
		eventSofiaGatewayParamDelete,
		eventSofiaGatewayVarAdd,
		eventSofiaGatewayVarUpdate,
		eventSofiaGatewayVarSwitch,
		eventSofiaGatewayVarDelete,
		eventSofiaGatewayAdd,
		eventSofiaGatewayDelete,
		eventSofiaGatewayRename,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesCDRConfigFamily(t *testing.T) {
	events := []string{
		eventCdrPgCsvGet,
		eventCdrPgCsvParamAdd,
		eventCdrPgCsvParamUpdate,
		eventCdrPgCsvParamSwitch,
		eventCdrPgCsvParamDelete,
		eventCdrPgCsvFieldAdd,
		eventCdrPgCsvFieldUpdate,
		eventCdrPgCsvFieldSwitch,
		eventCdrPgCsvFieldDelete,
		eventOdbcCdrGet,
		eventOdbcCdrFieldGet,
		eventOdbcCdrParamAdd,
		eventOdbcCdrParamUpdate,
		eventOdbcCdrParamSwitch,
		eventOdbcCdrParamDelete,
		eventOdbcCdrTableAdd,
		eventOdbcCdrTableUpdate,
		eventOdbcCdrTableSwitch,
		eventOdbcCdrTableDelete,
		eventOdbcCdrFieldAdd,
		eventOdbcCdrFieldUpdate,
		eventOdbcCdrFieldSwitch,
		eventOdbcCdrFieldDelete,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesLCRFamily(t *testing.T) {
	events := []string{
		eventLCRGet,
		eventLCRProfileParamsGet,
		eventLCRParamUpdate,
		eventLCRParamSwitch,
		eventLCRParamAdd,
		eventLCRParamDelete,
		eventLCRProfileParamAdd,
		eventLCRProfileParamDelete,
		eventLCRProfileParamSwitch,
		eventLCRProfileParamUpdate,
		eventLCRProfileAdd,
		eventLCRProfileUpdate,
		eventLCRProfileDelete,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesSimpleModuleSettingFamilies(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, simpleModuleSettingEvents())
}

func TestCoreRegistryIncludesPostSwitchFamily(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, postSwitchRegistryEvents())
}

func TestCoreRegistryIncludesDirectoryConfigFamily(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, directoryConfigRegistryEvents())
}

func TestCoreRegistryIncludesFifoFamily(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, fifoRegistryEvents())
}

func TestCoreRegistryIncludesTelephonyModuleFamilies(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, telephonyModuleRegistryEvents())
}

func TestCoreRegistryIncludesConferenceFamily(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, conferenceRegistryEvents())
}

func TestCoreRegistryIncludesBackendVertoConfigFamily(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, vertoConfigRegistryEvents())
}

func TestCoreRegistryIncludesRemainingConfigFamilies(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, remainingConfigRegistryEvents())
}

func TestCoreRegistryIncludesConversationFamily(t *testing.T) {
	events := []string{
		eventGetConvPrivateMessages,
		eventGetConvPrivateCalls,
		eventGetConvRoomMessages,
		eventSendConvPrivateMessage,
		eventSendConvPrivateCall,
		eventSendConvRoomMessage,
	}
	assertAdminOnlyEventsDispatch(t, events)
}

func TestCoreRegistryIncludesDialplanFamily(t *testing.T) {
	assertAdminOnlyEventsDispatch(t, dialplanRegistryEvents())
}

func dialplanRegistryEvents() []string {
	return []string{
		eventSwitchDialplanNoProceed,
		eventDialplanGetSettings,
		eventDialplanGetContexts,
		eventDialplanImport,
		eventDialplanGetExtensions,
		eventDialplanGetConditions,
		eventDialplanGetExtDetails,
		eventDialplanMoveExtension,
		eventDialplanMoveCondition,
		eventDialplanMoveAction,
		eventDialplanMoveAntiaction,
		eventDialplanAddRegex,
		eventDialplanAddAction,
		eventDialplanAddAntiaction,
		eventDialplanDeleteRegex,
		eventDialplanDeleteAction,
		eventDialplanDeleteAntiaction,
		eventDialplanUpdateRegex,
		eventDialplanUpdateAction,
		eventDialplanUpdateAntiaction,
		eventDialplanSwitchRegex,
		eventDialplanSwitchAction,
		eventDialplanSwitchAntiaction,
		eventDialplanAddContext,
		eventDialplanAddExtension,
		eventDialplanAddCondition,
		eventDialplanRenameContext,
		eventDialplanRenameExtension,
		eventDialplanDeleteContext,
		eventDialplanDeleteExtension,
		eventDialplanSwitchExtContinue,
		eventDialplanUpdateCondition,
		eventDialplanSwitchCondition,
		eventDialplanDeleteCondition,
	}
}

func assertAdminOnlyEventsDispatch(t *testing.T, events []string) {
	t.Helper()
	for _, event := range events {
		event := event
		calls := 0
		r := buildCoreEvents(map[string]eventHandler{
			event: func(data *webStruct.MessageData) webStruct.UserResponse {
				calls++
				return webStruct.UserResponse{MessageType: data.Event}
			},
		})
		ctx := adminContext()
		data := messageData(ctx, event)

		resp, ok := r.Dispatch(data, ctx)

		if !ok {
			t.Fatalf("%s event was not dispatched", event)
		}
		if calls != 1 {
			t.Fatalf("%s handler calls = %d, want 1", event, calls)
		}
		if resp.MessageType != event {
			t.Fatalf("%s response message type = %q", event, resp.MessageType)
		}

		userCtx := userContext()
		userData := messageData(userCtx, event)
		userResp, userOK := r.Dispatch(userData, userCtx)

		if !userOK {
			t.Fatalf("%s event was not dispatched for access check", event)
		}
		if userResp.MessageType != "no_access" {
			t.Fatalf("%s user message type = %q, want no_access", event, userResp.MessageType)
		}
		if calls != 1 {
			t.Fatalf("%s unauthorized handler call changed calls to %d", event, calls)
		}
	}
}

func TestHEPDetailsRegistryPreservesEmptyPayloadBehavior(t *testing.T) {
	ctx := adminContext()
	data := messageData(ctx, eventGetHEPDetails)

	resp, ok := coreEvents.Dispatch(data, ctx)

	if !ok {
		t.Fatal("GetHEPDetails was not dispatched by registry")
	}
	if resp.MessageType != eventGetHEPDetails || resp.Error != "empty data" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestUpdateSettingsNormalizesWebSocketConfig(t *testing.T) {
	oldConfig := cfg.CustomPbx
	t.Cleanup(func() { cfg.CustomPbx = oldConfig })
	t.Setenv("CUSTOMPBX_CONFIG", t.TempDir()+"/config.json")

	payload := validSettingsPayload()
	payload.Web.Route = "ws"
	payload.XMLCurl.Route = "conf/config"
	payload.Web.WriteTimeoutSeconds = -1
	payload.Web.ReadTimeoutSeconds = 1
	payload.Web.PingIntervalSeconds = 100
	payload.Web.WebSocketQueueSize = cfg.MaxWebSocketQueueSize + 1
	payload.Web.OriginPolicy = cfg.OriginPolicySameOrigin

	resp := updateSettings(&webStruct.MessageData{Event: eventSetSettings, Payload: payload})

	if resp.Error != "" {
		t.Fatalf("unexpected error: %s", resp.Error)
	}
	if resp.Settings == nil {
		t.Fatal("settings response is nil")
	}
	if resp.Settings.Web.Route != "/ws" || resp.Settings.XMLCurl.Route != "/conf/config" {
		t.Fatalf("routes were not normalized: web=%q xml=%q", resp.Settings.Web.Route, resp.Settings.XMLCurl.Route)
	}
	if resp.Settings.Web.WriteTimeoutSeconds != cfg.DefaultWSWriteTimeoutSeconds ||
		resp.Settings.Web.ReadTimeoutSeconds != cfg.DefaultWSReadTimeoutSeconds ||
		resp.Settings.Web.PingIntervalSeconds != cfg.DefaultWSPingIntervalSeconds ||
		resp.Settings.Web.WebSocketQueueSize != cfg.MaxWebSocketQueueSize {
		t.Fatalf("websocket settings were not normalized: %+v", resp.Settings.Web)
	}
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

func managerContext() *webStruct.WsContext {
	ctx := webStruct.CreateWsContext(nil)
	ctx.SetUser(&mainStruct.WebUser{Id: 3, Login: "manager", GroupId: mainStruct.GetManagerId()})
	return ctx
}

func messageData(ctx *webStruct.WsContext, event string) *webStruct.MessageData {
	return &webStruct.MessageData{Event: event, Context: ctx}
}

func validSettingsPayload() cfg.GeneralCfg {
	return cfg.GeneralCfg{
		Fs: cfg.FreeSWITCH{
			Esl: cfg.Esl{Host: "127.0.0.1", Port: 8021, Pass: "change-me"},
		},
		Db: cfg.Database{
			Host: "127.0.0.1",
			Port: 5432,
			Name: "custompbx",
			User: "custompbx",
			Pass: "change-me",
		},
		Web: cfg.WebServer{
			Host:  "127.0.0.1",
			Port:  8080,
			Route: "/ws",
		},
		XMLCurl: cfg.WebServer{
			Host:  "127.0.0.1",
			Port:  8081,
			Route: "/conf/config",
		},
	}
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
