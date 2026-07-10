package web

import (
	"custompbx/altStruct"
	"custompbx/webStruct"
)

func registerCoreSofiaEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventSofiaGlobalSettingsGet, getSofiaGlobalSettings, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventSofiaGlobalSettingUpdate, Switch: eventSofiaGlobalSettingSwitch, Add: eventSofiaGlobalSettingAdd, Delete: eventSofiaGlobalSettingDelete},
		&altStruct.ConfigSofiaGlobalSetting{},
	)
	mustRegisterAdmin(r, webStruct.GetSofiaProfiles, getSofiaProfiles, overrides)
	mustRegisterAdmin(r, eventSofiaProfileParamsGet, getSofiaProfileParams, overrides)
	registerParentedParamConfigMutationsForSample(r, overrides,
		parentedParamConfigEvents{Add: eventSofiaProfileParamAdd, Delete: eventSofiaProfileParamDelete, Switch: eventSofiaProfileParamSwitch, Update: eventSofiaProfileParamUpdate},
		&altStruct.ConfigSofiaProfileParameter{},
		func(data *webStruct.MessageData) interface{} { return &altStruct.ConfigSofiaProfile{Id: data.Id} },
	)
	mustRegisterAdmin(r, eventSofiaProfileGatewaysGet, getSofiaProfileGateways, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayVarsGet, getSofiaGatewayVariables, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayParamsGet, getSofiaGatewayParameters, overrides)
	registerParentedParamConfigMutationsForSample(r, overrides,
		parentedParamConfigEvents{Add: eventSofiaGatewayParamAdd, Delete: eventSofiaGatewayParamDelete, Switch: eventSofiaGatewayParamSwitch, Update: eventSofiaGatewayParamUpdate},
		&altStruct.ConfigSofiaProfileGatewayParameter{},
		func(data *webStruct.MessageData) interface{} {
			return &altStruct.ConfigSofiaProfileGateway{Id: data.Id}
		},
	)
	mustRegisterAdmin(r, eventSofiaGatewayVarAdd, addSofiaGatewayVar, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayVarUpdate, updateSofiaGatewayVar, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayVarSwitch, switchSofiaGatewayVar, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayVarDelete, deleteSofiaGatewayVar, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayAdd, addSofiaGateway, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayDelete, deleteSofiaGateway, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayRename, renameSofiaGateway, overrides)
	mustRegisterAdmin(r, eventSofiaProfileDomainsGet, getSofiaProfileDomains, overrides)
	mustRegisterAdmin(r, eventSofiaProfileDomainAdd, addSofiaProfileDomain, overrides)
	mustRegisterAdmin(r, eventSofiaProfileDomainDelete, deleteSofiaProfileDomain, overrides)
	mustRegisterAdmin(r, eventSofiaProfileDomainSwitch, switchSofiaProfileDomain, overrides)
	mustRegisterAdmin(r, eventSofiaProfileDomainUpdate, updateSofiaProfileDomain, overrides)
	mustRegisterAdmin(r, eventSofiaProfileAliasesGet, getSofiaProfileAliases, overrides)
	mustRegisterAdmin(r, eventSofiaProfileAliasAdd, addSofiaProfileAlias, overrides)
	mustRegisterAdmin(r, eventSofiaProfileAliasDelete, deleteSofiaProfileAlias, overrides)
	mustRegisterAdmin(r, eventSofiaProfileAliasSwitch, switchSofiaProfileAlias, overrides)
	mustRegisterAdmin(r, eventSofiaProfileAliasUpdate, updateSofiaProfileAlias, overrides)
	registerNamedConfigMutationsForSample(r, overrides,
		namedConfigEvents{Add: eventSofiaProfileAdd, Update: eventSofiaProfileRename, Delete: eventSofiaProfileDelete},
		&altStruct.ConfigSofiaProfile{},
		func(data *webStruct.MessageData) string { return data.Name },
		func(_ *webStruct.MessageData) interface{} { return configParentFor(&altStruct.ConfigSofiaProfile{}) },
	)
	mustRegisterAdmin(r, eventSofiaProfileCommand, runProfileCommand, overrides)
	mustRegisterAdmin(r, eventSofiaProfileSwitch, switchSofiaProfile, overrides)
}

func getSofiaGlobalSettings(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaGlobalSetting{}, adminOnly())
}

func getSofiaProfiles(data *webStruct.MessageData) webStruct.UserResponse {
	return setProfileStatuses(getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfile{}, adminOnly()))
}

func getSofiaProfileParams(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileParameter{}, adminOnly())
}

func getSofiaProfileGateways(data *webStruct.MessageData) webStruct.UserResponse {
	return setGatewayStatuses(getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileGateway{}, adminOnly()))
}

func getSofiaGatewayVariables(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileGatewayVariable{}, adminOnly())
}

func getSofiaGatewayParameters(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileGatewayParameter{}, adminOnly())
}

func addSofiaGatewayVar(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfileGatewayVariable{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigSofiaProfileGateway{Id: data.Id}}, adminOnly())
}

func updateSofiaGatewayVar(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileGatewayVariable{Id: data.Variable.Id, Name: data.Variable.Name, Value: data.Variable.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchSofiaGatewayVar(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileGatewayVariable{Id: data.Variable.Id, Enabled: data.Variable.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteSofiaGatewayVar(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfileGatewayVariable{Id: data.Variable.Id}, adminOnly())
}

func addSofiaGateway(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfileGateway{Name: data.Name, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: data.Id}}, adminOnly())
}

func deleteSofiaGateway(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfileGateway{Id: data.Id}, adminOnly())
}

func renameSofiaGateway(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileGateway{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func getSofiaProfileDomains(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileDomain{}, adminOnly())
}

func addSofiaProfileDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfileDomain{Name: data.SofiaDomain.Name, Alias: data.SofiaDomain.Alias, Parse: data.SofiaDomain.Parse, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: data.Id}}, adminOnly())
}

func deleteSofiaProfileDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfileDomain{Id: data.SofiaDomain.Id}, adminOnly())
}

func switchSofiaProfileDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileDomain{Id: data.SofiaDomain.Id, Enabled: data.SofiaDomain.Enabled}, []string{"Enabled"}}, adminOnly())
}

func updateSofiaProfileDomain(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileDomain{Id: data.SofiaDomain.Id, Name: data.SofiaDomain.Name, Alias: data.SofiaDomain.Alias, Parse: data.SofiaDomain.Parse}, []string{"Name", "Alias", "Parse"}}, adminOnly())
}

func getSofiaProfileAliases(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileAlias{}, adminOnly())
}

func addSofiaProfileAlias(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfileAlias{Name: data.SofiaAlias.Name, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: data.Id}}, adminOnly())
}

func deleteSofiaProfileAlias(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfileAlias{Id: data.SofiaAlias.Id}, adminOnly())
}

func switchSofiaProfileAlias(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileAlias{Id: data.SofiaAlias.Id, Enabled: data.SofiaAlias.Enabled}, []string{"Enabled"}}, adminOnly())
}

func updateSofiaProfileAlias(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileAlias{Id: data.SofiaAlias.Id, Name: data.SofiaAlias.Name}, []string{"Name"}}, adminOnly())
}

func switchSofiaProfile(data *webStruct.MessageData) webStruct.UserResponse {
	return setProfileStatuses(getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfile{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly()))
}
