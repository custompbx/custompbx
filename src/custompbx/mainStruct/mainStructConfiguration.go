package mainStruct

import (
	"reflect"
	"strings"
)

// will add all the fields og any configs
type Configuration struct {
	Name        string         `xml:"name,attr"`
	Description string         `xml:"description,attr"`
	XMLProfiles *[]interface{} `xml:"profiles>profile,omitempty"`
	AnyXML      interface{}
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

func CheckQueueCommand(command string) bool {
	comMap := map[string]bool{
		CommandCallcenterQueueLoad:   true,
		CommandCallcenterQueueReload: true,
		CommandCallcenterQueueUnload: true,
	}
	_, ok := comMap[command]

	return ok
}
