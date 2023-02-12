package web

import (
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"encoding/json"
	"log"
)

func GetItems(data *webStruct.MessageData, user *mainStruct.WebUser, item *interface{}) webStruct.UserResponse {
	res, err := intermediateDB.GetAllFromDBAsMap(item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddItem(data *webStruct.MessageData, user *mainStruct.WebUser, item *interface{}) webStruct.UserResponse {
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	_, err = intermediateDB.InsertItem(&item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	return webStruct.UserResponse{MessageType: data.Event}
}

func DelItem(data *webStruct.MessageData, user *mainStruct.WebUser, item *interface{}) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err := intermediateDB.DeleteById(item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func UpdateItem(data *webStruct.MessageData, user *mainStruct.WebUser, item *interface{}) webStruct.UserResponse {
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
