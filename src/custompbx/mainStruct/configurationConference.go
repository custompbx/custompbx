package mainStruct

import (
	"encoding/xml"
	"log"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

type ConfigList struct {
	Id          int64  `xml:"-" json:"id"`
	Enabled     bool   `xml:"-" json:"enabled"`
	Loaded      bool   `xml:"-" json:"loaded"`
	Name        string `xml:"-" json:"name"`
	ModName     string `xml:"-" json:"mod_name"`
	Description string `xml:"-" json:"description"`
}

type Conference struct {
	ConfigList
	Advertise               *Advertise                   `json:"-" xml:"-"`
	XmlAdvertise            interface{}                  `json:"-" xml:"advertise>room,omitempty"`
	CallerControlsGroups    *CallerControlsGroups        `json:"-" xml:"-"`
	XmlCallerControls       interface{}                  `json:"-" xml:"caller-controls>group,omitempty"`
	ChatPermissions         *ChatPermissions             `json:"-" xml:"-"`
	XmlChatPermissions      interface{}                  `json:"-" xml:"chat-permissions>group,omitempty"`
	Profiles                *ConferenceProfiles          `json:"-" xml:"-"`
	XmlProfiles             []interface{}                `json:"-" xml:"profiles>profile,omitempty"`
	Controls                *Controls                    `json:"-" xml:"-"`
	ConferenceProfileParams *ConferenceProfileParams     `json:"-" xml:"-"`
	Users                   *ChatPermissionsProfileUsers `json:"-" xml:"-"`
}

type Advertise struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceAdvertiseRooms
	byName map[string]*ConfigConferenceAdvertiseRooms
	Parent *Conference
}

type ConfigConferenceAdvertiseRooms struct {
	Id      int64      `xml:"-" json:"id" customsql:"pkey:id"`
	Name    string     `xml:"name,attr" json:"name" customsql:"room_name;unique"`
	Status  string     `xml:"status,attr" json:"status" customsql:"room_state"`
	Enabled bool       `xml:"-" json:"enabled" customsql:"enabled"`
	Parent  *Advertise `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type CallerControlsGroups struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceCallerControlsGroups
	byName map[string]*ConfigConferenceCallerControlsGroups
	Parent *Conference
}

type ConfigConferenceCallerControlsGroups struct {
	Id          int64                                     `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled     bool                                      `xml:"-" json:"enabled" customsql:"enabled"`
	Name        string                                    `xml:"name,attr" json:"name" customsql:"group_name;unique"`
	Controls    *Controls                                 `xml:"-" json:"-"`
	XMLControls []*ConfigConferenceCallerControlsControls `xml:"control" json:"-"`
	Parent      *CallerControlsGroups                     `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type Controls struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceCallerControlsControls
	byName map[string]*ConfigConferenceCallerControlsControls
	Parent *ConfigConferenceCallerControlsGroups
}

type ConfigConferenceCallerControlsControls struct {
	Id      int64     `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled bool      `xml:"-" json:"enabled" customsql:"enabled"`
	Action  string    `xml:"action,attr" json:"action" customsql:"control_action;unique"`
	Digits  string    `xml:"digits,attr" json:"digits" customsql:"control_digits"`
	Parent  *Controls `xml:"-" json:"-" customsql:"fkey:group_id;unique"`
}

type ConferenceProfiles struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceProfiles
	byName map[string]*ConfigConferenceProfiles
	Parent *Conference
}

type ConfigConferenceProfiles struct {
	Id       int64                             `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled  bool                              `xml:"-" json:"enabled" customsql:"enabled"`
	Name     string                            `xml:"name,attr" json:"name" customsql:"profile_name;unique"`
	Params   *ConferenceProfileParams          `xml:"-" json:"-"`
	XMLParam []*ConfigConferenceProfilesParams `xml:"param" json:"-"`
	Parent   *ConferenceProfiles               `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type ConferenceProfileParams struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceProfilesParams
	byName map[string]*ConfigConferenceProfilesParams
	Parent *ConfigConferenceProfiles
}

