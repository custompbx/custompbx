package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

const (
	eventConferenceGet = "GetConference"

	eventConferenceRoomUpdate = "UpdateConferenceRoom"
	eventConferenceRoomSwitch = "SwitchConferenceRoom"
	eventConferenceRoomAdd    = "AddConferenceRoom"
	eventConferenceRoomDelete = "DelConferenceRoom"

	eventConferenceCallerControlsGet        = "GetConferenceCallerControls"
	eventConferenceCallerControlAdd         = "AddConferenceCallerControl"
	eventConferenceCallerControlDelete      = "DelConferenceCallerControl"
	eventConferenceCallerControlSwitch      = "SwitchConferenceCallerControl"
	eventConferenceCallerControlUpdate      = "UpdateConferenceCallerControl"
	eventConferenceCallerControlGroupAdd    = "AddConferenceCallerControlGroup"
	eventConferenceCallerControlGroupUpdate = "UpdateConferenceCallerControlGroup"
	eventConferenceCallerControlGroupDelete = "DelConferenceCallerControlGroup"
	eventConferenceProfileParametersGet     = "GetConferenceProfileParameters"
	eventConferenceProfileParameterAdd      = "AddConferenceProfileParameter"
	eventConferenceProfileParameterDelete   = "DelConferenceProfileParameter"
	eventConferenceProfileParameterSwitch   = "SwitchConferenceProfileParameter"
	eventConferenceProfileParameterUpdate   = "UpdateConferenceProfileParameter"
	eventConferenceProfileAdd               = "AddConferenceProfile"
	eventConferenceProfileUpdate            = "UpdateConferenceProfile"
	eventConferenceProfileDelete            = "DelConferenceProfile"
	eventConferenceChatPermissionUsersGet   = "GetConferenceChatPermissionUsers"
	eventConferenceChatPermissionUserAdd    = "AddConferenceChatPermissionUser"
	eventConferenceChatPermissionUserDelete = "DelConferenceChatPermissionUser"
	eventConferenceChatPermissionUserSwitch = "SwitchConferenceChatPermissionUser"
	eventConferenceChatPermissionUserUpdate = "UpdateConferenceChatPermissionUser"
	eventConferenceChatPermissionAdd        = "AddConferenceChatPermission"
	eventConferenceChatPermissionUpdate     = "UpdateConferenceChatPermission"
	eventConferenceChatPermissionDelete     = "DelConferenceChatPermission"
	eventConferenceLayoutsGet               = "GetConferenceLayouts"
	eventConferenceLayoutUpdate             = "UpdateConferenceLayout"
	eventConferenceLayout3DUpdate           = "UpdateConferenceLayout3D"
	eventConferenceLayoutSwitch             = "SwitchConferenceLayout"
	eventConferenceLayoutAdd                = "AddConferenceLayout"
	eventConferenceLayoutDelete             = "DelConferenceLayout"
	eventConferenceLayoutGroupUpdate        = "UpdateConferenceLayoutGroup"
	eventConferenceLayoutGroupSwitch        = "SwitchConferenceLayoutGroup"
	eventConferenceLayoutGroupAdd           = "AddConferenceLayoutGroup"
	eventConferenceLayoutGroupDelete        = "DelConferenceLayoutGroup"
	eventConferenceLayoutGroupLayoutsGet    = "GetConferenceLayoutGroupLayouts"
	eventConferenceLayoutGroupLayoutAdd     = "AddConferenceLayoutGroupLayout"
	eventConferenceLayoutGroupLayoutDelete  = "DelConferenceLayoutGroupLayout"
	eventConferenceLayoutGroupLayoutSwitch  = "SwitchConferenceLayoutGroupLayout"
	eventConferenceLayoutGroupLayoutUpdate  = "UpdateConferenceLayoutGroupLayout"
	eventConferenceLayoutImagesGet          = "GetConferenceLayoutImages"
	eventConferenceLayoutImageAdd           = "AddConferenceLayoutImage"
	eventConferenceLayoutImageDelete        = "DelConferenceLayoutImage"
	eventConferenceLayoutImageSwitch        = "SwitchConferenceLayoutImage"
	eventConferenceLayoutImageUpdate        = "UpdateConferenceLayoutImage"
)

func registerCoreConferenceEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventConferenceGet, getConferenceConfig, overrides)

	mustRegisterAdmin(r, eventConferenceRoomUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceAdvertiseRoom{Id: data.Param.Id, Name: data.Param.Name, Status: data.Param.Value}
	}, "Name", "Status"), overrides)
	mustRegisterAdmin(r, eventConferenceRoomSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceAdvertiseRoom{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceRoomAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceAdvertiseRoom{Name: data.Param.Name, Status: data.Param.Value, Enabled: true, Parent: configParentFor(&altStruct.ConfigConferenceAdvertiseRoom{})}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceRoomDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceAdvertiseRoom{Id: data.Param.Id}
	}), overrides)

	mustRegisterAdmin(r, eventConferenceCallerControlsGet, configGet(&altStruct.ConfigConferenceCallerControlGroupControl{}), overrides)
	mustRegisterAdmin(r, eventConferenceCallerControlAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceCallerControlGroupControl{Action: data.Param.Name, Digits: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceCallerControlGroup{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceCallerControlDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceCallerControlGroupControl{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceCallerControlSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceCallerControlGroupControl{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceCallerControlUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceCallerControlGroupControl{Id: data.Param.Id, Action: data.Param.Name, Digits: data.Param.Value}
	}, "Action", "Digits"), overrides)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventConferenceCallerControlGroupAdd, Update: eventConferenceCallerControlGroupUpdate, Delete: eventConferenceCallerControlGroupDelete},
		&altStruct.ConfigConferenceCallerControlGroup{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} {
			return configParentFor(&altStruct.ConfigConferenceCallerControlGroup{})
		},
	)

	mustRegisterAdmin(r, eventConferenceProfileParametersGet, configGet(&altStruct.ConfigConferenceProfileParameter{}), overrides)
	registerParentedParamConfigMutationsForSample(r, overrides,
		parentedParamConfigEvents{Add: eventConferenceProfileParameterAdd, Delete: eventConferenceProfileParameterDelete, Switch: eventConferenceProfileParameterSwitch, Update: eventConferenceProfileParameterUpdate},
		&altStruct.ConfigConferenceProfileParameter{},
		func(data *webStruct.MessageData) interface{} { return &altStruct.ConfigConferenceProfile{Id: data.Id} },
	)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventConferenceProfileAdd, Update: eventConferenceProfileUpdate, Delete: eventConferenceProfileDelete},
		&altStruct.ConfigConferenceProfile{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} {
			return configParentFor(&altStruct.ConfigConferenceProfile{})
		},
	)

	mustRegisterAdmin(r, eventConferenceChatPermissionUsersGet, configGet(&altStruct.ConfigConferenceChatPermissionProfileUser{}), overrides)
	mustRegisterAdmin(r, eventConferenceChatPermissionUserAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceChatPermissionProfileUser{Name: data.Param.Name, Commands: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigConferenceChatPermissionProfile{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceChatPermissionUserDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceChatPermissionProfileUser{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceChatPermissionUserSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceChatPermissionProfileUser{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceChatPermissionUserUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceChatPermissionProfileUser{Id: data.Param.Id, Name: data.Param.Name, Commands: data.Param.Value}
	}, "Name", "Commands"), overrides)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventConferenceChatPermissionAdd, Update: eventConferenceChatPermissionUpdate, Delete: eventConferenceChatPermissionDelete},
		&altStruct.ConfigConferenceChatPermissionProfile{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} {
			return configParentFor(&altStruct.ConfigConferenceChatPermissionProfile{})
		},
	)

	mustRegisterAdmin(r, eventConferenceLayoutsGet, getConferenceLayouts, overrides)
	mustRegisterAdmin(r, eventConferenceLayoutUpdate, configUpdateWithFields(&altStruct.ConfigConferenceLayout{}, []string{"Name"}, configDataIDName), overrides)
	mustRegisterAdmin(r, eventConferenceLayout3DUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayout{Id: data.Id, Auto3dPosition: data.Value}
	}, "Auto3dPosition"), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayout{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayout{Name: data.Name, Enabled: true, Parent: getConferenceLayoutConfig()}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutDelete, configDeleteWithFields(&altStruct.ConfigConferenceLayout{}, configGetNamedID), overrides)

	mustRegisterAdmin(r, eventConferenceLayoutGroupUpdate, configUpdateWithFields(&altStruct.ConfigConferenceLayoutGroup{}, []string{"Name"}, configDataIDName), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutGroup{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutGroup{Name: data.Name, Enabled: true, Parent: getConferenceLayoutConfig()}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupDelete, configDeleteWithFields(&altStruct.ConfigConferenceLayoutGroup{}, configGetNamedID), overrides)

	mustRegisterAdmin(r, eventConferenceLayoutGroupLayoutsGet, configGet(&altStruct.ConfigConferenceLayoutGroupLayout{}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupLayoutAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutGroupLayout{Body: data.Value, Enabled: true, Parent: &altStruct.ConfigConferenceLayoutGroup{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupLayoutDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutGroupLayout{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupLayoutSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutGroupLayout{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutGroupLayoutUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutGroupLayout{Id: data.Param.Id, Body: data.Param.Name}
	}, "Body"), overrides)

	mustRegisterAdmin(r, eventConferenceLayoutImagesGet, configGet(&altStruct.ConfigConferenceLayoutImage{}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutImageAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutImage{
			X:             data.LayoutImages.X,
			Y:             data.LayoutImages.Y,
			Scale:         data.LayoutImages.Scale,
			Floor:         data.LayoutImages.Floor,
			FloorOnly:     data.LayoutImages.FloorOnly,
			Hscale:        data.LayoutImages.Hscale,
			Overlap:       data.LayoutImages.Overlap,
			ReservationId: data.LayoutImages.ReservationId,
			Zoom:          data.LayoutImages.Zoom,
			Enabled:       true,
			Parent:        &altStruct.ConfigConferenceLayout{Id: data.Id},
		}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutImageDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutImage{Id: data.Param.Id}
	}), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutImageSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutImage{Id: data.Param.Id, Enabled: data.Param.Enabled}
	}, "Enabled"), overrides)
	mustRegisterAdmin(r, eventConferenceLayoutImageUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigConferenceLayoutImage{
			Id:            data.LayoutImages.Id,
			X:             data.LayoutImages.X,
			Y:             data.LayoutImages.Y,
			Scale:         data.LayoutImages.Scale,
			Floor:         data.LayoutImages.Floor,
			FloorOnly:     data.LayoutImages.FloorOnly,
			Hscale:        data.LayoutImages.Hscale,
			Overlap:       data.LayoutImages.Overlap,
			ReservationId: data.LayoutImages.ReservationId,
			Zoom:          data.LayoutImages.Zoom,
		}
	}, "X", "Y", "Scale", "Floor", "FloorOnly", "Hscale", "Overlap", "ReservationId", "Zoom"), overrides)
}

