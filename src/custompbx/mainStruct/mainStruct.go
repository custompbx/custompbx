package mainStruct

import (
	"database/sql"
	"sync"
	"time"
)

const (
	WebRTCLibVerto   = "verto"
	WebRTCLibSipJs   = "sipjs"
	Version          = "0.0.102"
	NoProceedName    = "NoProceed"
	CustomPBXVersion = "custompbx_version"
)

// ESL
type CallcenterAgentsJSONResponse struct {
	Status   string       `json:"status"`
	Response []KnownAgent `json:"response"`
}

type CallcenterTiersJSONResponse struct {
	Status   string      `json:"status"`
	Response []KnownTier `json:"response"`
}

type CallcenterMembersJSONResponse struct {
	Status   string   `json:"status"`
	Response []Member `json:"response"`
}

type DBRequest struct {
	Limit   int      `json:"limit,omitempty"`
	Offset  int      `json:"offset,omitempty"`
	Filters []Filter `json:"filters"`
	Order   Order    `json:"order"`
}

type Filter struct {
	Field      string `json:"field"`
	Operand    string `json:"operand"`
	FieldValue string `json:"field_value"`
}

type Order struct {
	Desc   bool     `json:"desc"`
	Fields []string `json:"fields"`
}

// WEB
type WebUser struct {
	Id           int64         `json:"id"`
	Login        string        `json:"login"`
	SipId        sql.NullInt64 `json:"sip_id"`
	WebRTCLib    string        `json:"webrtc_lib"`
	Ws           string        `json:"ws"`
	VertoWs      string        `json:"verto_ws"`
	Stun         string        `json:"stun"`
	Avatar       string        `json:"-"`
	AvatarFormat string        `json:"avatar_format"`
	Enabled      bool          `json:"enabled"`
	Lang         uint          `json:"lang"`
	Key          string        `json:"-"`
	Tokens       WebUserTokens `json:"-"`
	GroupId      int           `json:"group_id"`
}

type WebUserGroup struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//type WebUserGroupId uint

type WebUserToken struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Token   string `json:"token"`
	Created string `json:"created"`
	Purpose string `json:"purpose"`
}

type WebUsers struct {
	mx      sync.RWMutex
	byLogin map[string]*WebUser
	byId    map[int64]*WebUser
	byToken map[string]*WebUser
}

type WebUserTokens struct {
	mx     sync.RWMutex
	tokens map[string]string
}

func (w *WebUserTokens) GetList() []string {
	w.mx.RLock()
	defer w.mx.RUnlock()
	var items []string
	for _, v := range w.tokens {
		items = append(items, v)
	}
	return items
}

func (w *WebUserTokens) Set(token string) {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.tokens[token] = token
}

func (w *WebUserTokens) Delete(token string) {
	w.mx.RLock()
	defer w.mx.RUnlock()
	delete(w.tokens, token)
}

func NewWebUserTokens() WebUserTokens {
	return WebUserTokens{
		tokens: make(map[string]string),
	}
}

func (w *WebUsers) GetByLogin(key string) *WebUser {
	w.mx.RLock()
	defer w.mx.RUnlock()
	return w.byLogin[key]
}

func (w *WebUsers) Count() int {
	w.mx.RLock()
	defer w.mx.RUnlock()
	return len(w.byId)
}

func (w *WebUsers) GetById(key int64) *WebUser {
	w.mx.RLock()
	defer w.mx.RUnlock()
	val := w.byId[key]
	return val
}

func (w *WebUsers) GetByToken(key string) (*WebUser, bool) {
	w.mx.RLock()
	defer w.mx.RUnlock()
	val, ok := w.byToken[key]
	return val, ok
}

func (w *WebUsers) Set(value *WebUser) {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.byLogin[value.Login] = value
	w.byId[value.Id] = value
	for _, v := range value.Tokens.GetList() {
		w.byToken[v] = value
	}
}

func (w *WebUsers) ClearCache(key WebUser) {
	w.mx.RLock()
	defer w.mx.RUnlock()
	delete(w.byLogin, key.Login)
	delete(w.byId, key.Id)
	for _, v := range key.Tokens.GetList() {
		delete(w.byToken, v)
		key.Tokens.Delete(v)
	}
}

