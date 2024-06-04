package altStruct

type ConfigVoicemailSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigVoicemailSetting) GetTableName() string {
	return "config_voicemail_settings"
}

type ConfigVoicemailProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"profile_name;unique_1;check(profile_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigVoicemailProfile) GetTableName() string {
	return "config_voicemail_profiles"
}

type ConfigVoicemailProfileParameter struct {
	Id          int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                  `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string                  `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string                  `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigVoicemailProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigVoicemailProfileParameter) GetTableName() string {
	return "config_voicemail_profile_parameters"
}

type ConfigVoicemailProfileEmailParameter struct {
	Id          int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                  `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string                  `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string                  `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigVoicemailProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigVoicemailProfileEmailParameter) GetTableName() string {
	return "config_voicemail_profile_email_parameters"
}
