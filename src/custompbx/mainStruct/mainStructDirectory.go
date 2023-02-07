package mainStruct

import (
	"log"
	"sync"
)

type Domains struct {
	mx     sync.RWMutex
	byName map[string]*Domain
	byId   map[int64]*Domain
}

type Groups struct {
	mx     sync.RWMutex
	byName map[string]*Group
	byId   map[int64]*Group
}

type GroupUsers struct {
	mx     sync.RWMutex
	byName map[string]*GroupUser
	byId   map[int64]*GroupUser
}

type DomainParams struct {
	mx     sync.RWMutex
	byName map[string]*DomainParam
	byId   map[int64]*DomainParam
}

type DomainVars struct {
	mx     sync.RWMutex
	byName map[string]*DomainVariable
	byId   map[int64]*DomainVariable
}

type UserParams struct {
	mx     sync.RWMutex
	byName map[string]*UserParam
	byId   map[int64]*UserParam
}

type GatewayParams struct {
	mx     sync.RWMutex
	byName map[string]*GatewayParam
	byId   map[int64]*GatewayParam
}

type GatewayVars struct {
	mx     sync.RWMutex
	byName map[string]*GatewayVariable
	byId   map[int64]*GatewayVariable
}

type UserVars struct {
	mx     sync.RWMutex
	byName map[string]*UserVariable
	byId   map[int64]*UserVariable
}

type Users struct {
	mx     sync.RWMutex
	byName map[string]*User
	byId   map[int64]*User
}

type UserGateways struct {
	mx     sync.RWMutex
	byName map[string]*UserGateway
	byId   map[int64]*UserGateway
}

type DirectoryItems struct {
	Domains       *Domains
	DomainParams  *DomainParams
	DomainVars    *DomainVars
	Groups        *Groups
	GroupUsers    *GroupUsers
	Users         *Users
	UserParams    *UserParams
	UserVars      *UserVars
	UserGateways  *UserGateways
	GatewayParams *GatewayParams
	GatewayVars   *GatewayVars
}

type Domain struct {
	Id             int64            `xml:"-" json:"id"`
	Enabled        bool             `xml:"-" json:"enabled"`
	Name           string           `xml:"name,attr" json:"name"`
	Params         *DomainParams    `xml:"-" json:"-"`
	Vars           *DomainVars      `xml:"-" json:"-"`
	Users          *Users           `xml:"-" json:"-"`
	Groups         *Groups          `xml:"-" json:"-"`
	XMLParams      []DomainParam    `xml:"params>param,omitempty" json:"-"`
	XMLVars        []DomainVariable `xml:"variables>variable,omitempty" json:"-"`
	XMLUsers       []User           `xml:"users>user,omitempty" json:"-"`
	XMLGroups      []Group          `xml:"groups>group" json:"-"`
	SipRegsCounter int              `xml:"-" json:"sip_regs_counter,omitempty"`
	// VertoRegsCounter int              `xml:"-" json:"verto_regs_counter"`
}

type XMLDomain struct {
	Name      string           `xml:"name,attr" json:"-"`
	XMLParams []DomainParam    `xml:"params>param,omitempty" json:"-"`
	XMLVars   []DomainVariable `xml:"variables>variable,omitempty" json:"-"`
	XMLGroups interface{}      `xml:"groups>group,omitempty" json:"-"`
}

type XMLGroup struct {
	Name     string      `xml:"name,attr" json:"-"`
	XMLUsers interface{} `xml:"users>user" json:"-"`
}

type Group struct {
	Id       int64       `xml:"-" json:"id"`
	Enabled  bool        `xml:"-" json:"enabled"`
	Domain   *Domain     `xml:"-" json:"-"`
	Name     string      `xml:"name,attr" json:"name"`
	Users    *GroupUsers `xml:"-" json:"users"`
	XMLUsers []GroupUser `xml:"users>user" json:"-"`
}

type GroupUser struct {
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Group   *Group `xml:"-" json:"-"`
	UserId  int64  `xml:"-" json:"user_id"`
	Name    string `xml:"id,attr" json:"name"`
	Type    string `xml:"type,attr" json:"type"`
}

type DomainParam struct {
	Domain  *Domain `xml:"-" json:"-"`
	Id      int64   `xml:"-" json:"id"`
	Enabled bool    `xml:"-" json:"enabled"`
	Name    string  `xml:"name,attr" json:"name"`
	Value   string  `xml:"value,attr" json:"value"`
}

type DomainVariable struct {
	Domain  *Domain `xml:"-" json:"-"`
	Id      int64   `xml:"-" json:"id"`
	Enabled bool    `xml:"-" json:"enabled"`
	Name    string  `xml:"name,attr" json:"name"`
	Value   string  `xml:"value,attr" json:"value"`
}

