package altWeb

import (
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
)

func GetConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
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
