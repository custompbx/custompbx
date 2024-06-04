package altStruct

type ConfigOdbcCdrSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigOdbcCdrSetting) GetTableName() string {
	return "config_odbc_cdr_settings"
}

type ConfigOdbcCdrTable struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	LogLeg      string              `xml:"log_leg,attr" json:"log_leg" customsql:"log_leg"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigOdbcCdrTable) GetTableName() string {
	return "config_odbc_cdr_tables"
}

type ConfigOdbcCdrTableField struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	ChanVarName string              `xml:"chan-var-name,attr" json:"chan_var_name" customsql:"chan_var_name"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigOdbcCdrTable `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigOdbcCdrTableField) GetTableName() string {
	return "config_odbc_cdr_tables_fields"
}
