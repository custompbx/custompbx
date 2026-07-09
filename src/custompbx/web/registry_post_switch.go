package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventPostSwitchGet                 = "GetPostSwitch"
	eventPostSwitchParamUpdate         = "UpdatePostSwitchParameter"
	eventPostSwitchParamSwitch         = "SwitchPostSwitchParameter"
	eventPostSwitchParamAdd            = "AddPostSwitchParameter"
	eventPostSwitchParamDelete         = "DelPostSwitchParameter"
	eventPostSwitchCliKeybindingUpdate = "UpdatePostSwitchCliKeybinding"
	eventPostSwitchCliKeybindingSwitch = "SwitchPostSwitchCliKeybinding"
	eventPostSwitchCliKeybindingAdd    = "AddPostSwitchCliKeybinding"
	eventPostSwitchCliKeybindingDelete = "DelPostSwitchCliKeybinding"
	eventPostSwitchDefaultPtimeUpdate  = "UpdatePostSwitchDefaultPtime"
	eventPostSwitchDefaultPtimeSwitch  = "SwitchPostSwitchDefaultPtime"
	eventPostSwitchDefaultPtimeAdd     = "AddPostSwitchDefaultPtime"
	eventPostSwitchDefaultPtimeDelete  = "DelPostSwitchDefaultPtime"
)

func registerCorePostSwitchEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventPostSwitchGet, getPostSwitchConfig, overrides)

	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventPostSwitchParamUpdate, Switch: eventPostSwitchParamSwitch, Add: eventPostSwitchParamAdd, Delete: eventPostSwitchParamDelete},
		&altStruct.ConfigPostLoadSwitchSetting{},
	)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventPostSwitchCliKeybindingUpdate, Switch: eventPostSwitchCliKeybindingSwitch, Add: eventPostSwitchCliKeybindingAdd, Delete: eventPostSwitchCliKeybindingDelete},
		&altStruct.ConfigPostLoadSwitchCliKeybinding{},
	)

	mustRegisterAdmin(r, eventPostSwitchDefaultPtimeUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigPostLoadSwitchDefaultPtime{
			Id:         data.Param.Id,
			CodecName:  data.Param.Name,
			CodecPtime: data.Param.Value,
		}
	}, "Name", "CodecName", "CodecPtime"), overrides)
	mustRegisterAdmin(r, eventPostSwitchDefaultPtimeSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigPostLoadSwitchDefaultPtime{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventPostSwitchDefaultPtimeAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigPostLoadSwitchDefaultPtime{
			CodecName:  data.Param.Name,
			CodecPtime: data.Param.Value,
			Enabled:    true,
			Parent:     configParentFor(&altStruct.ConfigPostLoadSwitchDefaultPtime{}),
		}
	}), overrides)
	mustRegisterAdmin(r, eventPostSwitchDefaultPtimeDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigPostLoadSwitchDefaultPtime{Id: data.Param.Id}
	}), overrides)
}

func getPostSwitchConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigPostLoadSwitchSetting{}, adminOnly())
	keybindings := getUserForConfig(data, getConfig, &altStruct.ConfigPostLoadSwitchCliKeybinding{}, adminOnly())
	defaultPtimes := getUserForConfig(data, getConfig, &altStruct.ConfigPostLoadSwitchDefaultPtime{}, adminOnly())

	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "cli_keybinding", data: keybindings.Data},
		responseDataPair{name: "default_ptime", data: defaultPtimes.Data},
	)
}

func postSwitchRegistryEvents() []string {
	return []string{
		eventPostSwitchGet,
		eventPostSwitchParamUpdate,
		eventPostSwitchParamSwitch,
		eventPostSwitchParamAdd,
		eventPostSwitchParamDelete,
		eventPostSwitchCliKeybindingUpdate,
		eventPostSwitchCliKeybindingSwitch,
		eventPostSwitchCliKeybindingAdd,
		eventPostSwitchCliKeybindingDelete,
		eventPostSwitchDefaultPtimeUpdate,
		eventPostSwitchDefaultPtimeSwitch,
		eventPostSwitchDefaultPtimeAdd,
		eventPostSwitchDefaultPtimeDelete,
	}
}
