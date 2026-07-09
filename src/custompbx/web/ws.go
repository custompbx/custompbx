package web

import (
	"crypto/rand"
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/cdrDb"
	"custompbx/cfg"
	"custompbx/daemonCache"
	"custompbx/db"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/logsafe"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"custompbx/webcache"
	"encoding/json"
	"fmt"
	"github.com/custompbx/customorm"
	"github.com/custompbx/hepparser"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

const (
	eventLogin                      = "login"
	eventSubscriptionList           = "SubscriptionList"
	eventPersistentSubscription     = "PersistentSubscription"
	eventRelogin                    = "relogin"
	eventLogOut                     = "[Auth] Logout"
	eventSwitchDialplanDebug        = "[Dialplan][Switch] Debug"
	eventGetSettings                = "get_settings"
	eventSetSettings                = "set_settings"
	eventGetInstances               = "GetInstances"
	eventUpdateInstanceDescription  = "UpdateInstanceDescription"
	eventGetWebSettings             = "GetWebSettings"
	eventSaveWebSettings            = "SaveWebSettings"
	eventGetCDR                     = "[CDR] Get"
	eventGetHEP                     = "GetHEP"
	eventGetHEPDetails              = "GetHEPDetails"
	eventGetLogs                    = "GetLogs"
	eventGetPhoneCreds              = "[Phone][Get] Creds"
	eventSendFSCLICommand           = "SendFSCLICommand"
	eventRealFSCLIConnect           = "RealFSCLIConnect"
	eventRealFSCLICommand           = "RealFSCLICommand"
	eventUpdateWebUserGroup         = "UpdateWebUserGroup"
	eventGetWebDirUserTemplates     = "GetWebDirectoryUsersTemplates"
	eventAddWebDirUserTemplate      = "AddWebDirectoryUsersTemplate"
	eventDelWebDirUserTemplate      = "DelWebDirectoryUsersTemplate"
	eventUpdateWebDirUserTemplate   = "UpdateWebDirectoryUsersTemplate"
	eventSwitchWebDirUserTemplate   = "SwitchWebDirectoryUsersTemplate"
	eventGetWebDirUserTplParams     = "GetWebDirectoryUsersTemplateParameters"
	eventAddWebDirUserTplParam      = "AddWebDirectoryUsersTemplateParameter"
	eventDelWebDirUserTplParam      = "DelWebDirectoryUsersTemplateParameter"
	eventSwitchWebDirUserTplParam   = "SwitchWebDirectoryUsersTemplateParameter"
	eventUpdateWebDirUserTplParam   = "UpdateWebDirectoryUsersTemplateParameter"
	eventGetWebDirUserTplVars       = "GetWebDirectoryUsersTemplateVariables"
	eventAddWebDirUserTplVar        = "AddWebDirectoryUsersTemplateVariable"
	eventDelWebDirUserTplVar        = "DelWebDirectoryUsersTemplateVariable"
	eventSwitchWebDirUserTplVar     = "SwitchWebDirectoryUsersTemplateVariable"
	eventUpdateWebDirUserTplVar     = "UpdateWebDirectoryUsersTemplateVariable"
	eventGetWebDirUserTplList       = "GetWebDirectoryUsersTemplatesList"
	eventGetWebDirUserTplForm       = "GetWebDirectoryUsersTemplateForm"
	eventCreateWebDirUserByTpl      = "CreateWebDirectoryUsersByTemplate"
	eventSettingsUsersGet           = "[Settings][Users] Get"
	eventGetWebUsersByDirectory     = "GetWebUsersByDirectory"
	eventSettingsUsersAdd           = "[Settings][Users] Add"
	eventSettingsUsersRename        = "[Settings][Users] Rename"
	eventSettingsUsersDelete        = "[Settings][Users] Delete"
	eventSettingsUsersSwitch        = "[Settings][Users][Switch] Web user"
	eventSettingsUsersUpdatePass    = "[Settings][Users][Update] Password"
	eventSettingsUsersUpdateLang    = "[Settings][Users][Update] Lang"
	eventSettingsUsersUpdateSip     = "[Settings][Users][Update] Sip user"
	eventSettingsUsersUpdateWS      = "[Settings][Users][Update] Ws"
	eventSettingsUsersUpdateVerto   = "[Settings][Users][Update] Verto Ws"
	eventSettingsUsersUpdateRTC     = "[Settings][Users][Update] WebRTC Lib"
	eventSettingsUsersUpdateStun    = "[Settings][Users][Update] Stun"
	eventSettingsUsersUpdateAvatar  = "[Settings][Users][Update] Avatar"
	eventSettingsUsersClearAvatar   = "[Settings][Users][Clear] Avatar"
	eventGetConvPrivateMessages     = "GetConversationPrivateMessages"
	eventGetConvPrivateCalls        = "GetConversationPrivateCalls"
	eventGetConvRoomMessages        = "GetConversationRoomMessages"
	eventSendConvPrivateMessage     = "SendConversationPrivateMessage"
	eventSendConvPrivateCall        = "SendConversationPrivateCall"
	eventSendConvRoomMessage        = "SendConversationRoomMessage"
	eventSwitchDialplanNoProceed    = "DialplanChangeNotProceed"
	eventDialplanGetSettings        = "DialplanGetSettings"
	eventDialplanGetContexts        = "[Dialplan][Get] Contexts"
	eventDialplanImport             = "[Dialplan][Import]"
	eventDialplanGetExtensions      = "[Dialplan][Get] Extensions"
	eventDialplanGetConditions      = "[Dialplan][Get] Conditions"
	eventDialplanGetExtDetails      = "[Dialplan][Get] Extension details"
	eventDialplanMoveExtension      = "[Dialplan][Move] Extension"
	eventDialplanMoveCondition      = "[Dialplan][Move] Condition"
	eventDialplanMoveAction         = "[Dialplan][Move] Action"
	eventDialplanMoveAntiaction     = "[Dialplan][Move] Antiaction"
	eventDialplanAddRegex           = "[Dialplan][Add] Regex"
	eventDialplanAddAction          = "[Dialplan][Add] Action"
	eventDialplanAddAntiaction      = "[Dialplan][Add] Antiaction"
	eventDialplanDeleteRegex        = "[Dialplan][Delete] Regex"
	eventDialplanDeleteAction       = "[Dialplan][Delete] Action"
	eventDialplanDeleteAntiaction   = "[Dialplan][Delete] Antiaction"
	eventDialplanUpdateRegex        = "[Dialplan][Update] Regex"
	eventDialplanUpdateAction       = "[Dialplan][Update] Action"
	eventDialplanUpdateAntiaction   = "[Dialplan][Update] Antiaction"
	eventDialplanSwitchRegex        = "[Dialplan][Switch] Regex"
	eventDialplanSwitchAction       = "[Dialplan][Switch] Action"
	eventDialplanSwitchAntiaction   = "[Dialplan][Switch] Antiaction"
	eventDialplanAddContext         = "[Dialplan][Add] Context"
	eventDialplanAddExtension       = "[Dialplan][Add] Extension"
	eventDialplanAddCondition       = "[Dialplan][Add] Condition"
	eventDialplanRenameContext      = "[Dialplan][Rename] Context"
	eventDialplanRenameExtension    = "[Dialplan][Rename] Extension"
	eventDialplanDeleteContext      = "[Dialplan][Delete] Context"
	eventDialplanDeleteExtension    = "[Dialplan][Delete] Extension"
	eventDialplanSwitchExtContinue  = "[Dialplan][Switch] Extension Continue"
	eventDialplanUpdateCondition    = "[Dialplan][Update] Condition"
	eventDialplanSwitchCondition    = "[Dialplan][Switch] Condition"
	eventDialplanDeleteCondition    = "[Dialplan][Delete] Condition"
	eventDirDomainsGet              = "GetDirectoryDomains"
	eventDirImport                  = "ImportDirectory"
	eventDirDomainImportXML         = "ImportXMLDomain"
	eventDirDomainAdd               = "AddDirectoryDomain"
	eventDirDomainRename            = "RenameDirectoryDomain"
	eventDirDomainSwitch            = "SwitchDirectoryDomain"
	eventDirDomainDelete            = "DeleteDirectoryDomain"
	eventDirDomainDetails           = "GetDirectoryDomainDetails"
	eventDirDomainAddParam          = "AddDirectoryDomainParameter"
	eventDirDomainUpdateParam       = "UpdateDirectoryDomainParameter"
	eventDirDomainSwitchParam       = "SwitchDirectoryDomainParameter"
	eventDirDomainDeleteParam       = "DeleteDirectoryDomainParameter"
	eventDirDomainAddVar            = "AddDirectoryDomainVariable"
	eventDirDomainUpdateVar         = "UpdateDirectoryDomainVariable"
	eventDirDomainSwitchVar         = "SwitchDirectoryDomainVariable"
	eventDirDomainDeleteVar         = "DeleteDirectoryDomainVariable"
	eventDirUserDetails             = "GetDirectoryUserDetails"
	eventDirUserAddParam            = "AddDirectoryUserParameter"
	eventDirUserAddVar              = "AddDirectoryUserVariable"
	eventDirUserDeleteParam         = "DeleteDirectoryUserParameter"
	eventDirUserDeleteVar           = "DeleteDirectoryUserVariable"
	eventDirUserUpdateParam         = "UpdateDirectoryUserParameter"
	eventDirUserUpdateVar           = "UpdateDirectoryUserVariable"
	eventDirUserUpdateCache         = "UpdateDirectoryUserCache"
	eventDirUserUpdateCidr          = "UpdateDirectoryUserCidr"
	eventDirUserUpdateNumberAlias   = "UpdateDirectoryUserNumberAlias"
	eventDirUserAdd                 = "AddDirectoryUser"
	eventDirUserImportXML           = "ImportXMLDomainUser"
	eventDirUserDelete              = "DeleteDirectoryUser"
	eventDirUserUpdateName          = "UpdateDirectoryUserName"
	eventDirUserSwitch              = "SwitchDirectoryUser"
	eventDirUserSwitchParam         = "SwitchDirectoryUserParameter"
	eventDirUserSwitchVar           = "SwitchDirectoryUserVariable"
	eventDirGroupsGet               = "GetDirectoryGroups"
	eventDirGroupUsersGet           = "GetDirectoryGroupUsers"
	eventDirGroupAdd                = "AddNewDirectoryGroup"
	eventDirGroupDelete             = "DeleteDirectoryGroup"
	eventDirGroupUpdateName         = "UpdateDirectoryGroupName"
	eventDirGroupUserAdd            = "AddDirectoryGroupUser"
	eventDirGroupUserDelete         = "DeleteDirectoryGroupUser"
	eventDirUserGatewaysGet         = "GetDirectoryUserGateways"
	eventDirUserGatewayDetails      = "GetDirectoryUserGatewayDetails"
	eventDirUserGatewayAddParam     = "AddDirectoryUserGatewayParameter"
	eventDirUserGatewayDeleteParam  = "DeleteDirectoryUserGatewayParameter"
	eventDirUserGatewayUpdateParam  = "UpdateDirectoryUserGatewayParameter"
	eventDirUserGatewaySwitchParam  = "SwitchDirectoryUserGatewayParameter"
	eventDirUserGatewayAddVar       = "AddDirectoryUserGatewayVariable"
	eventDirUserGatewayUpdateVar    = "UpdateDirectoryUserGatewayVariable"
	eventDirUserGatewaySwitchVar    = "SwitchDirectoryUserGatewayVariable"
	eventDirUserGatewayDeleteVar    = "DeleteDirectoryUserGatewayVariable"
	eventDirUserGatewayAdd          = "AddDirectoryUserGateway"
	eventDirUserGatewayDelete       = "DeleteDirectoryUserGateway"
	eventDirUserGatewayUpdateName   = "UpdateDirectoryUserGatewayName"
	eventConfigModuleReload         = "[Config][Reload] Module"
	eventConfigModuleUnload         = "[Config][Unload] Module"
	eventConfigModuleLoad           = "[Config][Load] Module"
	eventConfigModuleSwitch         = "[Config][Switch] Module"
	eventConfigModuleFromScratch    = "[Config][From scratch] Module"
	eventConfigModuleImport         = "[Config][Import] Module"
	eventConfigModulesImportAll     = "[Config][Import] All Modules"
	eventConfigModuleTruncate       = "TruncateModuleConfig"
	eventConfigModuleImportXML      = "ImportXMLModuleConfig"
	eventConfigModuleAutoload       = "[Config][Autoload] Module"
	eventACLListsGet                = "GetAclLists"
	eventACLListAdd                 = "AddAclList"
	eventACLListUpdateDefault       = "UpdateAclListDefault"
	eventACLListUpdate              = "UpdateAclList"
	eventACLListDelete              = "DelAclList"
	eventACLListConfigUpdateDefault = "[Config] Update_acl_list_default"
	eventACLNodesGet                = "GetAclNodes"
	eventACLNodeAdd                 = "AddAclNode"
	eventACLNodeDelete              = "DelAclNode"
	eventACLNodeUpdate              = "UpdateAclNode"
	eventACLNodeSwitch              = "SwitchAclNode"
	eventACLNodeMove                = "MoveAclListNode"
	eventSofiaGlobalSettingsGet     = "[Config] Get_sofia_global_settings"
	eventSofiaGlobalSettingUpdate   = "[Config] Update_sofia_global_setting"
	eventSofiaGlobalSettingSwitch   = "[Config] Switch_sofia_global_setting"
	eventSofiaGlobalSettingAdd      = "[Config] Add_sofia_global_setting"
	eventSofiaGlobalSettingDelete   = "[Config] Del_sofia_global_setting"
	eventSofiaProfileParamsGet      = "[Config] Get_sofia_profiles_params"
	eventSofiaProfileParamAdd       = "[Config] Add_sofia_profile_param"
	eventSofiaProfileParamDelete    = "[Config] Del_sofia_profile_param"
	eventSofiaProfileParamSwitch    = "[Config] Switch_sofia_profile_param"
	eventSofiaProfileParamUpdate    = "[Config] Update_sofia_profile_param"
	eventSofiaProfileGatewaysGet    = "[Config] Get_sofia_profile_gateways"
	eventSofiaGatewayVarsGet        = "GetSofiaProfileGatewayVariables"
	eventSofiaGatewayParamsGet      = "GetSofiaProfileGatewayParameters"
	eventSofiaGatewayParamAdd       = "[Config] Add_sofia_profile_gateway_param"
	eventSofiaGatewayParamUpdate    = "[Config] Update_sofia_profile_gateway_param"
	eventSofiaGatewayParamSwitch    = "[Config] Switch_sofia_profile_gateway_param"
	eventSofiaGatewayParamDelete    = "[Config] Del_sofia_profile_gateway_param"
	eventSofiaGatewayVarAdd         = "[Config] Add_sofia_profile_gateway_var"
	eventSofiaGatewayVarUpdate      = "[Config] Update_sofia_profile_gateway_var"
	eventSofiaGatewayVarSwitch      = "[Config] Switch_sofia_profile_gateway_var"
	eventSofiaGatewayVarDelete      = "[Config] Del_sofia_profile_gateway_var"
	eventSofiaGatewayAdd            = "[Config] Add_sofia_profile_gateway"
	eventSofiaGatewayDelete         = "[Config] Del_sofia_profile_gateway"
	eventSofiaGatewayRename         = "[Config] Rename_sofia_profile_gateway"
	eventSofiaProfileDomainsGet     = "[Config] Get_sofia_profile_domains"
	eventSofiaProfileDomainAdd      = "[Config] Add_sofia_profile_domain"
	eventSofiaProfileDomainDelete   = "[Config] Del_sofia_profile_domain"
	eventSofiaProfileDomainSwitch   = "[Config] Switch_sofia_profile_domain"
	eventSofiaProfileDomainUpdate   = "[Config] Update_sofia_profile_domain"
	eventSofiaProfileAliasesGet     = "[Config] Get_sofia_profile_aliases"
	eventSofiaProfileAliasAdd       = "[Config] Add_sofia_profile_alias"
	eventSofiaProfileAliasDelete    = "[Config] Del_sofia_profile_alias"
	eventSofiaProfileAliasSwitch    = "[Config] Switch_sofia_profile_alias"
	eventSofiaProfileAliasUpdate    = "[Config] Update_sofia_profile_alias"
	eventSofiaProfileAdd            = "[Config] Add_sofia_profile"
	eventSofiaProfileRename         = "[Config] Rename_sofia_profile"
	eventSofiaProfileDelete         = "[Config] Del_sofia_profile"
	eventSofiaProfileCommand        = "[API] Sofia profile command"
	eventSofiaProfileSwitch         = "[Config] Switch_sofia_profile"
	eventCdrPgCsvGet                = "[Config][Get] Cdr_Pg_Csv"
	eventCdrPgCsvParamAdd           = "[Config][Add] Cdr_Pg_Csv Parameter"
	eventCdrPgCsvParamUpdate        = "[Config][Update] Cdr_Pg_Csv Parameter"
	eventCdrPgCsvParamSwitch        = "[Config][Switch] Cdr_Pg_Csv Parameter"
	eventCdrPgCsvParamDelete        = "[Config][Delete] Cdr_Pg_Csv Parameter"
	eventCdrPgCsvFieldAdd           = "[Config][Add] Cdr_Pg_Csv Field"
	eventCdrPgCsvFieldUpdate        = "[Config][Update] Cdr_Pg_Csv Field"
	eventCdrPgCsvFieldSwitch        = "[Config][Switch] Cdr_Pg_Csv Field"
	eventCdrPgCsvFieldDelete        = "[Config][Delete] Cdr_Pg_Csv Field"
	eventOdbcCdrGet                 = "GetOdbcCdr"
	eventOdbcCdrFieldGet            = "GetOdbcCdrField"
	eventOdbcCdrParamAdd            = "AddOdbcCdrParameter"
	eventOdbcCdrParamUpdate         = "UpdateOdbcCdrParameter"
	eventOdbcCdrParamSwitch         = "SwitchOdbcCdrParameter"
	eventOdbcCdrParamDelete         = "DeleteOdbcCdrParameter"
	eventOdbcCdrTableAdd            = "AddOdbcCdrTable"
	eventOdbcCdrTableUpdate         = "UpdateOdbcCdrTable"
	eventOdbcCdrTableSwitch         = "SwitchOdbcCdrTable"
	eventOdbcCdrTableDelete         = "DeleteOdbcCdrTable"
	eventOdbcCdrFieldAdd            = "AddOdbcCdrField"
	eventOdbcCdrFieldUpdate         = "UpdateOdbcCdrField"
	eventOdbcCdrFieldSwitch         = "SwitchOdbcCdrField"
	eventOdbcCdrFieldDelete         = "DeleteOdbcCdrField"
)

var eventChannel chan interface{}
var messageUserLookup = findUser

func onlyAdminGroup() []int {
	return []int{mainStruct.GetAdminId()}
}

func onlyAdminAndManagerGroup() []int {
	return []int{mainStruct.GetAdminId(), mainStruct.GetManagerId()}
}

func adminOnly() []int {
	return onlyAdminGroup()
}

func adminOrManager() []int {
	return onlyAdminAndManagerGroup()
}

func onlyAdminManagerAndUserGroup() []int {
	return []int{mainStruct.GetAdminId(), mainStruct.GetManagerId(), mainStruct.GetUserId()}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return CheckWebSocketOrigin(cfg.CustomPbx.Web.OriginPolicy, cfg.CustomPbx.Web.AllowedOrigins, r)
	},
}

func SetBroadcastChannel(brChannel chan interface{}) {
	eventChannel = brChannel
}

var b = webStruct.NewWsHub()

func TimeEvents() {
	twoSecondsTick := time.Tick(2 * time.Second)

	for {
		select {
		case event := <-eventChannel:
			switch event.(type) {
			case *mainStruct.Dashboard:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetDashboard, Dashboard: event.(*mainStruct.Dashboard)})
			case *altStruct.ConfigSofiaProfile:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetSofiaProfiles, Data: map[int64]*altStruct.ConfigSofiaProfile{event.(*altStruct.ConfigSofiaProfile).Id: event.(*altStruct.ConfigSofiaProfile)}})
			case *altStruct.ConfigSofiaProfileGateway:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetSofiaProfileGateways, Data: map[int64]*altStruct.ConfigSofiaProfileGateway{event.(*altStruct.ConfigSofiaProfileGateway).Id: event.(*altStruct.ConfigSofiaProfileGateway)}})
			case *altStruct.Configurations:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetModules, Module: event.(*altStruct.Configurations)})
			case *mainStruct.DialplanDebug:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.DialplanDebug, DialplanDebug: event.(*mainStruct.DialplanDebug)})
			case *altStruct.DirectoryDomainUser:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetDirectoryUser, Data: struct {
					A interface{} `json:"directory_users"`
				}{A: event.(*altStruct.DirectoryDomainUser)}})
			case *map[int64]*mainStruct.Agent:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.SubscribeCallcenterAgents, CallcenterAgentsList: event.(*map[int64]*mainStruct.Agent)})
			case *mainStruct.DaemonState:
				b.Broadcast(webStruct.UserResponse{MessageType: webStruct.BroadcastConnection, Daemon: event.(*mainStruct.DaemonState)})
			case *hepparser.HEP:
				// b.Broadcast(webStruct.UserResponse{MessageType: webStruct.SubscribeHepPackages, HEPs: event.(*hepparser.HEP)})
			default:
				log.Printf("Unknown event type: %T - %+v\n", event, event)
			}

		case <-twoSecondsTick:
			b.Broadcast(webStruct.UserResponse{MessageType: webStruct.GetDashboard, Dashboard: &mainStruct.Dashboard{DashboardData: webcache.GetDashboardData()}})
		}
	}
}

func StartWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("ERROR on StartWS: %+v", err.Error())
		return
	}
	fmt.Println("NEW WS CONNECTION")
	wsContext := webStruct.CreateWsContext(ws)
	b.Register(wsContext)

	fmt.Println("STARTING GOROUTINES")
	go wsContext.SendWaiter()
	go wsContext.ReadWaiter(messageHandler)
}

func PostAPIRequest(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var msg webStruct.Message
	err := decoder.Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}
	if err := msg.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}

	var resp webStruct.UserResponse
	msg.Data.Trim()
	msg.Data.Event = msg.Event
	msg.Data.Context = webStruct.CreateWsContext(nil)

	// find user by token
	user, _ := messageUserLookup(msg.Data)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 Unauthorized"))
		return
	}
	msg.Data.Context.SetUser(user)

	resp = dispatchMessage(msg.Data, msg.Data.Context)
	res, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Server Error"))
		return
	}
	_, err = w.Write(res)
	if err != nil {
		log.Printf("%+v", err)
	}
}

func HubMetrics(w http.ResponseWriter, r *http.Request) {
	user, status := UserFromBearer(r)
	if status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}
	if status := RequireGroups(user, mainStruct.GetAdminId()); status != http.StatusOK {
		http.Error(w, http.StatusText(status), status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(b.Metrics()); err != nil {
		log.Printf("component=websocket operation=write_metrics error=%q", err)
	}
}

func tokenGenerator() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func messageHandler(msg *webStruct.Message, wsContext *webStruct.WsContext) {
	defer func() {
		if r := recover(); r != nil {
			wsContext.RecordHandlerFailure()
			log.Printf("component=websocket connection_id=%d user_id=%d operation=message_handler error=%v", wsContext.ID, wsContext.UserID(), r)
			log.Printf("component=websocket connection_id=%d user_id=%d operation=message_handler stacktrace=%q", wsContext.ID, wsContext.UserID(), string(debug.Stack()))
		}
	}()
	if err := msg.Validate(); err != nil {
		wsContext.RecordHandlerFailure()
		sendWSResponse(wsContext, &webStruct.UserResponse{Error: "invalid message", MessageType: "none"})
		return
	}
	if !daemonCache.State.DatabaseConnection {
		sendWSResponse(wsContext, &webStruct.UserResponse{Daemon: daemonCache.State, MessageType: webStruct.BroadcastConnection})
		return
	}

	msg.Data.Trim()
	msg.Data.Event = msg.Event
	msg.Data.Context = wsContext

	// first check if it login request
	if msg.Event == eventLogin {
		resp := checkLogin(msg.Data)
		sendWSResponse(wsContext, &resp)
		return
	}
	// allow without token
	if msg.Event == "get_status" {
		resp := webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "connection"}
		sendWSResponse(wsContext, &resp)
		return
	}
	log.Println("EVENT: ", msg.Event)

	// find user by token
	user, response := messageUserLookup(msg.Data)
	if user == nil {
		log.Println("EVENT: ", msg.Event, "NO USER")
		sendWSResponse(wsContext, &response)
		return
	}
	subsResp := subscribeUser(msg.Data)
	if subsResp != nil {
		log.Println("EVENT: ", msg.Event, "NO SUBS")
		sendWSResponse(wsContext, subsResp)
		return
	}

	resp := dispatchMessage(msg.Data, wsContext)
	sendWSResponse(wsContext, &resp)
}

func dispatchMessage(data *webStruct.MessageData, wsContext *webStruct.WsContext) webStruct.UserResponse {
	if registeredResponse, ok := coreEvents.Dispatch(data, wsContext); ok {
		return registeredResponse
	}
	return messageMainHandler(data)
}

func sendWSResponse(wsContext *webStruct.WsContext, resp *webStruct.UserResponse) {
	if !wsContext.Enqueue(resp) {
		_ = wsContext.CloseWithReason("outbound queue full")
	}
}

