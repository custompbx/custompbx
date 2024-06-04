package altStruct

import "custompbx/mainStruct"

type DirectoryDomain struct {
	Id          int64                  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                  `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                   `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                 `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Description string                 `xml:"-" json:"description,omitempty" customsql:"description"`
	Parent      *mainStruct.FsInstance `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`

	SipRegsCounter int `xml:"-" json:"sip_regs_counter"`
}

func (w *DirectoryDomain) GetTableName() string {
	return "directory_domains"
}

type DirectoryDomainParameter struct {
	Id          int64            `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64            `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool             `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string           `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Value       string           `xml:"value,attr" json:"value" customsql:"value"`
	Description string           `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomain `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainParameter) GetTableName() string {
	return "directory_domain_parameters"
}

type DirectoryDomainVariable struct {
	Id          int64            `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64            `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool             `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string           `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Value       string           `xml:"value,attr" json:"value" customsql:"value"`
	Description string           `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomain `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainVariable) GetTableName() string {
	return "directory_domain_variables"
}

type DirectoryDomainUser struct {
	Id          int64            `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64            `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool             `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string           `xml:"id,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Cache       uint             `xml:"cacheable,attr,omitempty" json:"cache" customsql:"cache"`
	Cidr        string           `xml:"cidr,attr,omitempty" json:"cidr" customsql:"cidr"`
	NumberAlias string           `xml:"number-alias,attr,omitempty" json:"number_alias" customsql:"number_alias"`
	Description string           `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomain `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`

	CallDate      int64  `xml:"-" json:"call_date"`
	InCall        bool   `xml:"-" json:"in_call"`
	Talking       bool   `xml:"-" json:"talking"`
	LastUuid      string `xml:"-" json:"last_uuid"`
	CallDirection string `xml:"-" json:"call_direction"`
	SipRegister   bool   `xml:"-" json:"sip_register"`
	VertoRegister bool   `xml:"-" json:"verto_register"`
	CCAgent       int64  `xml:"-" json:"cc_agent,omitempty"`
}

func (w *DirectoryDomainUser) GetTableName() string {
	return "directory_domain_users"
}

type DirectoryDomainUserParameter struct {
	Id          int64                `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                 `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string               `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Value       string               `xml:"value,attr" json:"value" customsql:"value"`
	Description string               `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomainUser `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainUserParameter) GetTableName() string {
	return "directory_domain_user_parameters"
}

type DirectoryDomainUserVariable struct {
	Id          int64                `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                 `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string               `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Value       string               `xml:"value,attr" json:"value" customsql:"value"`
	Description string               `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomainUser `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainUserVariable) GetTableName() string {
	return "directory_domain_user_variables"
}

type DirectoryDomainUserGateway struct {
	Id          int64                `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                 `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string               `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Description string               `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomainUser `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainUserGateway) GetTableName() string {
	return "directory_domain_user_gateways"
}

type DirectoryDomainUserGatewayParameter struct {
	Id          int64                       `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                       `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                        `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                      `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Value       string                      `xml:"value,attr" json:"value" customsql:"value"`
	Description string                      `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomainUserGateway `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainUserGatewayParameter) GetTableName() string {
	return "directory_domain_user_gateway_parameters"
}

type DirectoryDomainUserGatewayVariable struct {
	Id          int64                       `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                       `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                        `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                      `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Value       string                      `xml:"value,attr" json:"value" customsql:"value"`
	Direction   string                      `xml:"direction,attr" json:"direction" customsql:"direction"`
	Description string                      `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomainUserGateway `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainUserGatewayVariable) GetTableName() string {
	return "directory_domain_user_gateway_variables"
}

type DirectoryDomainGroup struct {
	Id          int64            `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64            `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool             `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string           `xml:"name,attr" json:"name" customsql:"name;unique_1;check(name <> '')"`
	Description string           `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomain `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *DirectoryDomainGroup) GetTableName() string {
	return "directory_domain_groups"
}

type DirectoryDomainGroupUser struct {
	Id          int64                 `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                 `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                  `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Type        string                `xml:"type,attr" json:"type" customsql:"type;default='pointer'"`
	Description string                `xml:"-" json:"description" customsql:"description"`
	Parent      *DirectoryDomainGroup `xml:"-" json:"parent" customsql:"fkey:parent_id;check(parent_id <> 0)"`
	UserId      *DirectoryDomainUser  `xml:"-" json:"user" customsql:"fkey:user_id;check(user_id <> 0)"`
	Name        string                `xml:"id,attr" json:"-"`
}

func (w *DirectoryDomainGroupUser) GetTableName() string {
	return "directory_domain_group_users"
}
