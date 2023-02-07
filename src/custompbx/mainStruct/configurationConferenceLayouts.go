package mainStruct

import (
	"encoding/xml"
	"log"
	"reflect"
	"sync"
)

type ConferenceLayouts struct {
	ConfigList
	ConferenceLayoutsGroups *ConferenceLayoutsGroups `json:"-" xml:"-"`
	XmlLayoutsGroups        interface{}              `json:"-" xml:"groups>group,omitempty"`
	Layouts                 *Layouts                 `json:"-" xml:"-"`
	XmlLayouts              interface{}              `json:"-" xml:"layouts>layout,omitempty"`
	GroupLayouts            *GroupLayouts            `json:"-" xml:"-"`
	LayoutsImages           *LayoutsImages           `json:"-" xml:"-"`
}

type ConferenceLayoutsGroups struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceLayoutsGroups
	byName map[string]*ConfigConferenceLayoutsGroups
	Parent *ConferenceLayouts
}

type ConfigConferenceLayoutsGroups struct {
	Id         int64                                  `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled    bool                                   `xml:"-" json:"enabled" customsql:"enabled"`
	Name       string                                 `xml:"name,attr" json:"name" customsql:"group_name;unique"`
	Layouts    *GroupLayouts                          `xml:"-" json:"-"`
	XMLLayouts []*ConfigConferenceLayoutsGroupLayouts `xml:"control" json:"-"`
	Parent     *ConferenceLayoutsGroups               `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type GroupLayouts struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceLayoutsGroupLayouts
	Parent *ConfigConferenceLayoutsGroups
}

type ConfigConferenceLayoutsGroupLayouts struct {
	Id      int64         `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled bool          `xml:"-" json:"enabled" customsql:"enabled"`
	Body    string        `xml:",chardata"  json:"body" customsql:"layout_body;unique"`
	Parent  *GroupLayouts `xml:"-" json:"-" customsql:"fkey:group_id;unique"`
}

type Layouts struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceLayouts
	byName map[string]*ConfigConferenceLayouts
	Parent *ConferenceLayouts
}

type ConfigConferenceLayouts struct {
	Id             int64                            `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled        bool                             `xml:"-" json:"enabled" customsql:"enabled"`
	Name           string                           `xml:"name,attr" json:"name" customsql:"layout_name;unique"`
	Auto3dPosition string                           `xml:"auto-3d-position,attr" json:"auto-3d-position" customsql:"auto_3d_position"`
	LayoutsImages  *LayoutsImages                   `xml:"-" json:"-"`
	XMLLayouts     []*ConfigConferenceLayoutsImages `xml:"layouts" json:"-"`
	Parent         *Layouts                         `xml:"-" json:"-" customsql:"fkey:conf_id;unique"`
}

type LayoutsImages struct {
	mx     sync.RWMutex
	byId   map[int64]*ConfigConferenceLayoutsImages
	Parent *ConfigConferenceLayouts
}

type ConfigConferenceLayoutsImages struct {
	Id            int64          `xml:"-" json:"id" customsql:"pkey:id"`
	Enabled       bool           `xml:"-" json:"enabled" customsql:"enabled"`
	X             string         `xml:"x,attr"  json:"x" customsql:"image_x"`
	Y             string         `xml:"y,attr"  json:"y" customsql:"image_y"`
	Scale         string         `xml:"scale,attr"  json:"scale" customsql:"image_scale"`
	Floor         string         `xml:"floor,attr"  json:"floor" customsql:"image_floor"`
	FloorOnly     string         `xml:"floor-only,attr"  json:"floor_only" customsql:"image_floor_only"`
	Hscale        string         `xml:"hscale,attr"  json:"hscale" customsql:"image_hscale"`
	Overlap       string         `xml:"overlap,attr"  json:"overlap" customsql:"image_overlap"`
	ReservationId string         `xml:"reservation_id,attr"  json:"reservation_id" customsql:"image_reservation_id"`
	Zoom          string         `xml:"zoom,attr"  json:"zoom" customsql:"image_zoom"`
	Parent        *LayoutsImages `xml:"-" json:"-" customsql:"fkey:layout_id;unique"`
}

