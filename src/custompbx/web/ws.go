package web

import (
	"custompbx/daemonCache"
	"custompbx/webStruct"
)

func messageMainHandler(msg *webStruct.MessageData) webStruct.UserResponse {
	if !daemonCache.State.DatabaseConnection {
		return webStruct.UserResponse{Daemon: daemonCache.State, MessageType: webStruct.BroadcastConnection}
	}

	return webStruct.UserResponse{Error: "Wrong event", MessageType: "none"}
}
