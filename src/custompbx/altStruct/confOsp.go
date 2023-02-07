package altStruct

type ConfigOspSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigOspSetting) GetTableName() string {
	return "config_osp_settings"
}

type ConfigOspProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigOspProfile) GetTableName() string {
	return "config_osp_profiles"
}

type ConfigOspProfileParameter struct {
	Id          int64             `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64             `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool              `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string            `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string            `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string            `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigOspProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigOspProfileParameter) GetTableName() string {
	return "config_osp_profile_params"
}
