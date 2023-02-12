package mainStruct

import (
	"custompbx/xmlStruct"
	"encoding/xml"
	"log"
	"sort"
	"sync"
)

// Dialplan

type Dialplans struct {
	Contexts    *Contexts    `json:"-"`
	Extensions  *Extensions  `json:"-"`
	Conditions  *Conditions  `json:"-"`
	Regexes     *Regexes     `json:"-"`
	Actions     *Actions     `json:"-"`
	AntiActions *AntiActions `json:"-"`
	NoProceed   bool         `json:"no_proceed"`
	EnableDebug bool         `json:"enable_debug"`
}

type Contexts struct {
	mx     sync.RWMutex
	byName map[string]*Context
	byId   map[int64]*Context
}

type Extensions struct {
	mx     sync.RWMutex
	byName map[string]*Extension
	byId   map[int64]*Extension
}

type Conditions struct {
	mx   sync.RWMutex
	byId map[int64]*Condition
}

type Regexes struct {
	mx   sync.RWMutex
	byId map[int64]*Regex
}

type Actions struct {
	mx   sync.RWMutex
	byId map[int64]*Action
}

type AntiActions struct {
	mx   sync.RWMutex
	byId map[int64]*AntiAction
}

type DialplanDebug struct {
	Log     []string `json:"log"`
	Actions []Action `json:"actions"`
}

type ConcurrentByteSlice struct {
	mx    sync.RWMutex
	bytes []byte
}

func (c *ConcurrentByteSlice) Set(data []byte) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.bytes = data
}

func (c *ConcurrentByteSlice) Get() []byte {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.bytes
}

type Context struct {
	Id           int64               `xml:"-" json:"id"`
	Enabled      bool                `xml:"-" json:"enabled"`
	Name         string              `xml:"name,attr" json:"name"`
	Extensions   *Extensions         `xml:"-" json:"-"`
	XMLExtension Extension           `xml:"extension" json:"-"`
	FullXMLCache ConcurrentByteSlice `xml:"-" json:"-"`
	Dialplan     *Dialplans          `xml:"-" json:"-"`
}

type Extension struct {
	Id           int64       `xml:"-" json:"id"`
	Enabled      bool        `xml:"-" json:"enabled"`
	Position     int64       `xml:"-" json:"position"`
	Name         string      `xml:"name,attr" json:"name"`
	Continue     string      `xml:"continue,attr,omitempty" json:"continue"`
	Conditions   *Conditions `xml:"-" json:"-"`
	XMLCondition Condition   `xml:"condition" json:"-"`
	Context      *Context    `xml:"-" json:"-"`
}

type Condition struct {
	Id            int64        `xml:"-" json:"id"`
	Enabled       bool         `xml:"-" json:"enabled"`
	Position      int64        `xml:"-" json:"position"`
	Break         string       `xml:"break,attr,omitempty" json:"break"`
	Field         string       `xml:"field,attr,omitempty" json:"field"`
	Expression    string       `xml:"expression,attr,omitempty" json:"expression"`
	Actions       *Actions     `xml:"-" json:"-"`
	XMLAction     []Action     `xml:"action,omitempty" json:"-"`
	AntiActions   *AntiActions `xml:"-" json:"-"`
	XMLAntiAction []AntiAction `xml:"antiaction,omitempty" json:"-"`
	Extension     *Extension   `xml:"-" json:"-"`
	Hour          string       `xml:"hour,attr,omitempty" json:"hour"`
	Mday          string       `xml:"mday,attr,omitempty" json:"mday"`
	Mon           string       `xml:"mon,attr,omitempty" json:"mon"`
	Mweek         string       `xml:"mweek,attr,omitempty" json:"mweek"`
	Wday          string       `xml:"wday,attr,omitempty" json:"wday"`
	DateTime      string       `xml:"date-time,attr,omitempty" json:"date_time"`
	TimeOfDay     string       `xml:"time-of-day,attr,omitempty" json:"time_of_day"`
	Year          string       `xml:"year,attr,omitempty" json:"year"`
	Minute        string       `xml:"minute,attr,omitempty" json:"minute"`
	Week          string       `xml:"week,attr,omitempty" json:"week"`
	Yday          string       `xml:"yday,attr,omitempty" json:"yday"`
	Minday        string       `xml:"minday,attr,omitempty" json:"minday"`
	TzOffset      string       `xml:"tz-offset,attr,omitempty" json:"tz_offset"`
	Dst           string       `xml:"dst,attr,omitempty" json:"dst"`
	Regex         string       `xml:"regex,attr,omitempty" json:"regex"`
	Regexes       *Regexes     `xml:"-" json:"-"`
}

