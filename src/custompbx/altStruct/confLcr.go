package altStruct

type ConfigLcrSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigLcrSetting) GetTableName() string {
	return "config_lcr_settings"
}

type ConfigLcrProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigLcrProfile) GetTableName() string {
	return "config_lcr_profiles"
}

type ConfigLcrProfileParameter struct {
	Id          int64             `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64             `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool              `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string            `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string            `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string            `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigLcrProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigLcrProfileParameter) GetTableName() string {
	return "config_lcr_profile_parameters"
}
