package webStruct

import (
	"custompbx/altStruct"
	"custompbx/cfg"
	"custompbx/mainStruct"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"strings"
	"sync"
	"sync/atomic"
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
	Username string       `json:"username"`
	Event    string       `json:"event"`
	Data     *MessageData `json:"data"`
}

var ErrInvalidMessage = errors.New("invalid websocket message")

func (m *Message) Validate() error {
	if m == nil || m.Data == nil || strings.TrimSpace(m.Event) == "" {
		return ErrInvalidMessage
	}
	return nil
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
	LayoutImages     altStruct.ConfigConferenceLayoutImage     `json:"layout_images,omitempty"`

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
	mx         sync.RWMutex
	byName     map[string]bool
	persistent map[string]bool
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

// SetPersistent adds a persistent subscription for a given name.
func (s *Subscriptions) SetPersistent(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[name] = true
	s.persistent[name] = true
}

// Del removes a subscription for a given name.
func (s *Subscriptions) Del(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, name)
	delete(s.persistent, name)
}

// Clear removes all subscriptions.
func (s *Subscriptions) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName = make(map[string]bool, len(s.persistent))
	for name := range s.persistent {
		s.byName[name] = true
	}
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
	ws               *websocket.Conn
	Subscriptions    *Subscriptions
	ID               uint64
	User             *mainStruct.WebUser
	userID           atomic.Int64
	send             chan *UserResponse
	done             chan struct{}
	closeOnce        sync.Once
	onClose          func(*WsContext)
	onWriteFailure   func()
	onHandlerFailure func()
	writeTimeout     time.Duration
	readTimeout      time.Duration
	pingInterval     time.Duration
}

// Close closes the WebSocket connection and clears the send channel.
func (c *WsContext) Close() error {
	return c.CloseWithReason("closed")

}

func (c *WsContext) CloseWithReason(reason string) (closeErr error) {
	c.closeOnce.Do(func() {
		close(c.done)
		if c.ws != nil {
			closeErr = c.ws.Close()
		}
		if c.onClose != nil {
			c.onClose(c)
		}
		log.Printf("component=websocket connection_id=%d user_id=%d operation=close reason=%q", c.ID, c.UserID(), reason)
	})
	return closeErr
}

func (c *WsContext) SetUser(user *mainStruct.WebUser) {
	c.User = user
	if user == nil {
		c.userID.Store(0)
		return
	}
	c.userID.Store(user.Id)
}

func (c *WsContext) UserID() int64 {
	return c.userID.Load()
}

func (c *WsContext) Enqueue(event *UserResponse) bool {
	select {
	case <-c.done:
		return false
	default:
	}
	select {
	case c.send <- event:
		return true
	default:
		return false
	}
}

func (c *WsContext) RecordHandlerFailure() {
	if c.onHandlerFailure != nil {
		c.onHandlerFailure()
	}
}

// SendWaiter listens on the SendChannel and sends messages through the WebSocket connection.
func (c *WsContext) SendWaiter() {
	ping := time.NewTicker(c.pingInterval)
	defer ping.Stop()
	for {
		select {
		case <-c.done:
			return
		case <-ping.C:
			if err := c.ws.WriteControl(websocket.PingMessage, nil, time.Now().Add(c.writeTimeout)); err != nil {
				if c.onWriteFailure != nil {
					c.onWriteFailure()
				}
				c.CloseWithReason("ping failure")
				return
			}
		case event := <-c.send:
			// log.Printf("[WEBSOCKET] Send message: %+v\n", v)
			err := c.ws.SetWriteDeadline(time.Now().Add(c.writeTimeout))
			if err != nil {
				if c.onWriteFailure != nil {
					c.onWriteFailure()
				}
				log.Printf("component=websocket connection_id=%d user_id=%d operation=set_write_deadline error=%q", c.ID, c.UserID(), err)
				c.CloseWithReason("write deadline failure")
				return
			}
			eventMsg, err := json.Marshal(event)
			if err != nil {
				if c.onWriteFailure != nil {
					c.onWriteFailure()
				}
				log.Printf("component=websocket connection_id=%d user_id=%d operation=marshal_response error=%q", c.ID, c.UserID(), err)
				continue
			}

			err = c.ws.WriteMessage(websocket.TextMessage, eventMsg)
			// err = c.ws.WriteJSON(event)
			if err != nil {
				// fmt.Printf("ERROR on WriteJSON: %+v\n", err)
				log.Printf("component=websocket connection_id=%d user_id=%d operation=write_message error=%q", c.ID, c.UserID(), err)
				c.CloseWithReason("write failure")
				return
			}
		}
	}
}

