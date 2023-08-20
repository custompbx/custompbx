package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"custompbx/xmlStruct"
	"fmt"
	"github.com/custompbx/customorm"
	"log"
)

func getConfParent(name string) *altStruct.ConfigurationsList {
	log.Printf("%+v", name)
	conf, err := altData.GetModuleByName(mainStruct.GetModuleNameByConfName(name))
	if err != nil {
		return nil
	}
	return conf
}

func getConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	var parent interface{}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true}}
	if data.Id == 0 {
		parent = getConfParent(altData.GetConfNameByStruct(item))
		if parent == (*altStruct.ConfigurationsList)(nil) {
			exists := false
			return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
		}
	} else {
		parent = nil
		filter["Parent"] = customorm.FilterFields{Flag: true, UseValue: true, Value: data.Id}
	}

	module := altData.GetConfInstanceByStruct(item, parent)
	if module == nil {
		return webStruct.UserResponse{Error: "unknown config", MessageType: data.Event}
	}
	var res interface{}
	var err error
	if data.DBRequest.Limit != 0 || data.DBRequest.Filters != nil {
		limit := data.DBRequest.Limit
		if limit < 25 || limit > 250 {
			limit = 25
		}
		offset := 0
		if data.DBRequest.Offset > 0 {
			offset = data.DBRequest.Offset * limit
		}
		for _, v := range data.DBRequest.Filters {
			filter[v.Field] = customorm.FilterFields{Flag: true, UseValue: true, Value: v.FieldValue, Operand: v.Operand}
		}
		filterStr := customorm.Filters{
			Fields: filter,
			Limit:  limit,
			Offset: offset,
			Order:  customorm.Order{Desc: data.DBRequest.Order.Desc, Fields: data.DBRequest.Order.Fields}}
		res, err = intermediateDB.GetByFilteredValues(
			module,
			filterStr,
		)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
		}
		//TODO: with total all the time
		filterStr.Count = true
		resCount, err := intermediateDB.GetByFilteredValues(
			module,
			filterStr,
		)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
		}
		if len(resCount) == 0 {
			return webStruct.UserResponse{Error: "can't count total", MessageType: data.Event}
		}
		total, ok := resCount[0].(int64)
		if !ok {
			return webStruct.UserResponse{Error: "can't get total", MessageType: data.Event}
		}
		res = struct {
			Items interface{} `json:"items"`
			Total int64       `json:"total"`
		}{Items: res, Total: total}
	} else {
		res, err = intermediateDB.GetByValuesAsMap(
			module,
			filter,
		)
	}

	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: res, MessageType: data.Event}
}

func getConfigInnerParent(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true}}
	log.Printf("%+v", item)
	res, err := intermediateDB.GetByValuesAsMap(
		item,
		filter,
	)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &res, MessageType: data.Event}
}

func delConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	/*	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}*/
	res, err := intermediateDB.GetByIdFromDB(item)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	err = intermediateDB.DeleteById(item)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &res, MessageType: data.Event}
}

func setConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	log.Printf("%+v", item)
	res, err := intermediateDB.InsertItem(item)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	result, err := intermediateDB.GetByIdArg(item, res)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &result, MessageType: data.Event}
}

func updateConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	items, ok := item.(struct {
		S interface{}
		A []string
	})
	if !ok {
		return webStruct.UserResponse{Error: "no mandatory params", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(items.S, items.A)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	var returnList bool
	for _, v := range items.A {
		if v != "Position" {
			continue
		}
		returnList = true
		break
	}
	if returnList {
		return getConfigInnerParent(data, result)
	}

	return webStruct.UserResponse{Data: &result, MessageType: data.Event}
}

func GetConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	result := pbxcache.GetConfigs(res)

	return webStruct.UserResponse{MessageType: data.Event, Modules: result}
}

func reloadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Id: data.Id},
		map[string]bool{"Id": true},
	)
	if err != nil || len(res) == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	module, ok := res[0].(altStruct.ConfigurationsList)
	if !ok {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	comm := "reload " + mainStruct.GetModuleNameByConfName(module.Name)
	if mainStruct.GetModuleNameByConfName(module.Name) == mainStruct.ConfAcl {
		comm = "reloadacl"
	}
	_, err = fsesl.SendBgapiCmd(comm)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func unloadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Id: data.Id},
		map[string]bool{"Id": true},
	)
	if err != nil || len(res) == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	module, ok := res[0].(altStruct.ConfigurationsList)

	if !ok {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	comm := "unload " + mainStruct.GetModuleNameByConfName(module.Name)
	if mainStruct.GetModuleNameByConfName(module.Name) == mainStruct.ConfAcl {
		comm = "reloadacl"
	}
	_, err = fsesl.SendBgapiCmd(comm)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func loadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Id: data.Id},
		map[string]bool{"Id": true},
	)

	if err != nil || len(res) == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	module, ok := res[0].(altStruct.ConfigurationsList)

	if !ok {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	comm := "load " + mainStruct.GetModuleNameByConfName(module.Name)
	if mainStruct.GetModuleNameByConfName(module.Name) == mainStruct.ConfAcl {
		comm = "reloadacl"
	}
	_, err = fsesl.SendBgapiCmd(comm)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}
	return webStruct.UserResponse{MessageType: data.Event}
}

func switchConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 || data.Enabled == nil {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	err := intermediateDB.UpdateByIdByValuesMap(
		&altStruct.ConfigurationsList{Id: data.Id, Enabled: *data.Enabled},
		map[string]bool{"Enabled": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	result := pbxcache.GetConfigs(res)

	return webStruct.UserResponse{MessageType: data.Event, Modules: result}
}

func importConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	confName := data.Name + ".conf"
	err := fsesl.GetXMLModuleConfiguration(confName)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func TruncateModuleConfig(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	err := intermediateDB.DeleteById(&altStruct.ConfigurationsList{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	result := pbxcache.GetConfigs(res)
	return webStruct.UserResponse{MessageType: data.Event, Modules: result}
}

func ImportXMLModuleConfig(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.File == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	err := fsesl.ParseConfigXML(data.File)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	result := pbxcache.GetConfigs(res)

	return webStruct.UserResponse{MessageType: data.Event, Modules: result}
}

func importConfAllModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	fsesl.GetXMLConfigurations()

	return webStruct.UserResponse{MessageType: data.Event}
}

func fromScratchConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	confName := data.Name + ".conf"
	fsesl.InitConfigModule(&xmlStruct.Configuration{Attrname: confName})

	return webStruct.UserResponse{MessageType: data.Event}
}

func runProfileCommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profileI, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigSofiaProfile{Id: data.Id})
	if err != nil || profileI == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}
	profile, ok := profileI.(altStruct.ConfigSofiaProfile)
	if !ok {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}
	if data.Name == "" || !mainStruct.CheckCommand(data.Name) {
		return webStruct.UserResponse{Error: "unknown command", MessageType: data.Event}
	}

	if (data.Name == mainStruct.CommandSofiaProfileKillgw || data.Name == mainStruct.CommandSofiaProfileStartgw) && data.IdInt == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	gatewayName := ""
	if data.Name == mainStruct.CommandSofiaProfileKillgw || data.Name == mainStruct.CommandSofiaProfileStartgw {
		gatewayI, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigSofiaProfileGateway{Id: data.IdInt})
		if err != nil || gatewayI == nil {
			return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
		}
		gateway, ok := gatewayI.(altStruct.ConfigSofiaProfileGateway)
		if !ok {
			return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
		}
		gatewayName = gateway.Name
	}

	command := fmt.Sprintf("sofia profile %s %s %s", profile.Name, data.Name, gatewayName)

	_, err = fsesl.SendBgapiCmd(command)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

/*
	func ImportCallcenterAgentsAdnTiers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
		fsesl.GetCallcenterAgents()
		fsesl.GetCallcenterTiers()
		//fsesl.GetCallcenterMembers()

		resp := GetCallcenterAgents(data, user)
		tiers := GetCallcenterTiers(data, user)
		resp.CallcenterTiers = tiers.CallcenterTiers
		resp.AltTotal = tiers.Total
		return resp
	}
*/
func ImportCallcenterAgentsAdnTiers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	fsesl.GetCallcenterAgents()
	fsesl.GetCallcenterTiers()
	return webStruct.UserResponse{}
}

func runCallcenterQueueCommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	queueI, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigCallcenterQueue{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}
	queue, ok := queueI.(altStruct.ConfigCallcenterQueue)
	if !ok {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}
	if data.Name == "" || !mainStruct.CheckQueueCommand(data.Name) {
		return webStruct.UserResponse{Error: "unknown command", MessageType: data.Event}
	}

	command := fmt.Sprintf("callcenter_config queue %s %s", data.Name, queue.Name)

	_, err = fsesl.SendBgapiCmd(command)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}