func NewUsersCache() *WebUsers {
	return &WebUsers{
		byLogin: make(map[string]*WebUser),
		byId:    make(map[int64]*WebUser),
		byToken: make(map[string]*WebUser),
	}
}

func (w *WebUsers) GetList() map[int64]*WebUser {
	w.mx.RLock()
	defer w.mx.RUnlock()
	list := make(map[int64]*WebUser)
	for _, v := range w.byId {
		list[v.Id] = v
	}
	return list
}

func (w *WebUsers) GetListByDirectory() map[int64]*interface{} {
	w.mx.RLock()
	defer w.mx.RUnlock()
	list := make(map[int64]*interface{})
	for _, v := range w.byId {
		if !v.SipId.Valid {
			continue
		}
		var data interface{}
		data = struct {
			Id           int64  `json:"id"`
			Login        string `json:"login"`
			AvatarFormat string `json:"avatar_format"`
		}{
			Id:           v.Id,
			Login:        v.Login,
			AvatarFormat: v.AvatarFormat,
		}
		list[v.SipId.Int64] = &data
	}
	return list
}

func (w *WebUsers) Remove(key *WebUser) {
	w.mx.Lock()
	defer w.mx.Unlock()
	delete(w.byId, key.Id)
	delete(w.byLogin, key.Login)
	for _, v := range key.Tokens.GetList() {
		delete(w.byToken, v)
		key.Tokens.Delete(v)
	}
}

func (w *WebUsers) Rename(oldName, newName string) {
	w.mx.Lock()
	defer w.mx.Unlock()
	if w.byLogin[oldName] == nil {
		return
	}
	w.byLogin[newName] = w.byLogin[oldName]
	w.byLogin[newName].Login = newName
	delete(w.byLogin, oldName)
}

type Dashboard struct {
	*DashboardData
	*FSMetrics
}

type DashboardData struct {
	Timestamp      time.Time      `json:"timestamp"`
	Hostname       string         `json:"hostname,omitempty"`
	OS             string         `json:"os,omitempty"`
	Platform       string         `json:"platform,omitempty"`
	CPUModel       string         `json:"cpu_model,omitempty"`
	CPUFrequency   float64        `json:"cpu_frequency,omitempty"`
	DynamicMetrics DynamicMetrics `json:"dynamic_metrics"`
}

type DynamicMetrics struct {
	TotalMemory          uint64    `json:"total_memory,omitempty"`
	FreeMemory           uint64    `json:"free_memory,omitempty"`
	PercentageUsedMemory float64   `json:"percentage_used_memory,omitempty"`
	TotalDiscSpace       uint64    `json:"total_disc_space,omitempty"`
	FreeDiskSpace        uint64    `json:"free_disk_space,omitempty"`
	PercentageDiskUsage  float64   `json:"percentage_disk_usage,omitempty"`
	CoreUtilization      []float64 `json:"core_utilization,omitempty"`
}

type FSMetrics struct {
	DomainSipRegs   map[string]int `json:"domain_sip_regs,omitempty"`
	DomainVertoRegs map[string]int `json:"domain_verto_regs,omitempty"`
	CallsCounter    map[string]int `json:"calls_counter,omitempty"`
	SofiaProfiles   interface{}    `json:"sofia_profiles,omitempty"`
	SofiaGateways   interface{}    `json:"sofia_gateways,omitempty"`
}

func NewDashboard() *Dashboard {
	return &Dashboard{&DashboardData{}, &FSMetrics{}}
}

// PBX
type Calls struct {
	mx     sync.RWMutex
	byUuid map[string]*Call
}

type Channels struct {
	mx       sync.RWMutex
	byUuid   map[string]*Channel
	Answered int
	Total    int
}