type Regex struct {
	Id         int64      `xml:"-" json:"id"`
	Enabled    bool       `xml:"-" json:"enabled"`
	Field      string     `xml:"field,attr,omitempty" json:"field"`
	Expression string     `xml:"expression,attr,omitempty" json:"expression"`
	Condition  *Condition `xml:"-" json:"-"`
}

type Action struct {
	Id          int64  `xml:"-" json:"id"`
	Enabled     bool   `xml:"-" json:"enabled"`
	Position    int64  `xml:"-" json:"position"`
	Application string `xml:"application,attr" json:"application"`
	Data        string `xml:"data,attr" json:"data"`
	Inline      bool   `xml:"inline,attr,omitempty" json:"inline"`
	/*	Function    string `xml:"function,attr" json:"function"`
		Method      string `xml:"method,attr" json:"method"`
		Phrase      string `xml:"phrase,attr" json:"phrase"`
		ActionType  string `xml:"type,attr" json:"type"`*/
	Condition *Condition `xml:"-" json:"-"`
}

type AntiAction struct {
	Id          int64      `xml:"-" json:"id"`
	Enabled     bool       `xml:"-" json:"enabled"`
	Position    int64      `xml:"-" json:"position"`
	Application string     `xml:"application,attr" json:"application"`
	Data        string     `xml:"data,attr" json:"data"`
	Condition   *Condition `xml:"-" json:"-"`
}

func (c *Contexts) Set(value *Context) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byName[value.Name] = value
	c.byId[value.Id] = value
}

func (e *Extensions) Set(value *Extension) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.byName[value.Name] = value
	e.byId[value.Id] = value
}

func (c *Conditions) Set(value *Condition) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.byId[value.Id] = value
}

func (r *Regexes) Set(value *Regex) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.byId[value.Id] = value
}

func (a *Actions) Set(value *Action) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
}

func (a *AntiActions) Set(value *AntiAction) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.byId[value.Id] = value
}

func (c *Contexts) GetById(key int64) (*Context, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.byId[key]
	return val, ok
}

func (e *Extensions) GetById(key int64) (*Extension, bool) {
	e.mx.RLock()
	defer e.mx.RUnlock()
	val, ok := e.byId[key]
	return val, ok
}

func (c *Conditions) GetById(key int64) (*Condition, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.byId[key]
	return val, ok
}

func (r *Regexes) GetById(key int64) (*Regex, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	val, ok := r.byId[key]
	return val, ok
}

func (a *Actions) GetById(key int64) (*Action, bool) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	val, ok := a.byId[key]
	return val, ok
}

func (a *AntiActions) GetById(key int64) (*AntiAction, bool) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	val, ok := a.byId[key]
	return val, ok
}

func NewDialplanItems() *Dialplans {
	return &Dialplans{
		Contexts:    NewContexts(),
		Extensions:  NewExtensions(),
		Conditions:  NewConditions(),
		Regexes:     NewRegexes(),
		Actions:     NewActions(),
		AntiActions: NewAntiActions(),
	}
}

func NewContexts() *Contexts {
	return &Contexts{
		byName: make(map[string]*Context),
		byId:   make(map[int64]*Context),
	}
}

