package mainStruct

import (
	"encoding/xml"
	"sort"
	"sync"
)

const (
	dynamicGlobalVariableHostname       = "hostname"
	dynamicGlobalVariableLocalIpV4      = "local_ip_v4"
	dynamicGlobalVariableLocalMaskV4    = "local_mask_v4"
	dynamicGlobalVariableLocalIpV6      = "local_ip_v6"
	dynamicGlobalVariableSwitchSerial   = "switch_serial"
	dynamicGlobalVariableBaseDir        = "base_dir"
	dynamicGlobalVariableRecordingsDir  = "recordings_dir"
	dynamicGlobalVariableSoundPrefix    = "sound_prefix"
	dynamicGlobalVariableSoundsDir      = "sounds_dir"
	dynamicGlobalVariableConfDir        = "conf_dir"
	dynamicGlobalVariableLogDir         = "log_dir"
	dynamicGlobalVariableRunDir         = "run_dir"
	dynamicGlobalVariableDbDir          = "db_dir"
	dynamicGlobalVariableModDir         = "mod_dir"
	dynamicGlobalVariableHtdocsDir      = "htdocs_dir"
	dynamicGlobalVariableScriptDir      = "script_dir"
	dynamicGlobalVariableTempDir        = "temp_dir"
	dynamicGlobalVariableGrammarDir     = "grammar_dir"
	dynamicGlobalVariableCertsDir       = "certs_dir"
	dynamicGlobalVariableStorageDir     = "storage_dir"
	dynamicGlobalVariableCacheDir       = "cache_dir"
	dynamicGlobalVariableCoreUuid       = "core_uuid"
	dynamicGlobalVariableZrtpEnabled    = "zrtp_enabled"
	dynamicGlobalVariableNatPublicAddr  = "nat_public_addr"
	dynamicGlobalVariableNatPrivateAddr = "nat_private_addr"
	dynamicGlobalVariableNatType        = "nat_type"
)

type X_PRE_PROCESS struct {
	XMLName      xml.Name `xml:"X-PRE-PROCESS,omitempty" json:"X-PRE-PROCESS,omitempty"`
	Attrcmd      string   `xml:"cmd,attr"  json:",omitempty"`
	Attrdata     string   `xml:"data,attr"  json:",omitempty"`
	Attrmetatype string   `xml:"meta_type,attr"  json:",omitempty"`
}
type Includer struct {
	XMLName  xml.Name         `xml:"include,omitempty" json:",omitempty"`
	Includes []*X_PRE_PROCESS `xml:"X-PRE-PROCESS,omitempty" json:"X-PRE-PROCESS,omitempty"`
}

func IsDynamicGlobalVar(name string) bool {
	switch name {
	case dynamicGlobalVariableHostname:
		fallthrough
	case dynamicGlobalVariableLocalIpV4:
		fallthrough
	case dynamicGlobalVariableLocalMaskV4:
		fallthrough
	case dynamicGlobalVariableLocalIpV6:
		fallthrough
	case dynamicGlobalVariableSwitchSerial:
		fallthrough
	case dynamicGlobalVariableBaseDir:
		fallthrough
	case dynamicGlobalVariableRecordingsDir:
		fallthrough
	case dynamicGlobalVariableSoundPrefix:
		fallthrough
	case dynamicGlobalVariableSoundsDir:
		fallthrough
	case dynamicGlobalVariableConfDir:
		fallthrough
	case dynamicGlobalVariableLogDir:
		fallthrough
	case dynamicGlobalVariableRunDir:
		fallthrough
	case dynamicGlobalVariableDbDir:
		fallthrough
	case dynamicGlobalVariableModDir:
		fallthrough
	case dynamicGlobalVariableHtdocsDir:
		fallthrough
	case dynamicGlobalVariableScriptDir:
		fallthrough
	case dynamicGlobalVariableTempDir:
		fallthrough
	case dynamicGlobalVariableGrammarDir:
		fallthrough
	case dynamicGlobalVariableCertsDir:
		fallthrough
	case dynamicGlobalVariableStorageDir:
		fallthrough
	case dynamicGlobalVariableCacheDir:
		fallthrough
	case dynamicGlobalVariableCoreUuid:
		fallthrough
	case dynamicGlobalVariableZrtpEnabled:
		fallthrough
	case dynamicGlobalVariableNatPublicAddr:
		fallthrough
	case dynamicGlobalVariableNatPrivateAddr:
		fallthrough
	case dynamicGlobalVariableNatType:
		return true
	}
	return false
}

type GlobalVariable struct {
	Id       int64  `xml:"-" json:"id"`
	Enabled  bool   `xml:"-" json:"enabled"`
	Dynamic  bool   `xml:"-" json:"dynamic"`
	Name     string `xml:"-" json:"name"`
	Value    string `xml:"-" json:"value"`
	Type     string `xml:"-" json:"type"`
	Position int64  `xml:"-" json:"position"`
}

type GlobalVariables struct {
	mx     sync.RWMutex
	byName map[string]*GlobalVariable
	byId   map[int64]*GlobalVariable
}

func NewGlobalVariables() *GlobalVariables {
	return &GlobalVariables{
		byName: make(map[string]*GlobalVariable),
		byId:   make(map[int64]*GlobalVariable),
	}
}

func (g *GlobalVariables) GetByName(key string) *GlobalVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val := g.byName[key]
	return val
}

func (g *GlobalVariables) GetById(key int64) *GlobalVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	val := g.byId[key]
	return val
}

func (g *GlobalVariables) Remove(key *GlobalVariable) {
	g.mx.Lock()
	defer g.mx.Unlock()
	delete(g.byName, key.Name)
	delete(g.byId, key.Id)
}

func (g *GlobalVariables) Set(value *GlobalVariable) {
	g.mx.Lock()
	defer g.mx.Unlock()
	g.byName[value.Name] = value
	g.byId[value.Id] = value
}

func (g *GlobalVariables) GetList() map[int64]*GlobalVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[int64]*GlobalVariable)
	for _, v := range g.byId {
		list[v.Id] = v
	}
	return list
}

func (g *GlobalVariables) Props() []*GlobalVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	var items []*GlobalVariable
	for _, v := range g.byId {
		items = append(items, v)
	}
	return items
}

func (g *GlobalVariables) NotDynamicsProps() []*GlobalVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	var items []*GlobalVariable
	for _, v := range g.byId {
		if v.Dynamic {
			continue
		}
		items = append(items, v)
	}
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Position < items[j].Position
	})
	return items
}

func (g *GlobalVariables) GetNamedList() map[string]*GlobalVariable {
	g.mx.RLock()
	defer g.mx.RUnlock()
	list := make(map[string]*GlobalVariable)
	for _, v := range g.byName {
		list[v.Name] = v
	}
	return list
}
