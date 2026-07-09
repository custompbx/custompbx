package web

import (
	"custompbx/altStruct"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/webStruct"
	"github.com/custompbx/customorm"
)

// like getConfig
func getByStruct(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	if data.Id == 0 {
		return errorResponse(data.Event, "no parent id")
	}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true, UseValue: true, Value: data.Id}}

	var res interface{}
	var err error
	if hasPagedRequest(data.DBRequest) {
		filterStr := buildFilteredConfigRequest(filter, data.DBRequest)
		res, err = intermediateDB.GetByFilteredValues(
			item,
			filterStr,
		)
		if err != nil {
			return errorResponse(data.Event, err.Error())
		}
		//TODO: with total all the time
		filterStr.Count = true
		resCount, err := intermediateDB.GetByFilteredValues(
			item,
			filterStr,
		)
		if err != nil {
			return errorResponse(data.Event, err.Error())
		}
		if len(resCount) == 0 {
			return errorResponse(data.Event, "can't count total")
		}
		total, ok := resCount[0].(int64)
		if !ok {
			return errorResponse(data.Event, "can't get total")
		}
		res = paginatedResult{Items: res, Total: total}
	} else {
		res, err = intermediateDB.GetByValuesAsMap(
			item,
			filter,
		)
	}

	if err != nil {
		return errorResponse(data.Event, err.Error())
	}

	return dataResponse(data.Event, res)
}

func setProfileStatuses(resp webStruct.UserResponse) webStruct.UserResponse {
	profiles, ok := resp.Data.(map[int64]interface{})
	if ok {
		profilesX := fsesl.GetSofiaProfilesStatuses()
		for _, profileI := range profiles {
			profile, ok := profileI.(altStruct.ConfigSofiaProfile)
			if !ok {
				continue
			}
			profileX := profilesX[profile.Id]
			if profileX == nil {
				continue
			}
			profile.Started = profileX.Started
			profile.State = profileX.State
			profile.Uri = profileX.Uri
			profiles[profile.Id] = profile
		}
		resp.Data = profiles
	}
	return resp
}

func setGatewayStatuses(resp webStruct.UserResponse) webStruct.UserResponse {
	gateways, ok := resp.Data.(map[int64]interface{})
	if ok {
		gatewaysX := fsesl.GetSofiaGatewaysStatuses()
		for _, gatewayI := range gateways {
			gateway, ok := gatewayI.(altStruct.ConfigSofiaProfileGateway)
			if !ok {
				continue
			}
			gatewayX := gatewaysX[gateway.Id]
			if gatewayX == nil {
				continue
			}
			gateway.Started = gatewayX.Started
			gateway.State = gatewayX.State
			gateways[gateway.Id] = gateway
		}
		resp.Data = gateways
	}
	return resp
}