type Call struct {
	Uuid             string `json:"uuid"`
	Direction        string `json:"direction"`
	Created          string `json:"created"`
	CreatedEpoch     string `json:"created_epoch"`
	Name             string `json:"name"`
	State            string `json:"state"`
	CidName          string `json:"cid_name"`
	CidNum           string `json:"cid_num"`
	IpAddr           string `json:"ip_addr"`
	Dest             string `json:"dest"`
	PresenceId       string `json:"presence_id"`
	PresenceData     string `json:"presence_data"`
	Accountcode      string `json:"accountcode"`
	Callstate        string `json:"callstate"`
	CalleeName       string `json:"callee_name"`
	CalleeNum        string `json:"callee_num"`
	CalleeDirection  string `json:"callee_direction"`
	CallUuid         string `json:"call_uuid"`
	Hostname         string `json:"hostname"`
	SentCalleeName   string `json:"sent_callee_name"`
	SentCalleeNum    string `json:"sent_callee_num"`
	BUuid            string `json:"b_uuid"`
	BDirection       string `json:"b_direction"`
	BCreated         string `json:"b_created"`
	BCreatedEpoch    string `json:"b_created_epoch"`
	BName            string `json:"b_name"`
	BState           string `json:"b_state"`
	BCidName         string `json:"b_cid_name"`
	BCidNum          string `json:"b_cid_num"`
	BIpAddr          string `json:"b_ip_addr"`
	BDest            string `json:"b_dest"`
	BPresenceId      string `json:"b_presence_id"`
	BPresenceData    string `json:"b_presence_data"`
	BAccountcode     string `json:"b_accountcode"`
	BCallstate       string `json:"b_callstate"`
	BCalleeName      string `json:"b_callee_name"`
	BCalleeNum       string `json:"b_callee_num"`
	BCalleeDirection string `json:"b_callee_direction"`
	BSentCalleeName  string `json:"b_sent_callee_name"`
	BSentCalleeNum   string `json:"b_sent_callee_num"`
	CallCreatedEpoch string `json:"call_created_epoch"`
}

type Channel struct {
	Uuid            string `json:"uuid"`
	Direction       string `json:"direction"`
	Created         string `json:"created"`
	CreatedEpoch    string `json:"created_epoch"`
	Name            string `json:"name"`
	State           string `json:"state"`
	CidName         string `json:"cid_name"`
	CidNum          string `json:"cid_num"`
	IpAddr          string `json:"ip_addr"`
	Dest            string `json:"dest"`
	Application     string `json:"application"`
	ApplicationData string `json:"application_data"`
	Dialplan        string `json:"dialplan"`
	Context         string `json:"context"`
	ReadCodec       string `json:"read_codec"`
	ReadRate        string `json:"read_rate"`
	ReadBitRate     string `json:"read_bit_rate"`
	WriteCodec      string `json:"write_codec"`
	WriteRate       string `json:"write_rate"`
	WriteBitRate    string `json:"write_bit_rate"`
	Secure          string `json:"secure"`
	Hostname        string `json:"hostname"`
	PresenceId      string `json:"presence_id"`
	PresenceData    string `json:"presence_data"`
	Accountcode     string `json:"accountcode"`
	Callstate       string `json:"callstate"`
	CalleeName      string `json:"callee_name"`
	CalleeNum       string `json:"callee_num"`
	CalleeDirection string `json:"callee_direction"`
	CallUuid        string `json:"call_uuid"`
	SentCalleeName  string `json:"sent_callee_name"`
	SentCalleeNum   string `json:"sent_callee_num"`
	InitialCidName  string `json:"initial_cid_name"`
	InitialCidNum   string `json:"initial_cid_num"`
	InitialIpAddr   string `json:"initial_ip_addr"`
	InitialDest     string `json:"initial_dest"`
	InitialDialplan string `json:"initial_dialplan"`
	InitialContext  string `json:"initial_context"`
}

func NewChannelsCache() *Channels {
	return &Channels{
		byUuid: make(map[string]*Channel),
	}
}

func (c *Channels) Set(value *Channel) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byUuid[value.Uuid] = value
}

func (c *Channels) GetByUuid(key string) *Channel {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val := c.byUuid[key]
	return val
}

func (c *Channels) Remove(key *Channel) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.byUuid, key.Uuid)
}

func (c *Channels) GetLength() (int, int) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var answered int
	for _, v := range c.byUuid {
		if v.Callstate == "ACTIVE" {
			answered++
		}
	}

	return len(c.byUuid), answered
}

