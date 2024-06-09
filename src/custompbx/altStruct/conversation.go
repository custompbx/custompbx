package altStruct

import (
	"custompbx/mainStruct"
	"time"
)

type ConversationRoom struct {
	Id       int64  `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position int64  `json:"position" customsql:"position;position"`
	Enabled  bool   `json:"enabled" customsql:"enabled;default=TRUE"`
	Title    string `xml:"title,attr" json:"name" customsql:"title;unique_1;check(title <> '')"`
}

func (w *ConversationRoom) GetTableName() string {
	return "conversation_rooms"
}

type ConversationRoomParticipant struct {
	Id       int64               `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position int64               `json:"position" customsql:"position;position"`
	Enabled  bool                `json:"enabled" customsql:"enabled;default=TRUE"`
	User     *mainStruct.WebUser `json:"user_id" customsql:"fkey:user_id;unique_1;check(user_id <> 0)"`
	Room     *ConversationRoom   `json:"room_id" customsql:"fkey:room_id;unique_1"`
}

func (w *ConversationRoomParticipant) GetTableName() string {
	return "conversation_room_participants"
}

type ConversationRoomMessage struct {
	Id          int64                        `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Participant *ConversationRoomParticipant `json:"participant_id" customsql:"fkey:conversation_room_participant_id;unique_1"`
	Room        *ConversationRoom            `json:"room_id" customsql:"fkey:conversation_room_id;unique_1"`
	CreatedAt   time.Time                    `json:"created_at" customsql:"created_at;index"`
	DeletedAt   time.Time                    `json:"deleted_at" customsql:"deleted_at;index"`
	Text        string                       `json:"text" customsql:"text"`
}

func (w *ConversationRoomMessage) GetTableName() string {
	return "conversation_room_messages"
}

type ConversationPrivateMessage struct {
	Id        int64               `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Sender    *mainStruct.WebUser `json:"sender_id" customsql:"fkey:sender_id;check(user_id <> 0)"`
	Receiver  *mainStruct.WebUser `json:"receiver_id" customsql:"fkey:receiver_id;check(user_id <> 0)"`
	CreatedAt time.Time           `json:"created_at" customsql:"created_at;index"`
	DeletedAt time.Time           `json:"deleted_at" customsql:"deleted_at;index"`
	Text      string              `json:"text" customsql:"text"`
}

func (w *ConversationPrivateMessage) GetTableName() string {
	return "conversation_private_messages"
}

type ConversationPrivateCall struct {
	Id        int64               `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Sender    *mainStruct.WebUser `json:"sender_id" customsql:"fkey:sender_id;check(user_id <> 0)"`
	Receiver  *mainStruct.WebUser `json:"receiver_id" customsql:"fkey:receiver_id;check(user_id <> 0)"`
	CreatedAt time.Time           `json:"created_at" customsql:"created_at;index"`
	DeletedAt time.Time           `json:"deleted_at" customsql:"deleted_at;index"`
	Duration  uint                `json:"duration" customsql:"duration;default=0"`
}

func (w *ConversationPrivateCall) GetTableName() string {
	return "conversation_private_calls"
}

type ConversationPrivateCallMessage struct {
	Id        int64                    `json:"id" customsql:"pkey:id;check(id <> 0)"`
	Sender    *mainStruct.WebUser      `json:"sender_id" customsql:"fkey:sender_id;check(user_id <> 0)"`
	Receiver  *mainStruct.WebUser      `json:"receiver_id" customsql:"fkey:receiver_id;check(user_id <> 0)"`
	CreatedAt time.Time                `json:"created_at" customsql:"created_at;index"`
	Text      string                   `json:"text" customsql:"text"`
	Call      *ConversationPrivateCall `xml:"-" json:"call" customsql:"fkey:call_id;check(call_id <> 0)"`
}

func (w *ConversationPrivateCallMessage) GetTableName() string {
	return "conversation_private_call_messages"
}