type ConfigConferenceProfilesParams struct {
	Id      int64                    `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled bool                     `xml:"-" json:"enabled" customsql:"enabled"`
	Name    string                   `xml:"name,attr" json:"name" customsql:"param_name;unique"`
	Value   string                   `xml:"value,attr" json:"value" customsql:"param_value"`
	Parent  *ConferenceProfileParams `xml:"-" json:"-" customsql:"fkey:profile_id;unique"`
}

type ChatPermissions struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceChatPermissions
	byName map[string]*ConfigConferenceChatPermissions
	Parent *Conference
}

type ConfigConferenceChatPermissions struct {
	Id       int64                                  `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled  bool                                   `xml:"-" json:"enabled" customsql:"enabled"`
	Name     string                                 `xml:"name,attr" json:"name" customsql:"profile_name;unique"`
	Users    *ChatPermissionsProfileUsers           `xml:"-" json:"-"`
	XMLParam []*ConfigConferenceChatPermissionUsers `xml:"param" json:"-"`
	Parent   *ChatPermissions                       `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type ChatPermissionsProfileUsers struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceChatPermissionUsers
	byName map[string]*ConfigConferenceChatPermissionUsers
	Parent *ConfigConferenceChatPermissions
}

type ConfigConferenceChatPermissionUsers struct {
	Id       int64                        `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled  bool                         `xml:"-" json:"enabled" customsql:"enabled"`
	Name     string                       `xml:"name,attr" json:"name" customsql:"user_name;unique"`
	Commands string                       `xml:"commands,attr" json:"commands" customsql:"user_commands"`
	Parent   *ChatPermissionsProfileUsers `xml:"-" json:"-" customsql:"fkey:profile_id;unique"`
}

func (c *Configurations) NewConference(id int64, enabled bool) {
	if c.Conference != nil {
		return
	}
	var conference Conference

	conference.Id = id
	conference.Enabled = enabled
	conference.Advertise = NewAdvertiseRooms(&conference)
	conference.CallerControlsGroups = NewCallerControlsGroups(&conference)
	conference.ChatPermissions = NewConferenceChatPermissions(&conference)
	conference.Profiles = NewConferenceProfiles(&conference)
	conference.Controls = NewControls(nil)
	conference.ConferenceProfileParams = NewConferenceProfileParams(nil)
	conference.Users = NewConferenceChatPermissionUsers(nil)

	c.Conference = &conference
}

func NewAdvertiseRooms(c *Conference) *Advertise {
	return &Advertise{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceAdvertiseRooms),
		byName: make(map[string]*ConfigConferenceAdvertiseRooms),
	}
}

func NewCallerControlsGroups(c *Conference) *CallerControlsGroups {
	return &CallerControlsGroups{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceCallerControlsGroups),
		byName: make(map[string]*ConfigConferenceCallerControlsGroups),
	}
}

func NewConferenceProfiles(c *Conference) *ConferenceProfiles {
	return &ConferenceProfiles{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceProfiles),
		byName: make(map[string]*ConfigConferenceProfiles),
	}
}

func NewConferenceChatPermissions(c *Conference) *ChatPermissions {
	return &ChatPermissions{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceChatPermissions),
		byName: make(map[string]*ConfigConferenceChatPermissions),
	}
}

func NewConferenceChatPermissionUsers(c *ConfigConferenceChatPermissions) *ChatPermissionsProfileUsers {
	return &ChatPermissionsProfileUsers{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceChatPermissionUsers),
		byName: make(map[string]*ConfigConferenceChatPermissionUsers),
	}
}

func NewControls(c *ConfigConferenceCallerControlsGroups) *Controls {
	return &Controls{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceCallerControlsControls),
		byName: make(map[string]*ConfigConferenceCallerControlsControls),
	}
}

func NewConferenceProfileParams(c *ConfigConferenceProfiles) *ConferenceProfileParams {
	return &ConferenceProfileParams{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceProfilesParams),
		byName: make(map[string]*ConfigConferenceProfilesParams),
	}
}

func (c *Conference) Reload() string {
	return "reload " + c.GetModuleName()
}
func (c *Conference) Unload() string {
	return "unload " + c.GetModuleName()
}
func (c *Conference) Load() string {
	return "load " + c.GetModuleName()
}
func (c *Conference) Switch(enabled bool) {
	c.Enabled = enabled
}
func (c *Conference) AutoLoad() {

}
func (c *Conference) GetId() int64 {
	return c.Id
}
func (c *Conference) SetLoadStatus(status bool) {
	c.Loaded = status
}
func (c *Conference) GetConfig() *Configurations {
	return &Configurations{Conference: c}
}
func (c *Conference) GetModuleName() string {
	return ModConference
}
func (c *Conference) IsNil() bool {
	return c == nil
}