func messageMainHandler(msg *webStruct.MessageData) webStruct.UserResponse {
	// if !daemonCache.State.DatabaseConnection || !daemonCache.State.ESLConnection {
	if !daemonCache.State.DatabaseConnection {
		return webStruct.UserResponse{Daemon: daemonCache.State, MessageType: webStruct.BroadcastConnection}
	}

	var resp webStruct.UserResponse

	switch msg.Event {
	//Doc started ---- (//Request:.*parent":\{.*id":\d+)[^\}]+ //(Request:.*),"description":""(.+) //(Request:(?!.*Switch)+)(.+)(,|\{)"enabled":(?:false|true)(?:,|\{) //(Request:(.*Switch)+.+),"parent":\{"id":\d+\} //(Request:(?!.*Move)+.*),"position":\d+
	//## Directory
	//### Domains
	//Request:{"event":"GetDirectoryDomains","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetDirectoryDomains","data":{"4":{"id":4,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0}}}
	//Errors:no id, DB error
	//Request:{"event":"ImportDirectory","data":{"token":"example-token"}}
	//Response:{"MessageType":"ImportDirectory"}
	//Errors:empty data, can't parse
	//Request:{"event":"ImportXMLDomain","data":{"token":"example-token","file":"\r\n  <!--the domain or ip (the right hand side of the @ in the addr-->\r\n  <domain name=\"test_do\">\r\n    <params>\r\n      <param name=\"jsonrpc-allowed-methods\" value=\"verto\"/>>\r\n    </params>\r\n    <variables>\r\n      <variable name=\"record_stereo\" value=\"true\"/>\r\n      <variable name=\"default_gateway\" value=\"$${default_provider}\"/>\r\n    </variables>\r\n    <groups>\r\n      <group name=\"default\">\r\n\t<users>\r\n\t</users>\r\n      </group>\r\n    </groups>\r\n  </domain>"}}
	//Response:{"MessageType":"ImportXMLDomain","data":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0},"7":{"id":7,"position":2,"enabled":true,"name":"test_do","parent":{"id":1},"sip_regs_counter":0}}}
	//Errors:can't parse
	//Request:{"event":"AddDirectoryDomain","data":{"token":"example-token","name":"test"}}
	//Response:{"MessageType":"AddDirectoryDomain","data":{"id":9,"position":4,"enabled":true,"name":"test","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	//Request:{"event":"RenameDirectoryDomain","data":{"token":"example-token","id":9,"name":"test2"}}
	//Response:{"MessageType":"RenameDirectoryDomain","data":{"id":9,"position":4,"enabled":true,"name":"test2","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	//Request:{"event":"SwitchDirectoryDomain","data":{"token":"example-token","id":8,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryDomain","data":{"id":8,"position":3,"enabled":false,"name":"test_do2","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	//Request:{"event":"DeleteDirectoryDomain","data":{"token":"example-token","id":9}}
	//Response:{"MessageType":"DeleteDirectoryDomain","data":{"id":9,"position":4,"enabled":true,"name":"test2","parent":{"id":1},"sip_regs_counter":0}}
	//Errors:DB error
	//Request:{"event":"GetDirectoryDomainDetails","data":{"token":"example-token","id":6}}
	//Response:{"MessageType":"GetDirectoryDomainDetails","data":{"parameters":{"10":{"id":10,"position":2,"enabled":true,"name":"jsonrpc-allowed-methods","value":"verto","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"9":{"id":9,"position":1,"enabled":true,"name":"dial-string","value":"{^^:sip_invite_domain=${dialed_domain}:presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(*/${dialed_user}@${dialed_domain})},${verto_contact(${dialed_user}@${dialed_domain})}","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}},"variables":{"13":{"id":13,"position":1,"enabled":true,"name":"record_stereo","value":"true","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"14":{"id":14,"position":2,"enabled":true,"name":"default_gateway","value":"example.com","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"15":{"id":15,"position":3,"enabled":true,"name":"default_areacode","value":"918","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"16":{"id":16,"position":4,"enabled":true,"name":"transfer_fallback_extension","value":"operator","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}}}
	//Errors:DB error
	//Request:{"event":"AddDirectoryDomainParameter","data":{"token":"example-token","id":6,"name":"paramn","value":"paramv"}}
	//Response:{"MessageType":"AddDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:DB error
	//Request:{"event":"UpdateDirectoryDomainParameter","data":{"token":"example-token","id":15,"name":"paramn1","value":"paramv1"}}
	//Response:{"MessageType":"UpdateDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":true,"name":"paramn1","value":"paramv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"SwitchDirectoryDomainParameter","data":{"token":"example-token","id":15,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":false,"name":"paramn1","value":"paramv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryDomainParameter","data":{"token":"example-token","id":1}}
	//Response:{"MessageType":"DeleteDirectoryDomainParameter","data":{"id":15,"position":3,"enabled":false,"name":"paramn1","value":"paramv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"AddDirectoryDomainVariable","data":{"token":"example-token","id":6,"name":"varn","value":"varv"}}
	//Response:{"MessageType":"AddDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":true,"name":"varn","value":"varv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryDomainVariable","data":{"token":"example-token","id":21,"name":"varn1","value":"varv1"}}
	//Response:{"MessageType":"UpdateDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":true,"name":"varn1","value":"varv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"SwitchDirectoryDomainVariable","data":{"token":"example-token","id":21,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":false,"name":"varn1","value":"varv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryDomainVariable","data":{"token":"example-token","id":2}}
	//Response:{"MessageType":"DeleteDirectoryDomainVariable","data":{"id":21,"position":5,"enabled":false,"name":"varn1","value":"varv1","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//### Users
	//Request:{"event":"GetDirectoryUsers","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetDirectoryUsers","data":{"domains":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0},"7":{"id":7,"position":2,"enabled":true,"name":"test_do","parent":{"id":1},"sip_regs_counter":0},"8":{"id":8,"position":3,"enabled":true,"name":"test_do2","parent":{"id":1},"sip_regs_counter":0}},"directory_users":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}}
	//Errors:
	//Request:{"event":"GetDirectoryUserDetails","data":{"token":"example-token","id":91}}
	//Response:{"MessageType":"GetDirectoryUserDetails","data":{"parameters":{"93":{"id":93,"position":1,"enabled":true,"name":"password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"94":{"id":94,"position":2,"enabled":true,"name":"vm-password","value":"1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}},"variables":{"346":{"id":346,"position":1,"enabled":true,"name":"toll_allow","value":"domestic,international,local","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"347":{"id":347,"position":2,"enabled":true,"name":"accountcode","value":"1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"348":{"id":348,"position":3,"enabled":true,"name":"user_context","value":"default","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"349":{"id":349,"position":4,"enabled":true,"name":"effective_caller_id_name","value":"Extension 1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"350":{"id":350,"position":5,"enabled":true,"name":"effective_caller_id_number","value":"1000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"351":{"id":351,"position":6,"enabled":true,"name":"outbound_caller_id_name","value":"FreeSWITCH","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"352":{"id":352,"position":7,"enabled":true,"name":"outbound_caller_id_number","value":"0000000000","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"353":{"id":353,"position":8,"enabled":true,"name":"callgroup","value":"techsupport","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}},"user":{"id":91,"position":1,"enabled":true,"name":"1000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"AddDirectoryUserParameter","data":{"token":"example-token","id":91,"name":"paramn","value":"paramv"}}
	//Response:{"MessageType":"AddDirectoryUserParameter","data":{"id":137,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"AddDirectoryUserVariable","data":{"token":"example-token","id":91,"name":"varn","value":"varv"}}
	//Response:{"MessageType":"AddDirectoryUserVariable","data":{"id":514,"position":9,"enabled":true,"name":"varn","value":"varv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryUserParameter","data":{"token":"example-token","id":13}}
	//Response:{"MessageType":"DeleteDirectoryUserParameter","data":{"id":137,"position":3,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryUserVariable","data":{"token":"example-token","id":51}}
	//Response:{"MessageType":"DeleteDirectoryUserVariable","data":{"id":514,"position":9,"enabled":true,"name":"varn","value":"varv","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserParameter","data":{"token":"example-token","id":138,"name":"paramn1","value":"paramv1"}}
	//Response:{"MessageType":"UpdateDirectoryUserParameter","data":{"id":138,"position":3,"enabled":true,"name":"paramn1","value":"paramv1","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserVariable","data":{"token":"example-token","id":515,"name":"varn1","value":"varv1"}}
	//Response:{"MessageType":"UpdateDirectoryUserVariable","data":{"id":515,"position":9,"enabled":true,"name":"varn1","value":"varv1","description":"","parent":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserCache","data":{"token":"example-token","value":"3000","id":91}}
	//Response:{"MessageType":"UpdateDirectoryUserCache","data":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserCidr","data":{"token":"example-token","value":"0.0.0.0","id":91}}
	//Response:{"MessageType":"UpdateDirectoryUserCidr","data":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"0.0.0.0","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserNumberAlias","data":{"token":"example-token","value":"555","id":91}}
	//Response:{"MessageType":"UpdateDirectoryUserNumberAlias","data":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"0.0.0.0","number_alias":"555","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	//Request:{"event":"AddDirectoryUser","data":{"token":"example-token","name":"5000","id":6}}
	//Response:{"MessageType":"AddDirectoryUser","data":{"115":{"id":115,"position":25,"enabled":true,"name":"5000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"ImportXMLDomainUser","data":{"token":"example-token","file":"<user id=\"1099\">\r\n    <params>\r\n      <param name=\"password\" value=\"$${default_password}\"/>\r\n      <param name=\"vm-password\" value=\"1099\"/>\r\n    </params>\r\n    <variables>\r\n      <variable name=\"toll_allow\" value=\"domestic,international,local\"/>\r\n      <variable name=\"accountcode\" value=\"1099\"/>\r\n      <variable name=\"user_context\" value=\"default\"/>\r\n      <variable name=\"effective_caller_id_name\" value=\"Extension 1990\"/>\r\n      <variable name=\"effective_caller_id_number\" value=\"1099\"/>\r\n      <variable name=\"outbound_caller_id_name\" value=\"$${outbound_caller_name}\"/>\r\n      <variable name=\"outbound_caller_id_number\" value=\"$${outbound_caller_id}\"/>\r\n      <variable name=\"callgroup\" value=\"techsupport\"/>\r\n    </variables>\r\n  </user>","id":6}}
	//Response:{"MessageType":"ImportXMLDomainUser","data":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"102":{"id":102,"position":12,"enabled":true,"name":"1011","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"103":{"id":103,"position":13,"enabled":true,"name":"1012","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"104":{"id":104,"position":14,"enabled":true,"name":"1013","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"105":{"id":105,"position":15,"enabled":true,"name":"1014","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"106":{"id":106,"position":16,"enabled":true,"name":"1015","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"107":{"id":107,"position":17,"enabled":true,"name":"1016","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"108":{"id":108,"position":18,"enabled":true,"name":"1017","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"109":{"id":109,"position":19,"enabled":true,"name":"1018","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"110":{"id":110,"position":20,"enabled":true,"name":"1019","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"111":{"id":111,"position":21,"enabled":true,"name":"brian","cache":1000,"cidr":"192.0.2.0/24","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"112":{"id":112,"position":22,"enabled":true,"name":"default","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"113":{"id":113,"position":23,"enabled":true,"name":"example.com","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"114":{"id":114,"position":24,"enabled":true,"name":"SEP001120AABBCC","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"115":{"id":115,"position":25,"enabled":true,"name":"5000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"116":{"id":116,"position":26,"enabled":true,"name":"1099","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"91":{"id":91,"position":1,"enabled":true,"name":"1000","cache":3000,"cidr":"0.0.0.0","number_alias":"555","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"92":{"id":92,"position":2,"enabled":true,"name":"1001","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"93":{"id":93,"position":3,"enabled":true,"name":"1002","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"94":{"id":94,"position":4,"enabled":true,"name":"1003","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"95":{"id":95,"position":5,"enabled":true,"name":"1004","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"96":{"id":96,"position":6,"enabled":true,"name":"1005","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"97":{"id":97,"position":7,"enabled":true,"name":"1006","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"98":{"id":98,"position":8,"enabled":true,"name":"1007","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"99":{"id":99,"position":9,"enabled":true,"name":"1008","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryUser","data":{"token":"example-token","id":11}}
	//Response:{"MessageType":"DeleteDirectoryUser","data":{"id":115,"position":25,"enabled":true,"name":"5000","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserName","data":{"token":"example-token","id":116,"name":"1098"}}
	//Response:{"MessageType":"UpdateDirectoryUserName","data":{"id":116,"position":26,"enabled":true,"name":"1098","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	//Request:{"event":"SwitchDirectoryUser","data":{"token":"example-token","id":116,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUser","data":{"id":116,"position":26,"enabled":false,"name":"1098","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}
	//Errors:
	//Request:{"event":"SwitchDirectoryUserParameter","data":{"token":"example-token","id":140,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserParameter","data":{"id":140,"position":2,"enabled":false,"name":"vm-password","value":"1099","description":"","parent":{"id":116,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"SwitchDirectoryUserVariable","data":{"token":"example-token","id":516,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserVariable","data":{"id":516,"position":1,"enabled":false,"name":"toll_allow","value":"domestic,international,local","description":"","parent":{"id":116,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//### Groups
	//Request:{"event":"GetDirectoryGroups","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetDirectoryGroups","data":{"domains":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0},"7":{"id":7,"position":2,"enabled":true,"name":"test_do","parent":{"id":1},"sip_regs_counter":0},"8":{"id":8,"position":3,"enabled":true,"name":"test_do2","parent":{"id":1},"sip_regs_counter":0}},"list":{"10":{"id":10,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"11":{"id":11,"position":2,"enabled":true,"name":"sales","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"12":{"id":12,"position":3,"enabled":true,"name":"billing","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"13":{"id":13,"position":4,"enabled":true,"name":"support","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"14":{"id":14,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":7,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}},"15":{"id":15,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":8,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}}}
	//Errors:
	//Request:{"event":"GetDirectoryGroupUsers","data":{"token":"example-token","id":12}}
	//Response:{"MessageType":"GetDirectoryGroupUsers","data":{"group_users":{"114":{"id":114,"position":2,"enabled":true,"type":"","description":"","parent":{"id":12,"position":0,"enabled":false,"name":"","description":"","parent":null},"user":{"id":100,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}},"users":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}}
	//Errors:
	//Request:{"event":"AddNewDirectoryGroup","data":{"token":"example-token","id":6,"name":"new_group"}}
	//Response:{"MessageType":"AddNewDirectoryGroup","data":{"id":16,"position":5,"enabled":true,"name":"new_group","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryGroup","data":{"token":"example-token","id":1}}
	//Response:{"MessageType":"DeleteDirectoryGroup","data":{"id":16,"position":5,"enabled":true,"name":"new_group","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryGroupName","data":{"token":"example-token","id":17,"name":"newnew"}}
	//Response:{"MessageType":"UpdateDirectoryGroupName","data":{"id":17,"position":5,"enabled":true,"name":"newnew","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0}}}
	//Errors:
	//Request:{"event":"AddDirectoryGroupUser","data":{"token":"example-token","id_int":91,"id":13}}
	//Response:{"MessageType":"AddDirectoryGroupUser","data":{"id":120,"position":3,"enabled":true,"type":"","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","description":"","parent":null},"user":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryGroupUser","data":{"token":"example-token","id":12}}
	//Response:{"MessageType":"DeleteDirectoryGroupUser","data":{"id":120,"position":3,"enabled":true,"type":"","description":"","parent":{"id":13,"position":0,"enabled":false,"name":"","description":"","parent":null},"user":{"id":91,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//### Gateways
	//Request:{"event":"GetDirectoryUserGateways","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetDirectoryUserGateways","data":{"domains":{"6":{"id":6,"position":1,"enabled":true,"name":"45.61.54.76","parent":{"id":1},"sip_regs_counter":0}},"directory_users":{"100":{"id":100,"position":10,"enabled":true,"name":"1009","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"101":{"id":101,"position":11,"enabled":true,"name":"1010","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"102":{"id":102,"position":12,"enabled":true,"name":"1011","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"103":{"id":103,"position":13,"enabled":true,"name":"1012","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"104":{"id":104,"position":14,"enabled":true,"name":"1013","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false},"105":{"id":105,"position":15,"enabled":true,"name":"1014","cache":1000,"cidr":"","number_alias":"","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","parent":null,"sip_regs_counter":0},"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}},"user_gateways":{"5":{"id":5,"position":1,"enabled":true,"name":"example.com","description":"","parent":{"id":113,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}}}
	//Errors:
	//Request:{"event":"GetDirectoryUserGatewayDetails","data":{"token":"example-token","id":5}}
	//Response:{"MessageType":"GetDirectoryUserGatewayDetails","data":{"parameters":{"29":{"id":29,"position":1,"enabled":true,"name":"username","value":"joeuser","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"30":{"id":30,"position":2,"enabled":true,"name":"password","value":"password","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"31":{"id":31,"position":3,"enabled":true,"name":"from-user","value":"joeuser","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"32":{"id":32,"position":4,"enabled":true,"name":"from-domain","value":"example.com","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"33":{"id":33,"position":5,"enabled":true,"name":"expire-seconds","value":"600","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"34":{"id":34,"position":6,"enabled":true,"name":"register","value":"false","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"35":{"id":35,"position":7,"enabled":true,"name":"retry-seconds","value":"30","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"36":{"id":36,"position":8,"enabled":true,"name":"extension","value":"5000","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}},"37":{"id":37,"position":9,"enabled":true,"name":"context","value":"public","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}},"variables":{}}}
	//Errors:
	//Request:{"event":"AddDirectoryUserGatewayParameter","data":{"token":"example-token","id":5,"name":"paramn","value":"paramv"}}
	//Response:{"MessageType":"AddDirectoryUserGatewayParameter","data":{"id":38,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryUserGatewayParameter","data":{"token":"example-token","id":3}}
	//Response:{"MessageType":"DeleteDirectoryUserGatewayParameter","data":{"id":38,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserGatewayParameter","data":{"token":"example-token","id":39,"name":"param","value":"param_new_val"}}
	//Response:{"MessageType":"UpdateDirectoryUserGatewayParameter","data":{"id":39,"position":10,"enabled":true,"name":"param","value":"param_new_val","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"AddDirectoryUserGatewayVariable","data":{"token":"example-token","id":5,"name":"varn","value":"varv","direction":"vard"}}
	//Response:{"MessageType":"AddDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":true,"name":"varn","value":"varv","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserGatewayVariable","data":{"token":"example-token","id":4,"name":"varn","value":"varv2222","direction":"vard"}}
	//Response:{"MessageType":"UpdateDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":true,"name":"varn","value":"varv2222","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"SwitchDirectoryUserGatewayVariable","data":{"token":"example-token","id":4,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":false,"name":"varn","value":"varv2222","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryUserGatewayVariable","data":{"token":"example-token","id":4}}
	//Response:{"MessageType":"DeleteDirectoryUserGatewayVariable","data":{"id":4,"position":1,"enabled":false,"name":"varn","value":"varv2222","direction":"vard","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//Request:{"event":"AddDirectoryUserGateway","data":{"token":"example-token","name":"new_gw","id":93}}
	//Response:{"MessageType":"AddDirectoryUserGateway","data":{"id":6,"position":1,"enabled":true,"name":"new_gw","description":"","parent":{"id":93,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"DeleteDirectoryUserGateway","data":{"token":"example-token","id":6}}
	//Response:{"MessageType":"DeleteDirectoryUserGateway","data":{"id":6,"position":1,"enabled":true,"name":"new_gw2","description":"","parent":{"id":93,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"UpdateDirectoryUserGatewayName","data":{"token":"example-token","id":6,"name":"new_gw2"}}
	//Response:{"MessageType":"UpdateDirectoryUserGatewayName","data":{"id":6,"position":1,"enabled":true,"name":"new_gw2","description":"","parent":{"id":93,"position":0,"enabled":false,"name":"","cache":0,"cidr":"","number_alias":"","description":"","parent":null,"call_date":0,"in_call":false,"talking":false,"last_uuid":"","call_direction":"","sip_register":false,"verto_register":false}}}
	//Errors:
	//Request:{"event":"SwitchDirectoryUserGatewayParameter","data":{"token":"example-token","id":39,"enabled":false}}
	//Response:{"MessageType":"SwitchDirectoryUserGatewayParameter","data":{"id":39,"position":10,"enabled":false,"name":"param","value":"param_new_val","description":"","parent":{"id":5,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	//## Configuration
	//### Modules
	//Request:{"event":"[Config][Get] Modules","data":{"token":"example-token"}}
	//Response:{"MessageType":"[Config][Get] Modules","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":{"id":2,"position":2,"enabled":false,"name":"alsa.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"amr":{"id":52,"position":52,"enabled":true,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	//Request:{"event":"[Config][Reload] Module","data":{"token":"example-token","id":52}}
	//Response:{"MessageType":"[Config][Reload] Module"}
	//Errors:
	//Request:{"event":"[Config][Unload] Module","data":{"token":"example-token","id":52}}
	//Response:{"MessageType":"[Config][Unload] Module"}
	//Errors:
	//Request:{"event":"[Config][Load] Module","data":{"token":"example-token","id":52}}
	//Response:{"MessageType":"[Config][Load] Module"}
	//Errors:
	//Request:{"event":"[Config][Switch] Module","data":{"token":"example-token","id":52,"enabled":false}}
	//Response:{"MessageType":"[Config][Switch] Module","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":{"id":2,"position":2,"enabled":false,"name":"alsa.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"amr":{"id":52,"position":52,"enabled":false,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	//Request:{"event":"[Config][From scratch] Module","data":{"token":"example-token","name":"alsa"}}
	//Response:{"MessageType":"[Config][From scratch] Module"}
	//Errors:
	//Request:{"event":"[Config][Import] Module","data":{"token":"example-token","name":"alsa"}}
	//Response:{"MessageType":"[Config][Import] Module"}
	//Errors:
	//Request:{"event":"[Config][Import] All Modules","data":{"token":"example-token"}}
	//Response:{"MessageType":"[Config][Import] All Modules"}
	//Errors:
	//Request:{"event":"TruncateModuleConfig","data":{"token":"example-token","id":54}}
	//Response:{"MessageType":"TruncateModuleConfig","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":null,"amr":{"id":52,"position":52,"enabled":false,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	//Request:{"event":"ImportXMLModuleConfig","data":{"token":"example-token","file":"<configuration name=\"alsa.conf\" description=\"Soundcard Endpoint\">\r\n  <settings>\r\n    <!--Default dialplan and caller-id info -->\r\n    <param name=\"dialplan\" value=\"XML\"/>\r\n    <param name=\"cid-name\" value=\"N800 Alsa\"/>\r\n    <param name=\"cid-num\" value=\"5555551212\"/>\r\n\r\n    <!--audio sample rate and interval -->\r\n    <param name=\"sample-rate\" value=\"8000\"/>\r\n    <param name=\"codec-ms\" value=\"20\"/>\r\n  </settings>\r\n</configuration>"}}
	//Response:{"MessageType":"ImportXMLModuleConfig","modules":{"post_load_switch":{"id":43,"position":43,"enabled":true,"name":"post_load_switch.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"acl":{"id":1,"position":1,"enabled":true,"name":"acl.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"callcenter":{"id":6,"position":6,"enabled":true,"name":"callcenter.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_pg_csv":{"id":8,"position":8,"enabled":true,"name":"cdr_pg_csv.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"odbc_cdr":{"id":51,"position":51,"enabled":true,"name":"odbc_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lcr":{"id":24,"position":24,"enabled":true,"name":"lcr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sofia":{"id":42,"position":42,"enabled":true,"name":"sofia.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"verto":{"id":46,"position":46,"enabled":true,"name":"verto.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"shout":{"id":40,"position":40,"enabled":true,"name":"shout.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"redis":{"id":38,"position":38,"enabled":true,"name":"redis.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"nibblebill":{"id":29,"position":29,"enabled":true,"name":"nibblebill.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"db":{"id":14,"position":14,"enabled":true,"name":"db.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"distributor":{"id":17,"position":17,"enabled":true,"name":"distributor.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"memcache":{"id":26,"position":26,"enabled":true,"name":"memcache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"avmd":{"id":5,"position":5,"enabled":true,"name":"avmd.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"tts_commandline":{"id":44,"position":44,"enabled":true,"name":"tts_commandline.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cdr_mongodb":{"id":7,"position":7,"enabled":true,"name":"cdr_mongodb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"http_cache":{"id":23,"position":23,"enabled":true,"name":"http_cache.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"opus":{"id":31,"position":31,"enabled":true,"name":"opus.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"python":{"id":37,"position":37,"enabled":true,"name":"python.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"alsa":{"id":55,"position":53,"enabled":true,"name":"alsa.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"amr":{"id":52,"position":52,"enabled":false,"name":"amr.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"amrwb":{"id":4,"position":4,"enabled":true,"name":"amrwb.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cepstral":{"id":9,"position":9,"enabled":true,"name":"cepstral.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"cidlookup":{"id":10,"position":10,"enabled":true,"name":"cidlookup.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"curl":{"id":13,"position":13,"enabled":true,"name":"curl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"dialplan_directory":{"id":15,"position":15,"enabled":true,"name":"dialplan_directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"easyroute":{"id":18,"position":18,"enabled":true,"name":"easyroute.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"erlang_event":{"id":19,"position":19,"enabled":true,"name":"erlang_event.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"event_multicast":{"id":20,"position":20,"enabled":true,"name":"event_multicast.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fax":{"id":21,"position":21,"enabled":true,"name":"fax.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"lua":{"id":25,"position":25,"enabled":true,"name":"lua.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"mongo":{"id":27,"position":27,"enabled":true,"name":"mongo.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"msrp":{"id":28,"position":28,"enabled":true,"name":"msrp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"oreka":{"id":32,"position":32,"enabled":true,"name":"oreka.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"perl":{"id":34,"position":34,"enabled":true,"name":"perl.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"pocketsphinx":{"id":35,"position":35,"enabled":true,"name":"pocketsphinx.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sangoma_codec":{"id":39,"position":39,"enabled":true,"name":"sangoma_codec.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"sndfile":{"id":41,"position":41,"enabled":true,"name":"sndfile.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"xml_cdr":{"id":48,"position":48,"enabled":true,"name":"xml_cdr.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"xml_rpc":{"id":49,"position":49,"enabled":true,"name":"xml_rpc.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"zeroconf":{"id":50,"position":50,"enabled":true,"name":"zeroconf.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"directory":{"id":16,"position":16,"enabled":true,"name":"directory.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"fifo":{"id":22,"position":22,"enabled":true,"name":"fifo.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"opal":{"id":30,"position":30,"enabled":true,"name":"opal.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"osp":{"id":33,"position":33,"enabled":true,"name":"osp.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"unicall":{"id":45,"position":45,"enabled":true,"name":"unicall.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"conference":{"id":11,"position":11,"enabled":true,"name":"conference.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}},"conference_layouts":{"id":12,"position":12,"enabled":true,"name":"conference_layouts.conf","module":"","loaded":false,"unloadable":false,"parent":{"id":1}},"post_load_modules":{"id":36,"position":36,"enabled":true,"name":"post_load_modules.conf","module":"","loaded":false,"unloadable":true,"parent":{"id":1}},"voicemail":{"id":47,"position":47,"enabled":true,"name":"voicemail.conf","module":"","loaded":true,"unloadable":false,"parent":{"id":1}}}}
	//Errors:
	//Request:{"event":"[Config][Autoload] Module","data":{"token":"example-token","id":55}}
	//Response:{"MessageType":"[Config][Autoload] Module","data":{"id":15,"position":12,"enabled":true,"name":"mod_alsa","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	//### Acl/Sofia/CDR
	// Migrated to the WebSocket handler registry.
	//### Verto
	//Request:{"event":"[Config][Verto][Get]","data":{"token":"example-token"}}
	//Response:{"MessageType":"[Config][Verto][Get]","data":{"settings":{"1":{"id":1,"position":1,"enabled":true,"name":"debug","value":"0","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default-v4","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"default-v6","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "[Config][Verto][Get]":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigVertoSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigVertoProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"[Config][Verto][Profile][Parameters][Get]","data":{"token":"example-token","id":1}}
	//Response:{"MessageType":"[Config][Verto][Profile][Parameters][Get]","data":{"1":{"id":1,"position":1,"enabled":true,"name":"bind-local","value":"45.61.54.76:8081","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":9,"enabled":true,"name":"rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":10,"enabled":true,"name":"ext-rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":11,"enabled":true,"name":"local-network","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"13":{"id":13,"position":12,"enabled":true,"name":"outbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"14":{"id":14,"position":13,"enabled":true,"name":"inbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"15":{"id":15,"position":14,"enabled":true,"name":"apply-candidate-acl","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"19":{"id":19,"position":15,"enabled":true,"name":"timer-name","value":"soft","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":2,"enabled":true,"name":"force-register-domain","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":7,"enabled":true,"name":"secure-combined","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":3,"enabled":true,"name":"secure-chain","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":4,"enabled":true,"name":"userauth","value":"true","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":5,"enabled":true,"name":"blind-reg","value":"false","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":6,"enabled":true,"name":"mcast-ip","value":"224.1.1.1","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":8,"enabled":true,"name":"mcast-port","value":"1337","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "[Config][Verto][Profile][Parameters][Get]":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVertoProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Settings][Update]","data":{"token":"example-token","param":{"id":1,"name":"debug","value":"1"}}}
	//Response:{"MessageType":"[Config][Verto][Settings][Update]","data":{"id":1,"position":1,"enabled":true,"name":"debug","value":"1","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Settings][Update]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Setting][Switch]","data":{"token":"example-token","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"[Config][Verto][Setting][Switch]","data":{"id":1,"position":1,"enabled":false,"name":"debug","value":"1","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Setting][Switch]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Setting][Add]","data":{"token":"example-token","param":{"name":"param","value":"0"}}}
	//Response:{"MessageType":"[Config][Verto][Setting][Add]","data":{"id":5,"position":2,"enabled":true,"name":"param","value":"0","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Setting][Add]":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVertoSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVertoSetting{}))}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Setting][Del]","data":{"token":"example-token","param":{"id":5}}}
	//Response:{"MessageType":"[Config][Verto][Setting][Del]","data":{"id":5,"position":2,"enabled":true,"name":"param","value":"0","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Setting][Del]":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVertoSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Param][Add]","data":{"token":"example-token","id":1,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Param][Add]","data":{"id":39,"position":16,"enabled":true,"name":"paramn","value":"paramv","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Param][Add]":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVertoProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Secure: msg.Param.Secure, Enabled: true, Parent: &altStruct.ConfigVertoProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"MoveVertoProfileParameter","data":{"token":"example-token","previous_index":16,"current_index":14,"id":39}}
	//Response:{"MessageType":"MoveVertoProfileParameter","data":{"1":{"id":1,"position":1,"enabled":true,"name":"bind-local","value":"45.61.54.76:8081","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":9,"enabled":true,"name":"rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":10,"enabled":true,"name":"ext-rtp-ip","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":11,"enabled":true,"name":"local-network","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"13":{"id":13,"position":12,"enabled":true,"name":"outbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"14":{"id":14,"position":13,"enabled":true,"name":"inbound-codec-string","value":"opus,h264,vp8","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"15":{"id":15,"position":15,"enabled":true,"name":"apply-candidate-acl","value":"localnet.auto","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"19":{"id":19,"position":16,"enabled":true,"name":"timer-name","value":"soft","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":2,"enabled":true,"name":"force-register-domain","value":"45.61.54.76","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"39":{"id":39,"position":14,"enabled":true,"name":"paramn","value":"paramv","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":7,"enabled":true,"name":"secure-combined","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":3,"enabled":true,"name":"secure-chain","value":"/etc/freeswitch/tls/wss.pem","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":4,"enabled":true,"name":"userauth","value":"true","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":5,"enabled":true,"name":"blind-reg","value":"false","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":6,"enabled":true,"name":"mcast-ip","value":"224.1.1.1","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":8,"enabled":true,"name":"mcast-port","value":"1337","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "MoveVertoProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfileParameter{Id: msg.Id, Position: msg.CurrentIndex}, []string{"Position"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Param][Del]","data":{"token":"example-token","param":{"id":39}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Param][Del]","data":{"id":39,"position":14,"enabled":true,"name":"paramn","value":"paramv","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Param][Del]":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVertoProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Param][Switch]","data":{"token":"example-token","param":{"id":19,"enabled":false}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Param][Switch]","data":{"id":19,"position":16,"enabled":false,"name":"timer-name","value":"soft","secure":"","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Param][Switch]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Update]","data":{"token":"example-token","param":{"id":19,"name":"timer-name","value":"hard","secure":""}}}
	//Response:{"MessageType":"[Config][Verto][Profile][Update]","data":{"id":19,"position":19,"enabled":true,"name":"timer-name","value":"hard","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null,"started":false,"state":"","uri":""}}}
	//Errors:
	case "[Config][Verto][Profile][Update]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value, Secure: msg.Param.Secure}, []string{"Name", "Value", "secure"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Add]","data":{"token":"example-token","name":"new_profile"}}
	//Response:{"MessageType":"[Config][Verto][Profile][Add]","data":{"id":4,"position":3,"enabled":true,"name":"new_profile","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Add]":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVertoProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVertoProfile{}))}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Rename]","data":{"token":"example-token","id":4,"name":"new_profile2"}}
	//Response:{"MessageType":"[Config][Verto][Profile][Rename]","data":{"id":4,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Rename]":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVertoProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"[Config][Verto][Profile][Del]","data":{"token":"example-token","id":4}}
	//Response:{"MessageType":"[Config][Verto][Profile][Del]","data":{"id":4,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":46,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "[Config][Verto][Profile][Del]":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVertoProfile{Id: msg.Id}, onlyAdminGroup())
	//### Callcenter
	//Request:{"event":"GetCallcenterQueues","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetCallcenterQueues","data":{"2":{"id":2,"position":2,"enabled":true,"name":"ddaaaaw","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ggdsf","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"a","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCallcenterQueues":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCallcenterQueue{}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterQueuesParams","data":{"token":"example-token","id":2}}
	//Response:{"MessageType":"GetCallcenterQueuesParams","data":{"1":{"id":1,"position":1,"enabled":true,"name":"ddd","value":"ddd","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetCallcenterQueuesParams":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCallcenterQueueParameter{}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterSettings","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetCallcenterSettings","data":{"1":{"id":1,"position":1,"enabled":true,"name":"qqq","value":"qqq","description":"qqq","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetCallcenterSettings":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigCallcenterSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterSettings","data":{"token":"example-token","param":{"id":1,"name":"qqq2","value":"qqq2","description":"qqq"}}}
	//Response:{"MessageType":"UpdateCallcenterSettings","data":{"id":1,"position":1,"enabled":true,"name":"qqq2","value":"qqq2","description":"qqq","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateCallcenterSettings":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchCallcenterSettings","data":{"token":"example-token","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchCallcenterSettings","data":{"id":1,"position":1,"enabled":false,"name":"qqq2","value":"qqq2","description":"qqq","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchCallcenterSettings":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterSettings","data":{"token":"example-token","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCallcenterSettings","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCallcenterSettings":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCallcenterSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCallcenterSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterSettings","data":{"token":"example-token","param":{"id":5}}}
	//Response:{"MessageType":"DelCallcenterSettings","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCallcenterSettings":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCallcenterSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterQueueParam","data":{"token":"example-token","id":2,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddCallcenterQueueParam","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddCallcenterQueueParam":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCallcenterQueueParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigCallcenterQueue{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterQueueParam","data":{"token":"example-token","param":{"id":5}}}
	//Response:{"MessageType":"DelCallcenterQueueParam","data":{"id":5,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelCallcenterQueueParam":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCallcenterQueueParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchCallcenterQueueParam","data":{"token":"example-token","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchCallcenterQueueParam","data":{"id":1,"position":1,"enabled":false,"name":"ddd","value":"ddd","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchCallcenterQueueParam":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterQueueParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterQueueParam","data":{"token":"example-token","param":{"id":1,"name":"new_param","value":"new_value"}}}
	//Response:{"MessageType":"UpdateCallcenterQueueParam","data":{"id":1,"position":1,"enabled":true,"name":"new_param","value":"new_value","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateCallcenterQueueParam":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterQueueParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterQueue","data":{"token":"example-token","name":"new_queue"}}
	//Response:{"MessageType":"AddCallcenterQueue","data":{"id":5,"position":5,"enabled":true,"name":"new_queue","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddCallcenterQueue":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigCallcenterQueue{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCallcenterQueue{}))}, onlyAdminGroup())
	//Request:{"event":"RenameCallcenterQueue","data":{"token":"example-token","id":5,"name":"new_queue2"}}
	//Response:{"MessageType":"RenameCallcenterQueue","data":{"id":5,"position":5,"enabled":true,"name":"new_queue2","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "RenameCallcenterQueue":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigCallcenterQueue{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterQueue","data":{"token":"example-token","id":5}}
	//Response:{"MessageType":"DelCallcenterQueue","data":{"id":5,"position":5,"enabled":true,"name":"new_queue2","description":"","parent":{"id":6,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelCallcenterQueue":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigCallcenterQueue{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"ImportCallcenterAgentsAndTiers","data":{"token":"example-token"}}
	//Response:{"MessageType":"ImportCallcenterAgentsAndTiers","data":{"callcenter_agents":{"items":null,"total":0},"callcenter_tiers":{"items":[{"id":3,"agent":"1007@45.61.54.76","queue":"n","level":4,"position":4,"state":"Ready"}],"total":1}}}
	//Errors:
	case "ImportCallcenterAgentsAndTiers":
		//TODO: replace
		getUser(msg, ImportCallcenterAgentsAdnTiers, onlyAdminGroup())
		msg.DBRequest = mainStruct.DBRequest{Limit: 25}
		resp1 := getUserForConfig(msg, getConfig, &altStruct.Agent{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.Tier{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event,
			Data: struct {
				S interface{} `json:"callcenter_agents"`
				P interface{} `json:"callcenter_tiers"`
			}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"GetCallcenterAgents","data":{"token":"example-token","db_request":{"limit":25,"offset":0,"filters":[],"order":{"fields":[],"desc":false}}}}
	//Response:{"MessageType":"GetCallcenterAgents","data":{"items":[{"id":34,"name":"30","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":3,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10},{"id":32,"name":"28","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":7,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10},{"id":6,"name":"2","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":4,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10},{"id":7,"name":"1000@45.61.54.76","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":0,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}],"total":30}}
	//Errors:
	case "GetCallcenterAgents":
		resp = getUserForConfig(msg, getConfig, &altStruct.Agent{}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterAgent","data":{"token":"example-token","name":"new_agent"}}
	//Response:{"MessageType":"AddCallcenterAgent","data":{"id":35,"name":"new_agent","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":0,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}}
	//Errors:
	case "AddCallcenterAgent":
		resp = getUserForConfig(msg, setConfig, &altStruct.Agent{
			Name:              msg.Name,
			Type:              "callback",
			System:            "single_box",
			InstanceId:        "single_box",
			Status:            "On Break",
			State:             "Waiting",
			WrapUpTime:        10,
			ReadyTime:         10,
			BusyDelayTime:     10,
			NoAnswerDelayTime: 10,
		}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterAgent","data":{"token":"example-token","param":{"id":37,"name":"max_no_answer","value":"5"}}}
	//Response:{"MessageType":"UpdateCallcenterAgent","data":{"id":37,"name":"agent@domain","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":5,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}}
	//Errors:
	case "UpdateCallcenterAgent":
		resp = getCallcenterAgents(msg)
		/*
					switch name {
					//Title:
			//Request:
			//Response:
			//Errors:
			case "state":
						eventChannel <- &map[int64]*mainStruct.Agent{agent.Id: agent}
					//Title:
			//Request:
			//Response:
			//Errors:
			case "status":
						agent.LastStatusChange = time.Now().Unix()
						eventChannel <- &map[int64]*mainStruct.Agent{agent.Id: agent}
					}
		*/
		// resp = getUser(msg, UpdateCallcenterAgent, onlyAdminGroup())
	//Request:{"event":"DelCallcenterAgent","data":{"token":"example-token","id":3}}
	//Response:{"MessageType":"DelCallcenterAgent","data":{"id":37,"name":"agent@domain","type":"callback","system":"single_box","instance_id":"single_box","uuid":"","contact":"","status":"On Break","state":"Waiting","max_no_answer":5,"wrap_up_time":10,"reject_delay_time":0,"busy_delay_time":10,"no_answer_delay_time":10,"last_bridge_start":0,"last_bridge_end":0,"last_offered_call":0,"last_status_change":0,"no_answer_count":0,"calls_answered":0,"talk_time":0,"ready_time":10}}
	//Errors:
	case "DelCallcenterAgent":
		resp = getUserForConfig(msg, delConfig, &altStruct.Agent{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterTiers","data":{"token":"example-token","db_request":{"limit":25,"offset":0,"filters":[],"order":{"fields":[],"desc":false}}}}
	//Response:{"MessageType":"GetCallcenterTiers","data":{"items":[{"id":4,"agent":"agent","queue":"ddaaaaw","level":1,"position":1,"state":"Ready"}],"total":1}}
	//Errors:
	case "GetCallcenterTiers":
		resp = getUserForConfig(msg, getConfig, &altStruct.Tier{}, onlyAdminGroup())
	//Request:{"event":"AddCallcenterTier","data":{"token":"example-token","id":2,"name":"new_agent"}}
	//Response:{"MessageType":"AddCallcenterTier","data":{"id":5,"agent":"new_agent","queue":"ddaaaaw","level":1,"position":1,"state":"Ready"}}
	//Errors:
	case "AddCallcenterTier":
		queueI, err := intermediateDB.GetByIdArg(&altStruct.ConfigCallcenterQueue{}, msg.Id)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: msg.Event}
		}
		queue, ok := queueI.(altStruct.ConfigCallcenterQueue)
		if !ok {
			return webStruct.UserResponse{Error: "queue not found", MessageType: msg.Event}
		}
		resp = getUserForConfig(msg, setConfig, &altStruct.Tier{Queue: queue.Name, Agent: msg.Name, State: "Ready", Position: 1, Level: 1}, onlyAdminGroup())
	//Request:{"event":"UpdateCallcenterTier","data":{"token":"example-token","param":{"id":5,"name":"level","value":"2"}}}
	//Response:{"MessageType":"UpdateCallcenterTier","data":{"id":5,"agent":"new_agent","queue":"ddaaaaw","level":2,"position":7,"state":"Ready"}}
	//Errors:
	case "UpdateCallcenterTier":
		resp = getCallcenterTiers(msg)
	//Request:{"event":"DelCallcenterTier","data":{"token":"example-token","id":5}}
	//Response:{"MessageType":"DelCallcenterTier","data":{"id":5,"agent":"new_agent","queue":"ddaaaaw","level":2,"position":7,"state":"Ready"}}
	//Errors:
	case "DelCallcenterTier":
		resp = getUserForConfig(msg, delConfig, &altStruct.Tier{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetCallcenterMembers","data":{"token":"example-token","db_request":{"limit":25,"offset":0,"filters":[],"order":{"fields":[],"desc":false}}}}
	//Response:{"MessageType":"GetCallcenterMembers","data":{"items":null,"total":0}}
	//Errors:
	case "GetCallcenterMembers":
		resp = getUserForConfig(msg, getConfig, &altStruct.Member{}, onlyAdminGroup())
	//Request:{"event":"DelCallcenterMember","data":{"token":"example-token","id":5}}
	//Response:{"MessageType":"DelCallcenterMember","data":{"id":5}}
	//Errors:
	case "DelCallcenterMember":
		resp = getUserForConfig(msg, delConfig, &altStruct.Member{Uuid: msg.Uuid}, onlyAdminGroup())
	//Request:{"event":"SendCallcenterCommand","data":{"token":"example-token","name":"load","id":2}}
	//Response:{"MessageType":"SendCallcenterCommand"}
	//Errors:
	case "SendCallcenterCommand":
		resp = getUser(msg, runCallcenterQueueCommand, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case webStruct.SubscribeCallcenterAgents:
		//TODO: replace
		resp = getUserForConfig(msg, getConfig, &altStruct.Agent{}, onlyAdminGroup())
	//Request:
	//Response:
	//Errors:
	case webStruct.SubscribeCallcenterTiers:
		//TODO: replace
		resp = getUserForConfig(msg, getConfig, &altStruct.Tier{}, onlyAdminGroup())
	//### LCR
	// Migrated to the WebSocket handler registry.
	//### Shout/Redis/Nibblebill/Db/Memcache/Avmd/TtsCommandline/CdrMongodb
	// Migrated to the WebSocket handler registry.
	//### HTTP Cache
	//Request:{"event":"GetHttpCache","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetHttpCache","data":{"1":{"id":1,"position":1,"enabled":true,"name":"enable-file-formats","value":"false","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"max-urls","value":"10000","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"location","value":"/var/cache/freeswitch","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"default-max-age","value":"86400","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"prefetch-thread-count","value":"8","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"prefetch-queue-size","value":"100","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"ssl-cacert","value":"/etc/freeswitch/tls/cacert.pem","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"ssl-verifypeer","value":"true","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"ssl-verifyhost","value":"true","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetHttpCache":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheSetting{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"settings"`
			P interface{} `json:"profiles"`
		}{S: resp1.Data, P: resp2.Data}}
	//Request:{"event":"UpdateHttpCacheParameter","data":{"token":"example-token","param":{"id":9,"name":"ssl-verifyhost","value":"false"}}}
	//Response:{"MessageType":"UpdateHttpCacheParameter","data":{"id":9,"position":9,"enabled":true,"name":"ssl-verifyhost","value":"false","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateHttpCacheParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchHttpCacheParameter","data":{"token":"example-token","param":{"id":9,"enabled":false}}}
	//Response:{"MessageType":"SwitchHttpCacheParameter","data":{"id":9,"position":9,"enabled":false,"name":"ssl-verifyhost","value":"false","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchHttpCacheParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddHttpCacheParameter","data":{"token":"example-token","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddHttpCacheParameter","data":{"id":12,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddHttpCacheParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigHttpCacheSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigHttpCacheSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelHttpCacheParameter","data":{"token":"example-token","param":{"id":12}}}
	//Response:{"MessageType":"DelHttpCacheParameter","data":{"id":12,"position":10,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":23,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelHttpCacheParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigHttpCacheSetting{Id: msg.Param.Id}, onlyAdminGroup())
	case "GetHttpCacheProfile":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfile{}, onlyAdminGroup())
	case "AddHttpCacheProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigHttpCacheProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigHttpCacheProfile{}))}, onlyAdminGroup())
	case "RenameHttpCacheProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	case "DelHttpCacheProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigHttpCacheProfile{Id: msg.Id}, onlyAdminGroup())
	case "GetHttpCacheProfileParameters":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfileDomain{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfileAzureBlob{}, onlyAdminGroup())
		resp3 := getUserForConfig(msg, getConfig, &altStruct.ConfigHttpCacheProfileAWSS3{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			S interface{} `json:"domains"`
			P interface{} `json:"azure"`
			R interface{} `json:"aws_s3"`
		}{S: resp1.Data, P: resp2.Data, R: resp3.Data}}
	case "AddHttpCacheProfileDomain":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigHttpCacheProfileDomain{Name: msg.Param.Name, Enabled: true, Parent: &altStruct.ConfigHttpCacheProfile{Id: msg.Id}}, onlyAdminGroup())
	case "DelHttpCacheProfileDomain":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigHttpCacheProfileDomain{Id: msg.Param.Id}, onlyAdminGroup())
	case "SwitchHttpCacheProfileDomain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileDomain{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	case "UpdateHttpCacheProfileDomain":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileDomain{Id: msg.Param.Id, Name: msg.Param.Name}, []string{"Name"}}, onlyAdminGroup())
	case "UpdateHttpCacheProfileAws":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileAWSS3{Id: msg.AwsS3.Id,
			AccessKeyId:     msg.AwsS3.AccessKeyId,
			SecretAccessKey: msg.AwsS3.SecretAccessKey,
			BaseDomain:      msg.AwsS3.BaseDomain,
			Region:          msg.AwsS3.Region,
			Expires:         msg.AwsS3.Expires,
		}, []string{"AccessKeyId", "SecretAccessKey", "BaseDomain", "Region", "Expires"}}, onlyAdminGroup())
	case "UpdateHttpCacheProfileAzure":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigHttpCacheProfileAzureBlob{Id: msg.Azure.Id,
			SecretAccessKey: msg.Azure.SecretAccessKey,
		}, []string{"SecretAccessKey"}}, onlyAdminGroup())

	//### Opus/Python/Alsa/Amr/Amrwb/Cepstral/Cidlookup/Curl/DialplanDirectory/Easyroute/ErlangEvent/EventMulticast/Fax/Lua/Mongo/Msrp/Oreka/Perl/Pocketsphinx/SangomaCodec
	// Migrated to the WebSocket handler registry.
	//### Sndfile/XML CDR/XML RPC/Zeroconf
	// Migrated to the WebSocket handler registry.
	//### Post Load Switch
	// Migrated to the WebSocket handler registry.
	//### Distributor
	//Request:{"event":"GetDistributorConfig","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetDistributorConfig","data":{"1":{"id":1,"position":1,"enabled":true,"name":"test","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetDistributorConfig":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDistributorList{}, onlyAdminGroup())
	//Request:{"event":"AddDistributorList","data":{"token":"example-token","name":"new_list"}}
	//Response:{"MessageType":"AddDistributorList","data":{"id":7,"position":2,"enabled":true,"name":"new_list","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddDistributorList":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDistributorList{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigDistributorList{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateDistributorList","data":{"token":"example-token","id":7,"name":"new_list2"}}
	//Response:{"MessageType":"UpdateDistributorList","data":{"id":7,"position":2,"enabled":true,"name":"new_list2","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateDistributorList":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDistributorList{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelDistributorList","data":{"token":"example-token","id":7}}
	//Response:{"MessageType":"DelDistributorList","data":{"id":7,"position":2,"enabled":true,"name":"new_list2","description":"","parent":{"id":17,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelDistributorList":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDistributorList{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetDistributorNodes","data":{"token":"example-token","id":1}}
	//Response:{"MessageType":"GetDistributorNodes","data":{"1":{"id":1,"position":1,"enabled":true,"name":"foo1","weight":"1","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"foo2","weight":"9","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetDistributorNodes":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigDistributorListNode{}, onlyAdminGroup())
	//Request:{"event":"AddDistributorNode","data":{"token":"example-token","id":1,"distributor_node":{"name":"paramn","weight":"paramv"}}}
	//Response:{"MessageType":"AddDistributorNode","data":{"id":15,"position":3,"enabled":true,"name":"paramn","weight":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddDistributorNode":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigDistributorListNode{Name: msg.DistributorNode.Name, Weight: msg.DistributorNode.Weight, Enabled: true, Parent: &altStruct.ConfigDistributorList{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelDistributorNode","data":{"token":"example-token","distributor_node":{"id":15}}}
	//Response:{"MessageType":"DelDistributorNode","data":{"id":15,"position":3,"enabled":true,"name":"paramn","weight":"paramv","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelDistributorNode":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigDistributorListNode{Id: msg.DistributorNode.Id}, onlyAdminGroup())
	//Request:{"event":"UpdateDistributorNode","data":{"token":"example-token","distributor_node":{"id":2,"name":"foo2","weight":"2"}}}
	//Response:{"MessageType":"UpdateDistributorNode","data":{"id":2,"position":2,"enabled":true,"name":"foo2","weight":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateDistributorNode":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDistributorListNode{Id: msg.DistributorNode.Id, Name: msg.DistributorNode.Name, Weight: msg.DistributorNode.Weight}, []string{"Name", "Weight"}}, onlyAdminGroup())
	//Request:{"event":"SwitchDistributorNode","data":{"token":"example-token","distributor_node":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchDistributorNode","data":{"id":2,"position":2,"enabled":false,"name":"foo2","weight":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchDistributorNode":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigDistributorListNode{Id: msg.DistributorNode.Id, Enabled: msg.DistributorNode.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//### Directory/Fifo
	// Migrated to the WebSocket handler registry.
	//### Opal/Osp/Unicall
	// Migrated to the WebSocket handler registry.
	//### Conference
	//Request:{"event":"GetConference","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetConference","data":{"conference_rooms":{"1":{"id":1,"position":1,"enabled":true,"name":"3001@45.61.54.76","status":"FreeSWITCH","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"conference_profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"a1","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"wideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ultrawideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"cdquality","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"video-mcu-stereo","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"video-mcu-stereo-720","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"video-mcu-stereo-480","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"video-mcu-stereo-320","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"sla","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"conference_caller_control_groups":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"s2","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}},"conference_chat_permissions_profiles":{"1":{"id":1,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"a1","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"wideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"ultrawideband","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"cdquality","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"video-mcu-stereo","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"video-mcu-stereo-720","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"video-mcu-stereo-480","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"video-mcu-stereo-320","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"sla","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}}
	//Errors:
	case "GetConference":
		resp1 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceAdvertiseRoom{}, onlyAdminGroup())
		resp2 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceProfile{}, onlyAdminGroup())
		resp3 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceCallerControlGroup{}, onlyAdminGroup())
		resp4 := getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceChatPermissionProfile{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			A interface{} `json:"conference_rooms"`
			B interface{} `json:"conference_profiles"`
			C interface{} `json:"conference_caller_control_groups"`
			D interface{} `json:"conference_chat_permissions_profiles"`
		}{A: resp1.Data, B: resp2.Data, C: resp3.Data, D: resp4.Data}}
	//Request:{"event":"UpdateConferenceRoom","data":{"token":"example-token","param":{"id":1,"enabled":true,"name":"3001@45.61.54.76","status":"FreeSWITCH2"}}}
	//Response:{"MessageType":"UpdateConferenceRoom","data":{"id":1,"position":1,"enabled":true,"name":"3001@45.61.54.76","status":"","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceRoom":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceAdvertiseRoom{Id: msg.Param.Id, Name: msg.Param.Name, Status: msg.Param.Value}, []string{"Name", "Status"}}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceRoom","data":{"token":"example-token","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceRoom","data":{"id":1,"position":1,"enabled":false,"name":"3001@45.61.54.76","status":"","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchConferenceRoom":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceAdvertiseRoom{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceRoom","data":{"token":"example-token","param":{"name":"room","value":"status"}}}
	//Response:{"MessageType":"AddConferenceRoom","data":{"id":7,"position":2,"enabled":true,"name":"room","status":"status","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceRoom":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceAdvertiseRoom{Name: msg.Param.Name, Status: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceAdvertiseRoom{}))}, onlyAdminGroup())
	//Request:{"event":"DelConferenceRoom","data":{"token":"example-token","param":{"id":7}}}
	//Response:{"MessageType":"DelConferenceRoom","data":{"id":7,"position":2,"enabled":true,"name":"room","status":"status","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceRoom":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceAdvertiseRoom{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetConferenceCallerControls","data":{"token":"example-token","id":1}}
	//Response:{"MessageType":"GetConferenceCallerControls","data":{"1":{"id":1,"position":1,"enabled":true,"action":"mute","digits":"0","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":10,"enabled":true,"action":"vol listen zero","digits":"5","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":11,"enabled":true,"action":"vol listen dn","digits":"4","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":12,"enabled":true,"action":"hangup","digits":"#","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"action":"deaf mute","digits":"*","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":3,"enabled":true,"action":"energy up","digits":"9","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":4,"enabled":true,"action":"energy equ","digits":"8","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":5,"enabled":true,"action":"energy dn","digits":"7","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":6,"enabled":true,"action":"vol talk up","digits":"3","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":7,"enabled":true,"action":"vol talk zero","digits":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":8,"enabled":true,"action":"vol talk dn","digits":"1","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":9,"enabled":true,"action":"vol listen up","digits":"6","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetConferenceCallerControls":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceCallerControlGroupControl{}, onlyAdminGroup())
	//Request:{"event":"AddConferenceCallerControl","data":{"token":"example-token","id":1,"param":{"name":"action","value":"2"}}}
	//Response:{"MessageType":"AddConferenceCallerControl","data":{"id":19,"position":13,"enabled":true,"action":"action","digits":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddConferenceCallerControl":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceCallerControlGroupControl{Action: msg.Param.Name, Digits: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceCallerControlGroup{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceCallerControl","data":{"token":"example-token","param":{"id":19}}}
	//Response:{"MessageType":"DelConferenceCallerControl","data":{"id":19,"position":13,"enabled":true,"action":"action","digits":"2","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelConferenceCallerControl":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceCallerControlGroupControl{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceCallerControl","data":{"token":"example-token","param":{"id":12,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceCallerControl","data":{"id":12,"position":12,"enabled":false,"action":"hangup","digits":"#","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchConferenceCallerControl":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceCallerControlGroupControl{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceCallerControl","data":{"token":"example-token","param":{"id":11,"name":"vol listen dn","value":"4"}}}
	//Response:{"MessageType":"UpdateConferenceCallerControl","data":{"id":11,"position":11,"enabled":true,"action":"vol listen dn","digits":"4","description":"","parent":{"id":1,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateConferenceCallerControl":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceCallerControlGroupControl{Id: msg.Param.Id, Action: msg.Param.Name, Digits: msg.Param.Value}, []string{"Action", "Digits"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceCallerControlGroup","data":{"token":"example-token","name":"new_group"}}
	//Response:{"MessageType":"AddConferenceCallerControlGroup","data":{"id":4,"position":3,"enabled":true,"name":"new_group","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceCallerControlGroup":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceCallerControlGroup{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceCallerControlGroup{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceCallerControlGroup","data":{"token":"example-token","id":4,"name":"new_group2"}}
	//Response:{"MessageType":"UpdateConferenceCallerControlGroup","data":{"id":4,"position":3,"enabled":true,"name":"new_group2","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceCallerControlGroup":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceCallerControlGroup{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceCallerControlGroup","data":{"token":"example-token","id":4}}
	//Response:{"MessageType":"DelConferenceCallerControlGroup","data":{"id":4,"position":3,"enabled":true,"name":"new_group2","description":"","parent":{"id":11,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceCallerControlGroup":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceCallerControlGroup{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetConferenceProfileParameters","data":{"token":"example-token","id":10}}
	//Response:{"MessageType":"GetConferenceProfileParameters","data":{"204":{"id":204,"position":1,"enabled":true,"name":"domain","value":"45.61.54.76","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"205":{"id":205,"position":2,"enabled":true,"name":"rate","value":"8000","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"206":{"id":206,"position":3,"enabled":true,"name":"interval","value":"20","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"207":{"id":207,"position":4,"enabled":true,"name":"energy-level","value":"100","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"208":{"id":208,"position":5,"enabled":true,"name":"muted-sound","value":"conference/conf-muted.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"209":{"id":209,"position":6,"enabled":true,"name":"unmuted-sound","value":"conference/conf-unmuted.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"210":{"id":210,"position":7,"enabled":true,"name":"alone-sound","value":"conference/conf-alone.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"211":{"id":211,"position":8,"enabled":true,"name":"moh-sound","value":"local_stream://moh","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"212":{"id":212,"position":9,"enabled":true,"name":"enter-sound","value":"tone_stream://%(200,0,500,600,700)","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"213":{"id":213,"position":10,"enabled":true,"name":"exit-sound","value":"tone_stream://%(500,0,300,200,100,50,25)","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"214":{"id":214,"position":11,"enabled":true,"name":"kicked-sound","value":"conference/conf-kicked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"215":{"id":215,"position":12,"enabled":true,"name":"locked-sound","value":"conference/conf-locked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"216":{"id":216,"position":13,"enabled":true,"name":"is-locked-sound","value":"conference/conf-is-locked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"217":{"id":217,"position":14,"enabled":true,"name":"is-unlocked-sound","value":"conference/conf-is-unlocked.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"218":{"id":218,"position":15,"enabled":true,"name":"pin-sound","value":"conference/conf-pin.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"219":{"id":219,"position":16,"enabled":true,"name":"bad-pin-sound","value":"conference/conf-bad-pin.wav","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"220":{"id":220,"position":17,"enabled":true,"name":"caller-id-name","value":"FreeSWITCH","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"221":{"id":221,"position":18,"enabled":true,"name":"caller-id-number","value":"0000000000","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}},"222":{"id":222,"position":19,"enabled":true,"name":"comfort-noise","value":"true","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetConferenceProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"AddConferenceProfileParameter","data":{"token":"example-token","id":10,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddConferenceProfileParameter","data":{"id":407,"position":20,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddConferenceProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceProfileParameter","data":{"token":"example-token","param":{"id":407}}}
	//Response:{"MessageType":"DelConferenceProfileParameter","data":{"id":407,"position":20,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelConferenceProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceProfileParameter","data":{"token":"example-token","param":{"id":222,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceProfileParameter","data":{"id":222,"position":19,"enabled":false,"name":"comfort-noise","value":"true","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchConferenceProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceProfileParameter","data":{"token":"example-token","param":{"id":222,"name":"comfort-noise","value":"false"}}}
	//Response:{"MessageType":"UpdateConferenceProfileParameter","data":{"id":222,"position":19,"enabled":true,"name":"comfort-noise","value":"false","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateConferenceProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceProfile","data":{"token":"example-token","name":"new_profile"}}
	//Response:{"MessageType":"AddConferenceProfile","data":{"id":19,"position":10,"enabled":true,"name":"new_profile","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceProfile","data":{"token":"example-token","id":19,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateConferenceProfile","data":{"id":19,"position":10,"enabled":true,"name":"new_profile2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceProfile","data":{"token":"example-token","id":1}}
	//Response:{"MessageType":"DelConferenceProfile","data":{"id":19,"position":10,"enabled":true,"name":"new_profile2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceProfile{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetConferenceChatPermissionUsers","data":{"token":"example-token","id":10}}
	//Response:{"MessageType":"GetConferenceChatPermissionUsers","data":{"1":{"id":1,"position":1,"enabled":true,"name":"paramn","commands":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetConferenceChatPermissionUsers":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceChatPermissionProfileUser{}, onlyAdminGroup())
	//Request:{"event":"AddConferenceChatPermissionUser","data":{"token":"example-token","id":10,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":true,"name":"paramn","commands":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddConferenceChatPermissionUser":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceChatPermissionProfileUser{Name: msg.Param.Name, Commands: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceChatPermissionProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceChatPermissionUser","data":{"token":"example-token","param":{"id":1}}}
	//Response:{"MessageType":"DelConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":true,"name":"paramn2","commands":"","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelConferenceChatPermissionUser":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceChatPermissionProfileUser{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchConferenceChatPermissionUser","data":{"token":"example-token","param":{"id":1,"enabled":false}}}
	//Response:{"MessageType":"SwitchConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":false,"name":"paramn","commands":"paramv","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchConferenceChatPermissionUser":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceChatPermissionProfileUser{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceChatPermissionUser","data":{"token":"example-token","param":{"id":1,"name":"paramn2","commands":"paramv2"}}}
	//Response:{"MessageType":"UpdateConferenceChatPermissionUser","data":{"id":1,"position":1,"enabled":true,"name":"paramn2","commands":"","description":"","parent":{"id":10,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateConferenceChatPermissionUser":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceChatPermissionProfileUser{Id: msg.Param.Id, Name: msg.Param.Name, Commands: msg.Param.Value}, []string{"Name", "Commands"}}, onlyAdminGroup())
	//Request:{"event":"AddConferenceChatPermission","data":{"token":"example-token","name":"new_permission"}}
	//Response:{"MessageType":"AddConferenceChatPermission","data":{"id":20,"position":10,"enabled":true,"name":"new_permission","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddConferenceChatPermission":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceChatPermissionProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigConferenceChatPermissionProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateConferenceChatPermission","data":{"token":"example-token","id":20,"name":"new_permission2"}}
	//Response:{"MessageType":"UpdateConferenceChatPermission","data":{"id":20,"position":10,"enabled":true,"name":"new_permission2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateConferenceChatPermission":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceChatPermissionProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelConferenceChatPermission","data":{"token":"example-token","id":2}}
	//Response:{"MessageType":"DelConferenceChatPermission","data":{"id":20,"position":10,"enabled":true,"name":"new_permission2","description":"","parent":{"id":57,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelConferenceChatPermission":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceChatPermissionProfile{Id: msg.Id}, onlyAdminGroup())

	case "GetConferenceLayouts":
		msg.Name = "layout"
		resp1 := getUserForConfig(msg, getConferenceLayoutsConfig, &altStruct.ConfigConferenceLayout{}, onlyAdminGroup())
		msg.Name = "group"
		resp2 := getUserForConfig(msg, getConferenceLayoutsConfig, &altStruct.ConfigConferenceLayoutGroup{}, onlyAdminGroup())
		resp = webStruct.UserResponse{MessageType: msg.Event, Data: struct {
			A interface{} `json:"conference_layouts"`
			B interface{} `json:"conference_layouts_groups"`
		}{A: resp1.Data, B: resp2.Data}}
	case "UpdateConferenceLayout":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayout{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	case "UpdateConferenceLayout3D":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayout{Id: msg.Id, Auto3dPosition: msg.Value}, []string{"Auto3dPosition"}}, onlyAdminGroup())
	case "SwitchConferenceLayout":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayout{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	case "AddConferenceLayout":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceLayout{Name: msg.Name, Enabled: true, Parent: getConferenceLayoutConfig()}, onlyAdminGroup())
	case "DelConferenceLayout":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceLayout{Id: msg.Id}, onlyAdminGroup())

	case "UpdateConferenceLayoutGroup":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayoutGroup{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	case "SwitchConferenceLayoutGroup":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayoutGroup{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	case "AddConferenceLayoutGroup":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceLayoutGroup{Name: msg.Name, Enabled: true, Parent: getConferenceLayoutConfig()}, onlyAdminGroup())
	case "DelConferenceLayoutGroup":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceLayoutGroup{Id: msg.Id}, onlyAdminGroup())

	case "GetConferenceLayoutGroupLayouts":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceLayoutGroupLayout{}, onlyAdminGroup())
	case "AddConferenceLayoutGroupLayout":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceLayoutGroupLayout{Body: msg.Value, Enabled: true, Parent: &altStruct.ConfigConferenceLayoutGroup{Id: msg.Id}}, onlyAdminGroup())
	case "DelConferenceLayoutGroupLayout":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceLayoutGroupLayout{Id: msg.Param.Id}, onlyAdminGroup())
	case "SwitchConferenceLayoutGroupLayout":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayoutGroupLayout{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	case "UpdateConferenceLayoutGroupLayout":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayoutGroupLayout{Id: msg.Param.Id, Body: msg.Param.Name}, []string{"Body"}}, onlyAdminGroup())

	case "GetConferenceLayoutImages":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigConferenceLayoutImage{}, onlyAdminGroup())
	case "AddConferenceLayoutImage":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigConferenceLayoutImage{
			X:             msg.LayoutImages.X,
			Y:             msg.LayoutImages.Y,
			Scale:         msg.LayoutImages.Scale,
			Floor:         msg.LayoutImages.Floor,
			FloorOnly:     msg.LayoutImages.FloorOnly,
			Hscale:        msg.LayoutImages.Hscale,
			Overlap:       msg.LayoutImages.Overlap,
			ReservationId: msg.LayoutImages.ReservationId,
			Zoom:          msg.LayoutImages.Zoom,
			Enabled:       true,
			Parent:        &altStruct.ConfigConferenceLayout{Id: msg.Id},
		}, onlyAdminGroup())
	case "DelConferenceLayoutImage":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigConferenceLayoutImage{Id: msg.Param.Id}, onlyAdminGroup())
	case "SwitchConferenceLayoutImage":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayoutImage{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	case "UpdateConferenceLayoutImage":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigConferenceLayoutImage{
			Id:            msg.LayoutImages.Id,
			X:             msg.LayoutImages.X,
			Y:             msg.LayoutImages.Y,
			Scale:         msg.LayoutImages.Scale,
			Floor:         msg.LayoutImages.Floor,
			FloorOnly:     msg.LayoutImages.FloorOnly,
			Hscale:        msg.LayoutImages.Hscale,
			Overlap:       msg.LayoutImages.Overlap,
			ReservationId: msg.LayoutImages.ReservationId,
			Zoom:          msg.LayoutImages.Zoom,
		}, []string{"X", "Y", "Scale", "Floor", "FloorOnly", "Hscale", "Overlap", "ReservationId", "Zoom"}}, onlyAdminGroup())

	//### Post Load Modules
	//Request:{"event":"GetPostLoadModules","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetPostLoadModules","data":{"1":{"id":1,"position":1,"enabled":false,"name":"mod_sofia","description":" ","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"10":{"id":10,"position":7,"enabled":false,"name":"mod_unicall","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"11":{"id":11,"position":8,"enabled":false,"name":"mod_xml_cdr","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"12":{"id":12,"position":9,"enabled":false,"name":"mod_xml_rpc","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"13":{"id":13,"position":10,"enabled":true,"name":"mod_shout","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"14":{"id":14,"position":11,"enabled":true,"name":"mod_pocketsphinx","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"15":{"id":15,"position":12,"enabled":true,"name":"mod_alsa","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"2":{"id":2,"position":2,"enabled":false,"name":"mod_amr","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"4":{"id":4,"position":3,"enabled":false,"name":"mod_db","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":4,"enabled":false,"name":"mod_verto","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"8":{"id":8,"position":5,"enabled":true,"name":"mod_voicemail","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"9":{"id":9,"position":6,"enabled":false,"name":"mod_zeroconf","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetPostLoadModules":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigPostLoadModule{}, onlyAdminGroup())
	//Request:{"event":"UpdatePostLoadModule","data":{"token":"example-token","param":{"id":13,"name":"mod_shout"}}}
	//Response:{"MessageType":"UpdatePostLoadModule","data":{"id":13,"position":10,"enabled":true,"name":"mod_shout","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdatePostLoadModule":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadModule{Id: msg.Param.Id, Name: msg.Param.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"SwitchPostLoadModule","data":{"token":"example-token","param":{"id":13,"enabled":false}}}
	//Response:{"MessageType":"SwitchPostLoadModule","data":{"id":13,"position":10,"enabled":false,"name":"mod_shout","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case webStruct.SwitchPostLoadModule:
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigPostLoadModule{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"GetVoicemailSettings","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetVoicemailSettings","data":{"2":{"id":2,"position":1,"enabled":true,"name":"dsfsdf2","value":"sdfsfs2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case webStruct.AddPostLoadModule:
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigPostLoadModule{Name: msg.Param.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigPostLoadModule{}))}, onlyAdminGroup())
	//Request:{"event":"DelPostLoadModule","data":{"token":"example-token","param":{"id":16}}}
	//Response:{"MessageType":"DelPostLoadModule","data":{"id":16,"position":13,"enabled":true,"name":"mod_fifo","description":"","parent":{"id":36,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelPostLoadModule":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigPostLoadModule{Id: msg.Param.Id}, onlyAdminGroup())
	//### Voicemail
	//Request:{"event":"GetVoicemailProfiles","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetVoicemailProfiles","data":{"2":{"id":2,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":2,"enabled":true,"name":"ccc","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetVoicemailSettings":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailSetting{}, onlyAdminGroup())
	//Request:{"event":"UpdateVoicemailSetting","data":{"token":"example-token","param":{"id":2,"name":"dsfsdf2","value":"sdfsfs2"}}}
	//Response:{"MessageType":"UpdateVoicemailSetting","data":{"id":2,"position":1,"enabled":true,"name":"dsfsdf2","value":"sdfsfs2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateVoicemailSetting":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailSetting{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
	//Request:{"event":"SwitchVoicemailSetting","data":{"token":"example-token","param":{"id":2,"enabled":false}}}
	//Response:{"MessageType":"SwitchVoicemailSetting","data":{"id":2,"position":1,"enabled":false,"name":"dsfsdf2","value":"sdfsfs2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "SwitchVoicemailSetting":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailSetting{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"AddVoicemailSetting","data":{"token":"example-token","param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddVoicemailSetting","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddVoicemailSetting":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVoicemailSetting{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVoicemailSetting{}))}, onlyAdminGroup())
	//Request:{"event":"DelVoicemailSetting","data":{"token":"example-token","param":{"id":4}}}
	//Response:{"MessageType":"DelVoicemailSetting","data":{"id":4,"position":2,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelVoicemailSetting":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVoicemailSetting{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"GetVoicemailProfiles","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetVoicemailProfiles","data":{"2":{"id":2,"position":1,"enabled":true,"name":"default","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}},"5":{"id":5,"position":2,"enabled":true,"name":"ccc","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}}
	//Errors:
	case "GetVoicemailProfiles":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailProfile{}, onlyAdminGroup())
	//Request:{"event":"AddVoicemailProfile","data":{"token":"example-token","name":"new_profile"}}
	//Response:{"MessageType":"AddVoicemailProfile","data":{"id":8,"position":3,"enabled":true,"name":"new_profile","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "AddVoicemailProfile":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVoicemailProfile{Name: msg.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigVoicemailProfile{}))}, onlyAdminGroup())
	//Request:{"event":"UpdateVoicemailProfile","data":{"token":"example-token","id":8,"name":"new_profile2"}}
	//Response:{"MessageType":"UpdateVoicemailProfile","data":{"id":8,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "UpdateVoicemailProfile":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailProfile{Id: msg.Id, Name: msg.Name}, []string{"Name"}}, onlyAdminGroup())
	//Request:{"event":"DelVoicemailProfile","data":{"token":"example-token","id":8}}
	//Response:{"MessageType":"DelVoicemailProfile","data":{"id":8,"position":3,"enabled":true,"name":"new_profile2","description":"","parent":{"id":47,"position":0,"enabled":false,"name":"","module":"","loaded":false,"unloadable":false,"parent":null}}}
	//Errors:
	case "DelVoicemailProfile":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVoicemailProfile{Id: msg.Id}, onlyAdminGroup())
	//Request:{"event":"GetVoicemailProfileParameters","data":{"token":"example-token","id":2}}
	//Response:{"MessageType":"GetVoicemailProfileParameters","data":{"1":{"id":1,"position":1,"enabled":true,"name":"file-extension","value":"wav","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"10":{"id":10,"position":10,"enabled":true,"name":"callback-context","value":"default","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"11":{"id":11,"position":11,"enabled":true,"name":"play-new-messages-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"12":{"id":12,"position":12,"enabled":true,"name":"play-saved-messages-key","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"13":{"id":13,"position":13,"enabled":true,"name":"login-keys","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"14":{"id":14,"position":14,"enabled":true,"name":"main-menu-key","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"15":{"id":15,"position":15,"enabled":true,"name":"config-menu-key","value":"5","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"16":{"id":16,"position":16,"enabled":true,"name":"record-greeting-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"17":{"id":17,"position":17,"enabled":true,"name":"choose-greeting-key","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"18":{"id":18,"position":18,"enabled":true,"name":"change-pass-key","value":"6","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"19":{"id":19,"position":19,"enabled":true,"name":"record-name-key","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"2":{"id":2,"position":2,"enabled":true,"name":"terminator-key","value":"#","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"20":{"id":20,"position":20,"enabled":true,"name":"record-file-key","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"21":{"id":21,"position":21,"enabled":true,"name":"listen-file-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"22":{"id":22,"position":22,"enabled":true,"name":"save-file-key","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"23":{"id":23,"position":23,"enabled":true,"name":"delete-file-key","value":"7","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"24":{"id":24,"position":24,"enabled":true,"name":"undelete-file-key","value":"8","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"25":{"id":25,"position":25,"enabled":true,"name":"email-key","value":"4","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"26":{"id":26,"position":26,"enabled":true,"name":"pause-key","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"27":{"id":27,"position":27,"enabled":true,"name":"restart-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"28":{"id":28,"position":28,"enabled":true,"name":"ff-key","value":"6","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"29":{"id":29,"position":29,"enabled":true,"name":"rew-key","value":"4","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"3":{"id":3,"position":3,"enabled":true,"name":"max-login-attempts","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"30":{"id":30,"position":30,"enabled":true,"name":"skip-greet-key","value":"#","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"31":{"id":31,"position":31,"enabled":true,"name":"previous-message-key","value":"1","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"32":{"id":32,"position":32,"enabled":true,"name":"next-message-key","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"33":{"id":33,"position":33,"enabled":true,"name":"skip-info-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"34":{"id":34,"position":34,"enabled":true,"name":"repeat-message-key","value":"0","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"35":{"id":35,"position":35,"enabled":true,"name":"record-silence-threshold","value":"200","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"36":{"id":36,"position":36,"enabled":true,"name":"record-silence-hits","value":"2","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"37":{"id":37,"position":37,"enabled":true,"name":"web-template-file","value":"web-vm.tpl","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"38":{"id":38,"position":38,"enabled":true,"name":"db-password-override","value":"false","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"39":{"id":39,"position":39,"enabled":true,"name":"allow-empty-password-auth","value":"true","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"4":{"id":4,"position":4,"enabled":true,"name":"digit-timeout","value":"10000","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"40":{"id":40,"position":40,"enabled":true,"name":"operator-extension","value":"operator XML default","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"41":{"id":41,"position":41,"enabled":true,"name":"operator-key","value":"9","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"42":{"id":42,"position":42,"enabled":true,"name":"vmain-extension","value":"vmain XML default","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"43":{"id":43,"position":43,"enabled":true,"name":"vmain-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"5":{"id":5,"position":5,"enabled":true,"name":"min-record-len","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"6":{"id":6,"position":6,"enabled":true,"name":"max-record-len","value":"300","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"7":{"id":7,"position":7,"enabled":true,"name":"max-retries","value":"3","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"8":{"id":8,"position":8,"enabled":true,"name":"tone-spec","value":"%(1000, 0, 640)","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}},"9":{"id":9,"position":9,"enabled":true,"name":"callback-dialplan","value":"XML","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}}
	//Errors:
	case "GetVoicemailProfileParameters":
		resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailProfileParameter{}, onlyAdminGroup())
	//Request:{"event":"AddVoicemailProfileParameter","data":{"token":"example-token","id":2,"param":{"name":"paramn","value":"paramv"}}}
	//Response:{"MessageType":"AddVoicemailProfileParameter","data":{"id":48,"position":44,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "AddVoicemailProfileParameter":
		resp = getUserForConfig(msg, setConfig, &altStruct.ConfigVoicemailProfileParameter{Name: msg.Param.Name, Value: msg.Param.Value, Enabled: true, Parent: &altStruct.ConfigVoicemailProfile{Id: msg.Id}}, onlyAdminGroup())
	//Request:{"event":"DelVoicemailProfileParameter","data":{"token":"example-token","param":{"id":48}}}
	//Response:{"MessageType":"DelVoicemailProfileParameter","data":{"id":48,"position":44,"enabled":true,"name":"paramn","value":"paramv","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "DelVoicemailProfileParameter":
		resp = getUserForConfig(msg, delConfig, &altStruct.ConfigVoicemailProfileParameter{Id: msg.Param.Id}, onlyAdminGroup())
	//Request:{"event":"SwitchVoicemailProfileParameter","data":{"token":"example-token","param":{"id":43,"enabled":false}}}
	//Response:{"MessageType":"SwitchVoicemailProfileParameter","data":{"id":43,"position":43,"enabled":false,"name":"vmain-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "SwitchVoicemailProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailProfileParameter{Id: msg.Param.Id, Enabled: msg.Param.Enabled}, []string{"Enabled"}}, onlyAdminGroup())
	//Request:{"event":"UpdateVoicemailProfileParameter","data":{"token":"example-token","param":{"id":43,"name":"vmain-key","value":"*"}}}
	//Response:{"MessageType":"UpdateVoicemailProfileParameter","data":{"id":43,"position":43,"enabled":true,"name":"vmain-key","value":"*","description":"","parent":{"id":2,"position":0,"enabled":false,"name":"","description":"","parent":null}}}
	//Errors:
	case "UpdateVoicemailProfileParameter":
		resp = getUserForConfig(msg, updateConfig, struct {
			S interface{}
			A []string
		}{&altStruct.ConfigVoicemailProfileParameter{Id: msg.Param.Id, Name: msg.Param.Name, Value: msg.Param.Value}, []string{"Name", "Value"}}, onlyAdminGroup())
		/*
			case "GetVoicemailProfileParameters":
				resp = getUserForConfig(msg, getConfig, &altStruct.ConfigVoicemailProfileEmailParameter{}, onlyAdminGroup())
			case "AddVoicemailProfileParameter":
				resp = getUser(msg, AddVoicemailProfileParameter, onlyAdminGroup())
			case "DelVoicemailProfileParameter":
				resp = getUser(msg, DelVoicemailProfileParameter, onlyAdminGroup())
			case "SwitchVoicemailProfileParameter":
				resp = getUser(msg, SwitchVoicemailProfileParameter, onlyAdminGroup())
			case "UpdateVoicemailProfileParameter":
				resp = getUser(msg, UpdateVoicemailProfileParameter, onlyAdminGroup())
		*/
	//## Global Variables
	//Request:{"event":"GetGlobalVariables","data":{"token":"example-token"}}
	//Response:{"MessageType":"GetGlobalVariables","global_variables":{"1":{"id":1,"enabled":true,"dynamic":true,"name":"hostname","value":"debian-05","type":"set","position":1},"10":{"id":10,"enabled":true,"dynamic":true,"name":"log_dir","value":"/var/log/freeswitch","type":"set","position":10},"100":{"id":100,"enabled":true,"dynamic":false,"name":"video_mute_png","value":"/var/lib/freeswitch/images/default-mute.png","type":"set","position":100},"101":{"id":101,"enabled":true,"dynamic":false,"name":"video_no_avatar_png","value":"/var/lib/freeswitch/images/default-avatar.png","type":"set","position":101},"102":{"id":102,"enabled":true,"dynamic":false,"name":"rtp_liberal_dtmf","value":"true","type":"set","position":102},"103":{"id":103,"enabled":true,"dynamic":false,"name":"AT_EPENT1","value":"0 0 0 -1 -1 0 -1 0 -1 -1 0 -1","type":"set","position":103},"104":{"id":104,"enabled":true,"dynamic":false,"name":"AT_EPENT2","value":"1 1 1 -1 -1 1 -1 1 -1 -1 1 -1","type":"set","position":104},"105":{"id":105,"enabled":true,"dynamic":false,"name":"AT_CPENT1","value":"0 -1 -1 0 -1 0 0 0 -1 -1 0 -1","type":"set","position":105},"106":{"id":106,"enabled":true,"dynamic":false,"name":"AT_CPENT2","value":"1 -1 -1 1 -1 1 1 1 -1 -1 1 -1","type":"set","position":106},"107":{"id":107,"enabled":true,"dynamic":false,"name":"AT_CMAJ1","value":"0 -1 0 0 -1 0 -1 0 0 -1 0 -1","type":"set","position":107},"108":{"id":108,"enabled":true,"dynamic":false,"name":"AT_CMAJ2","value":"1 -1 1 1 -1 1 -1 1 1 -1 1 -1","type":"set","position":109},"109":{"id":109,"enabled":true,"dynamic":false,"name":"AT_BBLUES","value":"1 -1 1 -1 -1 1 -1 1 1 1 -1 -1","type":"set","position":110},"11":{"id":11,"enabled":true,"dynamic":true,"name":"run_dir","value":"/var/run/freeswitch","type":"set","position":11},"110":{"id":110,"enabled":true,"dynamic":false,"name":"ATGPENT2","value":"-1 1 -1 1 -1 1 -1 -1 1 -1 1 -1","type":"set","position":111},"111":{"id":111,"enabled":true,"dynamic":true,"name":"zrtp_enabled","value":"false","type":"set","position":112},"112":{"id":112,"enabled":true,"dynamic":true,"name":"core_uuid","value":"set","type":"set","position":113},"113":{"id":113,"enabled":true,"dynamic":false,"name":"sfsdfsdf","value":"dsfcsdfsfsdfsd","type":"set","position":108},"12":{"id":12,"enabled":true,"dynamic":true,"name":"db_dir","value":"/var/lib/freeswitch/db","type":"set","position":12},"13":{"id":13,"enabled":true,"dynamic":true,"name":"mod_dir","value":"/usr/lib/freeswitch/mod","type":"set","position":13},"14":{"id":14,"enabled":true,"dynamic":true,"name":"htdocs_dir","value":"/usr/share/freeswitch/htdocs","type":"set","position":14},"15":{"id":15,"enabled":true,"dynamic":true,"name":"script_dir","value":"/usr/share/freeswitch/scripts","type":"set","position":15},"16":{"id":16,"enabled":true,"dynamic":true,"name":"temp_dir","value":"/tmp","type":"set","position":16},"17":{"id":17,"enabled":true,"dynamic":true,"name":"grammar_dir","value":"/usr/share/freeswitch/grammar","type":"set","position":17},"18":{"id":18,"enabled":true,"dynamic":true,"name":"certs_dir","value":"/etc/freeswitch/tls","type":"set","position":18},"19":{"id":19,"enabled":true,"dynamic":true,"name":"storage_dir","value":"/var/lib/freeswitch/storage","type":"set","position":19},"2":{"id":2,"enabled":true,"dynamic":true,"name":"local_ip_v4","value":"45.61.54.76","type":"set","position":2},"20":{"id":20,"enabled":true,"dynamic":true,"name":"cache_dir","value":"/var/cache/freeswitch","type":"set","position":20},"21":{"id":21,"enabled":true,"dynamic":true,"name":"switch_serial","value":"2d3d364cd6cc","type":"set","position":21},"22":{"id":22,"enabled":true,"dynamic":false,"name":"fonts_dir","value":"/usr/share/freeswitch/fonts","type":"set","position":22},"23":{"id":23,"enabled":true,"dynamic":false,"name":"images_dir","value":"/var/lib/freeswitch/images","type":"set","position":23},"24":{"id":24,"enabled":true,"dynamic":false,"name":"data_dir","value":"/usr/share/freeswitch","type":"set","position":24},"25":{"id":25,"enabled":true,"dynamic":false,"name":"localstate_dir","value":"/var/lib/freeswitch","type":"set","position":25},"26":{"id":26,"enabled":true,"dynamic":false,"name":"default_password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","type":"set","position":26},"27":{"id":27,"enabled":true,"dynamic":false,"name":"domain","value":"45.61.54.76","type":"set","position":27},"28":{"id":28,"enabled":true,"dynamic":false,"name":"domain_name","value":"45.61.54.76","type":"set","position":28},"29":{"id":29,"enabled":true,"dynamic":false,"name":"hold_music","value":"local_stream://moh","type":"set","position":29},"3":{"id":3,"enabled":true,"dynamic":true,"name":"local_mask_v4","value":"255.255.255.0","type":"set","position":3},"30":{"id":30,"enabled":true,"dynamic":false,"name":"use_profile","value":"external","type":"set","position":30},"31":{"id":31,"enabled":true,"dynamic":false,"name":"rtp_sdes_suites","value":"AEAD_AES_256_GCM_8|AEAD_AES_128_GCM_8|AES_CM_256_HMAC_SHA1_80|AES_CM_192_HMAC_SHA1_80|AES_CM_128_HMAC_SHA1_80|AES_CM_256_HMAC_SHA1_32|AES_CM_192_HMAC_SHA1_32|AES_CM_128_HMAC_SHA1_32|AES_CM_128_NULL_AUTH","type":"set","position":31},"32":{"id":32,"enabled":true,"dynamic":false,"name":"zrtp_secure_media","value":"true","type":"set","position":32},"33":{"id":33,"enabled":true,"dynamic":false,"name":"global_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":33},"34":{"id":34,"enabled":true,"dynamic":false,"name":"outbound_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":34},"35":{"id":35,"enabled":true,"dynamic":false,"name":"xmpp_client_profile","value":"xmppc","type":"set","position":35},"36":{"id":36,"enabled":true,"dynamic":false,"name":"xmpp_server_profile","value":"xmpps","type":"set","position":36},"37":{"id":37,"enabled":true,"dynamic":false,"name":"bind_server_ip","value":"auto","type":"set","position":37},"38":{"id":38,"enabled":true,"dynamic":false,"name":"external_rtp_ip","value":"45.61.54.76","type":"set","position":38},"39":{"id":39,"enabled":true,"dynamic":false,"name":"external_sip_ip","value":"45.61.54.76","type":"set","position":39},"4":{"id":4,"enabled":true,"dynamic":true,"name":"local_ip_v6","value":"::1","type":"set","position":4},"40":{"id":40,"enabled":true,"dynamic":false,"name":"unroll_loops","value":"true","type":"set","position":40},"41":{"id":41,"enabled":true,"dynamic":false,"name":"outbound_caller_name","value":"FreeSWITCH","type":"set","position":41},"42":{"id":42,"enabled":true,"dynamic":false,"name":"outbound_caller_id","value":"0000000000","type":"set","position":42},"43":{"id":43,"enabled":true,"dynamic":false,"name":"call_debug","value":"false","type":"set","position":43},"44":{"id":44,"enabled":true,"dynamic":false,"name":"console_loglevel","value":"info","type":"set","position":44},"45":{"id":45,"enabled":true,"dynamic":false,"name":"default_areacode","value":"918","type":"set","position":45},"46":{"id":46,"enabled":true,"dynamic":false,"name":"default_country","value":"US","type":"set","position":46},"47":{"id":47,"enabled":true,"dynamic":false,"name":"presence_privacy","value":"false","type":"set","position":47},"48":{"id":48,"enabled":true,"dynamic":false,"name":"au-ring","value":"%(400,200,383,417);%(400,2000,383,417)","type":"set","position":48},"49":{"id":49,"enabled":true,"dynamic":false,"name":"be-ring","value":"%(1000,3000,425)","type":"set","position":49},"5":{"id":5,"enabled":true,"dynamic":true,"name":"base_dir","value":"/usr","type":"set","position":5},"50":{"id":50,"enabled":true,"dynamic":false,"name":"ca-ring","value":"%(2000,4000,440,480)","type":"set","position":50},"51":{"id":51,"enabled":true,"dynamic":false,"name":"cn-ring","value":"%(1000,4000,450)","type":"set","position":51},"52":{"id":52,"enabled":true,"dynamic":false,"name":"cy-ring","value":"%(1500,3000,425)","type":"set","position":52},"53":{"id":53,"enabled":true,"dynamic":false,"name":"cz-ring","value":"%(1000,4000,425)","type":"set","position":53},"54":{"id":54,"enabled":true,"dynamic":false,"name":"de-ring","value":"%(1000,4000,425)","type":"set","position":54},"55":{"id":55,"enabled":true,"dynamic":false,"name":"dk-ring","value":"%(1000,4000,425)","type":"set","position":55},"56":{"id":56,"enabled":true,"dynamic":false,"name":"dz-ring","value":"%(1500,3500,425)","type":"set","position":56},"57":{"id":57,"enabled":true,"dynamic":false,"name":"eg-ring","value":"%(2000,1000,475,375)","type":"set","position":57},"58":{"id":58,"enabled":true,"dynamic":false,"name":"es-ring","value":"%(1500,3000,425)","type":"set","position":58},"59":{"id":59,"enabled":true,"dynamic":false,"name":"fi-ring","value":"%(1000,4000,425)","type":"set","position":59},"6":{"id":6,"enabled":true,"dynamic":true,"name":"recordings_dir","value":"/var/lib/freeswitch/recordings","type":"set","position":6},"60":{"id":60,"enabled":true,"dynamic":false,"name":"fr-ring","value":"%(1500,3500,440)","type":"set","position":60},"61":{"id":61,"enabled":true,"dynamic":false,"name":"hk-ring","value":"%(400,200,440,480);%(400,3000,440,480)","type":"set","position":61},"62":{"id":62,"enabled":true,"dynamic":false,"name":"hu-ring","value":"%(1250,3750,425)","type":"set","position":62},"63":{"id":63,"enabled":true,"dynamic":false,"name":"il-ring","value":"%(1000,3000,400)","type":"set","position":63},"64":{"id":64,"enabled":true,"dynamic":false,"name":"in-ring","value":"%(400,200,425,375);%(400,2000,425,375)","type":"set","position":64},"65":{"id":65,"enabled":true,"dynamic":false,"name":"jp-ring","value":"%(1000,2000,420,380)","type":"set","position":65},"66":{"id":66,"enabled":true,"dynamic":false,"name":"ko-ring","value":"%(1000,2000,440,480)","type":"set","position":66},"67":{"id":67,"enabled":true,"dynamic":false,"name":"pk-ring","value":"%(1000,2000,400)","type":"set","position":67},"68":{"id":68,"enabled":true,"dynamic":false,"name":"pl-ring","value":"%(1000,4000,425)","type":"set","position":68},"69":{"id":69,"enabled":true,"dynamic":false,"name":"ro-ring","value":"%(1850,4150,475,425)","type":"set","position":69},"7":{"id":7,"enabled":true,"dynamic":true,"name":"sound_prefix","value":"/usr/share/freeswitch/sounds","type":"set","position":7},"70":{"id":70,"enabled":true,"dynamic":false,"name":"rs-ring","value":"%(1000,4000,425)","type":"set","position":70},"71":{"id":71,"enabled":true,"dynamic":false,"name":"ru-ring","value":"%(800,3200,425)","type":"set","position":71},"72":{"id":72,"enabled":true,"dynamic":false,"name":"sa-ring","value":"%(1200,4600,425)","type":"set","position":72},"73":{"id":73,"enabled":true,"dynamic":false,"name":"tr-ring","value":"%(2000,4000,450)","type":"set","position":73},"74":{"id":74,"enabled":true,"dynamic":false,"name":"uk-ring","value":"%(400,200,400,450);%(400,2000,400,450)","type":"set","position":74},"75":{"id":75,"enabled":true,"dynamic":false,"name":"us-ring","value":"%(2000,4000,440,480)","type":"set","position":75},"76":{"id":76,"enabled":true,"dynamic":false,"name":"bong-ring","value":"v","type":"set","position":76},"77":{"id":77,"enabled":true,"dynamic":false,"name":"beep","value":"%(1000,0,640)","type":"set","position":77},"78":{"id":78,"enabled":true,"dynamic":false,"name":"sit","value":"%(274,0,913.8);%(274,0,1370.6);%(380,0,1776.7)","type":"set","position":78},"79":{"id":79,"enabled":true,"dynamic":false,"name":"df_us_ssn","value":"(?!219099999|078051120)(?!666|000|9\\d{2})\\d{3}(?!00)\\d{2}(?!0{4})\\d{4}","type":"set","position":79},"8":{"id":8,"enabled":true,"dynamic":true,"name":"sounds_dir","value":"/usr/share/freeswitch/sounds","type":"set","position":8},"80":{"id":80,"enabled":true,"dynamic":false,"name":"df_luhn","value":"?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\\d{3})\\d{11}","type":"set","position":80},"81":{"id":81,"enabled":true,"dynamic":false,"name":"default_provider","value":"example.com","type":"set","position":81},"82":{"id":82,"enabled":true,"dynamic":false,"name":"default_provider_username","value":"joeuser","type":"set","position":82},"83":{"id":83,"enabled":true,"dynamic":false,"name":"default_provider_password","value":"password","type":"set","position":83},"84":{"id":84,"enabled":true,"dynamic":false,"name":"default_provider_from_domain","value":"example.com","type":"set","position":84},"85":{"id":85,"enabled":true,"dynamic":false,"name":"default_provider_register","value":"false","type":"set","position":85},"86":{"id":86,"enabled":true,"dynamic":false,"name":"default_provider_contact","value":"5000","type":"set","position":86},"87":{"id":87,"enabled":true,"dynamic":false,"name":"sip_tls_version","value":"tlsv1,tlsv1.1,tlsv1.2","type":"set","position":87},"88":{"id":88,"enabled":true,"dynamic":false,"name":"sip_tls_ciphers","value":"ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH","type":"set","position":88},"89":{"id":89,"enabled":true,"dynamic":false,"name":"internal_auth_calls","value":"true","type":"set","position":89},"9":{"id":9,"enabled":true,"dynamic":true,"name":"conf_dir","value":"/etc/freeswitch","type":"set","position":9},"90":{"id":90,"enabled":true,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":90},"91":{"id":91,"enabled":true,"dynamic":false,"name":"internal_tls_port","value":"5061","type":"set","position":91},"92":{"id":92,"enabled":true,"dynamic":false,"name":"internal_ssl_enable","value":"false","type":"set","position":92},"93":{"id":93,"enabled":true,"dynamic":false,"name":"external_auth_calls","value":"false","type":"set","position":93},"94":{"id":94,"enabled":true,"dynamic":false,"name":"external_sip_port","value":"5080","type":"set","position":94},"95":{"id":95,"enabled":true,"dynamic":false,"name":"external_tls_port","value":"5081","type":"set","position":95},"96":{"id":96,"enabled":true,"dynamic":false,"name":"external_ssl_enable","value":"false","type":"set","position":96},"97":{"id":97,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_in","value":"3mb","type":"set","position":97},"98":{"id":98,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_out","value":"3mb","type":"set","position":98},"99":{"id":99,"enabled":true,"dynamic":false,"name":"suppress_cng","value":"true","type":"set","position":99}}}
	//Errors:
	case "GetGlobalVariables":
		resp = getUser(msg, GetGlobalVariables, onlyAdminGroup())
	//Request:{"event":"UpdateGlobalVariable","data":{"token":"example-token","variable":{"id":117,"dynamic":false,"name":"new_var2","value":"new_val2","type":"set"}}}
	//Response:{"MessageType":"UpdateGlobalVariable","global_variables":{"117":{"id":117,"enabled":true,"dynamic":false,"name":"new_var2","value":"new_val2","type":"set","position":114}}}
	//Errors:
	case "UpdateGlobalVariable":
		resp = getUser(msg, UpdateGlobalVariable, onlyAdminGroup())
	//Request:{"event":"SwitchGlobalVariable","data":{"token":"example-token","variable":{"id":90,"enabled":false}}}
	//Response:{"MessageType":"SwitchGlobalVariable","global_variables":{"90":{"id":90,"enabled":false,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":90}}}
	//Errors:
	case "SwitchGlobalVariable":
		resp = getUser(msg, SwitchGlobalVariable, onlyAdminGroup())
	//Request:{"event":"AddGlobalVariable","data":{"token":"example-token","variable":{"name":"new_var","value":"new_val","type":"set"}}}
	//Response:{"MessageType":"AddGlobalVariable","global_variables":{"117":{"id":117,"enabled":true,"dynamic":false,"name":"new_var","value":"new_val","type":"set","position":114}}}
	//Errors:
	case "AddGlobalVariable":
		resp = getUser(msg, AddGlobalVariable, onlyAdminGroup())
	//Request:{"event":"DelGlobalVariable","data":{"token":"example-token","variable":{"id":117}}}
	//Response:{"MessageType":"DelGlobalVariable","id":117}
	//Errors:
	case "DelGlobalVariable":
		resp = getUser(msg, DelGlobalVariable, onlyAdminGroup())
	//Request:{"event":"MoveGlobalVariable","data":{"token":"example-token","previous_index":111,"current_index":108,"id":110}}
	//Response:{"MessageType":"MoveGlobalVariable","global_variables":{"1":{"id":1,"enabled":true,"dynamic":true,"name":"hostname","value":"debian-05","type":"set","position":1},"10":{"id":10,"enabled":true,"dynamic":true,"name":"log_dir","value":"/var/log/freeswitch","type":"set","position":10},"100":{"id":100,"enabled":true,"dynamic":false,"name":"video_mute_png","value":"/var/lib/freeswitch/images/default-mute.png","type":"set","position":100},"101":{"id":101,"enabled":true,"dynamic":false,"name":"video_no_avatar_png","value":"/var/lib/freeswitch/images/default-avatar.png","type":"set","position":101},"102":{"id":102,"enabled":true,"dynamic":false,"name":"rtp_liberal_dtmf","value":"true","type":"set","position":102},"103":{"id":103,"enabled":true,"dynamic":false,"name":"AT_EPENT1","value":"0 0 0 -1 -1 0 -1 0 -1 -1 0 -1","type":"set","position":103},"104":{"id":104,"enabled":true,"dynamic":false,"name":"AT_EPENT2","value":"1 1 1 -1 -1 1 -1 1 -1 -1 1 -1","type":"set","position":104},"105":{"id":105,"enabled":true,"dynamic":false,"name":"AT_CPENT1","value":"0 -1 -1 0 -1 0 0 0 -1 -1 0 -1","type":"set","position":105},"106":{"id":106,"enabled":true,"dynamic":false,"name":"AT_CPENT2","value":"1 -1 -1 1 -1 1 1 1 -1 -1 1 -1","type":"set","position":106},"107":{"id":107,"enabled":true,"dynamic":false,"name":"AT_CMAJ1","value":"0 -1 0 0 -1 0 -1 0 0 -1 0 -1","type":"set","position":107},"108":{"id":108,"enabled":true,"dynamic":false,"name":"AT_CMAJ2","value":"1 -1 1 1 -1 1 -1 1 1 -1 1 -1","type":"set","position":110},"109":{"id":109,"enabled":true,"dynamic":false,"name":"AT_BBLUES","value":"1 -1 1 -1 -1 1 -1 1 1 1 -1 -1","type":"set","position":111},"11":{"id":11,"enabled":true,"dynamic":true,"name":"run_dir","value":"/var/run/freeswitch","type":"set","position":11},"110":{"id":110,"enabled":true,"dynamic":false,"name":"ATGPENT2","value":"-1 1 -1 1 -1 1 -1 -1 1 -1 1 -1","type":"set","position":108},"111":{"id":111,"enabled":true,"dynamic":true,"name":"zrtp_enabled","value":"false","type":"set","position":112},"112":{"id":112,"enabled":true,"dynamic":true,"name":"core_uuid","value":"set","type":"set","position":113},"113":{"id":113,"enabled":true,"dynamic":false,"name":"sfsdfsdf","value":"dsfcsdfsfsdfsd","type":"set","position":109},"12":{"id":12,"enabled":true,"dynamic":true,"name":"db_dir","value":"/var/lib/freeswitch/db","type":"set","position":12},"13":{"id":13,"enabled":true,"dynamic":true,"name":"mod_dir","value":"/usr/lib/freeswitch/mod","type":"set","position":13},"14":{"id":14,"enabled":true,"dynamic":true,"name":"htdocs_dir","value":"/usr/share/freeswitch/htdocs","type":"set","position":14},"15":{"id":15,"enabled":true,"dynamic":true,"name":"script_dir","value":"/usr/share/freeswitch/scripts","type":"set","position":15},"16":{"id":16,"enabled":true,"dynamic":true,"name":"temp_dir","value":"/tmp","type":"set","position":16},"17":{"id":17,"enabled":true,"dynamic":true,"name":"grammar_dir","value":"/usr/share/freeswitch/grammar","type":"set","position":17},"18":{"id":18,"enabled":true,"dynamic":true,"name":"certs_dir","value":"/etc/freeswitch/tls","type":"set","position":18},"19":{"id":19,"enabled":true,"dynamic":true,"name":"storage_dir","value":"/var/lib/freeswitch/storage","type":"set","position":19},"2":{"id":2,"enabled":true,"dynamic":true,"name":"local_ip_v4","value":"45.61.54.76","type":"set","position":2},"20":{"id":20,"enabled":true,"dynamic":true,"name":"cache_dir","value":"/var/cache/freeswitch","type":"set","position":20},"21":{"id":21,"enabled":true,"dynamic":true,"name":"switch_serial","value":"2d3d364cd6cc","type":"set","position":21},"22":{"id":22,"enabled":true,"dynamic":false,"name":"fonts_dir","value":"/usr/share/freeswitch/fonts","type":"set","position":22},"23":{"id":23,"enabled":true,"dynamic":false,"name":"images_dir","value":"/var/lib/freeswitch/images","type":"set","position":23},"24":{"id":24,"enabled":true,"dynamic":false,"name":"data_dir","value":"/usr/share/freeswitch","type":"set","position":24},"25":{"id":25,"enabled":true,"dynamic":false,"name":"localstate_dir","value":"/var/lib/freeswitch","type":"set","position":25},"26":{"id":26,"enabled":true,"dynamic":false,"name":"default_password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","type":"set","position":26},"27":{"id":27,"enabled":true,"dynamic":false,"name":"domain","value":"45.61.54.76","type":"set","position":27},"28":{"id":28,"enabled":true,"dynamic":false,"name":"domain_name","value":"45.61.54.76","type":"set","position":28},"29":{"id":29,"enabled":true,"dynamic":false,"name":"hold_music","value":"local_stream://moh","type":"set","position":29},"3":{"id":3,"enabled":true,"dynamic":true,"name":"local_mask_v4","value":"255.255.255.0","type":"set","position":3},"30":{"id":30,"enabled":true,"dynamic":false,"name":"use_profile","value":"external","type":"set","position":30},"31":{"id":31,"enabled":true,"dynamic":false,"name":"rtp_sdes_suites","value":"AEAD_AES_256_GCM_8|AEAD_AES_128_GCM_8|AES_CM_256_HMAC_SHA1_80|AES_CM_192_HMAC_SHA1_80|AES_CM_128_HMAC_SHA1_80|AES_CM_256_HMAC_SHA1_32|AES_CM_192_HMAC_SHA1_32|AES_CM_128_HMAC_SHA1_32|AES_CM_128_NULL_AUTH","type":"set","position":31},"32":{"id":32,"enabled":true,"dynamic":false,"name":"zrtp_secure_media","value":"true","type":"set","position":32},"33":{"id":33,"enabled":true,"dynamic":false,"name":"global_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":33},"34":{"id":34,"enabled":true,"dynamic":false,"name":"outbound_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":34},"35":{"id":35,"enabled":true,"dynamic":false,"name":"xmpp_client_profile","value":"xmppc","type":"set","position":35},"36":{"id":36,"enabled":true,"dynamic":false,"name":"xmpp_server_profile","value":"xmpps","type":"set","position":36},"37":{"id":37,"enabled":true,"dynamic":false,"name":"bind_server_ip","value":"auto","type":"set","position":37},"38":{"id":38,"enabled":true,"dynamic":false,"name":"external_rtp_ip","value":"45.61.54.76","type":"set","position":38},"39":{"id":39,"enabled":true,"dynamic":false,"name":"external_sip_ip","value":"45.61.54.76","type":"set","position":39},"4":{"id":4,"enabled":true,"dynamic":true,"name":"local_ip_v6","value":"::1","type":"set","position":4},"40":{"id":40,"enabled":true,"dynamic":false,"name":"unroll_loops","value":"true","type":"set","position":40},"41":{"id":41,"enabled":true,"dynamic":false,"name":"outbound_caller_name","value":"FreeSWITCH","type":"set","position":41},"42":{"id":42,"enabled":true,"dynamic":false,"name":"outbound_caller_id","value":"0000000000","type":"set","position":42},"43":{"id":43,"enabled":true,"dynamic":false,"name":"call_debug","value":"false","type":"set","position":43},"44":{"id":44,"enabled":true,"dynamic":false,"name":"console_loglevel","value":"info","type":"set","position":44},"45":{"id":45,"enabled":true,"dynamic":false,"name":"default_areacode","value":"918","type":"set","position":45},"46":{"id":46,"enabled":true,"dynamic":false,"name":"default_country","value":"US","type":"set","position":46},"47":{"id":47,"enabled":true,"dynamic":false,"name":"presence_privacy","value":"false","type":"set","position":47},"48":{"id":48,"enabled":true,"dynamic":false,"name":"au-ring","value":"%(400,200,383,417);%(400,2000,383,417)","type":"set","position":48},"49":{"id":49,"enabled":true,"dynamic":false,"name":"be-ring","value":"%(1000,3000,425)","type":"set","position":49},"5":{"id":5,"enabled":true,"dynamic":true,"name":"base_dir","value":"/usr","type":"set","position":5},"50":{"id":50,"enabled":true,"dynamic":false,"name":"ca-ring","value":"%(2000,4000,440,480)","type":"set","position":50},"51":{"id":51,"enabled":true,"dynamic":false,"name":"cn-ring","value":"%(1000,4000,450)","type":"set","position":51},"52":{"id":52,"enabled":true,"dynamic":false,"name":"cy-ring","value":"%(1500,3000,425)","type":"set","position":52},"53":{"id":53,"enabled":true,"dynamic":false,"name":"cz-ring","value":"%(1000,4000,425)","type":"set","position":53},"54":{"id":54,"enabled":true,"dynamic":false,"name":"de-ring","value":"%(1000,4000,425)","type":"set","position":54},"55":{"id":55,"enabled":true,"dynamic":false,"name":"dk-ring","value":"%(1000,4000,425)","type":"set","position":55},"56":{"id":56,"enabled":true,"dynamic":false,"name":"dz-ring","value":"%(1500,3500,425)","type":"set","position":56},"57":{"id":57,"enabled":true,"dynamic":false,"name":"eg-ring","value":"%(2000,1000,475,375)","type":"set","position":57},"58":{"id":58,"enabled":true,"dynamic":false,"name":"es-ring","value":"%(1500,3000,425)","type":"set","position":58},"59":{"id":59,"enabled":true,"dynamic":false,"name":"fi-ring","value":"%(1000,4000,425)","type":"set","position":59},"6":{"id":6,"enabled":true,"dynamic":true,"name":"recordings_dir","value":"/var/lib/freeswitch/recordings","type":"set","position":6},"60":{"id":60,"enabled":true,"dynamic":false,"name":"fr-ring","value":"%(1500,3500,440)","type":"set","position":60},"61":{"id":61,"enabled":true,"dynamic":false,"name":"hk-ring","value":"%(400,200,440,480);%(400,3000,440,480)","type":"set","position":61},"62":{"id":62,"enabled":true,"dynamic":false,"name":"hu-ring","value":"%(1250,3750,425)","type":"set","position":62},"63":{"id":63,"enabled":true,"dynamic":false,"name":"il-ring","value":"%(1000,3000,400)","type":"set","position":63},"64":{"id":64,"enabled":true,"dynamic":false,"name":"in-ring","value":"%(400,200,425,375);%(400,2000,425,375)","type":"set","position":64},"65":{"id":65,"enabled":true,"dynamic":false,"name":"jp-ring","value":"%(1000,2000,420,380)","type":"set","position":65},"66":{"id":66,"enabled":true,"dynamic":false,"name":"ko-ring","value":"%(1000,2000,440,480)","type":"set","position":66},"67":{"id":67,"enabled":true,"dynamic":false,"name":"pk-ring","value":"%(1000,2000,400)","type":"set","position":67},"68":{"id":68,"enabled":true,"dynamic":false,"name":"pl-ring","value":"%(1000,4000,425)","type":"set","position":68},"69":{"id":69,"enabled":true,"dynamic":false,"name":"ro-ring","value":"%(1850,4150,475,425)","type":"set","position":69},"7":{"id":7,"enabled":true,"dynamic":true,"name":"sound_prefix","value":"/usr/share/freeswitch/sounds","type":"set","position":7},"70":{"id":70,"enabled":true,"dynamic":false,"name":"rs-ring","value":"%(1000,4000,425)","type":"set","position":70},"71":{"id":71,"enabled":true,"dynamic":false,"name":"ru-ring","value":"%(800,3200,425)","type":"set","position":71},"72":{"id":72,"enabled":true,"dynamic":false,"name":"sa-ring","value":"%(1200,4600,425)","type":"set","position":72},"73":{"id":73,"enabled":true,"dynamic":false,"name":"tr-ring","value":"%(2000,4000,450)","type":"set","position":73},"74":{"id":74,"enabled":true,"dynamic":false,"name":"uk-ring","value":"%(400,200,400,450);%(400,2000,400,450)","type":"set","position":74},"75":{"id":75,"enabled":true,"dynamic":false,"name":"us-ring","value":"%(2000,4000,440,480)","type":"set","position":75},"76":{"id":76,"enabled":true,"dynamic":false,"name":"bong-ring","value":"v","type":"set","position":76},"77":{"id":77,"enabled":true,"dynamic":false,"name":"beep","value":"%(1000,0,640)","type":"set","position":77},"78":{"id":78,"enabled":true,"dynamic":false,"name":"sit","value":"%(274,0,913.8);%(274,0,1370.6);%(380,0,1776.7)","type":"set","position":78},"79":{"id":79,"enabled":true,"dynamic":false,"name":"df_us_ssn","value":"(?!219099999|078051120)(?!666|000|9\\d{2})\\d{3}(?!00)\\d{2}(?!0{4})\\d{4}","type":"set","position":79},"8":{"id":8,"enabled":true,"dynamic":true,"name":"sounds_dir","value":"/usr/share/freeswitch/sounds","type":"set","position":8},"80":{"id":80,"enabled":true,"dynamic":false,"name":"df_luhn","value":"?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\\d{3})\\d{11}","type":"set","position":80},"81":{"id":81,"enabled":true,"dynamic":false,"name":"default_provider","value":"example.com","type":"set","position":81},"82":{"id":82,"enabled":true,"dynamic":false,"name":"default_provider_username","value":"joeuser","type":"set","position":82},"83":{"id":83,"enabled":true,"dynamic":false,"name":"default_provider_password","value":"password","type":"set","position":83},"84":{"id":84,"enabled":true,"dynamic":false,"name":"default_provider_from_domain","value":"example.com","type":"set","position":84},"85":{"id":85,"enabled":true,"dynamic":false,"name":"default_provider_register","value":"false","type":"set","position":85},"86":{"id":86,"enabled":true,"dynamic":false,"name":"default_provider_contact","value":"5000","type":"set","position":86},"87":{"id":87,"enabled":true,"dynamic":false,"name":"sip_tls_version","value":"tlsv1,tlsv1.1,tlsv1.2","type":"set","position":87},"88":{"id":88,"enabled":true,"dynamic":false,"name":"sip_tls_ciphers","value":"ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH","type":"set","position":88},"89":{"id":89,"enabled":true,"dynamic":false,"name":"internal_auth_calls","value":"true","type":"set","position":89},"9":{"id":9,"enabled":true,"dynamic":true,"name":"conf_dir","value":"/etc/freeswitch","type":"set","position":9},"90":{"id":90,"enabled":true,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":90},"91":{"id":91,"enabled":true,"dynamic":false,"name":"internal_tls_port","value":"5061","type":"set","position":91},"92":{"id":92,"enabled":true,"dynamic":false,"name":"internal_ssl_enable","value":"false","type":"set","position":92},"93":{"id":93,"enabled":true,"dynamic":false,"name":"external_auth_calls","value":"false","type":"set","position":93},"94":{"id":94,"enabled":true,"dynamic":false,"name":"external_sip_port","value":"5080","type":"set","position":94},"95":{"id":95,"enabled":true,"dynamic":false,"name":"external_tls_port","value":"5081","type":"set","position":95},"96":{"id":96,"enabled":true,"dynamic":false,"name":"external_ssl_enable","value":"false","type":"set","position":96},"97":{"id":97,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_in","value":"3mb","type":"set","position":97},"98":{"id":98,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_out","value":"3mb","type":"set","position":98},"99":{"id":99,"enabled":true,"dynamic":false,"name":"suppress_cng","value":"true","type":"set","position":99}}}
	//Errors:
	case "MoveGlobalVariable":
		resp = getUser(msg, MoveGlobalVariable, onlyAdminGroup())
	//Request:{"event":"ImportGlobalVariables","data":{"token":"example-token"}}
	//Response:{"MessageType":"ImportGlobalVariables","global_variables":{"1":{"id":1,"enabled":true,"dynamic":true,"name":"hostname","value":"debian-05","type":"set","position":1},"10":{"id":10,"enabled":true,"dynamic":true,"name":"log_dir","value":"/var/log/freeswitch","type":"set","position":10},"100":{"id":100,"enabled":true,"dynamic":false,"name":"video_no_avatar_png","value":"/var/lib/freeswitch/images/default-avatar.png","type":"set","position":100},"101":{"id":101,"enabled":true,"dynamic":false,"name":"rtp_liberal_dtmf","value":"true","type":"set","position":101},"102":{"id":102,"enabled":true,"dynamic":false,"name":"sfsdfsdf","value":"dsfcsdfsfsdfsd","type":"set","position":102},"103":{"id":103,"enabled":true,"dynamic":false,"name":"AT_EPENT1","value":"0 0 0 -1 -1 0 -1 0 -1 -1 0 -1","type":"set","position":103},"104":{"id":104,"enabled":true,"dynamic":false,"name":"AT_EPENT2","value":"1 1 1 -1 -1 1 -1 1 -1 -1 1 -1","type":"set","position":104},"105":{"id":105,"enabled":true,"dynamic":false,"name":"AT_CPENT1","value":"0 -1 -1 0 -1 0 0 0 -1 -1 0 -1","type":"set","position":105},"106":{"id":106,"enabled":true,"dynamic":false,"name":"AT_CPENT2","value":"1 -1 -1 1 -1 1 1 1 -1 -1 1 -1","type":"set","position":106},"107":{"id":107,"enabled":true,"dynamic":false,"name":"AT_CMAJ1","value":"0 -1 0 0 -1 0 -1 0 0 -1 0 -1","type":"set","position":107},"108":{"id":108,"enabled":true,"dynamic":false,"name":"AT_CMAJ2","value":"1 -1 1 1 -1 1 -1 1 1 -1 1 -1","type":"set","position":108},"109":{"id":109,"enabled":true,"dynamic":false,"name":"AT_BBLUES","value":"1 -1 1 -1 -1 1 -1 1 1 1 -1 -1","type":"set","position":109},"11":{"id":11,"enabled":true,"dynamic":true,"name":"run_dir","value":"/var/run/freeswitch","type":"set","position":11},"110":{"id":110,"enabled":true,"dynamic":false,"name":"ATGPENT2","value":"-1 1 -1 1 -1 1 -1 -1 1 -1 1 -1","type":"set","position":110},"111":{"id":111,"enabled":true,"dynamic":true,"name":"core_uuid","value":"4ee847e9-b9fb-49a8-99be-11e42a8cfdd4","type":"set","position":111},"112":{"id":112,"enabled":true,"dynamic":true,"name":"zrtp_enabled","value":"false","type":"set","position":112},"113":{"id":113,"enabled":true,"dynamic":false,"name":"internal_sip_port","value":"5060","type":"set","position":113},"12":{"id":12,"enabled":true,"dynamic":true,"name":"db_dir","value":"/var/lib/freeswitch/db","type":"set","position":12},"13":{"id":13,"enabled":true,"dynamic":true,"name":"mod_dir","value":"/usr/lib/freeswitch/mod","type":"set","position":13},"14":{"id":14,"enabled":true,"dynamic":true,"name":"htdocs_dir","value":"/usr/share/freeswitch/htdocs","type":"set","position":14},"15":{"id":15,"enabled":true,"dynamic":true,"name":"script_dir","value":"/usr/share/freeswitch/scripts","type":"set","position":15},"16":{"id":16,"enabled":true,"dynamic":true,"name":"temp_dir","value":"/tmp","type":"set","position":16},"17":{"id":17,"enabled":true,"dynamic":true,"name":"grammar_dir","value":"/usr/share/freeswitch/grammar","type":"set","position":17},"18":{"id":18,"enabled":true,"dynamic":true,"name":"certs_dir","value":"/etc/freeswitch/tls","type":"set","position":18},"19":{"id":19,"enabled":true,"dynamic":true,"name":"storage_dir","value":"/var/lib/freeswitch/storage","type":"set","position":19},"2":{"id":2,"enabled":true,"dynamic":true,"name":"local_ip_v4","value":"45.61.54.76","type":"set","position":2},"20":{"id":20,"enabled":true,"dynamic":true,"name":"cache_dir","value":"/var/cache/freeswitch","type":"set","position":20},"21":{"id":21,"enabled":true,"dynamic":true,"name":"switch_serial","value":"2d3d364cd6cc","type":"set","position":21},"22":{"id":22,"enabled":true,"dynamic":false,"name":"fonts_dir","value":"/usr/share/freeswitch/fonts","type":"set","position":22},"23":{"id":23,"enabled":true,"dynamic":false,"name":"images_dir","value":"/var/lib/freeswitch/images","type":"set","position":23},"24":{"id":24,"enabled":true,"dynamic":false,"name":"data_dir","value":"/usr/share/freeswitch","type":"set","position":24},"25":{"id":25,"enabled":true,"dynamic":false,"name":"localstate_dir","value":"/var/lib/freeswitch","type":"set","position":25},"26":{"id":26,"enabled":true,"dynamic":false,"name":"default_password","value":"12345asdqwe123asd213fsfd3qrsd3qrrfd32rffd5uhr6","type":"set","position":26},"27":{"id":27,"enabled":true,"dynamic":false,"name":"domain","value":"45.61.54.76","type":"set","position":27},"28":{"id":28,"enabled":true,"dynamic":false,"name":"domain_name","value":"45.61.54.76","type":"set","position":28},"29":{"id":29,"enabled":true,"dynamic":false,"name":"hold_music","value":"local_stream://moh","type":"set","position":29},"3":{"id":3,"enabled":true,"dynamic":true,"name":"local_mask_v4","value":"255.255.255.0","type":"set","position":3},"30":{"id":30,"enabled":true,"dynamic":false,"name":"use_profile","value":"external","type":"set","position":30},"31":{"id":31,"enabled":true,"dynamic":false,"name":"rtp_sdes_suites","value":"AEAD_AES_256_GCM_8|AEAD_AES_128_GCM_8|AES_CM_256_HMAC_SHA1_80|AES_CM_192_HMAC_SHA1_80|AES_CM_128_HMAC_SHA1_80|AES_CM_256_HMAC_SHA1_32|AES_CM_192_HMAC_SHA1_32|AES_CM_128_HMAC_SHA1_32|AES_CM_128_NULL_AUTH","type":"set","position":31},"32":{"id":32,"enabled":true,"dynamic":false,"name":"zrtp_secure_media","value":"true","type":"set","position":32},"33":{"id":33,"enabled":true,"dynamic":false,"name":"global_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":33},"34":{"id":34,"enabled":true,"dynamic":false,"name":"outbound_codec_prefs","value":"OPUS,G722,PCMU,PCMA,H264,VP8","type":"set","position":34},"35":{"id":35,"enabled":true,"dynamic":false,"name":"xmpp_client_profile","value":"xmppc","type":"set","position":35},"36":{"id":36,"enabled":true,"dynamic":false,"name":"xmpp_server_profile","value":"xmpps","type":"set","position":36},"37":{"id":37,"enabled":true,"dynamic":false,"name":"bind_server_ip","value":"auto","type":"set","position":37},"38":{"id":38,"enabled":true,"dynamic":false,"name":"external_rtp_ip","value":"45.61.54.76","type":"set","position":38},"39":{"id":39,"enabled":true,"dynamic":false,"name":"external_sip_ip","value":"45.61.54.76","type":"set","position":39},"4":{"id":4,"enabled":true,"dynamic":true,"name":"local_ip_v6","value":"::1","type":"set","position":4},"40":{"id":40,"enabled":true,"dynamic":false,"name":"unroll_loops","value":"true","type":"set","position":40},"41":{"id":41,"enabled":true,"dynamic":false,"name":"outbound_caller_name","value":"FreeSWITCH","type":"set","position":41},"42":{"id":42,"enabled":true,"dynamic":false,"name":"outbound_caller_id","value":"0000000000","type":"set","position":42},"43":{"id":43,"enabled":true,"dynamic":false,"name":"call_debug","value":"false","type":"set","position":43},"44":{"id":44,"enabled":true,"dynamic":false,"name":"console_loglevel","value":"info","type":"set","position":44},"45":{"id":45,"enabled":true,"dynamic":false,"name":"default_areacode","value":"918","type":"set","position":45},"46":{"id":46,"enabled":true,"dynamic":false,"name":"default_country","value":"US","type":"set","position":46},"47":{"id":47,"enabled":true,"dynamic":false,"name":"presence_privacy","value":"false","type":"set","position":47},"48":{"id":48,"enabled":true,"dynamic":false,"name":"au-ring","value":"%(400,200,383,417);%(400,2000,383,417)","type":"set","position":48},"49":{"id":49,"enabled":true,"dynamic":false,"name":"be-ring","value":"%(1000,3000,425)","type":"set","position":49},"5":{"id":5,"enabled":true,"dynamic":true,"name":"base_dir","value":"/usr","type":"set","position":5},"50":{"id":50,"enabled":true,"dynamic":false,"name":"ca-ring","value":"%(2000,4000,440,480)","type":"set","position":50},"51":{"id":51,"enabled":true,"dynamic":false,"name":"cn-ring","value":"%(1000,4000,450)","type":"set","position":51},"52":{"id":52,"enabled":true,"dynamic":false,"name":"cy-ring","value":"%(1500,3000,425)","type":"set","position":52},"53":{"id":53,"enabled":true,"dynamic":false,"name":"cz-ring","value":"%(1000,4000,425)","type":"set","position":53},"54":{"id":54,"enabled":true,"dynamic":false,"name":"de-ring","value":"%(1000,4000,425)","type":"set","position":54},"55":{"id":55,"enabled":true,"dynamic":false,"name":"dk-ring","value":"%(1000,4000,425)","type":"set","position":55},"56":{"id":56,"enabled":true,"dynamic":false,"name":"dz-ring","value":"%(1500,3500,425)","type":"set","position":56},"57":{"id":57,"enabled":true,"dynamic":false,"name":"eg-ring","value":"%(2000,1000,475,375)","type":"set","position":57},"58":{"id":58,"enabled":true,"dynamic":false,"name":"es-ring","value":"%(1500,3000,425)","type":"set","position":58},"59":{"id":59,"enabled":true,"dynamic":false,"name":"fi-ring","value":"%(1000,4000,425)","type":"set","position":59},"6":{"id":6,"enabled":true,"dynamic":true,"name":"recordings_dir","value":"/var/lib/freeswitch/recordings","type":"set","position":6},"60":{"id":60,"enabled":true,"dynamic":false,"name":"fr-ring","value":"%(1500,3500,440)","type":"set","position":60},"61":{"id":61,"enabled":true,"dynamic":false,"name":"hk-ring","value":"%(400,200,440,480);%(400,3000,440,480)","type":"set","position":61},"62":{"id":62,"enabled":true,"dynamic":false,"name":"hu-ring","value":"%(1250,3750,425)","type":"set","position":62},"63":{"id":63,"enabled":true,"dynamic":false,"name":"il-ring","value":"%(1000,3000,400)","type":"set","position":63},"64":{"id":64,"enabled":true,"dynamic":false,"name":"in-ring","value":"%(400,200,425,375);%(400,2000,425,375)","type":"set","position":64},"65":{"id":65,"enabled":true,"dynamic":false,"name":"jp-ring","value":"%(1000,2000,420,380)","type":"set","position":65},"66":{"id":66,"enabled":true,"dynamic":false,"name":"ko-ring","value":"%(1000,2000,440,480)","type":"set","position":66},"67":{"id":67,"enabled":true,"dynamic":false,"name":"pk-ring","value":"%(1000,2000,400)","type":"set","position":67},"68":{"id":68,"enabled":true,"dynamic":false,"name":"pl-ring","value":"%(1000,4000,425)","type":"set","position":68},"69":{"id":69,"enabled":true,"dynamic":false,"name":"ro-ring","value":"%(1850,4150,475,425)","type":"set","position":69},"7":{"id":7,"enabled":true,"dynamic":true,"name":"sound_prefix","value":"/usr/share/freeswitch/sounds","type":"set","position":7},"70":{"id":70,"enabled":true,"dynamic":false,"name":"rs-ring","value":"%(1000,4000,425)","type":"set","position":70},"71":{"id":71,"enabled":true,"dynamic":false,"name":"ru-ring","value":"%(800,3200,425)","type":"set","position":71},"72":{"id":72,"enabled":true,"dynamic":false,"name":"sa-ring","value":"%(1200,4600,425)","type":"set","position":72},"73":{"id":73,"enabled":true,"dynamic":false,"name":"tr-ring","value":"%(2000,4000,450)","type":"set","position":73},"74":{"id":74,"enabled":true,"dynamic":false,"name":"uk-ring","value":"%(400,200,400,450);%(400,2000,400,450)","type":"set","position":74},"75":{"id":75,"enabled":true,"dynamic":false,"name":"us-ring","value":"%(2000,4000,440,480)","type":"set","position":75},"76":{"id":76,"enabled":true,"dynamic":false,"name":"bong-ring","value":"v","type":"set","position":76},"77":{"id":77,"enabled":true,"dynamic":false,"name":"beep","value":"%(1000,0,640)","type":"set","position":77},"78":{"id":78,"enabled":true,"dynamic":false,"name":"sit","value":"%(274,0,913.8);%(274,0,1370.6);%(380,0,1776.7)","type":"set","position":78},"79":{"id":79,"enabled":true,"dynamic":false,"name":"df_us_ssn","value":"(?!219099999|078051120)(?!666|000|9\\d{2})\\d{3}(?!00)\\d{2}(?!0{4})\\d{4}","type":"set","position":79},"8":{"id":8,"enabled":true,"dynamic":true,"name":"sounds_dir","value":"/usr/share/freeswitch/sounds","type":"set","position":8},"80":{"id":80,"enabled":true,"dynamic":false,"name":"df_luhn","value":"?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\\d{3})\\d{11}","type":"set","position":80},"81":{"id":81,"enabled":true,"dynamic":false,"name":"default_provider","value":"example.com","type":"set","position":81},"82":{"id":82,"enabled":true,"dynamic":false,"name":"default_provider_username","value":"joeuser","type":"set","position":82},"83":{"id":83,"enabled":true,"dynamic":false,"name":"default_provider_password","value":"password","type":"set","position":83},"84":{"id":84,"enabled":true,"dynamic":false,"name":"default_provider_from_domain","value":"example.com","type":"set","position":84},"85":{"id":85,"enabled":true,"dynamic":false,"name":"default_provider_register","value":"false","type":"set","position":85},"86":{"id":86,"enabled":true,"dynamic":false,"name":"default_provider_contact","value":"5000","type":"set","position":86},"87":{"id":87,"enabled":true,"dynamic":false,"name":"sip_tls_version","value":"tlsv1,tlsv1.1,tlsv1.2","type":"set","position":87},"88":{"id":88,"enabled":true,"dynamic":false,"name":"sip_tls_ciphers","value":"ALL:!ADH:!LOW:!EXP:!MD5:@STRENGTH","type":"set","position":88},"89":{"id":89,"enabled":true,"dynamic":false,"name":"internal_auth_calls","value":"true","type":"set","position":89},"9":{"id":9,"enabled":true,"dynamic":true,"name":"conf_dir","value":"/etc/freeswitch","type":"set","position":9},"90":{"id":90,"enabled":true,"dynamic":false,"name":"internal_tls_port","value":"5061","type":"set","position":90},"91":{"id":91,"enabled":true,"dynamic":false,"name":"internal_ssl_enable","value":"false","type":"set","position":91},"92":{"id":92,"enabled":true,"dynamic":false,"name":"external_auth_calls","value":"false","type":"set","position":92},"93":{"id":93,"enabled":true,"dynamic":false,"name":"external_sip_port","value":"5080","type":"set","position":93},"94":{"id":94,"enabled":true,"dynamic":false,"name":"external_tls_port","value":"5081","type":"set","position":94},"95":{"id":95,"enabled":true,"dynamic":false,"name":"external_ssl_enable","value":"false","type":"set","position":95},"96":{"id":96,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_in","value":"3mb","type":"set","position":96},"97":{"id":97,"enabled":true,"dynamic":false,"name":"rtp_video_max_bandwidth_out","value":"3mb","type":"set","position":97},"98":{"id":98,"enabled":true,"dynamic":false,"name":"suppress_cng","value":"true","type":"set","position":98},"99":{"id":99,"enabled":true,"dynamic":false,"name":"video_mute_png","value":"/var/lib/freeswitch/images/default-mute.png","type":"set","position":99}}}
	//Errors:
	case "ImportGlobalVariables":
		resp = getUser(msg, ImportGlobalVariables, onlyAdminGroup())

	//Request:
	//Response:
	//Errors:
	default:
		resp = webStruct.UserResponse{Error: "Wrong event", MessageType: "none"}
	}

	return resp
}