func NewExtensions() *Extensions {
	return &Extensions{
		byName: make(map[string]*Extension),
		byId:   make(map[int64]*Extension),
	}
}

func NewConditions() *Conditions {
	return &Conditions{
		byId: make(map[int64]*Condition),
	}
}

func NewRegexes() *Regexes {
	return &Regexes{
		byId: make(map[int64]*Regex),
	}
}

func NewActions() *Actions {
	return &Actions{
		byId: make(map[int64]*Action),
	}
}

func NewAntiActions() *AntiActions {
	return &AntiActions{
		byId: make(map[int64]*AntiAction),
	}
}

func (e *Extensions) GetByName(key string) *Extension {
	e.mx.RLock()
	defer e.mx.RUnlock()
	val, _ := e.byName[key]

	return val
}

func (c *Contexts) Names() []string {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var keys []string
	for k := range c.byName {
		keys = append(keys, k)
	}
	return keys
}

func (c *Contexts) Props() []*Context {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*Context
	for _, v := range c.byId {
		items = append(items, v)
	}
	return items
}

func (c *Contexts) GetByName(key string) (*Context, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.byName[key]
	return val, ok
}

func (e *Extensions) Names() []string {
	e.mx.RLock()
	defer e.mx.RUnlock()
	var keys []string
	for k := range e.byName {
		keys = append(keys, k)
	}
	return keys
}

func (e *Extensions) Props() []*Extension {
	e.mx.RLock()
	defer e.mx.RUnlock()
	var items []*Extension
	for _, v := range e.byId {
		items = append(items, v)
	}
	return items
}

func (c *Conditions) Ids() []int64 {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var keys []int64
	for k := range c.byId {
		keys = append(keys, k)
	}
	return keys
}

func (c *Conditions) Props() []*Condition {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var items []*Condition
	for _, v := range c.byId {
		items = append(items, v)
	}
	return items
}

func (a *Actions) Props() []*Action {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var items = make([]*Action, 0, len(a.byId))
	for _, v := range a.byId {
		items = append(items, v)
	}
	return items
}

func (a *AntiActions) Props() []*AntiAction {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var items = make([]*AntiAction, 0, len(a.byId))
	for _, v := range a.byId {
		items = append(items, v)
	}
	return items
}

func (r *Regexes) Props() []*Regex {
	r.mx.RLock()
	defer r.mx.RUnlock()
	var items = make([]*Regex, 0, len(r.byId))
	for _, v := range r.byId {
		items = append(items, v)
	}
	return items
}

func (e *Extensions) ExtensionsList() []*Extension {
	e.mx.RLock()
	defer e.mx.RUnlock()
	var extensions = make([]*Extension, 0, len(e.byId))
	for _, ex := range e.byId {
		if !ex.Enabled {
			continue
		}
		extensions = append(extensions, ex)
	}
	sort.SliceStable(extensions, func(i, j int) bool {
		return extensions[i].Position < extensions[j].Position
	})
	return extensions
}

func (c *Conditions) ConditionsList() []*Condition {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var conditions = make([]*Condition, 0, len(c.byId))
	for _, ex := range c.byId {
		if !ex.Enabled {
			continue
		}
		conditions = append(conditions, ex)
	}
	sort.SliceStable(conditions, func(i, j int) bool {
		return conditions[i].Position < conditions[j].Position
	})
	return conditions
}

func (r *Regexes) RegexList() []*Regex {
	r.mx.RLock()
	defer r.mx.RUnlock()
	var regexes = make([]*Regex, 0, len(r.byId))
	for _, ex := range r.byId {
		if !ex.Enabled {
			continue
		}
		regexes = append(regexes, ex)
	}
	return regexes
}

func (a *Actions) ActionsList() []*Action {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var actions = make([]*Action, 0, len(a.byId))
	for _, ex := range a.byId {
		if !ex.Enabled {
			continue
		}
		actions = append(actions, ex)
	}
	sort.SliceStable(actions, func(i, j int) bool {
		return actions[i].Position < actions[j].Position
	})
	return actions
}

