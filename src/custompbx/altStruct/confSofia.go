package altStruct

type ConfigSofiaGlobalSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigSofiaGlobalSetting) GetTableName() string {
	return "config_sofia_global_settings"
}

type ConfigSofiaProfile struct {
	Id       int64 `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position int64 `xml:"-" json:"position" customsql:"position;position"`
	Enabled  bool  `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	//TODO: later
	//Name        string              `xml:"name,attr" json:"name" customsql:"profile_name;unique;check(profile_name <> '')"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
	Started     bool                `xml:"-" json:"started"`
	State       string              `xml:"-" json:"state"`
	Uri         string              `xml:"-" json:"uri"`
}

func (w *ConfigSofiaProfile) GetTableName() string {
	return "config_sofia_profiles"
}

type ConfigSofiaProfileParameter struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigSofiaProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigSofiaProfileParameter) GetTableName() string {
	return "config_sofia_profile_parameters"
}

type ConfigSofiaProfileAlias struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigSofiaProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigSofiaProfileAlias) GetTableName() string {
	return "config_sofia_profile_aliases"
}

type ConfigSofiaProfileDomain struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Alias       bool                `xml:"alias,attr" json:"alias" customsql:"alias"`
	Parse       bool                `xml:"parse,attr" json:"parse" customsql:"parse"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigSofiaProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigSofiaProfileDomain) GetTableName() string {
	return "config_sofia_profile_domains"
}

type ConfigSofiaProfileGateway struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigSofiaProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
	Started     bool                `xml:"-" json:"started"`
	State       string              `xml:"-" json:"state"`
}

func (w *ConfigSofiaProfileGateway) GetTableName() string {
	return "config_sofia_profile_gateways"
}

type ConfigSofiaProfileGatewayParameter struct {
	Id          int64                      `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                      `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                       `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                     `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string                     `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string                     `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigSofiaProfileGateway `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigSofiaProfileGatewayParameter) GetTableName() string {
	return "config_sofia_profile_gateway_params"
}

type ConfigSofiaProfileGatewayVariable struct {
	Id          int64                      `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                      `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                       `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                     `xml:"name,attr" json:"name" customsql:"var_name;unique;check(var_name <> '')"`
	Value       string                     `xml:"value,attr" json:"value" customsql:"var_value"`
	Description string                     `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigSofiaProfileGateway `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigSofiaProfileGatewayVariable) GetTableName() string {
	return "config_sofia_profile_gateway_vars"
}