type UserParam struct {
	User    *User  `xml:"-" json:"-"`
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Name    string `xml:"name,attr" json:"name"`
	Value   string `xml:"value,attr" json:"value"`
}

type GatewayParam struct {
	Gateway *UserGateway `xml:"-" json:"-"`
	Id      int64        `xml:"-" json:"id"`
	Enabled bool         `xml:"-" json:"enabled"`
	Name    string       `xml:"name,attr" json:"name"`
	Value   string       `xml:"value,attr" json:"value"`
}

type GatewayVariable struct {
	Gateway   *UserGateway `xml:"-" json:"-"`
	Id        int64        `xml:"-" json:"id"`
	Enabled   bool         `xml:"-" json:"enabled"`
	Name      string       `xml:"name,attr" json:"name"`
	Value     string       `xml:"value,attr" json:"value"`
	Direction string       `xml:"direction,attr" json:"direction"`
}

type UserVariable struct {
	User    *User  `xml:"-" json:"-"`
	Id      int64  `xml:"-" json:"id"`
	Enabled bool   `xml:"-" json:"enabled"`
	Name    string `xml:"name,attr" json:"name"`
	Value   string `xml:"value,attr" json:"value"`
}

type User struct {
	Domain        *Domain        `xml:"-" json:"-"`
	Id            int64          `xml:"-" json:"id"`
	Enabled       bool           `xml:"-" json:"enabled"`
	Name          string         `xml:"id,attr" json:"name"`
	Cache         uint           `xml:"cacheable,attr,omitempty" json:"cache"`
	Cidr          string         `xml:"cidr,attr,omitempty" json:"cidr"`
	NumberAlias   string         `xml:"number-alias,attr,omitempty" json:"number_alias"`
	XMLParams     []UserParam    `xml:"params>param,omitempty" json:"-"`
	XMLVars       []UserVariable `xml:"variables>variable,omitempty" json:"-"`
	XMLGateways   []UserGateway  `xml:"gateways>gateway,omitempty" json:"-"`
	Params        *UserParams    `xml:"-" json:"-"`
	Vars          *UserVars      `xml:"-" json:"-"`
	Gateways      *UserGateways  `xml:"-" json:"-"`
	CallDate      int64          `xml:"-" json:"call_date"`
	InCall        bool           `xml:"-" json:"in_call"`
	Talking       bool           `xml:"-" json:"talking"`
	LastUuid      string         `xml:"-" json:"last_uuid"`
	CallDirection string         `xml:"-" json:"call_direction"`
	SipRegister   bool           `xml:"-" json:"sip_register"`
	VertoRegister bool           `xml:"-" json:"verto_register"`
	CCAgent       int64          `xml:"-" json:"cc_agent,omitempty"`
}

type UserGateway struct {
	User      *User             `xml:"-" json:"-"`
	Id        int64             `xml:"-" json:"id"`
	Enabled   bool              `xml:"-" json:"enabled"`
	Name      string            `xml:"name,attr" json:"name"`
	Params    *GatewayParams    `xml:"-" json:"-"`
	XMLParams []GatewayParam    `xml:"param,omitempty" json:"-"`
	Vars      *GatewayVars      `xml:"-" json:"-"`
	XMLVars   []GatewayVariable `xml:"variable,omitempty" json:"-"`
}

func NewDomains() *Domains {
	return &Domains{
		byName: make(map[string]*Domain),
		byId:   make(map[int64]*Domain),
	}
}

func NewDomainParams() *DomainParams {
	return &DomainParams{
		byName: make(map[string]*DomainParam),
		byId:   make(map[int64]*DomainParam),
	}
}

func NewUserParams() *UserParams {
	return &UserParams{
		byName: make(map[string]*UserParam),
		byId:   make(map[int64]*UserParam),
	}
}

func NewGatewayParams() *GatewayParams {
	return &GatewayParams{
		byName: make(map[string]*GatewayParam),
		byId:   make(map[int64]*GatewayParam),
	}
}

func NewGatewayVars() *GatewayVars {
	return &GatewayVars{
		byName: make(map[string]*GatewayVariable),
		byId:   make(map[int64]*GatewayVariable),
	}
}

func NewDomainVars() *DomainVars {
	return &DomainVars{
		byName: make(map[string]*DomainVariable),
		byId:   make(map[int64]*DomainVariable),
	}
}

func NewUserGateways() *UserGateways {
	return &UserGateways{
		byName: make(map[string]*UserGateway),
		byId:   make(map[int64]*UserGateway),
	}
}

