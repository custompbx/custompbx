package web

import (
	"custompbx/altData"
	"custompbx/cache"
	"custompbx/dataSourceAdapter"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"custompbx/webcache"
)

func getDashboardData(data *webStruct.MessageData) webStruct.UserResponse {
	webcache.DashBoardSetSipRegs(cache.GetDomainSipRegsCounter())

	profiles, gateways := altData.GetSofiaProfilesAndGateways()
	profiles = dataSourceAdapter.UpdateSofiaProfileStatuses(profiles)
	gateways = dataSourceAdapter.UpdateSofiaGatewayStatuses(gateways)
	webcache.DashBoardSetSofiaData(profiles, gateways)
	webcache.DashBoardSetCallsCounter(pbxcache.GetChannelsCounter())
	dashboardData := webcache.GetDashboardData()
	FSMetrics := webcache.GetDashboardFSMetrics()

	return webStruct.UserResponse{MessageType: data.Event, Dashboard: &mainStruct.Dashboard{DashboardData: dashboardData, FSMetrics: FSMetrics}}
}
