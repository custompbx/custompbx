package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventOpalGet                 = "GetOpal"
	eventOpalParamUpdate         = "UpdateOpalParameter"
	eventOpalParamSwitch         = "SwitchOpalParameter"
	eventOpalParamAdd            = "AddOpalParameter"
	eventOpalParamDelete         = "DelOpalParameter"
	eventOpalListenerParamsGet   = "GetOpalListenerParameters"
	eventOpalListenerParamAdd    = "AddOpalListenerParameter"
	eventOpalListenerParamDelete = "DelOpalListenerParameter"
	eventOpalListenerParamSwitch = "SwitchOpalListenerParameter"
	eventOpalListenerParamUpdate = "UpdateOpalListenerParameter"
	eventOpalListenerAdd         = "AddOpalListener"
	eventOpalListenerUpdate      = "UpdateOpalListener"
	eventOpalListenerDelete      = "DelOpalListener"

	eventOspGet                = "GetOsp"
	eventOspParamUpdate        = "UpdateOspParameter"
	eventOspParamSwitch        = "SwitchOspParameter"
	eventOspParamAdd           = "AddOspParameter"
	eventOspParamDelete        = "DelOspParameter"
	eventOspProfileParamsGet   = "GetOspProfileParameters"
	eventOspProfileParamAdd    = "AddOspProfileParameter"
	eventOspProfileParamDelete = "DelOspProfileParameter"
	eventOspProfileParamSwitch = "SwitchOspProfileParameter"
	eventOspProfileParamUpdate = "UpdateOspProfileParameter"
	eventOspProfileAdd         = "AddOspProfile"
	eventOspProfileUpdate      = "UpdateOspProfile"
	eventOspProfileDelete      = "DelOspProfile"

	eventUnicallGet             = "GetUnicall"
	eventUnicallParamUpdate     = "UpdateUnicallParameter"
	eventUnicallParamSwitch     = "SwitchUnicallParameter"
	eventUnicallParamAdd        = "AddUnicallParameter"
	eventUnicallParamDelete     = "DelUnicallParameter"
	eventUnicallSpanParamsGet   = "GetUnicallSpanParameters"
	eventUnicallSpanParamAdd    = "AddUnicallSpanParameter"
	eventUnicallSpanParamDelete = "DelUnicallSpanParameter"
	eventUnicallSpanParamSwitch = "SwitchUnicallSpanParameter"
	eventUnicallSpanParamUpdate = "UpdateUnicallSpanParameter"
	eventUnicallSpanAdd         = "AddUnicallSpan"
	eventUnicallSpanUpdate      = "UpdateUnicallSpan"
	eventUnicallSpanDelete      = "DelUnicallSpan"
)

func registerCoreTelephonyModuleEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	registerCoreOpalEvents(r, overrides)
	registerCoreOspEvents(r, overrides)
	registerCoreUnicallEvents(r, overrides)
}

func registerCoreOpalEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventOpalGet, getOpalConfig, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventOpalParamUpdate, Switch: eventOpalParamSwitch, Add: eventOpalParamAdd, Delete: eventOpalParamDelete},
		&altStruct.ConfigOpalSetting{},
	)
	mustRegisterAdmin(r, eventOpalListenerParamsGet, configGet(&altStruct.ConfigOpalListenerParameter{}), overrides)
	mustRegisterAdmin(r, eventOpalListenerParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListenerParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigOpalListener{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventOpalListenerParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListenerParameter{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventOpalListenerParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListenerParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventOpalListenerParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListenerParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, eventOpalListenerAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListener{Name: data.Name, Enabled: true, Parent: configParentFor(&altStruct.ConfigOpalListener{})}
	}), overrides)
	mustRegisterAdmin(r, eventOpalListenerUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListener{Id: data.Id, Name: data.Name}
	}, "Name"), overrides)
	mustRegisterAdmin(r, eventOpalListenerDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOpalListener{Id: data.Id}
	}), overrides)
}

func registerCoreOspEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventOspGet, getOspConfig, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventOspParamUpdate, Switch: eventOspParamSwitch, Add: eventOspParamAdd, Delete: eventOspParamDelete},
		&altStruct.ConfigOspSetting{},
	)
	mustRegisterAdmin(r, eventOspProfileParamsGet, configGet(&altStruct.ConfigOspProfileParameter{}), overrides)
	mustRegisterAdmin(r, eventOspProfileParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfileParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigOspProfile{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventOspProfileParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfileParameter{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventOspProfileParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfileParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventOspProfileParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfileParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, eventOspProfileAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfile{Name: data.Name, Enabled: true, Parent: configParentFor(&altStruct.ConfigOspProfile{})}
	}), overrides)
	mustRegisterAdmin(r, eventOspProfileUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfile{Id: data.Id, Name: data.Name}
	}, "Name"), overrides)
	mustRegisterAdmin(r, eventOspProfileDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigOspProfile{Id: data.Id}
	}), overrides)
}

func registerCoreUnicallEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventUnicallGet, getUnicallConfig, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventUnicallParamUpdate, Switch: eventUnicallParamSwitch, Add: eventUnicallParamAdd, Delete: eventUnicallParamDelete},
		&altStruct.ConfigUnicallSetting{},
	)
	mustRegisterAdmin(r, eventUnicallSpanParamsGet, configGet(&altStruct.ConfigUnicallSpanParameter{}), overrides)
	mustRegisterAdmin(r, eventUnicallSpanParamAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpanParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigUnicallSpan{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventUnicallSpanParamDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpanParameter{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventUnicallSpanParamSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpanParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventUnicallSpanParamUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpanParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}
	}, "Name", "Value"), overrides)
	mustRegisterAdmin(r, eventUnicallSpanAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpan{SpanId: data.Name, Enabled: true, Parent: configParentFor(&altStruct.ConfigUnicallSpan{})}
	}), overrides)
	mustRegisterAdmin(r, eventUnicallSpanUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpan{Id: data.Id, SpanId: data.Name}
	}, "SpanId"), overrides)
	mustRegisterAdmin(r, eventUnicallSpanDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigUnicallSpan{Id: data.Id}
	}), overrides)
}

func getOpalConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigOpalSetting{}, adminOnly())
	listeners := getUserForConfig(data, getConfig, &altStruct.ConfigOpalListener{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "listeners", data: listeners.Data},
	)
}

func getOspConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigOspSetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigOspProfile{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}

func getUnicallConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigUnicallSetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigUnicallSpan{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}

func telephonyModuleRegistryEvents() []string {
	return append(append(opalRegistryEvents(), ospRegistryEvents()...), unicallRegistryEvents()...)
}

func opalRegistryEvents() []string {
	return []string{
		eventOpalGet,
		eventOpalParamUpdate, eventOpalParamSwitch, eventOpalParamAdd, eventOpalParamDelete,
		eventOpalListenerParamsGet, eventOpalListenerParamAdd, eventOpalListenerParamDelete, eventOpalListenerParamSwitch, eventOpalListenerParamUpdate,
		eventOpalListenerAdd, eventOpalListenerUpdate, eventOpalListenerDelete,
	}
}

func ospRegistryEvents() []string {
	return []string{
		eventOspGet,
		eventOspParamUpdate, eventOspParamSwitch, eventOspParamAdd, eventOspParamDelete,
		eventOspProfileParamsGet, eventOspProfileParamAdd, eventOspProfileParamDelete, eventOspProfileParamSwitch, eventOspProfileParamUpdate,
		eventOspProfileAdd, eventOspProfileUpdate, eventOspProfileDelete,
	}
}

func unicallRegistryEvents() []string {
	return []string{
		eventUnicallGet,
		eventUnicallParamUpdate, eventUnicallParamSwitch, eventUnicallParamAdd, eventUnicallParamDelete,
		eventUnicallSpanParamsGet, eventUnicallSpanParamAdd, eventUnicallSpanParamDelete, eventUnicallSpanParamSwitch, eventUnicallSpanParamUpdate,
		eventUnicallSpanAdd, eventUnicallSpanUpdate, eventUnicallSpanDelete,
	}
}