func NewUserVars() *UserVars {
	return &UserVars{
		byName: make(map[string]*UserVariable),
		byId:   make(map[int64]*UserVariable),
	}
}

func NewUsers() *Users {
	return &Users{
		byName: make(map[string]*User),
		byId:   make(map[int64]*User),
	}
}

func NewGroupUsers() *GroupUsers {
	return &GroupUsers{
		byName: make(map[string]*GroupUser),
		byId:   make(map[int64]*GroupUser),
	}
}

func NewGroups() *Groups {
	return &Groups{
		byName: make(map[string]*Group),
		byId:   make(map[int64]*Group),
	}
}

func NewDirectoryItems() *DirectoryItems {
	return &DirectoryItems{
		Domains:       NewDomains(),
		DomainParams:  NewDomainParams(),
		DomainVars:    NewDomainVars(),
		Groups:        NewGroups(),
		GroupUsers:    NewGroupUsers(),
		Users:         NewUsers(),
		UserParams:    NewUserParams(),
		UserVars:      NewUserVars(),
		UserGateways:  NewUserGateways(),
		GatewayParams: NewGatewayParams(),
		GatewayVars:   NewGatewayVars(),
	}
}
func (d *DirectoryItems) ClearUp() {
	log.Printf("going to clean\n")
	d.DomainParams.clearUp(d)
	d.DomainVars.clearUp(d)
	d.ClearDirectoryGroups()
	d.ClearDirectoryUsers()
	log.Printf("cleaned\n")
}

func (d *DirectoryItems) ClearDirectoryGroups() {
	d.Groups.clearUp(d)
	d.GroupUsers.clearUp(d)
}

func (d *DirectoryItems) ClearDirectoryUsers() {
	d.Users.clearUp(d)
	d.UserParams.clearUp(d)
	d.UserVars.clearUp(d)
	d.ClearDirectoryUserGateways()
}

func (d *DirectoryItems) ClearDirectoryUserGateways() {
	d.UserGateways.clearUp(d)
	d.GatewayParams.clearUp(d)
	d.GatewayVars.clearUp(d)
}

func (d *DomainParams) clearUp(directory *DirectoryItems) {
	d.mx.Lock()
	defer d.mx.Unlock()
	for _, v := range d.byId {
		domain := directory.Domains.GetById(v.Domain.Id)
		if domain == nil {
			delete(d.byName, v.Name)
			delete(d.byId, v.Id)
		}
	}
}

func (d *DomainVars) clearUp(directory *DirectoryItems) {
	d.mx.Lock()
	defer d.mx.Unlock()
	for _, v := range d.byId {
		domain := directory.Domains.GetById(v.Domain.Id)
		if domain == nil {
			delete(d.byName, v.Name)
			delete(d.byId, v.Id)
		}
	}
}

func (g *Groups) clearUp(directory *DirectoryItems) {
	g.mx.Lock()
	defer g.mx.Unlock()
	for _, v := range g.byId {
		domain := directory.Domains.GetById(v.Domain.Id)
		if domain == nil {
			delete(g.byName, v.Name)
			delete(g.byId, v.Id)
		}
	}
}

func (g *GroupUsers) clearUp(directory *DirectoryItems) {
	g.mx.Lock()
	defer g.mx.Unlock()
	for _, v := range g.byId {
		_, ok := directory.Groups.GetById(v.Group.Id)
		if !ok {
			delete(g.byName, v.Name)
			delete(g.byId, v.Id)
		}
	}
}

func (u *Users) clearUp(directory *DirectoryItems) {
	u.mx.Lock()
	defer u.mx.Unlock()
	for _, v := range u.byId {
		domain := directory.Domains.GetById(v.Domain.Id)
		if domain == nil {
			delete(u.byName, v.Name)
			delete(u.byId, v.Id)
		}
	}
}

func (u *UserParams) clearUp(directory *DirectoryItems) {
	u.mx.Lock()
	defer u.mx.Unlock()
	for _, v := range u.byId {
		user := directory.Users.GetById(v.User.Id)
		if user == nil {
			delete(u.byName, v.Name)
			delete(u.byId, v.Id)
		}
	}
}

func (u *UserVars) clearUp(directory *DirectoryItems) {
	u.mx.Lock()
	defer u.mx.Unlock()
	for _, v := range u.byId {
		user := directory.Users.GetById(v.User.Id)
		if user == nil {
			delete(u.byName, v.Name)
			delete(u.byId, v.Id)
		}
	}
}

