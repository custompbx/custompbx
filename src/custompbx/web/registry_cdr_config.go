package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/webStruct"
)

func registerCoreCDREvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventCdrPgCsvGet, getCdrPgCsv, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvParamAdd, addCdrPgCsvParam, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvParamUpdate, updateCdrPgCsvParam, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvParamSwitch, switchCdrPgCsvParam, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvParamDelete, deleteCdrPgCsvParam, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldAdd, addCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldUpdate, updateCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldSwitch, switchCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldDelete, deleteCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventOdbcCdrGet, getOdbcCdr, overrides)
	mustRegisterAdmin(r, eventOdbcCdrFieldGet, getOdbcCdrField, overrides)
	mustRegisterAdmin(r, eventOdbcCdrParamAdd, addOdbcCdrParam, overrides)
	mustRegisterAdmin(r, eventOdbcCdrParamUpdate, updateOdbcCdrParam, overrides)
	mustRegisterAdmin(r, eventOdbcCdrParamSwitch, switchOdbcCdrParam, overrides)
	mustRegisterAdmin(r, eventOdbcCdrParamDelete, deleteOdbcCdrParam, overrides)
	mustRegisterAdmin(r, eventOdbcCdrTableAdd, addOdbcCdrTable, overrides)
	mustRegisterAdmin(r, eventOdbcCdrTableUpdate, updateOdbcCdrTable, overrides)
	mustRegisterAdmin(r, eventOdbcCdrTableSwitch, switchOdbcCdrTable, overrides)
	mustRegisterAdmin(r, eventOdbcCdrTableDelete, deleteOdbcCdrTable, overrides)
	mustRegisterAdmin(r, eventOdbcCdrFieldAdd, addOdbcCdrField, overrides)
	mustRegisterAdmin(r, eventOdbcCdrFieldUpdate, updateOdbcCdrField, overrides)
	mustRegisterAdmin(r, eventOdbcCdrFieldSwitch, switchOdbcCdrField, overrides)
	mustRegisterAdmin(r, eventOdbcCdrFieldDelete, deleteOdbcCdrField, overrides)
}

func getCdrPgCsv(data *webStruct.MessageData) webStruct.UserResponse {
	resp1 := getUserForConfig(data, getConfig, &altStruct.ConfigCdrPgCsvSetting{}, adminOnly())
	resp2 := getUserForConfig(data, getConfig, &altStruct.ConfigCdrPgCsvSchema{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: resp1.Data},
		responseDataPair{name: "schemas", data: resp2.Data},
	)
}

func addCdrPgCsvParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigCdrPgCsvSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCdrPgCsvSetting{}))}, adminOnly())
}

func updateCdrPgCsvParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigCdrPgCsvSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchCdrPgCsvParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigCdrPgCsvSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteCdrPgCsvParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigCdrPgCsvSetting{Id: data.Param.Id}, adminOnly())
}

func addCdrPgCsvField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigCdrPgCsvSchema{Var: data.Field.Var, Column: data.Field.Column, Quote: data.Field.Quote, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigCdrPgCsvSchema{}))}, adminOnly())
}

func updateCdrPgCsvField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigCdrPgCsvSchema{Id: data.Field.Id, Var: data.Field.Var, Column: data.Field.Column, Quote: data.Field.Quote}, []string{"Var", "Column", "Quote"}}, adminOnly())
}

func switchCdrPgCsvField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigCdrPgCsvSchema{Id: data.Field.Id, Enabled: data.Field.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteCdrPgCsvField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigCdrPgCsvSchema{Id: data.Field.Id}, adminOnly())
}

func getOdbcCdr(data *webStruct.MessageData) webStruct.UserResponse {
	resp1 := getUserForConfig(data, getConfig, &altStruct.ConfigOdbcCdrSetting{}, adminOnly())
	resp2 := getUserForConfig(data, getConfig, &altStruct.ConfigOdbcCdrTable{}, adminOnly())
	return combinedDataResponse(data.Event,
		responseDataPair{name: "settings", data: resp1.Data},
		responseDataPair{name: "tables", data: resp2.Data},
	)
}

func getOdbcCdrField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigOdbcCdrTableField{Parent: &altStruct.ConfigOdbcCdrTable{Id: data.Id}}, adminOnly())
}

func addOdbcCdrParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigOdbcCdrSetting{Name: data.Param.Name, Value: data.Param.Value, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOdbcCdrSetting{}))}, adminOnly())
}

func updateOdbcCdrParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigOdbcCdrSetting{Id: data.Param.Id, Name: data.Param.Name, Value: data.Param.Value}, []string{"Name", "Value"}}, adminOnly())
}

func switchOdbcCdrParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigOdbcCdrSetting{Id: data.Param.Id, Enabled: data.Param.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteOdbcCdrParam(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigOdbcCdrSetting{Id: data.Param.Id}, adminOnly())
}

func addOdbcCdrTable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigOdbcCdrTable{Name: data.Table.Name, LogLeg: data.Table.LogLeg, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigOdbcCdrTable{}))}, adminOnly())
}

func updateOdbcCdrTable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigOdbcCdrTable{Id: data.Table.Id, Name: data.Table.Name, LogLeg: data.Table.LogLeg}, []string{"Name", "LogLeg"}}, adminOnly())
}

func switchOdbcCdrTable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigOdbcCdrTable{Id: data.Table.Id, Enabled: data.Table.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteOdbcCdrTable(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigOdbcCdrTable{Id: data.Table.Id}, adminOnly())
}

func addOdbcCdrField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigOdbcCdrTableField{Name: data.OdbcCdrField.Name, ChanVarName: data.OdbcCdrField.ChanVarName, Enabled: true, Parent: &altStruct.ConfigOdbcCdrTable{Id: data.Id}}, adminOnly())
}

func updateOdbcCdrField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigOdbcCdrTableField{Id: data.OdbcCdrField.Id, Name: data.OdbcCdrField.Name, ChanVarName: data.OdbcCdrField.ChanVarName}, []string{"Name", "ChanVarName"}}, adminOnly())
}

func switchOdbcCdrField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigOdbcCdrTableField{Id: data.OdbcCdrField.Id, Enabled: data.OdbcCdrField.Enabled}, []string{"Enabled"}}, adminOnly())
}

func deleteOdbcCdrField(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigOdbcCdrTableField{Id: data.OdbcCdrField.Id}, adminOnly())
}
