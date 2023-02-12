package altWeb

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"encoding/json"
	"log"
)

func GetPostSwitch(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	config, err := altData.GetCurrentConfPostSwitch()
	if err != nil {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}
	params, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConfigPostLoadSwitchSetting{Parent: &altStruct.ConfigurationsList{Id: config.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		var exists bool
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	keyBindings, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConfigPostLoadSwitchCliKeybinding{Parent: &altStruct.ConfigurationsList{Id: config.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		var exists bool
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	ptimes, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConfigPostLoadSwitchDefaultPtime{Parent: &altStruct.ConfigurationsList{Id: config.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		var exists bool
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	item := struct {
		PostSwitchParameters    interface{} `json:"post_switch_parameters"`
		PostSwitchCliKeybinding interface{} `json:"post_switch_cli_keybinding"`
		PostSwitchDefaultPtime  interface{} `json:"post_switch_default_ptime"`
	}{
		PostSwitchParameters:    &params,
		PostSwitchCliKeybinding: &keyBindings,
		PostSwitchDefaultPtime:  &ptimes,
	}

	return webStruct.UserResponse{Data: item, MessageType: data.Event}
}

func GetPostSwitchSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	config, err := altData.GetCurrentConfPostSwitch()
	if err != nil {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}
	params, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConfigPostLoadSwitchSetting{Parent: &altStruct.ConfigurationsList{Id: config.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		var exists bool
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &params, MessageType: data.Event}
}

func GetPostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	config, err := altData.GetCurrentConfPostSwitch()
	if err != nil {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}

	keyBindings, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConfigPostLoadSwitchCliKeybinding{Parent: &altStruct.ConfigurationsList{Id: config.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		var exists bool
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &keyBindings, MessageType: data.Event}
}

func GetPostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	config, err := altData.GetCurrentConfPostSwitch()
	if err != nil {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}

	ptimes, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConfigPostLoadSwitchDefaultPtime{Parent: &altStruct.ConfigurationsList{Id: config.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		var exists bool
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &ptimes, MessageType: data.Event}
}

func UpdatePostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := altStruct.ConfigPostLoadSwitchSetting{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func AddPostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := altStruct.ConfigPostLoadSwitchSetting{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelPostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := altStruct.ConfigPostLoadSwitchSetting{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdatePostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := altStruct.ConfigPostLoadSwitchCliKeybinding{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func AddPostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := altStruct.ConfigPostLoadSwitchCliKeybinding{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelPostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := altStruct.ConfigPostLoadSwitchCliKeybinding{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdatePostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := altStruct.ConfigPostLoadSwitchDefaultPtime{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func AddPostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := altStruct.ConfigPostLoadSwitchDefaultPtime{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelPostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchDefaultPtime(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPostSwitchDefaultPtime(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}