func (a *Advertise) NewSubItem() *ConfigConferenceAdvertiseRooms {
	return &ConfigConferenceAdvertiseRooms{Parent: a}
}

func (c *CallerControlsGroups) NewSubItem() *ConfigConferenceCallerControlsGroups {
	return &ConfigConferenceCallerControlsGroups{Parent: c}
}

func (a *ConferenceProfiles) NewSubItem() *ConfigConferenceProfiles {
	return &ConfigConferenceProfiles{Parent: a}
}

func (a *ConferenceProfileParams) NewSubItem() *ConfigConferenceProfilesParams {
	return &ConfigConferenceProfilesParams{Parent: a}
}

func (a *Controls) NewSubItem() *ConfigConferenceCallerControlsControls {
	return &ConfigConferenceCallerControlsControls{Parent: a}
}

func (a *ChatPermissions) NewSubItem() *ConfigConferenceChatPermissions {
	return &ConfigConferenceChatPermissions{Parent: a}
}

func (a *ChatPermissionsProfileUsers) NewSubItem() *ConfigConferenceChatPermissionUsers {
	return &ConfigConferenceChatPermissionUsers{Parent: a}
}

func (a *Advertise) Set(value *ConfigConferenceAdvertiseRooms) {
	value.Parent = a
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
	a.byName[value.Name] = value
}

func (c *CallerControlsGroups) Set(value *ConfigConferenceCallerControlsGroups) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value
}

func (c *ConferenceProfiles) Set(value *ConfigConferenceProfiles) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value
}

func (c *Controls) Set(value *ConfigConferenceCallerControlsControls) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Action] = value

	root := c.Parent.Parent.Parent.Controls
	root.mx.RLock()
	defer root.mx.RUnlock()
	root.byId[value.Id] = value
	root.byName[value.Action] = value
}

func (c *ChatPermissions) Set(value *ConfigConferenceChatPermissions) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value
}

func (c *ChatPermissionsProfileUsers) Set(value *ConfigConferenceChatPermissionUsers) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value

	root := c.Parent.Parent.Parent.Users
	root.mx.RLock()
	defer root.mx.RUnlock()
	root.byId[value.Id] = value
	root.byName[value.Name] = value
}

func (c *ConferenceProfileParams) Set(value *ConfigConferenceProfilesParams) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value

	root := c.Parent.Parent.Parent.ConferenceProfileParams
	root.mx.RLock()
	defer root.mx.RUnlock()
	root.byId[value.Id] = value
	root.byName[value.Name] = value
}

func (c *Configurations) XMLConference() *Configuration {
	if c.Conference == nil || !c.Conference.Enabled {
		return nil
	}
	c.Conference.XMLItems()
	currentConfig := Configuration{
		Name:        ConfConference,
		Description: "Conference Config",
		XMLProfiles: &c.Conference.XmlProfiles,
		AnyXML: []interface{}{
			struct {
				XMLName xml.Name    `xml:"advertise,omitempty"`
				Inner   interface{} `xml:"room"`
			}{Inner: c.Conference.XmlAdvertise},
			struct {
				XMLName xml.Name    `xml:"caller-controls,omitempty"`
				Inner   interface{} `xml:"group"`
			}{Inner: c.Conference.XmlCallerControls},
			struct {
				XMLName xml.Name    `xml:"chat-permissions,omitempty"`
				Inner   interface{} `xml:"profile"`
			}{Inner: c.Conference.XmlChatPermissions},
		},
	}
	return &currentConfig
}

func (v *Conference) XMLItems() {
	v.XmlAdvertise = v.Advertise.XMLItems()
	v.XmlCallerControls = v.CallerControlsGroups.XMLItems()
	v.XmlProfiles = v.Profiles.XMLItems()
	v.XmlChatPermissions = v.ChatPermissions.XMLItems()
}