type WebMetaData struct {
	WssUris     []string `json:"wss_uris"`
	VertoWsUris []string `json:"verto_ws_uris"`
}

func NewWebMetaData() *WebMetaData {
	return &WebMetaData{
		WssUris:     []string{},
		VertoWsUris: []string{},
	}
}

func (w *WebMetaData) GetWssUris() []string {
	return w.WssUris
}

func (w *WebMetaData) SetWssUris(data []string) {
	w.WssUris = data
}

func (w *WebMetaData) GetVertoWsUris() []string {
	return w.VertoWsUris
}

func (w *WebMetaData) SetVertoWsUris(data []string) {
	w.VertoWsUris = data
}

type WebSettings struct {
	mx     sync.RWMutex
	byName map[string]string
}

func (w *WebSettings) Set(key, value string) {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.byName[key] = value
}

func (w *WebSettings) Get(key string) string {
	w.mx.RLock()
	defer w.mx.RUnlock()
	return w.byName[key]
}

func NewWebSettings() *WebSettings {
	return &WebSettings{
		byName: map[string]string{},
	}
}

type LogType struct {
	Created     time.Time `json:"created"`
	LogFile     string    `json:"log_file"`
	LogFunc     string    `json:"log_func"`
	LogLine     int       `json:"log_line"`
	LogLevel    int       `json:"log_level"`
	TextChannel int       `json:"text_channel"`
	UserData    string    `json:"user_data"`
	Body        string    `json:"body"`
}

