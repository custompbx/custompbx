package mainStruct

import (
	"encoding/xml"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// will add all the fields og any configs
type Configuration struct {
	Name              string          `xml:"name,attr"`
	Description       string          `xml:"description,attr"`
	XMLLists          *[]List         `xml:"network-lists>list,omitempty"`
	XMLSettings       *[]Param        `xml:"settings>param,omitempty"`
	XMLGlobalSettings *[]Param        `xml:"global_settings>param,omitempty"`
	XMLQueues         *[]Queue        `xml:"queues>queue,omitempty"`
	XMLProfiles       *[]interface{}  `xml:"profiles>profile,omitempty"`
	XMLSchema         *[]Field        `xml:"schema>field,omitempty"`
	XMLTables         *[]Table        `xml:"tables>table,omitempty"`
	XMLCliKeybindings *[]Param        `xml:"cli-keybindings>key,omitempty"`
	XMLDefaultPtimes  *[]DefaultPtime `xml:"default-ptimes>codec,omitempty"`
	AnyXML            interface{}
}

func NewParams() *Params {
	return &Params{
		byName: make(map[string]*Param),
		byId:   make(map[int64]*Param),
	}
}

func NewLcrProfiles() *LcrProfiles {
	return &LcrProfiles{
		byName: make(map[string]*LcrProfile),
		byId:   make(map[int64]*LcrProfile),
	}
}

func NewLcrProfileParams() *LcrProfileParams {
	return &LcrProfileParams{
		byName: make(map[string]*LcrProfileParam),
		byId:   make(map[int64]*LcrProfileParam),
	}
}

func NewSofiaGatewayParams() *SofiaGatewayParams {
	return &SofiaGatewayParams{
		byName: make(map[string]*SofiaGatewayParam),
		byId:   make(map[int64]*SofiaGatewayParam),
	}
}

func NewSofiaGatewayVars() *SofiaGatewayVars {
	return &SofiaGatewayVars{
		byName: make(map[string]*SofiaGatewayVariable),
		byId:   make(map[int64]*SofiaGatewayVariable),
	}
}

func NewSofiaProfileParams() *SofiaProfileParams {
	return &SofiaProfileParams{
		byName: make(map[string]*SofiaProfileParam),
		byId:   make(map[int64]*SofiaProfileParam),
	}
}

func NewVertoProfiles() *VertoProfiles {
	return &VertoProfiles{
		byName: make(map[string]*VertoProfile),
		byId:   make(map[int64]*VertoProfile),
	}
}

func NewVertoProfileParams() *VertoProfileParams {
	return &VertoProfileParams{
		byName: make(map[string]*VertoProfileParam),
		byId:   make(map[int64]*VertoProfileParam),
	}
}

func NewQueueParams() *QueueParams {
	return &QueueParams{
		byName: make(map[string]*QueueParam),
		byId:   make(map[int64]*QueueParam),
	}
}

func (v *Vars) XMLItems() []Variable {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var variable []Variable
	for _, val := range v.byName {
		if !val.Enabled {
			continue
		}
		variable = append(variable, *val)
	}
	return variable
}

func (p *Params) XMLItems() []Param {
	p.mx.RLock()
	defer p.mx.RUnlock()
	var param []Param
	for _, v := range p.byName {
		if !v.Enabled {
			continue
		}
		param = append(param, *v)
	}
	return param
}

type Params struct {
	mx     sync.RWMutex
	byName map[string]*Param
	byId   map[int64]*Param
}

type Vars struct {
	mx     sync.RWMutex
	byName map[string]*Variable
	byId   map[int64]*Variable
}

type SofiaGatewayParams struct {
	mx     sync.RWMutex
	byName map[string]*SofiaGatewayParam
	byId   map[int64]*SofiaGatewayParam
}

type SofiaGatewayVars struct {
	mx     sync.RWMutex
	byName map[string]*SofiaGatewayVariable
	byId   map[int64]*SofiaGatewayVariable
}

type SofiaProfileParams struct {
	mx     sync.RWMutex
	byName map[string]*SofiaProfileParam
	byId   map[int64]*SofiaProfileParam
}

type Tiers struct {
	mx     sync.RWMutex
	byName map[string]*Tier
	byId   map[int64]*Tier
}

type Members struct {
	mx     sync.RWMutex
	byUuid map[string]*Member
}

type Lists struct {
	mx     sync.RWMutex
	byName map[string]*List
	byId   map[int64]*List
}

type Nodes struct {
	mx     sync.RWMutex
	byName map[string]*Node
	byId   map[int64]*Node
}

type Queues struct {
	mx     sync.RWMutex
	byName map[string]*Queue
	byId   map[int64]*Queue
}

type Agents struct {
	mx     sync.RWMutex
	byName map[string]*Agent
	byId   map[int64]*Agent
}

type SofiaProfiles struct {
	mx     sync.RWMutex
	byName map[string]*SofiaProfile
	byId   map[int64]*SofiaProfile
}

type SofiaGateways struct {
	mx     sync.RWMutex
	byName map[string]*SofiaGateway
	byId   map[int64]*SofiaGateway
}

type Aliases struct {
	mx     sync.RWMutex
	byName map[string]*Alias
	byId   map[int64]*Alias
}

type SofiaDomains struct {
	mx     sync.RWMutex
	byName map[string]*SofiaDomain
	byId   map[int64]*SofiaDomain
}

type Alias struct {
	Id      int64         `xml:"-" json:"id"`
	Enabled bool          `xml:"-" json:"enabled"`
	Name    string        `xml:"name,attr" json:"name"`
	Profile *SofiaProfile `xml:"-" json:"-"`
}

type Param struct {
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Name    string `xml:"name,attr" json:"name"`
	Value   string `xml:"value,attr" json:"value"`
}

type Variable struct {
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Name    string `xml:"name,attr" json:"name"`
	Value   string `xml:"value,attr" json:"value"`
}

type SofiaGatewayParam struct {
	Id      int64         `xml:"-" json:"id"`
	Enabled bool          `xml:"-" json:"enabled"`
	Name    string        `xml:"name,attr" json:"name"`
	Value   string        `xml:"value,attr" json:"value"`
	Gateway *SofiaGateway `xml:"-" json:"-"`
}

type SofiaGatewayVariable struct {
	Id        int64         `xml:"-" json:"id"`
	Enabled   bool          `xml:"-" json:"enabled"`
	Name      string        `xml:"name,attr" json:"name"`
	Value     string        `xml:"value,attr" json:"value"`
	Direction string        `xml:"direction,attr" json:"direction"`
	Gateway   *SofiaGateway `xml:"-" json:"-"`
}

type SofiaProfileParam struct {
	Id      int64         `xml:"-" json:"id"`
	Enabled bool          `xml:"-" json:"enabled"`
	Name    string        `xml:"name,attr" json:"name"`
	Value   string        `xml:"value,attr" json:"value"`
	Profile *SofiaProfile `xml:"-" json:"-"`
}

type SofiaGateway struct {
	Id        int64               `xml:"-" json:"id"`
	Enabled   bool                `xml:"-" json:"enabled"`
	Name      string              `xml:"name,attr" json:"name"`
	Params    *SofiaGatewayParams `xml:"-" json:"-"`
	XmlParams []SofiaGatewayParam `xml:"param,omitempty" json:"-"`
	Vars      *SofiaGatewayVars   `xml:"-" json:"-"`
	XmlVars   []SofiaGatewayVars  `xml:"variable,omitempty" json:"-"`
	Profile   *SofiaProfile       `xml:"-" json:"-"`
	Started   bool                `xml:"-" json:"started"`
	State     string              `xml:"-" json:"state"`
}

type SofiaDomain struct {
	Id      int64         `xml:"-" json:"id"`
	Enabled bool          `xml:"-" json:"enabled"`
	Name    string        `xml:"name,attr" json:"name"`
	Alias   bool          `xml:"alias,attr" json:"alias"`
	Parse   bool          `xml:"parse,attr" json:"parse"`
	Profile *SofiaProfile `xml:"-" json:"-" `
}

type VertoProfiles struct {
	mx     sync.RWMutex
	byName map[string]*VertoProfile
	byId   map[int64]*VertoProfile
}

type VertoProfileParams struct {
	mx     sync.RWMutex
	byName map[string]*VertoProfileParam
	byId   map[int64]*VertoProfileParam
}

type Acl struct {
	Id         int64  `xml:"-" json:"id"`
	Enabled    bool   `xml:"-" json:"enabled"`
	Loaded     bool   `xml:"-" json:"loaded"`
	Unloadable bool   `xml:"-" json:"unloadable"`
	Lists      *Lists `xml:"-" json:"-"`
	XMLLists   []List `xml:"network-lists>list" json:"-"`
	Nodes      *Nodes `xml:"-" json:"-"`
}

type List struct {
	Id       int64  `xml:"-" json:"id"`
	Enabled  bool   `xml:"-" json:"enabled"`
	Default  string `xml:"default,attr" json:"default,omitempty"`
	Name     string `xml:"name,attr"  json:"name,omitempty"`
	Nodes    *Nodes `xml:"-" json:"-"`
	XMLNodes []Node `xml:"node" json:"-"`
}

type Node struct {
	Id       int64  `xml:"-" json:"id"`
	Position int64  `xml:"-" json:"position"`
	Cidr     string `xml:"cidr,attr,omitempty"  json:"cidr,omitempty"`
	Domain   string `xml:"domain,attr,omitempty"  json:"domain,omitempty"`
	Name     string `xml:"name,attr,omitempty"  json:"name,omitempty"`
	Type     string `xml:"type,attr,omitempty"  json:"type,omitempty"`
	Enabled  bool   `xml:"-" json:"enabled"`
	List     *List  `xml:"-" json:"-"`
}

type Callcenter struct {
	Id          int64        `xml:"-" json:"id"`
	Enabled     bool         `xml:"-" json:"enabled"`
	Loaded      bool         `xml:"-" json:"loaded"`
	Settings    *Params      `xml:"-" json:"-"`
	QueueParams *QueueParams `xml:"-" json:"-"`
	Queues      *Queues      `xml:"-" json:"-"`
	Agents      *Agents      `xml:"-" json:"-"`
	Tiers       *Tiers       `xml:"-" json:"-"`
	Members     *Members     `xml:"-" json:"-"`
	XMLSettings []Param      `xml:"settings>param,omitempty" json:"-"`
	XmlQueues   []Queue      `xml:"queues>queue,omitempty" json:"-"`
}

type Queue struct {
	Id        int64        `xml:"-" json:"id"`
	Enabled   bool         `xml:"-" json:"enabled"`
	Name      string       `xml:"name,attr" json:"name"`
	Params    *QueueParams `xml:"-"  json:"-"`
	XmlParams []QueueParam `xml:"param,omitempty" json:"-"`
}

type Agent struct {
	Id                int64  `xml:"-" json:"id"`
	Name              string `xml:"-" json:"name"`
	Type              string `xml:"-" json:"type"`
	System            string `xml:"-" json:"system"`
	InstanceId        string `xml:"-" json:"instance_id"`
	Uuid              string `xml:"-" json:"uuid"`
	Contact           string `xml:"-" json:"contact"`
	Status            string `xml:"-" json:"status"`
	State             string `xml:"-" json:"state"`
	MaxNoAnswer       int64  `xml:"-" json:"max_no_answer"`
	WrapUpTime        int64  `xml:"-" json:"wrap_up_time"`
	RejectDelayTime   int64  `xml:"-" json:"reject_delay_time"`
	BusyDelayTime     int64  `xml:"-" json:"busy_delay_time"`
	NoAnswerDelayTime int64  `xml:"-" json:"no_answer_delay_time"`
	LastBridgeStart   int64  `xml:"-" json:"last_bridge_start"`
	LastBridgeEnd     int64  `xml:"-" json:"last_bridge_end"`
	LastOfferedCall   int64  `xml:"-" json:"last_offered_call"`
	LastStatusChange  int64  `xml:"-" json:"last_status_change"`
	NoAnswerCount     int64  `xml:"-" json:"no_answer_count"`
	CallsAnswered     int64  `xml:"-" json:"calls_answered"`
	TalkTime          int64  `xml:"-" json:"talk_time"`
	ReadyTime         int64  `xml:"-" json:"ready_time"`
}

type CCAgentLight struct {
	Id                int64  `xml:"-" json:"id"`
	Name              string `xml:"-" json:"name,omitempty"`
	Type              string `xml:"-" json:"type,omitempty"`
	System            string `xml:"-" json:"system,omitempty"`
	InstanceId        string `xml:"-" json:"instance_id,omitempty"`
	Uuid              string `xml:"-" json:"uuid,omitempty"`
	Contact           string `xml:"-" json:"contact,omitempty"`
	Status            string `xml:"-" json:"status,omitempty"`
	State             string `xml:"-" json:"state,omitempty"`
	MaxNoAnswer       int64  `xml:"-" json:"max_no_answer,omitempty"`
	WrapUpTime        int64  `xml:"-" json:"wrap_up_time,omitempty"`
	RejectDelayTime   int64  `xml:"-" json:"reject_delay_time,omitempty"`
	BusyDelayTime     int64  `xml:"-" json:"busy_delay_time,omitempty"`
	NoAnswerDelayTime int64  `xml:"-" json:"no_answer_delay_time,omitempty"`
	LastBridgeStart   int64  `xml:"-" json:"last_bridge_start,omitempty"`
	LastBridgeEnd     int64  `xml:"-" json:"last_bridge_end,omitempty"`
	LastOfferedCall   int64  `xml:"-" json:"last_offered_call,omitempty"`
	LastStatusChange  int64  `xml:"-" json:"last_status_change,omitempty"`
	NoAnswerCount     int64  `xml:"-" json:"no_answer_count,omitempty"`
	CallsAnswered     int64  `xml:"-" json:"calls_answered,omitempty"`
	TalkTime          int64  `xml:"-" json:"talk_time,omitempty"`
	ReadyTime         int64  `xml:"-" json:"ready_time,omitempty"`
}

type KnownAgent struct {
	Name              string `xml:"-" json:"name"`
	Type              string `xml:"-" json:"type"`
	System            string `xml:"-" json:"system"`
	InstanceId        string `xml:"-" json:"instance_id"`
	Uuid              string `xml:"-" json:"uuid"`
	Contact           string `xml:"-" json:"contact"`
	Status            string `xml:"-" json:"status"`
	State             string `xml:"-" json:"state"`
	MaxNoAnswer       string `xml:"-" json:"max_no_answer"`
	WrapUpTime        string `xml:"-" json:"wrap_up_time"`
	RejectDelayTime   string `xml:"-" json:"reject_delay_time"`
	BusyDelayTime     string `xml:"-" json:"busy_delay_time"`
	NoAnswerDelayTime string `xml:"-" json:"no_answer_delay_time"`
	LastBridgeStart   string `xml:"-" json:"last_bridge_start"`
	LastBridgeEnd     string `xml:"-" json:"last_bridge_end"`
	LastOfferedCall   string `xml:"-" json:"last_offered_call"`
	LastStatusChange  string `xml:"-" json:"last_status_change"`
	NoAnswerCount     string `xml:"-" json:"no_answer_count"`
	CallsAnswered     string `xml:"-" json:"calls_answered"`
	TalkTime          string `xml:"-" json:"talk_time"`
	ReadyTime         string `xml:"-" json:"ready_time"`
}

type Tier struct {
	Id       int64  `xml:"-" json:"id"`
	Agent    string `xml:"-" json:"agent"`
	Queue    string `xml:"-" json:"queue"`
	Level    int64  `xml:"-" json:"level"`
	Position int64  `xml:"-" json:"position"`
	State    string `xml:"-" json:"state"`
}

type CCTierLight struct {
	Id       int64  `xml:"-" json:"id"`
	Agent    string `xml:"-" json:"agent,omitempty"`
	Queue    string `xml:"-" json:"queue,omitempty"`
	Level    int64  `xml:"-" json:"level,omitempty"`
	Position int64  `xml:"-" json:"position,omitempty"`
	State    string `xml:"-" json:"state,omitempty"`
}

type KnownTier struct {
	Agent    string `xml:"-" json:"agent"`
	Queue    string `xml:"-" json:"queue"`
	Level    string `xml:"-" json:"level"`
	Position string `xml:"-" json:"position"`
	State    string `xml:"-" json:"state"`
}

type Member struct {
	Uuid           string `xml:"-" json:"uuid"`
	Queue          string `xml:"-" json:"queue"`
	InstanceId     string `xml:"-" json:"instance_id"`
	SessionUuid    string `xml:"-" json:"session_uuid"`
	CidNumber      string `xml:"-" json:"cid_number"`
	CidName        string `xml:"-" json:"cid_name"`
	SystemEpoch    int64  `xml:"-" json:"system_epoch"`
	JoinedEpoch    int64  `xml:"-" json:"joined_epoch"`
	RejoinedEpoch  int64  `xml:"-" json:"rejoined_epoch"`
	BridgeEpoch    int64  `xml:"-" json:"bridge_epoch"`
	AbandonedEpoch int64  `xml:"-" json:"abandoned_epoch"`
	BaseScore      int64  `xml:"-" json:"base_score"`
	SkillScore     int64  `xml:"-" json:"skill_score"`
	ServingAgent   string `xml:"-" json:"serving_agent"`
	ServingSystem  string `xml:"-" json:"serving_system"`
	State          string `xml:"-" json:"state"`
}

type SofiaProfile struct {
	Id          int64               `xml:"-" json:"id"`
	Enabled     bool                `xml:"-" json:"enabled"`
	Name        string              `xml:"name,attr"  json:"name"`
	Params      *SofiaProfileParams `xml:"-" json:"-"`
	XmlParams   []SofiaProfileParam `xml:"settings>param,omitempty" json:"-"`
	Aliases     *Aliases            `xml:"-" json:"-"`
	XmlAliases  []Alias             `xml:"aliases>alias,omitempty" json:"-"`
	Gateways    *SofiaGateways      `xml:"-" json:"-"`
	XmlGateways []SofiaGateway      `xml:"gateways>gateway,omitempty" json:"-"`
	Domains     *SofiaDomains       `xml:"-" json:"-"`
	XmlDomains  []SofiaDomain       `xml:"domains>domain,omitempty" json:"-"`
	Started     bool                `xml:"-" json:"started"`
	State       string              `xml:"-" json:"state"`
	Uri         string              `xml:"-" json:"uri"`
}

type VertoProfile struct {
	Id        int64               `xml:"-" json:"id"`
	Enabled   bool                `xml:"-" json:"enabled"`
	Name      string              `xml:"name,attr"  json:"name"`
	Params    *VertoProfileParams `xml:"-" json:"-"`
	XmlParams []VertoProfileParam `xml:"param,omitempty" json:"-"`
	Started   bool                `xml:"-" json:"started"`
	State     string              `xml:"-" json:"state"`
	Uri       string              `xml:"-" json:"uri"`
}

type VertoProfileParam struct {
	Id       int64         `xml:"-" json:"id"`
	Enabled  bool          `xml:"-" json:"enabled"`
	Name     string        `xml:"name,attr" json:"name"`
	Value    string        `xml:"value,attr" json:"value"`
	Secure   string        `xml:"secure,attr" json:"secure"`
	Position int64         `xml:"-" json:"position"`
	Profile  *VertoProfile `xml:"-" json:"-"`
}

type EventSocket struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type FormatCdr struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type Httapi struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type HttapiCache struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type Ivr struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type Logfile struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type Modules struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type Sofia struct {
	Id                int64               `xml:"-" json:"id"`
	Enabled           bool                `json:"enabled" xml:"-"`
	Loaded            bool                `xml:"-" json:"loaded"`
	GlobalSettings    *Params             `json:"-" xml:"-"`
	XmlGlobalSettings []Param             `json:"-" xml:"param,omitempty"`
	Profiles          *SofiaProfiles      `json:"-" xml:"-"`
	XmlProfiles       []interface{}       `json:"-" xml:"profiles>profile,omitempty"`
	ProfileParams     *SofiaProfileParams `json:"-" xml:"-"`
	ProfileAliases    *Aliases            `json:"-" xml:"-"`
	ProfileGateways   *SofiaGateways      `json:"-" xml:"-"`
	ProfileDomains    *SofiaDomains       `json:"-" xml:"-"`
	GatewayParams     *SofiaGatewayParams `json:"-" xml:"-"`
	GatewayVars       *SofiaGatewayVars   `json:"-" xml:"-"`
}

type Lcr struct {
	Id            int64             `xml:"-" json:"id"`
	Enabled       bool              `json:"enabled" xml:"-"`
	Loaded        bool              `xml:"-" json:"loaded"`
	Settings      *Params           `json:"-" xml:"-"`
	XmlSettings   []Param           `json:"-" xml:"param,omitempty"`
	Profiles      *LcrProfiles      `json:"-" xml:"-"`
	XmlProfiles   []interface{}     `json:"-" xml:"profiles>profile,omitempty"`
	ProfileParams *LcrProfileParams `json:"-" xml:"-"`
}

type LcrProfiles struct {
	mx     sync.RWMutex
	byName map[string]*LcrProfile
	byId   map[int64]*LcrProfile
}

type LcrProfile struct {
	Id        int64             `xml:"-" json:"id"`
	Enabled   bool              `json:"enabled" xml:"-"`
	Name      string            `json:"name" xml:"name,attr"`
	Params    *LcrProfileParams `json:"-" xml:"-"`
	XmlParams []LcrProfileParam `json:"-" xml:"param,omitempty"`
}

type LcrProfileParams struct {
	mx     sync.RWMutex
	byName map[string]*LcrProfileParam
	byId   map[int64]*LcrProfileParam
}

type LcrProfileParam struct {
	Id      int64       `xml:"-" json:"id"`
	Enabled bool        `xml:"-" json:"enabled"`
	Name    string      `xml:"name,attr" json:"name"`
	Value   string      `xml:"value,attr" json:"value"`
	Profile *LcrProfile `xml:"-" json:"-"`
}

type Switch struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

type Verto struct {
	Id            int64               `xml:"-" json:"id"`
	Enabled       bool                `json:"enabled" xml:"-"`
	Loaded        bool                `xml:"-" json:"loaded"`
	Settings      *Params             `json:"-" xml:"-"`
	XmlSettings   []Param             `json:"-" xml:"param,omitempty"`
	Profiles      *VertoProfiles      `json:"-" xml:"-"`
	XmlProfiles   []interface{}       `json:"-" xml:"profiles>profile,omitempty"`
	ProfileParams *VertoProfileParams `json:"-" xml:"-"`
}

type XmlCurl struct {
	Id      int64 `xml:"-"`
	Enabled bool  `xml:"-"`
}

func (v *Vars) Remove(key *Variable) {
	v.mx.RLock()
	defer v.mx.RUnlock()
	delete(v.byName, key.Name)
	delete(v.byId, key.Id)
}

func (p *Params) Remove(key *Param) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	delete(p.byName, key.Name)
	delete(p.byId, key.Id)
}

