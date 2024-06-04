package altStruct

type ConfigFifoSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigFifoSetting) GetTableName() string {
	return "config_fifo_settings"
}

type ConfigFifoFifo struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Importance  string              `xml:"importance,attr" json:"importance" customsql:"importance"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigFifoFifo) GetTableName() string {
	return "config_fifo_fifos"
}

type ConfigFifoFifoMember struct {
	Id          int64           `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64           `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool            `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Timeout     string          `xml:"timeout,attr" json:"timeout" customsql:"timeout"`
	Simo        string          `xml:"simo,attr" json:"simo" customsql:"simo"`
	Lag         string          `xml:"lag,attr" json:"lag" customsql:"lag"`
	Body        string          `xml:",chardata" json:"body" customsql:"body"`
	Description string          `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigFifoFifo `xml:"-" json:"parent" customsql:"fkey:parent_id;check(parent_id <> 0)"`
}

func (w *ConfigFifoFifoMember) GetTableName() string {
	return "config_fifo_fifo_members"
}