func (u *UserGateways) clearUp(directory *DirectoryItems) {
	u.mx.Lock()
	defer u.mx.Unlock()
	for _, v := range u.byId {
		user := directory.Users.GetById(v.User.Id)
		if user == nil {
			delete(u.byName, v.Name)
			delete(u.byId, v.Id)
		}
	}
}

func (g *GatewayParams) clearUp(directory *DirectoryItems) {
	g.mx.Lock()
	defer g.mx.Unlock()
	for _, v := range g.byId {
		_, ok := directory.UserGateways.GetById(v.Gateway.Id)
		if !ok {
			delete(g.byName, v.Name)
			delete(g.byId, v.Id)
		}
	}
}

func (g *GatewayVars) clearUp(directory *DirectoryItems) {
	g.mx.Lock()
	defer g.mx.Unlock()
	for _, v := range g.byId {
		_, ok := directory.UserGateways.GetById(v.Gateway.Id)
		if !ok {
			delete(g.byName, v.Name)
			delete(g.byId, v.Id)
		}
	}
}

func (d *Domain) NewParams() {
	d.Params = &DomainParams{
		byName: make(map[string]*DomainParam),
		byId:   make(map[int64]*DomainParam),
	}
}

func (d *Domain) NewVars() {
	d.Vars = &DomainVars{
		byName: make(map[string]*DomainVariable),
		byId:   make(map[int64]*DomainVariable),
	}
}

func (d *Domain) NewUsers() {
	d.Users = &Users{
		byName: make(map[string]*User),
		byId:   make(map[int64]*User),
	}
}

func (d *Domain) NewGroups() {
	d.Groups = &Groups{
		byName: make(map[string]*Group),
		byId:   make(map[int64]*Group),
	}
}

func (g *Group) NewGroupUsers() {
	g.Users = &GroupUsers{
		byName: make(map[string]*GroupUser),
		byId:   make(map[int64]*GroupUser),
	}
}

func (u *User) NewUserParams() {
	u.Params = &UserParams{
		byName: make(map[string]*UserParam),
		byId:   make(map[int64]*UserParam),
	}
}

func (u *User) NewUserVars() {
	u.Vars = &UserVars{
		byName: make(map[string]*UserVariable),
		byId:   make(map[int64]*UserVariable),
	}
}

func (u *User) NewUserGateways() {
	u.Gateways = &UserGateways{
		byName: make(map[string]*UserGateway),
		byId:   make(map[int64]*UserGateway),
	}
}

func (g *UserGateway) NewGatewayParams() {
	g.Params = &GatewayParams{
		byName: make(map[string]*GatewayParam),
		byId:   make(map[int64]*GatewayParam),
	}
}

func (g *UserGateway) NewGatewayVars() {
	g.Vars = &GatewayVars{
		byName: make(map[string]*GatewayVariable),
		byId:   make(map[int64]*GatewayVariable),
	}
}

func (d *Domains) GetByName(key string) *Domain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	val := d.byName[key]
	return val
}

func (u *UserGateways) GetByName(key string) (*UserGateway, bool) {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val, ok := u.byName[key]
	return val, ok
}

func (g *Groups) GetByName(key string) (*Group, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byName[key]
	return val, ok
}

func (g *GroupUsers) GetByName(key string) (*GroupUser, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byName[key]
	return val, ok
}

func (d *DomainParams) GetByName(key string) *DomainParam {
	d.mx.RLock()
	defer d.mx.RUnlock()
	return d.byName[key]
}

func (g *GatewayParams) GetByName(key string) (*GatewayParam, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byName[key]
	return val, ok
}

func (u *UserParams) GetByName(key string) *UserParam {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val, _ := u.byName[key]
	return val
}

func (d *DomainVars) GetByName(key string) (*DomainVariable, bool) {
	d.mx.RLock()
	defer d.mx.RUnlock()
	val, ok := d.byName[key]
	return val, ok
}

func (u *UserVars) GetByName(key string) (*UserVariable, bool) {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val, ok := u.byName[key]
	return val, ok
}

func (u *Users) GetByName(key string) *User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val := u.byName[key]
	return val
}

func (d *Domains) GetById(key int64) *Domain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	return d.byId[key]
}

func (u *UserGateways) GetById(key int64) (*UserGateway, bool) {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val, ok := u.byId[key]
	return val, ok
}

func (g *Groups) GetById(key int64) (*Group, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byId[key]
	return val, ok
}

func (g *GroupUsers) GetById(key int64) (*GroupUser, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byId[key]
	return val, ok
}

func (d *DomainParams) GetById(key int64) (*DomainParam, bool) {
	d.mx.RLock()
	defer d.mx.RUnlock()
	val, ok := d.byId[key]
	return val, ok
}