func (v *Vars) HasName(key string) bool {
	v.mx.RLock()
	_, ok := v.byName[key]
	v.mx.RUnlock()
	return ok
}

func (p *Params) HasName(key string) bool {
	p.mx.RLock()
	_, ok := p.byName[key]
	p.mx.RUnlock()
	return ok
}

func (v *Vars) Set(value *Variable) {
	v.mx.Lock()
	defer v.mx.Unlock()
	v.byName[value.Name] = value
	v.byId[value.Id] = value
}

func (p *Params) GetById(key int64) *Param {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (s *SofiaGatewayParams) GetById(key int64) *SofiaGatewayParam {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byId[key]
	return val
}

func (s *SofiaGatewayVars) GetById(key int64) *SofiaGatewayVariable {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byId[key]
	return val
}

func (p *Params) GetByName(key string) *Param {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (v *Vars) GetByName(key string) *Variable {
	v.mx.RLock()
	defer v.mx.RUnlock()
	val := v.byName[key]
	return val
}

func (v *Vars) GetById(key int64) *Variable {
	v.mx.RLock()
	defer v.mx.RUnlock()
	val := v.byId[key]
	return val
}

func (p *Params) Set(value *Param) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.byName[value.Name] = value
	p.byId[value.Id] = value
}

func (s *SofiaGatewayParams) Set(value *SofiaGatewayParam) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[value.Name] = value
	s.byId[value.Id] = value
}

func (s *SofiaGatewayVars) Set(value *SofiaGatewayVariable) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[value.Name] = value
	s.byId[value.Id] = value
}

func (s *SofiaProfileParams) Set(value *SofiaProfileParam) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[value.Name] = value
	s.byId[value.Id] = value
}

func (c *Configurations) NewAcl(id int64, enabled bool) {
	if c.Acl != nil {
		return
	}
	c.Acl = &Acl{Id: id, Enabled: enabled, Lists: NewLists(), Nodes: NewNodes(), Unloadable: true}
}

func (c *Configurations) NewCallcenter(id int64, enabled bool) {
	if c.Callcenter != nil {
		return
	}
	c.Callcenter = &Callcenter{Id: id, Enabled: enabled, Settings: NewParams(), Queues: NewQueues(), Agents: NewAgents(), Tiers: NewTiers(), Members: NewMembers(), QueueParams: NewQueueParams()}
}

/*
	func (c *Configurations) NewEventSocket(id int64, enabled bool) {
		if c.EventSocket != nil {
			return
		}
		c.EventSocket = &EventSocket{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewFormatCdr(id int64, enabled bool) {
		if c.FormatCdr != nil {
			return
		}
		c.FormatCdr = &FormatCdr{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewHttapi(id int64, enabled bool) {
		if c.Httapi != nil {
			return
		}
		c.Httapi = &Httapi{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewHttapiCache(id int64, enabled bool) {
		if c.HttapiCache != nil {
			return
		}
		c.HttapiCache = &HttapiCache{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewIvr(id int64, enabled bool) {
		if c.Ivr != nil {
			return
		}
		c.Ivr = &Ivr{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewLogfile(id int64, enabled bool) {
		if c.Logfile != nil {
			return
		}
		c.Logfile = &Logfile{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewModules(id int64, enabled bool) {
		if c.Modules != nil {
			return
		}
		c.Modules = &Modules{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewNibblebill(id int64, enabled bool) {
		if c.Nibblebill != nil {
			return
		}
		c.Nibblebill = &Nibblebill{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}

	func (c *Configurations) NewPocketsphinx(id int64, enabled bool) {
		if c.Pocketsphinx != nil {
			return
		}
		c.Pocketsphinx = &Pocketsphinx{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
	}
*/
func (c *Configurations) NewSofia(id int64, enabled bool) {
	if c.Sofia != nil {
		return
	}
	c.Sofia = &Sofia{
		Id:              id,
		Enabled:         enabled,
		GlobalSettings:  NewParams(),
		Profiles:        NewProfiles(),
		ProfileParams:   NewSofiaProfileParams(),
		ProfileDomains:  NewSofiaDomains(),
		ProfileGateways: NewSofiaGateways(),
		ProfileAliases:  NewAliases(),
		GatewayParams:   NewSofiaGatewayParams(),
		GatewayVars:     NewSofiaGatewayVars(),
	}
}

func (c *Configurations) NewLcr(id int64, enabled bool) {
	if c.Lcr != nil {
		return
	}
	c.Lcr = &Lcr{
		Id:            id,
		Enabled:       enabled,
		Settings:      NewParams(),
		Profiles:      NewLcrProfiles(),
		ProfileParams: NewLcrProfileParams(),
	}
}

func (c *Configurations) NewVerto(id int64, enabled bool) {
	if c.Verto != nil {
		return
	}
	c.Verto = &Verto{
		Id:            id,
		Enabled:       enabled,
		Settings:      NewParams(),
		Profiles:      NewVertoProfiles(),
		ProfileParams: NewVertoProfileParams(),
	}
}

/*
func (c *Configurations) NewSwitch(id int64, enabled bool) {
	if c.Switch != nil {
		return
	}
	c.Switch = &Switch{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
}

func (c *Configurations) NewVoicemail(id int64, enabled bool) {
	if c.Voicemail != nil {
		return
	}
	c.Voicemail = &Voicemail{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
}

func (c *Configurations) NewXmlCurl(id int64, enabled bool) {
	if c.XmlCurl != nil {
		return
	}
	c.XmlCurl = &XmlCurl{Id: id, Enabled: enabled, Name: name, Lists: *NewLists()}
}*/

func NewLists() *Lists {
	return &Lists{
		byName: make(map[string]*List),
		byId:   make(map[int64]*List),
	}
}

func NewNodes() *Nodes {
	return &Nodes{
		byName: make(map[string]*Node),
		byId:   make(map[int64]*Node),
	}
}

func NewQueues() *Queues {
	return &Queues{
		byName: make(map[string]*Queue),
		byId:   make(map[int64]*Queue),
	}
}

func NewTiers() *Tiers {
	return &Tiers{
		byName: make(map[string]*Tier),
		byId:   make(map[int64]*Tier),
	}
}

func NewMembers() *Members {
	return &Members{
		byUuid: make(map[string]*Member),
	}
}

func NewAgents() *Agents {
	return &Agents{
		byName: make(map[string]*Agent),
		byId:   make(map[int64]*Agent),
	}
}

func NewProfiles() *SofiaProfiles {
	return &SofiaProfiles{
		byName: make(map[string]*SofiaProfile),
		byId:   make(map[int64]*SofiaProfile),
	}
}

func NewAliases() *Aliases {
	return &Aliases{
		byName: make(map[string]*Alias),
		byId:   make(map[int64]*Alias),
	}
}

func NewSofiaGateways() *SofiaGateways {
	return &SofiaGateways{
		byName: make(map[string]*SofiaGateway),
		byId:   make(map[int64]*SofiaGateway),
	}
}

func NewSofiaDomains() *SofiaDomains {
	return &SofiaDomains{
		byName: make(map[string]*SofiaDomain),
		byId:   make(map[int64]*SofiaDomain),
	}
}

func (l *Lists) GetByName(key string) *List {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byName[key]
	return val
}

func (n *Nodes) GetByName(key string) *Node {
	n.mx.RLock()
	defer n.mx.RUnlock()
	val := n.byName[key]
	return val
}

func (q *Queues) GetByName(key string) *Queue {
	q.mx.RLock()
	defer q.mx.RUnlock()
	val := q.byName[key]
	return val
}

func (a *Agents) GetByName(key string) *Agent {
	a.mx.RLock()
	defer a.mx.RUnlock()
	val := a.byName[key]
	return val
}

func (t *Tiers) GetByName(key string) *Tier {
	t.mx.RLock()
	defer t.mx.RUnlock()
	val := t.byName[key]
	return val
}

func (l *Lists) GetById(key int64) *List {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (n *Nodes) GetById(key int64) *Node {
	n.mx.RLock()
	defer n.mx.RUnlock()
	val := n.byId[key]
	return val
}

func (q *Queues) GetById(key int64) *Queue {
	q.mx.RLock()
	defer q.mx.RUnlock()
	val := q.byId[key]
	return val
}

func (a *Agents) GetById(key int64) *Agent {
	a.mx.RLock()
	defer a.mx.RUnlock()
	val := a.byId[key]
	return val
}

func (t *Tiers) GetById(key int64) *Tier {
	t.mx.RLock()
	defer t.mx.RUnlock()
	val := t.byId[key]
	return val
}

func (m *Members) GetByUuid(key string) *Member {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val := m.byUuid[key]
	return val
}

func (s *SofiaProfiles) GetById(key int64) *SofiaProfile {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byId[key]
	return val
}

func (s *SofiaGateways) GetById(key int64) *SofiaGateway {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byId[key]
	return val
}

func (s *SofiaDomains) GetById(key int64) *SofiaDomain {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byId[key]
	return val
}

func (a *Aliases) GetById(key int64) *Alias {
	a.mx.RLock()
	defer a.mx.RUnlock()
	val := a.byId[key]
	return val
}

func (s *SofiaProfiles) GetByName(name string) *SofiaProfile {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byName[name]
	return val
}

func (s *SofiaGateways) GetByName(name string) *SofiaGateway {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byName[name]
	return val
}

func (l *Lists) Set(value *List) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (n *Nodes) Set(value *Node) {
	n.mx.Lock()
	defer n.mx.Unlock()
	n.byName[value.Cidr+value.Domain] = value
	n.byId[value.Id] = value
}

func (q *Queues) Set(value *Queue) {
	q.mx.Lock()
	defer q.mx.Unlock()
	q.byName[value.Name] = value
	q.byId[value.Id] = value
}

func (a *Agents) Set(value *Agent) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byName[value.Name] = value
	a.byId[value.Id] = value
}

func (t *Tiers) Set(value *Tier) {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.byName[value.Queue+value.Agent] = value
	t.byId[value.Id] = value
}

func (m *Members) Set(value *Member) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.byUuid[value.Uuid] = value
}

func (s *SofiaProfiles) Set(value *SofiaProfile) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[value.Name] = value
	s.byId[value.Id] = value
}

func (s *SofiaGateways) Set(value *SofiaGateway) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[value.Name] = value
	s.byId[value.Id] = value
}

func (a *Aliases) Set(value *Alias) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byName[value.Name] = value
	a.byId[value.Id] = value
}

func (s *SofiaDomains) Set(value *SofiaDomain) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.byName[value.Name] = value
	s.byId[value.Id] = value
}

func (l *Lists) Names() []string {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var lists []string
	for k := range l.byName {
		lists = append(lists, k)
	}
	return lists
}

func (l *Lists) Props() []*List {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*List
	for _, v := range l.byId {
		items = append(items, v)
	}
	return items
}

func (n *Nodes) Props() []*Node {
	n.mx.RLock()
	defer n.mx.RUnlock()
	var items []*Node
	for _, v := range n.byId {
		items = append(items, v)
	}
	return items
}

func (q *Queues) Names() []string {
	q.mx.RLock()
	defer q.mx.RUnlock()
	var keys []string
	for k := range q.byName {
		keys = append(keys, k)
	}
	return keys
}

func (q *Queues) Props() []*Queue {
	q.mx.RLock()
	defer q.mx.RUnlock()
	var items []*Queue
	for _, v := range q.byId {
		items = append(items, v)
	}
	return items
}

func (s *SofiaProfiles) Props() []*SofiaProfile {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []*SofiaProfile
	for _, v := range s.byId {
		items = append(items, v)
	}
	return items
}

func (s *SofiaGateways) Props() []*SofiaGateway {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []*SofiaGateway
	for _, v := range s.byId {
		items = append(items, v)
	}
	return items
}

func (c *Configurations) XMLAcl() *Configuration {
	if c.Acl == nil || !c.Acl.Enabled {
		return nil
	}
	c.Acl.XMLItems()
	currentConfig := Configuration{Name: ConfAcl, Description: "ACL Config", XMLLists: &c.Acl.XMLLists}
	return &currentConfig
}

func (a *Acl) XMLItems() {
	a.XMLLists = a.Lists.XMLItems()
}

func (l *Lists) XMLItems() []List {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var list []List
	for _, v := range l.byName {
		if !v.Enabled {
			continue
		}
		v.XMLNodes = v.Nodes.XMLItems()
		list = append(list, *v)
	}
	return list
}

func (n *Nodes) XMLItems() []Node {
	n.mx.RLock()
	defer n.mx.RUnlock()
	var node []Node
	for _, v := range n.byName {
		if !v.Enabled {
			continue
		}
		node = append(node, *v)
	}
	sort.SliceStable(node, func(i, j int) bool {
		return node[i].Position < node[j].Position
	})
	return node
}

func (c *Configurations) XMLCallcenter(queueName string) *Configuration {
	if c.Callcenter == nil || !c.Callcenter.Enabled {
		return nil
	}
	c.Callcenter.XMLItems(queueName)
	currentConfig := Configuration{
		Name:        ConfCallcenter,
		Description: "Callcenter Config",
		XMLSettings: &c.Callcenter.XMLSettings,
		XMLQueues:   &c.Callcenter.XmlQueues,
	}
	return &currentConfig
}

func (c *Callcenter) XMLItems(queueName string) {
	c.XMLSettings = c.Settings.XMLItems()
	if queueName == "" {
		c.XmlQueues = c.Queues.XMLItems()
	} else {
		c.XmlQueues = c.Queues.XMLItemsByName(queueName)
	}
}

func (q *Queues) XMLItems() []Queue {
	q.mx.RLock()
	defer q.mx.RUnlock()
	var queue []Queue
	for _, v := range q.byName {
		if !v.Enabled {
			continue
		}
		v.XmlParams = v.Params.XMLItems()
		queue = append(queue, *v)
	}
	return queue
}

func (q *Queues) XMLItemsByName(queueName string) []Queue {
	q.mx.RLock()
	defer q.mx.RUnlock()
	var queues []Queue
	queue, ok := q.byName[queueName]
	if !ok || queue == nil {
		return queues
	}
	if !queue.Enabled {
		return queues
	}
	queue.XmlParams = queue.Params.XMLItems()
	queues = append(queues, *queue)

	return queues
}

// BY ID ONLY!
func (l *Lists) GetList() map[int64]*List {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*List)
	// BY ID ONLY!
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (d *Nodes) GetList() map[int64]*Node {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[int64]*Node)
	for _, val := range d.byId {
		list[val.Id] = val
	}
	return list
}

func (n *Nodes) Remove(key *Node) {
	n.mx.Lock()
	defer n.mx.Unlock()
	delete(n.byName, key.Name)
	delete(n.byId, key.Id)
}