func (v *ChatPermissions) XMLItems() []interface{} {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var profile []interface{}
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		val.XMLParam = val.Users.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (v *ChatPermissionsProfileUsers) XMLItems() []*ConfigConferenceChatPermissionUsers {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []*ConfigConferenceChatPermissionUsers
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, val)
	}
	return param
}

func (v *ConferenceProfiles) XMLItems() []interface{} {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var profile []interface{}
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		val.XMLParam = val.Params.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (v *ConferenceProfileParams) XMLItems() []*ConfigConferenceProfilesParams {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []*ConfigConferenceProfilesParams
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, val)
	}
	return param
}

func (c *CallerControlsGroups) XMLItems() []interface{} {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var profile []interface{}
	for _, val := range c.byId {
		if !val.Enabled {
			continue
		}
		val.XMLControls = val.Controls.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (v *Controls) XMLItems() []*ConfigConferenceCallerControlsControls {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []*ConfigConferenceCallerControlsControls
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, val)
	}
	return param
}

func (v *Advertise) XMLItems() []*ConfigConferenceAdvertiseRooms {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []*ConfigConferenceAdvertiseRooms
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, val)
	}
	return param
}

func SQLSchema(i interface{}) []string {
	t := reflect.TypeOf(i).Elem()
	f, ok := t.FieldByName("byId")
	if !ok {
		return []string{}
	}
	res, ok := f.Tag.Lookup("customsql")
	if !ok {
		return []string{}
	}
	fields := []string{res}

	f, ok = t.FieldByName("Parent")
	if !ok {
		return []string{}
	}
	res, ok = f.Tag.Lookup("customsql")
	if !ok {
		return []string{}
	}
	fields = append(fields, res)

	return fields
}

func GetTableName(i interface{}) string {
	name := getType(i)
	return ToSnakeCase(name)
}

type TableMethods interface {
	NewSubItemInterface() RowItem
	SetFromInterface(interface{})
}

func (a *Advertise) NewSubItemInterface() RowItem {
	return &ConfigConferenceAdvertiseRooms{Parent: a}
}

func (c *CallerControlsGroups) NewSubItemInterface() RowItem {
	return &ConfigConferenceCallerControlsGroups{Parent: c}
}

func (a *Controls) NewSubItemInterface() RowItem {
	return &ConfigConferenceCallerControlsControls{Parent: a}
}

func (a *ConferenceProfiles) NewSubItemInterface() RowItem {
	return &ConfigConferenceProfiles{Parent: a}
}

func (a *ConferenceProfileParams) NewSubItemInterface() RowItem {
	return &ConfigConferenceProfilesParams{Parent: a}
}

func (a *ChatPermissions) NewSubItemInterface() RowItem {
	return &ConfigConferenceChatPermissions{Parent: a}
}

func (a *ChatPermissionsProfileUsers) NewSubItemInterface() RowItem {
	return &ConfigConferenceChatPermissionUsers{Parent: a}
}

func (a *Advertise) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceAdvertiseRooms:
		a.Set(value.(*ConfigConferenceAdvertiseRooms))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (c *CallerControlsGroups) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceCallerControlsGroups:
		t := value.(*ConfigConferenceCallerControlsGroups)
		t.Controls = NewControls(t)
		c.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (a *Controls) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceCallerControlsControls:
		a.Set(value.(*ConfigConferenceCallerControlsControls))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (a *ConferenceProfiles) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceProfiles:
		t := value.(*ConfigConferenceProfiles)
		t.Params = NewConferenceProfileParams(t)
		a.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (a *ConferenceProfileParams) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceProfilesParams:
		a.Set(value.(*ConfigConferenceProfilesParams))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (a *ChatPermissions) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceChatPermissions:
		t := value.(*ConfigConferenceChatPermissions)
		t.Users = NewConferenceChatPermissionUsers(t)
		a.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (a *ChatPermissionsProfileUsers) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceChatPermissionUsers:
		a.Set(value.(*ConfigConferenceChatPermissionUsers))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (l *Advertise) Props() []*ConfigConferenceAdvertiseRooms {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ConfigConferenceAdvertiseRooms
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (c *CallerControlsGroups) Props() []*ConfigConferenceCallerControlsGroups {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*ConfigConferenceCallerControlsGroups
	for _, val := range c.byId {
		items = append(items, val)
	}
	return items
}

func (l *Controls) Props() []*ConfigConferenceCallerControlsControls {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ConfigConferenceCallerControlsControls
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *ConferenceProfiles) Props() []*ConfigConferenceProfiles {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ConfigConferenceProfiles
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *ConferenceProfileParams) Props() []*ConfigConferenceProfilesParams {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ConfigConferenceProfilesParams
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *ChatPermissions) Props() []*ConfigConferenceChatPermissions {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ConfigConferenceChatPermissions
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *ChatPermissionsProfileUsers) Props() []*ConfigConferenceChatPermissionUsers {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ConfigConferenceChatPermissionUsers
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *Advertise) GetList() map[int64]*ConfigConferenceAdvertiseRooms {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceAdvertiseRooms)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (c *CallerControlsGroups) GetList() map[int64]*ConfigConferenceCallerControlsGroups {
	c.mx.RLock()
	defer c.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceCallerControlsGroups)
	for _, v := range c.byId {
		list[v.Id] = v
	}
	return list
}

