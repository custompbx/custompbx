package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webStruct"
	"github.com/custompbx/customorm"
	"regexp"
	"strconv"
	"strings"
)

func getDirectoryByParent(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true, UseValue: true, Value: data.Id}}

	res, err := intermediateDB.GetByValuesAsMap(
		item,
		filter,
	)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: res, MessageType: data.Event}
}

func getDirectoryByParents(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	if len(data.IntSlice) == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	filter := map[string]customorm.FilterFields{"Parent": {Flag: true, UseValue: true, Value: data.IntSlice, Operand: "IN"}}

	res, err := intermediateDB.GetByValuesAsMap(
		item,
		filter,
	)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: res, MessageType: data.Event}
}

func getDirectoryById(data *webStruct.MessageData, item interface{}) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	res, err := intermediateDB.GetByIdArg(item, data.Id)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Data: res, MessageType: data.Event}
}

func importDirectory(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	fsesl.GetXMLDirectory()
	return webStruct.UserResponse{MessageType: data.Event}
}

func ImportXMLDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.File == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	err := fsesl.ParseDirectoryXML(data.File)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	//items := pbxcache.GetDomains()

	return webStruct.UserResponse{MessageType: data.Event}
}

func getDomains(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetDomains()
	return webStruct.UserResponse{Domains: &items, MessageType: data.Event}
}

func setDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	domain, err := pbxcache.SetDomain(data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.Domain)
	items[domain.Id] = domain

	return webStruct.UserResponse{Domains: &items, MessageType: data.Event}
}

func renameDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetDomain(data.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateDomain(data.Id, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	items := make(map[int64]*mainStruct.Domain)
	items[data.Id] = domain

	return webStruct.UserResponse{Domains: &items, MessageType: data.Event}
}

func delDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}
	ok := pbxcache.DelDirectoryDomain(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete domain", MessageType: data.Event}
	}

	return webStruct.UserResponse{Id: &data.Id, MessageType: data.Event}
}

