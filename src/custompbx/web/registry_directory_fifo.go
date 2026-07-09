package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventDirectoryGet                = "GetDirectory"
	eventDirectoryParamUpdate        = "UpdateDirectoryParameter"
	eventDirectoryParamSwitch        = "SwitchDirectoryParameter"
	eventDirectoryParamAdd           = "AddDirectoryParameter"
	eventDirectoryParamDelete        = "DelDirectoryParameter"
	eventDirectoryProfileParamsGet   = "GetDirectoryProfileParameters"
	eventDirectoryProfileParamAdd    = "AddDirectoryProfileParameter"
	eventDirectoryProfileParamDelete = "DelDirectoryProfileParameter"
	eventDirectoryProfileParamSwitch = "SwitchDirectoryProfileParameter"
	eventDirectoryProfileParamUpdate = "UpdateDirectoryProfileParameter"
	eventDirectoryProfileAdd         = "AddDirectoryProfile"
	eventDirectoryProfileUpdate      = "UpdateDirectoryProfile"
	eventDirectoryProfileDelete      = "DelDirectoryProfile"

	eventFifoGet              = "GetFifo"
	eventFifoParamUpdate      = "UpdateFifoParameter"
	eventFifoParamSwitch      = "SwitchFifoParameter"
	eventFifoParamAdd         = "AddFifoParameter"
	eventFifoParamDelete      = "DelFifoParameter"
	eventFifoMembersGet       = "GetFifoFifoMembers"
	eventFifoMemberAdd        = "AddFifoFifoMember"
	eventFifoMemberDelete     = "DelFifoFifoMember"
	eventFifoMemberSwitch     = "SwitchFifoFifoMember"
	eventFifoMemberUpdate     = "UpdateFifoFifoMember"
	eventFifoAdd              = "AddFifoFifo"
	eventFifoUpdate           = "UpdateFifoFifo"
	eventFifoDelete           = "DelFifoFifo"
	eventFifoImportanceUpdate = "UpdateFifoFifoImportance"
)

func registerCoreDirectoryConfigEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventDirectoryGet, getDirectoryConfig, overrides)

	// Preserve existing behavior: UpdateDirectoryParameter updates ConfigDirectoryProfileParameter.
	mustRegisterAdmin(r, eventDirectoryParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfileParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, eventDirectoryParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectorySetting{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventDirectoryParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectorySetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigDirectorySetting{})}
	}), overrides)
	mustRegisterAdmin(r, eventDirectoryParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectorySetting{Id: data.Param.Id}
	}), overrides)

	mustRegisterAdmin(r, eventDirectoryProfileParamsGet, configGet(&altStruct.ConfigDirectoryProfileParameter{}), overrides)
	mustRegisterAdmin(r, eventDirectoryProfileParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfileParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigDirectoryProfile{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventDirectoryProfileParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfileParameter{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventDirectoryProfileParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfileParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventDirectoryProfileParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfileParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)

	mustRegisterAdmin(r, eventDirectoryProfileAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfile{Name: data.Name, Enabled: true, Parent: configParentFor(&altStruct.ConfigDirectoryProfile{})}
	}), overrides)
	mustRegisterAdmin(r, eventDirectoryProfileUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfile{Id: data.Id, Name: data.Name}
	}, "Name"), overrides)
	mustRegisterAdmin(r, eventDirectoryProfileDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDirectoryProfile{Id: data.Id}
	}), overrides)
}

func registerCoreFifoEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventFifoGet, getFifoConfig, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventFifoParamUpdate, Switch: eventFifoParamSwitch, Add: eventFifoParamAdd, Delete: eventFifoParamDelete},
		&altStruct.ConfigFifoSetting{},
	)

	mustRegisterAdmin(r, eventFifoMembersGet, configGet(&altStruct.ConfigFifoFifoMember{}), overrides)
	mustRegisterAdmin(r, eventFifoMemberAdd, addFifoMember, overrides)
	mustRegisterAdmin(r, eventFifoMemberDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigFifoFifoMember{Id: data.FifoFifoMember.Id}
	}), overrides)
	mustRegisterAdmin(r, eventFifoMemberSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigFifoFifoMember{Id: data.FifoFifoMember.Id, Enabled: data.FifoFifoMember.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventFifoMemberUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigFifoFifoMember{
			Id:      data.FifoFifoMember.Id,
			Timeout: data.FifoFifoMember.Timeout,
			Simo:    data.FifoFifoMember.Simo,
			Lag:     data.FifoFifoMember.Lag,
			Body:    data.FifoFifoMember.Body,
		}
	}, "Timeout", "Simo", "Lag", "Body"), overrides)

	mustRegisterAdmin(r, eventFifoAdd, addFifo, overrides)
	mustRegisterAdmin(r, eventFifoUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigFifoFifo{Id: data.Id, Name: data.Name}
	}, "Name"), overrides)
	mustRegisterAdmin(r, eventFifoDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigFifoFifo{Id: data.Id}
	}), overrides)
	mustRegisterAdmin(r, eventFifoImportanceUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigFifoFifo{Id: data.Id, Importance: data.Value}
	}, "Importance"), overrides)
}

func getDirectoryConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigDirectorySetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigDirectoryProfile{}, adminOnly())

	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}

func getFifoConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigFifoSetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigFifoFifo{}, adminOnly())

	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}

func addFifoMember(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigFifoFifoMember{
		Id:      data.FifoFifoMember.Id,
		Timeout: data.FifoFifoMember.Timeout,
		Simo:    data.FifoFifoMember.Simo,
		Lag:     data.FifoFifoMember.Lag,
		Body:    data.FifoFifoMember.Body,
		Enabled: true,
		Parent:  &altStruct.ConfigFifoFifo{Id: data.Id},
	}, adminOnly())
}

func addFifo(data *webStruct.MessageData) webStruct.UserResponse {
	importance := data.Importance
	if importance == "" {
		importance = "0"
	}
	return getUserForConfig(data, setConfig, &altStruct.ConfigFifoFifo{
		Name:       data.Name,
		Importance: importance,
		Enabled:    true,
		Parent:     configParentFor(&altStruct.ConfigFifoFifo{}),
	}, adminOnly())
}

func directoryConfigRegistryEvents() []string {
	return []string{
		eventDirectoryGet,
		eventDirectoryParamUpdate,
		eventDirectoryParamSwitch,
		eventDirectoryParamAdd,
		eventDirectoryParamDelete,
		eventDirectoryProfileParamsGet,
		eventDirectoryProfileParamAdd,
		eventDirectoryProfileParamDelete,
		eventDirectoryProfileParamSwitch,
		eventDirectoryProfileParamUpdate,
		eventDirectoryProfileAdd,
		eventDirectoryProfileUpdate,
		eventDirectoryProfileDelete,
	}
}

func fifoRegistryEvents() []string {
	return []string{
		eventFifoGet,
		eventFifoParamUpdate,
		eventFifoParamSwitch,
		eventFifoParamAdd,
		eventFifoParamDelete,
		eventFifoMembersGet,
		eventFifoMemberAdd,
		eventFifoMemberDelete,
		eventFifoMemberSwitch,
		eventFifoMemberUpdate,
		eventFifoAdd,
		eventFifoUpdate,
		eventFifoDelete,
		eventFifoImportanceUpdate,
	}
}