func (a *AntiActions) ActionsList() []*AntiAction {
	a.mx.RLock()
	defer a.mx.RUnlock()
	var antiActions = make([]*AntiAction, 0, len(a.byId))
	for _, ex := range a.byId {
		if !ex.Enabled {
			continue
		}
		antiActions = append(antiActions, ex)
	}
	sort.SliceStable(antiActions, func(i, j int) bool {
		return antiActions[i].Position < antiActions[j].Position
	})
	return antiActions
}

func (c *Contexts) GetList() map[int64]*Context {
	c.mx.RLock()
	defer c.mx.RUnlock()
	list := make(map[int64]*Context)
	// BY ID ONLY!
	for _, v := range c.byId {
		list[v.Id] = v
	}
	return list
}

func (c *Conditions) Remove(key *Condition) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.byId, key.Id)
}

func (r *Regexes) Remove(key *Regex) {
	r.mx.Lock()
	defer r.mx.Unlock()
	delete(r.byId, key.Id)
}

func (a *Actions) Remove(key *Action) {
	a.mx.Lock()
	defer a.mx.Unlock()
	delete(a.byId, key.Id)
}

func (a *AntiActions) Remove(key *AntiAction) {
	a.mx.Lock()
	defer a.mx.Unlock()
	delete(a.byId, key.Id)
}

func (c *Contexts) Rename(oldName, newName string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	if c.byName[oldName] == nil {
		return
	}
	c.byName[newName] = c.byName[oldName]
	c.byName[newName].Name = newName
	delete(c.byName, oldName)
}

func (e *Extensions) Rename(oldName, newName string) {
	e.mx.Lock()
	defer e.mx.Unlock()
	if e.byName[oldName] == nil {
		return
	}
	e.byName[newName] = e.byName[oldName]
	e.byName[newName].Name = newName
	delete(e.byName, oldName)
}

func (c *Contexts) HasName(key string) bool {
	c.mx.RLock()
	_, ok := c.byName[key]
	c.mx.RUnlock()
	return ok
}

func (e *Extensions) HasName(key string) bool {
	e.mx.RLock()
	_, ok := e.byName[key]
	e.mx.RUnlock()
	return ok
}

func (c *Contexts) Remove(key *Context) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.byName, key.Name)
	delete(c.byId, key.Id)
}

func (e *Extensions) Remove(key *Extension) {
	e.mx.Lock()
	defer e.mx.Unlock()
	delete(e.byName, key.Name)
	delete(e.byId, key.Id)
}

func (d *Dialplans) ClearUp() {
	d.ClearExtensions()
}

func (d *Dialplans) ClearExtensions() {
	d.Extensions.clearUp(d)
	d.ClearConditions()
}

func (d *Dialplans) ClearConditions() {
	d.Conditions.clearUp(d)
	d.Regexes.clearUp(d)
	d.Actions.clearUp(d)
	d.AntiActions.clearUp(d)
}

func (e *Extensions) clearUp(dialplan *Dialplans) {
	e.mx.Lock()
	defer e.mx.Unlock()
	for _, v := range e.byId {
		_, ok := dialplan.Contexts.GetById(v.Context.Id)
		if !ok {
			delete(e.byName, v.Name)
			delete(e.byId, v.Id)
		}
	}

}

func (c *Conditions) clearUp(dialpan *Dialplans) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for _, v := range c.byId {
		_, ok := dialpan.Extensions.GetById(v.Extension.Id)
		if !ok {
			delete(c.byId, v.Id)
		}
	}

}

func (r *Regexes) clearUp(dialpan *Dialplans) {
	r.mx.Lock()
	defer r.mx.Unlock()
	for _, v := range r.byId {
		_, ok := dialpan.Conditions.GetById(v.Condition.Id)
		if !ok {
			delete(r.byId, v.Id)
		}
	}
}

