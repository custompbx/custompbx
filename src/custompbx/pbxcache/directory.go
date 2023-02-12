package pbxcache

import (
	"custompbx/cache"
	"custompbx/db"
	"custompbx/mainStruct"
	"errors"
)

func GetCachedDomains() []mainStruct.XMLDomain {
	return directory.Domains.XMLItems()
}

func GetCachedDomainUser(domainName, userName string) ([]mainStruct.XMLDomain, bool) {
	domain := directory.Domains.GetByName(domainName)
	if domain == nil {
		return []mainStruct.XMLDomain{}, false
	}
	user := domain.Users.GetByName(userName)
	if user == nil {
		return []mainStruct.XMLDomain{}, false
	}
	newXmlDomain := domain.XMLItems()
	newXmlDomain.XMLGroups = append([]interface{}{mainStruct.XMLGroup{Name: "CustomPbxGroupForFS1.6", XMLUsers: []*mainStruct.User{user.XMLItems()}}})

	return []mainStruct.XMLDomain{*newXmlDomain}, true
}

func GetCachedDomainGroup(domainName, groupName string) ([]mainStruct.XMLDomain, bool) {
	domain := directory.Domains.GetByName(domainName)
	if domain == nil {
		return []mainStruct.XMLDomain{}, false
	}
	group, ok := domain.Groups.GetByName(groupName)
	if !ok {
		return []mainStruct.XMLDomain{}, false
	}
	newGroup := group.XMLItems()
	emptyDomain := mainStruct.XMLDomain{Name: domain.Name, XMLGroups: []mainStruct.XMLGroup{{Name: newGroup.Name, XMLUsers: newGroup.XMLUsers}}}

	return []mainStruct.XMLDomain{emptyDomain}, true
}

func GetCachedDomainUsersList() []mainStruct.XMLDomain {
	return directory.Domains.XMLSafe()
}

func GetCachedDomainUsersListOnly() []mainStruct.XMLDomain {
	return directory.Domains.XMLFullSafe()
}

// SET
func SetDomain(domainName string) (*mainStruct.Domain, error) {
	if domainName == "" {
		return nil, errors.New("empty domain name")
	}
	if directory.Domains.HasName(domainName) {
		return nil, errors.New("domain name already exists")
	}
	res, err := db.SetDomain(domainName, cache.GetCurrentInstanceId())
	if err != nil {
		return nil, err
	}
	domain := &mainStruct.Domain{Id: res, Name: domainName, Enabled: true, Params: mainStruct.NewDomainParams(), Vars: mainStruct.NewDomainVars(), Users: mainStruct.NewUsers(), Groups: mainStruct.NewGroups()}
	directory.Domains.Set(domain)

	return domain, err
}

func SetDomainParameter(domain *mainStruct.Domain, paramName, paramValue string) (*mainStruct.DomainParam, error) {
	if domain == nil {
		return nil, errors.New("domain doesn't exists")
	}
	if domain.Params.HasName(paramName) {
		return nil, errors.New("parameter name already exists")
	}

	res, err := db.SetDomainParameter(domain.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}
	param := &mainStruct.DomainParam{Id: res, Name: paramName, Enabled: true, Value: paramValue, Domain: domain}
	domain.Params.Set(param)
	directory.DomainParams.Set(param)

	return param, err
}

func SetDomainVariable(domain *mainStruct.Domain, varName, varValue string) (*mainStruct.DomainVariable, error) {
	if domain == nil {
		return nil, errors.New("domain doesn't exists")
	}
	if domain.Vars.HasName(varName) {
		return nil, errors.New("variable name already exists")
	}
	res, err := db.SetDomainVariable(domain.Id, varName, varValue)
	if err != nil {
		return nil, err
	}
	variable := &mainStruct.DomainVariable{Id: res, Name: varName, Enabled: true, Value: varValue, Domain: domain}
	domain.Vars.Set(variable)
	directory.DomainVars.Set(variable)

	return variable, err
}

