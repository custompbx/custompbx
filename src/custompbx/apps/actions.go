package apps

import (
	"custompbx/altStruct"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"encoding/json"
	"github.com/custompbx/customorm"
	"log"
)

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerCompanies(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	res, err := intermediateDB.GetAllFromDBAsMap(&AutoDialerCompany{})
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerCompany(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerCompany{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Name == "" || item.Domain == nil || item.Domain.Id == 0 {
		return webStruct.UserResponse{Error: "no name", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerCompany(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err := intermediateDB.DeleteById(&AutoDialerCompany{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerCompany(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerCompany{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(&item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerTeams(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&AutoDialerTeam{Domain: &altStruct.DirectoryDomain{Id: data.Id}},
		map[string]bool{"Domain": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerTeam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerTeam{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Name == "" {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerTeam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := AutoDialerTeam{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerTeam(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerTeam{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerTeamMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&AutoDialerTeamMember{Parent: &AutoDialerTeam{Id: data.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerTeamMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	filter := map[string]customorm.FilterFields{"Parent": {Flag: true, UseValue: true, Value: data.Id}}

	_, err := intermediateDB.GetByValuesAsMap(
		&AutoDialerTeamMember{Parent: &AutoDialerTeam{Id: data.Id}},
		filter,
	)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	//TODO: remake after
	item := AutoDialerTeamMember{Parent: &AutoDialerTeam{Id: data.Id}}
	err = intermediateDB.DeleteRows(&item, map[string]bool{"Parent": true})
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	for _, id := range data.Ids {
		item = AutoDialerTeamMember{User: &altStruct.DirectoryDomainUser{Id: id}, Enabled: true, Parent: &AutoDialerTeam{Id: data.Id}}
		_, err := intermediateDB.InsertItem(&item)
		if err != nil {
			log.Println(err)
		}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: GetAutoDialerTeamMembers(data, user).Data}
}

func AddAutoDialerTeamMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerTeamMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Parent.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerTeamMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := AutoDialerTeamMember{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerTeamMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerTeamMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerLists(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&AutoDialerList{Domain: &altStruct.DirectoryDomain{Id: data.Id}},
		map[string]bool{"Domain": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerList{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Name == "" || item.Domain.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := AutoDialerList{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerList(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerList{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerListMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&AutoDialerListMember{Parent: &AutoDialerList{Id: data.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerListMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerListMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Parent.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func AddAutoDialerListMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerListMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Parent.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerListMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := AutoDialerListMember{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerListMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerListMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerReducers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerReducer{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Domain.Id == 0 {
		return webStruct.UserResponse{Error: "no domain id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&AutoDialerReducer{Domain: &altStruct.DirectoryDomain{Id: item.Domain.Id}},
		map[string]bool{"Domain": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerReducer(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerReducer{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Name == "" || item.Domain.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerReducer(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := AutoDialerReducer{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerReducer(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerReducer{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	result, err := intermediateDB.UpdateFunc(item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAutoDialerReducerMembers(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&AutoDialerReducerMember{Parent: &AutoDialerReducer{Id: data.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddAutoDialerReducerMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	var id int64
	item := AutoDialerReducerMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Application == "" || item.Parent.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	id, err = intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	result, _ := intermediateDB.GetByValueAsMap(
		&AutoDialerReducerMember{Id: id},
		map[string]bool{"Id": true},
	)

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelAutoDialerReducerMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := AutoDialerReducerMember{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateAutoDialerReducerMember(data *webStruct.MessageData, user *mainStruct.WebUser) webStruct.UserResponse {
	item := AutoDialerReducerMember{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	result, err := intermediateDB.UpdateFunc(&item, data.Fields)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	if contains(data.Fields, "position") {
		result, err = intermediateDB.GetByValueAsMap(
			&AutoDialerReducerMember{Parent: item.Parent},
			map[string]bool{"Parent": true},
		)
		if err != nil {
			log.Println(err)
			return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
		}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
