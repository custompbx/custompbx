package intermediateDB

import "custompbx/altStruct"

func InitServicesDB() {

	corm := GetCORM()

	corm.CreateTable(&altStruct.ConversationRoom{})
	corm.CreateTable(&altStruct.ConversationRoomParticipant{})
	corm.CreateTable(&altStruct.ConversationRoomMessage{})
	corm.CreateTable(&altStruct.ConversationPrivateMessage{})
}
