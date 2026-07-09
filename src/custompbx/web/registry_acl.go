package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/webStruct"
)

func registerCoreACLEvents(r *handlerRegistry, overrides map[string]eventHandler) {
	mustRegisterAdmin(r, eventACLListsGet, getACLLists, overrides)
	mustRegisterAdmin(r, eventACLListAdd, addACLList, overrides)
	mustRegisterAdmin(r, eventACLListUpdateDefault, updateACLListDefault, overrides)
	mustRegisterAdmin(r, eventACLListUpdate, updateACLList, overrides)
	mustRegisterAdmin(r, eventACLListDelete, deleteACLList, overrides)
	mustRegisterAdmin(r, eventACLListConfigUpdateDefault, updateACLListConfigDefault, overrides)
	mustRegisterAdmin(r, eventACLNodesGet, getACLNodes, overrides)
	mustRegisterAdmin(r, eventACLNodeAdd, addACLNode, overrides)
	mustRegisterAdmin(r, eventACLNodeDelete, deleteACLNode, overrides)
	mustRegisterAdmin(r, eventACLNodeUpdate, updateACLNode, overrides)
	mustRegisterAdmin(r, eventACLNodeSwitch, switchACLNode, overrides)
	mustRegisterAdmin(r, eventACLNodeMove, moveACLNode, overrides)
}

func getACLLists(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigAclList{}, adminOnly())
}

func addACLList(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigAclList{Name: data.Name, Default: data.Default, Enabled: true, Parent: getConfParent(altData.GetConfNameByStruct(&altStruct.ConfigAclList{}))}, adminOnly())
}

func updateACLListDefault(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigAclList{Id: data.Id, Default: data.Default}, []string{"Default"}}, adminOnly())
}

func updateACLList(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigAclList{Id: data.Id, Name: data.Name}, []string{"Name"}}, adminOnly())
}

func deleteACLList(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigAclList{Id: data.Id}, adminOnly())
}

func updateACLListConfigDefault(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigAclList{Id: data.Id, Default: data.Value}, []string{"Default"}}, adminOnly())
}

func getACLNodes(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, getConfig, &altStruct.ConfigAclNode{}, adminOnly())
}

func addACLNode(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, setConfig, &altStruct.ConfigAclNode{Type: data.Node.Type, Cidr: data.Node.Cidr, Domain: data.Node.Domain, Enabled: true, Parent: &altStruct.ConfigAclList{Id: data.Id}}, adminOnly())
}

func deleteACLNode(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, delConfig, &altStruct.ConfigAclNode{Id: data.Id}, adminOnly())
}

func updateACLNode(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigAclNode{Id: data.Node.Id, Type: data.Node.Type, Cidr: data.Node.Cidr, Domain: data.Node.Domain}, []string{"Type", "Cidr", "Domain"}}, adminOnly())
}

func switchACLNode(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigAclNode{Id: data.Node.Id, Enabled: data.Node.Enabled}, []string{"Enabled"}}, adminOnly())
}

func moveACLNode(data *webStruct.MessageData) webStruct.UserResponse {
	return getUserForConfig(data, updateConfig, struct {
		S interface{}
		A []string
	}{&altStruct.ConfigAclNode{Id: data.Id, Position: data.CurrentIndex}, []string{"Position"}}, adminOnly())
}