func SetDomainGroup(domain *mainStruct.Domain, groupName string) (*mainStruct.Group, error) {
	if domain == nil {
		return nil, errors.New("domain doesn't exists")
	}
	if domain.Groups.HasName(groupName) {
		return nil, errors.New("group name already exists")
	}
	res, err := db.SetDomainGroup(domain.Id, groupName)
	if err != nil {
		return nil, err
	}
	group := &mainStruct.Group{Id: res, Name: groupName, Enabled: true, Users: mainStruct.NewGroupUsers(), Domain: domain}
	domain.Groups.Set(group)
	directory.Groups.Set(group)
	return group, err
}

func SetDomainGroupUser(group *mainStruct.Group, userName string) (*mainStruct.GroupUser, error) {
	if group == nil {
		return nil, errors.New("group name doesn't exists")
	}
	if group.Users.HasName(userName) {
		return nil, errors.New("user in group already exists")
	}
	domain := group.Domain
	user := domain.Users.GetByName(userName)
	if user == nil {
		return nil, errors.New("user name doesn't exists")
	}
	res, err := db.SetGroupMember(group.Id, user.Id)
	if err != nil {
		return nil, err
	}
	groupUser := &mainStruct.GroupUser{Id: res, Name: userName, Enabled: true, Type: "pointer", Group: group, UserId: user.Id}
	group.Users.Set(groupUser)
	directory.GroupUsers.Set(groupUser)
	return groupUser, err
}

func SetDomainGroupNewUser(group *mainStruct.Group, user *mainStruct.User) (*mainStruct.GroupUser, error) {
	if group == nil || user == nil {
		return nil, errors.New("wrong data")
	}

	res, err := db.SetGroupMember(group.Id, user.Id)
	if err != nil {
		return nil, err
	}
	groupUser := &mainStruct.GroupUser{Id: res, Name: user.Name, Enabled: true, Type: "pointer", Group: group, UserId: user.Id}
	group.Users.Set(groupUser)
	directory.GroupUsers.Set(groupUser)
	return groupUser, err
}

func SetDomainUser(domain *mainStruct.Domain, userName, cidr, numberAlias string) (*mainStruct.User, error) {
	if domain == nil {
		return nil, errors.New("domain doesn't exists")
	}
	directoryUser := domain.Users.GetByName(userName)
	if directoryUser != nil {
		return nil, errors.New("user name already exists")
	}
	res, err := db.SetDomainUser(domain.Id, userName, cidr, numberAlias)
	if err != nil {
		return nil, err
	}
	user := &mainStruct.User{Name: userName, Enabled: true, Id: res, Cidr: cidr, NumberAlias: numberAlias, Params: mainStruct.NewUserParams(), Vars: mainStruct.NewUserVars(), Gateways: mainStruct.NewUserGateways(), Domain: domain}
	domain.Users.Set(user)
	directory.Users.Set(user)
	return user, err
}

func SetUserParameter(user *mainStruct.User, paramName, paramValue string) (*mainStruct.UserParam, error) {
	if user == nil {
		return nil, errors.New("user name doesn't exists")
	}
	if user.Params.HasName(paramName) {
		return nil, errors.New("parameter name already exists")
	}
	res, err := db.SetUserParameter(user.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}
	param := &mainStruct.UserParam{Id: res, Enabled: true, Name: paramName, Value: paramValue, User: user}
	user.Params.Set(param)
	directory.UserParams.Set(param)
	return param, err
}

func SetUserVariable(user *mainStruct.User, varName, varValue string) (*mainStruct.UserVariable, error) {
	if user == nil {
		return nil, errors.New("user name doesn't exists")
	}
	if user.Vars.HasName(varName) {
		return nil, errors.New("parameter name already exists")
	}
	res, err := db.SetUserVariable(user.Id, varName, varValue)
	if err != nil {
		return nil, err
	}
	variable := &mainStruct.UserVariable{Id: res, Enabled: true, Name: varName, Value: varValue, User: user}
	user.Vars.Set(variable)
	directory.UserVars.Set(variable)
	return variable, err
}