func (d *DomainVars) GetById(key int64) (*DomainVariable, bool) {
	d.mx.RLock()
	defer d.mx.RUnlock()
	val, ok := d.byId[key]
	return val, ok
}

func (u *Users) GetById(key int64) *User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	return u.byId[key]
}

func (u *UserParams) GetById(key int64) (*UserParam, bool) {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val, ok := u.byId[key]
	return val, ok
}

func (u *UserVars) GetById(key int64) (*UserVariable, bool) {
	u.mx.RLock()
	defer u.mx.RUnlock()
	val, ok := u.byId[key]
	return val, ok
}

func (g *GatewayParams) GetById(key int64) (*GatewayParam, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byId[key]
	return val, ok
}

func (g *GatewayVars) GetById(key int64) (*GatewayVariable, bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val, ok := g.byId[key]
	return val, ok
}

func (g *GatewayVars) Remove(key *GatewayVariable) {
	g.mx.Lock()
	defer g.mx.Unlock()
	delete(g.byName, key.Name)
	delete(g.byId, key.Id)
}

func (g *Groups) Set(value *Group) {
	g.mx.Lock()
	defer g.mx.Unlock()
	g.byName[value.Name] = value
	g.byId[value.Id] = value
}

func (g *GroupUsers) Set(value *GroupUser) {
	g.mx.Lock()
	defer g.mx.Unlock()
	g.byName[value.Name] = value
	g.byId[value.Id] = value
}

func (d *Domains) Set(value *Domain) {
	d.mx.Lock()
	defer d.mx.Unlock()
	d.byName[value.Name] = value
	d.byId[value.Id] = value
}

func (u *UserGateways) Set(value *UserGateway) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.byName[value.Name] = value
	u.byId[value.Id] = value
}

func (d *DomainParams) Set(value *DomainParam) {
	d.mx.Lock()
	defer d.mx.Unlock()
	d.byName[value.Name] = value
	d.byId[value.Id] = value
}

func (u *UserParams) Set(value *UserParam) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.byName[value.Name] = value
	u.byId[value.Id] = value
}

func (g *GatewayParams) Set(value *GatewayParam) {
	g.mx.Lock()
	defer g.mx.Unlock()
	g.byName[value.Name] = value
	g.byId[value.Id] = value
}

func (g *GatewayVars) Set(value *GatewayVariable) {
	g.mx.Lock()
	defer g.mx.Unlock()
	g.byName[value.Name] = value
	g.byId[value.Id] = value
}

func (d *DomainVars) Set(value *DomainVariable) {
	d.mx.Lock()
	defer d.mx.Unlock()
	d.byName[value.Name] = value
	d.byId[value.Id] = value
}

func (u *UserVars) Set(value *UserVariable) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.byName[value.Name] = value
	u.byId[value.Id] = value
}

func (u *Users) Set(value *User) {
	u.mx.Lock()
	defer u.mx.Unlock()
	u.byName[value.Name] = value
	u.byId[value.Id] = value
}

// BY ID ONLY!
func (d *Domains) GetList() map[int64]string {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[int64]string)
	// BY ID ONLY!
	for _, v := range d.byId {
		list[v.Id] = v.Name
	}
	return list
}

func (d *Domains) GetSipRegCounterList() map[string]int {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[string]int)
	// BY ID ONLY!
	for _, v := range d.byId {
		list[v.Name] = v.SipRegsCounter
	}
	return list
}

func (u *Users) GetList() map[int64]string {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]string)
	for _, v := range u.byId {
		list[v.Id] = v.Name
	}
	return list
}

func (u *Users) GetDomainsList() map[int64]map[int64]*User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]map[int64]*User)
	for _, v := range u.byId {
		if _, ok := list[v.Domain.Id]; !ok {
			list[v.Domain.Id] = make(map[int64]*User)
		}
		list[v.Domain.Id][v.Id] = v
	}
	return list
}

func (g *Groups) GetDomainsList() map[int64]map[int64]string {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]map[int64]string)
	for _, v := range g.byId {
		if _, ok := list[v.Domain.Id]; !ok {
			list[v.Domain.Id] = make(map[int64]string)
		}
		list[v.Domain.Id][v.Id] = v.Name
	}
	return list
}

func (g *Groups) GetUsersList() map[int64]map[int64]string {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]map[int64]string)
	for _, v := range g.byId {
		if _, ok := list[v.Domain.Id]; !ok {
			list[v.Domain.Id] = make(map[int64]string)
		}
		list[v.Domain.Id][v.Id] = v.Name
	}
	return list
}

func (g *Groups) GetList() map[int64]string {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]string)
	for _, v := range g.byId {
		list[v.Id] = v.Name
	}
	return list
}

