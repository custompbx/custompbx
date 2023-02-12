package web

import (
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"log"
)

func GetVoicemailSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigVoicemailSettings()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{VoicemailSettings: &items, MessageType: data.Event}
}

func UpdateVoicemailSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailSettings(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subject not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(subject, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.VoicemailSettingsParameter{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailSettings: &item}
}

func SwitchVoicemailSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailSettings(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(subject, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VoicemailSettingsParameter{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailSettings: &item}
}

func AddVoicemailSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailSettingByName(data.Param.Name)
	if subject != nil {
		return webStruct.UserResponse{Error: "subject name already exists", MessageType: data.Event}
	}

	subject, err := pbxcache.SetVoicemailSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VoicemailSettingsParameter{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailSettings: &item}
}

func DelVoicemailSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailSettings(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(subject)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &subject.Id}
}

func GetVoicemailProfiles(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigVoicemailProfiles()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{VoicemailProfiles: &items, MessageType: data.Event}
}

func UpdateVoicemailProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 || data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfile(data.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subject not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(subject, data.Name)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.VoicemailProfile{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailProfiles: &item}
}

func SwitchVoicemailProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfile(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(subject, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VoicemailProfile{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailProfiles: &item}
}

func AddVoicemailProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id != 0 || data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfileByName(data.Name)
	if subject != nil {
		return webStruct.UserResponse{Error: "subject name already exists", MessageType: data.Event}
	}

	subject, err := pbxcache.SetVoicemailProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VoicemailProfile{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailProfiles: &item}
}

func DelVoicemailProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfile(data.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(subject)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &subject.Id}
}

func GetVoicemailProfileParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	exists := pbxcache.CheckVoicemailConfigExists()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profile := pbxcache.GetVoicemailProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()
	return webStruct.UserResponse{VoicemailProfilesParameters: &params, MessageType: data.Event, Id: &profile.Id}
}

func UpdateVoicemailProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfileParam(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subject not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(subject, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.VoicemailProfilesParameter{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailProfilesParameters: &item, Id: &subject.Parent.Parent.Id}
}

func SwitchVoicemailProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfileParam(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(subject, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VoicemailProfilesParameter{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailProfilesParameters: &item, Id: &subject.Parent.Parent.Id}
}

func AddVoicemailProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	profile := pbxcache.GetVoicemailProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfileParamByName(data.Param.Name)
	if subject != nil {
		return webStruct.UserResponse{Error: "subject name already exists", MessageType: data.Event}
	}

	subject, err := pbxcache.SetVoicemailProfileParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VoicemailProfilesParameter{subject.Id: subject}

	return webStruct.UserResponse{MessageType: data.Event, VoicemailProfilesParameters: &item, Id: &subject.Parent.Parent.Id}
}

func DelVoicemailProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	subject := pbxcache.GetVoicemailProfileParam(data.Param.Id)
	if subject == nil {
		return webStruct.UserResponse{Error: "subjecteter not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(subject)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: subject.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &subject.Parent.Parent.Id}
}
