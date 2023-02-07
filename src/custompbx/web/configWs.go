package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"custompbx/xmlStruct"
	"fmt"
	"github.com/custompbx/customorm"
	"log"
	"reflect"
	"strconv"
)

func getConfParent(name string) *altStruct.ConfigurationsList {
	log.Printf("%+v", name)
	conf, err := altData.GetModuleByName(mainStruct.GetModuleNameByConfName(name))
	if err != nil {
		return nil
	}
	return conf
}

func getConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	var parent interface{}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true}}
	if data.Id == 0 {
		parent = getConfParent(altData.GetConfNameByStruct(item))
		if parent == (*altStruct.ConfigurationsList)(nil) {
			exists := false
			return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
		}
	} else {
		parent = nil
		filter["Parent"] = customorm.FilterFields{Flag: true, UseValue: true, Value: data.Id}
	}

	module := altData.GetConfInstanceByStruct(item, parent)
	if module == nil {
		return webStruct.UserResponse{Error: "unknown config", MessageType: data.Event}
	}
	var res interface{}
	var err error
	if data.DBRequest.Limit != 0 || data.DBRequest.Filters != nil {
		limit := data.DBRequest.Limit
		if limit < 25 || limit > 250 {
			limit = 25
		}
		offset := 0
		if data.DBRequest.Offset > 0 {
			offset = data.DBRequest.Offset * limit
		}
		for _, v := range data.DBRequest.Filters {
			filter[v.Field] = customorm.FilterFields{Flag: true, UseValue: true, Value: v.FieldValue, Operand: v.Operand}
		}
		filterStr := customorm.Filters{
			Fields: filter,
			Limit:  limit,
			Offset: offset,
			Order:  customorm.Order{Desc: data.DBRequest.Order.Desc, Fields: data.DBRequest.Order.Fields}}
		res, err = intermediateDB.GetByFilteredValues(
			module,
			filterStr,
		)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
		}
		//TODO: with total all the time
		filterStr.Count = true
		resCount, err := intermediateDB.GetByFilteredValues(
			module,
			filterStr,
		)
		if err != nil {
			return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
		}
		if len(resCount) == 0 {
			return webStruct.UserResponse{Error: "can't count total", MessageType: data.Event}
		}
		total, ok := resCount[0].(int64)
		if !ok {
			return webStruct.UserResponse{Error: "can't get total", MessageType: data.Event}
		}
		res = struct {
			Items interface{} `json:"items"`
			Total int64       `json:"total"`
		}{Items: res, Total: total}
	} else {
		res, err = intermediateDB.GetByValuesAsMap(
			module,
			filter,
		)
	}

	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: res, MessageType: data.Event}
}

func getConfigInnerParent(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true}}
	log.Printf("%+v", item)
	res, err := intermediateDB.GetByValuesAsMap(
		item,
		filter,
	)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &res, MessageType: data.Event}
}

func delConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	/*	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}*/
	res, err := intermediateDB.GetByIdFromDB(item)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	err = intermediateDB.DeleteById(item)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &res, MessageType: data.Event}
}

func setConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	log.Printf("%+v", item)
	res, err := intermediateDB.InsertItem(item)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	result, err := intermediateDB.GetByIdArg(item, res)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: &result, MessageType: data.Event}
}

func updateConfig(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	items, ok := item.(struct {
		S interface{}
		A []string
	})
	if !ok {
		return webStruct.UserResponse{Error: "no mandatory params", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(items.S, items.A)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	var returnList bool
	for _, v := range items.A {
		if v != "Position" {
			continue
		}
		returnList = true
		break
	}
	if returnList {
		return getConfigInnerParent(data, result)
	}

	return webStruct.UserResponse{Data: &result, MessageType: data.Event}
}

func updateConfigAclListDefault(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateAclListDefault(data.Id, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Value: data.Value}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item}
}

func delConfigAclNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Node.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	node := pbxcache.GetAclNode(data.Node.Id)
	if node == nil {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	res := pbxcache.DelAclNode(node)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: node.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func updateConfigAclNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Node.Id == 0 || data.Node.Type == "" || (data.Node.Cidr == "" && data.Node.Domain == "") {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	node := pbxcache.GetAclNode(data.Node.Id)
	if node == nil {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateAclNode(node, data.Node.Type, data.Node.Cidr, data.Node.Domain)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	nodes := map[int64]*mainStruct.Node{node.Id: node}
	return webStruct.UserResponse{MessageType: data.Event, AclListNodes: &nodes, Id: &node.List.Id}
}

func switchConfigAclNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Node.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	node := pbxcache.GetAclNode(data.Node.Id)
	if node == nil {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchAclNode(node, data.Node.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	nodes := map[int64]*mainStruct.Node{node.Id: node}
	return webStruct.UserResponse{MessageType: data.Event, AclListNodes: &nodes, Id: &node.List.Id}
}

func addConfigAclNodes(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Node.Id != 0 || data.Node.Type == "" || (data.Node.Cidr == "" && data.Node.Domain == "") {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	list := pbxcache.GetAclList(data.Id)
	if list == nil {
		return webStruct.UserResponse{Error: "acl list not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfAclNode(list, data.Node.Type, data.Node.Cidr, data.Node.Domain)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	node := list.Nodes.GetById(res)
	if node == nil {
		return webStruct.UserResponse{Error: "not added", MessageType: data.Event}
	}
	nodes := map[int64]*mainStruct.Node{list.Id: node}

	return webStruct.UserResponse{AclListNodes: &nodes, MessageType: data.Event}
}

func MoveAclListNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	node := pbxcache.GetAclNode(data.Id)
	if node == nil || node.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	err := pbxcache.MoveAclListNode(node, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move action", MessageType: data.Event}
	}

	nodes := node.List.Nodes.GetList()
	return webStruct.UserResponse{MessageType: data.Event, AclListNodes: &nodes, Id: &node.List.Id}
}

func addAclList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	res, err := pbxcache.SetConfAclList(data.Name, data.Default)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.List)
	items[res.Id] = res

	return webStruct.UserResponse{AclLists: &items, MessageType: data.Event}
}

func delAclList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "ACL list not found", MessageType: data.Event}
	}
	ok := pbxcache.DelAclList(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete ACL list", MessageType: data.Event}
	}

	return webStruct.UserResponse{Id: &data.Id, MessageType: data.Event}
}

func renameAclList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "ACL list not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	err := pbxcache.UpdateAclList(data.Id, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	res := pbxcache.GetAclList(data.Id)
	items := make(map[int64]*mainStruct.List)
	items[res.Id] = res

	return webStruct.UserResponse{AclLists: &items, MessageType: data.Event}
}

func getConfigSofiaGlobalSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetSofiaGlobalSettings()
	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func getConfigSofiaProfiles(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetSofiaProfilesLists()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{SofiaProfiles: &items, MessageType: data.Event}
}

func getConfigSofiaProfileParams(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, SofiaProfileParams: &params, Id: &profile.Id}
}

func updateConfigSofiaGlobalSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaGlobalSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateSofiaGlobalSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func switchConfigSofiaGlobalSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaGlobalSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchSofiaGlobalSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func addConfigSofiaGlobalSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigSofiaGlobalSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func delConfigSofiaGlobalSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaGlobalSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelSofiaGlobalSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func addConfigSofiaProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfSofiaProfileParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.SofiaProfileParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, SofiaProfileParams: &item, Id: &profile.Id}
}

func delConfigSofiaProfileSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfileParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func switchSofiaProfileSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchProfileParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.SofiaProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, SofiaProfileParams: &params, Id: &param.Profile.Id}
}

func updateSofiaProfileSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateProfileParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.SofiaProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, SofiaProfileParams: &params, Id: &param.Profile.Id}
}

func getConfigSofiaProfileGateways(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetSofiaParentsGateways()
	return webStruct.UserResponse{SofiaGateways: &items, MessageType: data.Event}
}

func getConfigSofiaProfileGatewaysDetails(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	gateway := pbxcache.GetSofiaProfileGateway(data.Id)
	if gateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[gateway.Id] = &webStruct.SofiaGatewayDetails{Params: gateway.Params.GetList(), Vars: gateway.Vars.GetList()}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &gateway.Profile.Id}
}

func addConfigSofiaProfileGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	gateway := pbxcache.GetSofiaProfileGateway(data.Id)
	if gateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigSofiaGatewayParam(gateway, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayParam{param.Id: param}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[gateway.Id] = &webStruct.SofiaGatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &gateway.Profile.Id}
}

func updateConfigSofiaProfileGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaProfileGatewayParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateProfileGatewayParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayParam{param.Id: param}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[param.Gateway.Id] = &webStruct.SofiaGatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &param.Gateway.Profile.Id}
}

func switchConfigSofiaProfileGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaProfileGatewayParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchProfileGatewayParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayParam{param.Id: param}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[param.Gateway.Id] = &webStruct.SofiaGatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &param.Gateway.Profile.Id}
}

func delConfigSofiaProfileGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSofiaProfileGatewayParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfileGatewayParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayParam{param.Id: param}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[param.Gateway.Id] = &webStruct.SofiaGatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &param.Gateway.Profile.Id}
}

func addConfigSofiaProfileGatewayVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Variable.Id != 0 || data.Variable.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	gateway := pbxcache.GetSofiaProfileGateway(data.Id)
	if gateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	variable, err := pbxcache.SetConfigSofiaGatewayVar(gateway, data.Variable.Name, data.Variable.Value, data.Variable.Direction)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[gateway.Id] = &webStruct.SofiaGatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &gateway.Profile.Id}
}

func updateConfigSofiaProfileGatewayVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 || data.Variable.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetSofiaProfileGatewayVariable(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateProfileGatewayVariable(variable, data.Variable.Name, data.Variable.Value, data.Variable.Direction)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[variable.Gateway.Id] = &webStruct.SofiaGatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &variable.Gateway.Profile.Id}
}

func switchConfigSofiaProfileGatewayVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetSofiaProfileGatewayVariable(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchProfileGatewayVariable(variable, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[variable.Gateway.Id] = &webStruct.SofiaGatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &variable.Gateway.Profile.Id}
}

func delConfigSofiaProfileGatewayVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetSofiaProfileGatewayVariable(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfileGatewayVariable(variable)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.SofiaGatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.SofiaGatewayDetails)
	item[variable.Gateway.Id] = &webStruct.SofiaGatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, SofiaGatewayDetails: &item, Id: &variable.Gateway.Profile.Id}
}

func getConfigSofiaProfileDomains(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Domains.GetList()

	return webStruct.UserResponse{MessageType: data.Event, SofiaDomains: &params, Id: &profile.Id}
}

func addConfigSofiaProfileDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.SofiaDomain.Id != 0 || data.SofiaDomain.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	var aliasValue = "false"
	var parseValue = "false"
	if data.SofiaDomain.Alias {
		aliasValue = "true"
	}
	if data.SofiaDomain.Parse {
		parseValue = "true"
	}

	res, err := pbxcache.SetConfigSofiaProfileDomain(profile, data.SofiaDomain.Name, aliasValue, parseValue)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.SofiaDomain{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, SofiaDomains: &item, Id: &profile.Id}
}

func delConfigSofiaProfileDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.SofiaDomain.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetSofiaProfileDomain(data.SofiaDomain.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfileDomain(domain)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.SofiaDomain{domain.Id: domain}
	return webStruct.UserResponse{MessageType: data.Event, SofiaDomains: &item, Id: &domain.Profile.Id}
}

func switchSofiaProfileDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.SofiaDomain.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetSofiaProfileDomain(data.SofiaDomain.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchProfileDomain(domain, data.SofiaDomain.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	domains := map[int64]*mainStruct.SofiaDomain{domain.Id: domain}

	return webStruct.UserResponse{MessageType: data.Event, SofiaDomains: &domains, Id: &domain.Profile.Id}
}

func updateSofiaProfileDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.SofiaDomain.Id == 0 || data.SofiaDomain.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetSofiaProfileDomain(data.SofiaDomain.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateProfileDomain(domain, data.SofiaDomain.Name, data.SofiaDomain.Alias, data.SofiaDomain.Parse)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	domains := map[int64]*mainStruct.SofiaDomain{domain.Id: domain}

	return webStruct.UserResponse{MessageType: data.Event, SofiaDomains: &domains, Id: &domain.Profile.Id}
}

func getConfigSofiaProfileAliases(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Aliases.GetList()

	return webStruct.UserResponse{MessageType: data.Event, SofiaAliases: &params, Id: &profile.Id}
}

func addConfigSofiaProfileAlias(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.SofiaAlias.Id != 0 || data.SofiaAlias.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigSofiaProfileAliases(profile, data.SofiaAlias.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Alias{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, SofiaAliases: &item, Id: &profile.Id}
}

func delConfigSofiaProfileAlias(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.SofiaAlias.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	alias := pbxcache.GetSofiaProfileAlias(data.SofiaAlias.Id)
	if alias == nil {
		return webStruct.UserResponse{Error: "alias not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfileAlias(alias)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Alias{alias.Id: alias}
	return webStruct.UserResponse{MessageType: data.Event, SofiaAliases: &item, Id: &alias.Profile.Id}
}

func switchSofiaProfileAlias(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.SofiaAlias.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	alias := pbxcache.GetSofiaProfileAlias(data.SofiaAlias.Id)
	if alias == nil {
		return webStruct.UserResponse{Error: "alias not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchProfileAlias(alias, data.SofiaAlias.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	aliases := map[int64]*mainStruct.Alias{alias.Id: alias}

	return webStruct.UserResponse{MessageType: data.Event, SofiaAliases: &aliases, Id: &alias.Profile.Id}
}

func updateSofiaProfileAlias(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.SofiaAlias.Id == 0 || data.SofiaAlias.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	alias := pbxcache.GetSofiaProfileAlias(data.SofiaAlias.Id)
	if alias == nil {
		return webStruct.UserResponse{Error: "alias not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateProfileAlias(alias, data.SofiaAlias.Name)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	aliass := map[int64]*mainStruct.Alias{alias.Id: alias}

	return webStruct.UserResponse{MessageType: data.Event, SofiaAliases: &aliass, Id: &alias.Profile.Id}
}

func addConfigSofiaProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigSofiaProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.SofiaProfile)
	items[res.Id] = res

	return webStruct.UserResponse{SofiaProfiles: &items, MessageType: data.Event}
}

func addConfigSofiaProfileGateway(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigSofiaGateway(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.SofiaGateway{res.Id: res}
	items := map[int64]map[int64]*mainStruct.SofiaGateway{profile.Id: item}

	return webStruct.UserResponse{SofiaGateways: &items, MessageType: data.Event}
}

func updateConfigSofiaProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateSofiaProfile(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.SofiaProfile{profile.Id: profile}

	return webStruct.UserResponse{SofiaProfiles: &item, MessageType: data.Event}
}

func delConfigSofiaProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfile(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func updateConfigSofiaProfileGateway(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	gateway := pbxcache.GetSofiaProfileGateway(data.Id)
	if gateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateSofiaProfileGateway(gateway, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.SofiaGateway{gateway.Id: gateway}
	items := map[int64]map[int64]*mainStruct.SofiaGateway{gateway.Profile.Id: item}

	return webStruct.UserResponse{SofiaGateways: &items, MessageType: data.Event}
}

func delConfigSofiaProfileGateway(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	gateway := pbxcache.GetSofiaProfileGateway(data.Id)
	if gateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	res := pbxcache.DelProfileGateway(gateway)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: gateway.Id}

	return webStruct.UserResponse{MessageType: data.Event, Id: &gateway.Profile.Id, Item: &item}
}

func reloadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Id: data.Id},
		map[string]bool{"Id": true},
	)
	if err != nil || len(res) == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	module, ok := res[0].(altStruct.ConfigurationsList)
	if !ok {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	comm := "reload " + mainStruct.GetModuleNameByConfName(module.Name)
	if mainStruct.GetModuleNameByConfName(module.Name) == mainStruct.ConfAcl {
		comm = "reloadacl"
	}
	_, err = fsesl.SendBgapiCmd(comm)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func unloadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Id: data.Id},
		map[string]bool{"Id": true},
	)
	if err != nil || len(res) == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	module, ok := res[0].(altStruct.ConfigurationsList)

	if !ok {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	comm := "unload " + mainStruct.GetModuleNameByConfName(module.Name)
	if mainStruct.GetModuleNameByConfName(module.Name) == mainStruct.ConfAcl {
		comm = "reloadacl"
	}
	_, err = fsesl.SendBgapiCmd(comm)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func loadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValue(
		&altStruct.ConfigurationsList{Id: data.Id},
		map[string]bool{"Id": true},
	)

	if err != nil || len(res) == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}
	module, ok := res[0].(altStruct.ConfigurationsList)

	if !ok {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	comm := "load " + mainStruct.GetModuleNameByConfName(module.Name)
	if mainStruct.GetModuleNameByConfName(module.Name) == mainStruct.ConfAcl {
		comm = "reloadacl"
	}
	_, err = fsesl.SendBgapiCmd(comm)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}
	return webStruct.UserResponse{MessageType: data.Event}
}

func switchConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 || data.Enabled == nil {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	err := intermediateDB.UpdateByIdByValuesMap(
		&altStruct.ConfigurationsList{Id: data.Id, Enabled: *data.Enabled},
		map[string]bool{"Enabled": true},
	)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

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

func importConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	confName := data.Name + ".conf"
	err := fsesl.GetXMLModuleConfiguration(confName)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func TruncateModuleConfig(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
	}

	err := intermediateDB.DeleteById(&altStruct.ConfigurationsList{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

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

func ImportXMLModuleConfig(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.File == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	err := fsesl.ParseConfigXML(data.File)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

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

/*
	func autoloadConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
		if data.Id == 0 {
			return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
		}

		module := pbxcash.GetModule(data.Id)
		if module == nil {
			return webStruct.UserResponse{Error: "configuration not found", MessageType: data.Event}
		}

		module.AutoLoad()

		return webStruct.UserResponse{MessageType: data.Event}
	}
*/
func importConfAllModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	fsesl.GetXMLConfigurations()

	return webStruct.UserResponse{MessageType: data.Event}
}

func fromScratchConfModules(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	confName := data.Name + ".conf"
	fsesl.InitConfigModule(&xmlStruct.Configuration{Attrname: confName})

	return webStruct.UserResponse{MessageType: data.Event}
}

func runProfileCommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profileI, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigSofiaProfile{Id: data.Id})
	if err != nil || profileI == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}
	profile, ok := profileI.(altStruct.ConfigSofiaProfile)
	if !ok {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}
	if data.Name == "" || !mainStruct.CheckCommand(data.Name) {
		return webStruct.UserResponse{Error: "unknown command", MessageType: data.Event}
	}

	if (data.Name == mainStruct.CommandSofiaProfileKillgw || data.Name == mainStruct.CommandSofiaProfileStartgw) && data.IdInt == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	gatewayName := ""
	if data.Name == mainStruct.CommandSofiaProfileKillgw || data.Name == mainStruct.CommandSofiaProfileStartgw {
		gatewayI, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigSofiaProfileGateway{Id: data.IdInt})
		if err != nil || gatewayI == nil {
			return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
		}
		gateway, ok := gatewayI.(altStruct.ConfigSofiaProfileGateway)
		if !ok {
			return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
		}
		gatewayName = gateway.Name
	}

	command := fmt.Sprintf("sofia profile %s %s %s", profile.Name, data.Name, gatewayName)

	_, err = fsesl.SendBgapiCmd(command)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func switchSofiaProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetSofiaProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Enabled == nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchProfile(profile, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := &map[int64]*mainStruct.SofiaProfile{profile.Id: profile}

	return webStruct.UserResponse{MessageType: data.Event, SofiaProfiles: item}
}

func getConfigCdrPgCsv(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetCdrPgCsvSettings()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	fields, _ := pbxcache.GetCdrPgCsvSchema()

	return webStruct.UserResponse{Parameters: &items, CdrPgCsvSchema: &fields, MessageType: data.Event}
}

func addConfigCdrPgCsvParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfCdrPgCsvSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func addConfigCdrPgCsvField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Field.Id != 0 || data.Field.Var == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field, err := pbxcache.SetConfCdrPgCsvField(data.Field.Var, data.Field.Column)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Field{field.Id: field}
	return webStruct.UserResponse{MessageType: data.Event, CdrPgCsvSchema: &item}
}

func switchCdrPgCsvParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrPgCsvParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchCdrPgCsvParam(param, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}
	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func switchCdrPgCsvField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Field.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field := pbxcache.GetCdrPgCsvField(data.Field.Id)
	if field == nil {
		return webStruct.UserResponse{Error: "field not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchCdrPgCsvField(field, data.Field.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Field{field.Id: field}
	return webStruct.UserResponse{MessageType: data.Event, CdrPgCsvSchema: &item}
}

func updateCdrPgCsvParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrPgCsvParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateCdrPgCsvParam(param, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}
	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func updateCdrPgCsvField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Field.Id == 0 || data.Field.Var == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field := pbxcache.GetCdrPgCsvField(data.Field.Id)
	if field == nil {
		return webStruct.UserResponse{Error: "field not update", MessageType: data.Event}
	}

	err := pbxcache.UpdateCdrPgCsvField(field, data.Field.Var, data.Field.Column)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Field{field.Id: field}
	return webStruct.UserResponse{MessageType: data.Event, CdrPgCsvSchema: &item}
}

func delConfigCdrPgCsvParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrPgCsvParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelCdrPgCsvParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func delConfigCdrPgCsvField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Field.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field := pbxcache.GetCdrPgCsvField(data.Field.Id)
	if field == nil {
		return webStruct.UserResponse{Error: "field not found", MessageType: data.Event}
	}

	res := pbxcache.DelCdrPgCsvField(field)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func getConfigVerto(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigVerto()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetProfiles()

	return webStruct.UserResponse{Parameters: &items, VertoProfiles: &profiles, MessageType: data.Event}
}

func getVertoProfileParams(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetVertoProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, VertoProfileParams: &params, Id: &profile.Id}
}

func updateConfigVertoSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateVertoSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func switchConfigVertoSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchVertoSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func addConfigVertoSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigVertoSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func delConfigVertoSetting(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelVertoSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func addConfigVertoProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetVertoProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil && data.Param.Name != "bind-local" && data.Param.Name != "apply-candidate-acl" {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	if data.Param.Name == "bind-local" {
		data.Param.Secure = ""
	}

	res, err := pbxcache.SetConfigVertoProfileParam(profile, data.Param.Name, data.Param.Value, data.Param.Secure)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.VertoProfileParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, VertoProfileParams: &item, Id: &profile.Id}
}

func MoveVertoProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.CurrentIndex == 0 || data.PreviousIndex == 0 {
		return webStruct.UserResponse{Error: "wrong position", MessageType: data.Event}
	}

	node := pbxcache.GetVertoProfileParam(data.Id)
	if node == nil || node.Position != data.PreviousIndex {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	err := pbxcache.MoveVertoProfileParam(node, data.CurrentIndex)
	if err != nil {
		return webStruct.UserResponse{Error: "can't move param", MessageType: data.Event}
	}

	nodes := node.Profile.Params.GetList()
	return webStruct.UserResponse{MessageType: data.Event, VertoProfileParams: &nodes, Id: &node.Profile.Id}
}

func delConfigVertoProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelVertoProfileParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func switchVertoProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchVertoProfileParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.VertoProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, VertoProfileParams: &params, Id: &param.Profile.Id}
}

func updateVertoProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetVertoProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateVertoProfileParam(param, data.Param.Name, data.Param.Value, data.Param.Secure)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.VertoProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, VertoProfileParams: &params, Id: &param.Profile.Id}
}

func addConfigVertoProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetVertoProfileByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigVertoProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.VertoProfile)
	items[res.Id] = res

	return webStruct.UserResponse{VertoProfiles: &items, MessageType: data.Event}
}

func updateConfigVertoProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetVertoProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateVertoProfile(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.VertoProfile{profile.Id: profile}

	return webStruct.UserResponse{VertoProfiles: &item, MessageType: data.Event}
}

func delConfigVertoProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetVertoProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelVertoProfile(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetCallcenterQueues(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetCallcenterQueuesLists()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{CallcenterQueues: &items, MessageType: data.Event}
}

func GetCallcenterQueuesParams(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	queue := pbxcache.GetCallcenterQueue(data.Id)
	if queue == nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	params := queue.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, CallcenterQueuesParams: &params, Id: &queue.Id}
}

func UpdateCallcenterSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCallcenterSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateCallcenterSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchCallcenterSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCallcenterSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchCallcenterSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddCallcenterSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfCallcenterSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelCallcenterSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCallcenterSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelCallcenterSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddCallcenterQueueParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	queue := pbxcache.GetCallcenterQueue(data.Id)
	if queue == nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfCallcenterQueueParam(queue, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.QueueParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, CallcenterQueuesParams: &item, Id: &queue.Id}
}

func DelCallcenterQueueParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCallcenterQueueParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelCallcenterQueueParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchCallcenterQueueParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCallcenterQueueParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchCallcenterQueueParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.QueueParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, CallcenterQueuesParams: &params, Id: &param.Queue.Id}
}

func UpdateCallcenterQueueParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCallcenterQueueParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateCallcenterQueueParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.QueueParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, CallcenterQueuesParams: &params, Id: &param.Queue.Id}
}

func AddCallcenterQueue(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	queue := pbxcache.GetCallcenterQueueByName(data.Name)
	if queue != nil {
		return webStruct.UserResponse{Error: "queue already exists", MessageType: data.Event}
	}
	queue, err := pbxcache.SetConfCallcenterQueue(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.Queue)
	items[queue.Id] = queue

	return webStruct.UserResponse{CallcenterQueues: &items, MessageType: data.Event}
}

