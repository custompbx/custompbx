package web

import (
	"custompbx/fsesl"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"log"
)

func GetGlobalVariables(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetGlobalVariableList()

	return webStruct.UserResponse{GlobalVariables: &items, MessageType: data.Event}
}

func UpdateGlobalVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 || data.Variable.Name == "" || data.Variable.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	if mainStruct.IsDynamicGlobalVar(data.Variable.Name) {
		return webStruct.UserResponse{Error: "Dynamic variable name", MessageType: data.Event}
	}

	variable := pbxcache.GetGlobalVariableById(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	if variable.Type == "set" {
		if variable.Name != data.Variable.Name {
			log.Println(fsesl.SendBgapiCmd("global_setvar " + variable.Name + "="))
		}
		log.Println(fsesl.SendBgapiCmd("global_setvar " + data.Variable.Name + "=" + data.Variable.Value))
	}

	res, err := pbxcache.UpdateGlobalVariable(variable, data.Variable.Name, data.Variable.Value, data.Variable.Type)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.GlobalVariable{variable.Id: variable}
	SaveGlobalVarsToFile(false)
	return webStruct.UserResponse{MessageType: data.Event, GlobalVariables: &item}
}

func SwitchGlobalVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	if mainStruct.IsDynamicGlobalVar(data.Variable.Name) {
		return webStruct.UserResponse{Error: "Dynamic variable name", MessageType: data.Event}
	}

	variable := pbxcache.GetGlobalVariableById(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variableeter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchGlobalVariable(variable, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}
	if variable.Type == "set" {
		if data.Variable.Enabled {
			log.Println(fsesl.SendBgapiCmd("global_setvar " + variable.Name + "=" + variable.Value))
		} else {
			log.Println(fsesl.SendBgapiCmd("global_setvar " + variable.Name + "="))
		}
	}

	item := map[int64]*mainStruct.GlobalVariable{variable.Id: variable}
	SaveGlobalVarsToFile(false)
	return webStruct.UserResponse{MessageType: data.Event, GlobalVariables: &item}
}

func AddGlobalVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id != 0 || data.Variable.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	if data.Variable.Type == "" {
		data.Variable.Type = "set"
	}

	if mainStruct.IsDynamicGlobalVar(data.Variable.Name) {
		return webStruct.UserResponse{Error: "Dynamic variable name", MessageType: data.Event}
	}

	variable := pbxcache.GetGlobalVariableByName(data.Variable.Name)
	if variable != nil {
		return webStruct.UserResponse{Error: "variable name already exists", MessageType: data.Event}
	}

	variable, err := pbxcache.SetGlobalVariable(data.Variable.Name, data.Variable.Value, data.Variable.Type, false)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	if variable.Type == "set" {
		log.Println(fsesl.SendBgapiCmd("global_setvar " + data.Variable.Name + "=" + data.Variable.Value))
	}
	item := map[int64]*mainStruct.GlobalVariable{variable.Id: variable}
	SaveGlobalVarsToFile(false)
	return webStruct.UserResponse{MessageType: data.Event, GlobalVariables: &item}
}

func DelGlobalVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	if mainStruct.IsDynamicGlobalVar(data.Variable.Name) {
		return webStruct.UserResponse{Error: "Dynamic variable name", MessageType: data.Event}
	}

	variable := pbxcache.GetGlobalVariableById(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variableeter not found", MessageType: data.Event}
	}

	res := pbxcache.DelGlobalVariable(variable)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}
	SaveGlobalVarsToFile(false)
	if variable.Type == "set" {
		log.Println(fsesl.SendBgapiCmd("global_setvar " + variable.Name + "="))
	}
	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func ImportGlobalVariables(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	fsesl.ImportGlobalVariables()
	SaveGlobalVarsToFile(false)
	return GetGlobalVariables(data, user)
}

func MoveGlobalVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	variable := pbxcache.GetGlobalVariableById(data.Id)
	if variable == nil || variable.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	err := pbxcache.MoveGlobalVariable(variable, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move action", MessageType: data.Event}
	}

	items := pbxcache.GetGlobalVariableList()
	SaveGlobalVarsToFile(false)
	return webStruct.UserResponse{GlobalVariables: &items, MessageType: data.Event}
}

func SaveGlobalVarsToFile(withOs bool) {
	body, path, fileName := pbxcache.GlobalVariablesDropToFile()
	fsesl.SaveToFile(body, path, fileName, withOs)
}
