package pbxcache

import (
	"custompbx/cache"
	"custompbx/db"
	"custompbx/mainStruct"
	"encoding/xml"
	"errors"
	"log"
	"strings"
)

func UpdateGlobalVariable(variable *mainStruct.GlobalVariable, name, value, varType string) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	switch varType {
	case "env-set":
	case "exec-set":
	case "stun-set":
	default:
		varType = "set"
	}
	err := db.UpdateGlobalVariable(variable.Id, name, value, varType)
	if err != nil {
		return 0, err
	}
	//GlobalVariablesDropToFile()
	variable.Name = name
	variable.Value = value
	variable.Type = varType
	return variable.Id, err
}

func SwitchGlobalVariable(variable *mainStruct.GlobalVariable, switcher bool) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	err := db.SwitchGlobalVariable(variable.Id, switcher)
	if err != nil {
		return 0, err
	}
	variable.Enabled = switcher
	return variable.Id, err
}

func DelGlobalVariable(variable *mainStruct.GlobalVariable) int64 {
	id := variable.Id
	ok := db.DelGlobalVariable(variable.Id)
	if !ok {
		return 0
	}

	globalVariables.Remove(variable)
	return id
}

func SetGlobalVariable(variableName, variableValue, varType string, dynamic bool) (*mainStruct.GlobalVariable, error) {
	res, position, err := db.SetGlobalVariable(variableName, variableValue, varType, cache.GetCurrentInstanceId(), dynamic)
	if err != nil {
		return nil, err
	}

	variable := &mainStruct.GlobalVariable{Id: res, Name: variableName, Value: variableValue, Type: varType, Enabled: true, Dynamic: dynamic, Position: position}
	globalVariables.Set(variable)
	return variable, err
}

func MoveGlobalVariable(variable *mainStruct.GlobalVariable, newPosition int64) error {
	if variable == nil || newPosition == 0 {
		return errors.New("node doesn't exists")
	}
	err := db.MoveGlobalVariable(GetGlobalVariables(), variable, newPosition, cache.GetCurrentInstanceId())
	if err != nil {
		return err
	}

	return err
}

func GlobalVariablesDropToFile() (string, string, string) {
	path := GetGlobalVariableByName("conf_dir")
	if path == nil {
		return "", "", ""
	}

	var newVars []*mainStruct.X_PRE_PROCESS
	for _, value := range GetGlobalVariableNotDynamicsProps() {
		if !value.Enabled {
			continue
		}
		newVars = append(newVars, &mainStruct.X_PRE_PROCESS{XMLName: xml.Name{Space: "", Local: "X-PRE-PROCESS"}, Attrcmd: value.Type, Attrdata: value.Name + "=" + value.Value, Attrmetatype: "custompbx"})
	}
	/*raw, err := SendBgapiCmd("system cat " + path + "/vars.xml")
	if err != nil {
		log.Println("Cant get global vars: " + err.Error())
		return
	}

	var parsed Includer
	reader := strings.NewReader(raw)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&parsed)
	if err != nil {
		//log.Println(err.Error())
		return
	}
	var own []*X_PRE_PROCESS
	var legacy []*X_PRE_PROCESS
	for _, v := range parsed.Includes {
		if v.Attrmetatype == "custompbx" {
			own = append(own, v)
			continue
		}
		legacy = append(legacy, v)
	}
	if len(own) == 0 {
		_, err = SendBgapiCmd("system cp " + path + "/vars.xml " + path + "/vars.xml.dump")
		if err != nil {
			log.Println("Cant get global vars: " + err.Error())
			return
		}
	}*/
	/*	if len(newVars) == 0 {
		return
	}*/
	/*_, err = SendBgapiCmd("system grep -v \"\\meta_type=\"custompbx\\\"\" " + path + "/vars.xml >  " + path + "/vars.xml.proc")
	if err != nil {
		log.Println("Cant get global vars: " + err.Error())
		return
	}*/
	var inc = mainStruct.Includer{XMLName: xml.Name{Space: "", Local: "include"}, Includes: newVars}
	output, err := xml.MarshalIndent(inc, "", "  ")
	if err != nil {
		log.Printf("%+v", err)
		return "", "", ""
	}
	str := strings.Replace(string(output), "'", "\\'", -1)
	str = strings.Replace(str, "\n", "\\n", -1)
	str = strings.Replace(str, "'", "\\'", -1)

	fileName := "vars.xml"

	return str, path.Value, fileName
}
