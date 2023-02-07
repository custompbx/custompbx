package dataSourceAdapter

import (
	"custompbx/altStruct"
	"custompbx/fsesl"
)

func UpdateSofiaProfileStatuses(profiles []interface{}) []interface{} {
	profilesX := fsesl.GetSofiaProfilesStatuses()
	for k, profileI := range profiles {
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
		profiles[k] = profile
	}
	return profiles
}

func UpdateSofiaGatewayStatuses(gateways []interface{}) []interface{} {
	gatewaysX := fsesl.GetSofiaGatewaysStatuses()
	for k, gatewayI := range gateways {
		gateway, ok := gatewayI.(altStruct.ConfigSofiaProfileGateway)
		if !ok {
			continue
		}
		profileX := gatewaysX[gateway.Id]
		if profileX == nil {
			continue
		}
		gateway.Started = profileX.Started
		gateway.State = profileX.State
		gateways[k] = gateway
	}
	return gateways
}
