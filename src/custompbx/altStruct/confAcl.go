package altStruct

type ConfigAclList struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"list_name;unique;check(list_name <> '')"`
	Default     string              `xml:"default,attr" json:"default" customsql:"list_default"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigAclList) GetTableName() string {
	return "config_acl_lists"
}

type ConfigAclNode struct {
	Id          int64          `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64          `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool           `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Type        string         `xml:"type,attr" json:"type,omitempty" customsql:"node_type;check(node_type <> '')"`
	Cidr        string         `xml:"cidr,attr" json:"cidr" customsql:"cidr"`
	Domain      string         `xml:"domain,attr" json:"domain" customsql:"domain"`
	Description string         `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigAclList `xml:"-" json:"parent" customsql:"fkey:parent_id;check(parent_id <> 0)"`
}

func (w *ConfigAclNode) GetTableName() string {
	return "config_acl_nodes"
}