func (l *Controls) GetList() map[int64]*ConfigConferenceCallerControlsControls {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceCallerControlsControls)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (l *ConferenceProfiles) GetList() map[int64]*ConfigConferenceProfiles {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceProfiles)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (l *ConferenceProfileParams) GetList() map[int64]*ConfigConferenceProfilesParams {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceProfilesParams)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (l *ChatPermissions) GetList() map[int64]*ConfigConferenceChatPermissions {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceChatPermissions)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (l *ChatPermissionsProfileUsers) GetList() map[int64]*ConfigConferenceChatPermissionUsers {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceChatPermissionUsers)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (p *Advertise) GetById(key int64) *ConfigConferenceAdvertiseRooms {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (c *CallerControlsGroups) GetById(key int64) *ConfigConferenceCallerControlsGroups {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val := c.byId[key]
	return val
}

func (p *Controls) GetById(key int64) *ConfigConferenceCallerControlsControls {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *ConferenceProfiles) GetById(key int64) *ConfigConferenceProfiles {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *ConferenceProfileParams) GetById(key int64) *ConfigConferenceProfilesParams {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *ChatPermissions) GetById(key int64) *ConfigConferenceChatPermissions {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *ChatPermissionsProfileUsers) GetById(key int64) *ConfigConferenceChatPermissionUsers {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *Advertise) GetByName(key string) *ConfigConferenceAdvertiseRooms {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (c *CallerControlsGroups) GetByName(key string) *ConfigConferenceCallerControlsGroups {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, _ := c.byName[key]
	return val
}

func (p *Controls) GetByName(key string) *ConfigConferenceCallerControlsControls {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (p *ConferenceProfiles) GetByName(key string) *ConfigConferenceProfiles {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (p *ConferenceProfileParams) GetByName(key string) *ConfigConferenceProfilesParams {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (p *ChatPermissions) GetByName(key string) *ConfigConferenceChatPermissions {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (p *ChatPermissionsProfileUsers) GetByName(key string) *ConfigConferenceChatPermissionUsers {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (r *ConfigConferenceAdvertiseRooms) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceCallerControlsGroups) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceCallerControlsControls) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceProfiles) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceProfilesParams) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceChatPermissions) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceChatPermissionUsers) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceAdvertiseRooms) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (r *ConfigConferenceCallerControlsGroups) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (r *ConfigConferenceCallerControlsControls) GetFKTableName() string {
	return GetTableName(ConfigConferenceCallerControlsGroups{})
}

func (r *ConfigConferenceProfiles) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (r *ConfigConferenceProfilesParams) GetFKTableName() string {
	return GetTableName(ConfigConferenceProfiles{})
}

func (r *ConfigConferenceChatPermissions) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (r *ConfigConferenceChatPermissionUsers) GetFKTableName() string {
	return GetTableName(ConfigConferenceChatPermissions{})
}

func (a *Advertise) Remove(key *ConfigConferenceAdvertiseRooms) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	delete(a.byName, key.Name)
	delete(a.byId, key.Id)
}

func (c *CallerControlsGroups) Remove(key *ConfigConferenceCallerControlsGroups) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	delete(c.byName, key.Name)
	delete(c.byId, key.Id)
}

