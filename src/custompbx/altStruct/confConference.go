package altStruct

type ConfigConferenceAdvertiseRoom struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"room_name;unique_1;check(room_name <> '')"`
	Status      string              `xml:"status,attr" json:"status" customsql:"room_status"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceAdvertiseRoom) GetTableName() string {
	return "config_conference_advertise_rooms"
}

type ConfigConferenceCallerControlGroup struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"group_name;unique_1;check(group_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceCallerControlGroup) GetTableName() string {
	return "config_conference_caller_control_groups"
}

type ConfigConferenceCallerControlGroupControl struct {
	Id          int64                               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Action      string                              `xml:"action,attr" json:"action" customsql:"control_action;unique_1;check(control_action <> '')"`
	Digits      string                              `xml:"digits,attr" json:"digits" customsql:"control_digits"`
	Description string                              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigConferenceCallerControlGroup `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceCallerControlGroupControl) GetTableName() string {
	return "config_conference_caller_control_group_controls"
}

type ConfigConferenceChatPermissionProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"profile_name;unique_1;check(profile_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceChatPermissionProfile) GetTableName() string {
	return "config_conference_profiles"
}

type ConfigConferenceChatPermissionProfileUser struct {
	Id          int64                                  `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                                  `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                                   `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                                 `xml:"name,attr" json:"name" customsql:"user_name;unique_1;check(user_name <> '')"`
	Commands    string                                 `xml:"commands,attr" json:"commands" customsql:"user_commands"`
	Description string                                 `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigConferenceChatPermissionProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceChatPermissionProfileUser) GetTableName() string {
	return "config_conference_chat_permission_profile_users"
}

type ConfigConferenceProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"profile_name;unique_1;check(profile_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceProfile) GetTableName() string {
	return "config_conference_profiles"
}

type ConfigConferenceProfileParameter struct {
	Id          int64                    `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                    `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                     `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                   `xml:"name,attr" json:"name" customsql:"param_name;unique_1;check(param_name <> '')"`
	Value       string                   `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string                   `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigConferenceProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique_1;check(parent_id <> 0)"`
}

func (w *ConfigConferenceProfileParameter) GetTableName() string {
	return "config_conference_profile_parameters"
}
