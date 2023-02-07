package webStruct

import (
	"custompbx/altStruct"
	"custompbx/cfg"
	"custompbx/mainStruct"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"strings"
	"sync"
	"time"
)

const (
	GetDashboard              = "[Dashboard]"
	GetSofiaProfiles          = "[Config] Get_sofia_profiles"
	GetSofiaProfileGateways   = "[Config] Get_sofia_profile_gateways"
	GetModules                = "[Config][Get] Modules"
	DialplanDebug             = "[Dialplan] Debug"
	GetDirectoryUser          = "GetDirectoryUsers"
	Unsubscribe               = "[UnSubscribe]"
	SubscribeCallcenterAgents = "SubscribeCallcenterAgents"
	SubscribeCallcenterTiers  = "SubscribeCallcenterTiers"
	SubscribeHepPackages      = "SubscribeHepPackages"
	BroadcastConnection       = "connection"
	AddPostLoadModule         = "AddPostLoadModule"
	SwitchPostLoadModule      = "SwitchPostLoadModule"
)

type Message struct {
	Username    string       `json:"username"`
	MessageType string       `json:"MessageType"`
	Event       string       `json:"event"`
	Data        *MessageData `json:"data"`
}

type MessageData struct {
	Login            string                     `json:"login"`
	Payload          cfg.GeneralCfg             `json:"payload,omitempty"`
	Password         string                     `json:"password"`
	Token            string                     `json:"token,omitempty"`
	Id               int64                      `json:"id,omitempty"`
	IdInt            int64                      `json:"id_int,omitempty"`
	Name             string                     `json:"name,omitempty"`
	Value            string                     `json:"value,omitempty"`
	Direction        string                     `json:"direction,omitempty"`
	Default          string                     `json:"default,omitempty"`
	Node             mainStruct.Node            `json:"node,omitempty"`
	Param            Param                      `json:"param,omitempty"`
	SofiaDomain      mainStruct.SofiaDomain     `json:"sofia_domain,omitempty"`
	SofiaAlias       mainStruct.Alias           `json:"sofia_alias,omitempty"`
	Variable         MessageDataVariable        `json:"variable,omitempty"`
	Regex            mainStruct.Regex           `json:"regex,omitempty"`
	Action           mainStruct.Action          `json:"action,omitempty"`
	AntiAction       mainStruct.AntiAction      `json:"antiaction,omitempty"`
	Condition        mainStruct.Condition       `json:"condition,omitempty"`
	Enabled          *bool                      `json:"enabled,omitempty"`
	KeepSubscription *bool                      `json:"keep_subscription,omitempty"`
	PreviousIndex    int64                      `json:"previous_index,omitempty"`
	CurrentIndex     int64                      `json:"current_index,omitempty"`
	ParamId          int64                      `json:"param_id,omitempty"`
	Bulk             int                        `json:"bulk,omitempty"`
	GroupId          int                        `json:"group_id,omitempty"`
	Field            mainStruct.Field           `json:"field,omitempty"`
	DBRequest        mainStruct.DBRequest       `json:"db_request,omitempty"`
	File             string                     `json:"file,omitempty"`
	Uuid             string                     `json:"uuid,omitempty"`
	ArrVal           []string                   `json:"values,omitempty"`
	Table            mainStruct.Table           `json:"table,omitempty"`
	OdbcCdrField     mainStruct.ODBCField       `json:"odbc_cdr_field,omitempty"`
	WebSettings      map[string]string          `json:"web_settings,omitempty"`
	DistributorNode  mainStruct.DistributorNode `json:"distributor_node,omitempty"`
	Importance       string                     `json:"importance,omitempty"`
	FifoFifoMember   mainStruct.FifoFifoMember  `json:"fifo_fifo_member,omitempty"`
	Fields           []string                   `json:"fields,omitempty"`
	Ids              []int64                    `json:"ids,omitempty"`

	Data     json.RawMessage `json:"data,omitempty"`
	IntSlice []int64         `json:"-"`
	//WebDirectoryUsersTemplate          mainStruct.WebDirectoryUsersTemplate          `json:"web_directory_users_template,omitempty"`
	//WebDirectoryUsersTemplateParameter mainStruct.WebDirectoryUsersTemplateParameter `json:"web_directory_users_template_parameter,omitempty"`
	//WebDirectoryUsersTemplateVariable  mainStruct.WebDirectoryUsersTemplateVariable  `json:"web_directory_users_template_variable,omitempty"`

	Event   string
	Context *WsContext
}

