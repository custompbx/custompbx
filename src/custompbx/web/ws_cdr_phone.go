package web

import (
	"custompbx/altStruct"
	"custompbx/cdrDb"
	"custompbx/cfg"
	"custompbx/daemonCache"
	"custompbx/intermediateDB"
	"custompbx/webStruct"
	"strconv"
)

func getCDR(data *webStruct.MessageData) webStruct.UserResponse {
	limit := data.DBRequest.Limit
	if limit == 0 || limit > 250 {
		limit = 25
	}
	offset := data.DBRequest.Offset * limit
	cdr, err := cdrDb.GetList(limit, offset, data.DBRequest.Filters, data.DBRequest.Order)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}
	/*
		if cdr == nil {
			return webStruct.UserResponse{Error: "nothing", MessageType: data.Event}
		}*/

	return webStruct.UserResponse{CDR: &cdr, MessageType: data.Event}
}

func getPhoneCreds(data *webStruct.MessageData) webStruct.UserResponse {
	if !data.Context.User.SipId.Valid {
		return webStruct.UserResponse{Error: "no config", MessageType: data.Event}
	}

	userI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: data.Context.User.SipId.Int64})
	if err != nil || userI == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	directoryUser, ok := userI.(altStruct.DirectoryDomainUser)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	param, err := intermediateDB.GetByValue(
		&altStruct.DirectoryDomainUserParameter{Name: "password", Parent: &directoryUser},
		map[string]bool{"Parent": true, "Name": true},
	)
	if err != nil || len(param) == 0 {
		return webStruct.UserResponse{Error: "user password not found", MessageType: data.Event}
	}
	directoryUserParam, ok := param[0].(altStruct.DirectoryDomainUserParameter)
	if !ok {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	domainI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomain{Id: directoryUser.Parent.Id})
	if err != nil || domainI == nil {
		return webStruct.UserResponse{Error: "user domain not found", MessageType: data.Event}
	}
	domain, ok := domainI.(altStruct.DirectoryDomain)
	if !ok {
		return webStruct.UserResponse{Error: "user domain not found", MessageType: data.Event}
	}

	password := directoryUserParam.Value

	if password == "" {
		paramI, err := intermediateDB.GetByValue(
			&altStruct.DirectoryDomainParameter{Name: "password", Parent: &domain},
			map[string]bool{"Parent": true, "Name": true},
		)
		if err != nil || len(paramI) == 0 {
			return webStruct.UserResponse{Error: "domain directory password not found", MessageType: data.Event}
		}
		domainParam, ok := paramI[0].(altStruct.DirectoryDomainParameter)
		if !ok {
			return webStruct.UserResponse{Error: "domain directory not found", MessageType: data.Event}
		}

		password = domainParam.Value
	}
	if password == "" || (data.Context.User.Ws == "" && data.Context.User.VertoWs == "") /*|| user.Stun == ""*/ {
		return webStruct.UserResponse{Error: "no enough params params", MessageType: data.Event}
	}

	creds := webStruct.PhoneCreds{}
	creds.UserName = directoryUser.Name
	creds.Password = password
	creds.Domain = domain.Name
	creds.WebRTCLib = data.Context.User.WebRTCLib
	creds.Ws = data.Context.User.Ws
	creds.VertoWs = data.Context.User.VertoWs
	creds.Stun = data.Context.User.Stun
	if creds.Stun == "" && daemonCache.State.StunServerStatus {
		creds.Stun = "stun:" + cfg.CustomPbx.Web.Host + ":" + strconv.Itoa(cfg.CustomPbx.Web.StunPort)
	}

	return webStruct.UserResponse{PhoneCreds: &creds, MessageType: data.Event}
}