func getConferenceConfig(data *webStruct.MessageData) webStruct.UserResponse {
	rooms := getUserForConfig(data, getConfig, &altStruct.ConfigConferenceAdvertiseRoom{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigConferenceProfile{}, adminOnly())
	callerControlGroups := getUserForConfig(data, getConfig, &altStruct.ConfigConferenceCallerControlGroup{}, adminOnly())
	chatPermissionProfiles := getUserForConfig(data, getConfig, &altStruct.ConfigConferenceChatPermissionProfile{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "conference_rooms", data: rooms.Data},
		responseDataPair{name: "conference_profiles", data: profiles.Data},
		responseDataPair{name: "conference_caller_control_groups", data: callerControlGroups.Data},
		responseDataPair{name: "conference_chat_permissions_profiles", data: chatPermissionProfiles.Data},
	)
}

func getConferenceLayouts(data *webStruct.MessageData) webStruct.UserResponse {
	layoutData := *data
	layoutData.Name = "layout"
	layouts := getUserForConfig(&layoutData, getConferenceLayoutsConfig, &altStruct.ConfigConferenceLayout{}, adminOnly())

	groupData := *data
	groupData.Name = "group"
	groups := getUserForConfig(&groupData, getConferenceLayoutsConfig, &altStruct.ConfigConferenceLayoutGroup{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "conference_layouts", data: layouts.Data},
		responseDataPair{name: "conference_layouts_groups", data: groups.Data},
	)
}

func conferenceRegistryEvents() []string {
	return []string{
		eventConferenceGet,
		eventConferenceRoomUpdate,
		eventConferenceRoomSwitch,
		eventConferenceRoomAdd,
		eventConferenceRoomDelete,
		eventConferenceCallerControlsGet,
		eventConferenceCallerControlAdd,
		eventConferenceCallerControlDelete,
		eventConferenceCallerControlSwitch,
		eventConferenceCallerControlUpdate,
		eventConferenceCallerControlGroupAdd,
		eventConferenceCallerControlGroupUpdate,
		eventConferenceCallerControlGroupDelete,
		eventConferenceProfileParametersGet,
		eventConferenceProfileParameterAdd,
		eventConferenceProfileParameterDelete,
		eventConferenceProfileParameterSwitch,
		eventConferenceProfileParameterUpdate,
		eventConferenceProfileAdd,
		eventConferenceProfileUpdate,
		eventConferenceProfileDelete,
		eventConferenceChatPermissionUsersGet,
		eventConferenceChatPermissionUserAdd,
		eventConferenceChatPermissionUserDelete,
		eventConferenceChatPermissionUserSwitch,
		eventConferenceChatPermissionUserUpdate,
		eventConferenceChatPermissionAdd,
		eventConferenceChatPermissionUpdate,
		eventConferenceChatPermissionDelete,
		eventConferenceLayoutsGet,
		eventConferenceLayoutUpdate,
		eventConferenceLayout3DUpdate,
		eventConferenceLayoutSwitch,
		eventConferenceLayoutAdd,
		eventConferenceLayoutDelete,
		eventConferenceLayoutGroupUpdate,
		eventConferenceLayoutGroupSwitch,
		eventConferenceLayoutGroupAdd,
		eventConferenceLayoutGroupDelete,
		eventConferenceLayoutGroupLayoutsGet,
		eventConferenceLayoutGroupLayoutAdd,
		eventConferenceLayoutGroupLayoutDelete,
		eventConferenceLayoutGroupLayoutSwitch,
		eventConferenceLayoutGroupLayoutUpdate,
		eventConferenceLayoutImagesGet,
		eventConferenceLayoutImageAdd,
		eventConferenceLayoutImageDelete,
		eventConferenceLayoutImageSwitch,
		eventConferenceLayoutImageUpdate,
	}
}
