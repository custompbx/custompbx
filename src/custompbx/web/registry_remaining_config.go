package web

import (
	"custompbx/altStruct"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"reflect"
	"strconv"
)

const (
	eventCallcenterQueuesGet          = "GetCallcenterQueues"
	eventCallcenterQueueParamsGet     = "GetCallcenterQueuesParams"
	eventCallcenterSettingsGet        = "GetCallcenterSettings"
	eventCallcenterSettingsUpdate     = "UpdateCallcenterSettings"
	eventCallcenterSettingsSwitch     = "SwitchCallcenterSettings"
	eventCallcenterSettingsAdd        = "AddCallcenterSettings"
	eventCallcenterSettingsDelete     = "DelCallcenterSettings"
	eventCallcenterQueueParamAdd      = "AddCallcenterQueueParam"
	eventCallcenterQueueParamDelete   = "DelCallcenterQueueParam"
	eventCallcenterQueueParamSwitch   = "SwitchCallcenterQueueParam"
	eventCallcenterQueueParamUpdate   = "UpdateCallcenterQueueParam"
	eventCallcenterQueueAdd           = "AddCallcenterQueue"
	eventCallcenterQueueRename        = "RenameCallcenterQueue"
	eventCallcenterQueueDelete        = "DelCallcenterQueue"
	eventCallcenterAgentsAndTiersLoad = "ImportCallcenterAgentsAndTiers"
	eventCallcenterAgentsGet          = "GetCallcenterAgents"
	eventCallcenterAgentAdd           = "AddCallcenterAgent"
	eventCallcenterAgentUpdate        = "UpdateCallcenterAgent"
	eventCallcenterAgentDelete        = "DelCallcenterAgent"
	eventCallcenterTiersGet           = "GetCallcenterTiers"
	eventCallcenterTierAdd            = "AddCallcenterTier"
	eventCallcenterTierUpdate         = "UpdateCallcenterTier"
	eventCallcenterTierDelete         = "DelCallcenterTier"
	eventCallcenterMembersGet         = "GetCallcenterMembers"
	eventCallcenterMemberDelete       = "DelCallcenterMember"
	eventCallcenterCommandSend        = "SendCallcenterCommand"

	eventHttpCacheGet                 = "GetHttpCache"
	eventHttpCacheParamUpdate         = "UpdateHttpCacheParameter"
	eventHttpCacheParamSwitch         = "SwitchHttpCacheParameter"
	eventHttpCacheParamAdd            = "AddHttpCacheParameter"
	eventHttpCacheParamDelete         = "DelHttpCacheParameter"
	eventHttpCacheProfileGet          = "GetHttpCacheProfile"
	eventHttpCacheProfileAdd          = "AddHttpCacheProfile"
	eventHttpCacheProfileRename       = "RenameHttpCacheProfile"
	eventHttpCacheProfileDelete       = "DelHttpCacheProfile"
	eventHttpCacheProfileParamsGet    = "GetHttpCacheProfileParameters"
	eventHttpCacheProfileDomainAdd    = "AddHttpCacheProfileDomain"
	eventHttpCacheProfileDomainDelete = "DelHttpCacheProfileDomain"
	eventHttpCacheProfileDomainSwitch = "SwitchHttpCacheProfileDomain"
	eventHttpCacheProfileDomainUpdate = "UpdateHttpCacheProfileDomain"
	eventHttpCacheProfileAWSUpdate    = "UpdateHttpCacheProfileAws"
	eventHttpCacheProfileAzureUpdate  = "UpdateHttpCacheProfileAzure"

	eventDistributorGet        = "GetDistributorConfig"
	eventDistributorListAdd    = "AddDistributorList"
	eventDistributorListUpdate = "UpdateDistributorList"
	eventDistributorListDelete = "DelDistributorList"
	eventDistributorNodesGet   = "GetDistributorNodes"
	eventDistributorNodeAdd    = "AddDistributorNode"
	eventDistributorNodeDelete = "DelDistributorNode"
	eventDistributorNodeUpdate = "UpdateDistributorNode"
	eventDistributorNodeSwitch = "SwitchDistributorNode"

	eventPostLoadModulesGet   = "GetPostLoadModules"
	eventPostLoadModuleUpdate = "UpdatePostLoadModule"
	eventPostLoadModuleDelete = "DelPostLoadModule"

	eventVoicemailSettingsGet        = "GetVoicemailSettings"
	eventVoicemailSettingUpdate      = "UpdateVoicemailSetting"
	eventVoicemailSettingSwitch      = "SwitchVoicemailSetting"
	eventVoicemailSettingAdd         = "AddVoicemailSetting"
	eventVoicemailSettingDelete      = "DelVoicemailSetting"
	eventVoicemailProfilesGet        = "GetVoicemailProfiles"
	eventVoicemailProfileAdd         = "AddVoicemailProfile"
	eventVoicemailProfileUpdate      = "UpdateVoicemailProfile"
	eventVoicemailProfileDelete      = "DelVoicemailProfile"
	eventVoicemailProfileParamsGet   = "GetVoicemailProfileParameters"
	eventVoicemailProfileParamAdd    = "AddVoicemailProfileParameter"
	eventVoicemailProfileParamDelete = "DelVoicemailProfileParameter"
	eventVoicemailProfileParamSwitch = "SwitchVoicemailProfileParameter"
	eventVoicemailProfileParamUpdate = "UpdateVoicemailProfileParameter"

	eventGlobalVariablesGet    = "GetGlobalVariables"
	eventGlobalVariableUpdate  = "UpdateGlobalVariable"
	eventGlobalVariableSwitch  = "SwitchGlobalVariable"
	eventGlobalVariableAdd     = "AddGlobalVariable"
	eventGlobalVariableDelete  = "DelGlobalVariable"
	eventGlobalVariableMove    = "MoveGlobalVariable"
	eventGlobalVariablesImport = "ImportGlobalVariables"
)

func registerCoreRemainingConfigEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	registerCoreCallcenterEvents(r, overrides)
	registerCoreHttpCacheEvents(r, overrides)
	registerCoreDistributorEvents(r, overrides)
	registerCorePostLoadModuleEvents(r, overrides)
	registerCoreVoicemailEvents(r, overrides)
	registerCoreGlobalVariableEvents(r, overrides)
}

func registerCoreCallcenterEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventCallcenterQueuesGet, configGet(&altStruct.ConfigCallcenterQueue{}), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueParamsGet, configGet(&altStruct.ConfigCallcenterQueueParameter{}), overrides)
	mustRegisterAdmin(r, eventCallcenterSettingsGet, configGet(&altStruct.ConfigCallcenterSetting{}), overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventCallcenterSettingsUpdate, Switch: eventCallcenterSettingsSwitch, Add: eventCallcenterSettingsAdd, Delete: eventCallcenterSettingsDelete},
		&altStruct.ConfigCallcenterSetting{},
	)
	mustRegisterAdmin(r, eventCallcenterQueueParamAdd, configSetWithFields(&altStruct.ConfigCallcenterQueueParameter{}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Name": data.Param.Name, "Value": data.Param.Value, "Enabled": true, "Parent": &altStruct.ConfigCallcenterQueue{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueParamDelete, configDeleteWithFields(&altStruct.ConfigCallcenterQueueParameter{}, configGetParamID), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueParamSwitch, configUpdateWithFields(&altStruct.ConfigCallcenterQueueParameter{}, []string{"Enabled"}, configSwitchParamEnabled), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueParamUpdate, configUpdateWithFields(&altStruct.ConfigCallcenterQueueParameter{}, []string{"Name", "Value"}, configUpdateParamNameValue), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueAdd, configSetWithFields(&altStruct.ConfigCallcenterQueue{}, func(data *webStruct.MessageData) map[string]interface{} {
		return configSetTopLevelName(&altStruct.ConfigCallcenterQueue{}, data.Name)
	}), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueRename, configUpdateWithFields(&altStruct.ConfigCallcenterQueue{}, []string{"Name"}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Id": data.Id, "Name": data.Name}
	}), overrides)
	mustRegisterAdmin(r, eventCallcenterQueueDelete, configDeleteWithFields(&altStruct.ConfigCallcenterQueue{}, configGetNamedID), overrides)
	mustRegisterAdmin(r, eventCallcenterAgentsAndTiersLoad, importCallcenterAgentsAndTiers, overrides)
	mustRegisterAdmin(r, eventCallcenterAgentsGet, configGet(&altStruct.Agent{}), overrides)
	mustRegisterAdmin(r, eventCallcenterAgentAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.Agent{
			Name:              data.Name,
			Type:              "callback",
			System:            "single_box",
			InstanceId:        "single_box",
			Status:            "On Break",
			State:             "Waiting",
			WrapUpTime:        10,
			ReadyTime:         10,
			BusyDelayTime:     10,
			NoAnswerDelayTime: 10,
		}
	}), overrides)
	mustRegisterAdmin(r, eventCallcenterAgentUpdate, updateCallcenterAgent, overrides)
	mustRegisterAdmin(r, eventCallcenterAgentDelete, configDeleteWithFields(&altStruct.Agent{}, configGetNamedID), overrides)
	mustRegisterAdmin(r, eventCallcenterTiersGet, configGet(&altStruct.Tier{}), overrides)
	mustRegisterAdmin(r, eventCallcenterTierAdd, addCallcenterTier, overrides)
	mustRegisterAdmin(r, eventCallcenterTierUpdate, updateCallcenterTier, overrides)
	mustRegisterAdmin(r, eventCallcenterTierDelete, configDeleteWithFields(&altStruct.Tier{}, configGetNamedID), overrides)
	mustRegisterAdmin(r, eventCallcenterMembersGet, configGet(&altStruct.Member{}), overrides)
	mustRegisterAdmin(r, eventCallcenterMemberDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.Member{Uuid: data.Uuid}
	}), overrides)
	mustRegisterAdmin(r, eventCallcenterCommandSend, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, runCallcenterQueueCommand, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, webStruct.SubscribeCallcenterAgents, configGet(&altStruct.Agent{}), overrides)
	mustRegisterAdmin(r, webStruct.SubscribeCallcenterTiers, configGet(&altStruct.Tier{}), overrides)
}

func registerCoreHttpCacheEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventHttpCacheGet, getHttpCacheConfig, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventHttpCacheParamUpdate, Switch: eventHttpCacheParamSwitch, Add: eventHttpCacheParamAdd, Delete: eventHttpCacheParamDelete},
		&altStruct.ConfigHttpCacheSetting{},
	)
	mustRegisterAdmin(r, eventHttpCacheProfileGet, configGet(&altStruct.ConfigHttpCacheProfile{}), overrides)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventHttpCacheProfileAdd, Update: eventHttpCacheProfileRename, Delete: eventHttpCacheProfileDelete},
		&altStruct.ConfigHttpCacheProfile{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} {
			return configParentFor(&altStruct.ConfigHttpCacheProfile{})
		},
	)
	mustRegisterAdmin(r, eventHttpCacheProfileParamsGet, getHttpCacheProfileParameters, overrides)
	mustRegisterAdmin(r, eventHttpCacheProfileDomainAdd, configSetWithFields(&altStruct.ConfigHttpCacheProfileDomain{}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Name": data.Param.Name, "Enabled": true, "Parent": &altStruct.ConfigHttpCacheProfile{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventHttpCacheProfileDomainDelete, configDeleteWithFields(&altStruct.ConfigHttpCacheProfileDomain{}, configGetParamID), overrides)
	mustRegisterAdmin(r, eventHttpCacheProfileDomainSwitch, configUpdateWithFields(&altStruct.ConfigHttpCacheProfileDomain{}, []string{"Enabled"}, configSwitchParamEnabled), overrides)
	mustRegisterAdmin(r, eventHttpCacheProfileDomainUpdate, configUpdateWithFields(&altStruct.ConfigHttpCacheProfileDomain{}, []string{"Name"}, func(data *webStruct.MessageData) map[string]interface{} {
		return map[string]interface{}{"Id": data.Param.Id, "Name": data.Param.Name}
	}), overrides)
	mustRegisterAdmin(r, eventHttpCacheProfileAWSUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigHttpCacheProfileAWSS3{
			Id:              data.AwsS3.Id,
			AccessKeyId:     data.AwsS3.AccessKeyId,
			SecretAccessKey: data.AwsS3.SecretAccessKey,
			BaseDomain:      data.AwsS3.BaseDomain,
			Region:          data.AwsS3.Region,
			Expires:         data.AwsS3.Expires,
		}
	}, "AccessKeyId", "SecretAccessKey", "BaseDomain", "Region", "Expires"), overrides)
	mustRegisterAdmin(r, eventHttpCacheProfileAzureUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigHttpCacheProfileAzureBlob{Id: data.Azure.Id, SecretAccessKey: data.Azure.SecretAccessKey}
	}, "SecretAccessKey"), overrides)
}

func registerCoreDistributorEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventDistributorGet, configGet(&altStruct.ConfigDistributorList{}), overrides)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventDistributorListAdd, Update: eventDistributorListUpdate, Delete: eventDistributorListDelete},
		&altStruct.ConfigDistributorList{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} { return configParentFor(&altStruct.ConfigDistributorList{}) },
	)
	mustRegisterAdmin(r, eventDistributorNodesGet, configGet(&altStruct.ConfigDistributorListNode{}), overrides)
	mustRegisterAdmin(r, eventDistributorNodeAdd, configSet(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDistributorListNode{Name: data.DistributorNode.Name, Weight: data.DistributorNode.Weight, Enabled: true, Parent: &altStruct.ConfigDistributorList{Id: data.Id}}
	}), overrides)
	mustRegisterAdmin(r, eventDistributorNodeDelete, configDelete(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDistributorListNode{Id: data.DistributorNode.Id}
	}), overrides)
	mustRegisterAdmin(r, eventDistributorNodeUpdate, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDistributorListNode{Id: data.DistributorNode.Id, Name: data.DistributorNode.Name, Weight: data.DistributorNode.Weight}
	}, "Name", "Weight"), overrides)
	mustRegisterAdmin(r, eventDistributorNodeSwitch, configUpdate(func(data *webStruct.MessageData) interface{} {
		return &altStruct.ConfigDistributorListNode{Id: data.DistributorNode.Id, Enabled: data.DistributorNode.Enabled}
	}, "Enabled"), overrides)
}

func registerCorePostLoadModuleEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventPostLoadModulesGet, configGet(&altStruct.ConfigPostLoadModule{}), overrides)
	mustRegisterAdmin(r, eventPostLoadModuleUpdate, configUpdateWithFields(&altStruct.ConfigPostLoadModule{}, []string{"Name"}, configParamIDName), overrides)
	mustRegisterAdmin(r, webStruct.SwitchPostLoadModule, configUpdateWithFields(&altStruct.ConfigPostLoadModule{}, []string{"Enabled"}, configSwitchParamEnabled), overrides)
	mustRegisterAdmin(r, webStruct.AddPostLoadModule, configSetWithFields(&altStruct.ConfigPostLoadModule{}, func(data *webStruct.MessageData) map[string]interface{} {
		return configSetTopLevelName(&altStruct.ConfigPostLoadModule{}, data.Param.Name)
	}), overrides)
	mustRegisterAdmin(r, eventPostLoadModuleDelete, configDeleteWithFields(&altStruct.ConfigPostLoadModule{}, configGetParamID), overrides)
}

func registerCoreVoicemailEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventVoicemailSettingUpdate, Switch: eventVoicemailSettingSwitch, Add: eventVoicemailSettingAdd, Delete: eventVoicemailSettingDelete},
		&altStruct.ConfigVoicemailSetting{},
	)
	mustRegisterAdmin(r, eventVoicemailSettingsGet, configGet(&altStruct.ConfigVoicemailSetting{}), overrides)
	mustRegisterAdmin(r, eventVoicemailProfilesGet, configGet(&altStruct.ConfigVoicemailProfile{}), overrides)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventVoicemailProfileAdd, Update: eventVoicemailProfileUpdate, Delete: eventVoicemailProfileDelete},
		&altStruct.ConfigVoicemailProfile{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} {
			return configParentFor(&altStruct.ConfigVoicemailProfile{})
		},
	)
	mustRegisterAdmin(r, eventVoicemailProfileParamsGet, configGet(&altStruct.ConfigVoicemailProfileParameter{}), overrides)
	registerParentedParamConfigMutationsForSample(r, overrides,
		parentedParamConfigEvents{Add: eventVoicemailProfileParamAdd, Delete: eventVoicemailProfileParamDelete, Switch: eventVoicemailProfileParamSwitch, Update: eventVoicemailProfileParamUpdate},
		&altStruct.ConfigVoicemailProfileParameter{},
		func(data *webStruct.MessageData) interface{} { return &altStruct.ConfigVoicemailProfile{Id: data.Id} },
	)
}

func registerCoreGlobalVariableEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventGlobalVariablesGet, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, GetGlobalVariables, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, eventGlobalVariableUpdate, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, UpdateGlobalVariable, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, eventGlobalVariableSwitch, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, SwitchGlobalVariable, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, eventGlobalVariableAdd, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, AddGlobalVariable, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, eventGlobalVariableDelete, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, DelGlobalVariable, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, eventGlobalVariableMove, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, MoveGlobalVariable, adminOnly())
	}, overrides)
	mustRegisterAdmin(r, eventGlobalVariablesImport, func(data *webStruct.MessageData) webStruct.UserResponse {
		return getUser(data, ImportGlobalVariables, adminOnly())
	}, overrides)
}

func importCallcenterAgentsAndTiers(data *webStruct.MessageData) webStruct.UserResponse {
	getUser(data, ImportCallcenterAgentsAdnTiers, adminOnly())
	data.DBRequest = mainStruct.DBRequest{Limit: 25}
	agents := getUserForConfig(data, getConfig, &altStruct.Agent{}, adminOnly())
	tiers := getUserForConfig(data, getConfig, &altStruct.Tier{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "callcenter_agents", data: agents.Data},
		responseDataPair{name: "callcenter_tiers", data: tiers.Data},
	)
}

func addCallcenterTier(data *webStruct.MessageData) webStruct.UserResponse {
	queueI, err := intermediateDB.GetByIdArg(&altStruct.ConfigCallcenterQueue{}, data.Id)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	queue, ok := queueI.(altStruct.ConfigCallcenterQueue)
	if !ok {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}
	return getUserForConfig(data, setConfig, &altStruct.Tier{Queue: queue.Name, Agent: data.Name, State: "Ready", Position: 1, Level: 1}, adminOnly())
}

func updateCallcenterAgent(data *webStruct.MessageData) webStruct.UserResponse {
	payload, err := callcenterDynamicUpdatePayload(data, &altStruct.Agent{Id: data.Param.Id})
	if err != "" {
		return webStruct.UserResponse{Error: err, MessageType: data.Event}
	}
	return getUserForConfig(data, updateConfig, payload, adminOnly())
}

func updateCallcenterTier(data *webStruct.MessageData) webStruct.UserResponse {
	payload, err := callcenterDynamicUpdatePayload(data, &altStruct.Tier{Id: data.Param.Id})
	if err != "" {
		return webStruct.UserResponse{Error: err, MessageType: data.Event}
	}
	return getUserForConfig(data, updateConfig, payload, adminOnly())
}

func callcenterDynamicUpdatePayload(data *webStruct.MessageData, item interface{}) (interface{}, string) {
	if data.Param.Name == "" || data.Param.Value == "" {
		return nil, "wrong params"
	}
	fieldName := mainStruct.GetItemNameByTag(item, data.Param.Name)
	if data.Param.Name == "id" || fieldName == "id" || fieldName == "Id" {
		return nil, "please dont"
	}
	if fieldName == "" {
		return nil, "unknown field"
	}
	field := reflect.ValueOf(item).Elem().FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return nil, "unknown field"
	}
	switch field.Type().Name() {
	case "string":
		field.SetString(data.Param.Value)
	case "int", "int64":
		res, err := strconv.ParseInt(data.Param.Value, 10, 64)
		if err != nil {
			return nil, err.Error()
		}
		field.SetInt(res)
	case "bool":
		res, err := strconv.ParseBool(data.Param.Value)
		if err != nil {
			return nil, err.Error()
		}
		field.SetBool(res)
	default:
		return nil, "unsupported field type"
	}
	return struct {
		S interface{}
		A []string
	}{item, []string{fieldName}}, ""
}

