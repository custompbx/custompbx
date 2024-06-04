package altStruct

type ConfigCdrPgCsvSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigCdrPgCsvSetting) GetTableName() string {
	return "config_cdr_pg_csv_setting"
}

type ConfigCdrPgCsvSchema struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Var         string              `xml:"var,attr" json:"var" customsql:"var;unique_1;check(var <> '')"`
	Column      string              `xml:"column,attr" json:"column" customsql:"column_name"`
	Quote       string              `xml:"quote,attr" json:"quote" customsql:"quote"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigCdrPgCsvSchema) GetTableName() string {
	return "config_cdr_pg_csv_schema"
}