type MessageDataVariable struct {
	Id        int64  `json:"id"`
	Enabled   bool   `json:"enabled"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Direction string `json:"direction"`
	Type      string `json:"type"`
}

type MessageDataParameter struct {
}

type Subscriptions struct {
	mx     sync.RWMutex
	byName map[string]bool
}

func (s *Subscriptions) Get(name string) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.byName[name]
}

func (s *Subscriptions) Set(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[name] = true
}

func (s *Subscriptions) Del(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if !s.byName[name] {
		return
	}
	s.byName[name] = false
}

func (s *Subscriptions) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName = map[string]bool{}
}

type Param struct {
	Id      int64  `json:"id"`
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Secure  string `json:"secure"`
}

type UserResponse struct {
	MessageType             string                                                 `json:"MessageType"`
	Token                   string                                                 `json:"token,omitempty"`
	TokensList              *[]mainStruct.WebUserToken                             `json:"tokens_list,omitempty"`
	Uuid                    *string                                                `json:"uuid,omitempty"`
	Response                *string                                                `json:"response,omitempty"`
	FSInstances             *map[int64]*mainStruct.FsInstance                      `json:"fs_instances,omitempty"`
	Domains                 *map[int64]*mainStruct.Domain                          `json:"domains,omitempty"`
	DirectoryUsers          interface{}                                            `json:"directory_users,omitempty"`
	User                    *mainStruct.WebUser                                    `json:"user,omitempty"`
	Settings                *cfg.GeneralCfg                                        `json:"settings,omitempty"`
	Daemon                  *mainStruct.DaemonState                                `json:"daemon,omitempty"`
	Items                   *map[int64]string                                      `json:"items,omitempty"`
	List                    *map[int64]map[int64]string                            `json:"list,omitempty"`
	Item                    *Item                                                  `json:"item,omitempty"`
	DomainDetails           *map[int64]DomainDetails                               `json:"domain_details,omitempty"`
	UserDetails             *map[int64]UserDetails                                 `json:"user_details,omitempty"`
	GroupUsers              *map[int64]map[int64]*mainStruct.GroupUser             `json:"group_users,omitempty"`
	UserGateways            *map[int64]map[int64]map[int64]*mainStruct.UserGateway `json:"user_gateways,omitempty"`
	GatewayDetails          *map[int64]*GatewayDetails                             `json:"gateway_details,omitempty"`
	Error                   string                                                 `json:"error,omitempty"`
	Id                      *int64                                                 `json:"id,omitempty"`
	AffectedId              *int64                                                 `json:"affected_id,omitempty"`
	AclLists                *map[int64]*mainStruct.List                            `json:"acl_lists,omitempty"`
	AclListNodes            *map[int64]*mainStruct.Node                            `json:"acl_list_nodes,omitempty"`
	SofiaProfiles           *map[int64]*mainStruct.SofiaProfile                    `json:"sofia_profiles,omitempty"`
	Parameters              *map[int64]*mainStruct.Param                           `json:"parameters,omitempty"`
	SofiaProfileParams      *map[int64]*mainStruct.SofiaProfileParam               `json:"sofia_profile_params,omitempty"`
	SofiaGateways           *map[int64]map[int64]*mainStruct.SofiaGateway          `json:"sofia_gateways,omitempty"`
	SofiaGatewayDetails     *map[int64]*SofiaGatewayDetails                        `json:"sofia_gateway_details,omitempty"`
	SofiaDomains            *map[int64]*mainStruct.SofiaDomain                     `json:"sofia_profile_domains,omitempty"`
	SofiaAliases            *map[int64]*mainStruct.Alias                           `json:"sofia_profile_aliases,omitempty"`
	Modules                 *altStruct.Configurations                              `json:"modules,omitempty"`
	Module                  *altStruct.Configurations                              `json:"module,omitempty"`
	Exists                  *bool                                                  `json:"exists,omitempty"`
	DialplanContexts        *map[int64]*mainStruct.Context                         `json:"dialplan_contexts,omitempty"`
	DialplanSettings        *mainStruct.Dialplans                                  `json:"dialplan_settings,omitempty"`
	DialplanExtensions      *[]*mainStruct.Extension                               `json:"dialplan_extensions,omitempty"`
	DialplanConditions      *map[int64]*[]*mainStruct.Condition                    `json:"dialplan_conditions,omitempty"`
	DialplanDetails         *map[int64]*DialplanDetails                            `json:"dialplan_details,omitempty"`
	Dashboard               *mainStruct.Dashboard                                  `json:"dashboard_data,omitempty"`
	WebUsers                *map[int64]*mainStruct.WebUser                         `json:"web_users,omitempty"`
	WebUsersGroups          *map[int]mainStruct.WebUserGroup                       `json:"web_users_groups,omitempty"`
	Options                 *[]string                                              `json:"options,omitempty"`
	AltOptions              *[]string                                              `json:"alt_options,omitempty"`
	CdrPgCsvSchema          *map[int64]*mainStruct.Field                           `json:"cdr_pg_csv_schema,omitempty"`
	VertoProfiles           *map[int64]*mainStruct.VertoProfile                    `json:"verto_profiles,omitempty"`
	VertoProfileParams      *map[int64]*mainStruct.VertoProfileParam               `json:"verto_profile_params,omitempty"`
	LcrProfiles             *map[int64]*mainStruct.LcrProfile                      `json:"lcr_profiles,omitempty"`
	LcrProfileParams        *map[int64]*mainStruct.LcrProfileParam                 `json:"lcr_profile_params,omitempty"`
	PhoneCreds              *PhoneCreds                                            `json:"phone_creds,omitempty"`
	NoToken                 *bool                                                  `json:"no_token,omitempty"`
	DialplanDebug           *mainStruct.DialplanDebug                              `json:"dialplan_debug,omitempty"`
	Enabled                 *bool                                                  `json:"enabled,omitempty"`
	CallcenterQueues        *map[int64]*mainStruct.Queue                           `json:"callcenter_queues,omitempty"`
	CallcenterQueuesParams  *map[int64]*mainStruct.QueueParam                      `json:"callcenter_queues_params,omitempty"`
	CallcenterAgentsList    *map[int64]*mainStruct.Agent                           `json:"callcenter_agents_list,omitempty"`
	CallcenterAgents        *[]*mainStruct.Agent                                   `json:"callcenter_agents,omitempty"`
	CallcenterAgent         *mainStruct.CCAgentLight                               `json:"callcenter_agent,omitempty"`
	CallcenterTiersList     *map[int64]*mainStruct.Tier                            `json:"callcenter_tiers_list,omitempty"`
	CallcenterTiers         *[]*mainStruct.Tier                                    `json:"callcenter_tiers,omitempty"`
	CallcenterTier          *mainStruct.CCTierLight                                `json:"callcenter_tier,omitempty"`
	CallcenterMembers       *[]*mainStruct.Member                                  `json:"callcenter_members,omitempty"`
	Total                   *int                                                   `json:"total,omitempty"`
	AltTotal                *int                                                   `json:"alt_total,omitempty"`
	OdbcCdrTable            *map[int64]*mainStruct.Table                           `json:"odbc_cdr_tables,omitempty"`
	OdbcCdrTableField       *map[int64]map[int64]*mainStruct.ODBCField             `json:"odbc_cdr_table_fields,omitempty"`
	WebSettings             *map[string]string                                     `json:"web_settings,omitempty"`
	PostSwitchParameters    *map[int64]*mainStruct.Param                           `json:"post_load_switch_parameters,omitempty"`
	PostSwitchCliKeybinding *map[int64]*mainStruct.Param                           `json:"post_load_switch_cli_keybinding,omitempty"`
	PostSwitchDefaultPtime  *map[int64]*mainStruct.DefaultPtime                    `json:"post_load_switch_default_ptime,omitempty"`
	DistributorLists        *map[int64]*mainStruct.DistributorList                 `json:"distributor_lists,omitempty"`
	DistributorListNodes    *map[int64]*mainStruct.DistributorNode                 `json:"distributor_list_nodes,omitempty"`

	DirectoryProfiles                 *map[int64]*mainStruct.DirectoryProfile                       `json:"directory_profiles,omitempty"`
	DirectoryProfileParams            *map[int64]*mainStruct.DirectoryProfileParam                  `json:"directory_profile_params,omitempty"`
	FifoFifo                          *map[int64]*mainStruct.FifoFifo                               `json:"fifo_fifos,omitempty"`
	FifoFifosMembers                  *map[int64]*mainStruct.FifoFifoMember                         `json:"fifo_fifo_members,omitempty"`
	OpalListeners                     *map[int64]*mainStruct.OpalListener                           `json:"opal_listeners,omitempty"`
	OpalListenerParams                *map[int64]*mainStruct.OpalListenerParam                      `json:"opal_listener_params,omitempty"`
	OspProfiles                       *map[int64]*mainStruct.OspProfile                             `json:"osp_profiles,omitempty"`
	OspProfileParams                  *map[int64]*mainStruct.OspProfileParam                        `json:"osp_profile_params,omitempty"`
	UnicallSpans                      *map[int64]*mainStruct.UnicallSpan                            `json:"unicall_spans,omitempty"`
	UnicallSpanParams                 *map[int64]*mainStruct.UnicallSpanParam                       `json:"unicall_span_params,omitempty"`
	ConferenceRooms                   *map[int64]*mainStruct.ConfigConferenceAdvertiseRooms         `json:"conference_rooms,omitempty"`
	ConferenceCallerControlGroups     *map[int64]*mainStruct.ConfigConferenceCallerControlsGroups   `json:"conference_caller_control_groups,omitempty"`
	ConferenceCallerControl           *map[int64]*mainStruct.ConfigConferenceCallerControlsControls `json:"conference_caller_controls,omitempty"`
	ConferenceProfiles                *map[int64]*mainStruct.ConfigConferenceProfiles               `json:"conference_profiles,omitempty"`
	ConferenceProfileParams           *map[int64]*mainStruct.ConfigConferenceProfilesParams         `json:"conference_profile_params,omitempty"`
	ConferenceChatPermissionsUsers    *map[int64]*mainStruct.ConfigConferenceChatPermissionUsers    `json:"conference_chat_permissions_users,omitempty"`
	ConferenceChatPermissionsProfiles *map[int64]*mainStruct.ConfigConferenceChatPermissions        `json:"conference_chat_permissions_profiles,omitempty"`

	CDR            *[]map[string]interface{}  `json:"cdr,omitempty"`
	Logs           *[]map[string]*interface{} `json:"logs,omitempty"`
	HEPs           *[]map[string]*interface{} `json:"heps,omitempty"`
	HEPsDetails    *[]map[string]*interface{} `json:"hep_details,omitempty"`
	AdditionalData *map[int64]*interface{}    `json:"additional_data,omitempty"`

	GlobalVariables             *map[int64]*mainStruct.GlobalVariable             `json:"global_variables,omitempty"`
	ModuleTags                  *map[int64]*mainStruct.ModuleTag                  `json:"module_tags,omitempty"`
	VoicemailSettings           *map[int64]*mainStruct.VoicemailSettingsParameter `json:"voicemail_settings,omitempty"`
	VoicemailProfiles           *map[int64]*mainStruct.VoicemailProfile           `json:"voicemail_profiles,omitempty"`
	VoicemailProfilesParameters *map[int64]*mainStruct.VoicemailProfilesParameter `json:"voicemail_profiles_parameters,omitempty"`

	Data interface{} `json:"data,omitempty"`
}

type PhoneCreds struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Domain    string `json:"domain"`
	WebRTCLib string `json:"webrtc_lib"`
	Ws        string `json:"ws"`
	VertoWs   string `json:"verto_ws"`
	Stun      string `json:"stun"`
}

type DialplanDetails struct {
	//DialplanActions     *map[int64]*mainStruct.Action                          `json:"actions,omitempty"`
	//DialplanAntiActions *map[int64]*mainStruct.AntiAction                      `json:"antiactions,omitempty"`
	DialplanActions     *[]*mainStruct.Action     `json:"actions,omitempty"`
	DialplanAntiActions *[]*mainStruct.AntiAction `json:"antiactions,omitempty"`
	DialplanRegexes     *[]*mainStruct.Regex      `json:"regexes,omitempty"`
}

type DomainDetails struct {
	Params map[int64]*mainStruct.DomainParam    `json:"parameters,omitempty"`
	Vars   map[int64]*mainStruct.DomainVariable `json:"variables,omitempty"`
}

type GatewayDetails struct {
	Params map[int64]*mainStruct.GatewayParam    `json:"parameters,omitempty"`
	Vars   map[int64]*mainStruct.GatewayVariable `json:"variables,omitempty"`
}

type SofiaGatewayDetails struct {
	Params map[int64]*mainStruct.SofiaGatewayParam    `json:"parameters,omitempty"`
	Vars   map[int64]*mainStruct.SofiaGatewayVariable `json:"variables,omitempty"`
}

type UserDetails struct {
	Params      map[int64]*mainStruct.UserParam    `json:"parameters,omitempty"`
	Vars        map[int64]*mainStruct.UserVariable `json:"variables,omitempty"`
	Cache       uint                               `json:"cache"`
	Cidr        string                             `json:"cidr"`
	NumberAlias string                             `json:"number_alias"`
}

type Item struct {
	Id       int64  `json:"id,omitempty"`
	Position int64  `json:"position,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	Type     string `json:"type,omitempty"`
	Cidr     string `json:"cidr,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Weight   string `json:"weight,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
}