func SetUserGateway(user *mainStruct.User, gatewayName string) (*mainStruct.UserGateway, error) {
	if user == nil {
		return nil, errors.New("user name doesn't exists")
	}
	if user.Gateways.HasName(gatewayName) {
		return nil, errors.New("gateway name already exists")
	}
	res, err := db.SetUserGateway(user.Id, gatewayName)
	if err != nil {
		return nil, err
	}
	gateway := &mainStruct.UserGateway{Id: res, Enabled: true, Name: gatewayName, Params: mainStruct.NewGatewayParams(), Vars: mainStruct.NewGatewayVars(), User: user}
	user.Gateways.Set(gateway)
	directory.UserGateways.Set(gateway)
	return gateway, err
}

func SetUserGatewayParam(gateway *mainStruct.UserGateway, paramName, paramValue string) (*mainStruct.GatewayParam, error) {
	if gateway == nil {
		return nil, errors.New("gateway name doesn't exists")
	}

	if gateway.Params.HasName(paramName) {
		return nil, errors.New("parameter name already exists")
	}

	res, err := db.SetUserGatewayParam(gateway.Id, paramName, paramValue)
	if err != nil {
		return nil, err
	}
	param := &mainStruct.GatewayParam{Id: res, Enabled: true, Name: paramName, Value: paramValue, Gateway: gateway}
	gateway.Params.Set(param)
	directory.GatewayParams.Set(param)
	return param, err
}

func SetUserGatewayVar(gateway *mainStruct.UserGateway, varName, varValue, varDirection string) (*mainStruct.GatewayVariable, error) {
	if gateway == nil {
		return nil, errors.New("gateway name doesn't exists")
	}

	if gateway.Params.HasName(varName) {
		return nil, errors.New("variable name already exists")
	}

	res, err := db.SetUserGatewayVar(gateway.Id, varName, varValue, varDirection)
	if err != nil {
		return nil, err
	}
	variable := &mainStruct.GatewayVariable{Id: res, Name: varName, Value: varValue, Direction: varDirection, Enabled: true, Gateway: gateway}
	gateway.Vars.Set(variable)
	directory.GatewayVars.Set(variable)
	return variable, err
}

// GET
func GetDomain(id int64) *mainStruct.Domain {
	return directory.Domains.GetById(id)
}

func IsDirectoryEnabled() bool {
	return directory.Domains != nil
}

func GetDomainByName(name string) *mainStruct.Domain {
	if directory.Domains == nil {
		return nil
	}
	return directory.Domains.GetByName(name)
}

func GetDomains() map[int64]*mainStruct.Domain {
	return directory.Domains.GetCopy()
}

func GetUsers() map[int64]string {
	return directory.Users.GetList()
}

func GetDomainsUsers() map[int64]map[int64]*mainStruct.User {
	return directory.Users.GetDomainsList()
}

func GetDomainUsers(id int64) map[int64]string {
	domain := directory.Domains.GetById(id)
	if domain == nil {
		return map[int64]string{}
	}
	return domain.Users.GetList()
}

func GetUser(id int64) *mainStruct.User {
	return directory.Users.GetById(id)
}

func GetUserParameter(id int64) *mainStruct.UserParam {
	param, _ := directory.UserParams.GetById(id)
	return param
}

func GetUserVariable(id int64) *mainStruct.UserVariable {
	variable, _ := directory.UserVars.GetById(id)
	return variable
}

func GetUsersGateways() map[int64]*mainStruct.UserGateway {
	gateways := directory.UserGateways.GetList()
	return gateways
}

func GetUsersGatewaysParam(id int64) *mainStruct.GatewayParam {
	gateways, _ := directory.GatewayParams.GetById(id)
	return gateways
}