// ReadWaiter listens for incoming messages on the WebSocket connection and handles them using the provided handler function.
func (c *WsContext) ReadWaiter(handler func(*Message, *WsContext)) {
	c.ws.SetReadLimit(1 << 20)
	_ = c.ws.SetReadDeadline(time.Now().Add(c.readTimeout))
	c.ws.SetPongHandler(func(string) error { return c.ws.SetReadDeadline(time.Now().Add(c.readTimeout)) })
	for {
		var msg *Message
		_, r, err := c.ws.NextReader()
		if err != nil {
			log.Printf("component=websocket connection_id=%d user_id=%d operation=next_reader error=%q", c.ID, c.UserID(), err)
			c.CloseWithReason("read failure")
			return
		}

		err = json.NewDecoder(r).Decode(&msg)
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}

		if err != nil {
			log.Printf("component=websocket connection_id=%d user_id=%d operation=decode_message error=%q", c.ID, c.UserID(), err)
			if !c.Enqueue(&UserResponse{Error: "failed to read message", MessageType: "none"}) {
				c.RecordHandlerFailure()
				c.CloseWithReason("outbound queue full")
			}
			continue
		}
		if err := msg.Validate(); err != nil {
			log.Printf("component=websocket connection_id=%d user_id=%d operation=validate_message error=%q", c.ID, c.UserID(), err)
			if !c.Enqueue(&UserResponse{Error: "invalid message", MessageType: "none"}) {
				c.RecordHandlerFailure()
				c.CloseWithReason("outbound queue full")
			}
			continue
		}
		handler(msg, c)
		// log.Printf("[WEBSOCKET] Got message: %+v\n", msg)
	}
}

// WsHub manages a collection of WebSocket contexts and provides broadcast capabilities.
type WsHub struct {
	mx              sync.RWMutex
	connections     map[uint64]*WsContext
	active          atomic.Int64
	broadcasts      atomic.Uint64
	queueOverflows  atomic.Uint64
	failedWrites    atomic.Uint64
	handlerFailures atomic.Uint64
	shuttingDown    atomic.Bool
	shutdownOnce    sync.Once
}

type HubMetrics struct {
	Active          int64  `json:"active"`
	Broadcasts      uint64 `json:"broadcasts"`
	QueueOverflows  uint64 `json:"queue_overflows"`
	FailedWrites    uint64 `json:"failed_writes"`
	HandlerFailures uint64 `json:"handler_failures"`
	ShuttingDown    bool   `json:"shutting_down"`
}

func NewWsHub() *WsHub { return &WsHub{connections: make(map[uint64]*WsContext)} }

func (h *WsHub) Register(c *WsContext) {
	if c == nil {
		return
	}
	if h.shuttingDown.Load() {
		_ = c.CloseWithReason("hub shutdown")
		return
	}
	h.mx.Lock()
	if h.shuttingDown.Load() {
		h.mx.Unlock()
		_ = c.CloseWithReason("hub shutdown")
		return
	}
	if h.connections == nil {
		h.connections = make(map[uint64]*WsContext)
	}
	if existing := h.connections[c.ID]; existing != nil {
		h.active.Store(int64(len(h.connections)))
		h.mx.Unlock()
		if existing != c {
			_ = c.CloseWithReason("duplicate connection id")
		}
		return
	}
	h.connections[c.ID] = c
	c.onClose = h.unregister
	c.onWriteFailure = func() { h.failedWrites.Add(1) }
	c.onHandlerFailure = func() { h.handlerFailures.Add(1) }
	h.active.Store(int64(len(h.connections)))
	h.mx.Unlock()
}