func checkRelogin(data *webStruct.MessageData) webStruct.UserResponse {
	return webStruct.UserResponse{User: data.Context.User, Token: data.Token, MessageType: "relogin"}
}

func checkSettings(data *webStruct.MessageData) webStruct.UserResponse {
	return webStruct.UserResponse{Settings: &cfg.CustomPbx, MessageType: "settings"}
}

func setSettings(data *webStruct.MessageData) webStruct.UserResponse {
	log.Printf("settings update requested payload=%s", logsafe.Redact(data.Payload))
	if data.Payload.Fs.Esl.Pass == "" || data.Payload.Fs.Esl.Port == 0 || data.Payload.Fs.Esl.Host == "" ||
		data.Payload.Db.Host == "" || data.Payload.Db.Port == 0 || data.Payload.Db.Name == "" ||
		data.Payload.Db.User == "" || data.Payload.Db.Pass == "" ||
		data.Payload.Web.Host == "" || data.Payload.Web.Port == 0 ||
		data.Payload.Web.Route == "" ||
		data.Payload.XMLCurl.Host == "" || data.Payload.XMLCurl.Port == 0 ||
		data.Payload.XMLCurl.Route == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: "settings"}
	}
	cfg.CustomPbx.Fs.Esl = data.Payload.Fs.Esl
	cfg.CustomPbx.Db = data.Payload.Db
	cfg.CustomPbx.Web = data.Payload.Web
	cfg.CustomPbx.XMLCurl = data.Payload.XMLCurl
	conf, err := cfg.WD(cfg.CustomPbx)
	if err != nil {
		cfg.RD()
		return webStruct.UserResponse{Error: "can't save", MessageType: "settings"}
	}

	return webStruct.UserResponse{Settings: &conf, MessageType: "settings"}
}

