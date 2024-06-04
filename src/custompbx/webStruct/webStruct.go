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
	Login            string                                    `json:"login"`
	Payload          cfg.GeneralCfg                            `json:"payload,omitempty"`
	Password         string                                    `json:"password"`
	Token            string                                    `json:"token,omitempty"`
	Id               int64                                     `json:"id,omitempty"`
	IdInt            int64                                     `json:"id_int,omitempty"`
	Name             string                                    `json:"name,omitempty"`
	Value            string                                    `json:"value,omitempty"`
	Direction        string                                    `json:"direction,omitempty"`
	Default          string                                    `json:"default,omitempty"`
	Node             altStruct.ConfigAclNode                   `json:"node,omitempty"`
	Param            Param                                     `json:"param,omitempty"`
	SofiaDomain      altStruct.ConfigSofiaProfileDomain        `json:"sofia_domain,omitempty"`
	SofiaAlias       altStruct.ConfigSofiaProfileAlias         `json:"sofia_alias,omitempty"`
	Variable         MessageDataVariable                       `json:"variable,omitempty"`
	Regex            mainStruct.Regex                          `json:"regex,omitempty"`
	Action           mainStruct.Action                         `json:"action,omitempty"`
	AntiAction       mainStruct.AntiAction                     `json:"antiaction,omitempty"`
	Condition        mainStruct.Condition                      `json:"condition,omitempty"`
	Enabled          *bool                                     `json:"enabled,omitempty"`
	KeepSubscription *bool                                     `json:"keep_subscription,omitempty"`
	PreviousIndex    int64                                     `json:"previous_index,omitempty"`
	CurrentIndex     int64                                     `json:"current_index,omitempty"`
	ParamId          int64                                     `json:"param_id,omitempty"`
	Bulk             int                                       `json:"bulk,omitempty"`
	GroupId          int                                       `json:"group_id,omitempty"`
	Field            altStruct.ConfigCdrPgCsvSchema            `json:"field,omitempty"`
	DBRequest        mainStruct.DBRequest                      `json:"db_request,omitempty"`
	File             string                                    `json:"file,omitempty"`
	Uuid             string                                    `json:"uuid,omitempty"`
	ArrVal           []string                                  `json:"values,omitempty"`
	Table            altStruct.ConfigOdbcCdrTable              `json:"table,omitempty"`
	OdbcCdrField     altStruct.ConfigOdbcCdrTableField         `json:"odbc_cdr_field,omitempty"`
	WebSettings      map[string]string                         `json:"web_settings,omitempty"`
	DistributorNode  altStruct.ConfigDistributorListNode       `json:"distributor_node,omitempty"`
	Importance       string                                    `json:"importance,omitempty"`
	FifoFifoMember   altStruct.ConfigFifoFifoMember            `json:"fifo_fifo_member,omitempty"`
	Fields           []string                                  `json:"fields,omitempty"`
	Ids              []int64                                   `json:"ids,omitempty"`
	AwsS3            altStruct.ConfigHttpCacheProfileAWSS3     `json:"aws_s3,omitempty"`
	Azure            altStruct.ConfigHttpCacheProfileAzureBlob `json:"azure,omitempty"`

	Text     string    `json:"text,omitempty"`
	UpToTime time.Time `json:"up_to_time,omitempty"`

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

// Subscriptions manages a set of subscriptions with thread-safe operations.
type Subscriptions struct {
	mx     sync.RWMutex
	byName map[string]bool
}

// Get retrieves the subscription status for a given name.
func (s *Subscriptions) Get(name string) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.byName[name]
}

// Set adds a subscription for a given name.
func (s *Subscriptions) Set(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[name] = true
}

// Del removes a subscription for a given name.
func (s *Subscriptions) Del(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if !s.byName[name] {
		return
	}
	s.byName[name] = false
}

// Clear removes all subscriptions.
func (s *Subscriptions) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName = make(map[string]bool)
}

type Param struct {
	Id      int64  `json:"id"`
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Secure  string `json:"secure"`
}

