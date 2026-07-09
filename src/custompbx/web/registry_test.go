package web

import (
	"custompbx/cfg"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
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

func TestMessageMainHandlerInventoryHasNoRegistryOverlap(t *testing.T) {
	registered := stringSet(coreEvents.EventNames())
	remainingSwitchEvents := messageMainHandlerEventInventory(t)
	var overlap []string
	for _, event := range remainingSwitchEvents {
		if registered[event] {
			overlap = append(overlap, event)
		}
	}
	sort.Strings(overlap)
	if len(overlap) != 0 {
		t.Fatalf("events are both registered and still present in messageMainHandler fallback: %s", strings.Join(overlap, ", "))
	}
	t.Logf("event inventory: registered=%d remaining_switch=%d", len(registered), len(remainingSwitchEvents))
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
	events := []string{
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
	}
	assertAdminOnlyEventsDispatch(t, events)
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

func TestCombinedDataResponse(t *testing.T) {
	resp := combinedDataResponse("event",
		responseDataPair{name: "settings", data: "settings-data"},
		responseDataPair{name: "profiles", data: "profiles-data"},
	)
	if resp.MessageType != "event" {
		t.Fatalf("message type = %q, want event", resp.MessageType)
	}
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("data type = %T, want map[string]interface{}", resp.Data)
	}
	if data["settings"] != "settings-data" {
		t.Fatalf("settings data = %v", data["settings"])
	}
	if data["profiles"] != "profiles-data" {
		t.Fatalf("profiles data = %v", data["profiles"])
	}
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

func messageMainHandlerEventInventory(t *testing.T) []string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot resolve test source path")
	}
	wsPath := filepath.Join(filepath.Dir(file), "ws.go")
	body, err := os.ReadFile(wsPath)
	if err != nil {
		t.Fatal(err)
	}
	source := string(body)
	start := strings.Index(source, "func messageMainHandler")
	if start < 0 {
		t.Fatal("messageMainHandler not found")
	}
	end := strings.Index(source[start:], "func checkRelogin")
	if end < 0 {
		t.Fatal("messageMainHandler end marker not found")
	}
	block := source[start : start+end]
	caseRE := regexp.MustCompile(`(?m)^\s*case\s+([^:]+):`)
	quotedRE := regexp.MustCompile(`"([^"]+)"`)
	webStructEvents := map[string]string{
		"webStruct.GetDirectoryUser": webStruct.GetDirectoryUser,
	}
	events := map[string]bool{}
	for _, match := range caseRE.FindAllStringSubmatch(block, -1) {
		caseExpr := strings.TrimSpace(match[1])
		for _, quoted := range quotedRE.FindAllStringSubmatch(caseExpr, -1) {
			events[quoted[1]] = true
		}
		if resolved, ok := webStructEvents[caseExpr]; ok {
			events[resolved] = true
		}
	}
	names := make([]string, 0, len(events))
	for event := range events {
		names = append(names, event)
	}
	sort.Strings(names)
	return names
}

func stringSet(items []string) map[string]bool {
	set := make(map[string]bool, len(items))
	for _, item := range items {
		set[item] = true
	}
	return set
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