func RenameCallcenterQueue(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	queue := pbxcache.GetCallcenterQueue(data.Id)
	if queue == nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateCallcenterQueue(queue, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Queue{queue.Id: queue}

	return webStruct.UserResponse{CallcenterQueues: &item, MessageType: data.Event}
}

func DelCallcenterQueue(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	queue := pbxcache.GetCallcenterQueue(data.Id)
	if queue == nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	res := pbxcache.DelCallcenterQueue(queue)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &queue.Id}
}

func GetCallcenterSettings(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetCallcenterSettings()
	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

/*
	func ImportCallcenterAgentsAdnTiers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
		fsesl.GetCallcenterAgents()
		fsesl.GetCallcenterTiers()
		//fsesl.GetCallcenterMembers()

		resp := GetCallcenterAgents(data, user)
		tiers := GetCallcenterTiers(data, user)
		resp.CallcenterTiers = tiers.CallcenterTiers
		resp.AltTotal = tiers.Total
		return resp
	}
*/
func ImportCallcenterAgentsAdnTiers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	fsesl.GetCallcenterAgents()
	fsesl.GetCallcenterTiers()
	return webStruct.UserResponse{}
}

func GetCallcenterAgents(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit < 25 || limit > 250 {
		limit = 25
	}
	offset := 0
	if data.DBRequest.Offset > 0 {
		offset = data.DBRequest.Offset * limit
	}

	items, total, err := pbxcache.GetCallcenterAgentsListsByForm(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{CallcenterAgents: &items, MessageType: data.Event, Total: &total}
}

func GetCallcenterAgentsList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, err := pbxcache.GetCallcenterAgentsList()
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{CallcenterAgentsList: &items, MessageType: data.Event}
}

func AddCallcenterAgent(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	agent := pbxcache.GetCallcenterAgentByName(data.Name)
	if agent != nil {
		return webStruct.UserResponse{Error: "agent already exists", MessageType: data.Event}
	}
	agent, err := pbxcache.SetConfCallcenterAgent(data.Name, "callback", "single_box", "single_box",
		"", "", "On Break", "Waiting", 0, 10, 10,
		10, 10,
		0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := []*mainStruct.Agent{agent}

	return webStruct.UserResponse{CallcenterAgents: &items, MessageType: data.Event}
}

func UpdateCallcenterAgent(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	agent := pbxcache.GetCallcenterAgent(data.Param.Id)
	if agent == nil {
		return webStruct.UserResponse{Error: "agent not found", MessageType: data.Event}
	}
	name := agent.GetItemNameByTag(data.Param.Name)
	if name == "" {
		return webStruct.UserResponse{Error: "field not found", MessageType: data.Event}
	}

	ok, err := pbxcache.UpdateCallcenterAgent(agent, data.Param.Name, data.Param.Value, eventChannel)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	if !ok {
		return webStruct.UserResponse{Error: "not updated", MessageType: data.Event}
	}

	tempAgent := mainStruct.CCAgentLight{Id: agent.Id}
	f := reflect.ValueOf(&tempAgent).Elem().FieldByName(name)

	switch f.Type().Name() {
	case "string":
		f.SetString(data.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(data.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(data.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}

	return webStruct.UserResponse{CallcenterAgent: &tempAgent, MessageType: data.Event}
}

func DelCallcenterAgent(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "agent not found", MessageType: data.Event}
	}
	agent := pbxcache.GetCallcenterAgent(data.Id)
	if agent == nil {
		return webStruct.UserResponse{Error: "agent not found", MessageType: data.Event}
	}

	res := pbxcache.DelCallcenterAgent(agent)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &agent.Id}
}

func GetCallcenterTiers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 250 {
		limit = 25
	}
	offset := data.DBRequest.Offset * limit

	items, total, err := pbxcache.GetCallcenterTiersListsByForm(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{CallcenterTiers: &items, MessageType: data.Event, Total: &total}
}

func GetCallcenterTiersList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, err := pbxcache.GetCallcenterTiersList()
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{CallcenterTiersList: &items, MessageType: data.Event}
}

func AddCallcenterTier(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	var err error
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "tier not found", MessageType: data.Event}
	}
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	queue := pbxcache.GetCallcenterQueue(data.Id)
	if queue == nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}
	agent := pbxcache.GetCallcenterAgentByName(data.Name)
	if agent == nil {
		return webStruct.UserResponse{Error: "agent not found", MessageType: data.Event}
	}
	tier := pbxcache.GetCallcenterTierByName(queue.Name + data.Name)
	if tier != nil {
		return webStruct.UserResponse{Error: "tier already exists", MessageType: data.Event}
	}
	tier, err = pbxcache.SetConfCallcenterTier(queue.Name, data.Name, "Ready", 1, 1)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := []*mainStruct.Tier{tier}

	return webStruct.UserResponse{CallcenterTiers: &items, MessageType: data.Event}
}

func UpdateCallcenterTier(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	tier := pbxcache.GetCallcenterTier(data.Param.Id)
	if tier == nil {
		return webStruct.UserResponse{Error: "tier not found", MessageType: data.Event}
	}
	name := tier.GetItemNameByTag(data.Param.Name)
	if name == "" {
		return webStruct.UserResponse{Error: "field not found", MessageType: data.Event}
	}

	ok, err := pbxcache.UpdateCallcenterTier(tier, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	if !ok {
		return webStruct.UserResponse{Error: "not updated", MessageType: data.Event}
	}

	tempTier := mainStruct.CCTierLight{Id: tier.Id}
	f := reflect.ValueOf(&tempTier).Elem().FieldByName(name)

	switch f.Type().Name() {
	case "string":
		f.SetString(data.Param.Value)
	case "int":
		fallthrough
	case "int64":
		res, err := strconv.ParseInt(data.Param.Value, 10, 64)
		if err == nil {
			f.SetInt(res)
		}
	case "bool":
		res, err := strconv.ParseBool(data.Param.Value)
		if err == nil {
			f.SetBool(res)
		}
	}

	return webStruct.UserResponse{CallcenterTier: &tempTier, MessageType: data.Event}
}

func DelCallcenterTier(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "tier not found", MessageType: data.Event}
	}
	tier := pbxcache.GetCallcenterTier(data.Id)
	if tier == nil {
		return webStruct.UserResponse{Error: "tier not found", MessageType: data.Event}
	}

	res := pbxcache.DelCallcenterTier(tier)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &tier.Id}
}

func runQueueCommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	queue := pbxcache.GetCallcenterQueue(data.Id)
	if queue == nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}

	if data.Name == "" || !mainStruct.CheckQueueCommand(data.Name) {
		return webStruct.UserResponse{Error: "unknown command", MessageType: data.Event}
	}

	command := fmt.Sprintf("callcenter_config queue %s %s", data.Name, queue.Name)

	_, err := fsesl.SendBgapiCmd(command)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func runCallcenterQueueCommand(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	queueI, err := intermediateDB.GetByIdFromDB(&altStruct.ConfigCallcenterQueue{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}
	queue, ok := queueI.(altStruct.ConfigCallcenterQueue)
	if !ok {
		return webStruct.UserResponse{Error: "queue not found", MessageType: data.Event}
	}
	if data.Name == "" || !mainStruct.CheckQueueCommand(data.Name) {
		return webStruct.UserResponse{Error: "unknown command", MessageType: data.Event}
	}

	command := fmt.Sprintf("callcenter_config queue %s %s", data.Name, queue.Name)

	_, err = fsesl.SendBgapiCmd(command)
	if err != nil {
		return webStruct.UserResponse{MessageType: data.Event, Error: err.Error()}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func GetCallcenterMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 250 {
		limit = 25
	}
	offset := data.DBRequest.Offset * limit

	items, total, err := pbxcache.GetCallcenterMembersListsByForm(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{CallcenterMembers: &items, MessageType: data.Event, Total: &total}
}

func DelCallcenterMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Uuid == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	member := pbxcache.GetCallcenterMember(data.Uuid)
	if member == nil {
		return webStruct.UserResponse{Error: "member not found", MessageType: data.Event}
	}

	res := pbxcache.DelCallcenterMember(member)
	if res == "" {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Uuid: &member.Uuid}
}

func getConfigOdbcCdr(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetOdbcCdrSettings()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	tables := pbxcache.GetOdbcCdrTables()
	fields := pbxcache.GetOdbcCdrFields()

	return webStruct.UserResponse{Parameters: &items, OdbcCdrTable: &tables, OdbcCdrTableField: &fields, MessageType: data.Event}
}

func addConfigOdbcCdrParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfOdbcCdrSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func switchOdbcCdrParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOdbcCdrParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchOdbcCdrParam(param, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}
	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func updateOdbcCdrParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOdbcCdrParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateOdbcCdrParam(param, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}
	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func delConfigOdbcCdrParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOdbcCdrParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelOdbcCdrParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func addConfigOdbcCdrTable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Table.Id != 0 || data.Table.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	table, err := pbxcache.SetConfOdbcCdrTable(data.Table.Name, data.Table.LogLeg)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Table{table.Id: table}
	return webStruct.UserResponse{MessageType: data.Event, OdbcCdrTable: &item}
}

func updateOdbcCdrTable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Table.Id == 0 || data.Table.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	table := pbxcache.GetOdbcCdrTable(data.Table.Id)
	if table == nil {
		return webStruct.UserResponse{Error: "table not update", MessageType: data.Event}
	}

	err := pbxcache.UpdateOdbcCdrTable(table, data.Table.Name, data.Table.LogLeg)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Table{table.Id: table}
	return webStruct.UserResponse{MessageType: data.Event, OdbcCdrTable: &item}
}

