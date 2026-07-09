package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/webStruct"
)

func registerCoreSofiaEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventSofiaGlobalSettingsGet, getSofiaGlobalSettings, overrides)
	mustRegisterAdmin(r, eventSofiaGlobalSettingUpdate, updateSofiaGlobalSetting, overrides)
	mustRegisterAdmin(r, eventSofiaGlobalSettingSwitch, switchSofiaGlobalSetting, overrides)
	mustRegisterAdmin(r, eventSofiaGlobalSettingAdd, addSofiaGlobalSetting, overrides)
	mustRegisterAdmin(r, eventSofiaGlobalSettingDelete, deleteSofiaGlobalSetting, overrides)
	mustRegisterAdmin(r, webStruct.GetSofiaProfiles, getSofiaProfiles, overrides)
	mustRegisterAdmin(r, eventSofiaProfileParamsGet, getSofiaProfileParams, overrides)
	mustRegisterAdmin(r, eventSofiaProfileParamAdd, addSofiaProfileParam, overrides)
	mustRegisterAdmin(r, eventSofiaProfileParamDelete, deleteSofiaProfileParam, overrides)
	mustRegisterAdmin(r, eventSofiaProfileParamSwitch, switchSofiaProfileParam, overrides)
	mustRegisterAdmin(r, eventSofiaProfileParamUpdate, updateSofiaProfileParam, overrides)
	mustRegisterAdmin(r, eventSofiaProfileGatewaysGet, getSofiaProfileGateways, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayVarsGet, getSofiaGatewayVariables, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayParamsGet, getSofiaGatewayParameters, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayParamAdd, addSofiaGatewayParam, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayParamUpdate, updateSofiaGatewayParam, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayParamSwitch, switchSofiaGatewayParam, overrides)
	mustRegisterAdmin(r, eventSofiaGatewayParamDelete, deleteSofiaGatewayParam, overrides)
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
	mustRegisterAdmin(r, eventSofiaProfileAdd, addSofiaProfile, overrides)
	mustRegisterAdmin(r, eventSofiaProfileRename, renameSofiaProfile, overrides)
	mustRegisterAdmin(r, eventSofiaProfileDelete, deleteSofiaProfile, overrides)
	mustRegisterAdmin(r, eventSofiaProfileCommand, runProfileCommand, overrides)
	mustRegisterAdmin(r, eventSofiaProfileSwitch, switchSofiaProfile, overrides)
}

func getSofiaGlobalSettings(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaGlobalSetting{}, adminOnly())
}

func updateSofiaGlobalSetting(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaGlobalSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchSofiaGlobalSetting(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaGlobalSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}, []string{"Enabled"}}, adminOnly())
}

func addSofiaGlobalSetting(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaGlobalSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigSofiaGlobalSetting{}))}, adminOnly())
}

func deleteSofiaGlobalSetting(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaGlobalSetting{Id: data.Param.Id}, adminOnly())
}

func getSofiaProfiles(data *webStruct.MessageData) webStruct.UserResponse {
	return setProfileStatuses(getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfile{}, adminOnly()))
}

func getSofiaProfileParams(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigSofiaProfileParameter{}, adminOnly())
}

func addSofiaProfileParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfileParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigSofiaProfile{Id: data.Id}}, adminOnly())
}

func deleteSofiaProfileParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfileParameter{Id: data.Param.Id}, adminOnly())
}

func switchSofiaProfileParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}, []string{"Enabled"}}, adminOnly())
}

func updateSofiaProfileParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}, []string{"Name", "Value"}}, adminOnly())
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

func addSofiaGatewayParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfileGatewayParameter{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: &altStruct.ConfigSofiaProfileGateway{Id: data.Id}}, adminOnly())
}

func updateSofiaGatewayParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileGatewayParameter{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchSofiaGatewayParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfileGatewayParameter{Id: data.Param.Id, Enabled: data.Param.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteSofiaGatewayParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfileGatewayParameter{Id: data.Param.Id}, adminOnly())
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

func addSofiaProfile(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigSofiaProfile{Name: data.Name, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigSofiaProfile{}))}, adminOnly())
}

func renameSofiaProfile(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfile{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func deleteSofiaProfile(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigSofiaProfile{Id: data.Id}, adminOnly())
}

func switchSofiaProfile(data *webStruct.MessageData) webStruct.UserResponse {
	return setProfileStatuses(getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigSofiaProfile{Id: data.Id, Enabled: *data.Enabled}, []string{"Enabled"}}, adminOnly()))
}
