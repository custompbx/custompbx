package altStruct

type ConfigDirectorySetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigDirectorySetting) GetTableName() string {
	return "config_directory_settings"
}

type ConfigDirectoryProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigDirectoryProfile) GetTableName() string {
	return "config_directory_profiles"
}

type ConfigDirectoryProfileParameter struct {
	Id          int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                  `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string                  `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string                  `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigDirectoryProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigDirectoryProfileParameter) GetTableName() string {
	return "config_directory_profile_parameters"
}
