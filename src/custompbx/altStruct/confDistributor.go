package altStruct

type ConfigDistributorList struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"list_name;unique_1;check(list_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigDistributorList) GetTableName() string {
	return "config_distributor_lists"
}

type ConfigDistributorListNode struct {
	Id          int64                  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                  `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                   `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                 `xml:"name,attr" json:"name" customsql:"node_name;unique_1;check(node_name <> '')"`
	Weight      string                 `xml:"weight,attr" json:"weight" customsql:"node_weight"`
	Description string                 `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigDistributorList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigDistributorListNode) GetTableName() string {
	return "config_distributor_list_nodes"
}