func (c *Controls) Remove(key *ConfigConferenceCallerControlsControls) {
	native := c.Parent.Controls
	native.mx.RLock()
	defer native.mx.RUnlock()
	delete(native.byName, key.Action)
	delete(native.byId, key.Id)

	root := c.Parent.Parent.Parent.Controls
	root.mx.RLock()
	defer root.mx.RUnlock()
	delete(root.byName, key.Action)
	delete(root.byId, key.Id)
}

func (c *ConferenceProfiles) Remove(key *ConfigConferenceProfiles) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	delete(c.byName, key.Name)
	delete(c.byId, key.Id)
}

func (c *ConferenceProfileParams) Remove(key *ConfigConferenceProfilesParams) {
	native := c.Parent.Params
	native.mx.RLock()
	defer native.mx.RUnlock()
	delete(native.byName, key.Name)
	delete(native.byId, key.Id)

	root := c.Parent.Parent.Parent.ConferenceProfileParams
	root.mx.RLock()
	defer root.mx.RUnlock()
	delete(root.byName, key.Name)
	delete(root.byId, key.Id)
}

func (c *ChatPermissions) Remove(key *ConfigConferenceChatPermissions) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	delete(c.byName, key.Name)
	delete(c.byId, key.Id)
}

func (c *ChatPermissionsProfileUsers) Remove(key *ConfigConferenceChatPermissionUsers) {
	native := c.Parent.Users
	native.mx.RLock()
	defer native.mx.RUnlock()
	delete(native.byName, key.Name)
	delete(native.byId, key.Id)

	root := c.Parent.Parent.Parent.ConferenceProfileParams
	root.mx.RLock()
	defer root.mx.RUnlock()
	delete(root.byName, key.Name)
	delete(root.byId, key.Id)
}

func (r *ConfigConferenceAdvertiseRooms) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceCallerControlsGroups) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceCallerControlsControls) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceProfiles) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceProfilesParams) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceChatPermissions) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceChatPermissionUsers) Remove() {
	r.Parent.Remove(r)
}

type RowItem interface {
	Remove()
	GetId() int64
	SetEnabled(bool)
	ForUpdate([]string) interface{}
	Update([]string)
	GetTableName() string
	GetFKTableName() string
}

func (r *ConfigConferenceAdvertiseRooms) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceCallerControlsGroups) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceCallerControlsControls) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceProfiles) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceProfilesParams) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceChatPermissions) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceChatPermissionUsers) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceAdvertiseRooms) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceCallerControlsGroups) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceCallerControlsControls) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceProfiles) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceProfilesParams) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceChatPermissions) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceChatPermissionUsers) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceAdvertiseRooms) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceAdvertiseRooms{Id: r.Id, Name: e[0], Status: e[1]}
}

func (r *ConfigConferenceCallerControlsGroups) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceCallerControlsGroups{Id: r.Id, Name: e[0]}
}

func (r *ConfigConferenceCallerControlsControls) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceCallerControlsControls{Id: r.Id, Action: e[0], Digits: e[1]}
}

func (r *ConfigConferenceProfiles) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceProfiles{Id: r.Id, Name: e[0]}
}

func (r *ConfigConferenceProfilesParams) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceProfilesParams{Id: r.Id, Name: e[0], Value: e[1]}
}

func (r *ConfigConferenceChatPermissions) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceChatPermissions{Id: r.Id, Name: e[0]}
}

func (r *ConfigConferenceChatPermissionUsers) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceChatPermissionUsers{Id: r.Id, Name: e[0], Commands: e[1]}
}

func (r *ConfigConferenceAdvertiseRooms) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Name = e[0]
	r.Status = e[1]
}

func (r *ConfigConferenceCallerControlsGroups) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Name = e[0]
}

func (r *ConfigConferenceCallerControlsControls) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Action = e[0]
	r.Digits = e[1]
}

func (r *ConfigConferenceProfiles) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Name = e[0]
}

func (r *ConfigConferenceProfilesParams) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Name = e[0]
	r.Value = e[1]
}

func (r *ConfigConferenceChatPermissions) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Name = e[0]
}

func (r *ConfigConferenceChatPermissionUsers) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Name = e[0]
	r.Commands = e[1]
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("[*]?([A-Za-z]+)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("[*]?([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