func (l *Lists) Remove(key *List) {
	l.mx.Lock()
	defer l.mx.Unlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (n *Nodes) ClearUp(conf *Configurations) {
	n.mx.Lock()
	defer n.mx.Unlock()
	for _, v := range n.byId {
		list := conf.Acl.Lists.GetById(v.List.Id)
		if list == nil {
			delete(n.byName, v.Name)
			delete(n.byId, v.Id)
		}
	}
}

func (l *Lists) Rename(oldName, newName string) {
	l.mx.Lock()
	defer l.mx.Unlock()
	if l.byName[oldName] == nil {
		return
	}
	l.byName[newName] = l.byName[oldName]
	l.byName[newName].Name = newName
	delete(l.byName, oldName)
}

func (c *Configurations) XMLSofia(profile string) *Configuration {
	if c.Sofia == nil || !c.Sofia.Enabled {
		return nil
	}

	currentConfig := Configuration{Name: ConfSofia, Description: "Sofia Config"}
	xmlProfiles := c.Sofia.XMLItems(profile)
	currentConfig.XMLProfiles = &xmlProfiles

	//if profile == "" {
	xmlGlobalSettings := c.Sofia.GlobalSettings.XMLItems()
	currentConfig.XMLGlobalSettings = &xmlGlobalSettings
	//}

	return &currentConfig
}

func (s *Sofia) XMLItems(profile string) []interface{} {
	if profile != "" {
		prof := s.Profiles.GetByName(profile)
		if prof == nil {
			return []interface{}{}
		}
		return []interface{}{SofiaProfile{
			Name:        prof.Name,
			XmlAliases:  prof.Aliases.XMLItems(),
			XmlDomains:  prof.Domains.XMLItems(),
			XmlParams:   prof.Params.XMLItems(),
			XmlGateways: prof.Gateways.XMLItems(),
		}}
	}

	return s.Profiles.XMLItems() // self.Profiles.XMLProfilesNames()
}

func (s *SofiaProfiles) XMLProfilesNames() []interface{} {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []interface{}
	for _, v := range s.byName {
		items = append(items, &SofiaProfile{Name: v.Name})
	}
	return items
}

func (s *SofiaProfiles) XMLItems() []interface{} {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []interface{}
	for _, v := range s.byName {
		if v.Enabled {
			items = append(items, &SofiaProfile{
				Name:       v.Name,
				XmlParams:  v.Params.XMLItems(),
				XmlAliases: v.Aliases.XMLItems(),
				XmlDomains: v.Domains.XMLItems(),
			})
		}
	}
	return items
}

func (s *SofiaProfileParams) XMLItems() []SofiaProfileParam {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []SofiaProfileParam
	for _, v := range s.byId {
		if v.Enabled {
			items = append(items, *v)
		}
	}
	return items
}

func (a *Aliases) XMLItems() []Alias {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var items []Alias
	for _, v := range a.byId {
		if v.Enabled {
			items = append(items, *v)
		}
	}
	return items
}

func (s *SofiaDomains) XMLItems() []SofiaDomain {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []SofiaDomain
	for _, v := range s.byId {
		if v.Enabled {
			items = append(items, *v)
		}
	}
	return items
}

func (s *SofiaGateways) XMLItems() []SofiaGateway {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []SofiaGateway
	for _, v := range s.byId {
		if v.Enabled {
			if !v.Enabled {
				continue
			}
			items = append(items, SofiaGateway{Name: v.Name, XmlParams: v.Params.XMLItems()})
		}
	}
	return items
}

func (s *SofiaGatewayParams) XMLItems() []SofiaGatewayParam {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var items []SofiaGatewayParam
	for _, v := range s.byId {
		if v.Enabled {
			items = append(items, *v)
		}
	}
	return items
}

func (s *SofiaProfiles) GetList() map[int64]*SofiaProfile {
	s.mx.RLock()
	defer s.mx.RUnlock()
	list := make(map[int64]*SofiaProfile)
	// BY ID ONLY!
	for _, v := range s.byId {
		list[v.Id] = v
	}
	return list
}

func (s *SofiaProfileParams) GetList() map[int64]*SofiaProfileParam {
	s.mx.RLock()
	defer s.mx.RUnlock()
	list := make(map[int64]*SofiaProfileParam)
	// BY ID ONLY!
	for _, v := range s.byId {
		list[v.Id] = v
	}
	return list
}

func (s *SofiaDomains) GetList() map[int64]*SofiaDomain {
	s.mx.RLock()
	defer s.mx.RUnlock()
	list := make(map[int64]*SofiaDomain)
	// BY ID ONLY!
	for _, v := range s.byId {
		list[v.Id] = v
	}
	return list
}

func (a *Aliases) GetList() map[int64]*Alias {
	a.mx.RLock()
	defer a.mx.RUnlock()
	list := make(map[int64]*Alias)
	// BY ID ONLY!
	for _, v := range a.byId {
		list[v.Id] = v
	}
	return list
}

func (p *Params) GetList() map[int64]*Param {
	p.mx.RLock()
	defer p.mx.RUnlock()
	list := make(map[int64]*Param)
	// BY ID ONLY!
	for _, v := range p.byId {
		list[v.Id] = v
	}
	return list
}

func (s *SofiaProfileParams) GetById(key int64) *SofiaProfileParam {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val := s.byId[key]
	return val
}

func (s *SofiaProfileParams) Remove(key *SofiaProfileParam) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, key.Name)
	delete(s.byId, key.Id)
}

func (s *SofiaDomains) Remove(key *SofiaDomain) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, key.Name)
	delete(s.byId, key.Id)
}

func (s *SofiaProfiles) Remove(key *SofiaProfile) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, key.Name)
	delete(s.byId, key.Id)
}

func (s *SofiaGateways) Remove(key *SofiaGateway) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, key.Name)
	delete(s.byId, key.Id)
}

func (a *Aliases) Remove(key *Alias) {
	a.mx.Lock()
	defer a.mx.Unlock()
	delete(a.byName, key.Name)
	delete(a.byId, key.Id)
}

func (s *SofiaGatewayParams) Remove(key *SofiaGatewayParam) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, key.Name)
	delete(s.byId, key.Id)
}

func (s *SofiaGatewayVars) Remove(key *SofiaGatewayVariable) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.byName, key.Name)
	delete(s.byId, key.Id)
}

func (s *SofiaGateways) GetList() map[int64]*SofiaGateway {
	s.mx.RLock()
	defer s.mx.RUnlock()
	list := make(map[int64]*SofiaGateway)
	// BY ID ONLY!
	for _, v := range s.byId {
		list[v.Id] = v
	}
	return list
}

func (s *SofiaGateways) GetParentList() map[int64]map[int64]*SofiaGateway {
	s.mx.RLock()
	defer s.mx.RUnlock()
	parentList := make(map[int64]map[int64]*SofiaGateway)
	// BY ID ONLY!
	for _, v := range s.byId {
		if _, ok := parentList[v.Profile.Id]; !ok {
			parentList[v.Profile.Id] = make(map[int64]*SofiaGateway)
		}
		parentList[v.Profile.Id][v.Id] = v
	}
	return parentList
}

func (s *SofiaGatewayParams) GetList() map[int64]*SofiaGatewayParam {
	s.mx.RLock()
	defer s.mx.RUnlock()
	list := make(map[int64]*SofiaGatewayParam)
	// BY ID ONLY!
	for _, v := range s.byId {
		list[v.Id] = v
	}
	return list
}

func (s *SofiaGatewayVars) GetList() map[int64]*SofiaGatewayVariable {
	s.mx.RLock()
	defer s.mx.RUnlock()
	list := make(map[int64]*SofiaGatewayVariable)
	// BY ID ONLY!
	for _, v := range s.byId {
		list[v.Id] = v
	}
	return list
}

func (s *SofiaDomains) GetParentList() map[int64]map[int64]*SofiaDomain {
	s.mx.RLock()
	defer s.mx.RUnlock()
	parentList := make(map[int64]map[int64]*SofiaDomain)
	// BY ID ONLY!
	for _, v := range s.byId {
		if _, ok := parentList[v.Profile.Id]; !ok {
			parentList[v.Profile.Id] = make(map[int64]*SofiaDomain)
		}
		parentList[v.Profile.Id][v.Id] = v
	}
	return parentList
}

type Module interface {
	Reload() string
	Unload() string
	Load() string
	Switch(bool)
	AutoLoad()
	GetId() int64
	SetLoadStatus(bool)
	GetConfig() *Configurations
	GetModuleName() string
	IsNil() bool
}

func (a *Acl) Reload() string {
	return "reloadacl"
}
func (a *Acl) Unload() string {
	return "reloadacl"
}
func (a *Acl) Load() string {
	return "reloadacl"
}
func (a *Acl) Switch(enabled bool) {
	a.Enabled = enabled
}
func (a *Acl) AutoLoad() {

}
func (a *Acl) GetId() int64 {
	return a.Id
}
func (a *Acl) SetLoadStatus(status bool) {
	a.Loaded = status
}
func (a *Acl) GetConfig() *Configurations {
	return &Configurations{Acl: a}
}
func (a *Acl) GetModuleName() string {
	return ModAcl
}
func (a *Acl) IsNil() bool {
	return a == nil
}

func (s *Sofia) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Sofia) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Sofia) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Sofia) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Sofia) AutoLoad() {

}
func (s *Sofia) GetId() int64 {
	return s.Id
}
func (s *Sofia) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Sofia) GetConfig() *Configurations {
	return &Configurations{Sofia: s}
}
func (s *Sofia) GetModuleName() string {
	return ModSofia
}
func (s *Sofia) IsNil() bool {
	return s == nil
}

func (c *Callcenter) Reload() string {
	return "reload " + c.GetModuleName()

}
func (c *Callcenter) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *Callcenter) Load() string {
	return "load " + c.GetModuleName()
}
func (c *Callcenter) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *Callcenter) AutoLoad() {

}
func (c *Callcenter) GetId() int64 {
	return c.Id
}
func (c *Callcenter) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *Callcenter) GetConfig() *Configurations {
	return &Configurations{Callcenter: c}
}
func (c *Callcenter) GetModuleName() string {
	return ModCallcenter
}
func (c *Callcenter) IsNil() bool {
	return c == nil
}

func (c *CdrPgCsv) Reload() string {
	return "reload " + c.GetModuleName()
}
func (c *CdrPgCsv) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *CdrPgCsv) Load() string {
	return "load " + c.GetModuleName()
}

func (c *CdrPgCsv) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *CdrPgCsv) AutoLoad() {

}
func (c *CdrPgCsv) GetId() int64 {
	return c.Id
}
func (c *CdrPgCsv) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *CdrPgCsv) GetConfig() *Configurations {
	return &Configurations{CdrPgCsv: c}
}
func (c *CdrPgCsv) GetModuleName() string {
	return ModCdrPgCsv
}
func (c *CdrPgCsv) IsNil() bool {
	return c == nil
}

func (v *Verto) Reload() string {
	return "reload " + v.GetModuleName()

}
func (v *Verto) Unload() string {
	return "unload " + v.GetModuleName()
}
func (v *Verto) Load() string {
	return "load " + v.GetModuleName()
}
func (v *Verto) Switch(enabled bool) {
	v.Enabled = enabled
}
func (v *Verto) AutoLoad() {

}
func (v *Verto) GetId() int64 {
	return v.Id
}
func (v *Verto) SetLoadStatus(status bool) {
	v.Loaded = status
}

func (v *Verto) GetConfig() *Configurations {
	return &Configurations{Verto: v}
}
func (v *Verto) GetModuleName() string {
	return ModVerto
}
func (v *Verto) IsNil() bool {
	return v == nil
}

func (o *OdbcCdr) Reload() string {
	return "reload " + o.GetModuleName()
}
func (o *OdbcCdr) Unload() string {
	return "unload " + o.GetModuleName()
}
func (o *OdbcCdr) Load() string {
	return "load " + o.GetModuleName()
}

func (o *OdbcCdr) Switch(enabled bool) {
	o.Enabled = enabled
}
func (o *OdbcCdr) AutoLoad() {

}
func (o *OdbcCdr) GetId() int64 {
	return o.Id
}
func (o *OdbcCdr) SetLoadStatus(status bool) {
	o.Loaded = status
}
func (o *OdbcCdr) GetConfig() *Configurations {
	return &Configurations{OdbcCdr: o}
}
func (o *OdbcCdr) GetModuleName() string {
	return ModOdbcCdr
}
func (o *OdbcCdr) IsNil() bool {
	return o == nil
}

func (l *Lcr) Reload() string {
	return "reload " + l.GetModuleName()
}
func (l *Lcr) Unload() string {
	return "unload " + l.GetModuleName()
}
func (l *Lcr) Load() string {
	return "load " + l.GetModuleName()
}
func (l *Lcr) Switch(enabled bool) {
	l.Enabled = enabled
}
func (l *Lcr) AutoLoad() {

}
func (l *Lcr) GetId() int64 {
	return l.Id
}
func (l *Lcr) SetLoadStatus(status bool) {
	l.Loaded = status
}
func (l *Lcr) GetConfig() *Configurations {
	return &Configurations{Lcr: l}
}
func (l *Lcr) GetModuleName() string {
	return ModLcr
}
func (l *Lcr) IsNil() bool {
	return l == nil
}

func (s *Shout) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Shout) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Shout) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Shout) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Shout) AutoLoad() {

}
func (s *Shout) GetId() int64 {
	return s.Id
}
func (s *Shout) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Shout) GetConfig() *Configurations {
	return &Configurations{Shout: s}
}
func (s *Shout) GetModuleName() string {
	return ModShout
}
func (s *Shout) IsNil() bool {
	return s == nil
}

func (r *Redis) Reload() string {
	return "reload " + r.GetModuleName()
}
func (r *Redis) Unload() string {
	return "unload " + r.GetModuleName()
}
func (r *Redis) Load() string {
	return "load " + r.GetModuleName()
}
func (r *Redis) Switch(enabled bool) {
	r.Enabled = enabled
}
func (r *Redis) AutoLoad() {

}
func (r *Redis) GetId() int64 {
	return r.Id
}
func (r *Redis) SetLoadStatus(status bool) {
	r.Loaded = status
}
func (r *Redis) GetConfig() *Configurations {
	return &Configurations{Redis: r}
}
func (r *Redis) GetModuleName() string {
	return ModRedis
}
func (r *Redis) IsNil() bool {
	return r == nil
}

func (n *Nibblebill) Reload() string {
	return "reload " + n.GetModuleName()
}
func (n *Nibblebill) Unload() string {
	return "unload " + n.GetModuleName()
}
func (n *Nibblebill) Load() string {
	return "load " + n.GetModuleName()
}
func (n *Nibblebill) Switch(enabled bool) {
	n.Enabled = enabled
}
func (n *Nibblebill) AutoLoad() {

}
func (n *Nibblebill) GetId() int64 {
	return n.Id
}
func (n *Nibblebill) SetLoadStatus(status bool) {
	n.Loaded = status
}
func (n *Nibblebill) GetConfig() *Configurations {
	return &Configurations{Nibblebill: n}
}
func (n *Nibblebill) GetModuleName() string {
	return ModNibblebill
}
func (n *Nibblebill) IsNil() bool {
	return n == nil
}

func (d *Db) Reload() string {
	return "reload " + d.GetModuleName()
}
func (d *Db) Unload() string {
	return "unload " + d.GetModuleName()
}
func (d *Db) Load() string {
	return "load " + d.GetModuleName()
}
func (d *Db) Switch(enabled bool) {
	d.Enabled = enabled
}
func (d *Db) AutoLoad() {

}
func (d *Db) GetId() int64 {
	return d.Id
}
func (d *Db) SetLoadStatus(status bool) {
	d.Loaded = status
}
func (d *Db) GetConfig() *Configurations {
	return &Configurations{Db: d}
}
func (d *Db) GetModuleName() string {
	return ModDb
}
func (d *Db) IsNil() bool {
	return d == nil
}

func (m *Memcache) Reload() string {
	return "reload " + m.GetModuleName()
}
func (m *Memcache) Unload() string {
	return "unload " + m.GetModuleName()
}
func (m *Memcache) Load() string {
	return "load " + m.GetModuleName()
}
func (m *Memcache) Switch(enabled bool) {
	m.Enabled = enabled
}
func (m *Memcache) AutoLoad() {

}
func (m *Memcache) GetId() int64 {
	return m.Id
}
func (m *Memcache) SetLoadStatus(status bool) {
	m.Loaded = status
}
func (m *Memcache) GetConfig() *Configurations {
	return &Configurations{Memcache: m}
}
func (m *Memcache) GetModuleName() string {
	return ModMemcache
}
func (m *Memcache) IsNil() bool {
	return m == nil
}

func (a *Avmd) Reload() string {
	return "reload " + a.GetModuleName()
}
func (a *Avmd) Unload() string {
	return "unload " + a.GetModuleName()
}
func (a *Avmd) Load() string {
	return "load " + a.GetModuleName()
}
func (a *Avmd) Switch(enabled bool) {
	a.Enabled = enabled
}
func (a *Avmd) AutoLoad() {

}
func (a *Avmd) GetId() int64 {
	return a.Id
}
func (a *Avmd) SetLoadStatus(status bool) {
	a.Loaded = status
}
func (a *Avmd) GetConfig() *Configurations {
	return &Configurations{Avmd: a}
}
func (a *Avmd) GetModuleName() string {
	return ModAvmd
}
func (a *Avmd) IsNil() bool {
	return a == nil
}

func (t *TtsCommandline) Reload() string {
	return "reload " + t.GetModuleName()
}
func (t *TtsCommandline) Unload() string {
	return "unload " + t.GetModuleName()
}
func (t *TtsCommandline) Load() string {
	return "load " + t.GetModuleName()
}
func (t *TtsCommandline) Switch(enabled bool) {
	t.Enabled = enabled
}
func (t *TtsCommandline) AutoLoad() {

}
func (t *TtsCommandline) GetId() int64 {
	return t.Id
}
func (t *TtsCommandline) SetLoadStatus(status bool) {
	t.Loaded = status
}
func (t *TtsCommandline) GetConfig() *Configurations {
	return &Configurations{TtsCommandline: t}
}
func (t *TtsCommandline) GetModuleName() string {
	return ModTtsCommandline
}
func (t *TtsCommandline) IsNil() bool {
	return t == nil
}

func (c *CdrMongodb) Reload() string {
	return "reload " + c.GetModuleName()
}
func (c *CdrMongodb) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *CdrMongodb) Load() string {
	return "load " + c.GetModuleName()
}
func (c *CdrMongodb) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *CdrMongodb) AutoLoad() {

}
func (c *CdrMongodb) GetId() int64 {
	return c.Id
}
func (c *CdrMongodb) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *CdrMongodb) GetConfig() *Configurations {
	return &Configurations{CdrMongodb: c}
}
func (c *CdrMongodb) GetModuleName() string {
	return ModCdrMongodb
}
func (c *CdrMongodb) IsNil() bool {
	return c == nil
}

func (h *HttpCache) Reload() string {
	return "reload " + h.GetModuleName()
}
func (h *HttpCache) Unload() string {
	return "unload " + h.GetModuleName()
}
func (h *HttpCache) Load() string {
	return "load " + h.GetModuleName()
}
func (h *HttpCache) Switch(enabled bool) {
	h.Enabled = enabled
}
func (h *HttpCache) AutoLoad() {

}
func (h *HttpCache) GetId() int64 {
	return h.Id
}
func (h *HttpCache) SetLoadStatus(status bool) {
	h.Loaded = status
}
func (h *HttpCache) GetConfig() *Configurations {
	return &Configurations{HttpCache: h}
}
func (h *HttpCache) GetModuleName() string {
	return ModHttpCache
}
func (h *HttpCache) IsNil() bool {
	return h == nil
}

func (o *Opus) Reload() string {
	return "reload " + o.GetModuleName()
}
func (o *Opus) Unload() string {
	return "unload " + o.GetModuleName()
}
func (o *Opus) Load() string {
	return "load " + o.GetModuleName()
}
func (o *Opus) Switch(enabled bool) {
	o.Enabled = enabled
}
func (o *Opus) AutoLoad() {

}
func (o *Opus) GetId() int64 {
	return o.Id
}
func (o *Opus) SetLoadStatus(status bool) {
	o.Loaded = status
}
func (o *Opus) GetConfig() *Configurations {
	return &Configurations{Opus: o}
}
func (o *Opus) GetModuleName() string {
	return ModOpus
}
func (o *Opus) IsNil() bool {
	return o == nil
}

func (p *Python) Reload() string {
	return "reload " + p.GetModuleName()
}
func (p *Python) Unload() string {
	return "unload " + p.GetModuleName()
}
func (p *Python) Load() string {
	return "load " + p.GetModuleName()
}
func (p *Python) Switch(enabled bool) {
	p.Enabled = enabled
}
func (p *Python) AutoLoad() {

}
func (p *Python) GetId() int64 {
	return p.Id
}
func (p *Python) SetLoadStatus(status bool) {
	p.Loaded = status
}
func (p *Python) GetConfig() *Configurations {
	return &Configurations{Python: p}
}
func (p *Python) GetModuleName() string {
	return ModPython
}
func (p *Python) IsNil() bool {
	return p == nil
}

func (a *Alsa) Reload() string {
	return "reload " + a.GetModuleName()
}
func (a *Alsa) Unload() string {
	return "unload " + a.GetModuleName()
}
func (a *Alsa) Load() string {
	return "load " + a.GetModuleName()
}
func (a *Alsa) Switch(enabled bool) {
	a.Enabled = enabled
}
func (a *Alsa) AutoLoad() {

}
func (a *Alsa) GetId() int64 {
	return a.Id
}
func (a *Alsa) SetLoadStatus(status bool) {
	a.Loaded = status
}
func (a *Alsa) GetConfig() *Configurations {
	return &Configurations{Alsa: a}
}
func (a *Alsa) GetModuleName() string {
	return ModAlsa
}
func (a *Alsa) IsNil() bool {
	return a == nil
}

