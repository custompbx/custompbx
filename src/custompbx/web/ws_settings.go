package web

import (
	"custompbx/cfg"
	"custompbx/logsafe"
	"custompbx/webStruct"
	"log"
)

func checkRelogin(data *webStruct.MessageData) webStruct.UserResponse {
	return webStruct.UserResponse{User: data.Context.User, Token: data.Token, MessageType: "relogin"}
}

func checkSettings(data *webStruct.MessageData) webStruct.UserResponse {
	return webStruct.UserResponse{Settings: &cfg.CustomPbx, MessageType: "settings"}
}

func setSettings(data *webStruct.MessageData) webStruct.UserResponse {
	log.Printf("settings update requested payload=%s", logsafe.Redact(data.Payload))
	if data.Payload.Fs.Esl.Pass == "" || data.Payload.Fs.Esl.Port == 0 || data.Payload.Fs.Esl.Host == "" ||
		data.Payload.Db.Host == "" || data.Payload.Db.Port == 0 || data.Payload.Db.Name == "" ||
		data.Payload.Db.User == "" || data.Payload.Db.Pass == "" ||
		data.Payload.Web.Host == "" || data.Payload.Web.Port == 0 ||
		data.Payload.Web.Route == "" ||
		data.Payload.XMLCurl.Host == "" || data.Payload.XMLCurl.Port == 0 ||
		data.Payload.XMLCurl.Route == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: "settings"}
	}
	cfg.CustomPbx.Fs.Esl = data.Payload.Fs.Esl
	cfg.CustomPbx.Db = data.Payload.Db
	cfg.CustomPbx.Web = data.Payload.Web
	cfg.CustomPbx.XMLCurl = data.Payload.XMLCurl
	conf, err := cfg.WD(cfg.CustomPbx)
	if err != nil {
		cfg.RD()
		return webStruct.UserResponse{Error: "can't save", MessageType: "settings"}
	}

	return webStruct.UserResponse{Settings: &conf, MessageType: "settings"}
}