func (c *Configurations) NewConferenceLayouts(id int64, enabled bool) {
	if c.Conference != nil {
		return
	}
	var conferenceLayouts ConferenceLayouts

	conferenceLayouts.Id = id
	conferenceLayouts.Enabled = enabled
	conferenceLayouts.ConferenceLayoutsGroups = NewConferenceLayoutsGroups(&conferenceLayouts)
	conferenceLayouts.Layouts = NewConferenceLayoutsLayouts(&conferenceLayouts)
	//parents not exist at the moment so nil
	conferenceLayouts.GroupLayouts = NewGroupLayouts(nil)
	conferenceLayouts.LayoutsImages = NewLayoutsImages(nil)

	c.ConferenceLayouts = &conferenceLayouts
}

func NewConferenceLayoutsGroups(c *ConferenceLayouts) *ConferenceLayoutsGroups {
	return &ConferenceLayoutsGroups{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceLayoutsGroups),
		byName: make(map[string]*ConfigConferenceLayoutsGroups),
	}
}

func NewConferenceLayoutsLayouts(c *ConferenceLayouts) *Layouts {
	return &Layouts{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceLayouts),
		byName: make(map[string]*ConfigConferenceLayouts),
	}
}

func NewGroupLayouts(c *ConfigConferenceLayoutsGroups) *GroupLayouts {
	return &GroupLayouts{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceLayoutsGroupLayouts),
	}
}

func NewLayoutsImages(c *ConfigConferenceLayouts) *LayoutsImages {
	return &LayoutsImages{
		Parent: c,
		byId:   make(map[int64]*ConfigConferenceLayoutsImages),
	}
}

func (a *ConferenceLayoutsGroups) NewSubItem() *ConfigConferenceLayoutsGroups {
	return &ConfigConferenceLayoutsGroups{Parent: a}
}

func (a *GroupLayouts) NewSubItem() *ConfigConferenceLayoutsGroupLayouts {
	return &ConfigConferenceLayoutsGroupLayouts{Parent: a}
}

func (a *Layouts) NewSubItem() *ConfigConferenceLayouts {
	return &ConfigConferenceLayouts{Parent: a}
}

func (a *LayoutsImages) NewSubItem() *ConfigConferenceLayoutsImages {
	return &ConfigConferenceLayoutsImages{Parent: a}
}

func (a *ConfigConferenceLayouts) NewSubItem() *LayoutsImages {
	return &LayoutsImages{Parent: a}
}

func (r *ConfigConferenceLayoutsGroups) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceLayoutsGroupLayouts) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceLayouts) GetTableName() string {
	return GetTableName(r)
}

func (r *ConfigConferenceLayoutsImages) GetTableName() string {
	return GetTableName(r)
}

func (c *ConferenceLayoutsGroups) Set(value *ConfigConferenceLayoutsGroups) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value
}

func (c *GroupLayouts) Set(value *ConfigConferenceLayoutsGroupLayouts) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value

	root := c.Parent.Parent.Parent.GroupLayouts
	root.mx.RLock()
	defer root.mx.RUnlock()
	root.byId[value.Id] = value
}

func (c *Layouts) Set(value *ConfigConferenceLayouts) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
	c.byName[value.Name] = value
}

func (c *LayoutsImages) Set(value *ConfigConferenceLayoutsImages) {
	value.Parent = c
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value

	root := c.Parent.Parent.Parent.LayoutsImages
	root.mx.RLock()
	defer root.mx.RUnlock()
	root.byId[value.Id] = value
}

func (r *ConfigConferenceLayoutsGroups) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceLayoutsGroupLayouts) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceLayouts) Remove() {
	r.Parent.Remove(r)
}

func (r *ConfigConferenceLayoutsImages) Remove() {
	r.Parent.Remove(r)
}

func (c *ConferenceLayoutsGroups) Remove(key *ConfigConferenceLayoutsGroups) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	delete(c.byName, key.Name)
	delete(c.byId, key.Id)
}

func (c *GroupLayouts) Remove(key *ConfigConferenceLayoutsGroupLayouts) {
	native := c.Parent.Layouts
	native.mx.RLock()
	defer native.mx.RUnlock()
	delete(native.byId, key.Id)

	root := c.Parent.Parent.Parent.GroupLayouts
	root.mx.RLock()
	defer root.mx.RUnlock()
	delete(root.byId, key.Id)
}

func (c *Layouts) Remove(key *ConfigConferenceLayouts) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	delete(c.byName, key.Name)
	delete(c.byId, key.Id)
}

func (c *LayoutsImages) Remove(key *ConfigConferenceLayoutsImages) {
	native := c.Parent.LayoutsImages
	native.mx.RLock()
	defer native.mx.RUnlock()
	delete(native.byId, key.Id)

	root := c.Parent.Parent.Parent.LayoutsImages
	root.mx.RLock()
	defer root.mx.RUnlock()
	delete(root.byId, key.Id)
}