func getCDR(data *webStruct.MessageData) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 250 {
		limit = 25
	}
	offset := data.DBRequest.Offset * limit
	cdr, err := cdrDb.GetList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	/*
		if cdr == nil {
			return webStruct.UserResponse{Error: "nothing", MessageType: data.Event}
		}*/

	return webStruct.UserResponse{CDR: &cdr, MessageType: data.Event}
}

func getPhoneCreds(data *webStruct.MessageData) webStruct.UserResponse {
	if !data.Context.User.SipId.Valid {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}

	userI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: data.Context.User.SipId.Int64})
	if err != nil || userI == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	directoryUser, ok := userI.(altStruct.DirectoryDomainUser)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	param, err := intermediateDB.GetByValue(
		&altStruct.DirectoryDomainUserParameter{Name: "password", Parent: &directoryUser},
		map[string]bool{"Parent": true, "Name": true},
	)
	if err != nil || len(param) == 0 {
		return webStruct.UserResponse{Error: "user password not found", MessageType: data.Event}
	}
	directoryUserParam, ok := param[0].(altStruct.DirectoryDomainUserParameter)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	domainI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomain{Id: directoryUser.Parent.Id})
	if err != nil || domainI == nil {
		return webStruct.UserResponse{Error: "user domain not found", MessageType: data.Event}
	}
	domain, ok := domainI.(altStruct.DirectoryDomain)
	if !ok {
		return webStruct.UserResponse{Error: "user domain not found", MessageType: data.Event}
	}

	password := directoryUserParam.Value

	if password == "" {
		paramI, err := intermediateDB.GetByValue(
			&altStruct.DirectoryDomainParameter{Name: "password", Parent: &domain},
			map[string]bool{"Parent": true, "Name": true},
		)
		if err != nil || len(paramI) == 0 {
			return webStruct.UserResponse{Error: "domain directory password not found", MessageType: data.Event}
		}
		domainParam, ok := paramI[0].(altStruct.DirectoryDomainParameter)
		if !ok {
			return webStruct.UserResponse{Error: "domain directory not found", MessageType: data.Event}
		}

		password = domainParam.Value
	}
	if password == "" || (data.Context.User.Ws == "" && data.Context.User.VertoWs == "") /*|| user.Stun == ""*/ {
		return webStruct.UserResponse{Error: "no enough params params", MessageType: data.Event}
	}

	creds := webStruct.PhoneCreds{}
	creds.UserName = directoryUser.Name
	creds.Password = password
	creds.Domain = domain.Name
	creds.WebRTCLib = data.Context.User.WebRTCLib
	creds.Ws = data.Context.User.Ws
	creds.VertoWs = data.Context.User.VertoWs
	creds.Stun = data.Context.User.Stun
	if creds.Stun == "" && daemonCache.State.StunServerStatus {
		creds.Stun = "stun:" + cfg.CustomPbx.Web.Host + ":" + strconv.Itoa(cfg.CustomPbx.Web.StunPort)
	}

	return webStruct.UserResponse{PhoneCreds: &creds, MessageType: data.Event}
}