type FsInstance struct {
	Id          int64  `json:"id"`
	Name        string `json:"name,omitempty"`
	Host        string `json:"host,omitempty"`
	Port        int    `json:"port,omitempty"`
	Auth        string `json:"auth,omitempty"`
	Token       string `json:"token,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

func (w *FsInstance) GetTableName() string {
	return "fs_instances"
}

type FsInstances struct {
	mx     sync.RWMutex
	byId   map[int64]*FsInstance
	byName map[string]*FsInstance
}

func NewFsInstances() *FsInstances {
	return &FsInstances{
		byName: make(map[string]*FsInstance),
		byId:   make(map[int64]*FsInstance),
	}
}

func (f *FsInstances) GetByName(key string) *FsInstance {
	f.mx.RLock()
	defer f.mx.RUnlock()
	val := f.byName[key]
	return val
}

func (f *FsInstances) GetById(key int64) *FsInstance {
	f.mx.RLock()
	defer f.mx.RUnlock()
	return f.byId[key]
}

func (f *FsInstances) Set(value *FsInstance) {
	f.mx.Lock()
	defer f.mx.Unlock()
	f.byName[value.Name] = value
	f.byId[value.Id] = value
}

func (f *FsInstances) GetList() map[int64]*FsInstance {
	f.mx.RLock()
	defer f.mx.RUnlock()
	list := make(map[int64]*FsInstance)
	for _, val := range f.byId {
		list[val.Id] = val
	}
	return list
}

func (f *FsInstances) Rename(oldName, newName string) {
	f.mx.Lock()
	defer f.mx.Unlock()
	if f.byName[oldName] == nil {
		return
	}
	f.byName[newName] = f.byName[oldName]
	f.byName[newName].Name = newName
	delete(f.byName, oldName)
}

func (f *FsInstances) Remove(key *FsInstance) {
	f.mx.Lock()
	defer f.mx.Unlock()
	delete(f.byName, key.Name)
	delete(f.byId, key.Id)
}

func GetWebUserGroups() map[int]WebUserGroup {
	groups := map[int]WebUserGroup{}
	groups[GetBlockedId()] = WebUserGroup{Id: GetBlockedId(), Name: "Blocked"}
	groups[GetAdminId()] = WebUserGroup{Id: GetAdminId(), Name: "Admin"}
	groups[GetManagerId()] = WebUserGroup{Id: GetManagerId(), Name: "Manager"}
	groups[GetUserId()] = WebUserGroup{Id: GetUserId(), Name: "User"}

	return groups
}

func GetWebUserGroup(id int) WebUserGroup {
	group, ok := GetWebUserGroups()[id]
	if !ok {
		group = GetWebUserGroups()[GetBlockedId()]
	}

	return group
}

func GetWebUserBlockedGroup() WebUserGroup {
	return GetWebUserGroup(GetBlockedId())
}

func GetWebUserAdminGroup() WebUserGroup {
	return GetWebUserGroup(GetAdminId())
}

func GetWebUserManagerGroup() WebUserGroup {
	return GetWebUserGroup(GetManagerId())
}

func GetWebUserUserGroup() WebUserGroup {
	return GetWebUserGroup(GetUserId())
}

func GetBlockedId() int {
	return 0
}

func GetAdminId() int {
	return 1
}

func GetManagerId() int {
	return 2
}

func GetUserId() int {
	return 3
}

func (g *WebUserGroup) ValidateGroupAccess(groupList []int) bool {
	for _, v := range groupList {
		if v != g.Id {
			continue
		}
		return true
	}

	return false
}

// web new staff
type WebDirectoryUsersTemplate struct {
	Id            int64   `xml:"-" json:"id" customsql:"pkey:id"`
	Name          string  `xml:"-" json:"name,omitempty" customsql:"name;unique"`
	Cache         int     `xml:"-" json:"cache,omitempty" customsql:"cache"`
	Cidr          string  `xml:"-" json:"cidr,omitempty" customsql:"cidr"`
	NumberAlias   string  `xml:"-" json:"number_alias,omitempty" customsql:"number_alias"`
	Enabled       bool    `xml:"-" json:"enabled,omitempty" customsql:"enabled"`
	DirectoryName string  `xml:"-" json:"directory_name,omitempty" customsql:"directory_name"`
	Domain        *Domain `xml:"-" json:"domain,omitempty" customsql:"fkey:domain_id"`
}

func (w *WebDirectoryUsersTemplate) GetTableName() string {
	return "web_directory_users_templates"
}

type WebDirectoryUsersTemplateParameter struct {
	Id          int64                      `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled     bool                       `xml:"-" json:"enabled,omitempty" customsql:"enabled"`
	Name        string                     `xml:"-" json:"name,omitempty" customsql:"param_name;unique"`
	Value       string                     `xml:"-" json:"value,omitempty" customsql:"param_value"`
	Description string                     `xml:"-" json:"description,omitempty" customsql:"param_description"`
	Placeholder string                     `xml:"-" json:"placeholder,omitempty" customsql:"param_placeholder"`
	Editable    bool                       `xml:"-" json:"editable,omitempty" customsql:"param_editable"`
	Show        bool                       `xml:"-" json:"show,omitempty" customsql:"param_show"`
	Required    bool                       `xml:"-" json:"required,omitempty" customsql:"required"`
	Parent      *WebDirectoryUsersTemplate `xml:"-" json:"parent,omitempty" customsql:"fkey:parent_id;unique"`
}

func (w *WebDirectoryUsersTemplateParameter) GetTableName() string {
	return "web_directory_users_template_parameters"
}

type WebDirectoryUsersTemplateVariable struct {
	Id          int64                      `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled     bool                       `xml:"-" json:"enabled" customsql:"enabled"`
	Name        string                     `xml:"-" json:"name,omitempty" customsql:"var_name;unique"`
	Value       string                     `xml:"-" json:"value,omitempty" customsql:"var_value"`
	Description string                     `xml:"-" json:"description,omitempty" customsql:"var_description"`
	Placeholder string                     `xml:"-" json:"placeholder,omitempty" customsql:"var_placeholder"`
	Editable    bool                       `xml:"-" json:"editable,omitempty" customsql:"var_editable"`
	Show        bool                       `xml:"-" json:"show,omitempty" customsql:"var_show"`
	Required    bool                       `xml:"-" json:"required,omitempty" customsql:"required"`
	Parent      *WebDirectoryUsersTemplate `xml:"-" json:"parent" customsql:"fkey:parent_id;unique"`
}

func (w *WebDirectoryUsersTemplateVariable) GetTableName() string {
	return "web_directory_users_template_variables"
}

func (d *Domain) GetTableName() string {
	return "directory_domains"
}