func GetUsersGatewaysVariable(id int64) *mainStruct.GatewayVariable {
	gateways, _ := directory.GatewayVars.GetById(id)
	return gateways
}

func GetDomainsGroups() map[int64]map[int64]string {
	return directory.Groups.GetDomainsList()
}

func GetGroup(id int64) *mainStruct.Group {
	group, _ := directory.Groups.GetById(id)
	return group
}

func GetGroupUser(id int64) *mainStruct.GroupUser {
	group, _ := directory.GroupUsers.GetById(id)
	return group
}

func GetUserGateway(id int64) *mainStruct.UserGateway {
	gateway, _ := directory.UserGateways.GetById(id)
	return gateway
}

func GetDomainParameter(id int64) *mainStruct.DomainParam {
	parameter, _ := directory.DomainParams.GetById(id)
	return parameter
}

func GetDomainVariable(id int64) *mainStruct.DomainVariable {
	variable, _ := directory.DomainVars.GetById(id)
	return variable
}

func GetDomainSipRegsCounter() map[string]int {
	list := directory.Domains.GetSipRegCounterList()
	return list
}

func DashBoardSetSofiaData() (*[]*mainStruct.SofiaProfile, *[]*mainStruct.SofiaGateway) {
	profiles, _ := GetSofiaProfilesProps()
	gateways := GetSofiasGatewaysProps()
	return &profiles, &gateways
}

// DELETE
func DelDirectoryDomain(domainId int64) bool {
	domain := directory.Domains.GetById(domainId)
	if domain == nil {
		return false
	}
	ok := db.DelDomain(domain.Id)
	if !ok {
		return false
	}

	directory.Domains.Remove(*domain)
	directory.ClearUp()
	return true
}

func DelDomainParameter(paramId int64) bool {
	param, ok := directory.DomainParams.GetById(paramId)
	if !ok {
		return false
	}

	ok = db.DelDomainParameter(paramId)
	if !ok {
		return false
	}

	param.Domain.Params.Remove(*param)
	directory.DomainParams.Remove(*param)
	return true
}

func DelDomainVariable(varId int64) bool {
	variable, ok := directory.DomainVars.GetById(varId)
	if !ok {
		return false
	}

	ok = db.DelDomainVariable(varId)
	if !ok {
		return false
	}

	variable.Domain.Vars.Remove(*variable)
	directory.DomainVars.Remove(*variable)
	return true
}

func DelGroup(griupId int64) bool {
	group, ok := directory.Groups.GetById(griupId)
	if !ok {
		return false
	}

	ok = db.DelGroup(griupId)
	if !ok {
		return false
	}

	group.Domain.Groups.Remove(*group)
	directory.Groups.Remove(*group)
	directory.ClearDirectoryGroups()
	return true
}

func DelGroupUser(userId int64) bool {
	groupUser, ok := directory.GroupUsers.GetById(userId)
	if !ok {
		return false
	}

	ok = db.DelGroupMember(userId)
	if !ok {
		return false
	}

	groupUser.Group.Users.Remove(*groupUser)
	directory.GroupUsers.Remove(*groupUser)
	return true
}

func DelUser(userId int64) bool {
	user := directory.Users.GetById(userId)
	if user == nil {
		return false
	}

	ok := db.DelUser(userId)
	if !ok {
		return false
	}

	user.Domain.Users.Remove(*user)
	directory.Users.Remove(*user)
	directory.ClearDirectoryUsers()
	return true
}

func DelUserParameter(parametrId int64) bool {
	userParam, ok := directory.UserParams.GetById(parametrId)
	if !ok {
		return false
	}

	ok = db.DelUserParameter(parametrId)
	if !ok {
		return false
	}

	userParam.User.Params.Remove(*userParam)
	directory.UserParams.Remove(*userParam)
	return true
}

func DelUserVariable(varId int64) bool {
	variable, ok := directory.UserVars.GetById(varId)
	if !ok {
		return false
	}

	ok = db.DelUserVariable(varId)
	if !ok {
		return false
	}

	variable.User.Vars.Remove(*variable)
	directory.UserVars.Remove(*variable)
	return true
}

