package altStruct

type ConfigUnicallSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigUnicallSetting) GetTableName() string {
	return "config_unicall_settings"
}

type ConfigUnicallSpan struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	SpanId      string              `xml:"id,attr" json:"span_id" customsql:"span_id;unique;check(span_id <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigUnicallSpan) GetTableName() string {
	return "config_unicall_spans"
}

type ConfigUnicallSpanParameter struct {
	Id          int64              `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64              `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool               `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string             `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string             `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string             `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigUnicallSpan `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigUnicallSpanParameter) GetTableName() string {
	return "config_unicall_span_params"
}
