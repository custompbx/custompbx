package apps

import (
	"custompbx/altStruct"
)

type AutoDialerCompany struct {
	Id          int64                      `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                      `json:"position" customsql:"position;position"`
	Name        string                     `json:"name" customsql:"name;unique_1;check(name <> '')"`
	Enabled     bool                       `json:"enabled" customsql:"enabled;default=TRUE"`
	Predictive  bool                       `json:"predictive" customsql:"predictive;default=FALSE"`
	Description string                     `json:"description,omitempty" customsql:"description"`
	Domain      *altStruct.DirectoryDomain `json:"domain,omitempty" customsql:"fkey:domain_id;unique_1;check(parent_id <> 0)"`
	Reducer     *AutoDialerReducer         `json:"reducer,omitempty" customsql:"fkey:reducer_id;null"`
	Team        *AutoDialerTeam            `json:"team,omitempty" customsql:"fkey:team_id;null"`
	List        *AutoDialerList            `json:"list,omitempty" customsql:"fkey:list_id;null"`
}

func (w *AutoDialerCompany) GetTableName() string {
	return "app_auto_dialer_companies"
}

type AutoDialerTeam struct {
	Id          int64                      `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                      `json:"position" customsql:"position;position"`
	Name        string                     `json:"name" customsql:"name;unique_1;check(name <> '')"`
	Enabled     bool                       `json:"enabled" customsql:"enabled;default=TRUE"`
	Description string                     `json:"description,omitempty" customsql:"description"`
	Domain      *altStruct.DirectoryDomain `json:"domain,omitempty" customsql:"fkey:domain_id;unique_1;check(parent_id <> 0)"`
}

func (w *AutoDialerTeam) GetTableName() string {
	return "app_auto_dialer_teams"
}

// now limit to only directory members
type AutoDialerTeamMember struct {
	Id          int64                          `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                          `json:"position" customsql:"position;position"`
	Enabled     bool                           `json:"enabled" customsql:"enabled;default=TRUE"`
	Description string                         `json:"description,omitempty" customsql:"description"`
	User        *altStruct.DirectoryDomainUser `json:"user,omitempty" customsql:"fkey:user_id;unique_1;check(user_id <> 0)"`
	Parent      *AutoDialerTeam                `json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *AutoDialerTeamMember) GetTableName() string {
	return "app_auto_dialer_team_members"
}

type AutoDialerList struct {
	Id          int64                      `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                      `json:"position" customsql:"position;position"`
	Name        string                     `json:"name" customsql:"name;unique_1;check(name <> '')"`
	Enabled     bool                       `json:"enabled" customsql:"enabled;default=TRUE"`
	Description string                     `json:"description,omitempty" customsql:"description"`
	Domain      *altStruct.DirectoryDomain `json:"domain,omitempty" customsql:"fkey:domain_id;unique_1;check(parent_id <> 0)"`
}

func (w *AutoDialerList) GetTableName() string {
	return "app_auto_dialer_lists"
}

type AutoDialerListMember struct {
	Id         int64           `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position   int64           `json:"position" customsql:"position;position"`
	Name       string          `json:"name" customsql:"name"`
	ToNumber   string          `json:"to_number" customsql:"to_number;check(to_number <> '')"`
	FromNumber string          `json:"from_number" customsql:"from_number;check(from_number <> '')"`
	Retries    int64           `json:"retries" customsql:"retries;default=0"`
	CustomVars string          `json:"custom_vars" customsql:"custom_vars"`
	Enabled    bool            `json:"enabled" customsql:"enabled;default=TRUE"`
	Tried      int64           `json:"tried" customsql:"tried;default=0"`
	Talked     int64           `json:"talked" customsql:"talked;default=0"`
	RespReason string          `json:"resp_reason" customsql:"resp_reason"`
	Comment    string          `json:"comment" customsql:"comment"`
	Parent     *AutoDialerList `json:"parent" customsql:"fkey:parent_id;check(parent_id <> 0)"`
}

func (w *AutoDialerListMember) GetTableName() string {
	return "app_auto_dialer_list_members"
}

type AutoDialerReducer struct {
	Id          int64                      `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                      `json:"position" customsql:"position;position"`
	Name        string                     `json:"name" customsql:"name;unique_1;check(name <> '')"`
	Enabled     bool                       `json:"enabled" customsql:"enabled;default=TRUE"`
	Description string                     `json:"description,omitempty" customsql:"description"`
	Domain      *altStruct.DirectoryDomain `json:"domain,omitempty" customsql:"fkey:domain_id;unique_1;check(domain_id <> 0)""`
}

func (w *AutoDialerReducer) GetTableName() string {
	return "app_auto_dialer_reducers"
}

type AutoDialerReducerMember struct {
	Id          int64              `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64              `json:"position" customsql:"position;position"`
	Application string             `json:"application" customsql:"application"`
	Data        string             `json:"data,omitempty" customsql:"data"`
	Enabled     bool               `json:"enabled" customsql:"enabled;default=TRUE"`
	Description string             `json:"description,omitempty" customsql:"description"`
	Parent      *AutoDialerReducer `json:"parent,omitempty" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *AutoDialerReducerMember) GetTableName() string {
	return "app_auto_dialer_reducer_members"
}

type AutoDialerProceed struct {
	Id      int64 `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Started bool  `json:"running" customsql:"running"`
	Running bool  `json:"running" customsql:"running"`

	Parent *AutoDialerCompany    `json:"parent,omitempty" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
	Member *AutoDialerListMember `json:"member,omitempty" customsql:"fkey:member_id;unique_1;check(member_id <> 0)"`
}

func (w *AutoDialerProceed) GetTableName() string {
	return "app_auto_dialer_proceed"
}
