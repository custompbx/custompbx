package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventLCRGet                = "GetLcr"
	eventLCRProfileParamsGet   = "GetLcrProfileParameters"
	eventLCRParamUpdate        = "UpdateLcrParameter"
	eventLCRParamSwitch        = "SwitchLcrParameter"
	eventLCRParamAdd           = "AddLcrParameter"
	eventLCRParamDelete        = "DelLcrParameter"
	eventLCRProfileParamAdd    = "AddLcrProfileParameter"
	eventLCRProfileParamDelete = "DelLcrProfileParameter"
	eventLCRProfileParamSwitch = "SwitchLcrProfileParameter"
	eventLCRProfileParamUpdate = "UpdateLcrProfileParameter"
	eventLCRProfileAdd         = "AddLcrProfile"
	eventLCRProfileUpdate      = "UpdateLcrProfile"
	eventLCRProfileDelete      = "DelLcrProfile"
)

func registerCoreLCREvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventLCRGet, getLCR, overrides)
	mustRegisterAdmin(r, eventLCRProfileParamsGet, configGet(&altStruct.ConfigLcrProfileParameter{}), overrides)
	mustRegisterAdmin(r, eventLCRParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, eventLCRParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventLCRParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigLcrSetting{})}
	}), overrides)
	mustRegisterAdmin(r, eventLCRParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrSetting{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventLCRProfileParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfileParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigLcrProfile{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventLCRProfileParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfileParameter{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventLCRProfileParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfileParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventLCRProfileParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfileParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, eventLCRProfileAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfile{Name: data.Name, Enabled: true, Parent: configParentFor(&altStruct.ConfigLcrProfile{})}
	}), overrides)
	mustRegisterAdmin(r, eventLCRProfileUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfile{Id: data.Id, Name: data.Name}
	}, "Name"), overrides)
	mustRegisterAdmin(r, eventLCRProfileDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigLcrProfile{Id: data.Id}
	}), overrides)
}

func getLCR(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigLcrSetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigLcrProfile{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}