func (a *Amr) Reload() string {
	return "reload " + a.GetModuleName()
}
func (a *Amr) Unload() string {
	return "unload " + a.GetModuleName()
}
func (a *Amr) Load() string {
	return "load " + a.GetModuleName()
}
func (a *Amr) Switch(enabled bool) {
	a.Enabled = enabled
}
func (a *Amr) AutoLoad() {

}
func (a *Amr) GetId() int64 {
	return a.Id
}
func (a *Amr) SetLoadStatus(status bool) {
	a.Loaded = status
}
func (a *Amr) GetConfig() *Configurations {
	return &Configurations{Amr: a}
}
func (a *Amr) GetModuleName() string {
	return ModAmr
}
func (a *Amr) IsNil() bool {
	return a == nil
}

func (a *Amrwb) Reload() string {
	return "reload " + a.GetModuleName()
}
func (a *Amrwb) Unload() string {
	return "unload " + a.GetModuleName()
}
func (a *Amrwb) Load() string {
	return "load " + a.GetModuleName()
}
func (a *Amrwb) Switch(enabled bool) {
	a.Enabled = enabled
}
func (a *Amrwb) AutoLoad() {

}
func (a *Amrwb) GetId() int64 {
	return a.Id
}
func (a *Amrwb) SetLoadStatus(status bool) {
	a.Loaded = status
}
func (a *Amrwb) GetConfig() *Configurations {
	return &Configurations{Amrwb: a}
}
func (a *Amrwb) GetModuleName() string {
	return ModAmrwb
}
func (a *Amrwb) IsNil() bool {
	return a == nil
}

func (c *Cepstral) Reload() string {
	return "reload " + c.GetModuleName()
}
func (c *Cepstral) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *Cepstral) Load() string {
	return "load " + c.GetModuleName()
}
func (c *Cepstral) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *Cepstral) AutoLoad() {

}
func (c *Cepstral) GetId() int64 {
	return c.Id
}
func (c *Cepstral) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *Cepstral) GetConfig() *Configurations {
	return &Configurations{Cepstral: c}
}
func (c *Cepstral) GetModuleName() string {
	return ModCepstral
}
func (c *Cepstral) IsNil() bool {
	return c == nil
}

func (c *Cidlookup) Reload() string {
	return "reload " + c.GetModuleName()
}
func (c *Cidlookup) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *Cidlookup) Load() string {
	return "load " + c.GetModuleName()
}
func (c *Cidlookup) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *Cidlookup) AutoLoad() {

}
func (c *Cidlookup) GetId() int64 {
	return c.Id
}
func (c *Cidlookup) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *Cidlookup) GetConfig() *Configurations {
	return &Configurations{Cidlookup: c}
}
func (c *Cidlookup) GetModuleName() string {
	return ModCidlookup
}
func (c *Cidlookup) IsNil() bool {
	return c == nil
}

func (c *Curl) Reload() string {
	return "reload " + c.GetModuleName()
}
func (c *Curl) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *Curl) Load() string {
	return "load " + c.GetModuleName()
}
func (c *Curl) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *Curl) AutoLoad() {

}
func (c *Curl) GetId() int64 {
	return c.Id
}
func (c *Curl) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *Curl) GetConfig() *Configurations {
	return &Configurations{Curl: c}
}
func (c *Curl) GetModuleName() string {
	return ModCurl
}
func (c *Curl) IsNil() bool {
	return c == nil
}

func (d *DialplanDirectory) Reload() string {
	return "reload " + d.GetModuleName()
}
func (d *DialplanDirectory) Unload() string {
	return "unload " + d.GetModuleName()
}
func (d *DialplanDirectory) Load() string {
	return "load " + d.GetModuleName()
}
func (d *DialplanDirectory) Switch(enabled bool) {
	d.Enabled = enabled
}
func (d *DialplanDirectory) AutoLoad() {

}
func (d *DialplanDirectory) GetId() int64 {
	return d.Id
}
func (d *DialplanDirectory) SetLoadStatus(status bool) {
	d.Loaded = status
}
func (d *DialplanDirectory) GetConfig() *Configurations {
	return &Configurations{DialplanDirectory: d}
}
func (d *DialplanDirectory) GetModuleName() string {
	return ModDialplanDirectory
}
func (d *DialplanDirectory) IsNil() bool {
	return d == nil
}

func (e *Easyroute) Reload() string {
	return "reload " + e.GetModuleName()
}
func (e *Easyroute) Unload() string {
	return "unload " + e.GetModuleName()
}
func (e *Easyroute) Load() string {
	return "load " + e.GetModuleName()
}
func (e *Easyroute) Switch(enabled bool) {
	e.Enabled = enabled
}
func (e *Easyroute) AutoLoad() {

}
func (e *Easyroute) GetId() int64 {
	return e.Id
}
func (e *Easyroute) SetLoadStatus(status bool) {
	e.Loaded = status
}
func (e *Easyroute) GetConfig() *Configurations {
	return &Configurations{Easyroute: e}
}
func (e *Easyroute) GetModuleName() string {
	return ModEasyroute
}
func (e *Easyroute) IsNil() bool {
	return e == nil
}

func (s *ErlangEvent) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *ErlangEvent) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *ErlangEvent) Load() string {
	return "load " + s.GetModuleName()
}
func (s *ErlangEvent) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *ErlangEvent) AutoLoad() {

}
func (s *ErlangEvent) GetId() int64 {
	return s.Id
}
func (s *ErlangEvent) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *ErlangEvent) GetConfig() *Configurations {
	return &Configurations{ErlangEvent: s}
}
func (s *ErlangEvent) GetModuleName() string {
	return ModErlangEvent
}
func (s *ErlangEvent) IsNil() bool {
	return s == nil
}

func (s *EventMulticast) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *EventMulticast) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *EventMulticast) Load() string {
	return "load " + s.GetModuleName()
}
func (s *EventMulticast) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *EventMulticast) AutoLoad() {

}
func (s *EventMulticast) GetId() int64 {
	return s.Id
}
func (s *EventMulticast) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *EventMulticast) GetConfig() *Configurations {
	return &Configurations{EventMulticast: s}
}
func (s *EventMulticast) GetModuleName() string {
	return ModEventMulticast
}
func (s *EventMulticast) IsNil() bool {
	return s == nil
}

func (s *Fax) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Fax) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Fax) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Fax) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Fax) AutoLoad() {

}
func (s *Fax) GetId() int64 {
	return s.Id
}
func (s *Fax) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Fax) GetConfig() *Configurations {
	return &Configurations{Fax: s}
}
func (s *Fax) GetModuleName() string {
	return ModFax
}
func (s *Fax) IsNil() bool {
	return s == nil
}

func (s *Lua) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Lua) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Lua) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Lua) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Lua) AutoLoad() {

}
func (s *Lua) GetId() int64 {
	return s.Id
}
func (s *Lua) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Lua) GetConfig() *Configurations {
	return &Configurations{Lua: s}
}
func (s *Lua) GetModuleName() string {
	return ModLua
}
func (s *Lua) IsNil() bool {
	return s == nil
}

func (s *Mongo) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Mongo) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Mongo) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Mongo) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Mongo) AutoLoad() {

}
func (s *Mongo) GetId() int64 {
	return s.Id
}
func (s *Mongo) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Mongo) GetConfig() *Configurations {
	return &Configurations{Mongo: s}
}
func (s *Mongo) GetModuleName() string {
	return ModMongo
}
func (s *Mongo) IsNil() bool {
	return s == nil
}

func (s *Msrp) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Msrp) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Msrp) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Msrp) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Msrp) AutoLoad() {

}
func (s *Msrp) GetId() int64 {
	return s.Id
}
func (s *Msrp) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Msrp) GetConfig() *Configurations {
	return &Configurations{Msrp: s}
}
func (s *Msrp) GetModuleName() string {
	return ModMsrp
}
func (s *Msrp) IsNil() bool {
	return s == nil
}

func (s *Oreka) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Oreka) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Oreka) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Oreka) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Oreka) AutoLoad() {

}
func (s *Oreka) GetId() int64 {
	return s.Id
}
func (s *Oreka) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Oreka) GetConfig() *Configurations {
	return &Configurations{Oreka: s}
}
func (s *Oreka) GetModuleName() string {
	return ModOreka
}
func (s *Oreka) IsNil() bool {
	return s == nil
}

func (s *Perl) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Perl) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Perl) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Perl) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Perl) AutoLoad() {

}
func (s *Perl) GetId() int64 {
	return s.Id
}
func (s *Perl) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Perl) GetConfig() *Configurations {
	return &Configurations{Perl: s}
}
func (s *Perl) GetModuleName() string {
	return ModPerl
}
func (s *Perl) IsNil() bool {
	return s == nil
}

func (s *Pocketsphinx) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Pocketsphinx) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Pocketsphinx) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Pocketsphinx) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Pocketsphinx) AutoLoad() {

}
func (s *Pocketsphinx) GetId() int64 {
	return s.Id
}
func (s *Pocketsphinx) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Pocketsphinx) GetConfig() *Configurations {
	return &Configurations{Pocketsphinx: s}
}
func (s *Pocketsphinx) GetModuleName() string {
	return ModPocketsphinx
}
func (s *Pocketsphinx) IsNil() bool {
	return s == nil
}

func (s *SangomaCodec) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *SangomaCodec) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *SangomaCodec) Load() string {
	return "load " + s.GetModuleName()
}
func (s *SangomaCodec) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *SangomaCodec) AutoLoad() {

}
func (s *SangomaCodec) GetId() int64 {
	return s.Id
}
func (s *SangomaCodec) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *SangomaCodec) GetConfig() *Configurations {
	return &Configurations{SangomaCodec: s}
}
func (s *SangomaCodec) GetModuleName() string {
	return ModSangomaCodec
}
func (s *SangomaCodec) IsNil() bool {
	return s == nil
}

func (s *Sndfile) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Sndfile) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Sndfile) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Sndfile) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Sndfile) AutoLoad() {

}
func (s *Sndfile) GetId() int64 {
	return s.Id
}
func (s *Sndfile) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Sndfile) GetConfig() *Configurations {
	return &Configurations{Sndfile: s}
}
func (s *Sndfile) GetModuleName() string {
	return ModSndfile
}
func (s *Sndfile) IsNil() bool {
	return s == nil
}

func (s *XmlCdr) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *XmlCdr) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *XmlCdr) Load() string {
	return "load " + s.GetModuleName()
}
func (s *XmlCdr) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *XmlCdr) AutoLoad() {

}
func (s *XmlCdr) GetId() int64 {
	return s.Id
}
func (s *XmlCdr) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *XmlCdr) GetConfig() *Configurations {
	return &Configurations{XmlCdr: s}
}
func (s *XmlCdr) GetModuleName() string {
	return ModXmlCdr
}
func (s *XmlCdr) IsNil() bool {
	return s == nil
}

func (s *XmlRpc) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *XmlRpc) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *XmlRpc) Load() string {
	return "load " + s.GetModuleName()
}
func (s *XmlRpc) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *XmlRpc) AutoLoad() {

}
func (s *XmlRpc) GetId() int64 {
	return s.Id
}
func (s *XmlRpc) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *XmlRpc) GetConfig() *Configurations {
	return &Configurations{XmlRpc: s}
}
func (s *XmlRpc) GetModuleName() string {
	return ModXmlRpc
}
func (s *XmlRpc) IsNil() bool {
	return s == nil
}

func (s *Zeroconf) Reload() string {
	return "reload " + s.GetModuleName()
}
func (s *Zeroconf) Unload() string {
	return "unload " + s.GetModuleName()
}
func (s *Zeroconf) Load() string {
	return "load " + s.GetModuleName()
}
func (s *Zeroconf) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *Zeroconf) AutoLoad() {

}
func (s *Zeroconf) GetId() int64 {
	return s.Id
}
func (s *Zeroconf) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *Zeroconf) GetConfig() *Configurations {
	return &Configurations{Zeroconf: s}
}
func (s *Zeroconf) GetModuleName() string {
	return ModZeroconf
}
func (s *Zeroconf) IsNil() bool {
	return s == nil
}

func (s *PostLoadSwitch) Reload() string {
	return ""
}
func (s *PostLoadSwitch) Unload() string {
	return ""
}
func (s *PostLoadSwitch) Load() string {
	return ""
}
func (s *PostLoadSwitch) Switch(enabled bool) {
	s.Enabled = enabled
}
func (s *PostLoadSwitch) AutoLoad() {

}
func (s *PostLoadSwitch) GetId() int64 {
	return s.Id
}
func (s *PostLoadSwitch) SetLoadStatus(status bool) {
	s.Loaded = status
}
func (s *PostLoadSwitch) GetConfig() *Configurations {
	return &Configurations{PostSwitch: s}
}
func (s *PostLoadSwitch) GetModuleName() string {
	return ModPostLoadSwitch
}
func (s *PostLoadSwitch) IsNil() bool {
	return s == nil
}

func CheckCommand(command string) bool {
	comMap := map[string]bool{
		CommandSofiaProfileStart:   true,
		CommandSofiaProfileStop:    true,
		CommandSofiaProfileRestart: true,
		CommandSofiaProfileRescan:  true,
		CommandSofiaProfileStartgw: true,
		CommandSofiaProfileKillgw:  true,
	}
	_, ok := comMap[command]

	return ok
}

