package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventVertoGet                = "[Config][Verto][Get]"
	eventVertoProfileParamsGet   = "[Config][Verto][Profile][Parameters][Get]"
	eventVertoSettingUpdate      = "[Config][Verto][Settings][Update]"
	eventVertoSettingSwitch      = "[Config][Verto][Setting][Switch]"
	eventVertoSettingAdd         = "[Config][Verto][Setting][Add]"
	eventVertoSettingDelete      = "[Config][Verto][Setting][Del]"
	eventVertoProfileParamAdd    = "[Config][Verto][Profile][Param][Add]"
	eventVertoProfileParamMove   = "MoveVertoProfileParameter"
	eventVertoProfileParamDelete = "[Config][Verto][Profile][Param][Del]"
	eventVertoProfileParamSwitch = "[Config][Verto][Profile][Param][Switch]"
	eventVertoProfileParamUpdate = "[Config][Verto][Profile][Update]"
	eventVertoProfileAdd         = "[Config][Verto][Profile][Add]"
	eventVertoProfileRename      = "[Config][Verto][Profile][Rename]"
	eventVertoProfileDelete      = "[Config][Verto][Profile][Del]"
)

func registerCoreVertoConfigEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventVertoGet, getVertoConfig, overrides)
	mustRegisterAdmin(r, eventVertoProfileParamsGet, configGet(&altStruct.ConfigVertoProfileParameter{}), overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventVertoSettingUpdate, Switch: eventVertoSettingSwitch, Add: eventVertoSettingAdd, Delete: eventVertoSettingDelete},
		&altStruct.ConfigVertoSetting{},
	)
	mustRegisterAdmin(r, eventVertoProfileParamAdd, configSetWithFields(&altStruct.ConfigVertoProfileParameter{}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Name": data.Param.Name, "Value": data.Param.Value, "Secure": data.Param.Secure, "Enabled": true, "Parent": &altStruct.ConfigVertoProfile{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventVertoProfileParamMove, configUpdateWithFields(&altStruct.ConfigVertoProfileParameter{}, []string{"Position"}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Id": data.Id, "Position": data.CurrentIndex}
	}), overrides)
	mustRegisterAdmin(r, eventVertoProfileParamDelete, configDeleteWithFields(&altStruct.ConfigVertoProfileParameter{}, configGetParamID), overrides)
	mustRegisterAdmin(r, eventVertoProfileParamSwitch, configUpdateWithFields(&altStruct.ConfigVertoProfileParameter{}, []string{"Enabled"}, configSwitchParamEnabled), overrides)
	mustRegisterAdmin(r, eventVertoProfileParamUpdate, configUpdateWithFields(&altStruct.ConfigVertoProfileParameter{}, []string{"Name", "Value", "secure"}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Id": data.Param.Id, "Name": data.Param.Name, "Value": data.Param.Value, "Secure": data.Param.Secure}
	}), overrides)
	mustRegisterAdmin(r, eventVertoProfileAdd, configSetWithFields(&altStruct.ConfigVertoProfile{}, func(data *webStruct.MessageData) map[string]interface{} {
		return configSetTopLevelName(&altStruct.ConfigVertoProfile{}, data.Name)
	}), overrides)
	mustRegisterAdmin(r, eventVertoProfileRename, configUpdateWithFields(&altStruct.ConfigVertoProfile{}, []string{"Name"}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Id": data.Id, "Name": data.Name}
	}), overrides)
	mustRegisterAdmin(r, eventVertoProfileDelete, configDeleteWithFields(&altStruct.ConfigVertoProfile{}, configGetNamedID), overrides)
}

func getVertoConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigVertoSetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigVertoProfile{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}

func vertoConfigRegistryEvents() []string {
	return []string{
		eventVertoGet,
		eventVertoProfileParamsGet,
		eventVertoSettingUpdate,
		eventVertoSettingSwitch,
		eventVertoSettingAdd,
		eventVertoSettingDelete,
		eventVertoProfileParamAdd,
		eventVertoProfileParamMove,
		eventVertoProfileParamDelete,
		eventVertoProfileParamSwitch,
		eventVertoProfileParamUpdate,
		eventVertoProfileAdd,
		eventVertoProfileRename,
		eventVertoProfileDelete,
	}
}