func switchOdbcCdrTable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Table.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	table := pbxcache.GetOdbcCdrTable(data.Table.Id)
	if table == nil {
		return webStruct.UserResponse{Error: "table not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchOdbcCdrTable(table, data.Table.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Table{table.Id: table}
	return webStruct.UserResponse{MessageType: data.Event, OdbcCdrTable: &item}
}

func delConfigOdbcCdrTable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Table.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	table := pbxcache.GetOdbcCdrTable(data.Table.Id)
	if table == nil {
		return webStruct.UserResponse{Error: "table not found", MessageType: data.Event}
	}

	res := pbxcache.DelOdbcCdrTable(table)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func addConfigOdbcCdrField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.OdbcCdrField.Id != 0 || data.OdbcCdrField.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	table := pbxcache.GetOdbcCdrTable(data.Id)
	if table == nil {
		return webStruct.UserResponse{Error: "table not found", MessageType: data.Event}
	}

	field, err := pbxcache.SetConfOdbcCdrTableField(table, data.OdbcCdrField.Name, data.OdbcCdrField.ChanVarName)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := &map[int64]map[int64]*mainStruct.ODBCField{table.Id: {field.Id: field}}
	return webStruct.UserResponse{MessageType: data.Event, OdbcCdrTableField: item}
}

func switchOdbcCdrField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.OdbcCdrField.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field := pbxcache.GetOdbcCdrField(data.OdbcCdrField.Id)
	if field == nil {
		return webStruct.UserResponse{Error: "field not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchOdbcCdrField(field, data.OdbcCdrField.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := &map[int64]map[int64]*mainStruct.ODBCField{field.Table.Id: {field.Id: field}}
	return webStruct.UserResponse{MessageType: data.Event, OdbcCdrTableField: item}
}

func updateOdbcCdrField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.OdbcCdrField.Id == 0 || data.OdbcCdrField.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field := pbxcache.GetOdbcCdrField(data.OdbcCdrField.Id)
	if field == nil {
		return webStruct.UserResponse{Error: "field not update", MessageType: data.Event}
	}

	err := pbxcache.UpdateOdbcCdrField(field, data.OdbcCdrField.Name, data.OdbcCdrField.ChanVarName)
	if err != nil {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	item := &map[int64]map[int64]*mainStruct.ODBCField{field.Table.Id: {field.Id: field}}
	return webStruct.UserResponse{MessageType: data.Event, OdbcCdrTableField: item}
}

func delConfigOdbcCdrField(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.OdbcCdrField.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	field := pbxcache.GetOdbcCdrField(data.OdbcCdrField.Id)
	if field == nil {
		return webStruct.UserResponse{Error: "field not found", MessageType: data.Event}
	}

	res := pbxcache.DelOdbcCdrField(field)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &field.Table.Id, AffectedId: &res}
}

func GetLcr(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigLcr()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetLcrProfiles()

	return webStruct.UserResponse{Parameters: &items, LcrProfiles: &profiles, MessageType: data.Event}
}

func GetLcrProfileParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetLcrProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, LcrProfileParams: &params, Id: &profile.Id}
}

func UpdateLcrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateLcrSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchLcrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchLcrSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddLcrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigLcrSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelLcrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelLcrSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddLcrProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetLcrProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigLcrProfileParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.LcrProfileParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, LcrProfileParams: &item, Id: &profile.Id}
}

func DelLcrProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelLcrProfileParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchLcrProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchLcrProfileParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.LcrProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, LcrProfileParams: &params, Id: &param.Profile.Id}
}

func UpdateLcrProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLcrProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateLcrProfileParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.LcrProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, LcrProfileParams: &params, Id: &param.Profile.Id}
}

func AddLcrProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetLcrProfileByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigLcrProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.LcrProfile)
	items[res.Id] = res

	return webStruct.UserResponse{LcrProfiles: &items, MessageType: data.Event}
}

func UpdateLcrProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetLcrProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateLcrProfile(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.LcrProfile{profile.Id: profile}

	return webStruct.UserResponse{LcrProfiles: &item, MessageType: data.Event}
}

func DelLcrProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetLcrProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelLcrProfile(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetShout(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigShout()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateShoutParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetShoutSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateShoutSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchShoutParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetShoutSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchShoutSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddShoutParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetShoutSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigShoutSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelShoutParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetShoutSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelShoutSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetRedis(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigRedis()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateRedisParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetRedisSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateRedisSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchRedisParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetRedisSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchRedisSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddRedisParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetRedisSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigRedisSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelRedisParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetRedisSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelRedisSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetNibblebill(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigNibblebill()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateNibblebillParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetNibblebillSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateNibblebillSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchNibblebillParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetNibblebillSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchNibblebillSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddNibblebillParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetNibblebillSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigNibblebillSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelNibblebillParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetNibblebillSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelNibblebillSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetDb(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigDb()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateDbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDbSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchDbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDbSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddDbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDbSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigDbSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelDbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelDbSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetMemcache(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigMemcache()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateMemcacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMemcacheSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateMemcacheSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchMemcacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMemcacheSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchMemcacheSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddMemcacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMemcacheSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigMemcacheSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelMemcacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMemcacheSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelMemcacheSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetAvmd(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigAvmd()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateAvmdParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAvmdSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateAvmdSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchAvmdParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAvmdSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchAvmdSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddAvmdParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAvmdSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigAvmdSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelAvmdParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAvmdSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelAvmdSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetTtsCommandline(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigTtsCommandline()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateTtsCommandlineParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetTtsCommandlineSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateTtsCommandlineSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchTtsCommandlineParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetTtsCommandlineSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchTtsCommandlineSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddTtsCommandlineParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetTtsCommandlineSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigTtsCommandlineSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelTtsCommandlineParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetTtsCommandlineSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelTtsCommandlineSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetCdrMongodb(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigCdrMongodb()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateCdrMongodbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrMongodbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateCdrMongodbSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchCdrMongodbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrMongodbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchCdrMongodbSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddCdrMongodbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrMongodbSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigCdrMongodbSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelCdrMongodbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCdrMongodbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelCdrMongodbSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetHttpCache(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigHttpCache()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateHttpCacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetHttpCacheSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateHttpCacheSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchHttpCacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetHttpCacheSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchHttpCacheSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddHttpCacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetHttpCacheSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigHttpCacheSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelHttpCacheParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetHttpCacheSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelHttpCacheSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetOpus(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigOpus()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateOpusParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpusSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateOpusSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchOpusParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpusSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchOpusSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddOpusParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpusSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigOpusSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelOpusParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpusSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelOpusSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetPython(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigPython()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdatePythonParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPythonSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdatePythonSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchPythonParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPythonSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchPythonSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddPythonParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPythonSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigPythonSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelPythonParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPythonSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPythonSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetAlsa(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigAlsa()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateAlsaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAlsaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateAlsaSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchAlsaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAlsaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchAlsaSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddAlsaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAlsaSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigAlsaSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelAlsaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAlsaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelAlsaSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetAmr(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigAmr()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateAmrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateAmrSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchAmrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchAmrSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddAmrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigAmrSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelAmrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelAmrSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetAmrwb(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigAmrwb()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateAmrwbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrwbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateAmrwbSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchAmrwbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrwbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchAmrwbSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddAmrwbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrwbSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigAmrwbSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelAmrwbParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetAmrwbSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelAmrwbSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetCepstral(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigCepstral()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateCepstralParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCepstralSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateCepstralSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchCepstralParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCepstralSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchCepstralSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddCepstralParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCepstralSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigCepstralSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelCepstralParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCepstralSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelCepstralSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetCidlookup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigCidlookup()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateCidlookupParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCidlookupSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateCidlookupSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchCidlookupParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCidlookupSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchCidlookupSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddCidlookupParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCidlookupSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigCidlookupSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelCidlookupParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCidlookupSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelCidlookupSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetCurl(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigCurl()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateCurlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCurlSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateCurlSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchCurlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCurlSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchCurlSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddCurlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCurlSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigCurlSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelCurlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetCurlSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelCurlSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetDialplanDirectory(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigDialplanDirectory()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateDialplanDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDialplanDirectorySetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDialplanDirectorySetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchDialplanDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDialplanDirectorySetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDialplanDirectorySetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddDialplanDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDialplanDirectorySettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigDialplanDirectorySetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelDialplanDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDialplanDirectorySetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelDialplanDirectorySetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetEasyroute(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigEasyroute()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateEasyrouteParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEasyrouteSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateEasyrouteSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchEasyrouteParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEasyrouteSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchEasyrouteSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddEasyrouteParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEasyrouteSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigEasyrouteSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelEasyrouteParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEasyrouteSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelEasyrouteSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetErlangEvent(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigErlangEvent()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateErlangEventParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetErlangEventSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateErlangEventSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchErlangEventParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetErlangEventSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchErlangEventSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddErlangEventParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetErlangEventSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigErlangEventSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelErlangEventParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetErlangEventSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelErlangEventSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetEventMulticast(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigEventMulticast()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateEventMulticastParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEventMulticastSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateEventMulticastSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchEventMulticastParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEventMulticastSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchEventMulticastSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddEventMulticastParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEventMulticastSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigEventMulticastSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelEventMulticastParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetEventMulticastSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelEventMulticastSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetFax(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigFax()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateFaxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFaxSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateFaxSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchFaxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFaxSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchFaxSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddFaxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFaxSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigFaxSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelFaxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFaxSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelFaxSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetLua(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigLua()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateLuaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLuaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateLuaSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchLuaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLuaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchLuaSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddLuaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLuaSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigLuaSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelLuaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetLuaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelLuaSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetMongo(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigMongo()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateMongoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMongoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateMongoSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchMongoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMongoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchMongoSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddMongoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMongoSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigMongoSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelMongoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMongoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelMongoSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetMsrp(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigMsrp()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateMsrpParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMsrpSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateMsrpSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchMsrpParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMsrpSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchMsrpSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddMsrpParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMsrpSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigMsrpSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelMsrpParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetMsrpSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelMsrpSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetOreka(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigOreka()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateOrekaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOrekaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateOrekaSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchOrekaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOrekaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchOrekaSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddOrekaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOrekaSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigOrekaSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelOrekaParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOrekaSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelOrekaSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetPerl(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigPerl()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdatePerlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPerlSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdatePerlSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchPerlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPerlSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchPerlSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddPerlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPerlSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigPerlSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelPerlParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPerlSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPerlSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetPocketsphinx(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigPocketsphinx()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdatePocketsphinxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPocketsphinxSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdatePocketsphinxSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchPocketsphinxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPocketsphinxSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchPocketsphinxSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddPocketsphinxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPocketsphinxSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigPocketsphinxSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelPocketsphinxParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPocketsphinxSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPocketsphinxSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetSangomaCodec(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigSangomaCodec()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateSangomaCodecParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSangomaCodecSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateSangomaCodecSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchSangomaCodecParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSangomaCodecSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchSangomaCodecSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddSangomaCodecParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSangomaCodecSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigSangomaCodecSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelSangomaCodecParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSangomaCodecSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelSangomaCodecSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetSndfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigSndfile()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateSndfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSndfileSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateSndfileSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchSndfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSndfileSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchSndfileSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddSndfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSndfileSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigSndfileSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelSndfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetSndfileSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelSndfileSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetXmlCdr(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigXmlCdr()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateXmlCdrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlCdrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateXmlCdrSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchXmlCdrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlCdrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchXmlCdrSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddXmlCdrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlCdrSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigXmlCdrSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelXmlCdrParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlCdrSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelXmlCdrSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetXmlRpc(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigXmlRpc()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateXmlRpcParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlRpcSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateXmlRpcSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchXmlRpcParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlRpcSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchXmlRpcSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddXmlRpcParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlRpcSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigXmlRpcSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelXmlRpcParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetXmlRpcSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelXmlRpcSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetZeroconf(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigZeroconf()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{Parameters: &items, MessageType: data.Event}
}

func UpdateZeroconfParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetZeroconfSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateZeroconfSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchZeroconfParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetZeroconfSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchZeroconfSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddZeroconfParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetZeroconfSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigZeroconfSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelZeroconfParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetZeroconfSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelZeroconfSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetPostSwitch(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	params, exists := pbxcache.GetConfigPostSwitch()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	keyBindings, exists := pbxcache.GetConfigPostSwitchCliKeybinding()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	ptimes, exists := pbxcache.GetConfigPostSwitchDefaultPtime()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{PostSwitchParameters: &params, PostSwitchCliKeybinding: &keyBindings, PostSwitchDefaultPtime: &ptimes, MessageType: data.Event}
}

func UpdatePostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdatePostSwitchSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchPostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchPostSwitchSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddPostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigPostSwitchSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelPostSwitchParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPostSwitchSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func UpdatePostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchCliKeybinding(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdatePostSwitchCliKeybinding(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, PostSwitchCliKeybinding: &item}
}

func SwitchPostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchCliKeybinding(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchPostSwitchCliKeybinding(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, PostSwitchCliKeybinding: &item}
}

func AddPostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchCliKeybindingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigPostSwitchCliKeybinding(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, PostSwitchCliKeybinding: &item}
}

func DelPostSwitchCliKeybinding(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchCliKeybinding(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPostSwitchCliKeybinding(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func UpdatePostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchDefaultPtime(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdatePostSwitchDefaultPtime(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.DefaultPtime{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, PostSwitchDefaultPtime: &item}
}

func SwitchPostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchDefaultPtime(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchPostSwitchDefaultPtime(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.DefaultPtime{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, PostSwitchDefaultPtime: &item}
}

func AddPostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchDefaultPtimeByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigPostSwitchDefaultPtime(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.DefaultPtime{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, PostSwitchDefaultPtime: &item}
}

func DelPostSwitchDefaultPtime(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetPostSwitchDefaultPtime(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelPostSwitchDefaultPtime(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func GetDistributorConfig(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetDistributorLists()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}

	return webStruct.UserResponse{DistributorLists: &items, MessageType: data.Event}
}

func AddDistributorList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	res, err := pbxcache.SetConfDistributorList(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.DistributorList)
	items[res.Id] = res

	return webStruct.UserResponse{DistributorLists: &items, MessageType: data.Event}
}

func DelDistributorList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "distributor list not found", MessageType: data.Event}
	}
	ok := pbxcache.DelDistributorList(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete distributor list", MessageType: data.Event}
	}

	return webStruct.UserResponse{Id: &data.Id, MessageType: data.Event}
}

func UpdateDistributorList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "distributor list not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	err := pbxcache.UpdateDistributorList(data.Id, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	res := pbxcache.GetDistributorList(data.Id)
	items := make(map[int64]*mainStruct.DistributorList)
	items[res.Id] = res

	return webStruct.UserResponse{DistributorLists: &items, MessageType: data.Event}
}

func GetDistributorNodes(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	list := pbxcache.GetDistributorList(data.Id)
	if list == nil {
		return webStruct.UserResponse{Error: "list not found", MessageType: data.Event}
	}

	nodes := list.Nodes.GetList()

	return webStruct.UserResponse{MessageType: data.Event, DistributorListNodes: &nodes}
}

func DelDistributorNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.DistributorNode.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	node := pbxcache.GetDistributorNode(data.DistributorNode.Id)
	if node == nil {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	res := pbxcache.DelDistributorNode(node)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: node.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func UpdateDistributorNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.DistributorNode.Id == 0 || data.DistributorNode.Name == "" || data.DistributorNode.Weight == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	node := pbxcache.GetDistributorNode(data.DistributorNode.Id)
	if node == nil {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDistributorNode(node, data.DistributorNode.Name, data.DistributorNode.Weight)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := webStruct.Item{Id: node.Id, Name: node.Name, Weight: node.Weight, Enabled: node.Enabled}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchDistributorNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.DistributorNode.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	node := pbxcache.GetDistributorNode(data.DistributorNode.Id)
	if node == nil {
		return webStruct.UserResponse{Error: "node not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDistributorNode(node, data.DistributorNode.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := webStruct.Item{Id: node.Id, Name: node.Name, Weight: node.Weight, Enabled: node.Enabled}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func AddDistributorNode(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.DistributorNode.Name == "" || data.DistributorNode.Weight == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	list := pbxcache.GetDistributorList(data.Id)
	if list == nil {
		return webStruct.UserResponse{Error: "distributor list not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfDistributorNode(list, data.DistributorNode.Name, data.DistributorNode.Weight)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	node := list.Nodes.GetById(res)
	if node == nil {
		return webStruct.UserResponse{Error: "not added", MessageType: data.Event}
	}
	nodes := map[int64]*mainStruct.DistributorNode{list.Id: node}

	return webStruct.UserResponse{DistributorListNodes: &nodes, MessageType: data.Event}
}

func GetDirectory(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigDirectory()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetDirectoryProfiles()

	return webStruct.UserResponse{Parameters: &items, DirectoryProfiles: &profiles, MessageType: data.Event}
}

func GetDirectoryProfileParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetDirectoryProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, DirectoryProfileParams: &params, Id: &profile.Id}
}

func UpdateDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectorySetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDirectorySetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectorySetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDirectorySetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectorySettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigDirectorySetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelDirectoryParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectorySetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelDirectorySetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddDirectoryProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetDirectoryProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigDirectoryProfileParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.DirectoryProfileParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, DirectoryProfileParams: &item, Id: &profile.Id}
}

func DelDirectoryProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectoryProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelDirectoryProfileParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchDirectoryProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectoryProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDirectoryProfileParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.DirectoryProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, DirectoryProfileParams: &params, Id: &param.Profile.Id}
}

func UpdateDirectoryProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDirectoryProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDirectoryProfileParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.DirectoryProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, DirectoryProfileParams: &params, Id: &param.Profile.Id}
}

func AddDirectoryProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetDirectoryProfileByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigDirectoryProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.DirectoryProfile)
	items[res.Id] = res

	return webStruct.UserResponse{DirectoryProfiles: &items, MessageType: data.Event}
}

func UpdateDirectoryProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetDirectoryProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateDirectoryProfile(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.DirectoryProfile{profile.Id: profile}

	return webStruct.UserResponse{DirectoryProfiles: &item, MessageType: data.Event}
}

func DelDirectoryProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetDirectoryProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelDirectoryProfile(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetFifo(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigFifo()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetFifoFifos()

	return webStruct.UserResponse{Parameters: &items, FifoFifo: &profiles, MessageType: data.Event}
}

func GetFifoFifoMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetFifoFifo(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, FifoFifosMembers: &params, Id: &profile.Id}
}

func UpdateFifoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateFifoSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchFifoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchFifoSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddFifoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigFifoSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelFifoParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelFifoSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddFifoFifoMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.FifoFifoMember.Id != 0 || data.FifoFifoMember.Body == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetFifoFifo(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigFifoFifoParam(profile, data.FifoFifoMember.Timeout, data.FifoFifoMember.Simo, data.FifoFifoMember.Lag, data.FifoFifoMember.Body)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.FifoFifoMember{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, FifoFifosMembers: &item, Id: &profile.Id}
}

func DelFifoFifoMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.FifoFifoMember.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoFifoParam(data.FifoFifoMember.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelFifoFifoParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func UpdateFifoFifoImportance(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateFifoFifoImportance(data.Id, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Value: data.Value}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item}
}

func SwitchFifoFifoMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.FifoFifoMember.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoFifoParam(data.FifoFifoMember.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchFifoFifoParam(param, data.FifoFifoMember.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.FifoFifoMember{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, FifoFifosMembers: &params, Id: &param.Fifo.Id}
}

func UpdateFifoFifoMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.FifoFifoMember.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetFifoFifoParam(data.FifoFifoMember.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "member not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateFifoFifoParam(param, data.FifoFifoMember.Timeout, data.FifoFifoMember.Simo, data.FifoFifoMember.Lag, data.FifoFifoMember.Body)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.FifoFifoMember{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, FifoFifosMembers: &params, Id: &param.Fifo.Id}
}

func AddFifoFifo(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	if data.Importance == "" {
		data.Importance = "0"
	}

	profile := pbxcache.GetFifoFifoByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "fifo already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigFifoFifo(data.Name, data.Importance)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.FifoFifo)
	items[res.Id] = res

	return webStruct.UserResponse{FifoFifo: &items, MessageType: data.Event}
}

func UpdateFifoFifo(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "fifo not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetFifoFifo(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "fifo not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateFifoFifo(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.FifoFifo{profile.Id: profile}

	return webStruct.UserResponse{FifoFifo: &item, MessageType: data.Event}
}

func DelFifoFifo(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "fifo not found", MessageType: data.Event}
	}

	profile := pbxcache.GetFifoFifo(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "fifo not found", MessageType: data.Event}
	}

	res := pbxcache.DelFifoFifo(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetOpal(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigOpal()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetOpalListeners()

	return webStruct.UserResponse{Parameters: &items, OpalListeners: &profiles, MessageType: data.Event}
}

func GetOpalListenerParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetOpalListener(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, OpalListenerParams: &params, Id: &profile.Id}
}

func UpdateOpalParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateOpalSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchOpalParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchOpalSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddOpalParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigOpalSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelOpalParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelOpalSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddOpalListenerParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetOpalListener(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigOpalListenerParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.OpalListenerParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, OpalListenerParams: &item, Id: &profile.Id}
}

func DelOpalListenerParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalListenerParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelOpalListenerParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchOpalListenerParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalListenerParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchOpalListenerParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.OpalListenerParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, OpalListenerParams: &params, Id: &param.Listener.Id}
}

func UpdateOpalListenerParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOpalListenerParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateOpalListenerParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.OpalListenerParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, OpalListenerParams: &params, Id: &param.Listener.Id}
}

func AddOpalListener(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetOpalListenerByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigOpalListener(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.OpalListener)
	items[res.Id] = res

	return webStruct.UserResponse{OpalListeners: &items, MessageType: data.Event}
}

func UpdateOpalListener(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetOpalListener(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateOpalListener(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.OpalListener{profile.Id: profile}

	return webStruct.UserResponse{OpalListeners: &item, MessageType: data.Event}
}

func DelOpalListener(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetOpalListener(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelOpalListener(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetOsp(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigOsp()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetOspProfiles()

	return webStruct.UserResponse{Parameters: &items, OspProfiles: &profiles, MessageType: data.Event}
}

func GetOspProfileParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetOspProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, OspProfileParams: &params, Id: &profile.Id}
}

func UpdateOspParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateOspSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchOspParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchOspSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddOspParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigOspSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelOspParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelOspSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddOspProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetOspProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigOspProfileParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.OspProfileParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, OspProfileParams: &item, Id: &profile.Id}
}

func DelOspProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelOspProfileParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchOspProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchOspProfileParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.OspProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, OspProfileParams: &params, Id: &param.Profile.Id}
}

func UpdateOspProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetOspProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateOspProfileParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.OspProfileParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, OspProfileParams: &params, Id: &param.Profile.Id}
}

func AddOspProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetOspProfileByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigOspProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.OspProfile)
	items[res.Id] = res

	return webStruct.UserResponse{OspProfiles: &items, MessageType: data.Event}
}

func UpdateOspProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetOspProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateOspProfile(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.OspProfile{profile.Id: profile}

	return webStruct.UserResponse{OspProfiles: &item, MessageType: data.Event}
}

func DelOspProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetOspProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelOspProfile(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetUnicall(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigUnicall()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	profiles, _ := pbxcache.GetUnicallSpans()

	return webStruct.UserResponse{Parameters: &items, UnicallSpans: &profiles, MessageType: data.Event}
}

func GetUnicallSpanParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetUnicallSpan(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, UnicallSpanParams: &params, Id: &profile.Id}
}

func UpdateUnicallParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUnicallSetting(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func SwitchUnicallParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchUnicallSetting(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func AddUnicallParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSettingByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfigUnicallSetting(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.Param{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, Parameters: &item}
}

func DelUnicallParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSetting(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res := pbxcache.DelUnicallSetting(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &res}
}

func AddUnicallSpanParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetUnicallSpan(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfigUnicallSpanParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.UnicallSpanParam{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, UnicallSpanParams: &item, Id: &profile.Id}
}

func DelUnicallSpanParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSpanParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res := pbxcache.DelUnicallSpanParam(param)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &res}
}

func SwitchUnicallSpanParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSpanParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchUnicallSpanParam(param, data.Param.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.UnicallSpanParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, UnicallSpanParams: &params, Id: &param.Span.Id}
}

func UpdateUnicallSpanParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUnicallSpanParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUnicallSpanParam(param, data.Param.Name, data.Param.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.UnicallSpanParam{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, UnicallSpanParams: &params, Id: &param.Span.Id}
}

func AddUnicallSpan(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetUnicallSpanByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfigUnicallSpan(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.UnicallSpan)
	items[res.Id] = res

	return webStruct.UserResponse{UnicallSpans: &items, MessageType: data.Event}
}

func UpdateUnicallSpan(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetUnicallSpan(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	_, err := pbxcache.UpdateUnicallSpan(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.UnicallSpan{profile.Id: profile}

	return webStruct.UserResponse{UnicallSpans: &item, MessageType: data.Event}
}

func DelUnicallSpan(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetUnicallSpan(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	res := pbxcache.DelUnicallSpan(profile)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func GetConference(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items, exists := pbxcache.GetConfigConference()
	if !exists {
		return webStruct.UserResponse{Exists: &exists, MessageType: data.Event}
	}
	callerControls, _ := pbxcache.GetConferenceCallerControls()
	profiles, _ := pbxcache.GetConferenceProfiles()
	chatPermissionsprofiles, _ := pbxcache.GetConferenceChatPermissionsProfiles()

	return webStruct.UserResponse{ConferenceRooms: &items, ConferenceProfiles: &profiles, ConferenceCallerControlGroups: &callerControls, ConferenceChatPermissionsProfiles: &chatPermissionsprofiles, MessageType: data.Event}
}

func GetConferenceProfileParameters(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	params := profile.Params.GetList()

	return webStruct.UserResponse{MessageType: data.Event, ConferenceProfileParams: &params, Id: &profile.Id}
}

func UpdateConferenceRoom(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceRoom(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(param, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item := map[int64]*mainStruct.ConfigConferenceAdvertiseRooms{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceRooms: &item}
}

func SwitchConferenceRoom(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceRoom(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(param, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ConfigConferenceAdvertiseRooms{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceRooms: &item}
}

func AddConferenceRoom(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceRoomByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	param, err := pbxcache.SetConfConferenceAdvertise(data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ConfigConferenceAdvertiseRooms{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceRooms: &item}
}

func DelConferenceRoom(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceRoom(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(param)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &param.Id}
}

func AddConferenceProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	param := profile.Params.GetByName(data.Param.Name)
	if param != nil {
		return webStruct.UserResponse{Error: "param name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfConferenceProfileParam(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ConfigConferenceProfilesParams{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceProfileParams: &item, Id: &profile.Id}
}

func DelConferenceProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(param)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: param.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &param.Parent.Parent.Id}
}

func SwitchConferenceProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(param, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	params := map[int64]*mainStruct.ConfigConferenceProfilesParams{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceProfileParams: &params, Id: &param.Parent.Parent.Id}
}

func UpdateConferenceProfileParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetConferenceProfileParam(data.Param.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(param, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	params := map[int64]*mainStruct.ConfigConferenceProfilesParams{param.Id: param}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceProfileParams: &params, Id: &param.Parent.Parent.Id}
}

func AddConferenceProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceProfileByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfConferenceProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.ConfigConferenceProfiles)
	items[res.Id] = res

	return webStruct.UserResponse{ConferenceProfiles: &items, MessageType: data.Event}
}

func UpdateConferenceProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.ConfigConferenceProfiles{profile.Id: profile}

	return webStruct.UserResponse{ConferenceProfiles: &item, MessageType: data.Event}
}

func DelConferenceProfile(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(profile)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}

func AddConferenceCallerControl(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	group := pbxcache.GetConferenceCallerControlGroup(data.Id)
	if group == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	control := group.Controls.GetByName(data.Param.Name)
	if control != nil {
		return webStruct.UserResponse{Error: "control name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfConferenceCallerControlsGroupControl(group, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ConfigConferenceCallerControlsControls{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceCallerControl: &item, Id: &group.Id}
}

func DelConferenceCallerControl(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	control := pbxcache.GetConferenceCallerControl(data.Param.Id)
	if control == nil {
		return webStruct.UserResponse{Error: "control not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(control)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: control.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &control.Parent.Parent.Id}
}

func SwitchConferenceCallerControl(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	control := pbxcache.GetConferenceCallerControl(data.Param.Id)
	if control == nil {
		return webStruct.UserResponse{Error: "control not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(control, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	controls := map[int64]*mainStruct.ConfigConferenceCallerControlsControls{control.Id: control}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceCallerControl: &controls, Id: &control.Parent.Parent.Id}
}

func UpdateConferenceCallerControl(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	control := pbxcache.GetConferenceCallerControl(data.Param.Id)
	if control == nil {
		return webStruct.UserResponse{Error: "control not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(control, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	controls := map[int64]*mainStruct.ConfigConferenceCallerControlsControls{control.Id: control}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceCallerControl: &controls, Id: &control.Parent.Parent.Id}
}

func AddConferenceCallerControlGroup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	group := pbxcache.GetConferenceCallerControlGroup(data.Id)
	if group != nil {
		return webStruct.UserResponse{Error: "group already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfConferenceCallerControlsGroup(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.ConfigConferenceCallerControlsGroups)
	items[res.Id] = res

	return webStruct.UserResponse{ConferenceCallerControlGroups: &items, MessageType: data.Event}
}

func UpdateConferenceCallerControlGroup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	group := pbxcache.GetConferenceCallerControlGroup(data.Id)
	if group == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(group, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.ConfigConferenceCallerControlsGroups{group.Id: group}

	return webStruct.UserResponse{ConferenceCallerControlGroups: &item, MessageType: data.Event}
}

func DelConferenceCallerControlGroup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	group := pbxcache.GetConferenceCallerControlGroup(data.Id)
	if group == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(group)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &group.Id}
}

func GetConferenceCallerControls(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	group := pbxcache.GetConferenceCallerControlGroup(data.Id)
	if group == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	params := group.Controls.GetList()

	return webStruct.UserResponse{MessageType: data.Event, ConferenceCallerControl: &params, Id: &group.Id}
}

func GetConferenceChatPermissionUsers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceChatPermissionsProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	items := profile.Users.GetList()

	return webStruct.UserResponse{MessageType: data.Event, ConferenceChatPermissionsUsers: &items, Id: &profile.Id}
}

func AddConferenceChatPermissionUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Param.Id != 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceChatPermissionsProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	permissionUser := profile.Users.GetByName(data.Param.Name)
	if permissionUser != nil {
		return webStruct.UserResponse{Error: "user name already exists", MessageType: data.Event}
	}

	res, err := pbxcache.SetConfConferenceChatPermissionsUser(profile, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]*mainStruct.ConfigConferenceChatPermissionUsers{res.Id: res}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceChatPermissionsUsers: &item, Id: &profile.Id}
}

func DelConferenceChatPermissionUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	permissionUser := pbxcache.GetConferenceChatPermissionsUser(data.Param.Id)
	if permissionUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(permissionUser)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	item := webStruct.Item{Id: permissionUser.Id}
	return webStruct.UserResponse{MessageType: data.Event, Item: &item, Id: &permissionUser.Parent.Parent.Id}
}

func SwitchConferenceChatPermissionUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	permissionUser := pbxcache.GetConferenceChatPermissionsUser(data.Param.Id)
	if permissionUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	err := pbxcache.SwitchConfigRow(permissionUser, data.Param.Enabled)
	if err != nil {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	permissionUsers := map[int64]*mainStruct.ConfigConferenceChatPermissionUsers{permissionUser.Id: permissionUser}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceChatPermissionsUsers: &permissionUsers, Id: &permissionUser.Parent.Parent.Id}
}

func UpdateConferenceChatPermissionUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Param.Id == 0 || data.Param.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	permissionUser := pbxcache.GetConferenceChatPermissionsUser(data.Param.Id)
	if permissionUser == nil {
		return webStruct.UserResponse{Error: "permissionUser not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(permissionUser, data.Param.Name, data.Param.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	permissionUsers := map[int64]*mainStruct.ConfigConferenceChatPermissionUsers{permissionUser.Id: permissionUser}

	return webStruct.UserResponse{MessageType: data.Event, ConferenceChatPermissionsUsers: &permissionUsers, Id: &permissionUser.Parent.Parent.Id}
}

func AddConferenceChatPermission(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceChatPermissionsProfileByName(data.Name)
	if profile != nil {
		return webStruct.UserResponse{Error: "profile already exists", MessageType: data.Event}
	}
	res, err := pbxcache.SetConfConferenceChatPermissionsProfile(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.ConfigConferenceChatPermissions)
	items[res.Id] = res

	return webStruct.UserResponse{ConferenceChatPermissionsProfiles: &items, MessageType: data.Event}
}

func UpdateConferenceChatPermission(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceChatPermissionsProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateConfigRow(profile, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	item := map[int64]*mainStruct.ConfigConferenceChatPermissions{profile.Id: profile}

	return webStruct.UserResponse{ConferenceChatPermissionsProfiles: &item, MessageType: data.Event}
}

func DelConferenceChatPermission(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	profile := pbxcache.GetConferenceChatPermissionsProfile(data.Id)
	if profile == nil {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}

	ok := pbxcache.DelConfigRow(profile)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &profile.Id}
}
