package mainStruct

import (
	"encoding/xml"
	"log"
	"reflect"
	"sync"
)

type Voicemail struct {
	ConfigList
	Loaded        bool                         `xml:"-" json:"loaded"`
	Settings      *VoicemailSettings           `json:"-" xml:"-"`
	XMLSettings   []VoicemailSettingsParameter `json:"-" xml:"param,omitempty"`
	Profiles      *VoicemailProfiles           `json:"-" xml:"-"`
	XMLProfiles   []VoicemailProfile           `json:"-" xml:"profiles>profile,omitempty"`
	ProfileParams *VoicemailProfilesParameters `json:"-" xml:"-"`
}

type VoicemailSettings struct {
	mx     sync.RWMutex
	byName map[string]*VoicemailSettingsParameter
	byId   map[int64]*VoicemailSettingsParameter
	Parent *Voicemail
}

type VoicemailSettingsParameter struct {
	Id      int64              `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled bool               `xml:"-" json:"enabled" customsql:"enabled"`
	Name    string             `xml:"name,attr" json:"name,omitempty"  customsql:"param_name;unique"`
	Value   string             `xml:"value,attr" json:"value,omitempty"  customsql:"param_value"`
	Parent  *VoicemailSettings `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type VoicemailProfiles struct {
	mx     sync.RWMutex
	byName map[string]*VoicemailProfile
	byId   map[int64]*VoicemailProfile
	Parent *Voicemail
}

type VoicemailProfile struct {
	Id        int64                        `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled   bool                         `json:"enabled" xml:"-" customsql:"enabled"`
	Name      string                       `json:"name" xml:"name,attr" customsql:"param_name;unique"`
	Params    *VoicemailProfilesParameters `json:"-" xml:"-"`
	XmlParams []VoicemailProfilesParameter `json:"-" xml:"param,omitempty"`
	Parent    *VoicemailProfiles           `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type VoicemailProfilesParameters struct {
	mx     sync.RWMutex
	byName map[string]*VoicemailProfilesParameter
	byId   map[int64]*VoicemailProfilesParameter
	Parent *VoicemailProfile
}

type VoicemailProfilesParameter struct {
	Id      int64                        `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled bool                         `xml:"-" json:"enabled" customsql:"enabled"`
	Name    string                       `xml:"name,attr" json:"name,omitempty"  customsql:"param_name;unique"`
	Value   string                       `xml:"value,attr" json:"value,omitempty"  customsql:"param_value;unique"`
	Parent  *VoicemailProfilesParameters `xml:"-" json:"-" customsql:"fkey:profile_id;unique"`
}

func (p *Voicemail) Reload() string {
	return "reload " + p.GetModuleName()
}
func (p *Voicemail) Unload() string {
	return "unload " + p.GetModuleName()
}
func (p *Voicemail) Load() string {
	return "load " + p.GetModuleName()
}
func (p *Voicemail) Switch(enabled bool) {
	p.Enabled = enabled
}
func (p *Voicemail) AutoLoad() {

}
func (p *Voicemail) GetId() int64 {
	return p.Id
}
func (p *Voicemail) SetLoadStatus(status bool) {
	p.Loaded = status
}
func (p *Voicemail) GetConfig() *Configurations {
	return &Configurations{Voicemail: p}
}
func (p *Voicemail) GetModuleName() string {
	return ModVoicemail
}
func (p *Voicemail) IsNil() bool {
	return p == nil
}

func (p *Voicemail) XMLItems() {
	p.XMLSettings = p.Settings.XMLItems()
	p.XMLProfiles = p.Profiles.XMLItems()
}

func (c *Configurations) NewVoicemail(id int64, enabled bool) {
	if c.Voicemail != nil {
		return
	}
	var voicemail Voicemail
	voicemail.Id = id
	voicemail.Enabled = enabled
	voicemail.Settings = NewVoicemailSettings(&voicemail)
	voicemail.Profiles = NewVoicemailProfiles(&voicemail)
	voicemail.ProfileParams = NewVoicemailProfileParams(nil)

	c.Voicemail = &voicemail
}

func (c *Configurations) XMLVoicemail() *Configuration {
	if c.Voicemail == nil || !c.Voicemail.Enabled {
		return nil
	}
	c.Voicemail.XMLItems()
	currentConfig := Configuration{Name: ConfVoicemail, Description: "Voicemail Config", AnyXML: []interface{}{
		struct {
			XMLName xml.Name    `xml:"settings,omitempty"`
			Inner   interface{} `xml:"param"`
		}{Inner: &c.Voicemail.XMLSettings},
		struct {
			XMLName xml.Name    `xml:"profiles,omitempty"`
			Inner   interface{} `xml:"param"`
		}{Inner: &c.Voicemail.XMLProfiles},
	},
	}
	return &currentConfig
}

func NewVoicemailSettings(c *Voicemail) *VoicemailSettings {
	return &VoicemailSettings{
		byName: make(map[string]*VoicemailSettingsParameter),
		byId:   make(map[int64]*VoicemailSettingsParameter),
		Parent: c,
	}
}

func NewVoicemailProfiles(c *Voicemail) *VoicemailProfiles {
	return &VoicemailProfiles{
		byName: make(map[string]*VoicemailProfile),
		byId:   make(map[int64]*VoicemailProfile),
		Parent: c,
	}
}

func NewVoicemailProfileParams(c *VoicemailProfile) *VoicemailProfilesParameters {
	return &VoicemailProfilesParameters{
		Parent: c,
		byId:   make(map[int64]*VoicemailProfilesParameter),
		byName: make(map[string]*VoicemailProfilesParameter),
	}
}

func (m *VoicemailSettings) XMLItems() []VoicemailSettingsParameter {
	m.mx.RLock()
	defer m.mx.RUnlock()
	var list []VoicemailSettingsParameter
	for _, v := range m.byName {
		if !v.Enabled {
			continue
		}
		list = append(list, *v)
	}
	return list
}

func (a *VoicemailSettings) NewSubItem() *VoicemailSettingsParameter {
	return &VoicemailSettingsParameter{Parent: a}
}

func (a *VoicemailSettings) Set(value *VoicemailSettingsParameter) {
	value.Parent = a
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
	a.byName[value.Name] = value
}

func (a *VoicemailSettings) NewSubItemInterface() RowItem {
	return &VoicemailSettingsParameter{Parent: a}
}

func (a *VoicemailSettings) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *VoicemailSettingsParameter:
		a.Set(value.(*VoicemailSettingsParameter))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (l *VoicemailSettings) Props() []*VoicemailSettingsParameter {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*VoicemailSettingsParameter
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *VoicemailSettings) GetList() map[int64]*VoicemailSettingsParameter {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*VoicemailSettingsParameter)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (p *VoicemailSettings) GetById(key int64) *VoicemailSettingsParameter {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *VoicemailSettings) GetByName(key string) *VoicemailSettingsParameter {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (a *VoicemailSettings) Remove(key *VoicemailSettingsParameter) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	delete(a.byName, key.Name)
	delete(a.byId, key.Id)
}

func (r *VoicemailSettingsParameter) GetTableName() string {
	return "config_voicemail_settings"
}

func (r *VoicemailSettingsParameter) Remove() {
	r.Parent.Remove(r)
}

func (r *VoicemailSettingsParameter) GetId() int64 {
	return r.Id
}

func (r *VoicemailSettingsParameter) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *VoicemailSettingsParameter) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &VoicemailSettingsParameter{Id: r.Id, Name: e[0], Value: e[1]}
}

func (r *VoicemailSettingsParameter) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Name = e[0]
	r.Value = e[1]
}

func (r *VoicemailSettingsParameter) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (m *VoicemailProfilesParameters) XMLItems() []VoicemailProfilesParameter {
	m.mx.RLock()
	defer m.mx.RUnlock()
	var list []VoicemailProfilesParameter
	for _, v := range m.byName {
		if !v.Enabled {
			continue
		}
		list = append(list, *v)
	}
	return list
}

func (a *VoicemailProfilesParameters) NewSubItem() *VoicemailProfilesParameter {
	return &VoicemailProfilesParameter{Parent: a}
}

func (a *VoicemailProfilesParameters) Set(value *VoicemailProfilesParameter) {
	value.Parent = a
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
	a.byName[value.Name] = value

	root := a.Parent.Parent.Parent.ProfileParams
	root.mx.RLock()
	defer root.mx.RUnlock()
	root.byId[value.Id] = value
	root.byName[value.Name] = value
}

func (a *VoicemailProfilesParameters) NewSubItemInterface() RowItem {
	return &VoicemailProfilesParameter{Parent: a}
}

func (a *VoicemailProfilesParameters) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *VoicemailProfilesParameter:
		a.Set(value.(*VoicemailProfilesParameter))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (l *VoicemailProfilesParameters) Props() []*VoicemailProfilesParameter {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*VoicemailProfilesParameter
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *VoicemailProfilesParameters) GetList() map[int64]*VoicemailProfilesParameter {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*VoicemailProfilesParameter)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (p *VoicemailProfilesParameters) GetById(key int64) *VoicemailProfilesParameter {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *VoicemailProfilesParameters) GetByName(key string) *VoicemailProfilesParameter {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (a *VoicemailProfilesParameters) Remove(key *VoicemailProfilesParameter) {
	native := a.Parent.Params
	native.mx.RLock()
	defer native.mx.RUnlock()
	delete(native.byName, key.Name)
	delete(native.byId, key.Id)

	root := a.Parent.Parent.Parent.ProfileParams
	root.mx.RLock()
	defer root.mx.RUnlock()
	delete(root.byName, key.Name)
	delete(root.byId, key.Id)
}

func (r *VoicemailProfile) GetTableName() string {
	return "config_voicemail_profiles"
}

func (r *VoicemailProfile) Remove() {
	r.Parent.Remove(r)
}

func (r *VoicemailProfile) GetId() int64 {
	return r.Id
}

func (r *VoicemailProfile) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *VoicemailProfile) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &VoicemailProfile{Id: r.Id, Name: e[0]}
}

func (r *VoicemailProfile) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Name = e[0]
}

func (r *VoicemailProfile) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (m *VoicemailProfiles) XMLItems() []VoicemailProfile {
	m.mx.RLock()
	defer m.mx.RUnlock()
	var list []VoicemailProfile
	for _, v := range m.byName {
		if !v.Enabled {
			continue
		}
		list = append(list, *v)
	}
	return list
}

func (a *VoicemailProfiles) NewSubItem() *VoicemailProfile {
	return &VoicemailProfile{Parent: a}
}

func (a *VoicemailProfiles) Set(value *VoicemailProfile) {
	value.Parent = a
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
	a.byName[value.Name] = value
}

func (a *VoicemailProfiles) NewSubItemInterface() RowItem {
	return &VoicemailProfile{Parent: a}
}

func (a *VoicemailProfiles) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *VoicemailProfile:
		t := value.(*VoicemailProfile)
		t.Params = NewVoicemailProfileParams(t)
		a.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (l *VoicemailProfiles) Props() []*VoicemailProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*VoicemailProfile
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *VoicemailProfiles) GetList() map[int64]*VoicemailProfile {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*VoicemailProfile)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (p *VoicemailProfiles) GetById(key int64) *VoicemailProfile {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *VoicemailProfiles) GetByName(key string) *VoicemailProfile {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (a *VoicemailProfiles) Remove(key *VoicemailProfile) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	delete(a.byName, key.Name)
	delete(a.byId, key.Id)
}

func (r *VoicemailProfilesParameter) GetTableName() string {
	return "config_voicemail_profiles_parameters"
}

func (r *VoicemailProfilesParameter) Remove() {
	r.Parent.Remove(r)
}

func (r *VoicemailProfilesParameter) GetId() int64 {
	return r.Id
}

func (r *VoicemailProfilesParameter) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *VoicemailProfilesParameter) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &VoicemailProfilesParameter{Id: r.Id, Name: e[0], Value: e[1]}
}

func (r *VoicemailProfilesParameter) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Name = e[0]
	r.Value = e[1]
}

func (r *VoicemailProfilesParameter) GetFKTableName() string {
	return (&VoicemailProfile{}).GetTableName()
}
