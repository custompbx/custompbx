package altStruct

type WebDirectoryUsersTemplate struct {
	Id            int64            `xml:"-" json:"id" customsql:"pkey:id"`
	Name          string           `xml:"-" json:"name,omitempty" customsql:"name;unique_1"`
	Cache         int              `xml:"-" json:"cache,omitempty" customsql:"cache"`
	Cidr          string           `xml:"-" json:"cidr,omitempty" customsql:"cidr"`
	NumberAlias   string           `xml:"-" json:"number_alias,omitempty" customsql:"number_alias"`
	Enabled       bool             `xml:"-" json:"enabled,omitempty" customsql:"enabled"`
	DirectoryName string           `xml:"-" json:"directory_name,omitempty" customsql:"directory_name"`
	Domain        *DirectoryDomain `xml:"-" json:"domain,omitempty" customsql:"fkey:domain_id"`
}

func (w *WebDirectoryUsersTemplate) GetTableName() string {
	return "web_directory_users_templates"
}

type WebDirectoryUsersTemplateParameter struct {
	Id          int64                      `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled     bool                       `xml:"-" json:"enabled,omitempty" customsql:"enabled"`
	Name        string                     `xml:"-" json:"name,omitempty" customsql:"param_name;unique_1"`
	Value       string                     `xml:"-" json:"value,omitempty" customsql:"param_value"`
	Description string                     `xml:"-" json:"description,omitempty" customsql:"param_description"`
	Placeholder string                     `xml:"-" json:"placeholder,omitempty" customsql:"param_placeholder"`
	Editable    bool                       `xml:"-" json:"editable,omitempty" customsql:"param_editable"`
	Show        bool                       `xml:"-" json:"show,omitempty" customsql:"param_show"`
	Required    bool                       `xml:"-" json:"required,omitempty" customsql:"required"`
	Parent      *WebDirectoryUsersTemplate `xml:"-" json:"parent,omitempty" customsql:"fkey:parent_id;unique_1"`
}

func (w *WebDirectoryUsersTemplateParameter) GetTableName() string {
	return "web_directory_users_template_parameters"
}

type WebDirectoryUsersTemplateVariable struct {
	Id          int64                      `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled     bool                       `xml:"-" json:"enabled" customsql:"enabled"`
	Name        string                     `xml:"-" json:"name,omitempty" customsql:"var_name;unique_1"`
	Value       string                     `xml:"-" json:"value,omitempty" customsql:"var_value"`
	Description string                     `xml:"-" json:"description,omitempty" customsql:"var_description"`
	Placeholder string                     `xml:"-" json:"placeholder,omitempty" customsql:"var_placeholder"`
	Editable    bool                       `xml:"-" json:"editable,omitempty" customsql:"var_editable"`
	Show        bool                       `xml:"-" json:"show,omitempty" customsql:"var_show"`
	Required    bool                       `xml:"-" json:"required,omitempty" customsql:"required"`
	Parent      *WebDirectoryUsersTemplate `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1"`
}

func (w *WebDirectoryUsersTemplateVariable) GetTableName() string {
	return "web_directory_users_template_variables"
}
