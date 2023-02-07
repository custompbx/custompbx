package altData

import (
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"errors"
	"github.com/custompbx/customorm"
	"strconv"
)

func XMLDomainDirectoryBase(domainName, userName, groupName string, hidePass, onlyGateways, onlyCidr bool) []interface{} {
	filter := map[string]customorm.FilterFields{
		"Parent":  {Flag: true},
		"Enabled": {Flag: true},
		"Name":    {Flag: domainName != "", UseValue: true, Value: domainName, Operand: customorm.OperandEqual},
	}
	domains, _ := intermediateDB.GetByValues(
		&altStruct.DirectoryDomain{Parent: &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()}, Enabled: true},
		filter,
	)
	var listLists []interface{}
	for _, domainI := range domains {
		domain, ok := domainI.(altStruct.DirectoryDomain)
		if !ok {
			continue
		}
		var domainParams []interface{}
		var domainVariables []interface{}
		var groupLists []interface{}
		var userLists []interface{}

		if userName != "" {
			domainParams, _ = intermediateDB.GetByValue(
				&altStruct.DirectoryDomainParameter{Parent: &altStruct.DirectoryDomain{Id: domain.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
			domainVariables, _ = intermediateDB.GetByValue(
				&altStruct.DirectoryDomainVariable{Parent: &altStruct.DirectoryDomain{Id: domain.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
		}

		users := getDirectoryUsers(domain.Id, userName, hidePass, onlyGateways, onlyCidr)
		if userName != "" && len(users) == 0 {
			// if no user found by domain/name return nil to send "not found" then
			return nil
		}
		if groupName == "" {
			userLists = append(userLists, users...)
			groupLists = append(
				groupLists,
				struct {
					Name     string      `xml:"name,attr"`
					XMLUsers interface{} `xml:"users>user"`
				}{
					Name:     "CustomPbxGroupForFS1.6",
					XMLUsers: userLists,
				},
			)
		} else {
			groupLists = append(groupLists, getDirectoryGroups(domain.Id, groupName, users)...)
		}
		listLists = append(listLists, struct {
			*altStruct.DirectoryDomain
			Params interface{} `xml:"params>param"`
			Vars   interface{} `xml:"variables>variable"`
			Groups interface{} `xml:"groups>group"`
		}{
			&domain,
			domainParams,
			domainVariables,
			groupLists,
		})
	}

	return listLists
}

func XMLDomainDirectoryGateways() []interface{} {
	return XMLDomainDirectoryBase("", "", "", true, true, false)
}

func XMLDomainDirectoryNetworkLists(domainName string) []interface{} {
	return XMLDomainDirectoryBase(domainName, "", "", true, false, true)
}

func XMLDomainDirectoryDefault() []interface{} {
	return XMLDomainDirectoryBase("", "", "", true, false, false)
}

func XMLDomainDirectoryUser(domainName, userName string) []interface{} {
	if domainName == "" || userName == "" {
		return nil
	}
	return XMLDomainDirectoryBase(domainName, userName, "", true, false, false)
}

func XMLDomainDirectoryUserGroup(domainName, groupName string) []interface{} {
	if domainName == "" || groupName == "" {
		return nil
	}
	return XMLDomainDirectoryBase(domainName, "", groupName, true, false, false)
}

func getDirectoryGroups(domainId int64, groupName string, users []interface{}) []interface{} {
	var groupLists []interface{}
	filter := map[string]customorm.FilterFields{
		"Parent":  {Flag: true},
		"Enabled": {Flag: true},
		"Name":    {Flag: groupName != "", UseValue: true, Value: groupName, Operand: customorm.OperandEqual},
	}
	domainGroups, _ := intermediateDB.GetByValues(
		&altStruct.DirectoryDomainGroup{Parent: &altStruct.DirectoryDomain{Id: domainId}, Enabled: true},
		filter,
	)
	userNames := make(map[int64]string, len(users))
	for _, userI := range users {
		user, ok := userI.(altStruct.DirectoryDomainUser)
		if !ok {
			continue
		}
		userNames[user.Id] = user.Name
	}
	for _, groupI := range domainGroups {
		group, ok := groupI.(altStruct.DirectoryDomainGroup)
		if !ok {
			continue
		}
		groupUsers, _ := intermediateDB.GetByValue(
			&altStruct.DirectoryDomainGroupUser{Parent: &group, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		for _, guserI := range groupUsers {
			user, ok := guserI.(altStruct.DirectoryDomainGroupUser)
			if !ok {
				continue
			}
			user.Name = userNames[user.UserId.Id]
		}

		groupLists = append(groupLists, struct {
			*altStruct.DirectoryDomainGroup
			Users interface{} `xml:"users>user"`
		}{
			&group,
			groupUsers,
		})
	}
	return groupLists
}

func getDirectoryUsers(domainId int64, userName string, hidePass, onlyGateways, onlyCidr bool) []interface{} {
	var userLists []interface{}
	filter := map[string]customorm.FilterFields{
		"Parent":  {Flag: true},
		"Enabled": {Flag: true},
		"Cidr":    {Flag: onlyCidr, UseValue: true, Value: "", Operand: customorm.OperandNotEqual},
		"Name":    {Flag: userName != "", UseValue: true, Value: userName, Operand: customorm.OperandEqual},
	}

	directoryUser := &altStruct.DirectoryDomainUser{Parent: &altStruct.DirectoryDomain{Id: domainId}, Enabled: true}
	domainUsers, _ := intermediateDB.GetByValues(
		directoryUser,
		filter,
	)
	for _, userI := range domainUsers {
		user, ok := userI.(altStruct.DirectoryDomainUser)
		if !ok {
			continue
		}
		var userParams []interface{}
		var userVars []interface{}
		var userGateways []interface{}
		if !onlyGateways && !onlyCidr {
			userParamsFilter := map[string]customorm.FilterFields{
				"Parent":  {Flag: true},
				"Enabled": {Flag: true},
			}
			if userName == "" {
				userParamsFilter["Name"] = customorm.FilterFields{Flag: hidePass, UseValue: true, Value: "password", Operand: customorm.OperandNotEqual}
			}
			userParams, _ = intermediateDB.GetByValues(
				&altStruct.DirectoryDomainUserParameter{Parent: &altStruct.DirectoryDomainUser{Id: user.Id}, Enabled: true},
				userParamsFilter,
			)
			userVars, _ = intermediateDB.GetByValue(
				&altStruct.DirectoryDomainUserVariable{Parent: &altStruct.DirectoryDomainUser{Id: user.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
		}
		if (!hidePass || onlyGateways) && !onlyCidr {
			userGateways, _ = intermediateDB.GetByValue(
				&altStruct.DirectoryDomainUserGateway{Parent: &altStruct.DirectoryDomainUser{Id: user.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
		}
		var userGatewayLists []interface{}
		for _, userGI := range userGateways {
			g, ok := userGI.(altStruct.DirectoryDomainUserGateway)
			if !ok {
				continue
			}
			gatewayParams, _ := intermediateDB.GetByValue(
				&altStruct.DirectoryDomainUserGatewayParameter{Parent: &altStruct.DirectoryDomainUserGateway{Id: g.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
			gatewayVars, _ := intermediateDB.GetByValue(
				&altStruct.DirectoryDomainUserGatewayVariable{Parent: &altStruct.DirectoryDomainUserGateway{Id: g.Id}, Enabled: true},
				map[string]bool{"Parent": true, "Enabled": true},
			)
			userGatewayLists = append(userGatewayLists, struct {
				*altStruct.DirectoryDomainUserGateway
				Params interface{} `xml:"param"`
				Vars   interface{} `xml:"variable"`
			}{
				&g,
				gatewayParams,
				gatewayVars,
			})
		}
		if onlyGateways && len(userGatewayLists) == 0 {
			continue
		}
		userLists = append(userLists, struct {
			*altStruct.DirectoryDomainUser
			Params   interface{} `xml:"params>param"`
			Vars     interface{} `xml:"variables>variable"`
			Gateways interface{} `xml:"gateways>gateway"`
		}{
			&user,
			userParams,
			userVars,
			userGatewayLists,
		})
	}
	return userLists
}

func SetDirectoryDomain(name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomain{
		Name:    name,
		Parent:  &mainStruct.FsInstance{Id: cache.GetCurrentInstanceId()},
		Enabled: true,
	})
}

func SetDirectoryDomainParameter(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.DirectoryDomain{Id: parentId},
		Enabled: true,
	})
}

func SetDirectoryDomainVariable(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainVariable{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.DirectoryDomain{Id: parentId},
		Enabled: true,
	})
}

func SetDirectoryDomainUser(parentId int64, name, cache, cidr, numberAlias string) (int64, error) {
	c, err := strconv.ParseUint(cache, 10, 32)
	if err != nil {
		c = 1000
	}
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainUser{
		Name:        name,
		Cache:       uint(c),
		Cidr:        cidr,
		NumberAlias: numberAlias,
		Parent:      &altStruct.DirectoryDomain{Id: parentId},
		Enabled:     true,
	})
}

func SetDirectoryDomainGroup(parentId int64, name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainGroup{
		Name:    name,
		Parent:  &altStruct.DirectoryDomain{Id: parentId},
		Enabled: true,
	})
}

func SetDirectoryDomainGroupUser(groupId, domainId int64, userName string) (int64, error) {
	domainUser, _ := intermediateDB.GetByValue(
		&altStruct.DirectoryDomainUser{Parent: &altStruct.DirectoryDomain{Id: domainId}, Name: userName},
		map[string]bool{"Parent": true, "Name": true},
	)
	if len(domainUser) == 0 {
		return 0, errors.New("user not found")
	}
	user, ok := domainUser[0].(altStruct.DirectoryDomainUser)
	if !ok {
		return 0, errors.New("user not found")
	}
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainGroupUser{
		UserId:  &altStruct.DirectoryDomainUser{Id: user.Id},
		Parent:  &altStruct.DirectoryDomainGroup{Id: groupId},
		Enabled: true,
	})
}

func SetDirectoryDomainUserParameter(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainUserParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.DirectoryDomainUser{Id: parentId},
		Enabled: true,
	})
}

func SetDirectoryDomainUserVariable(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainUserVariable{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.DirectoryDomainUser{Id: parentId},
		Enabled: true,
	})
}

func SetDirectoryDomainUserGateway(parentId int64, name string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainUserGateway{
		Name:    name,
		Parent:  &altStruct.DirectoryDomainUser{Id: parentId},
		Enabled: true,
	})
}

func SetDirectoryDomainUserGatewayVariable(parentId int64, name, value, direction string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainUserGatewayVariable{
		Name:        name,
		Value:       value,
		Description: direction,
		Parent:      &altStruct.DirectoryDomainUserGateway{Id: parentId},
		Enabled:     true,
	})
}

func SetDirectoryDomainUserGatewayParameter(parentId int64, name, value string) (int64, error) {
	return intermediateDB.InsertItem(&altStruct.DirectoryDomainUserGatewayParameter{
		Name:    name,
		Value:   value,
		Parent:  &altStruct.DirectoryDomainUserGateway{Id: parentId},
		Enabled: true,
	})
}

func GetDirectoryNameAndInstanceByStruct(structure interface{}, par interface{}) interface{} {
	switch structure.(type) {
	case altStruct.DirectoryDomain, *altStruct.DirectoryDomain:
		parent, ok := par.(*mainStruct.FsInstance)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomain{Parent: parent}
	case altStruct.DirectoryDomainParameter, *altStruct.DirectoryDomainParameter:
		parent, ok := par.(*altStruct.DirectoryDomain)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainParameter{Parent: parent}
	case altStruct.DirectoryDomainVariable, *altStruct.DirectoryDomainVariable:
		parent, ok := par.(*altStruct.DirectoryDomain)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainVariable{Parent: parent}
	case altStruct.DirectoryDomainUser, *altStruct.DirectoryDomainUser:
		parent, ok := par.(*altStruct.DirectoryDomain)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainUser{Parent: parent}
	case altStruct.DirectoryDomainUserParameter, *altStruct.DirectoryDomainUserParameter:
		parent, ok := par.(*altStruct.DirectoryDomainUser)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainUserParameter{Parent: parent}
	case altStruct.DirectoryDomainUserVariable, *altStruct.DirectoryDomainUserVariable:
		parent, ok := par.(*altStruct.DirectoryDomainUser)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainUserVariable{Parent: parent}
	case altStruct.DirectoryDomainUserGateway, *altStruct.DirectoryDomainUserGateway:
		parent, ok := par.(*altStruct.DirectoryDomainUser)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainUserGateway{Parent: parent}
	case altStruct.DirectoryDomainUserGatewayParameter, *altStruct.DirectoryDomainUserGatewayParameter:
		parent, ok := par.(*altStruct.DirectoryDomainUserGateway)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainUserGatewayParameter{Parent: parent}
	case altStruct.DirectoryDomainUserGatewayVariable, *altStruct.DirectoryDomainUserGatewayVariable:
		parent, ok := par.(*altStruct.DirectoryDomainUserGateway)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainUserGatewayVariable{Parent: parent}
	case altStruct.DirectoryDomainGroup, *altStruct.DirectoryDomainGroup:
		parent, ok := par.(*altStruct.DirectoryDomain)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainGroup{Parent: parent}
	case altStruct.DirectoryDomainGroupUser, *altStruct.DirectoryDomainGroupUser:
		parent, ok := par.(*altStruct.DirectoryDomainGroup)
		if !ok {
			parent = nil
		}
		return &altStruct.DirectoryDomainGroupUser{Parent: parent}
	default:
		return nil
	}
}

func CountDirectoryDomains() int64 {
	id := cache.GetCurrentInstanceId()
	filterStr := customorm.Filters{
		Count:  true,
		Fields: map[string]customorm.FilterFields{"Parent": {Flag: true}},
	}
	resCount, err := intermediateDB.GetByFilteredValues(
		&altStruct.DirectoryDomain{Parent: &mainStruct.FsInstance{Id: id}},
		filterStr,
	)
	if err != nil {
		return 0
	}
	if len(resCount) == 0 {
		return 0
	}
	total, ok := resCount[0].(int64)
	if !ok {
		return 0
	}
	return total
}

func IsDirectoryEnabled() bool {
	return CountDirectoryDomains() > 0
}
