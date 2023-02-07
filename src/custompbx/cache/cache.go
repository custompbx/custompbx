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

/*
func (f *userCache) GetByName(key string) *userCacheStruct {
	f.mx.RLock()
	defer f.mx.RUnlock()
	val := f.byName[key]
	return val
}
*/

func (f *userCache) SetByData(id int64, name string, reg bool) *userCacheStruct {
	var value = &userCacheStruct{Id: id, Name: name, SipRegister: reg}
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
