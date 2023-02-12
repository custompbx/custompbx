package mainStruct

import (
	"encoding/xml"
	"log"
	"reflect"
	"sync"
)

type PostLoadModules struct {
	ConfigList
	Loaded     bool        `xml:"-" json:"loaded"`
	Modules    *ModuleTags `json:"-" xml:"-"`
	XmlModules []ModuleTag `json:"-" xml:"modules>load,omitempty"`
	Unloadable bool        `xml:"-" json:"unloadable"`
}

type ModuleTags struct {
	mx     sync.RWMutex
	byName map[string]*ModuleTag
	byId   map[int64]*ModuleTag
	Parent *PostLoadModules
}

type ModuleTag struct {
	Id      int64       `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled bool        `xml:"-" json:"enabled" customsql:"enabled"`
	Name    string      `xml:"module,attr" json:"name,omitempty"  customsql:"module_name;unique"`
	Parent  *ModuleTags `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

func (p *PostLoadModules) Reload() string {
	return ""
}
func (p *PostLoadModules) Unload() string {
	return ""
}
func (p *PostLoadModules) Load() string {
	return ""
}
func (p *PostLoadModules) Switch(enabled bool) {
	p.Enabled = enabled
}
func (p *PostLoadModules) AutoLoad() {

}
func (p *PostLoadModules) GetId() int64 {
	return p.Id
}
func (p *PostLoadModules) SetLoadStatus(status bool) {
	p.Loaded = status
}
func (p *PostLoadModules) GetConfig() *Configurations {
	return &Configurations{PostLoadModules: p}
}
func (p *PostLoadModules) GetModuleName() string {
	return ModPostLoadModules
}
func (p *PostLoadModules) IsNil() bool {
	return p == nil
}

func (p *PostLoadModules) XMLItems() {
	p.XmlModules = p.Modules.XMLItems()
}

func (c *Configurations) NewPostLoadModules(id int64, enabled bool) {
	if c.PostLoadModules != nil {
		return
	}
	var postLoadModules PostLoadModules
	postLoadModules.Id = id
	postLoadModules.Enabled = enabled
	postLoadModules.Unloadable = true
	postLoadModules.Modules = NewModuleTags(&postLoadModules)

	c.PostLoadModules = &postLoadModules
}

func (c *Configurations) XMLPostLoadModules() *Configuration {
	if c.PostLoadModules == nil || !c.PostLoadModules.Enabled {
		return nil
	}
	c.PostLoadModules.XMLItems()
	currentConfig := Configuration{Name: ConfPostLoadModules, Description: "PostLoadModules Config", AnyXML: struct {
		XMLName xml.Name    `xml:"modules,omitempty"`
		Inner   interface{} `xml:"load"`
	}{Inner: &c.PostLoadModules.XmlModules}}
	return &currentConfig
}

func NewModuleTags(c *PostLoadModules) *ModuleTags {
	return &ModuleTags{
		byName: make(map[string]*ModuleTag),
		byId:   make(map[int64]*ModuleTag),
		Parent: c,
	}
}

func (m *ModuleTags) XMLItems() []ModuleTag {
	m.mx.RLock()
	defer m.mx.RUnlock()
	var list []ModuleTag
	for _, v := range m.byName {
		if !v.Enabled {
			continue
		}
		list = append(list, *v)
	}
	return list
}

func (a *ModuleTags) NewSubItem() *ModuleTag {
	return &ModuleTag{Parent: a}
}

func (a *ModuleTags) Set(value *ModuleTag) {
	value.Parent = a
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
	a.byName[value.Name] = value
}

func (a *ModuleTags) NewSubItemInterface() RowItem {
	return &ModuleTag{Parent: a}
}

func (a *ModuleTags) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ModuleTag:
		a.Set(value.(*ModuleTag))
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (l *ModuleTags) Props() []*ModuleTag {
	l.mx.RLock()
	defer l.mx.RUnlock()
	var items []*ModuleTag
	for _, val := range l.byId {
		items = append(items, val)
	}
	return items
}

func (l *ModuleTags) GetList() map[int64]*ModuleTag {
	l.mx.RLock()
	defer l.mx.RUnlock()
	list := make(map[int64]*ModuleTag)
	for _, v := range l.byId {
		list[v.Id] = v
	}
	return list
}

func (p *ModuleTags) GetById(key int64) *ModuleTag {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *ModuleTags) GetByName(key string) *ModuleTag {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val, _ := p.byName[key]
	return val
}

func (r *ModuleTag) GetTableName() string {
	return "config_post_load_modules"
}

func (a *ModuleTags) Remove(key *ModuleTag) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	delete(a.byName, key.Name)
	delete(a.byId, key.Id)
}

func (r *ModuleTag) Remove() {
	r.Parent.Remove(r)
}

func (r *ModuleTag) GetId() int64 {
	return r.Id
}

func (r *ModuleTag) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ModuleTag) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &ModuleTag{Id: r.Id, Name: e[0]}
}

func (r *ModuleTag) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Name = e[0]
}

func (r *ModuleTag) GetFKTableName() string {
	return GetTableName(ConfigList{})
}