func (u *UserGateways) GetNamesList() map[int64]string {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]string)
	for _, v := range u.byId {
		list[v.Id] = v.Name
	}
	return list
}

func (d *Domains) GetCopy() map[int64]*Domain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[int64]*Domain)
	// BY ID ONLY!
	for _, v := range d.byId {
		list[v.Id] = v
	}
	return list
}

func (u *Users) GetCopy() map[int64]*User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]*User)
	// BY ID ONLY!
	for _, v := range u.byId {
		list[v.Id] = v
	}
	return list
}

func (g *Groups) GetCopy() map[int64]*Group {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]*Group)
	// BY ID ONLY!
	for _, v := range g.byId {
		list[v.Id] = v
	}
	return list
}

func (u *UserGateways) GetCopy() map[int64]*UserGateway {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]*UserGateway)
	// BY ID ONLY!
	for _, v := range u.byId {
		list[v.Id] = v
	}
	return list
}

func (d *Domains) Names() []string {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var keys []string
	for k := range d.byName {
		keys = append(keys, k)
	}
	return keys
}

func (d *Domains) Props() []*Domain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var items []*Domain
	for _, v := range d.byId {
		items = append(items, v)
	}
	return items
}

func (u *Users) Names() []string {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var keys []string
	for k := range u.byName {
		keys = append(keys, k)
	}
	return keys
}

func (u *Users) Props() []*User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var items []*User
	for _, v := range u.byId {
		items = append(items, v)
	}
	return items
}

func (g *Groups) Names() []string {
	g.mx.RLock()
	defer g.mx.RUnlock()
	var keys []string
	for k := range g.byName {
		keys = append(keys, k)
	}
	return keys
}

func (g *Groups) Props() []*Group {
	g.mx.RLock()
	defer g.mx.RUnlock()
	var items []*Group
	for _, v := range g.byId {
		items = append(items, v)
	}
	return items
}

func (u *UserGateways) Names() []string {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var keys []string
	for k := range u.byName {
		keys = append(keys, k)
	}
	return keys
}

func (u *UserGateways) Props() []*UserGateway {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var items []*UserGateway
	for _, v := range u.byId {
		items = append(items, v)
	}
	return items
}

func (d *Domains) HasName(key string) bool {
	d.mx.RLock()
	_, ok := d.byName[key]
	d.mx.RUnlock()
	return ok
}

func (d *DomainParams) HasName(key string) bool {
	d.mx.RLock()
	_, ok := d.byName[key]
	d.mx.RUnlock()
	return ok
}

func (u *UserParams) HasName(key string) bool {
	u.mx.RLock()
	_, ok := u.byName[key]
	u.mx.RUnlock()
	return ok
}

func (g *GatewayParams) HasName(key string) bool {
	g.mx.RLock()
	_, ok := g.byName[key]
	g.mx.RUnlock()
	return ok
}

func (d *DomainVars) HasName(key string) bool {
	d.mx.RLock()
	_, ok := d.byName[key]
	d.mx.RUnlock()
	return ok
}

func (u *UserVars) HasName(key string) bool {
	u.mx.RLock()
	_, ok := u.byName[key]
	u.mx.RUnlock()
	return ok
}

func (g *Groups) HasName(key string) bool {
	g.mx.RLock()
	_, ok := g.byName[key]
	g.mx.RUnlock()
	return ok
}

func (g *GroupUsers) HasName(key string) bool {
	g.mx.RLock()
	_, ok := g.byName[key]
	g.mx.RUnlock()
	return ok
}

func (u *UserGateways) HasName(key string) bool {
	u.mx.RLock()
	_, ok := u.byName[key]
	u.mx.RUnlock()
	return ok
}

func (d *DomainParams) GetList() map[int64]*DomainParam {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[int64]*DomainParam)
	for _, val := range d.byId {
		list[val.Id] = val
	}
	return list
}

func (u *UserParams) GetList() map[int64]*UserParam {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]*UserParam)
	for _, val := range u.byId {
		list[val.Id] = val
	}
	return list
}

func (g *GatewayParams) GetList() map[int64]*GatewayParam {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]*GatewayParam)
	for _, val := range g.byId {
		list[val.Id] = val
	}
	return list
}

func (g *GatewayVars) GetList() map[int64]*GatewayVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]*GatewayVariable)
	for _, val := range g.byId {
		list[val.Id] = val
	}
	return list
}

func (u *UserGateways) GetList() map[int64]*UserGateway {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]*UserGateway)
	for _, val := range u.byId {
		list[val.Id] = val
	}
	return list
}

