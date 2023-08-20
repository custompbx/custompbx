package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
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
