package web

import (
	"custompbx/cache"
	"custompbx/db"
	"custompbx/fsesl"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"strings"
)

func runCLICommand(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func RealFSCLIConnect(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func RealFSCLICommand(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty command", MessageType: data.Event}
	}
	res := fsesl.OneTimeConnectCommand(strings.TrimSpace(data.Name))

	return webStruct.UserResponse{MessageType: data.Event, Response: &res}
}

func GetLogs(data *webStruct.MessageData) webStruct.UserResponse {
	limit, offset := normalizePagination(data.DBRequest.Limit, data.DBRequest.Offset)
	logs, err := db.GetList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order, cache.GetCurrentInstanceId())
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Logs: &logs, MessageType: data.Event}
}

func getHEP(data *webStruct.MessageData) webStruct.UserResponse {
	limit, offset := normalizePagination(data.DBRequest.Limit, data.DBRequest.Offset)
	heps, err := db.GetHEPList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{HEPs: &heps, MessageType: data.Event}
}

func GetHEPDetails(data *webStruct.MessageData) webStruct.UserResponse {
	if len(data.ArrVal) == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	heps, err := db.GetHEPDetailsList(data.ArrVal, cache.GetCurrentInstanceId())
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{HEPsDetails: &heps, MessageType: data.Event}
}

func GetInstances(data *webStruct.MessageData) webStruct.UserResponse {
	cache.UpdateCacheInstances()
	var res = cache.GetFSInstances().GetList()
	var currentId = cache.GetCurrentInstanceId()
	return webStruct.UserResponse{FSInstances: &res, MessageType: "GetInstances", Id: &currentId}
}

func UpdateInstanceDescription(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	instance := cache.GetFSInstances().GetById(data.Id)
	if instance == nil {
		return webStruct.UserResponse{Error: "instance not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateFSInstanceDescription(instance, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.FsInstance{instance.Id: instance}

	return webStruct.UserResponse{MessageType: data.Event, FSInstances: &item}

}