type ErroeResponse struct {
	Error       string `json:"error,omitempty"`
	MessageType string `json:"MessageType"`
}

type WsContext struct {
	ws            *websocket.Conn
	Subscriptions *Subscriptions
	SendChannel   chan *UserResponse
}

func (c *WsContext) Close() {
	c.ws.Close()
	c.SendChannel = nil
}

func (c *WsContext) SendWaiter() {
	for event := range c.SendChannel {
		// log.Printf("[WEBSOCKET] Send message: %+v\n", v)
		err := c.ws.SetWriteDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			fmt.Printf("ERROR on SetWriteDeadline: %+v\n", err)
			c.Close()
			return
		}
		eventMsg, err := json.Marshal(event)
		if err != nil {
			fmt.Printf("ERROR on Marshal: %+v\n", err)
			continue
		}

		err = c.ws.WriteMessage(websocket.TextMessage, eventMsg)
		// err = c.ws.WriteJSON(event)
		if err != nil {
			// fmt.Printf("ERROR on WriteJSON: %+v\n", err)
			fmt.Printf("ERROR on WriteMessage: %+v\n", err)
			c.Close()
			return
		}
	}
}

func (c *WsContext) ReadWaiter(handler func(*Message, *WsContext)) {
	for {
		var msg *Message
		_, r, err := c.ws.NextReader()
		if err != nil {
			fmt.Printf("CLOSE WS BY ERROR on NextReader: %+v\n", err)
			c.Close()
			return
		}

		err = json.NewDecoder(r).Decode(&msg)
		if err == io.EOF {
			fmt.Printf("EOF reader\n")
			err = io.ErrUnexpectedEOF
		}

		if err != nil {
			fmt.Printf("ERROR on Decode: %+v\n", err)
			c.SendChannel <- &UserResponse{Error: "failed to read message", MessageType: "none"}
			continue
		}
		//messageHandler
		go handler(msg, c)
		// log.Printf("[WEBSOCKET] Got message: %+v\n", msg)
	}
}

