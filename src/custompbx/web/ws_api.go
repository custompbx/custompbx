package web

import (
	"custompbx/webStruct"
	"encoding/json"
	"log"
	"net/http"
)

func PostAPIRequest(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var msg webStruct.Message
	err := decoder.Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}
	if err := msg.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}

	var resp webStruct.UserResponse
	msg.Data.Trim()
	msg.Data.Event = msg.Event
	msg.Data.Context = webStruct.CreateWsContext(nil)

	// find user by token
	user, _ := messageUserLookup(msg.Data)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 Unauthorized"))
		return
	}
	msg.Data.Context.SetUser(user)

	resp = dispatchMessage(msg.Data, msg.Data.Context)
	res, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Server Error"))
		return
	}
	_, err = w.Write(res)
	if err != nil {
		log.Printf("%+v", err)
	}
}