func (r *ConfigConferenceLayoutsGroups) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceLayoutsGroupLayouts) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceLayouts) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceLayoutsImages) GetId() int64 {
	return r.Id
}

func (r *ConfigConferenceLayoutsGroups) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceLayoutsGroupLayouts) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceLayouts) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceLayoutsImages) SetEnabled(e bool) {
	r.Enabled = e
}

func (r *ConfigConferenceLayoutsGroups) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceLayoutsGroups{Id: r.Id, Name: e[0]}
}

func (r *ConfigConferenceLayoutsGroupLayouts) ForUpdate(e []string) interface{} {
	if len(e) != 1 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceLayoutsGroupLayouts{Id: r.Id, Body: e[0]}
}

func (r *ConfigConferenceLayouts) ForUpdate(e []string) interface{} {
	if len(e) != 2 || e[0] == "" {
		return nil
	}
	return &ConfigConferenceLayouts{Id: r.Id, Name: e[0], Auto3dPosition: e[1]}
}

func (r *ConfigConferenceLayoutsImages) ForUpdate(e []string) interface{} {
	if len(e) != 9 {
		return nil
	}
	return &ConfigConferenceLayoutsImages{Id: r.Id, X: e[0], Y: e[1], Scale: e[2], Floor: e[3], FloorOnly: e[4], Hscale: e[5], Overlap: e[6], ReservationId: e[7], Zoom: e[8]}
}

func (r *ConfigConferenceLayoutsGroups) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (r *ConfigConferenceLayoutsGroupLayouts) GetFKTableName() string {
	return GetTableName(ConfigConferenceLayoutsGroups{})
}

func (r *ConfigConferenceLayouts) GetFKTableName() string {
	return GetTableName(ConfigList{})
}

func (r *ConfigConferenceLayoutsImages) GetFKTableName() string {
	return GetTableName(ConfigConferenceLayouts{})
}

func (r *ConfigConferenceLayoutsGroups) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Name = e[0]
}

func (r *ConfigConferenceLayoutsGroupLayouts) Update(e []string) {
	if len(e) != 1 || e[0] == "" {
		return
	}
	r.Body = e[0]
}

func (r *ConfigConferenceLayouts) Update(e []string) {
	if len(e) != 2 || e[0] == "" {
		return
	}
	r.Name = e[0]
	r.Auto3dPosition = e[1]
}

func (r *ConfigConferenceLayoutsImages) Update(e []string) {
	if len(e) != 9 {
		return
	}
	r.X = e[0]
	r.Y = e[1]
	r.Scale = e[2]
	r.Floor = e[3]
	r.FloorOnly = e[4]
	r.Hscale = e[5]
	r.Overlap = e[6]
	r.ReservationId = e[7]
	r.Zoom = e[8]
}

func (c *ConferenceLayoutsGroups) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceLayoutsGroups:
		t := value.(*ConfigConferenceLayoutsGroups)
		t.Layouts = NewGroupLayouts(t)
		c.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (c *GroupLayouts) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceLayoutsGroupLayouts:
		t := value.(*ConfigConferenceLayoutsGroupLayouts)
		c.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (c *Layouts) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceLayouts:
		t := value.(*ConfigConferenceLayouts)
		t.LayoutsImages = NewLayoutsImages(t)
		c.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (c *LayoutsImages) SetFromInterface(value interface{}) {
	switch value.(type) {
	case *ConfigConferenceLayoutsImages:
		t := value.(*ConfigConferenceLayoutsImages)
		c.Set(t)
	default:
		log.Println(reflect.TypeOf(value).Name())
	}
}

func (c *ConferenceLayoutsGroups) NewSubItemInterface() RowItem {
	return &ConfigConferenceLayoutsGroups{Parent: c}
}

func (c *GroupLayouts) NewSubItemInterface() RowItem {
	return &ConfigConferenceLayoutsGroupLayouts{Parent: c}
}

func (c *Layouts) NewSubItemInterface() RowItem {
	return &ConfigConferenceLayouts{Parent: c}
}

func (c *LayoutsImages) NewSubItemInterface() RowItem {
	return &ConfigConferenceLayoutsImages{Parent: c}
}

func (c *ConferenceLayoutsGroups) Props() []*ConfigConferenceLayoutsGroups {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*ConfigConferenceLayoutsGroups
	for _, val := range c.byId {
		items = append(items, val)
	}
	return items
}

