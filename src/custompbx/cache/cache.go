package cache

import (
	"custompbx/altStruct"
	"custompbx/db"
	"custompbx/mainStruct"
	"sync"
)

var currentFsInstanceName string             //setting in esl connection
var currentFsInstance *mainStruct.FsInstance //setting in temp esl connection on start
var fsInstances *mainStruct.FsInstances
var directoryCache *directoryCacheStruct

type domainCache struct {
	mx   sync.RWMutex
	byId map[int64]*domainCacheStruct
	//byName map[string]*domainCacheStruct
}

type userCache struct {
	mx   sync.RWMutex
	byId map[int64]*userCacheStruct
	//byName map[string]*userCacheStruct
}

type directoryCacheStruct struct {
	DomainCache domainCache
	UserCache   userCache
}

type domainCacheStruct struct {
	Id             int64
	Name           string
	SipRegsCounter int
}

type userCacheStruct struct {
	Id             int64
	Name           string
	SipRegsCounter int

	CallDate      int64
	InCall        bool
	Talking       bool
	LastUuid      string
	CallDirection string
	SipRegister   bool
	VertoRegister bool
	CCAgent       int64
	activeCalls   map[string]UserCallState
}

type UserCallState struct {
	UUID      string
	CreatedAt int64
	Direction string
	Talking   bool
}

func GetDomainSipRegsCounter() map[string]int {
	return directoryCache.DomainCache.GetSipRegCounterList()
}

func (d *domainCache) GetSipRegCounterList() map[string]int {
	d.mx.RLock()
	defer d.mx.RUnlock()
	list := make(map[string]int)
	for _, v := range d.byId {
		list[v.Name] = v.SipRegsCounter
	}
	return list
}

func NewFsDirectoryCache() *directoryCacheStruct {
	return &directoryCacheStruct{
		DomainCache: domainCache{
			//byName: make(map[string]*domainCacheStruct),
			byId: make(map[int64]*domainCacheStruct),
		},
		UserCache: userCache{
			//byName: make(map[string]*userCacheStruct),
			byId: make(map[int64]*userCacheStruct),
		},
	}
}

func (f *domainCache) GetById(key int64) *domainCacheStruct {
	f.mx.RLock()
	defer f.mx.RUnlock()
	return f.byId[key]
}

/*
	func (f *domainCache) GetByName(key string) *domainCacheStruct {
		f.mx.RLock()
		defer f.mx.RUnlock()
		val := f.byName[key]
		return val
	}
*/
func (f *domainCache) Set(value *domainCacheStruct) {
	f.mx.Lock()
	defer f.mx.Unlock()
	//f.byName[value.Name] = value
	f.byId[value.Id] = value
}

func (f *domainCache) SetByData(id int64, name string, reg int) *domainCacheStruct {
	var value = &domainCacheStruct{Id: id, Name: name, SipRegsCounter: reg}
	f.mx.Lock()
	defer f.mx.Unlock()
	//f.byName[value.Name] = value
	f.byId[value.Id] = value
	return value
}

func (f *userCache) GetById(key int64) *userCacheStruct {
	f.mx.RLock()
	defer f.mx.RUnlock()
	return f.byId[key]
}

func (u *userCacheStruct) UpdateUser(user *altStruct.DirectoryDomainUser) {
	if user == nil {
		return
	}
	user.SipRegister = u.SipRegister
	user.VertoRegister = u.VertoRegister
	user.CallDate = u.CallDate
	user.CCAgent = u.CCAgent
	// user.SipRegister = u.SipRegsCounter
	user.Talking = u.Talking
	user.InCall = u.InCall
	user.CallDirection = u.CallDirection
	user.LastUuid = u.LastUuid
}

func (f *userCache) SetCall(user *altStruct.DirectoryDomainUser, call UserCallState, active bool) {
	if user == nil || user.Id == 0 || call.UUID == "" {
		return
	}
	f.mx.Lock()
	defer f.mx.Unlock()
	cached := f.byId[user.Id]
	if cached == nil {
		cached = &userCacheStruct{Id: user.Id, Name: user.Name, activeCalls: make(map[string]UserCallState)}
		f.byId[user.Id] = cached
	}
	if cached.activeCalls == nil {
		cached.activeCalls = make(map[string]UserCallState)
	}
	if active {
		cached.activeCalls[call.UUID] = call
	} else {
		delete(cached.activeCalls, call.UUID)
	}
	cached.refreshCallStateLocked()
	cached.updateUserLocked(user)
}