type WsHub struct {
	Hub []*WsContext
}

func (h *WsHub) Broadcast(data UserResponse) {
	for i := 0; i < len(h.Hub); i++ {
		if h.Hub[i] == nil || h.Hub[i].SendChannel == nil {
			h.Drop(i)
			i--
			continue
		}

		subscribed := h.Hub[i].Subscriptions.Get(data.MessageType)
		if !subscribed && data.MessageType != BroadcastConnection {
			// fmt.Println("NOT Subscribed")
			continue
		}

		select {
		case h.Hub[i].SendChannel <- &data:
			// fmt.Println("sent message", msg)
		default:
			h.Hub[i].Close()
			fmt.Println("Cant Send")
			h.Drop(i)
			i--
		}
	}
}

func (h *WsHub) Drop(index int) {
	if len(h.Hub) == index+1 {
		h.Hub = h.Hub[:index]
		return
	}
	h.Hub = append(h.Hub[:index], h.Hub[index+1:]...)
}

func (m *MessageData) Trim() {
	m.Login = strings.TrimSpace(m.Login)
	//m.Id = strings.TrimSpace(m.Id)
	m.Name = strings.TrimSpace(m.Name)
	m.Value = strings.TrimSpace(m.Value)
	m.Default = strings.TrimSpace(m.Default)
	m.Uuid = strings.TrimSpace(m.Uuid)
	m.Node.Name = strings.TrimSpace(m.Node.Name)
	m.Node.Type = strings.TrimSpace(m.Node.Type)
	m.Node.Domain = strings.TrimSpace(m.Node.Domain)
	m.Node.Cidr = strings.TrimSpace(m.Node.Cidr)
	m.Param.Name = strings.TrimSpace(m.Param.Name)
	m.Param.Value = strings.TrimSpace(m.Param.Value)
	m.Param.Secure = strings.TrimSpace(m.Param.Secure)
	m.SofiaDomain.Name = strings.TrimSpace(m.SofiaDomain.Name)
	m.SofiaAlias.Name = strings.TrimSpace(m.SofiaAlias.Name)
	m.Variable.Name = strings.TrimSpace(m.Variable.Name)
	m.Variable.Value = strings.TrimSpace(m.Variable.Value)
	m.Variable.Direction = strings.TrimSpace(m.Variable.Direction)
	m.Regex.Field = strings.TrimSpace(m.Regex.Field)
	m.Regex.Expression = strings.TrimSpace(m.Regex.Expression)
	m.Action.Data = strings.TrimSpace(m.Action.Data)
	m.Action.Application = strings.TrimSpace(m.Action.Application)
	m.AntiAction.Data = strings.TrimSpace(m.AntiAction.Data)
	m.AntiAction.Application = strings.TrimSpace(m.AntiAction.Application)
	m.Condition.Expression = strings.TrimSpace(m.Condition.Expression)
	m.Condition.Field = strings.TrimSpace(m.Condition.Field)
	m.Condition.Regex = strings.TrimSpace(m.Condition.Regex)
	m.Condition.Break = strings.TrimSpace(m.Condition.Break)
	m.Condition.Dst = strings.TrimSpace(m.Condition.Dst)
	m.Condition.Year = strings.TrimSpace(m.Condition.Year)
	m.Condition.TzOffset = strings.TrimSpace(m.Condition.TzOffset)
	m.Condition.DateTime = strings.TrimSpace(m.Condition.DateTime)
	m.Condition.Minday = strings.TrimSpace(m.Condition.Minday)
	m.Condition.Yday = strings.TrimSpace(m.Condition.Yday)
	m.Condition.Week = strings.TrimSpace(m.Condition.Week)
	m.Condition.Minute = strings.TrimSpace(m.Condition.Minute)
	m.Condition.TimeOfDay = strings.TrimSpace(m.Condition.TimeOfDay)
	m.Condition.Mday = strings.TrimSpace(m.Condition.Mday)
	m.Condition.Hour = strings.TrimSpace(m.Condition.Hour)
	m.Condition.Mon = strings.TrimSpace(m.Condition.Mon)
	m.Condition.Mweek = strings.TrimSpace(m.Condition.Mweek)
	m.Condition.Wday = strings.TrimSpace(m.Condition.Wday)
	m.Field.Column = strings.TrimSpace(m.Field.Column)
	m.Field.Var = strings.TrimSpace(m.Field.Var)
	m.Table.Name = strings.TrimSpace(m.Table.Name)
	m.Table.LogLeg = strings.TrimSpace(m.Table.LogLeg)
	m.OdbcCdrField.Name = strings.TrimSpace(m.OdbcCdrField.Name)
	m.OdbcCdrField.ChanVarName = strings.TrimSpace(m.OdbcCdrField.ChanVarName)
	m.Importance = strings.TrimSpace(m.Importance)
	m.FifoFifoMember.Timeout = strings.TrimSpace(m.FifoFifoMember.Timeout)
	m.FifoFifoMember.Simo = strings.TrimSpace(m.FifoFifoMember.Simo)
	m.FifoFifoMember.Lag = strings.TrimSpace(m.FifoFifoMember.Lag)
	m.FifoFifoMember.Body = strings.TrimSpace(m.FifoFifoMember.Body)
	m.DistributorNode.Name = strings.TrimSpace(m.DistributorNode.Name)
	m.DistributorNode.Weight = strings.TrimSpace(m.DistributorNode.Weight)
}

func newSubscriptions() *Subscriptions {
	return &Subscriptions{byName: make(map[string]bool)}
}

func CreateWsContext(ws *websocket.Conn) *WsContext {
	return &WsContext{
		ws:            ws,
		Subscriptions: newSubscriptions(),
		SendChannel:   make(chan *UserResponse, 42), //capacity - or got disconnect on 2 msg at one time
	}
}