type UserResponse struct {
	MessageType          string                              `json:"MessageType"`
	Token                string                              `json:"token,omitempty"`
	TokensList           *[]mainStruct.WebUserToken          `json:"tokens_list,omitempty"`
	Uuid                 *string                             `json:"uuid,omitempty"`
	Response             *string                             `json:"response,omitempty"`
	FSInstances          *map[int64]*mainStruct.FsInstance   `json:"fs_instances,omitempty"`
	DirectoryUsers       interface{}                         `json:"directory_users,omitempty"`
	User                 *mainStruct.WebUser                 `json:"user,omitempty"`
	Settings             *cfg.GeneralCfg                     `json:"settings,omitempty"`
	Daemon               *mainStruct.DaemonState             `json:"daemon,omitempty"`
	Items                *map[int64]string                   `json:"items,omitempty"`
	List                 *map[int64]map[int64]string         `json:"list,omitempty"`
	Item                 *Item                               `json:"item,omitempty"`
	Error                string                              `json:"error,omitempty"`
	Id                   *int64                              `json:"id,omitempty"`
	AffectedId           *int64                              `json:"affected_id,omitempty"`
	Modules              *altStruct.Configurations           `json:"modules,omitempty"`
	Module               *altStruct.Configurations           `json:"module,omitempty"`
	Exists               *bool                               `json:"exists,omitempty"`
	DialplanContexts     *map[int64]*mainStruct.Context      `json:"dialplan_contexts,omitempty"`
	DialplanSettings     *mainStruct.Dialplans               `json:"dialplan_settings,omitempty"`
	DialplanExtensions   *[]*mainStruct.Extension            `json:"dialplan_extensions,omitempty"`
	DialplanConditions   *map[int64]*[]*mainStruct.Condition `json:"dialplan_conditions,omitempty"`
	DialplanDetails      *map[int64]*DialplanDetails         `json:"dialplan_details,omitempty"`
	Dashboard            *mainStruct.Dashboard               `json:"dashboard_data,omitempty"`
	WebUsers             *map[int64]*mainStruct.WebUser      `json:"web_users,omitempty"`
	WebUsersGroups       *map[int]mainStruct.WebUserGroup    `json:"web_users_groups,omitempty"`
	Options              *[]string                           `json:"options,omitempty"`
	AltOptions           *[]string                           `json:"alt_options,omitempty"`
	PhoneCreds           *PhoneCreds                         `json:"phone_creds,omitempty"`
	NoToken              *bool                               `json:"no_token,omitempty"`
	DialplanDebug        *mainStruct.DialplanDebug           `json:"dialplan_debug,omitempty"`
	Enabled              *bool                               `json:"enabled,omitempty"`
	CallcenterAgentsList *map[int64]*mainStruct.Agent        `json:"callcenter_agents_list,omitempty"`
	Total                *int                                `json:"total,omitempty"`
	AltTotal             *int                                `json:"alt_total,omitempty"`
	WebSettings          *map[string]string                  `json:"web_settings,omitempty"`

	CDR            *[]map[string]interface{}  `json:"cdr,omitempty"`
	Logs           *[]map[string]*interface{} `json:"logs,omitempty"`
	HEPs           *[]map[string]*interface{} `json:"heps,omitempty"`
	HEPsDetails    *[]map[string]*interface{} `json:"hep_details,omitempty"`
	AdditionalData *map[int64]*interface{}    `json:"additional_data,omitempty"`

	GlobalVariables *map[int64]*mainStruct.GlobalVariable `json:"global_variables,omitempty"`

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

// WsContext represents a WebSocket context with its connection, subscriptions, and send channel.
type WsContext struct {
	ws            *websocket.Conn
	Subscriptions *Subscriptions
	SendChannel   chan *UserResponse
	User          *mainStruct.WebUser
}

// Close closes the WebSocket connection and clears the send channel.
func (c *WsContext) Close() error {
	err := c.ws.Close()
	if err != nil {
		return fmt.Errorf("failed to close WebSocket: %w", err)
	}
	c.SendChannel = nil
	return nil
}

// SendWaiter listens on the SendChannel and sends messages through the WebSocket connection.
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

// ReadWaiter listens for incoming messages on the WebSocket connection and handles them using the provided handler function.
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

// WsHub manages a collection of WebSocket contexts and provides broadcast capabilities.
type WsHub struct {
	Hub []*WsContext
}

// Broadcast sends a UserResponse message to all subscribed WebSocket contexts in the hub.
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
			fmt.Println("Can't Send")
			h.Drop(i)
			i--
		}
	}
}

// Unicast sends a UserResponse message to all WebSocket contexts belonging to one or more specified users.
func (h *WsHub) Unicast(data UserResponse, users []*mainStruct.WebUser) {
	userIDs := make(map[int64]bool)
	for _, user := range users {
		userIDs[user.Id] = true
	}

	for i := 0; i < len(h.Hub); i++ {
		if h.Hub[i] == nil || h.Hub[i].SendChannel == nil {
			h.Drop(i)
			i--
			continue
		}

		if h.Hub[i].User != nil && userIDs[h.Hub[i].User.Id] {
			select {
			case h.Hub[i].SendChannel <- &data:
			default:
				h.Hub[i].Close()
				fmt.Println("Can't Send")
				h.Drop(i)
				i--
			}
		}
	}
}

// Drop removes a WebSocket context at the specified index from the hub.
func (h *WsHub) Drop(index int) {
	if len(h.Hub) == index+1 {
		h.Hub = h.Hub[:index]
		return
	}
	h.Hub = append(h.Hub[:index], h.Hub[index+1:]...)
}

// Trim removes leading and trailing white spaces from various fields in MessageData.
func (m *MessageData) Trim() {
	m.Login = strings.TrimSpace(m.Login)
	//m.Id = strings.TrimSpace(m.Id)
	m.Name = strings.TrimSpace(m.Name)
	m.Value = strings.TrimSpace(m.Value)
	m.Default = strings.TrimSpace(m.Default)
	m.Uuid = strings.TrimSpace(m.Uuid)
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

// newSubscriptions creates and returns a new Subscriptions instance.
func newSubscriptions() *Subscriptions {
	return &Subscriptions{byName: make(map[string]bool)}
}

// CreateWsContext creates a new WsContext with the provided WebSocket connection.
func CreateWsContext(ws *websocket.Conn) *WsContext {
	return &WsContext{
		ws:            ws,
		Subscriptions: newSubscriptions(),
		SendChannel:   make(chan *UserResponse, 42), //capacity - or got disconnect on 2 msg at one time
	}
}
