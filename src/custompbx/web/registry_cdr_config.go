package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/webStruct"
)

func registerCoreCDREvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventCdrPgCsvGet, getCdrPgCsv, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventCdrPgCsvParamUpdate, Switch: eventCdrPgCsvParamSwitch, Add: eventCdrPgCsvParamAdd, Delete: eventCdrPgCsvParamDelete},
		&altStruct.ConfigCdrPgCsvSetting{},
	)
	mustRegisterAdmin(r, eventCdrPgCsvFieldAdd, addCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldUpdate, updateCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldSwitch, switchCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventCdrPgCsvFieldDelete, deleteCdrPgCsvField, overrides)
	mustRegisterAdmin(r, eventOdbcCdrGet, getOdbcCdr, overrides)
	mustRegisterAdmin(r, eventOdbcCdrFieldGet, getOdbcCdrField, overrides)
	registerSimpleParamConfigMutationsForSample(r, overrides,
		simpleParamConfigEvents{Update: eventOdbcCdrParamUpdate, Switch: eventOdbcCdrParamSwitch, Add: eventOdbcCdrParamAdd, Delete: eventOdbcCdrParamDelete},
		&altStruct.ConfigOdbcCdrSetting{},
	)
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