func getDomainsDetails(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	domain := pbxcache.GetDomain(data.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	directory[data.Id] = webStruct.DomainDetails{Vars: domain.Vars.GetList(), Params: domain.Params.GetList()}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func setDomainsParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetDomain(data.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	parameter, err := pbxcache.SetDomainParameter(domain, data.Name, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	parameters := map[int64]*mainStruct.DomainParam{parameter.Id: parameter}
	directory[domain.Id] = webStruct.DomainDetails{Params: parameters}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func setDomainsVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetDomain(data.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	variable, err := pbxcache.SetDomainVariable(domain, data.Name, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	variables := map[int64]*mainStruct.DomainVariable{variable.Id: variable}
	directory[domain.Id] = webStruct.DomainDetails{Vars: variables}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func delDomainsParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	ok := pbxcache.DelDomainParameter(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func delDomainsVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	ok := pbxcache.DelDomainVariable(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func updateDomainsParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetDomainParameter(data.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDomainParameter(data.Id, data.Name, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	parameters := map[int64]*mainStruct.DomainParam{param.Id: param}
	directory[param.Domain.Id] = webStruct.DomainDetails{Params: parameters}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func updateDomainsVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetDomainVariable(data.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateDomainVariable(data.Id, data.Name, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	variables := map[int64]*mainStruct.DomainVariable{variable.Id: variable}
	directory[variable.Domain.Id] = webStruct.DomainDetails{Vars: variables}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func getUsers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetDomains()
	list := pbxcache.GetDomainsUsers()
	return webStruct.UserResponse{Domains: &items, DirectoryUsers: &list, MessageType: data.Event}
}

func getUsersDetails(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.UserDetails)
	directory[data.Id] = webStruct.UserDetails{Vars: directoryUser.Vars.GetList(), Params: directoryUser.Params.GetList(), Cache: directoryUser.Cache, Cidr: directoryUser.Cidr, NumberAlias: directoryUser.NumberAlias}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

func setUsersParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	parameter, err := pbxcache.SetUserParameter(directoryUser, data.Name, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	directory := map[int64]webStruct.UserDetails{parameter.User.Id: {Params: map[int64]*mainStruct.UserParam{parameter.Id: parameter}}}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

func setUsersVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	variable, err := pbxcache.SetUserVariable(directoryUser, data.Name, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	directory := map[int64]webStruct.UserDetails{variable.User.Id: {Vars: map[int64]*mainStruct.UserVariable{variable.Id: variable}}}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

func delUsersParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	ok := pbxcache.DelUserParameter(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func delUsersVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	ok := pbxcache.DelUserVariable(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func updateUsersParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	parameter := pbxcache.GetUserParameter(data.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserParameter(data.Id, data.Name, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	directory := map[int64]webStruct.UserDetails{parameter.User.Id: {Params: map[int64]*mainStruct.UserParam{parameter.Id: parameter}}}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

func updateUsersVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetUserVariable(data.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserVariable(data.Id, data.Name, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	directory := map[int64]webStruct.UserDetails{variable.User.Id: {Vars: map[int64]*mainStruct.UserVariable{variable.Id: variable}}}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

func updateUsersCache(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	cacheValue, err := strconv.ParseUint(data.Value, 10, 32)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserCache(data.Id, uint(cacheValue))
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Value: data.Value}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item}
}

func updateUsersCidr(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserCidr(data.Id, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Value: data.Value}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item}
}

func UpdateUserNumberAlias(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserNumberAlias(data.Id, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Value: data.Value}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item}
}

func addNewUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	domainI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomain{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}
	domain, ok := domainI.(altStruct.DirectoryDomain)
	if !ok {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}
	if data.Bulk > 100 {
		return webStruct.UserResponse{Error: "to many at once", MessageType: data.Event}
	}
	if data.Bulk > 1 {
		counter := 0
		items := map[int64]interface{}{}
		re := regexp.MustCompile("[0-9]+")
		starts := re.FindAllString(data.Name, -1)
		i := 0
		if len(starts) != 0 {
			i, err = strconv.Atoi(starts[len(starts)-1])
		}
		userNamePrefix := ""
		userNameSuffix := ""
		if i != 0 {
			nameParts := strings.Split(data.Name, strconv.Itoa(i))
			for key, part := range nameParts {
				if key == len(nameParts)-1 {
					userNameSuffix = part
					break
				}
				if key != len(nameParts)-2 {
					part += strconv.Itoa(i)
				}
				userNamePrefix += part
			}
		}
		top := data.Bulk + i
		for ; i < top; i++ {
			userName := userNamePrefix + strconv.Itoa(i) + userNameSuffix
			res, err := altData.SetDirectoryDomainUser(domain.Id, userName, "", "", "")
			if err != nil {
				continue
			}

			userI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: res})
			if err != nil {
				continue
			}
			dUser, ok := userI.(altStruct.DirectoryDomainUser)
			if !ok {
				continue
			}
			counter++
			items[dUser.Id] = dUser
		}
		return webStruct.UserResponse{Data: &items, MessageType: data.Event, Total: &counter}
	}

	res, err := altData.SetDirectoryDomainUser(domain.Id, data.Name, "", "", "")
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	userI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: res})
	if err != nil {
		return webStruct.UserResponse{Error: "user not added", MessageType: data.Event}
	}
	dUser, ok := userI.(altStruct.DirectoryDomainUser)
	if !ok {
		return webStruct.UserResponse{Error: "user not added", MessageType: data.Event}
	}

	item := map[int64]interface{}{dUser.Id: dUser}

	return webStruct.UserResponse{Data: &item, MessageType: data.Event}
}

func ImportXMLDomainUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.File == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domainI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomain{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}
	domain, ok := domainI.(altStruct.DirectoryDomain)
	if !ok {
		return webStruct.UserResponse{Error: "profile not found", MessageType: data.Event}
	}
	err = fsesl.ParseDirectoryUsersXML(domain.Id, data.File)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event}
}

func delUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	domainId := directoryUser.Domain.Id

	ok := pbxcache.DelUser(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete user", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &domainId, Item: &item, MessageType: data.Event}
}

func renameUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateDomainUser(data.Id, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &directoryUser.Domain.Id, Item: &item, MessageType: data.Event}
}

func getGroups(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	items := pbxcache.GetDomains()
	list := pbxcache.GetDomainsGroups()
	return webStruct.UserResponse{Domains: &items, List: &list, MessageType: data.Event}
}

func getGroupUsers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	group := pbxcache.GetGroup(data.Id)
	if group == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	users := group.Users.GetList()
	groupUsers := make(map[int64]map[int64]*mainStruct.GroupUser)
	groupUsers[data.Id] = users

	item := pbxcache.GetDomainUsers(group.Domain.Id)

	list := map[int64]map[int64]string{group.Domain.Id: item}

	return webStruct.UserResponse{GroupUsers: &groupUsers, List: &list, MessageType: data.Event}
}

func addNewGroup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	domain := pbxcache.GetDomain(data.Id)
	if domain == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetDomainGroup(domain, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := map[int64]string{res.Id: data.Name}

	return webStruct.UserResponse{Id: &data.Id, Items: &item, MessageType: data.Event}
}

func delGroup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	directoryGroup := pbxcache.GetGroup(data.Id)
	if directoryGroup == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	domainId := directoryGroup.Domain.Id

	ok := pbxcache.DelGroup(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete group", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &domainId, Item: &item, MessageType: data.Event}
}

func renameGroup(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryGroup := pbxcache.GetGroup(data.Id)
	if directoryGroup == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateDomainGroup(data.Id, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &directoryGroup.Domain.Id, Item: &item, MessageType: data.Event}
}

func addNewGroupUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong group id", MessageType: data.Event}
	}
	userId, err := strconv.ParseInt(data.Value, 10, 64)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong user id", MessageType: data.Event}
	}

	group := pbxcache.GetGroup(data.Id)
	if group == nil {
		return webStruct.UserResponse{Error: "group not found", MessageType: data.Event}
	}

	userObj := pbxcache.GetUser(userId)
	if userObj == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	item, err := pbxcache.SetDomainGroupNewUser(group, userObj)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	groupUsers := make(map[int64]map[int64]*mainStruct.GroupUser)
	groupUsers[group.Id] = map[int64]*mainStruct.GroupUser{item.Id: item}

	return webStruct.UserResponse{GroupUsers: &groupUsers, Id: &group.Id, MessageType: data.Event}
}

func delGroupUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	userObj := pbxcache.GetGroupUser(data.Id)
	if userObj == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	res := pbxcache.DelGroupUser(data.Id)
	if !res {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	return webStruct.UserResponse{AffectedId: &userObj.Id, Id: &userObj.Group.Id, MessageType: data.Event}
}

func getUsersGateways(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	gateways := make(map[int64]map[int64]map[int64]*mainStruct.UserGateway)
	for key, value := range pbxcache.GetUsersGateways() {
		_, ok := gateways[value.User.Domain.Id]
		if !ok {
			gateways[value.User.Domain.Id] = make(map[int64]map[int64]*mainStruct.UserGateway)
		}
		_, ok2 := gateways[value.User.Domain.Id][value.User.Id]
		if !ok2 {
			gateways[value.User.Domain.Id][value.User.Id] = make(map[int64]*mainStruct.UserGateway)
		}
		gateways[value.User.Domain.Id][value.User.Id][key] = value
	}

	items := pbxcache.GetDomains()
	list := pbxcache.GetDomainsUsers()
	return webStruct.UserResponse{Domains: &items, DirectoryUsers: &list, UserGateways: &gateways, MessageType: data.Event}
}

func getUsersGatewaysDetails(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	directoryGateway := pbxcache.GetUserGateway(data.Id)
	if directoryGateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	item := make(map[int64]*webStruct.GatewayDetails)
	item[directoryGateway.Id] = &webStruct.GatewayDetails{Params: directoryGateway.Params.GetList(), Vars: directoryGateway.Vars.GetList()}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func setUsersGatewaysParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryGateway := pbxcache.GetUserGateway(data.Id)
	if directoryGateway == nil {
		return webStruct.UserResponse{Error: "user gateway not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetUserGatewayParam(directoryGateway, data.Name, data.Value)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	item := webStruct.Item{Id: res.Id, Name: data.Name, Value: data.Value, Enabled: true}

	return webStruct.UserResponse{MessageType: data.Event, Item: &item}
}

func setUsersGatewaysVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Variable.Id != 0 || data.Variable.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	gateway := pbxcache.GetUserGateway(data.Id)
	if gateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	variable, err := pbxcache.SetUserGatewayVar(gateway, data.Variable.Name, data.Variable.Value, data.Variable.Direction)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.GatewayDetails)
	item[gateway.Id] = &webStruct.GatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func updateUsersGatewaysVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 || data.Variable.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetUsersGatewaysVariable(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserGatewayVariable(variable, data.Variable.Name, data.Variable.Value, data.Variable.Direction)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.GatewayDetails)
	item[variable.Gateway.Id] = &webStruct.GatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func switchUserGatewayVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetUsersGatewaysVariable(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchUserGatewayVariable(variable, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.GatewayDetails)
	item[variable.Gateway.Id] = &webStruct.GatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func delConfigUserGatewayVar(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	variable := pbxcache.GetUsersGatewaysVariable(data.Variable.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res := pbxcache.DelUserGatewayVariable(variable)
	if res == 0 {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayVariable{variable.Id: variable}
	item := make(map[int64]*webStruct.GatewayDetails)
	item[variable.Gateway.Id] = &webStruct.GatewayDetails{Vars: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func delUsersGatewaysParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	ok := pbxcache.DelUserGatewayParam(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func updateUsersGatewaysParam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" || data.Value == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	param := pbxcache.GetUsersGatewaysParam(data.Id)
	if param == nil {
		return webStruct.UserResponse{Error: "param not found", MessageType: data.Event}
	}

	res, err := pbxcache.UpdateUserGatewayParameter(data.Id, data.Name, data.Value)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't update", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayParam{param.Id: param}
	item := make(map[int64]*webStruct.GatewayDetails)
	item[param.Gateway.Id] = &webStruct.GatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}

}

func addNewUsersGateway(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	res, err := pbxcache.SetUserGateway(directoryUser, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := webStruct.Item{Id: res.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &directoryUser.Domain.Id, AffectedId: &data.Id, Item: &item, MessageType: data.Event}
}

func delUsersGateway(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	directoryGateway := pbxcache.GetUserGateway(data.Id)
	if directoryGateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	ok := pbxcache.DelUserGateway(data.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete gateway", MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &directoryGateway.User.Domain.Id, AffectedId: &directoryGateway.User.Id, Item: &item, MessageType: data.Event}
}

func renameUsersGateway(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	directoryGateway := pbxcache.GetUserGateway(data.Id)
	if directoryGateway == nil {
		return webStruct.UserResponse{Error: "gateway not found", MessageType: data.Event}
	}

	err := pbxcache.UpdateDomainUserGateway(data.Id, data.Name)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	item := webStruct.Item{Id: data.Id, Name: data.Name}

	return webStruct.UserResponse{Id: &directoryGateway.User.Domain.Id, AffectedId: &directoryGateway.User.Id, Item: &item, MessageType: data.Event}
}

func switchUserGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	parameter := pbxcache.GetUsersGatewaysParam(data.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}
	if data.Enabled == nil {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}
	res, err := pbxcache.SwitchUserGatewayParameter(parameter, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayParam{parameter.Id: parameter}
	item := make(map[int64]*webStruct.GatewayDetails)
	item[parameter.Gateway.Id] = &webStruct.GatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func switchDomain(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	domain := pbxcache.GetDomain(data.Id)
	if domain == nil || data.Enabled == nil {
		return webStruct.UserResponse{Error: "domain not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDomain(domain, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := make(map[int64]*mainStruct.Domain)
	items[domain.Id] = domain

	return webStruct.UserResponse{Domains: &items, MessageType: data.Event}
}

func switchDomainParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	parameter := pbxcache.GetDomainParameter(data.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDomainParameter(parameter, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	parameters := map[int64]*mainStruct.DomainParam{parameter.Id: parameter}
	directory[parameter.Domain.Id] = webStruct.DomainDetails{Params: parameters}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func switchDomainVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	variable := pbxcache.GetDomainVariable(data.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "variable not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchDomainVariable(variable, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	directory := make(map[int64]webStruct.DomainDetails)
	variables := map[int64]*mainStruct.DomainVariable{variable.Id: variable}
	directory[variable.Domain.Id] = webStruct.DomainDetails{Vars: variables}

	return webStruct.UserResponse{MessageType: data.Event, DomainDetails: &directory}
}

func switchUser(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {

	if data.Id == 0 {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	directoryUser := pbxcache.GetUser(data.Id)
	if directoryUser == nil || data.Enabled == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchUser(directoryUser, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]map[int64]*mainStruct.User{directoryUser.Domain.Id: {directoryUser.Id: directoryUser}}

	return webStruct.UserResponse{DirectoryUsers: &items, MessageType: data.Event}
}

func switchUserParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	parameter := pbxcache.GetUserParameter(data.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchUserParameter(parameter, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	directory := map[int64]webStruct.UserDetails{parameter.User.Id: {Params: map[int64]*mainStruct.UserParam{parameter.Id: parameter}}}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

func switchUserVariable(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	variable := pbxcache.GetUserVariable(data.Id)
	if variable == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcache.SwitchUserVariable(variable, *data.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	directory := map[int64]webStruct.UserDetails{variable.User.Id: {Vars: map[int64]*mainStruct.UserVariable{variable.Id: variable}}}

	return webStruct.UserResponse{MessageType: data.Event, UserDetails: &directory}
}

/*
func switchConfigUserGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	parameter := pbxcash.GetUsersGatewaysParam(data.Variable.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcash.SwitchUserGatewayParameter(parameter, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayParam{parameter.Id: parameter}
	item := make(map[int64]*GatewayDetails)
	item[parameter.Gateway.Id] = &GatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func switchConfigUserGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	parameter := pbxcash.GetUsersGatewaysParam(data.Variable.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcash.SwitchUserGatewayParameter(parameter, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayParam{parameter.Id: parameter}
	item := make(map[int64]*GatewayDetails)
	item[parameter.Gateway.Id] = &GatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func switchConfigUserGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	parameter := pbxcash.GetUsersGatewaysParam(data.Variable.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcash.SwitchUserGatewayParameter(parameter, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayParam{parameter.Id: parameter}
	item := make(map[int64]*GatewayDetails)
	item[parameter.Gateway.Id] = &GatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}

func switchConfigUserGatewayParameter(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Variable.Id == 0 {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	parameter := pbxcash.GetUsersGatewaysParam(data.Variable.Id)
	if parameter == nil {
		return webStruct.UserResponse{Error: "parameter not found", MessageType: data.Event}
	}

	res, err := pbxcash.SwitchUserGatewayParameter(parameter, data.Variable.Enabled)
	if err != nil || res == 0 {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	subItem := map[int64]*mainStruct.GatewayParam{parameter.Id: parameter}
	item := make(map[int64]*GatewayDetails)
	item[parameter.Gateway.Id] = &GatewayDetails{Params: subItem}

	return webStruct.UserResponse{MessageType: data.Event, GatewayDetails: &item}
}
*/