func (f *userCache) ActiveCallUserIDs() []int64 {
	f.mx.RLock()
	defer f.mx.RUnlock()
	ids := make([]int64, 0)
	for id, user := range f.byId {
		if len(user.activeCalls) > 0 {
			ids = append(ids, id)
		}
	}
	return ids
}

func (f *userCache) ReplaceCalls(calls map[int64][]UserCallState, users map[int64]*altStruct.DirectoryDomainUser) {
	f.mx.Lock()
	defer f.mx.Unlock()
	for _, cached := range f.byId {
		cached.activeCalls = make(map[string]UserCallState)
		cached.refreshCallStateLocked()
	}
	for id, userCalls := range calls {
		cached := f.byId[id]
		if cached == nil {
			name := ""
			if users[id] != nil {
				name = users[id].Name
			}
			cached = &userCacheStruct{Id: id, Name: name, activeCalls: make(map[string]UserCallState)}
			f.byId[id] = cached
		}
		for _, call := range userCalls {
			if call.UUID != "" {
				cached.activeCalls[call.UUID] = call
			}
		}
		cached.refreshCallStateLocked()
	}
	for id, user := range users {
		if cached := f.byId[id]; cached != nil {
			cached.updateUserLocked(user)
		}
	}
}

func (f *userCache) ApplyToUser(user *altStruct.DirectoryDomainUser) {
	if user == nil {
		return
	}
	f.mx.RLock()
	defer f.mx.RUnlock()
	if cached := f.byId[user.Id]; cached != nil {
		cached.updateUserLocked(user)
	}
}

func (u *userCacheStruct) refreshCallStateLocked() {
	u.InCall = len(u.activeCalls) > 0
	u.Talking = false
	u.CallDate = 0
	u.LastUuid = ""
	u.CallDirection = ""
	var selectedTalkingAt int64
	for _, call := range u.activeCalls {
		if u.CallDate == 0 || (call.CreatedAt > 0 && call.CreatedAt < u.CallDate) {
			u.CallDate = call.CreatedAt
			u.CallDirection = call.Direction
			if !u.Talking {
				u.LastUuid = call.UUID
			}
		}
		if call.Talking && (!u.Talking || selectedTalkingAt == 0 || (call.CreatedAt > 0 && call.CreatedAt < selectedTalkingAt)) {
			u.Talking = true
			u.LastUuid = call.UUID
			u.CallDirection = call.Direction
			selectedTalkingAt = call.CreatedAt
		}
	}
}

func (u *userCacheStruct) updateUserLocked(user *altStruct.DirectoryDomainUser) {
	user.SipRegister = u.SipRegister
	user.VertoRegister = u.VertoRegister
	user.CallDate = u.CallDate
	user.CCAgent = u.CCAgent
	user.Talking = u.Talking
	user.InCall = u.InCall
	user.CallDirection = u.CallDirection
	user.LastUuid = u.LastUuid
}

/*
func (f *userCache) GetByName(key string) *userCacheStruct {
	f.mx.RLock()
	defer f.mx.RUnlock()
	val := f.byName[key]
	return val
}
*/

func (f *userCache) SetByData(id int64, name string, reg bool) *userCacheStruct {
	var value = &userCacheStruct{Id: id, Name: name, SipRegister: reg, activeCalls: make(map[string]UserCallState)}
	f.mx.Lock()
	defer f.mx.Unlock()
	//f.byName[value.Name] = value
	f.byId[value.Id] = value
	return value
}

func (f *userCache) Set(value *userCacheStruct) {
	f.mx.Lock()
	defer f.mx.Unlock()
	//f.byName[value.Name] = value
	f.byId[value.Id] = value
}

func GetDirectoryCache() *directoryCacheStruct {
	return directoryCache
}
func InitCacheObjects() {
	fsInstances = mainStruct.NewFsInstances()
	directoryCache = NewFsDirectoryCache()
}

func SetCurrentInstanceName(name string) {
	currentFsInstanceName = name
}

func GetCurrentInstanceName() string {
	return currentFsInstanceName
}

func SetCurrentInstance() {
	currentFsInstance = fsInstances.GetByName(currentFsInstanceName)
}

func GetCurrentInstanceId() int64 {
	if currentFsInstance == nil {
		return 0
	}
	return currentFsInstance.Id
}

func GetFSInstances() *mainStruct.FsInstances {
	return fsInstances
}

func GetInstances() {
	UpdateCacheInstances()
	SetCurrentInstance()
	if currentFsInstance != nil {
		return
	}
	db.SetFSInstance(currentFsInstanceName)
	UpdateCacheInstances()
	SetCurrentInstance()
}

func UpdateCacheInstances() {
	db.GetFsInstances(fsInstances)
}

func InitCache() {
	InitCacheObjects()
	GetInstances()
}
