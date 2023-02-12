package web

import (
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
)

func GetPostLoadModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigPostLoadModules()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{ModuleTags: &items, MessageType: data.Event}

	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	result := pbxcache.GetConfigs(res)

	return webStruct.UserResponse{MessageType: data.Event, Modules: result}

}

func UpdatePostLoadModule(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetPostLoadModule(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subject not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(subject, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.ModuleTag{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, ModuleTags: &item}
}

func SwitchPostLoadModule(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetPostLoadModule(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(subject, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ModuleTag{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, ModuleTags: &item}
}

func AddPostLoadModule(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetPostLoadModuleByName(data.Param.Name)
	if subject != nil {
		return webStruct.UserResponse{Error: "subject name already exists", MessageType: data.Event}
	}

	subject, err := pbxcache.SetPostLoadModule(data.Param.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ModuleTag{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, ModuleTags: &item}
}

func DelPostLoadModule(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetPostLoadModule(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(subject)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &subject.Id}
}

func AutoloadPostLoadModule(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	module := pbxcache.GetModule(data.Param.Id)
	if module == nil {
		return webStruct.UserResponse{Error: "module not found", MessageType: data.Event}
	}

	subject := pbxcache.GetPostLoadModuleByName(module.GetModuleName())
	if subject == nil {
		var err error
		subject, err = pbxcache.SetPostLoadModule(module.GetModuleName())
		if err != nil {
			return webStruct.UserResponse{Error: "SetPostLoadModule error", MessageType: data.Event}
		}
	} else {
		pbxcache.SwitchConfigRow(subject, !subject.Enabled)
	}

	item := map[int64]*mainStruct.ModuleTag{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, ModuleTags: &item}
}