func runCLICommand(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func RealFSCLIConnect(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func RealFSCLICommand(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func GetLogs(data *webStruct.MessageData) webStruct.UserResponse {
	limit, offset := normalizePagination(data.DBRequest.Limit, data.DBRequest.Offset)
	logs, err := db.GetList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order, cache.GetCurrentInstanceId())
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Logs: &logs, MessageType: data.Event}
}

func getHEP(data *webStruct.MessageData) webStruct.UserResponse {
	limit, offset := normalizePagination(data.DBRequest.Limit, data.DBRequest.Offset)
	heps, err := db.GetHEPList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{HEPs: &heps, MessageType: data.Event}
}

func GetHEPDetails(data *webStruct.MessageData) webStruct.UserResponse {
	if len(data.ArrVal) == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	heps, err := db.GetHEPDetailsList(data.ArrVal, cache.GetCurrentInstanceId())
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{HEPsDetails: &heps, MessageType: data.Event}
}

func GetInstances(data *webStruct.MessageData) webStruct.UserResponse {
	cache.UpdateCacheInstances()
	var res = cache.GetFSInstances().GetList()
	var currentId = cache.GetCurrentInstanceId()
	return webStruct.UserResponse{FSInstances: &res, MessageType: "GetInstances", Id: &currentId}
}

func UpdateInstanceDescription(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	instance := cache.GetFSInstances().GetById(data.Id)
	if instance == nil {
		return webStruct.UserResponse{Error: "instance not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateFSInstanceDescription(instance, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.FsInstance{instance.Id: instance}

	return webStruct.UserResponse{MessageType: data.Event, FSInstances: &item}

}

func getCallcenterAgents(msg *webStruct.MessageData) webStruct.UserResponse {
	if msg.Param.Name == "" || msg.Param.Value == "" {
		return webStruct.UserResponse{Error: "wrong params", MessageType: msg.Event}
	}
	item := &altStruct.Agent{Id: msg.Param.Id}
	fieldName := mainStruct.GetItemNameByTag(item, msg.Param.Name)
	if fieldName == "id" {
		return webStruct.UserResponse{Error: "please dont", MessageType: msg.Event}
	}
	f := reflect.ValueOf(item).Elem().FieldByName(fieldName)
	switch f.Type().Name() {
	case "string":
		f.SetString(msg.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(msg.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(msg.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}
	// and update with updateConfig func
	return getUserForConfig(msg, updateConfig, struct {
		S interface{}
		A []string
	}{item, []string{fieldName}}, onlyAdminGroup())
}

func getCallcenterTiers(msg *webStruct.MessageData) webStruct.UserResponse {
	if msg.Param.Name == "" || msg.Param.Value == "" {
		return webStruct.UserResponse{Error: "wrong params", MessageType: msg.Event}
	}
	item := &altStruct.Tier{Id: msg.Param.Id}
	fieldName := mainStruct.GetItemNameByTag(item, msg.Param.Name)
	if fieldName == "id" {
		return webStruct.UserResponse{Error: "please dont", MessageType: msg.Event}
	}
	f := reflect.ValueOf(item).Elem().FieldByName(fieldName)
	switch f.Type().Name() {
	case "string":
		f.SetString(msg.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(msg.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(msg.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}

	return getUserForConfig(msg, updateConfig, struct {
		S interface{}
		A []string
	}{item, []string{fieldName}}, onlyAdminGroup())
}

// like getConfig
func getByStruct(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	if data.Id == 0 {
		return errorResponse(data.Event, "no parent id")
	}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true, UseValue: true, Value: data.Id}}

	var res interface{}
	var err error
	if hasPagedRequest(data.DBRequest) {
		filterStr := buildFilteredConfigRequest(filter, data.DBRequest)
		res, err = intermediateDB.GetByFilteredValues(
			item,
			filterStr,
		)
		if err != nil {
			return errorResponse(data.Event, err.Error())
		}
		//TODO: with total all the time
		filterStr.Count = true
		resCount, err := intermediateDB.GetByFilteredValues(
			item,
			filterStr,
		)
		if err != nil {
			return errorResponse(data.Event, err.Error())
		}
		if len(resCount) == 0 {
			return errorResponse(data.Event, "can't count total")
		}
		total, ok := resCount[0].(int64)
		if !ok {
			return errorResponse(data.Event, "can't get total")
		}
		res = paginatedResult{Items: res, Total: total}
	} else {
		res, err = intermediateDB.GetByValuesAsMap(
			item,
			filter,
		)
	}

	if err != nil {
		return errorResponse(data.Event, err.Error())
	}

	return dataResponse(data.Event, res)
}

func setProfileStatuses(resp webStruct.UserResponse) webStruct.UserResponse {
	profiles, ok := resp.Data.(map[int64]interface{})
	if ok {
		profilesX := fsesl.GetSofiaProfilesStatuses()
		for _, profileI := range profiles {
			profile, ok := profileI.(altStruct.ConfigSofiaProfile)
			if !ok {
				continue
			}
			profileX := profilesX[profile.Id]
			if profileX == nil {
				continue
			}
			profile.Started = profileX.Started
			profile.State = profileX.State
			profile.Uri = profileX.Uri
			profiles[profile.Id] = profile
		}
		resp.Data = profiles
	}
	return resp
}

func setGatewayStatuses(resp webStruct.UserResponse) webStruct.UserResponse {
	gateways, ok := resp.Data.(map[int64]interface{})
	if ok {
		gatewaysX := fsesl.GetSofiaGatewaysStatuses()
		for _, gatewayI := range gateways {
			gateway, ok := gatewayI.(altStruct.ConfigSofiaProfileGateway)
			if !ok {
				continue
			}
			gatewayX := gatewaysX[gateway.Id]
			if gatewayX == nil {
				continue
			}
			gateway.Started = gatewayX.Started
			gateway.State = gatewayX.State
			gateways[gateway.Id] = gateway
		}
		resp.Data = gateways
	}
	return resp
}