type CdrPgCsv struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `xml:"-" json:"enabled"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `xml:"-" json:"-"`
	XMLSettings []Param `xml:"settings>param,omitempty" json:"-"`
	Schema      *Fields `xml:"-" json:"-"`
	XMLSchema   []Field `xml:"schema>field,omitempty" json:"-"`
}

type OdbcCdr struct {
	Id          int64       `xml:"-" json:"id"`
	Enabled     bool        `xml:"-" json:"enabled"`
	Loaded      bool        `xml:"-" json:"loaded"`
	Settings    *Params     `xml:"-" json:"-"`
	XMLSettings []Param     `xml:"settings>param,omitempty" json:"-"`
	Tables      *Tables     `xml:"-" json:"-"`
	TableFields *ODBCFields `xml:"-" json:"-"`
	XMLTables   []Table     `xml:"tables>table,omitempty" json:"-"`
}

type Tables struct {
	mx     sync.RWMutex
	byName map[string]*Table
	byId   map[int64]*Table
}

type Table struct {
	Id        int64       `xml:"-" json:"id"`
	Enabled   bool        `xml:"-" json:"enabled"`
	Name      string      `xml:"name,attr" json:"name"`
	LogLeg    string      `xml:"log-leg,attr" json:"log_leg"`
	Fields    *ODBCFields `xml:"-" json:"-"`
	XMLFields []ODBCField `xml:"field" json:"-"`
}

type ODBCFields struct {
	mx     sync.RWMutex
	byName map[string]*ODBCField
	byId   map[int64]*ODBCField
}

type ODBCField struct {
	Id          int64  `xml:"-" json:"id"`
	Enabled     bool   `xml:"-" json:"enabled"`
	Name        string `xml:"name,attr" json:"name"`
	ChanVarName string `xml:"chan-var-name,attr" json:"chan_var_name"`
	Table       *Table `xml:"-" json:"-"`
}

type Fields struct {
	mx     sync.RWMutex
	byName map[string]*Field
	byId   map[int64]*Field
}

type Field struct {
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Var     string `xml:"var,attr" json:"var"`
	Column  string `xml:"column,attr,omitempty" json:"column"`
}

func (c *Configurations) XMLCdrPgCsv() *Configuration {
	if c.CdrPgCsv == nil || !c.CdrPgCsv.Enabled {
		return nil
	}
	c.CdrPgCsv.XMLItems()
	currentConfig := Configuration{Name: ConfCdrPgCsv, Description: "Cdr_pg_csv Config", XMLSettings: &c.CdrPgCsv.XMLSettings, XMLSchema: &c.CdrPgCsv.XMLSchema}
	return &currentConfig
}

func (c *Configurations) NewCdrPgCsv(id int64, enabled bool) {
	if c.CdrPgCsv != nil {
		return
	}
	c.CdrPgCsv = &CdrPgCsv{Id: id, Enabled: enabled, Settings: NewParams(), Schema: NewSchema()}
}

func (c *Configurations) XMLOdbcCdr() *Configuration {
	if c.OdbcCdr == nil || !c.OdbcCdr.Enabled {
		return nil
	}
	c.OdbcCdr.XMLItems()
	currentConfig := Configuration{Name: ConfOdbcCdr, Description: "Odbc_cdr Config", XMLSettings: &c.OdbcCdr.XMLSettings, XMLTables: &c.OdbcCdr.XMLTables}
	return &currentConfig
}

func (c *Configurations) NewOdbcCdr(id int64, enabled bool) {
	if c.OdbcCdr != nil {
		return
	}
	c.OdbcCdr = &OdbcCdr{Id: id, Enabled: enabled, Settings: NewParams(), Tables: NewTables(), TableFields: NewOdbcFields()}
}

func (c *CdrPgCsv) XMLItems() {
	c.XMLSettings = c.Settings.XMLItems()
	c.XMLSchema = c.Schema.XMLItems()
}

func (o *OdbcCdr) XMLItems() {
	o.XMLSettings = o.Settings.XMLItems()
	o.XMLTables = o.Tables.XMLItems()
}

func (t *Tables) XMLItems() []Table {
	t.mx.RLock()
	defer t.mx.RUnlock()
	var tables []Table
	for _, v := range t.byName {
		if !v.Enabled {
			continue
		}
		v.XMLFields = v.Fields.XMLItems()
		tables = append(tables, *v)
	}
	return tables
}

func (o *ODBCFields) XMLItems() []ODBCField {
	o.mx.RLock()
	defer o.mx.RUnlock()
	var field []ODBCField
	for _, v := range o.byName {
		if !v.Enabled {
			continue
		}
		field = append(field, *v)
	}
	return field
}

func (f *Fields) XMLItems() []Field {
	f.mx.RLock()
	defer f.mx.RUnlock()
	var field []Field
	for _, v := range f.byName {
		if !v.Enabled {
			continue
		}
		field = append(field, *v)
	}
	return field
}

func NewSchema() *Fields {
	return &Fields{
		byName: make(map[string]*Field),
		byId:   make(map[int64]*Field),
	}
}

func NewTables() *Tables {
	return &Tables{
		byName: make(map[string]*Table),
		byId:   make(map[int64]*Table),
	}
}

func NewOdbcFields() *ODBCFields {
	return &ODBCFields{
		byName: make(map[string]*ODBCField),
		byId:   make(map[int64]*ODBCField),
	}
}

func (f *Fields) Set(value *Field) {
	f.mx.Lock()
	defer f.mx.Unlock()
	f.byName[value.Var] = value
	f.byId[value.Id] = value
}

func (f *Fields) GetList() map[int64]*Field {
	f.mx.RLock()
	defer f.mx.RUnlock()
	list := make(map[int64]*Field)
	// BY ID ONLY!
	for _, v := range f.byId {
		list[v.Id] = v
	}
	return list
}

func (f *Fields) GetById(key int64) *Field {
	f.mx.RLock()
	defer f.mx.RUnlock()
	val, _ := f.byId[key]

	return val
}

func (f *Fields) Remove(key *Field) {
	f.mx.Lock()
	defer f.mx.Unlock()
	delete(f.byName, key.Var)
	delete(f.byId, key.Id)
}

func (t *Tables) Set(value *Table) {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.byName[value.Name] = value
	t.byId[value.Id] = value
}

func (t *Tables) GetList() map[int64]*Table {
	t.mx.RLock()
	defer t.mx.RUnlock()
	list := make(map[int64]*Table)
	// BY ID ONLY!
	for _, v := range t.byId {
		list[v.Id] = v
	}
	return list
}

func (t *Tables) GetById(key int64) *Table {
	t.mx.RLock()
	defer t.mx.RUnlock()
	val, _ := t.byId[key]

	return val
}

func (t *Tables) GetByName(name string) *Table {
	t.mx.RLock()
	defer t.mx.RUnlock()
	val, _ := t.byName[name]

	return val
}

func (t *Tables) Remove(key *Table) {
	t.mx.Lock()
	defer t.mx.Unlock()
	delete(t.byName, key.Name)
	delete(t.byId, key.Id)
}

func (t *Tables) Props() []*Table {
	t.mx.RLock()
	defer t.mx.RUnlock()
	var items []*Table
	for _, v := range t.byId {
		items = append(items, v)
	}
	return items
}

func (o *ODBCFields) Set(value *ODBCField) {
	o.mx.Lock()
	defer o.mx.Unlock()
	o.byName[value.Name] = value
	o.byId[value.Id] = value
}

func (o *ODBCFields) GetList() map[int64]*ODBCField {
	o.mx.RLock()
	defer o.mx.RUnlock()
	list := make(map[int64]*ODBCField)
	// BY ID ONLY!
	for _, v := range o.byId {
		list[v.Id] = v
	}
	return list
}

func (o *ODBCFields) GetParentList() map[int64]map[int64]*ODBCField {
	o.mx.RLock()
	defer o.mx.RUnlock()
	parentList := make(map[int64]map[int64]*ODBCField)
	// BY ID ONLY!
	for _, v := range o.byId {
		if _, ok := parentList[v.Table.Id]; !ok {
			parentList[v.Table.Id] = make(map[int64]*ODBCField)
		}
		parentList[v.Table.Id][v.Id] = v
	}
	return parentList
}

func (o *ODBCFields) GetById(key int64) *ODBCField {
	o.mx.RLock()
	defer o.mx.RUnlock()
	val, _ := o.byId[key]

	return val
}

func (o *ODBCFields) Remove(key *ODBCField) {
	o.mx.Lock()
	defer o.mx.Unlock()
	delete(o.byName, key.Name)
	delete(o.byId, key.Id)
}

func (v *VertoProfiles) Set(value *VertoProfile) {
	v.mx.Lock()
	defer v.mx.Unlock()
	v.byName[value.Name] = value
	v.byId[value.Id] = value
}

func (v *VertoProfileParams) Set(value *VertoProfileParam) {
	v.mx.Lock()
	defer v.mx.Unlock()
	v.byName[value.Name] = value
	v.byId[value.Id] = value
}

func (v *VertoProfiles) Props() []*VertoProfile {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var items []*VertoProfile
	for _, val := range v.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLVerto() *Configuration {
	if c.Verto == nil || !c.Verto.Enabled {
		return nil
	}
	c.Verto.XMLItems()
	currentConfig := Configuration{Name: ConfVerto, Description: "Verto Config", XMLSettings: &c.Verto.XmlSettings, XMLProfiles: &c.Verto.XmlProfiles}
	return &currentConfig
}

func (v *Verto) XMLItems() {
	v.XmlSettings = v.Settings.XMLItems()
	v.XmlProfiles = v.Profiles.XMLItems()
}

func (v *VertoProfiles) XMLItems() []interface{} {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var profile []interface{}
	for _, val := range v.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (v *VertoProfileParams) XMLItems() []VertoProfileParam {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []VertoProfileParam
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	sort.SliceStable(param, func(i, j int) bool {
		return param[i].Position < param[j].Position
	})
	return param
}

func (v *VertoProfiles) GetList() map[int64]*VertoProfile {
	v.mx.RLock()
	defer v.mx.RUnlock()
	list := make(map[int64]*VertoProfile)
	for _, val := range v.byId {
		list[val.Id] = val
	}
	return list
}

func (v *VertoProfiles) GetById(key int64) *VertoProfile {
	v.mx.RLock()
	defer v.mx.RUnlock()
	val := v.byId[key]
	return val
}

func (v *VertoProfiles) GetByName(name string) *VertoProfile {
	v.mx.RLock()
	defer v.mx.RUnlock()
	val := v.byName[name]
	return val
}

func (v *VertoProfileParams) GetList() map[int64]*VertoProfileParam {
	v.mx.RLock()
	defer v.mx.RUnlock()
	list := make(map[int64]*VertoProfileParam)
	// BY ID ONLY!
	for _, val := range v.byId {
		list[val.Id] = val
	}
	return list
}

func (v *VertoProfileParams) GetById(key int64) *VertoProfileParam {
	v.mx.RLock()
	defer v.mx.RUnlock()
	val := v.byId[key]
	return val
}

func (v *VertoProfileParams) GetByName(name string) *VertoProfileParam {
	v.mx.RLock()
	defer v.mx.RUnlock()
	val, _ := v.byName[name]
	return val
}

func (v *VertoProfileParams) Remove(key *VertoProfileParam) {
	v.mx.Lock()
	defer v.mx.Unlock()
	delete(v.byName, key.Name)
	delete(v.byId, key.Id)
}

func (v *VertoProfiles) Remove(key *VertoProfile) {
	v.mx.Lock()
	defer v.mx.Unlock()
	delete(v.byName, key.Name)
	delete(v.byId, key.Id)
}

func (n *VertoProfileParams) Props() []*VertoProfileParam {
	n.mx.RLock()
	defer n.mx.RUnlock()
	var items []*VertoProfileParam
	for _, v := range n.byId {
		items = append(items, v)
	}
	return items
}

func (q *Queues) GetList() map[int64]*Queue {
	q.mx.RLock()
	defer q.mx.RUnlock()
	list := make(map[int64]*Queue)
	// BY ID ONLY!
	for _, v := range q.byId {
		list[v.Id] = v
	}
	return list
}

func (a *Agents) GetList() map[int64]*Agent {
	a.mx.RLock()
	defer a.mx.RUnlock()
	list := make(map[int64]*Agent)
	// BY ID ONLY!
	for _, v := range a.byId {
		list[v.Id] = v
	}
	return list
}

func (t *Tiers) GetList() map[int64]*Tier {
	t.mx.RLock()
	defer t.mx.RUnlock()
	list := make(map[int64]*Tier)
	// BY ID ONLY!
	for _, v := range t.byId {
		list[v.Id] = v
	}
	return list
}

func (q *Queues) Remove(key *Queue) {
	q.mx.Lock()
	defer q.mx.Unlock()
	delete(q.byName, key.Name)
	delete(q.byId, key.Id)
}

func (a *Agents) Remove(key *Agent) {
	a.mx.Lock()
	defer a.mx.Unlock()
	delete(a.byName, key.Name)
	delete(a.byId, key.Id)
}

func (t *Tiers) Remove(key *Tier) {
	t.mx.Lock()
	defer t.mx.Unlock()
	delete(t.byName, key.Queue+key.Agent)
	delete(t.byId, key.Id)
}

func (m *Members) Remove(key *Member) {
	m.mx.Lock()
	defer m.mx.Unlock()
	delete(m.byUuid, key.Uuid)
}

type QueueParams struct {
	mx     sync.RWMutex
	byName map[string]*QueueParam
	byId   map[int64]*QueueParam
}

type QueueParam struct {
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Name    string `xml:"name,attr" json:"name"`
	Value   string `xml:"value,attr" json:"value"`
	Queue   *Queue `xml:"-" json:"-"`
}

func (q *QueueParams) GetList() map[int64]*QueueParam {
	q.mx.RLock()
	defer q.mx.RUnlock()
	list := make(map[int64]*QueueParam)
	// BY ID ONLY!
	for _, v := range q.byId {
		list[v.Id] = v
	}
	return list
}

func (q *QueueParams) Remove(key *QueueParam) {
	q.mx.Lock()
	defer q.mx.Unlock()
	delete(q.byName, key.Name)
	delete(q.byId, key.Id)
}

func (q *QueueParams) GetById(key int64) *QueueParam {
	q.mx.RLock()
	defer q.mx.RUnlock()
	val := q.byId[key]
	return val
}

func (q *QueueParams) GetByName(name string) *QueueParam {
	q.mx.RLock()
	defer q.mx.RUnlock()
	val := q.byName[name]
	return val
}

func (q *QueueParams) Set(value *QueueParam) {
	q.mx.Lock()
	defer q.mx.Unlock()
	q.byName[value.Name] = value
	q.byId[value.Id] = value
}

func (q *QueueParams) XMLItems() []QueueParam {
	q.mx.RLock()
	defer q.mx.RUnlock()
	var param []QueueParam
	for _, v := range q.byName {
		if !v.Enabled {
			continue
		}
		param = append(param, *v)
	}
	return param
}

func (a *Agents) Props() []*Agent {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var items []*Agent
	for _, v := range a.byId {
		items = append(items, v)
	}
	return items
}

func (a *Agents) FilteredProps(limit, offset int, filters []Filter, order Order) ([]*Agent, int) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var items []*Agent
	var total = len(a.byId)
	var structTags = make(map[string]string)
	for _, v := range a.byId {
		if len(filters) > 0 {
			if len(structTags) == 0 {
				structTags = v.TagsList()
			}
			for _, filter := range filters {
				tagged := structTags[filter.Field]
				r := reflect.ValueOf(v)
				f := reflect.Indirect(r).FieldByName(tagged)
				switch f.Type().Name() {
				case "string":
					//log.Println(filter.FieldValue, filter.Operand, f.String(), items)
					switch filter.Operand {
					case "=":
						if filter.FieldValue == f.String() {
							items = append(items, v)
						}
					case ">":
						if filter.FieldValue < f.String() {
							items = append(items, v)
						}
					case "<":
						if filter.FieldValue > f.String() {
							items = append(items, v)
						}
					case "CONTAINS":
						if strings.Contains(f.String(), filter.FieldValue) {
							items = append(items, v)
						}
					}
				case "int64":
					fallthrough
				case "int":
					value, err := strconv.ParseInt(filter.FieldValue, 10, 64)
					if err != nil {
						continue
					}
					//log.Println(value, filter.Operand, f.Int(), items)
					switch filter.Operand {
					case "=":
						if value == f.Int() {
							items = append(items, v)
						}
					case ">":
						if value < f.Int() {
							items = append(items, v)
						}
					case "<":
						if value > f.Int() {
							items = append(items, v)
						}
					case "CONTAINS":
						if strings.Contains(filter.FieldValue, strconv.FormatInt(f.Int(), 10)) {
							items = append(items, v)
						}
					}
				}
			}
		} else {
			items = append(items, v)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		for _, orderItem := range order.Fields {
			rti := reflect.TypeOf(*items[i])
			rtj := reflect.TypeOf(*items[i])
			for t := 0; t < rti.NumField(); t++ {
				fi := rti.Field(t)
				fj := rtj.Field(t)

				fiVal := reflect.ValueOf(fi)
				fjVal := reflect.ValueOf(fj)
				structTag := strings.Split(fi.Tag.Get("json"), ",")[0]
				if structTag == orderItem {
					switch fiVal.Type().Name() {
					case "string":
						if fiVal.String() > fjVal.String() {
							return order.Desc
						}
						if fiVal.String() < fjVal.String() {
							return !order.Desc
						}
					case "int64":
						fallthrough
					case "int":
						if fiVal.Int() > fjVal.Int() {
							return order.Desc
						}
						if fiVal.Int() < fjVal.Int() {
							return !order.Desc
						}
					}
				}
			}
		}
		if order.Desc {
			return items[i].Id > items[j].Id
		}
		return items[i].Id < items[j].Id
	})

	if len(items) >= offset+limit {
		return items[offset : offset+limit], total
	} else if len(items) > offset {
		return items[offset:], total
	}
	return items, total
}

func (a *Agent) Update(key, value string) bool {
	var structItemName = a.GetItemNameByTag(key)
	r := reflect.ValueOf(a)
	f := reflect.Indirect(r).FieldByName(structItemName)
	switch f.Type().Name() {
	case "string":
		f.SetString(value)
		return true
	case "int":
		res, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}
		f.SetInt(res)
		return true
	case "int64":
		res, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}
		f.SetInt(res)
		return true
	case "bool":
		res, err := strconv.ParseBool(value)
		if err != nil {
			return false
		}
		f.SetBool(res)
		return true
	}
	return false
}

func (a *Agent) GetItemNameByTag(key string) string {
	rt := reflect.TypeOf(*a)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		structTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if structTag == key {
			return f.Name
		}
	}
	return ""
}

func GetItemNameByTag(a interface{}, key string) string {
	//rt := reflect.TypeOf(a)
	rt := reflect.ValueOf(a).Elem().Type()

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tags := strings.Split(f.Tag.Get("json"), ",")
		if len(tags) == 0 {
			return ""
		}
		structTag := tags[0]
		if structTag == key {
			return f.Name
		}
	}
	return ""
}

func (a *Agent) TagsList() map[string]string {
	var structTags = make(map[string]string)
	rt := reflect.TypeOf(*a)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		structTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if structTag != "" {
			structTags[structTag] = f.Name
		}
	}
	return structTags
}

func (a *Agents) Rename(oldName, newName string) {
	a.mx.Lock()
	defer a.mx.Unlock()
	if a.byName[oldName] == nil {
		return
	}
	a.byName[newName] = a.byName[oldName]
	a.byName[newName].Name = newName
	delete(a.byName, oldName)
}

func (t *Tiers) FilteredProps(limit, offset int, filters []Filter, order Order) ([]*Tier, int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	var items []*Tier
	var total = len(t.byId)
	var structTags = make(map[string]string)
	for _, v := range t.byId {
		if len(filters) > 0 {
			if len(structTags) == 0 {
				structTags = v.TagsList()
			}
			for _, filter := range filters {
				tagged := structTags[filter.Field]
				r := reflect.ValueOf(v)
				f := reflect.Indirect(r).FieldByName(tagged)
				switch f.Type().Name() {
				case "string":
					//log.Println(filter.FieldValue, filter.Operand, f.String(), items)
					switch filter.Operand {
					case "=":
						if filter.FieldValue == f.String() {
							items = append(items, v)
						}
					case ">":
						if filter.FieldValue < f.String() {
							items = append(items, v)
						}
					case "<":
						if filter.FieldValue > f.String() {
							items = append(items, v)
						}
					case "CONTAINS":
						if strings.Contains(f.String(), filter.FieldValue) {
							items = append(items, v)
						}
					}
				case "int64":
					fallthrough
				case "int":
					value, err := strconv.ParseInt(filter.FieldValue, 10, 64)
					if err != nil {
						continue
					}
					//log.Println(value, filter.Operand, f.Int(), items)
					switch filter.Operand {
					case "=":
						if value == f.Int() {
							items = append(items, v)
						}
					case ">":
						if value < f.Int() {
							items = append(items, v)
						}
					case "<":
						if value > f.Int() {
							items = append(items, v)
						}
					case "CONTAINS":
						if strings.Contains(filter.FieldValue, strconv.FormatInt(f.Int(), 10)) {
							items = append(items, v)
						}
					}
				}
			}
		} else {
			items = append(items, v)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		for _, orderItem := range order.Fields {
			rti := reflect.TypeOf(*items[i])
			rtj := reflect.TypeOf(*items[i])
			for t := 0; t < rti.NumField(); t++ {
				fi := rti.Field(t)
				fj := rtj.Field(t)

				fiVal := reflect.ValueOf(fi)
				fjVal := reflect.ValueOf(fj)
				structTag := strings.Split(fi.Tag.Get("json"), ",")[0]
				if structTag == orderItem {
					switch fiVal.Type().Name() {
					case "string":
						if fiVal.String() > fjVal.String() {
							return order.Desc
						}
						if fiVal.String() < fjVal.String() {
							return !order.Desc
						}
					case "int64":
						fallthrough
					case "int":
						if fiVal.Int() > fjVal.Int() {
							return order.Desc
						}
						if fiVal.Int() < fjVal.Int() {
							return !order.Desc
						}
					}
				}
			}
		}
		if order.Desc {
			return items[i].Id > items[j].Id
		}
		return items[i].Id < items[j].Id
	})

	if len(items) >= offset+limit {
		return items[offset : offset+limit], total
	} else if len(items) > offset {
		return items[offset:], total
	}
	return items, total
}

func (a *Tier) Update(key, value string) bool {
	var structItemName = a.GetItemNameByTag(key)
	r := reflect.ValueOf(a)
	f := reflect.Indirect(r).FieldByName(structItemName)
	switch f.Type().Name() {
	case "string":
		f.SetString(value)
		return true
	case "int":
		res, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}
		f.SetInt(res)
		return true
	case "int64":
		res, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}
		f.SetInt(res)
		return true
	case "bool":
		res, err := strconv.ParseBool(value)
		if err != nil {
			return false
		}
		f.SetBool(res)
		return true
	}
	return false
}

func (a *Tier) GetItemNameByTag(key string) string {
	rt := reflect.TypeOf(*a)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		structTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if structTag == key {
			return f.Name
		}
	}
	return ""
}

func (a *Tier) TagsList() map[string]string {
	var structTags = make(map[string]string)
	rt := reflect.TypeOf(*a)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		structTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if structTag != "" {
			structTags[structTag] = f.Name
		}
	}
	return structTags
}

func (t *Tiers) Rename(oldQueue, newQueue, oldAgent, newAgent string) {
	t.mx.Lock()
	defer t.mx.Unlock()
	if t.byName[oldQueue+oldAgent] == nil {
		return
	}
	t.byName[newQueue+newAgent] = t.byName[oldQueue+oldAgent]
	t.byName[newQueue+newAgent].Queue = newQueue
	t.byName[newQueue+newAgent].Agent = newAgent
	delete(t.byName, oldQueue+oldAgent)
}

func CheckQueueCommand(command string) bool {
	comMap := map[string]bool{
		CommandCallcenterQueueLoad:   true,
		CommandCallcenterQueueReload: true,
		CommandCallcenterQueueUnload: true,
	}
	_, ok := comMap[command]

	return ok
}

func (m *Members) FilteredProps(limit, offset int, filters []Filter, order Order) ([]*Member, int) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	var items []*Member
	var total = len(m.byUuid)
	var structTags = make(map[string]string)
	for _, v := range m.byUuid {
		if len(filters) > 0 {
			if len(structTags) == 0 {
				structTags = v.TagsList()
			}
			for _, filter := range filters {
				tagged := structTags[filter.Field]
				r := reflect.ValueOf(v)
				f := reflect.Indirect(r).FieldByName(tagged)
				var value string
				switch f.Type().Name() {
				case "string":
					value = f.String()
				case "int":
					value = strconv.Itoa(f.Interface().(int))
				case "int64":
					value = strconv.FormatInt(f.Interface().(int64), 10)
				}
				if value != "" || f.Type().Name() == "string" {
					switch filter.Operand {
					case "=":
						if filter.FieldValue == value {
							items = append(items, v)
						}
					case ">":
						if filter.FieldValue >= value {
							items = append(items, v)
						}
					case "<":
						if filter.FieldValue <= value {
							items = append(items, v)
						}
					case "CONTAINS":
						if strings.Contains(value, filter.FieldValue) {
							items = append(items, v)
						}
					}
					break
				}
			}
		} else {
			items = append(items, v)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		for _, orderItem := range order.Fields {
			rti := reflect.TypeOf(*items[i])
			rtj := reflect.TypeOf(*items[i])
			for t := 0; t < rti.NumField(); t++ {
				fi := rti.Field(t)
				fj := rtj.Field(t)
				structTag := strings.Split(fi.Tag.Get("json"), ",")[0]
				if structTag == orderItem {
					if reflect.ValueOf(fi).String() > reflect.ValueOf(fj).String() {
						return order.Desc
					}
					if reflect.ValueOf(fi).String() < reflect.ValueOf(fi).String() {
						return !order.Desc
					}
				}
			}
		}
		if order.Desc {
			return items[i].Uuid > items[j].Uuid
		}
		return items[i].Uuid < items[j].Uuid
	})

	if len(items) >= offset+limit {
		return items[offset : offset+limit], total
	} else if len(items) > offset {
		return items[offset:], total
	}
	return items, total
}

func (a *Member) TagsList() map[string]string {
	var structTags = make(map[string]string)
	rt := reflect.TypeOf(*a)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		structTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if structTag != "" {
			structTags[structTag] = f.Name
		}
	}
	return structTags
}

func (s *Sofia) ClearSofiaProfile() {
	s.ProfileParams.clearUp(s)
	s.ClearSofiaProfileGateways()
	s.ProfileAliases.clearUp(s)
	s.ProfileDomains.clearUp(s)
}

func (s *Sofia) ClearSofiaProfileGateways() {
	s.ProfileGateways.clearUp(s)
	s.GatewayParams.clearUp(s)
	s.GatewayVars.clearUp(s)
}

func (s *SofiaProfileParams) clearUp(conf *Sofia) {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, v := range s.byId {
		list := conf.Profiles.GetById(v.Profile.Id)
		if list == nil {
			delete(s.byName, v.Name)
			delete(s.byId, v.Id)
		}
	}
}

func (a *Aliases) clearUp(conf *Sofia) {
	a.mx.Lock()
	defer a.mx.Unlock()
	for _, v := range a.byId {
		list := conf.Profiles.GetById(v.Profile.Id)
		if list == nil {
			delete(a.byName, v.Name)
			delete(a.byId, v.Id)
		}
	}
}

func (s *SofiaDomains) clearUp(conf *Sofia) {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, v := range s.byId {
		list := conf.Profiles.GetById(v.Profile.Id)
		if list == nil {
			delete(s.byName, v.Name)
			delete(s.byId, v.Id)
		}
	}
}

func (s *SofiaGateways) clearUp(conf *Sofia) {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, v := range s.byId {
		list := conf.Profiles.GetById(v.Profile.Id)
		if list == nil {
			delete(s.byName, v.Name)
			delete(s.byId, v.Id)
		}
	}
}

func (s *SofiaGatewayParams) clearUp(conf *Sofia) {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, v := range s.byId {
		list := conf.Profiles.GetById(v.Gateway.Id)
		if list == nil {
			delete(s.byName, v.Name)
			delete(s.byId, v.Id)
		}
	}
}

func (s *SofiaGatewayVars) clearUp(conf *Sofia) {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, v := range s.byId {
		list := conf.Profiles.GetById(v.Gateway.Id)
		if list == nil {
			delete(s.byName, v.Name)
			delete(s.byId, v.Id)
		}
	}
}

func (v *VertoProfileParams) ClearUp(conf *Verto) {
	v.mx.Lock()
	defer v.mx.Unlock()
	for _, val := range v.byId {
		list := conf.Profiles.GetById(val.Profile.Id)
		if list == nil {
			delete(v.byName, val.Name)
			delete(v.byId, val.Id)
		}
	}
}

func (q *QueueParams) ClearUp(conf *Callcenter) {
	q.mx.Lock()
	defer q.mx.Unlock()
	for _, v := range q.byId {
		list := conf.Queues.GetById(v.Queue.Id)
		if list == nil {
			delete(q.byName, v.Name)
			delete(q.byId, v.Id)
		}
	}
}

func (o *ODBCFields) ClearUp(conf *OdbcCdr) {
	o.mx.Lock()
	defer o.mx.Unlock()
	for _, v := range o.byId {
		list := conf.Tables.GetById(v.Table.Id)
		if list == nil {
			delete(o.byName, v.Name)
			delete(o.byId, v.Id)
		}
	}
}

func (l *LcrProfiles) Set(value *LcrProfile) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *LcrProfileParams) Set(value *LcrProfileParam) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *LcrProfiles) Remove(key *LcrProfile) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *LcrProfileParams) Remove(key *LcrProfileParam) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *LcrProfiles) GetById(key int64) *LcrProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *LcrProfileParams) GetById(key int64) *LcrProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *LcrProfiles) GetByName(key string) *LcrProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *LcrProfileParams) GetByName(key string) *LcrProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *LcrProfileParams) ClearUp(conf *Lcr) {
	l.mx.Lock()
	defer l.mx.Unlock()
	for _, val := range l.byId {
		list := conf.Profiles.GetById(val.Profile.Id)
		if list == nil {
			delete(l.byName, val.Name)
			delete(l.byId, val.Id)
		}
	}
}

func (l *LcrProfiles) GetList() map[int64]*LcrProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*LcrProfile)
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func (l *LcrProfiles) Props() []*LcrProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*LcrProfile
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLLcr() *Configuration {
	if c.Lcr == nil || !c.Lcr.Enabled {
		return nil
	}
	c.Lcr.XMLItems()
	currentConfig := Configuration{Name: ConfLcr, Description: "Lcr Config", XMLSettings: &c.Lcr.XmlSettings, XMLProfiles: &c.Lcr.XmlProfiles}
	return &currentConfig
}

func (l *Lcr) XMLItems() {
	l.XmlSettings = l.Settings.XMLItems()
	l.XmlProfiles = l.Profiles.XMLItems()
}

func (l *LcrProfiles) XMLItems() []interface{} {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var profile []interface{}
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (l *LcrProfileParams) XMLItems() []LcrProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var param []LcrProfileParam
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	return param
}

func (l *LcrProfileParams) GetList() map[int64]*LcrProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*LcrProfileParam)
	// BY ID ONLY!
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

type Shout struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewShout(id int64, enabled bool) {
	if c.Shout != nil {
		return
	}
	c.Shout = &Shout{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLShout() *Configuration {
	if c.Shout == nil || !c.Shout.Enabled {
		return nil
	}
	c.Shout.XMLItems()
	currentConfig := Configuration{Name: ConfShout, Description: "Shout Config", XMLSettings: &c.Shout.XmlSettings}
	return &currentConfig
}

func (s *Shout) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Redis struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewRedis(id int64, enabled bool) {
	if c.Redis != nil {
		return
	}
	c.Redis = &Redis{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLRedis() *Configuration {
	if c.Redis == nil || !c.Redis.Enabled {
		return nil
	}
	c.Redis.XMLItems()
	currentConfig := Configuration{Name: ConfRedis, Description: "Redis Config", XMLSettings: &c.Redis.XmlSettings}
	return &currentConfig
}

func (r *Redis) XMLItems() {
	r.XmlSettings = r.Settings.XMLItems()
}

type Nibblebill struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewNibblebill(id int64, enabled bool) {
	if c.Nibblebill != nil {
		return
	}
	c.Nibblebill = &Nibblebill{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLNibblebill() *Configuration {
	if c.Nibblebill == nil || !c.Nibblebill.Enabled {
		return nil
	}
	c.Nibblebill.XMLItems()
	currentConfig := Configuration{Name: ConfNibblebill, Description: "Nibblebill Config", XMLSettings: &c.Nibblebill.XmlSettings}
	return &currentConfig
}

func (n *Nibblebill) XMLItems() {
	n.XmlSettings = n.Settings.XMLItems()
}

type Db struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewDb(id int64, enabled bool) {
	if c.Db != nil {
		return
	}
	c.Db = &Db{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLDb() *Configuration {
	if c.Db == nil || !c.Db.Enabled {
		return nil
	}
	c.Db.XMLItems()
	currentConfig := Configuration{Name: ConfDb, Description: "Db Config", XMLSettings: &c.Db.XmlSettings}
	return &currentConfig
}

func (d *Db) XMLItems() {
	d.XmlSettings = d.Settings.XMLItems()
}

type Memcache struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewMemcache(id int64, enabled bool) {
	if c.Memcache != nil {
		return
	}
	c.Memcache = &Memcache{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLMemcache() *Configuration {
	if c.Memcache == nil || !c.Memcache.Enabled {
		return nil
	}
	c.Memcache.XMLItems()
	currentConfig := Configuration{Name: ConfMemcache, Description: "Memcache Config", XMLSettings: &c.Memcache.XmlSettings}
	return &currentConfig
}

func (m *Memcache) XMLItems() {
	m.XmlSettings = m.Settings.XMLItems()
}

type Avmd struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewAvmd(id int64, enabled bool) {
	if c.Avmd != nil {
		return
	}
	c.Avmd = &Avmd{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLAvmd() *Configuration {
	if c.Avmd == nil || !c.Avmd.Enabled {
		return nil
	}
	c.Avmd.XMLItems()
	currentConfig := Configuration{Name: ConfAvmd, Description: "Avmd Config", XMLSettings: &c.Avmd.XmlSettings}
	return &currentConfig
}

func (a *Avmd) XMLItems() {
	a.XmlSettings = a.Settings.XMLItems()
}

type TtsCommandline struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewTtsCommandline(id int64, enabled bool) {
	if c.TtsCommandline != nil {
		return
	}
	c.TtsCommandline = &TtsCommandline{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLTtsCommandline() *Configuration {
	if c.TtsCommandline == nil || !c.TtsCommandline.Enabled {
		return nil
	}
	c.TtsCommandline.XMLItems()
	currentConfig := Configuration{Name: ConfTtsCommandline, Description: "TtsCommandline Config", XMLSettings: &c.TtsCommandline.XmlSettings}
	return &currentConfig
}

func (t *TtsCommandline) XMLItems() {
	t.XmlSettings = t.Settings.XMLItems()
}

type CdrMongodb struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewCdrMongodb(id int64, enabled bool) {
	if c.CdrMongodb != nil {
		return
	}
	c.CdrMongodb = &CdrMongodb{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLCdrMongodb() *Configuration {
	if c.CdrMongodb == nil || !c.CdrMongodb.Enabled {
		return nil
	}
	c.CdrMongodb.XMLItems()
	currentConfig := Configuration{Name: ConfCdrMongodb, Description: "CdrMongodb Config", XMLSettings: &c.CdrMongodb.XmlSettings}
	return &currentConfig
}

func (c *CdrMongodb) XMLItems() {
	c.XmlSettings = c.Settings.XMLItems()
}

type HttpCache struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewHttpCache(id int64, enabled bool) {
	if c.HttpCache != nil {
		return
	}
	c.HttpCache = &HttpCache{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLHttpCache() *Configuration {
	if c.HttpCache == nil || !c.HttpCache.Enabled {
		return nil
	}
	c.HttpCache.XMLItems()
	currentConfig := Configuration{Name: ConfHttpCache, Description: "HttpCache Config", XMLSettings: &c.HttpCache.XmlSettings}
	return &currentConfig
}

func (h *HttpCache) XMLItems() {
	h.XmlSettings = h.Settings.XMLItems()
}

type Opus struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewOpus(id int64, enabled bool) {
	if c.Opus != nil {
		return
	}
	c.Opus = &Opus{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLOpus() *Configuration {
	if c.Opus == nil || !c.Opus.Enabled {
		return nil
	}
	c.Opus.XMLItems()
	currentConfig := Configuration{Name: ConfOpus, Description: "Opus Config", XMLSettings: &c.Opus.XmlSettings}
	return &currentConfig
}

func (o *Opus) XMLItems() {
	o.XmlSettings = o.Settings.XMLItems()
}

type Python struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewPython(id int64, enabled bool) {
	if c.Python != nil {
		return
	}
	c.Python = &Python{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLPython() *Configuration {
	if c.Python == nil || !c.Python.Enabled {
		return nil
	}
	c.Python.XMLItems()
	currentConfig := Configuration{Name: ConfPython, Description: "Python Config", XMLSettings: &c.Python.XmlSettings}
	return &currentConfig
}

func (p *Python) XMLItems() {
	p.XmlSettings = p.Settings.XMLItems()
}

type Alsa struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewAlsa(id int64, enabled bool) {
	if c.Alsa != nil {
		return
	}
	c.Alsa = &Alsa{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLAlsa() *Configuration {
	if c.Alsa == nil || !c.Alsa.Enabled {
		return nil
	}
	c.Alsa.XMLItems()
	currentConfig := Configuration{Name: ConfAlsa, Description: "Alsa Config", XMLSettings: &c.Alsa.XmlSettings}
	return &currentConfig
}

func (a *Alsa) XMLItems() {
	a.XmlSettings = a.Settings.XMLItems()
}

type Amr struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewAmr(id int64, enabled bool) {
	if c.Amr != nil {
		return
	}
	c.Amr = &Amr{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLAmr() *Configuration {
	if c.Amr == nil || !c.Amr.Enabled {
		return nil
	}
	c.Amr.XMLItems()
	currentConfig := Configuration{Name: ConfAmr, Description: "Amr Config", XMLSettings: &c.Amr.XmlSettings}
	return &currentConfig
}

func (a *Amr) XMLItems() {
	a.XmlSettings = a.Settings.XMLItems()
}

type Amrwb struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewAmrwb(id int64, enabled bool) {
	if c.Amrwb != nil {
		return
	}
	c.Amrwb = &Amrwb{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLAmrwb() *Configuration {
	if c.Amrwb == nil || !c.Amrwb.Enabled {
		return nil
	}
	c.Amrwb.XMLItems()
	currentConfig := Configuration{Name: ConfAmrwb, Description: "Amrwb Config", XMLSettings: &c.Amrwb.XmlSettings}
	return &currentConfig
}

func (a *Amrwb) XMLItems() {
	a.XmlSettings = a.Settings.XMLItems()
}

type Cepstral struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewCepstral(id int64, enabled bool) {
	if c.Cepstral != nil {
		return
	}
	c.Cepstral = &Cepstral{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLCepstral() *Configuration {
	if c.Cepstral == nil || !c.Cepstral.Enabled {
		return nil
	}
	c.Cepstral.XMLItems()
	currentConfig := Configuration{Name: ConfCepstral, Description: "Cepstral Config", XMLSettings: &c.Cepstral.XmlSettings}
	return &currentConfig
}

func (c *Cepstral) XMLItems() {
	c.XmlSettings = c.Settings.XMLItems()
}

type Cidlookup struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewCidlookup(id int64, enabled bool) {
	if c.Cidlookup != nil {
		return
	}
	c.Cidlookup = &Cidlookup{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLCidlookup() *Configuration {
	if c.Cidlookup == nil || !c.Cidlookup.Enabled {
		return nil
	}
	c.Cidlookup.XMLItems()
	currentConfig := Configuration{Name: ConfCidlookup, Description: "Cidlookup Config", XMLSettings: &c.Cidlookup.XmlSettings}
	return &currentConfig
}

func (c *Cidlookup) XMLItems() {
	c.XmlSettings = c.Settings.XMLItems()
}

type Curl struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewCurl(id int64, enabled bool) {
	if c.Curl != nil {
		return
	}
	c.Curl = &Curl{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLCurl() *Configuration {
	if c.Curl == nil || !c.Curl.Enabled {
		return nil
	}
	c.Curl.XMLItems()
	currentConfig := Configuration{Name: ConfCurl, Description: "Curl Config", XMLSettings: &c.Curl.XmlSettings}
	return &currentConfig
}

func (c *Curl) XMLItems() {
	c.XmlSettings = c.Settings.XMLItems()
}

type DialplanDirectory struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewDialplanDirectory(id int64, enabled bool) {
	if c.DialplanDirectory != nil {
		return
	}
	c.DialplanDirectory = &DialplanDirectory{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLDialplanDirectory() *Configuration {
	if c.DialplanDirectory == nil || !c.DialplanDirectory.Enabled {
		return nil
	}
	c.DialplanDirectory.XMLItems()
	currentConfig := Configuration{Name: ConfDialplanDirectory, Description: "DialplanDirectory Config", XMLSettings: &c.DialplanDirectory.XmlSettings}
	return &currentConfig
}

func (d *DialplanDirectory) XMLItems() {
	d.XmlSettings = d.Settings.XMLItems()
}

type Easyroute struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewEasyroute(id int64, enabled bool) {
	if c.Easyroute != nil {
		return
	}
	c.Easyroute = &Easyroute{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLEasyroute() *Configuration {
	if c.Easyroute == nil || !c.Easyroute.Enabled {
		return nil
	}
	c.Easyroute.XMLItems()
	currentConfig := Configuration{Name: ConfEasyroute, Description: "Easyroute Config", XMLSettings: &c.Easyroute.XmlSettings}
	return &currentConfig
}

func (e *Easyroute) XMLItems() {
	e.XmlSettings = e.Settings.XMLItems()
}

type ErlangEvent struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewErlangEvent(id int64, enabled bool) {
	if c.ErlangEvent != nil {
		return
	}
	c.ErlangEvent = &ErlangEvent{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLErlangEvent() *Configuration {
	if c.ErlangEvent == nil || !c.ErlangEvent.Enabled {
		return nil
	}
	c.ErlangEvent.XMLItems()
	currentConfig := Configuration{Name: ConfErlangEvent, Description: "ErlangEvent Config", XMLSettings: &c.ErlangEvent.XmlSettings}
	return &currentConfig
}

func (s *ErlangEvent) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type EventMulticast struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewEventMulticast(id int64, enabled bool) {
	if c.EventMulticast != nil {
		return
	}
	c.EventMulticast = &EventMulticast{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLEventMulticast() *Configuration {
	if c.EventMulticast == nil || !c.EventMulticast.Enabled {
		return nil
	}
	c.EventMulticast.XMLItems()
	currentConfig := Configuration{Name: ConfEventMulticast, Description: "EventMulticast Config", XMLSettings: &c.EventMulticast.XmlSettings}
	return &currentConfig
}

func (s *EventMulticast) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Fax struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewFax(id int64, enabled bool) {
	if c.Fax != nil {
		return
	}
	c.Fax = &Fax{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLFax() *Configuration {
	if c.Fax == nil || !c.Fax.Enabled {
		return nil
	}
	c.Fax.XMLItems()
	currentConfig := Configuration{Name: ConfFax, Description: "Fax Config", XMLSettings: &c.Fax.XmlSettings}
	return &currentConfig
}

func (s *Fax) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Lua struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewLua(id int64, enabled bool) {
	if c.Lua != nil {
		return
	}
	c.Lua = &Lua{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLLua() *Configuration {
	if c.Lua == nil || !c.Lua.Enabled {
		return nil
	}
	c.Lua.XMLItems()
	currentConfig := Configuration{Name: ConfLua, Description: "Lua Config", XMLSettings: &c.Lua.XmlSettings}
	return &currentConfig
}

func (s *Lua) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Mongo struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewMongo(id int64, enabled bool) {
	if c.Mongo != nil {
		return
	}
	c.Mongo = &Mongo{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLMongo() *Configuration {
	if c.Mongo == nil || !c.Mongo.Enabled {
		return nil
	}
	c.Mongo.XMLItems()
	currentConfig := Configuration{Name: ConfMongo, Description: "Mongo Config", XMLSettings: &c.Mongo.XmlSettings}
	return &currentConfig
}

func (s *Mongo) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Msrp struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewMsrp(id int64, enabled bool) {
	if c.Msrp != nil {
		return
	}
	c.Msrp = &Msrp{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLMsrp() *Configuration {
	if c.Msrp == nil || !c.Msrp.Enabled {
		return nil
	}
	c.Msrp.XMLItems()
	currentConfig := Configuration{Name: ConfMsrp, Description: "Msrp Config", XMLSettings: &c.Msrp.XmlSettings}
	return &currentConfig
}

func (s *Msrp) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Oreka struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewOreka(id int64, enabled bool) {
	if c.Oreka != nil {
		return
	}
	c.Oreka = &Oreka{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLOreka() *Configuration {
	if c.Oreka == nil || !c.Oreka.Enabled {
		return nil
	}
	c.Oreka.XMLItems()
	currentConfig := Configuration{Name: ConfOreka, Description: "Oreka Config", XMLSettings: &c.Oreka.XmlSettings}
	return &currentConfig
}

func (s *Oreka) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Perl struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewPerl(id int64, enabled bool) {
	if c.Perl != nil {
		return
	}
	c.Perl = &Perl{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLPerl() *Configuration {
	if c.Perl == nil || !c.Perl.Enabled {
		return nil
	}
	c.Perl.XMLItems()
	currentConfig := Configuration{Name: ConfPerl, Description: "Perl Config", XMLSettings: &c.Perl.XmlSettings}
	return &currentConfig
}

func (s *Perl) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Pocketsphinx struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewPocketsphinx(id int64, enabled bool) {
	if c.Pocketsphinx != nil {
		return
	}
	c.Pocketsphinx = &Pocketsphinx{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLPocketsphinx() *Configuration {
	if c.Pocketsphinx == nil || !c.Pocketsphinx.Enabled {
		return nil
	}
	c.Pocketsphinx.XMLItems()
	currentConfig := Configuration{Name: ConfPocketsphinx, Description: "Pocketsphinx Config", XMLSettings: &c.Pocketsphinx.XmlSettings}
	return &currentConfig
}

func (s *Pocketsphinx) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type SangomaCodec struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewSangomaCodec(id int64, enabled bool) {
	if c.SangomaCodec != nil {
		return
	}
	c.SangomaCodec = &SangomaCodec{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLSangomaCodec() *Configuration {
	if c.SangomaCodec == nil || !c.SangomaCodec.Enabled {
		return nil
	}
	c.SangomaCodec.XMLItems()
	currentConfig := Configuration{Name: ConfSangomaCodec, Description: "SangomaCodec Config", XMLSettings: &c.SangomaCodec.XmlSettings}
	return &currentConfig
}

func (s *SangomaCodec) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Sndfile struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewSndfile(id int64, enabled bool) {
	if c.Sndfile != nil {
		return
	}
	c.Sndfile = &Sndfile{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLSndfile() *Configuration {
	if c.Sndfile == nil || !c.Sndfile.Enabled {
		return nil
	}
	c.Sndfile.XMLItems()
	currentConfig := Configuration{Name: ConfSndfile, Description: "Sndfile Config", XMLSettings: &c.Sndfile.XmlSettings}
	return &currentConfig
}

func (s *Sndfile) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type XmlCdr struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewXmlCdr(id int64, enabled bool) {
	if c.XmlCdr != nil {
		return
	}
	c.XmlCdr = &XmlCdr{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLXmlCdr() *Configuration {
	if c.XmlCdr == nil || !c.XmlCdr.Enabled {
		return nil
	}
	c.XmlCdr.XMLItems()
	currentConfig := Configuration{Name: ConfXmlCdr, Description: "XmlCdr Config", XMLSettings: &c.XmlCdr.XmlSettings}
	return &currentConfig
}

func (s *XmlCdr) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type XmlRpc struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewXmlRpc(id int64, enabled bool) {
	if c.XmlRpc != nil {
		return
	}
	c.XmlRpc = &XmlRpc{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLXmlRpc() *Configuration {
	if c.XmlRpc == nil || !c.XmlRpc.Enabled {
		return nil
	}
	c.XmlRpc.XMLItems()
	currentConfig := Configuration{Name: ConfXmlRpc, Description: "XmlRpc Config", XMLSettings: &c.XmlRpc.XmlSettings}
	return &currentConfig
}

func (s *XmlRpc) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type Zeroconf struct {
	Id          int64   `xml:"-" json:"id"`
	Enabled     bool    `json:"enabled" xml:"-"`
	Loaded      bool    `xml:"-" json:"loaded"`
	Settings    *Params `json:"-" xml:"-"`
	XmlSettings []Param `json:"-" xml:"param,omitempty"`
}

func (c *Configurations) NewZeroconf(id int64, enabled bool) {
	if c.Zeroconf != nil {
		return
	}
	c.Zeroconf = &Zeroconf{
		Id:       id,
		Enabled:  enabled,
		Settings: NewParams(),
	}
}

func (c *Configurations) XMLZeroconf() *Configuration {
	if c.Zeroconf == nil || !c.Zeroconf.Enabled {
		return nil
	}
	c.Zeroconf.XMLItems()
	currentConfig := Configuration{Name: ConfZeroconf, Description: "Zeroconf Config", XMLSettings: &c.Zeroconf.XmlSettings}
	return &currentConfig
}

func (s *Zeroconf) XMLItems() {
	s.XmlSettings = s.Settings.XMLItems()
}

type PostLoadSwitch struct {
	Id                int64          `xml:"-" json:"id"`
	Enabled           bool           `json:"enabled" xml:"-"`
	Loaded            bool           `xml:"-" json:"loaded"`
	CliKeybindings    *Params        `json:"-" xml:"-"`
	XmlCliKeybindings []Param        `json:"-" xml:"key,omitempty"`
	DefaultPtimes     *DefaultPtimes `json:"-" xml:"-"`
	XmlDefaultPtimes  []DefaultPtime `json:"-" xml:"codec,omitempty"`
	Settings          *Params        `json:"-" xml:"-"`
	XmlSettings       []Param        `json:"-" xml:"param,omitempty"`
	Unloadable        bool           `xml:"-" json:"unloadable"`
}

func (c *Configurations) NewPostSwitch(id int64, enabled bool) {
	if c.PostSwitch != nil {
		return
	}
	c.PostSwitch = &PostLoadSwitch{
		Id:             id,
		Enabled:        enabled,
		CliKeybindings: NewParams(),
		DefaultPtimes:  NewDefaultPtimes(),
		Settings:       NewParams(),
		Unloadable:     true,
	}
}

func (c *Configurations) XMLPostSwitch() *Configuration {
	if c.PostSwitch == nil || !c.PostSwitch.Enabled {
		return nil
	}
	c.PostSwitch.XMLItems()
	currentConfig := Configuration{Name: ConfPostLoadSwitch, Description: "Post load switch Config", XMLCliKeybindings: &c.PostSwitch.XmlCliKeybindings, XMLDefaultPtimes: &c.PostSwitch.XmlDefaultPtimes, XMLSettings: &c.PostSwitch.XmlSettings}
	return &currentConfig
}

func (p *PostLoadSwitch) XMLItems() {
	p.XmlCliKeybindings = p.CliKeybindings.XMLItems()
	p.XmlDefaultPtimes = p.DefaultPtimes.XMLItems()
	p.XmlSettings = p.Settings.XMLItems()
}

type DefaultPtimes struct {
	mx     sync.RWMutex
	byName map[string]*DefaultPtime
	byId   map[int64]*DefaultPtime
}

type DefaultPtime struct {
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Name    string `xml:"name,attr" json:"name"`
	Ptime   string `xml:"ptime,attr" json:"ptime"`
}

func NewDefaultPtimes() *DefaultPtimes {
	return &DefaultPtimes{
		byName: make(map[string]*DefaultPtime),
		byId:   make(map[int64]*DefaultPtime),
	}
}

func (p *DefaultPtimes) Set(value *DefaultPtime) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.byName[value.Name] = value
	p.byId[value.Id] = value
}

func (p *DefaultPtimes) GetList() map[int64]*DefaultPtime {
	p.mx.RLock()
	defer p.mx.RUnlock()
	list := make(map[int64]*DefaultPtime)
	// BY ID ONLY!
	for _, v := range p.byId {
		list[v.Id] = v
	}
	return list
}

func (p *DefaultPtimes) XMLItems() []DefaultPtime {
	p.mx.RLock()
	defer p.mx.RUnlock()
	var param []DefaultPtime
	for _, v := range p.byName {
		if !v.Enabled {
			continue
		}
		param = append(param, *v)
	}
	return param
}

func (p *DefaultPtimes) Remove(key *DefaultPtime) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	delete(p.byName, key.Name)
	delete(p.byId, key.Id)
}

func (p *DefaultPtimes) HasName(key string) bool {
	p.mx.RLock()
	_, ok := p.byName[key]
	p.mx.RUnlock()
	return ok
}

func (p *DefaultPtimes) GetById(key int64) *DefaultPtime {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *DefaultPtimes) GetByName(key string) *DefaultPtime {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

type Distributor struct {
	Id       int64             `xml:"-" json:"id"`
	Enabled  bool              `xml:"-" json:"enabled"`
	Loaded   bool              `xml:"-" json:"loaded"`
	Lists    *DistributorLists `xml:"-" json:"-"`
	XMLLists []DistributorList `xml:"lists>list" json:"-"`
	Nodes    *DistributorNodes `xml:"-" json:"-"`
}

type DistributorList struct {
	Id       int64             `xml:"-" json:"id"`
	Enabled  bool              `xml:"-" json:"enabled"`
	Name     string            `xml:"name,attr"  json:"name,omitempty"`
	Nodes    *DistributorNodes `xml:"-" json:"-"`
	XMLNodes []DistributorNode `xml:"node" json:"-"`
}

type DistributorNode struct {
	Id      int64            `xml:"-" json:"id"`
	Name    string           `xml:"name,attr,omitempty"  json:"name,omitempty"`
	Weight  string           `xml:"weight,attr,omitempty"  json:"weight,omitempty"`
	Enabled bool             `xml:"-" json:"enabled"`
	List    *DistributorList `xml:"-" json:"-"`
}

type DistributorLists struct {
	mx     sync.RWMutex
	byName map[string]*DistributorList
	byId   map[int64]*DistributorList
}

type DistributorNodes struct {
	mx     sync.RWMutex
	byName map[string]*DistributorNode
	byId   map[int64]*DistributorNode
}

func (c *Configurations) NewDistributor(id int64, enabled bool) {
	if c.Distributor != nil {
		return
	}
	c.Distributor = &Distributor{Id: id, Enabled: enabled, Lists: NewDistributorLists(), Nodes: NewDistributorNodes()}
}

func NewDistributorLists() *DistributorLists {
	return &DistributorLists{
		byName: make(map[string]*DistributorList),
		byId:   make(map[int64]*DistributorList),
	}
}

func NewDistributorNodes() *DistributorNodes {
	return &DistributorNodes{
		byName: make(map[string]*DistributorNode),
		byId:   make(map[int64]*DistributorNode),
	}
}

func (c *Configurations) XMLDistributor() *Configuration {
	if c.Distributor == nil || !c.Distributor.Enabled {
		return nil
	}
	c.Distributor.XMLItems()
	currentConfig := Configuration{Name: ConfDistributor, Description: "Distributor Config", AnyXML: &c.Distributor.XMLLists}
	return &currentConfig
}

func (d *Distributor) XMLItems() {
	d.XMLLists = d.Lists.XMLItems()
}

func (l *DistributorLists) XMLItems() []DistributorList {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var list []DistributorList
	for _, v := range l.byName {
		if !v.Enabled {
			continue
		}
		v.XMLNodes = v.Nodes.XMLItems()
		list = append(list, *v)
	}
	return list
}

func (d *DistributorNodes) XMLItems() []DistributorNode {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var node []DistributorNode
	for _, v := range d.byName {
		if v.Enabled {
			node = append(node, *v)
		}
	}
	return node
}

func (d *Distributor) Reload() string {
	return "reload " + d.GetModuleName()
}
func (d *Distributor) Unload() string {
	return "unload " + d.GetModuleName()
}
func (d *Distributor) Load() string {
	return "load " + d.GetModuleName()
}
func (d *Distributor) Switch(enabled bool) {
	d.Enabled = enabled
}
func (d *Distributor) AutoLoad() {

}
func (d *Distributor) GetId() int64 {
	return d.Id
}
func (d *Distributor) SetLoadStatus(status bool) {
	d.Loaded = status
}
func (d *Distributor) GetConfig() *Configurations {
	return &Configurations{Distributor: d}
}
func (d *Distributor) GetModuleName() string {
	return ModDistributor
}
func (d *Distributor) IsNil() bool {
	return d == nil
}

func (l *DistributorLists) GetByName(key string) *DistributorList {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byName[key]
	return val
}

func (l *DistributorLists) GetById(key int64) *DistributorList {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}
func (l *DistributorLists) Set(value *DistributorList) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *DistributorLists) Names() []string {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var lists []string
	for k := range l.byName {
		lists = append(lists, k)
	}
	return lists
}

func (l *DistributorLists) Props() []*DistributorList {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*DistributorList
	for _, v := range l.byId {
		items = append(items, v)
	}
	return items
}

func (l *DistributorLists) GetList() map[int64]*DistributorList {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*DistributorList)
	// BY ID ONLY!
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (l *DistributorLists) Rename(oldName, newName string) {
	l.mx.Lock()
	defer l.mx.Unlock()
	if l.byName[oldName] == nil {
		return
	}
	l.byName[newName] = l.byName[oldName]
	l.byName[newName].Name = newName
	delete(l.byName, oldName)
}

func (l *DistributorLists) Remove(key *DistributorList) {
	l.mx.Lock()
	defer l.mx.Unlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (d *DistributorNodes) GetList() map[int64]*DistributorNode {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[int64]*DistributorNode)
	for _, val := range d.byId {
		list[val.Id] = val
	}
	return list
}

func (d *DistributorNodes) GetById(key int64) *DistributorNode {
	d.mx.RLock()
	defer d.mx.RUnlock()
	val := d.byId[key]
	return val
}

func (d *DistributorNodes) GetByName(key string) *DistributorNode {
	d.mx.RLock()
	defer d.mx.RUnlock()
	val := d.byName[key]
	return val
}

func (d *DistributorNodes) Set(value *DistributorNode) {
	d.mx.Lock()
	defer d.mx.Unlock()
	d.byName[value.Name] = value
	d.byId[value.Id] = value
}

func (d *DistributorNodes) Props() []*DistributorNode {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var items []*DistributorNode
	for _, v := range d.byId {
		items = append(items, v)
	}
	return items
}

func (d *DistributorNodes) Remove(key *DistributorNode) {
	d.mx.Lock()
	defer d.mx.Unlock()
	delete(d.byName, key.Name)
	delete(d.byId, key.Id)
}

func (d *DistributorNodes) ClearUp(conf *Configurations) {
	d.mx.Lock()
	defer d.mx.Unlock()
	for _, v := range d.byId {
		list := conf.Acl.Lists.GetById(v.List.Id)
		if list == nil {
			delete(d.byName, v.Name)
			delete(d.byId, v.Id)
		}
	}
}

func NewDirectoryProfiles() *DirectoryProfiles {
	return &DirectoryProfiles{
		byName: make(map[string]*DirectoryProfile),
		byId:   make(map[int64]*DirectoryProfile),
	}
}

func NewDirectoryProfileParams() *DirectoryProfileParams {
	return &DirectoryProfileParams{
		byName: make(map[string]*DirectoryProfileParam),
		byId:   make(map[int64]*DirectoryProfileParam),
	}
}

type Directory struct {
	Id            int64                   `xml:"-" json:"id"`
	Enabled       bool                    `json:"enabled" xml:"-"`
	Loaded        bool                    `xml:"-" json:"loaded"`
	Settings      *Params                 `json:"-" xml:"-"`
	XmlSettings   []Param                 `json:"-" xml:"param,omitempty"`
	Profiles      *DirectoryProfiles      `json:"-" xml:"-"`
	XmlProfiles   []interface{}           `json:"-" xml:"profiles>profile,omitempty"`
	ProfileParams *DirectoryProfileParams `json:"-" xml:"-"`
}

type DirectoryProfiles struct {
	mx     sync.RWMutex
	byName map[string]*DirectoryProfile
	byId   map[int64]*DirectoryProfile
}

type DirectoryProfile struct {
	Id        int64                   `xml:"-" json:"id"`
	Enabled   bool                    `json:"enabled" xml:"-"`
	Name      string                  `json:"name" xml:"name,attr"`
	Params    *DirectoryProfileParams `json:"-" xml:"-"`
	XmlParams []DirectoryProfileParam `json:"-" xml:"param,omitempty"`
}

type DirectoryProfileParams struct {
	mx     sync.RWMutex
	byName map[string]*DirectoryProfileParam
	byId   map[int64]*DirectoryProfileParam
}

type DirectoryProfileParam struct {
	Id      int64             `xml:"-" json:"id"`
	Enabled bool              `xml:"-" json:"enabled"`
	Name    string            `xml:"name,attr" json:"name"`
	Value   string            `xml:"value,attr" json:"value"`
	Profile *DirectoryProfile `xml:"-" json:"-"`
}

func (c *Configurations) NewDirectory(id int64, enabled bool) {
	if c.Directory != nil {
		return
	}
	c.Directory = &Directory{
		Id:            id,
		Enabled:       enabled,
		Settings:      NewParams(),
		Profiles:      NewDirectoryProfiles(),
		ProfileParams: NewDirectoryProfileParams(),
	}
}

func (l *Directory) Reload() string {
	return "reload " + l.GetModuleName()
}
func (l *Directory) Unload() string {
	return "unload " + l.GetModuleName()
}
func (l *Directory) Load() string {
	return "load " + l.GetModuleName()
}
func (l *Directory) Switch(enabled bool) {
	l.Enabled = enabled
}
func (l *Directory) AutoLoad() {

}
func (l *Directory) GetId() int64 {
	return l.Id
}
func (l *Directory) SetLoadStatus(status bool) {
	l.Loaded = status
}
func (l *Directory) GetConfig() *Configurations {
	return &Configurations{Directory: l}
}
func (l *Directory) GetModuleName() string {
	return ModDirectory
}
func (l *Directory) IsNil() bool {
	return l == nil
}

func (l *DirectoryProfiles) Set(value *DirectoryProfile) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *DirectoryProfileParams) Set(value *DirectoryProfileParam) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *DirectoryProfiles) Remove(key *DirectoryProfile) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *DirectoryProfileParams) Remove(key *DirectoryProfileParam) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *DirectoryProfiles) GetById(key int64) *DirectoryProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *DirectoryProfileParams) GetById(key int64) *DirectoryProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *DirectoryProfiles) GetByName(key string) *DirectoryProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *DirectoryProfileParams) GetByName(key string) *DirectoryProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *DirectoryProfileParams) ClearUp(conf *Directory) {
	l.mx.Lock()
	defer l.mx.Unlock()
	for _, val := range l.byId {
		list := conf.Profiles.GetById(val.Profile.Id)
		if list == nil {
			delete(l.byName, val.Name)
			delete(l.byId, val.Id)
		}
	}
}

func (l *DirectoryProfiles) GetList() map[int64]*DirectoryProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*DirectoryProfile)
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func (l *DirectoryProfiles) Props() []*DirectoryProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*DirectoryProfile
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLDirectory() *Configuration {
	if c.Directory == nil || !c.Directory.Enabled {
		return nil
	}
	c.Directory.XMLItems()
	currentConfig := Configuration{Name: ConfDirectory, Description: "Directory Config", XMLSettings: &c.Directory.XmlSettings, AnyXML: struct {
		XMLName xml.Name    `xml:"profiles,omitempty"`
		Inner   interface{} `xml:"profile"`
	}{Inner: &c.Directory.XmlProfiles}}
	return &currentConfig
}

func (l *Directory) XMLItems() {
	l.XmlSettings = l.Settings.XMLItems()
	l.XmlProfiles = l.Profiles.XMLItems()
}

func (l *DirectoryProfiles) XMLItems() []interface{} {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var profile []interface{}
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (l *DirectoryProfileParams) XMLItems() []DirectoryProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var param []DirectoryProfileParam
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	return param
}

func (l *DirectoryProfileParams) GetList() map[int64]*DirectoryProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*DirectoryProfileParam)
	// BY ID ONLY!
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func NewFifoFifos() *FifoFifos {
	return &FifoFifos{
		byName: make(map[string]*FifoFifo),
		byId:   make(map[int64]*FifoFifo),
	}
}

func NewFifoFifoParams() *FifoFifoParams {
	return &FifoFifoParams{
		byId: make(map[int64]*FifoFifoMember),
	}
}

type Fifo struct {
	Id          int64           `xml:"-" json:"id"`
	Enabled     bool            `json:"enabled" xml:"-"`
	Loaded      bool            `xml:"-" json:"loaded"`
	Settings    *Params         `json:"-" xml:"-"`
	XmlSettings []Param         `json:"-" xml:"param,omitempty"`
	Fifos       *FifoFifos      `json:"-" xml:"-"`
	XmlFifos    []interface{}   `json:"-" xml:"fifo>member,omitempty"` //xml:"fifo>member,omitempty ?
	FifoParams  *FifoFifoParams `json:"-" xml:"-"`
}

type FifoFifos struct {
	mx     sync.RWMutex
	byName map[string]*FifoFifo
	byId   map[int64]*FifoFifo
}

type FifoFifo struct {
	Id         int64            `xml:"-" json:"id"`
	Enabled    bool             `json:"enabled" xml:"-"`
	Name       string           `json:"name" xml:"name,attr"`
	Importance string           `json:"importance" xml:"importance,attr"`
	Params     *FifoFifoParams  `json:"-" xml:"-"`
	XmlParams  []FifoFifoMember `json:"-" xml:"member,omitempty"`
}

type FifoFifoParams struct {
	mx   sync.RWMutex
	byId map[int64]*FifoFifoMember
}

type FifoFifoMember struct {
	Id      int64     `xml:"-" json:"id"`
	Enabled bool      `xml:"-" json:"enabled"`
	Timeout string    `xml:"timeout,attr" json:"timeout"`
	Simo    string    `xml:"simo,attr" json:"simo"`
	Lag     string    `xml:"lag,attr" json:"lag"`
	Body    string    `xml:",chardata" json:"body"`
	Fifo    *FifoFifo `xml:"-" json:"-"`
}

func (c *Configurations) NewFifo(id int64, enabled bool) {
	if c.Fifo != nil {
		return
	}
	c.Fifo = &Fifo{
		Id:         id,
		Enabled:    enabled,
		Settings:   NewParams(),
		Fifos:      NewFifoFifos(),
		FifoParams: NewFifoFifoParams(),
	}
}

func (l *Fifo) Reload() string {
	return "reload " + l.GetModuleName()
}
func (l *Fifo) Unload() string {
	return "unload " + l.GetModuleName()
}
func (l *Fifo) Load() string {
	return "load " + l.GetModuleName()
}
func (l *Fifo) Switch(enabled bool) {
	l.Enabled = enabled
}
func (l *Fifo) AutoLoad() {

}
func (l *Fifo) GetId() int64 {
	return l.Id
}
func (l *Fifo) SetLoadStatus(status bool) {
	l.Loaded = status
}
func (l *Fifo) GetConfig() *Configurations {
	return &Configurations{Fifo: l}
}
func (l *Fifo) GetModuleName() string {
	return ModFifo
}
func (l *Fifo) IsNil() bool {
	return l == nil
}

func (l *FifoFifos) Set(value *FifoFifo) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *FifoFifoParams) Set(value *FifoFifoMember) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byId[value.Id] = value
}

func (l *FifoFifos) Remove(key *FifoFifo) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *FifoFifoParams) Remove(key *FifoFifoMember) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byId, key.Id)
}

func (l *FifoFifos) GetById(key int64) *FifoFifo {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *FifoFifoParams) GetById(key int64) *FifoFifoMember {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *FifoFifos) GetByName(key string) *FifoFifo {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *FifoFifoParams) ClearUp(conf *Fifo) {
	l.mx.Lock()
	defer l.mx.Unlock()
	for _, val := range l.byId {
		list := conf.Fifos.GetById(val.Fifo.Id)
		if list == nil {
			delete(l.byId, val.Id)
		}
	}
}

func (l *FifoFifos) GetList() map[int64]*FifoFifo {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*FifoFifo)
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func (l *FifoFifos) Props() []*FifoFifo {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*FifoFifo
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLFifo() *Configuration {
	if c.Fifo == nil || !c.Fifo.Enabled {
		return nil
	}
	c.Fifo.XMLItems()
	currentConfig := Configuration{Name: ConfFifo, Description: "Fifo Config", XMLSettings: &c.Fifo.XmlSettings, AnyXML: struct {
		XMLName xml.Name    `xml:"fifos,omitempty"`
		Inner   interface{} `xml:"fifo"`
	}{Inner: &c.Fifo.XmlFifos}}
	return &currentConfig
}

func (l *Fifo) XMLItems() {
	l.XmlSettings = l.Settings.XMLItems()
	l.XmlFifos = l.Fifos.XMLItems()
}

func (l *FifoFifos) XMLItems() []interface{} {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var profile []interface{}
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (l *FifoFifoParams) XMLItems() []FifoFifoMember {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var param []FifoFifoMember
	for _, val := range l.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	return param
}

func (l *FifoFifoParams) GetList() map[int64]*FifoFifoMember {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*FifoFifoMember)
	// BY ID ONLY!
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func NewOpalListeners() *OpalListeners {
	return &OpalListeners{
		byName: make(map[string]*OpalListener),
		byId:   make(map[int64]*OpalListener),
	}
}

func NewOpalListenerParams() *OpalListenerParams {
	return &OpalListenerParams{
		byName: make(map[string]*OpalListenerParam),
		byId:   make(map[int64]*OpalListenerParam),
	}
}

type Opal struct {
	Id             int64               `xml:"-" json:"id"`
	Enabled        bool                `json:"enabled" xml:"-"`
	Loaded         bool                `xml:"-" json:"loaded"`
	Settings       *Params             `json:"-" xml:"-"`
	XmlSettings    []Param             `json:"-" xml:"param,omitempty"`
	Listeners      *OpalListeners      `json:"-" xml:"-"`
	XmlListeners   []interface{}       `json:"-" xml:"profiles>profile,omitempty"`
	ListenerParams *OpalListenerParams `json:"-" xml:"-"`
}

type OpalListeners struct {
	mx     sync.RWMutex
	byName map[string]*OpalListener
	byId   map[int64]*OpalListener
}

type OpalListener struct {
	Id        int64               `xml:"-" json:"id"`
	Enabled   bool                `json:"enabled" xml:"-"`
	Name      string              `json:"name" xml:"name,attr"`
	Params    *OpalListenerParams `json:"-" xml:"-"`
	XmlParams []OpalListenerParam `json:"-" xml:"param,omitempty"`
}

type OpalListenerParams struct {
	mx     sync.RWMutex
	byName map[string]*OpalListenerParam
	byId   map[int64]*OpalListenerParam
}

type OpalListenerParam struct {
	Id       int64         `xml:"-" json:"id"`
	Enabled  bool          `xml:"-" json:"enabled"`
	Name     string        `xml:"name,attr" json:"name"`
	Value    string        `xml:"value,attr" json:"value"`
	Listener *OpalListener `xml:"-" json:"-"`
}

func (c *Configurations) NewOpal(id int64, enabled bool) {
	if c.Opal != nil {
		return
	}
	c.Opal = &Opal{
		Id:             id,
		Enabled:        enabled,
		Settings:       NewParams(),
		Listeners:      NewOpalListeners(),
		ListenerParams: NewOpalListenerParams(),
	}
}

func (l *Opal) Reload() string {
	return "reload " + l.GetModuleName()
}
func (l *Opal) Unload() string {
	return "unload " + l.GetModuleName()
}
func (l *Opal) Load() string {
	return "load " + l.GetModuleName()
}
func (l *Opal) Switch(enabled bool) {
	l.Enabled = enabled
}
func (l *Opal) AutoLoad() {

}
func (l *Opal) GetId() int64 {
	return l.Id
}
func (l *Opal) SetLoadStatus(status bool) {
	l.Loaded = status
}
func (l *Opal) GetConfig() *Configurations {
	return &Configurations{Opal: l}
}
func (l *Opal) GetModuleName() string {
	return ModOpal
}
func (l *Opal) IsNil() bool {
	return l == nil
}

func (l *OpalListeners) Set(value *OpalListener) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *OpalListenerParams) Set(value *OpalListenerParam) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *OpalListeners) Remove(key *OpalListener) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *OpalListenerParams) Remove(key *OpalListenerParam) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *OpalListeners) GetById(key int64) *OpalListener {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *OpalListenerParams) GetById(key int64) *OpalListenerParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *OpalListeners) GetByName(key string) *OpalListener {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *OpalListenerParams) GetByName(key string) *OpalListenerParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *OpalListenerParams) ClearUp(conf *Opal) {
	l.mx.Lock()
	defer l.mx.Unlock()
	for _, val := range l.byId {
		list := conf.Listeners.GetById(val.Listener.Id)
		if list == nil {
			delete(l.byName, val.Name)
			delete(l.byId, val.Id)
		}
	}
}

func (l *OpalListeners) GetList() map[int64]*OpalListener {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*OpalListener)
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func (l *OpalListeners) Props() []*OpalListener {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*OpalListener
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLOpal() *Configuration {
	if c.Opal == nil || !c.Opal.Enabled {
		return nil
	}
	c.Opal.XMLItems()
	currentConfig := Configuration{Name: ConfOpal, Description: "Opal Config", XMLSettings: &c.Opal.XmlSettings, AnyXML: struct {
		XMLName xml.Name    `xml:"listeners,omitempty"`
		Inner   interface{} `xml:"listener"`
	}{Inner: &c.Opal.XmlListeners}}
	return &currentConfig
}

func (l *Opal) XMLItems() {
	l.XmlSettings = l.Settings.XMLItems()
	l.XmlListeners = l.Listeners.XMLItems()
}

func (l *OpalListeners) XMLItems() []interface{} {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var profile []interface{}
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (l *OpalListenerParams) XMLItems() []OpalListenerParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var param []OpalListenerParam
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	return param
}

func (l *OpalListenerParams) GetList() map[int64]*OpalListenerParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*OpalListenerParam)
	// BY ID ONLY!
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func NewOspProfiles() *OspProfiles {
	return &OspProfiles{
		byName: make(map[string]*OspProfile),
		byId:   make(map[int64]*OspProfile),
	}
}

func NewOspProfileParams() *OspProfileParams {
	return &OspProfileParams{
		byName: make(map[string]*OspProfileParam),
		byId:   make(map[int64]*OspProfileParam),
	}
}

type Osp struct {
	Id            int64             `xml:"-" json:"id"`
	Enabled       bool              `json:"enabled" xml:"-"`
	Loaded        bool              `xml:"-" json:"loaded"`
	Settings      *Params           `json:"-" xml:"-"`
	XmlSettings   []Param           `json:"-" xml:"param,omitempty"`
	Profiles      *OspProfiles      `json:"-" xml:"-"`
	XmlProfiles   []interface{}     `json:"-" xml:"profiles>profile,omitempty"`
	ProfileParams *OspProfileParams `json:"-" xml:"-"`
}

type OspProfiles struct {
	mx     sync.RWMutex
	byName map[string]*OspProfile
	byId   map[int64]*OspProfile
}

type OspProfile struct {
	Id        int64             `xml:"-" json:"id"`
	Enabled   bool              `json:"enabled" xml:"-"`
	Name      string            `json:"name" xml:"name,attr"`
	Params    *OspProfileParams `json:"-" xml:"-"`
	XmlParams []OspProfileParam `json:"-" xml:"param,omitempty"`
}

type OspProfileParams struct {
	mx     sync.RWMutex
	byName map[string]*OspProfileParam
	byId   map[int64]*OspProfileParam
}

type OspProfileParam struct {
	Id      int64       `xml:"-" json:"id"`
	Enabled bool        `xml:"-" json:"enabled"`
	Name    string      `xml:"name,attr" json:"name"`
	Value   string      `xml:"value,attr" json:"value"`
	Profile *OspProfile `xml:"-" json:"-"`
}

func (c *Configurations) NewOsp(id int64, enabled bool) {
	if c.Osp != nil {
		return
	}
	c.Osp = &Osp{
		Id:            id,
		Enabled:       enabled,
		Settings:      NewParams(),
		Profiles:      NewOspProfiles(),
		ProfileParams: NewOspProfileParams(),
	}
}

func (l *Osp) Reload() string {
	return "reload " + l.GetModuleName()
}
func (l *Osp) Unload() string {
	return "unload " + l.GetModuleName()
}
func (l *Osp) Load() string {
	return "load " + l.GetModuleName()
}
func (l *Osp) Switch(enabled bool) {
	l.Enabled = enabled
}
func (l *Osp) AutoLoad() {

}
func (l *Osp) GetId() int64 {
	return l.Id
}
func (l *Osp) SetLoadStatus(status bool) {
	l.Loaded = status
}
func (l *Osp) GetConfig() *Configurations {
	return &Configurations{Osp: l}
}
func (l *Osp) GetModuleName() string {
	return ModOsp
}
func (l *Osp) IsNil() bool {
	return l == nil
}

func (l *OspProfiles) Set(value *OspProfile) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *OspProfileParams) Set(value *OspProfileParam) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *OspProfiles) Remove(key *OspProfile) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *OspProfileParams) Remove(key *OspProfileParam) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *OspProfiles) GetById(key int64) *OspProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *OspProfileParams) GetById(key int64) *OspProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *OspProfiles) GetByName(key string) *OspProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *OspProfileParams) GetByName(key string) *OspProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *OspProfileParams) ClearUp(conf *Osp) {
	l.mx.Lock()
	defer l.mx.Unlock()
	for _, val := range l.byId {
		list := conf.Profiles.GetById(val.Profile.Id)
		if list == nil {
			delete(l.byName, val.Name)
			delete(l.byId, val.Id)
		}
	}
}

func (l *OspProfiles) GetList() map[int64]*OspProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*OspProfile)
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func (l *OspProfiles) Props() []*OspProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*OspProfile
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLOsp() *Configuration {
	if c.Osp == nil || !c.Osp.Enabled {
		return nil
	}
	c.Osp.XMLItems()
	currentConfig := Configuration{Name: ConfOsp, Description: "Osp Config", XMLSettings: &c.Osp.XmlSettings, AnyXML: struct {
		XMLName xml.Name    `xml:"profiles,omitempty"`
		Inner   interface{} `xml:"profile"`
	}{Inner: &c.Osp.XmlProfiles}}
	return &currentConfig
}

func (l *Osp) XMLItems() {
	l.XmlSettings = l.Settings.XMLItems()
	l.XmlProfiles = l.Profiles.XMLItems()
}

func (l *OspProfiles) XMLItems() []interface{} {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var profile []interface{}
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (l *OspProfileParams) XMLItems() []OspProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var param []OspProfileParam
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	return param
}

func (l *OspProfileParams) GetList() map[int64]*OspProfileParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*OspProfileParam)
	// BY ID ONLY!
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func NewUnicallSpans() *UnicallSpans {
	return &UnicallSpans{
		byName: make(map[string]*UnicallSpan),
		byId:   make(map[int64]*UnicallSpan),
	}
}

func NewUnicallSpanParams() *UnicallSpanParams {
	return &UnicallSpanParams{
		byName: make(map[string]*UnicallSpanParam),
		byId:   make(map[int64]*UnicallSpanParam),
	}
}

type Unicall struct {
	Id          int64              `xml:"-" json:"id"`
	Enabled     bool               `json:"enabled" xml:"-"`
	Loaded      bool               `xml:"-" json:"loaded"`
	Settings    *Params            `json:"-" xml:"-"`
	XmlSettings []Param            `json:"-" xml:"param,omitempty"`
	Spans       *UnicallSpans      `json:"-" xml:"-"`
	XmlSpans    []interface{}      `json:"-" xml:"spans>span,omitempty"`
	SpanParams  *UnicallSpanParams `json:"-" xml:"-"`
}

type UnicallSpans struct {
	mx     sync.RWMutex
	byName map[string]*UnicallSpan
	byId   map[int64]*UnicallSpan
}

type UnicallSpan struct {
	Id        int64              `xml:"-" json:"id"`
	Enabled   bool               `json:"enabled" xml:"-"`
	SpanId    string             `json:"span_id" xml:"id,attr"`
	Params    *UnicallSpanParams `json:"-" xml:"-"`
	XmlParams []UnicallSpanParam `json:"-" xml:"param,omitempty"`
}

type UnicallSpanParams struct {
	mx     sync.RWMutex
	byName map[string]*UnicallSpanParam
	byId   map[int64]*UnicallSpanParam
}

type UnicallSpanParam struct {
	Id      int64        `xml:"-" json:"id"`
	Enabled bool         `xml:"-" json:"enabled"`
	Name    string       `xml:"name,attr" json:"name"`
	Value   string       `xml:"value,attr" json:"value"`
	Span    *UnicallSpan `xml:"-" json:"-"`
}

func (c *Configurations) NewUnicall(id int64, enabled bool) {
	if c.Unicall != nil {
		return
	}
	c.Unicall = &Unicall{
		Id:         id,
		Enabled:    enabled,
		Settings:   NewParams(),
		Spans:      NewUnicallSpans(),
		SpanParams: NewUnicallSpanParams(),
	}
}

func (l *Unicall) Reload() string {
	return "reload " + l.GetModuleName()
}
func (l *Unicall) Unload() string {
	return "unload " + l.GetModuleName()
}
func (l *Unicall) Load() string {
	return "load " + l.GetModuleName()
}
func (l *Unicall) Switch(enabled bool) {
	l.Enabled = enabled
}
func (l *Unicall) AutoLoad() {

}
func (l *Unicall) GetId() int64 {
	return l.Id
}
func (l *Unicall) SetLoadStatus(status bool) {
	l.Loaded = status
}
func (l *Unicall) GetConfig() *Configurations {
	return &Configurations{Unicall: l}
}
func (l *Unicall) GetModuleName() string {
	return ModUnicall
}
func (l *Unicall) IsNil() bool {
	return l == nil
}

func (l *UnicallSpans) Set(value *UnicallSpan) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.SpanId] = value
	l.byId[value.Id] = value
}

func (l *UnicallSpanParams) Set(value *UnicallSpanParam) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.byName[value.Name] = value
	l.byId[value.Id] = value
}

func (l *UnicallSpans) Remove(key *UnicallSpan) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.SpanId)
	delete(l.byId, key.Id)
}

func (l *UnicallSpanParams) Remove(key *UnicallSpanParam) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	delete(l.byName, key.Name)
	delete(l.byId, key.Id)
}

func (l *UnicallSpans) GetById(key int64) *UnicallSpan {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *UnicallSpanParams) GetById(key int64) *UnicallSpanParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val := l.byId[key]
	return val
}

func (l *UnicallSpans) GetByName(key string) *UnicallSpan {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *UnicallSpanParams) GetByName(key string) *UnicallSpanParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	val, _ := l.byName[key]
	return val
}

func (l *UnicallSpanParams) ClearUp(conf *Unicall) {
	l.mx.Lock()
	defer l.mx.Unlock()
	for _, val := range l.byId {
		list := conf.Spans.GetById(val.Span.Id)
		if list == nil {
			delete(l.byName, val.Name)
			delete(l.byId, val.Id)
		}
	}
}

func (l *UnicallSpans) GetList() map[int64]*UnicallSpan {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*UnicallSpan)
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}

func (l *UnicallSpans) Props() []*UnicallSpan {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*UnicallSpan
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *Configurations) XMLUnicall() *Configuration {
	if c.Unicall == nil || !c.Unicall.Enabled {
		return nil
	}
	c.Unicall.XMLItems()
	currentConfig := Configuration{Name: ConfUnicall, Description: "Unicall Config", XMLSettings: &c.Unicall.XmlSettings, AnyXML: struct {
		XMLName xml.Name    `xml:"spans,omitempty"`
		Inner   interface{} `xml:"span"`
	}{Inner: &c.Unicall.XmlSpans}}
	return &currentConfig
}

func (l *Unicall) XMLItems() {
	l.XmlSettings = l.Settings.XMLItems()
	l.XmlSpans = l.Spans.XMLItems()
}

func (l *UnicallSpans) XMLItems() []interface{} {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var profile []interface{}
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		val.XmlParams = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (l *UnicallSpanParams) XMLItems() []UnicallSpanParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var param []UnicallSpanParam
	for _, val := range l.byName {
		if !val.Enabled {
			continue
		}
		param = append(param, *val)
	}
	return param
}

func (l *UnicallSpanParams) GetList() map[int64]*UnicallSpanParam {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*UnicallSpanParam)
	// BY ID ONLY!
	for _, val := range l.byId {
		list[val.Id] = val
	}
	return list
}