func DelUserGateway(gatewayId int64) bool {
	gateway, ok := directory.UserGateways.GetById(gatewayId)
	if !ok {
		return false
	}

	ok = db.DelUserGateway(gatewayId)
	if !ok {
		return false
	}

	gateway.User.Gateways.Remove(*gateway)
	directory.UserGateways.Remove(*gateway)
	directory.ClearDirectoryUserGateways()
	return true
}

func DelUserGatewayParam(dparametrId int64) bool {
	param, ok := directory.GatewayParams.GetById(dparametrId)
	if !ok {
		return false
	}

	ok = db.DelUserGatewayParam(param.Id)
	if !ok {
		return false
	}

	param.Gateway.Params.Remove(*param)
	directory.GatewayParams.Remove(*param)
	return true
}

// UPDATE
func UpdateDomain(domainId int64, newName string) error {
	domain := directory.Domains.GetById(domainId)
	if domain == nil {
		return errors.New("domain name doesn't exists")
	}
	err := db.UpdateDomain(domainId, newName)
	if err != nil {
		return err
	}
	directory.Domains.Rename(domain.Name, newName)
	return err
}

func UpdateDomainParameterValue(paramId int64, newValue string) error {
	parameter, ok := directory.DomainParams.GetById(paramId)
	if !ok {
		return errors.New("parameter name doesn't exists")
	}
	err := db.UpdateDomainParameterValue(parameter.Id, newValue)
	if err != nil {
		return err
	}
	parameter.Value = newValue
	return err
}

func UpdateDomainParameter(paramId int64, newName string, newValue string) (int64, error) {
	parameter, ok := directory.DomainParams.GetById(paramId)
	if !ok {
		return 0, errors.New("parameter name doesn't exists")
	}
	res, err := db.UpdateDomainParameter(parameter.Id, newName, newValue)
	if err != nil {
		return 0, err
	}
	parameter.Name = newName
	parameter.Value = newValue
	return res, err
}

func UpdateDomainVariableValue(varId int64, newValue string) error {
	variable, ok := directory.DomainVars.GetById(varId)
	if !ok {
		return errors.New("variable name doesn't exists")
	}
	err := db.UpdateDomainVariableValue(variable.Id, newValue)
	if err != nil {
		return err
	}
	variable.Value = newValue
	return err
}

func UpdateDomainVariable(varId int64, newName string, newValue string) (int64, error) {
	variable, ok := directory.DomainVars.GetById(varId)
	if !ok {
		return 0, errors.New("variable name doesn't exists")
	}
	res, err := db.UpdateDomainVariable(variable.Id, newName, newValue)
	if err != nil {
		return 0, err
	}
	variable.Name = newName
	variable.Value = newValue
	return res, err
}

func UpdateDomainUser(userId int64, newUserName string) error {
	user := directory.Users.GetById(userId)
	if user == nil {
		return errors.New("user name already exists")
	}
	err := db.UpdateDomainUser(user.Id, newUserName)
	if err != nil {
		return err
	}
	user.Domain.Users.Rename(user.Name, newUserName)
	return err
}

func UpdateDomainGroup(userId int64, newGroupName string) error {
	group, ok := directory.Groups.GetById(userId)
	if !ok {
		return errors.New("group name already exists")
	}
	err := db.UpdateGroupName(group.Id, newGroupName)
	if err != nil {
		return err
	}
	group.Domain.Groups.Rename(group.Name, newGroupName)
	return err
}

func UpdateUserParameterValue(paramId int64, newValue string) error {
	parameter, ok := directory.UserParams.GetById(paramId)
	if !ok {
		return errors.New("parameter name doesn't exists")
	}
	err := db.UpdateUserParameterValue(parameter.Id, newValue)
	if err != nil {
		return err
	}
	parameter.Value = newValue
	return err
}