func (d *DomainVars) GetList() map[int64]*DomainVariable {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[int64]*DomainVariable)
	for _, val := range d.byId {
		list[val.Id] = val
	}
	return list
}

func (u *UserVars) GetList() map[int64]*UserVariable {
	u.mx.RLock()
	defer u.mx.RUnlock()
	list := make(map[int64]*UserVariable)
	for _, val := range u.byId {
		list[val.Id] = val
	}
	return list
}

func (g *GroupUsers) GetList() map[int64]*GroupUser {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]*GroupUser)
	for _, val := range g.byId {
		list[val.Id] = val
	}
	return list
}

func (d *Domains) Rename(oldName, newName string) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if d.byName[oldName] == nil {
		return
	}
	d.byName[newName] = d.byName[oldName]
	d.byName[newName].Name = newName
	delete(d.byName, oldName)
}

func (u *Users) Rename(oldName, newName string) {
	u.mx.Lock()
	defer u.mx.Unlock()
	if u.byName[oldName] == nil {
		return
	}
	u.byName[newName] = u.byName[oldName]
	u.byName[newName].Name = newName
	delete(u.byName, oldName)
}

func (u *UserGateways) Rename(oldName, newName string) {
	u.mx.Lock()
	defer u.mx.Unlock()
	if u.byName[oldName] == nil {
		return
	}
	u.byName[newName] = u.byName[oldName]
	u.byName[newName].Name = newName
	delete(u.byName, oldName)
}

func (g *Groups) Rename(oldName, newName string) {
	g.mx.Lock()
	defer g.mx.Unlock()
	if g.byName[oldName] == nil {
		return
	}
	g.byName[newName] = g.byName[oldName]
	g.byName[newName].Name = newName
	delete(g.byName, oldName)
}

func (u *Users) Remove(key User) {
	u.mx.Lock()
	defer u.mx.Unlock()
	delete(u.byName, key.Name)
	delete(u.byId, key.Id)
}

func (g *Groups) Remove(key Group) {
	g.mx.Lock()
	defer g.mx.Unlock()
	delete(g.byName, key.Name)
	delete(g.byId, key.Id)
}

func (g *GroupUsers) Remove(key GroupUser) {
	g.mx.Lock()
	defer g.mx.Unlock()
	delete(g.byName, key.Name)
	delete(g.byId, key.Id)
}

func (u *UserGateways) Remove(key UserGateway) {
	u.mx.Lock()
	defer u.mx.Unlock()
	delete(u.byName, key.Name)
	delete(u.byId, key.Id)
}

func (d *DomainParams) Remove(key DomainParam) {
	d.mx.Lock()
	defer d.mx.Unlock()
	delete(d.byName, key.Name)
	delete(d.byId, key.Id)
}

func (u *UserParams) Remove(key UserParam) {
	u.mx.Lock()
	defer u.mx.Unlock()
	delete(u.byName, key.Name)
	delete(u.byId, key.Id)
}

func (g *GatewayParams) Remove(key GatewayParam) {
	g.mx.Lock()
	defer g.mx.Unlock()
	delete(g.byName, key.Name)
	delete(g.byId, key.Id)
}

func (d *DomainVars) Remove(key DomainVariable) {
	d.mx.Lock()
	defer d.mx.Unlock()
	delete(d.byName, key.Name)
	delete(d.byId, key.Id)
}

func (u *UserVars) Remove(key UserVariable) {
	u.mx.Lock()
	defer u.mx.Unlock()
	delete(u.byName, key.Name)
	delete(u.byId, key.Id)
}

func (d *Domains) Remove(key Domain) {
	d.mx.Lock()
	defer d.mx.Unlock()
	delete(d.byName, key.Name)
	delete(d.byId, key.Id)
}

func (d *Domains) XMLItems() []XMLDomain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var domains []XMLDomain
	for _, v := range d.byId {
		if !v.Enabled {
			continue
		}
		var domain XMLDomain
		domain.Name = v.Name
		domain.XMLParams = v.Params.XMLItems()
		domain.XMLVars = v.Vars.XMLItems()
		domain.XMLGroups = append([]interface{}{XMLGroup{Name: "CustomPbxGroupForFS1.6", XMLUsers: v.Users.XMLItems()}}, v.Groups.XMLItems())
		domains = append(domains, domain)
	}
	return domains
}

func (d *DomainParams) XMLItems() []DomainParam {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var params []DomainParam
	for _, v := range d.byId {
		if !v.Enabled {
			continue
		}
		params = append(params, *v)
	}
	return params
}

