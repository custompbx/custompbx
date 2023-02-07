package altStruct

type ConfigCallcenterSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigCallcenterSetting) GetTableName() string {
	return "config_callcenter_settings"
}

type ConfigCallcenterQueue struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"queue_name;unique;check(queue_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigCallcenterQueue) GetTableName() string {
	return "config_callcenter_queues"
}

type ConfigCallcenterQueueParameter struct {
	Id          int64                  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                  `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                   `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                 `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string                 `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string                 `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigCallcenterQueue `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigCallcenterQueueParameter) GetTableName() string {
	return "config_callcenter_queue_parameters"
}

type Agent struct {
	Id                int64  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Name              string `xml:"-" json:"name" customsql:"name"`
	Type              string `xml:"-" json:"type" customsql:"type"`
	System            string `xml:"-" json:"system" customsql:"system"`
	InstanceId        string `xml:"-" json:"instance_id" customsql:"instance_id"`
	Uuid              string `xml:"-" json:"uuid" customsql:"uuid"`
	Contact           string `xml:"-" json:"contact" customsql:"contact"`
	Status            string `xml:"-" json:"status" customsql:"status"`
	State             string `xml:"-" json:"state" customsql:"state"`
	MaxNoAnswer       int64  `xml:"-" json:"max_no_answer" customsql:"max_no_answer"`
	WrapUpTime        int64  `xml:"-" json:"wrap_up_time" customsql:"wrap_up_time"`
	RejectDelayTime   int64  `xml:"-" json:"reject_delay_time" customsql:"reject_delay_time"`
	BusyDelayTime     int64  `xml:"-" json:"busy_delay_time" customsql:"busy_delay_time"`
	NoAnswerDelayTime int64  `xml:"-" json:"no_answer_delay_time" customsql:"no_answer_delay_time"`
	LastBridgeStart   int64  `xml:"-" json:"last_bridge_start" customsql:"last_bridge_start"`
	LastBridgeEnd     int64  `xml:"-" json:"last_bridge_end" customsql:"last_bridge_end"`
	LastOfferedCall   int64  `xml:"-" json:"last_offered_call" customsql:"last_offered_call"`
	LastStatusChange  int64  `xml:"-" json:"last_status_change" customsql:"last_status_change"`
	NoAnswerCount     int64  `xml:"-" json:"no_answer_count" customsql:"no_answer_count"`
	CallsAnswered     int64  `xml:"-" json:"calls_answered" customsql:"calls_answered"`
	TalkTime          int64  `xml:"-" json:"talk_time" customsql:"talk_time"`
	ReadyTime         int64  `xml:"-" json:"ready_time" customsql:"ready_time"`
}

type Tier struct {
	Id       int64  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Agent    string `xml:"-" json:"agent" customsql:"agent"`
	Queue    string `xml:"-" json:"queue" customsql:"queue"`
	Level    int64  `xml:"-" json:"level" customsql:"level"`
	Position int64  `xml:"-" json:"position" customsql:"position"`
	State    string `xml:"-" json:"state" customsql:"state"`
}

type Member struct {
	Uuid           string `xml:"-" json:"uuid" customsql:"uuid"`
	Queue          string `xml:"-" json:"queue" customsql:"queue"`
	InstanceId     string `xml:"-" json:"instance_id" customsql:"instance_id"`
	SessionUuid    string `xml:"-" json:"session_uuid" customsql:"session_uuid"`
	CidNumber      string `xml:"-" json:"cid_number" customsql:"cid_number"`
	CidName        string `xml:"-" json:"cid_name" customsql:"cid_name"`
	SystemEpoch    int64  `xml:"-" json:"system_epoch" customsql:"system_epoch"`
	JoinedEpoch    int64  `xml:"-" json:"joined_epoch" customsql:"joined_epoch"`
	RejoinedEpoch  int64  `xml:"-" json:"rejoined_epoch" customsql:"rejoined_epoch"`
	BridgeEpoch    int64  `xml:"-" json:"bridge_epoch" customsql:"bridge_epoch"`
	AbandonedEpoch int64  `xml:"-" json:"abandoned_epoch" customsql:"abandoned_epoch"`
	BaseScore      int64  `xml:"-" json:"base_score" customsql:"base_score"`
	SkillScore     int64  `xml:"-" json:"skill_score" customsql:"skill_score"`
	ServingAgent   string `xml:"-" json:"serving_agent" customsql:"serving_agent"`
	ServingSystem  string `xml:"-" json:"serving_system" customsql:"serving_system"`
	State          string `xml:"-" json:"state" customsql:"state"`
}