func UpdateUserVariableValue(varId int64, newValue string) error {
	variable, ok := directory.UserVars.GetById(varId)
	if !ok {
		return errors.New("variable name doesn't exists")
	}
	err := db.UpdateUserVariableValue(variable.Id, newValue)
	if err != nil {
		return err
	}
	variable.Value = newValue
	return err
}

func UpdateUserParameter(paramId int64, newName string, newValue string) (int64, error) {
	parameter, ok := directory.UserParams.GetById(paramId)
	if !ok {
		return 0, errors.New("parameter name doesn't exists")
	}
	res, err := db.UpdateUserParameter(parameter.Id, newName, newValue)
	if err != nil {
		return 0, err
	}
	parameter.Name = newName
	parameter.Value = newValue
	return res, err
}

func UpdateUserVariable(varId int64, newName string, newValue string) (int64, error) {
	variable, ok := directory.UserVars.GetById(varId)
	if !ok {
		return 0, errors.New("variable name doesn't exists")
	}
	res, err := db.UpdateUserVariable(variable.Id, newName, newValue)
	if err != nil {
		return 0, err
	}
	variable.Name = newName
	variable.Value = newValue
	return res, err
}

func UpdateUserCache(userId int64, newValue uint) (int64, error) {
	user := directory.Users.GetById(userId)
	if user == nil {
		return 0, errors.New("user name doesn't exists")
	}
	res, err := db.UpdateUserCache(user.Id, newValue)
	if err != nil {
		return 0, err
	}
	user.Cache = newValue
	return res, err
}

func UpdateUserCidr(userId int64, newValue string) (int64, error) {
	user := directory.Users.GetById(userId)
	if user == nil {
		return 0, errors.New("user name doesn't exists")
	}
	res, err := db.UpdateUserCidr(user.Id, newValue)
	if err != nil {
		return 0, err
	}
	user.Cidr = newValue
	return res, err
}

func UpdateUserNumberAlias(userId int64, newValue string) (int64, error) {
	user := directory.Users.GetById(userId)
	if user == nil {
		return 0, errors.New("user name doesn't exists")
	}
	res, err := db.UpdateUserCidr(user.Id, newValue)
	if err != nil {
		return 0, err
	}
	user.NumberAlias = newValue
	return res, err
}

func UpdateUserGatewayParameter(paramId int64, newName string, newValue string) (int64, error) {
	parameter, ok := directory.GatewayParams.GetById(paramId)
	if !ok {
		return 0, errors.New("parameter name doesn't exists")
	}
	res, err := db.UpdateUserGatewayParameter(parameter.Id, newName, newValue)
	if err != nil {
		return 0, err
	}
	parameter.Name = newName
	parameter.Value = newValue
	return res, err
}

func UpdateUserGatewayVariable(variable *mainStruct.GatewayVariable, name, value, direction string) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	_, err := db.UpdateUserGatewayVariable(variable.Id, name, value, direction)
	if err != nil {
		return 0, err
	}
	variable.Name = name
	variable.Value = value
	variable.Direction = direction
	return variable.Gateway.Id, err
}

func SwitchUserGatewayVariable(variable *mainStruct.GatewayVariable, enabled bool) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	err := db.SwitchUserGatewayVariable(variable.Id, enabled)
	if err != nil {
		return 0, err
	}

	variable.Enabled = enabled
	return variable.Gateway.Id, err
}

func DelUserGatewayVariable(variable *mainStruct.GatewayVariable) int64 {
	if variable == nil {
		return 0
	}
	parentId := variable.Gateway.Id
	ok := db.DelUserGatewayVariable(variable.Id)
	if !ok {
		return 0
	}

	variable.Gateway.Vars.Remove(variable)
	directory.GatewayVars.Remove(variable)
	return parentId
}