func (c *GroupLayouts) Props() []*ConfigConferenceLayoutsGroupLayouts {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*ConfigConferenceLayoutsGroupLayouts
	for _, val := range c.byId {
		items = append(items, val)
	}
	return items
}

func (c *Layouts) Props() []*ConfigConferenceLayouts {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*ConfigConferenceLayouts
	for _, val := range c.byId {
		items = append(items, val)
	}
	return items
}

func (c *LayoutsImages) Props() []*ConfigConferenceLayoutsImages {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*ConfigConferenceLayoutsImages
	for _, val := range c.byId {
		items = append(items, val)
	}
	return items
}

func (p *ConferenceLayoutsGroups) GetById(key int64) *ConfigConferenceLayoutsGroups {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *GroupLayouts) GetById(key int64) *ConfigConferenceLayoutsGroupLayouts {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *Layouts) GetById(key int64) *ConfigConferenceLayouts {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *LayoutsImages) GetById(key int64) *ConfigConferenceLayoutsImages {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byId[key]
	return val
}

func (p *ConferenceLayoutsGroups) GetByName(key string) *ConfigConferenceLayoutsGroups {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byName[key]
	return val
}

func (p *Layouts) GetByName(key string) *ConfigConferenceLayouts {
	p.mx.RLock()
	defer p.mx.RUnlock()
	val := p.byName[key]
	return val
}

func (c *ConferenceLayoutsGroups) GetList() map[int64]*ConfigConferenceLayoutsGroups {
	c.mx.RLock()
	defer c.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceLayoutsGroups)
	for _, v := range c.byId {
		list[v.Id] = v
	}
	return list
}

func (c *GroupLayouts) GetList() map[int64]*ConfigConferenceLayoutsGroupLayouts {
	c.mx.RLock()
	defer c.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceLayoutsGroupLayouts)
	for _, v := range c.byId {
		list[v.Id] = v
	}
	return list
}

func (c *Layouts) GetList() map[int64]*ConfigConferenceLayouts {
	c.mx.RLock()
	defer c.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceLayouts)
	for _, v := range c.byId {
		list[v.Id] = v
	}
	return list
}

func (c *LayoutsImages) GetList() map[int64]*ConfigConferenceLayoutsImages {
	c.mx.RLock()
	defer c.mx.RUnlock()
	list := make(map[int64]*ConfigConferenceLayoutsImages)
	for _, v := range c.byId {
		list[v.Id] = v
	}
	return list
}

func (c *Configurations) XMLConferenceLayouts() *Configuration {
	if c.ConferenceLayouts == nil || !c.ConferenceLayouts.Enabled {
		return nil
	}
	c.ConferenceLayouts.XMLItems()
	currentConfig := Configuration{
		Name:        ConfConference,
		Description: "Conference Config",
		AnyXML: struct {
			LayoutSettings []interface{} `xml:"layout-settings"`
		}{
			LayoutSettings: []interface{}{
				struct {
					XMLName xml.Name    `xml:"layouts,omitempty"`
					Inner   interface{} `xml:"layout"`
				}{Inner: c.ConferenceLayouts.XmlLayouts},
				struct {
					XMLName xml.Name    `xml:"groups,omitempty"`
					Inner   interface{} `xml:"group"`
				}{Inner: c.ConferenceLayouts.XmlLayoutsGroups},
			},
		},
	}
	return &currentConfig
}

func (v *ConferenceLayouts) XMLItems() {
	v.XmlLayouts = v.Layouts.XMLItems()
	v.XmlLayoutsGroups = v.ConferenceLayoutsGroups.XMLItems()
}

func (v *Layouts) XMLItems() []interface{} {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var profile []interface{}
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		val.XMLLayouts = val.LayoutsImages.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (v *LayoutsImages) XMLItems() []*ConfigConferenceLayoutsImages {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []*ConfigConferenceLayoutsImages
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, val)
	}
	return param
}

func (v *ConferenceLayoutsGroups) XMLItems() []interface{} {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var profile []interface{}
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		val.XMLLayouts = val.Layouts.XMLItems()
		profile = append(profile, *val)
	}
	return profile
}

func (v *GroupLayouts) XMLItems() []*ConfigConferenceLayoutsGroupLayouts {
	v.mx.RLock()
	defer v.mx.RUnlock()
	var param []*ConfigConferenceLayoutsGroupLayouts
	for _, val := range v.byId {
		if !val.Enabled {
			continue
		}
		param = append(param, val)
	}
	return param
}