func getHttpCacheConfig(data *webStruct.MessageData) webStruct.UserResponse {
	settings := getUserForConfig(data, getConfig, &altStruct.ConfigHttpCacheSetting{}, adminOnly())
	profiles := getUserForConfig(data, getConfig, &altStruct.ConfigHttpCacheProfile{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: settings.Data},
		responseDataPair{name: "profiles", data: profiles.Data},
	)
}

func getHttpCacheProfileParameters(data *webStruct.MessageData) webStruct.UserResponse {
	domains := getUserForConfig(data, getConfig, &altStruct.ConfigHttpCacheProfileDomain{}, adminOnly())
	azure := getUserForConfig(data, getConfig, &altStruct.ConfigHttpCacheProfileAzureBlob{}, adminOnly())
	awsS3 := getUserForConfig(data, getConfig, &altStruct.ConfigHttpCacheProfileAWSS3{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "domains", data: domains.Data},
		responseDataPair{name: "azure", data: azure.Data},
		responseDataPair{name: "aws_s3", data: awsS3.Data},
	)
}

func remainingConfigRegistryEvents() []string {
	return []string{
		eventCallcenterQueuesGet, eventCallcenterQueueParamsGet, eventCallcenterSettingsGet,
		eventCallcenterSettingsUpdate, eventCallcenterSettingsSwitch, eventCallcenterSettingsAdd, eventCallcenterSettingsDelete,
		eventCallcenterQueueParamAdd, eventCallcenterQueueParamDelete, eventCallcenterQueueParamSwitch, eventCallcenterQueueParamUpdate,
		eventCallcenterQueueAdd, eventCallcenterQueueRename, eventCallcenterQueueDelete,
		eventCallcenterAgentsAndTiersLoad, eventCallcenterAgentsGet, eventCallcenterAgentAdd, eventCallcenterAgentUpdate, eventCallcenterAgentDelete,
		eventCallcenterTiersGet, eventCallcenterTierAdd, eventCallcenterTierUpdate, eventCallcenterTierDelete,
		eventCallcenterMembersGet, eventCallcenterMemberDelete, eventCallcenterCommandSend,
		webStruct.SubscribeCallcenterAgents, webStruct.SubscribeCallcenterTiers,
		eventHttpCacheGet, eventHttpCacheParamUpdate, eventHttpCacheParamSwitch, eventHttpCacheParamAdd, eventHttpCacheParamDelete,
		eventHttpCacheProfileGet, eventHttpCacheProfileAdd, eventHttpCacheProfileRename, eventHttpCacheProfileDelete, eventHttpCacheProfileParamsGet,
		eventHttpCacheProfileDomainAdd, eventHttpCacheProfileDomainDelete, eventHttpCacheProfileDomainSwitch, eventHttpCacheProfileDomainUpdate,
		eventHttpCacheProfileAWSUpdate, eventHttpCacheProfileAzureUpdate,
		eventDistributorGet, eventDistributorListAdd, eventDistributorListUpdate, eventDistributorListDelete,
		eventDistributorNodesGet, eventDistributorNodeAdd, eventDistributorNodeDelete, eventDistributorNodeUpdate, eventDistributorNodeSwitch,
		eventPostLoadModulesGet, eventPostLoadModuleUpdate, webStruct.SwitchPostLoadModule, webStruct.AddPostLoadModule, eventPostLoadModuleDelete,
		eventVoicemailSettingsGet, eventVoicemailSettingUpdate, eventVoicemailSettingSwitch, eventVoicemailSettingAdd, eventVoicemailSettingDelete,
		eventVoicemailProfilesGet, eventVoicemailProfileAdd, eventVoicemailProfileUpdate, eventVoicemailProfileDelete,
		eventVoicemailProfileParamsGet, eventVoicemailProfileParamAdd, eventVoicemailProfileParamDelete, eventVoicemailProfileParamSwitch, eventVoicemailProfileParamUpdate,
		eventGlobalVariablesGet, eventGlobalVariableUpdate, eventGlobalVariableSwitch, eventGlobalVariableAdd, eventGlobalVariableDelete,
		eventGlobalVariableMove, eventGlobalVariablesImport,
	}
}