func UpdateDomainUserGateway(userId int64, newUName string) error {
	gateway, ok := directory.UserGateways.GetById(userId)
	if !ok {
		return errors.New("gateway name already exists")
	}
	err := db.UpdateDomainUserGateway(gateway.Id, newUName)
	if err != nil {
		return err
	}
	gateway.User.Gateways.Rename(gateway.Name, newUName)
	return err
}

/*
func UpdateGroup(domainName, groupName, newName string) (error) {
	domain, ok := domains.GetByName(domainName)
	if !ok {
		return errors.New("domain name doesn't exists")
	}
	err := db.UpdateDomain(domain.Id, newName)
	if err != nil {
		return err
	}
	domain.Name = newName
	domains.Set(domain)
	domains.Remove(domainName)
	return err
}*/

func SwitchDomain(domain *mainStruct.Domain, enabled bool) (int64, error) {
	if domain == nil {
		return 0, errors.New("domain doesn't exists")
	}
	err := db.SwitchDomain(domain.Id, enabled)
	if err != nil {
		return 0, err
	}

	domain.Enabled = enabled
	return domain.Id, err
}

func SwitchDomainParameter(parameter *mainStruct.DomainParam, enabled bool) (int64, error) {
	if parameter == nil {
		return 0, errors.New("parameter doesn't exists")
	}
	err := db.SwitchDomainParameter(parameter.Id, enabled)
	if err != nil {
		return 0, err
	}

	parameter.Enabled = enabled
	return parameter.Domain.Id, err
}

func SwitchDomainVariable(variable *mainStruct.DomainVariable, enabled bool) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	err := db.SwitchDomainVariable(variable.Id, enabled)
	if err != nil {
		return 0, err
	}

	variable.Enabled = enabled
	return variable.Domain.Id, err
}

func SwitchUser(user *mainStruct.User, enabled bool) (int64, error) {
	if user == nil {
		return 0, errors.New("user doesn't exists")
	}
	err := db.SwitchUsers(user.Id, enabled)
	if err != nil {
		return 0, err
	}

	user.Enabled = enabled
	return user.Domain.Id, err
}

func SwitchUserParameter(parameter *mainStruct.UserParam, enabled bool) (int64, error) {
	if parameter == nil {
		return 0, errors.New("parameter doesn't exists")
	}
	err := db.SwitchUserParameter(parameter.Id, enabled)
	if err != nil {
		return 0, err
	}

	parameter.Enabled = enabled
	return parameter.User.Id, err
}

func SwitchUserVariable(variable *mainStruct.UserVariable, enabled bool) (int64, error) {
	if variable == nil {
		return 0, errors.New("variable doesn't exists")
	}
	err := db.SwitchUserVariable(variable.Id, enabled)
	if err != nil {
		return 0, err
	}

	variable.Enabled = enabled
	return variable.User.Id, err
}

func SwitchUserGateway(gateway *mainStruct.UserGateway, enabled bool) (int64, error) {
	if gateway == nil {
		return 0, errors.New("gateway doesn't exists")
	}
	err := db.SwitchUserGateway(gateway.Id, enabled)
	if err != nil {
		return 0, err
	}

	gateway.Enabled = enabled
	return gateway.User.Id, err
}

func SwitchUserGatewayParameter(parameter *mainStruct.GatewayParam, enabled bool) (int64, error) {
	if parameter == nil {
		return 0, errors.New("parameter doesn't exists")
	}
	err := db.SwitchUserGatewayParameter(parameter.Id, enabled)
	if err != nil {
		return 0, err
	}

	parameter.Enabled = enabled
	return parameter.Gateway.Id, err
}

func SwitchGroup(group *mainStruct.Group, enabled bool) (int64, error) {
	if group == nil {
		return 0, errors.New("group doesn't exists")
	}
	err := db.SwitchGroup(group.Id, enabled)
	if err != nil {
		return 0, err
	}

	group.Enabled = enabled
	return group.Domain.Id, err
}

func IsCallcenterEnabled() bool {
	return configs.Callcenter != nil
}