func (h *WsHub) unregister(c *WsContext) {
	if c == nil {
		return
	}
	h.mx.Lock()
	if h.connections[c.ID] == c {
		delete(h.connections, c.ID)
	}
	h.active.Store(int64(len(h.connections)))
	h.mx.Unlock()
}

func (h *WsHub) snapshot() []*WsContext {
	h.mx.RLock()
	defer h.mx.RUnlock()
	items := make([]*WsContext, 0, len(h.connections))
	for _, c := range h.connections {
		items = append(items, c)
	}
	return items
}

func (h *WsHub) Metrics() HubMetrics {
	return HubMetrics{
		Active:          h.active.Load(),
		Broadcasts:      h.broadcasts.Load(),
		QueueOverflows:  h.queueOverflows.Load(),
		FailedWrites:    h.failedWrites.Load(),
		HandlerFailures: h.handlerFailures.Load(),
		ShuttingDown:    h.shuttingDown.Load(),
	}
}

func (h *WsHub) Shutdown() {
	h.shutdownOnce.Do(func() {
		h.shuttingDown.Store(true)
		for _, c := range h.snapshot() {
			_ = c.CloseWithReason("hub shutdown")
		}
	})
}

// Broadcast sends a UserResponse message to all subscribed WebSocket contexts in the hub.
func (h *WsHub) Broadcast(data UserResponse) {
	if h.shuttingDown.Load() {
		return
	}
	h.broadcasts.Add(1)
	for _, connection := range h.snapshot() {
		subscribed := connection.Subscriptions.Get(data.MessageType)
		if !subscribed && data.MessageType != BroadcastConnection {
			// fmt.Println("NOT Subscribed")
			continue
		}

		h.enqueueOrDrop(connection, &data)
	}
}

// Unicast sends a UserResponse message to all WebSocket contexts belonging to one or more specified users.
func (h *WsHub) Unicast(data UserResponse, users []*mainStruct.WebUser) []int64 {
	userIDs := make(map[int64]bool)
	for _, user := range users {
		if user == nil {
			continue
		}
		userIDs[user.Id] = true
	}

	var sent []int64
	for _, connection := range h.snapshot() {
		connectionUserID := connection.UserID()
		if connectionUserID != 0 && userIDs[connectionUserID] {
			if h.enqueueOrDrop(connection, &data) {
				sent = append(sent, connectionUserID)
			}
		}
	}
	return sent
}

func (h *WsHub) enqueueOrDrop(connection *WsContext, data *UserResponse) bool {
	if connection.Enqueue(data) {
		return true
	}
	h.queueOverflows.Add(1)
	_ = connection.CloseWithReason("outbound queue full")
	return false
}

// Drop removes a WebSocket context at the specified index from the hub.
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
	return &Subscriptions{
		byName:     make(map[string]bool),
		persistent: make(map[string]bool),
	}
}

func durationFromConfig(seconds int, fallback time.Duration) time.Duration {
	if seconds <= 0 {
		return fallback
	}
	return time.Duration(seconds) * time.Second
}

// CreateWsContext creates a new WsContext with the provided WebSocket connection.
func CreateWsContext(ws *websocket.Conn) *WsContext {
	queueSize := cfg.NormalizeWebSocketQueueSize(cfg.CustomPbx.Web.WebSocketQueueSize)
	return &WsContext{
		ws:            ws,
		ID:            nextConnectionID.Add(1),
		Subscriptions: newSubscriptions(),
		send:          make(chan *UserResponse, queueSize),
		done:          make(chan struct{}),
		writeTimeout:  durationFromConfig(cfg.CustomPbx.Web.WriteTimeoutSeconds, 10*time.Second),
		readTimeout:   durationFromConfig(cfg.CustomPbx.Web.ReadTimeoutSeconds, 60*time.Second),
		pingInterval:  durationFromConfig(cfg.CustomPbx.Web.PingIntervalSeconds, 30*time.Second),
	}
}

var nextConnectionID atomic.Uint64