func (a *Actions) clearUp(dialpan *Dialplans) {
	a.mx.Lock()
	defer a.mx.Unlock()
	for _, v := range a.byId {
		_, ok := dialpan.Conditions.GetById(v.Condition.Id)
		if !ok {
			delete(a.byId, v.Id)
		}
	}
}

func (a *AntiActions) clearUp(dialpan *Dialplans) {
	a.mx.Lock()
	defer a.mx.Unlock()
	for _, v := range a.byId {
		_, ok := dialpan.Conditions.GetById(v.Condition.Id)
		if !ok {
			delete(a.byId, v.Id)
		}
	}
}

func (c *Context) CacheFullXML() {
	if !c.Dialplan.NoProceed {
		return
	}
	xmlContext := new(xmlStruct.Context)
	xmlContext.Attrname = c.Name
	for _, ext := range c.Extensions.ExtensionsList() {
		xmlExt := new(xmlStruct.Extension)
		xmlExt.Attrname = ext.Name
		xmlExt.Attrcontinue = ext.Continue
		for _, con := range ext.Conditions.ConditionsList() {
			xmlCon := new(xmlStruct.Condition)
			xmlCon.Attrexpression = con.Expression
			xmlCon.Attrfield = con.Field
			xmlCon.Attrbreak = con.Break
			xmlCon.Attrregex = con.Regex
			xmlCon.TimeOfDay = con.TimeOfDay
			xmlCon.Minute = con.Minute
			xmlCon.Week = con.Week
			xmlCon.Yday = con.Yday
			xmlCon.Yday = con.Yday
			xmlCon.Minday = con.Minday
			xmlCon.DateTime = con.DateTime
			xmlCon.TzOffset = con.TzOffset
			xmlCon.Year = con.Year
			xmlCon.Dst = con.Dst
			xmlCon.Attrwday = con.Wday
			xmlCon.Attrmweek = con.Mweek
			xmlCon.Attrmon = con.Mon
			xmlCon.Attrhour = con.Hour
			for _, act := range con.Actions.ActionsList() {
				xmlAct := new(xmlStruct.Action)
				xmlAct.Attrapplication = act.Application
				xmlAct.Attrdata = act.Data
				if act.Inline {
					xmlAct.Attrinline = "true"
				}
				xmlCon.Action = append(xmlCon.Action, xmlAct)
			}
			for _, aact := range con.AntiActions.ActionsList() {
				xmlAct := new(xmlStruct.AntiAction)
				xmlAct.Attrapplication = aact.Application
				xmlAct.Attrdata = aact.Data
				xmlCon.AntiAction = append(xmlCon.AntiAction, xmlAct)
			}
			for _, reg := range con.Regexes.RegexList() {
				xmlReg := new(xmlStruct.Regex)
				xmlReg.Attrexpression = reg.Expression
				xmlReg.Attrfield = reg.Field
				xmlCon.Regex = append(xmlCon.Regex, xmlReg)
			}
			xmlExt.Condition = append(xmlExt.Condition, xmlCon)
		}
		xmlContext.Extension = append(xmlContext.Extension, xmlExt)
	}

	xmlDocument := new(xmlStruct.Document)
	xmlDocument.Attrtype = "freeswitch/xml"
	xmlSection := new(xmlStruct.Section)
	xmlSection.Attrname = "dialplan"
	xmlSection.Context = []*xmlStruct.Context{xmlContext}
	xmlSections := []*xmlStruct.Section{xmlSection}
	xmlDocument.Section = xmlSections
	output, err := xml.MarshalIndent(xmlDocument, "", "  ")
	if err != nil {
		log.Printf("%+v", err)
		/*		output = []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
				<document type="freeswitch/xml">
				  <section name="result">
					<result status="not found" />
				  </section>
				</document>`)*/
		return
	}
	output = append(output, '\n')
	c.FullXMLCache.Set(output)
}