func (u *UserParams) XMLItems() []UserParam {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var params []UserParam
	for _, v := range u.byId {
		if !v.Enabled {
			continue
		}
		params = append(params, *v)
	}
	return params
}

func (d *DomainVars) XMLItems() []DomainVariable {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var variables []DomainVariable
	for _, v := range d.byId {
		if !v.Enabled {
			continue
		}
		variables = append(variables, *v)
	}
	return variables
}

func (u *UserVars) XMLItems() []UserVariable {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var variables []UserVariable
	for _, v := range u.byId {
		if !v.Enabled {
			continue
		}
		variables = append(variables, *v)
	}
	return variables
}

func (u *UserGateways) XMLItems() []UserGateway {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var gateways []UserGateway
	for _, v := range u.byId {
		if !v.Enabled {
			continue
		}
		gateways = append(gateways, *v)
	}
	return gateways
}

func (u *Users) XMLItems() []User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var users []User
	for _, v := range u.byId {
		if !v.Enabled {
			continue
		}
		var user User
		user.Name = v.Name
		user.Cache = v.Cache
		user.Cidr = v.Cidr
		user.NumberAlias = v.NumberAlias
		user.XMLParams = v.Params.XMLItems()
		user.XMLVars = v.Vars.XMLItems()
		user.XMLGateways = v.Gateways.XMLItems()
		users = append(users, user)
	}
	return users
}

func (g *Groups) XMLItems() []Group {
	g.mx.RLock()
	defer g.mx.RUnlock()
	var groups []Group
	for _, v := range g.byId {
		if !v.Enabled {
			continue
		}
		var group Group
		group.Name = v.Name
		group.XMLUsers = v.Users.XMLItems()
		groups = append(groups, group)
	}
	return groups
}

func (g *Group) XMLItems() *Group {
	if g.Enabled == false {
		return nil
	}
	var group Group
	group.Name = g.Name
	group.XMLUsers = g.Users.XMLItems()
	return &group
}

func (u *User) XMLItems() *User {
	if u.Enabled == false {
		return nil
	}
	var user User
	user.Name = u.Name
	user.Cache = u.Cache
	user.Cidr = u.Cidr
	user.NumberAlias = u.NumberAlias
	user.XMLParams = u.Params.XMLItems()
	user.XMLVars = u.Vars.XMLItems()
	user.XMLGateways = u.Gateways.XMLItems()
	return &user
}

func (d *Domain) XMLItems() *XMLDomain {
	if d.Enabled == false {
		return nil
	}
	var domain XMLDomain
	domain.Name = d.Name
	domain.XMLParams = d.Params.XMLItems()
	domain.XMLVars = d.Vars.XMLItems()
	return &domain
}

func (d *Domains) XMLSafe() []XMLDomain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var domains []XMLDomain
	for _, v := range d.byId {
		if !v.Enabled {
			continue
		}
		var domain XMLDomain
		domain.Name = v.Name
		domain.XMLGroups = append([]interface{}{XMLGroup{Name: "CustomPbxGroupForFS1.6", XMLUsers: v.Users.XMLItems()}}, v.Groups.XMLItems())
		domains = append(domains, domain)
	}
	return domains
}

func (d *Domains) XMLFullSafe() []XMLDomain {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var domains []XMLDomain
	for _, v := range d.byId {
		if !v.Enabled {
			continue
		}
		var domain XMLDomain
		domain.Name = v.Name
		domain.XMLGroups = []XMLGroup{{Name: "CustomPbxGroupForFS1.6", XMLUsers: v.Users.XMLSafe()}}
		domains = append(domains, domain)
	}
	return domains
}

func (u *Users) XMLSafe() []User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var users []User
	for _, v := range u.byId {
		if !v.Enabled {
			continue
		}
		var user User
		user.Name = v.Name
		user.Cache = v.Cache
		user.Cidr = v.Cidr
		user.NumberAlias = v.NumberAlias
		user.XMLVars = v.Vars.XMLItems()
		users = append(users, user)
	}
	return users
}

func (u *Users) XMLFullSafe() []User {
	u.mx.RLock()
	defer u.mx.RUnlock()
	var users []User
	for _, v := range u.byId {
		if !v.Enabled {
			continue
		}
		var user User
		user.Name = v.Name
		user.Cache = v.Cache
		user.Cidr = v.Cidr
		user.NumberAlias = v.NumberAlias
		users = append(users, user)
	}
	return users
}

func (g *GroupUsers) XMLItems() []GroupUser {
	g.mx.RLock()
	defer g.mx.RUnlock()
	var user []GroupUser
	for _, v := range g.byId {
		if !v.Enabled {
			continue
		}
		user = append(user, *v)
	}
	return user
}
