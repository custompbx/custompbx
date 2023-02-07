package altStruct

type ConfigPostLoadSwitchSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigPostLoadSwitchSetting) GetTableName() string {
	return "config_post_switch_settings"
}

type ConfigPostLoadSwitchCliKeybinding struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name,omitempty" customsql:"key_name;unique;check(key_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"key_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigPostLoadSwitchCliKeybinding) GetTableName() string {
	return "config_post_switch_cli_keybindings"
}

type ConfigPostLoadSwitchDefaultPtime struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	CodecName   string              `xml:"codec_name,attr" json:"codec_name" customsql:"codec_name;unique;check(codec_name <> '')"`
	CodecPtime  string              `xml:"codec_ptime,attr" json:"codec_ptime" customsql:"codec_ptime"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigPostLoadSwitchDefaultPtime) GetTableName() string {
	return "config_post_switch_default_ptimes"
}
