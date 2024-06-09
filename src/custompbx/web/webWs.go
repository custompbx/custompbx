package web

import (
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/cache"
	"custompbx/daemonCache"
	"custompbx/db"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/webStruct"
	"custompbx/webcache"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/custompbx/customorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"regexp"
	"time"
)

type templateItem struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Editable    bool   `json:"editable"`
}

type templateObj struct {
	Id         int64          `json:"id"`
	Name       string         `json:"name"`
	Variables  []templateItem `json:"variables"`
	Parameters []templateItem `json:"parameters"`
}

func checkLogin(data *webStruct.MessageData) webStruct.UserResponse {
	user := webcache.GetWebUser(data.Login)
	if user == nil {
		return webStruct.UserResponse{Error: "Unknown user", MessageType: data.Event}
	}

	if user.Login == "" || !CheckPassword(data.Password, []byte(user.Key)) {
		return webStruct.UserResponse{Error: "Wrong Login", MessageType: data.Event}
	}
	token := tokenGenerator()
	_, err := webcache.SaveWebUserToken(user, token, "gui")
	if err != nil {
		return webStruct.UserResponse{Error: "Cant set token", MessageType: data.Event}
	}

	data.Context.User = user
	return webStruct.UserResponse{User: user, Token: token, MessageType: data.Event}
}

func createAPIToken(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}
	neededUser := webcache.GetWebUserById(data.Id)
	if neededUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	token := tokenGenerator()
	tok, err := webcache.SaveWebUserToken(neededUser, token, "api")
	if err != nil {
		return webStruct.UserResponse{Error: "Cant set token", MessageType: data.Event}
	}

	tokSlice := []mainStruct.WebUserToken{tok}

	return webStruct.UserResponse{Id: &neededUser.Id, TokensList: &tokSlice, MessageType: data.Event}
}

func GetUserTokens(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}
	neededUser := webcache.GetWebUserById(data.Id)
	if neededUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	staff := webcache.GetWebUserTokens(neededUser)

	return webStruct.UserResponse{Id: &neededUser.Id, TokensList: &staff, MessageType: data.Event}
}

func RemoveUserToken(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	userId, err := webcache.DelWebUserTokenById(data.Id)
	if err != nil {
		return webStruct.UserResponse{Error: err.Error(), MessageType: data.Event}
	}

	return webStruct.UserResponse{Id: &userId, AffectedId: &data.Id, MessageType: data.Event}
}

func UserGetOwnTokens(data *webStruct.MessageData) webStruct.UserResponse {
	staff := webcache.GetWebUserTokens(data.Context.User)

	return webStruct.UserResponse{Id: &data.Context.User.Id, TokensList: &staff, MessageType: data.Event}
}

func loginOut(data *webStruct.MessageData) webStruct.UserResponse {
	err := webcache.DelWebUserToken(data.Context.User, data.Token)
	if err != nil {
		return webStruct.UserResponse{Error: "Cant delete token", MessageType: data.Event}
	}

	data.Context.User = nil
	return webStruct.UserResponse{MessageType: data.Event}
}

func findUser(data *webStruct.MessageData) (*mainStruct.WebUser, webStruct.UserResponse) {
	user, err := webcache.GetWebUserByToken(data.Token)
	if err != nil || user == nil || user.Login == "" {
		log.Println("EVENT: ", data.Event, "NO SUER BY TOKEN")
		if data.Context != nil {
			data.Context.Subscriptions.Clear()
		}
		noToken := true
		return nil, webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "connection", NoToken: &noToken}
	}
	if data.Context.User == nil {
		data.Context.User = user
	}
	if user.Id != data.Context.User.Id {
		log.Println("EVENT: ", data.Event, "USER: ", user.Login, " ACCESS DENIED! User changed in a single ws connection without logout")
		noToken := true
		return nil, webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "connection", NoToken: &noToken}
	}
	return user, webStruct.UserResponse{}
}

func checkAccessGroup(data *webStruct.MessageData, groupList []int) *webStruct.UserResponse {
	if data.Context.User == nil {
		log.Println("EVENT: ", data.Event, "no user!")
		return &webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "no_access"}
	}
	user := data.Context.User
	group := mainStruct.GetWebUserGroup(user.GroupId)
	if !group.ValidateGroupAccess(groupList) {
		log.Println("GROUP: ", user.GroupId, group, user.GroupId)
		log.Println("EVENT: ", data.Event, "USER: ", user.Login, " ACCESS DENIED!")
		return &webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "no_access"}
	}

	return nil
}

func subscribeUser(data *webStruct.MessageData) *webStruct.UserResponse {
	if data.Context.User == nil {
		log.Println("EVENT: ", data.Event, "no user!")
		return &webStruct.UserResponse{Daemon: daemonCache.State, MessageType: "no_access"}
	}
	log.Println("EVENT: ", data.Event, "USER: ", data.Context.User.Login)
	if data.KeepSubscription != nil && *data.KeepSubscription {
		data.Context.Subscriptions.Set(data.Event)
	}
	return nil
}

func getUser(data *webStruct.MessageData, foo func(*webStruct.MessageData) webStruct.UserResponse, groupList []int) webStruct.UserResponse {
	resp := checkAccessGroup(data, groupList)
	if resp != nil {
		return *resp
	}
	return foo(data)
}

func getUserForConfig(data *webStruct.MessageData, foo func(*webStruct.MessageData, interface{}) webStruct.UserResponse, conf interface{}, groupList []int) webStruct.UserResponse {
	resp := checkAccessGroup(data, groupList)
	if resp != nil {
		return *resp
	}
	return foo(data, conf)
}

func getWebUsers(data *webStruct.MessageData) webStruct.UserResponse {
	items := webcache.GetWebUsers()
	wssUris := webcache.GetWebMetaData().GetWssUris()
	VertoWsUris := webcache.GetWebMetaData().GetVertoWsUris()
	groups := mainStruct.GetWebUserGroups()
	return webStruct.UserResponse{WebUsers: &items, WebUsersGroups: &groups, Options: &wssUris, AltOptions: &VertoWsUris, MessageType: data.Event}
}

func GetWebUsersByDirectory(data *webStruct.MessageData) webStruct.UserResponse {
	items := webcache.GetWebUsersByDirectory()
	return webStruct.UserResponse{AdditionalData: &items, MessageType: data.Event}
}

func addWebUsers(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Login == "" || data.Password == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserByLogin(data.Login)
	if webUser != nil {
		return webStruct.UserResponse{Error: "user already exists", MessageType: data.Event}
	}
	group := mainStruct.GetWebUserGroup(data.GroupId)

	hashedPassword := HashPassword(data.Password)
	if hashedPassword == "" {
		return webStruct.UserResponse{Error: "unsuitable password", MessageType: data.Event}
	}
	webUser = webcache.AddWebUser(data.Login, hashedPassword, group.Id, cache.GetCurrentInstanceId())
	if webUser == nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func renameWebUsers(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Name == "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.RenameWebUser(webUser, data.Name)
	if !ok {
		return webStruct.UserResponse{Error: "can't rename", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func deleteWebUsers(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if webcache.CountWebUsers() < 2 {
		return webStruct.UserResponse{Error: "can't delete the last user", MessageType: data.Event}
	}
	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	if webUser.Id == data.Context.User.Id {
		return webStruct.UserResponse{Error: "you can't delete yourself", MessageType: data.Event}
	}

	ok := webcache.DelWebUser(webUser)
	if !ok {
		return webStruct.UserResponse{Error: "can't delete", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &webUser.Id}
}

func switchWebUser(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Enabled == nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	if webcache.CountWebUsers() < 2 && !*data.Enabled {
		return webStruct.UserResponse{Error: "can't disable the last user", MessageType: data.Event}
	}
	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	if webUser.Id == data.Context.User.Id && !*data.Enabled {
		return webStruct.UserResponse{Error: "you can't disable yourself", MessageType: data.Event}
	}

	ok := webcache.SwitchWebUser(webUser, *data.Enabled)
	if !ok {
		mes := "can't disable"
		if *data.Enabled {
			mes = "can't enable"
		}
		return webStruct.UserResponse{Error: mes, MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersPassword(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	if data.Password == "" || len(data.Password) < 6 {
		return webStruct.UserResponse{Error: "password is too short", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	hashedPassword := HashPassword(data.Password)
	if hashedPassword == "" {
		return webStruct.UserResponse{Error: "unsuitable password", MessageType: data.Event}
	}
	ok := webcache.UpdateWebUserPassword(webUser, hashedPassword)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersLang(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserLangId(webUser, data.ParamId)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersSipUser(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserSipId(webUser, data.ParamId)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersWs(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserWs(webUser, data.Value)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersVertoWs(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserVertoWs(webUser, data.Value)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func UpdateWebUserWebRTCLib(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserWebRTCLib(webUser, data.Value)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersStun(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserStun(webUser, data.Value)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func updateWebUsersAvatar(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	if data.File == "" {
		return webStruct.UserResponse{Error: "empty file string", MessageType: data.Event}
	}
	prefix := data.File[:30]
	r := regexp.MustCompile(`^(.*/(.+);.+,).*$`)
	match := r.FindStringSubmatch(prefix)
	if len(match) != 3 {
		return webStruct.UserResponse{Error: "file string not match", MessageType: data.Event}
	}
	fileString := data.File[len(match[1]):]
	res, err := base64.StdEncoding.DecodeString(fileString)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data " + err.Error(), MessageType: data.Event}
	}

	if len(res) > 512000 {
		return webStruct.UserResponse{Error: "file is too big", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserAvatar(webUser, fileString, match[2])
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func clearWebUsersAvatar(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}

	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}

	if data.File != "" {
		return webStruct.UserResponse{Error: "empty data", MessageType: data.Event}
	}

	ok := webcache.UpdateWebUserAvatar(webUser, "", "")
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

func GetWebSettings(data *webStruct.MessageData) webStruct.UserResponse {
	args := data.WebSettings
	if args == nil {
		return webStruct.UserResponse{Error: "empty request", MessageType: data.Event}
	}
	webSettings := map[string]string{}
	for key := range args {
		webSettings[key] = webcache.GetWebSetting(key)
	}
	return webStruct.UserResponse{WebSettings: &webSettings, MessageType: data.Event}
}

func SaveWebSettings(data *webStruct.MessageData) webStruct.UserResponse {
	args := data.WebSettings
	if args == nil {
		return webStruct.UserResponse{Error: "empty request", MessageType: data.Event}
	}
	webSettings := map[string]string{}
	for key, value := range args {
		oldValue := webcache.GetWebSetting(key)
		if value != oldValue {
			err := webcache.AddWebSetting(key, value, cache.GetCurrentInstanceId())
			if err != nil {
				continue
			}
		}
		webSettings[key] = value
	}
	return webStruct.UserResponse{WebSettings: &webSettings, MessageType: data.Event}
}

func UpdateWebUserGroup(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong id", MessageType: data.Event}
	}
	if data.Id == data.Context.User.Id {
		return webStruct.UserResponse{Error: "not allow to change own group", MessageType: data.Event}
	}
	webUser := webcache.GetWebUserById(data.Id)
	if webUser == nil {
		return webStruct.UserResponse{Error: "user not found", MessageType: data.Event}
	}
	group := mainStruct.GetWebUserGroup(data.GroupId)

	ok := webcache.ChangeWebUserGroup(webUser, group.Id)
	if !ok {
		return webStruct.UserResponse{Error: "can't change", MessageType: data.Event}
	}

	items := map[int64]*mainStruct.WebUser{webUser.Id: webUser}
	return webStruct.UserResponse{MessageType: data.Event, WebUsers: &items}
}

// /////////////////////////////////////////////////////////////////////////////////////////
func GetWebDirectoryUsersTemplates(data *webStruct.MessageData) webStruct.UserResponse {
	res, err := intermediateDB.GetAllFromDBAsMap(&altStruct.WebDirectoryUsersTemplate{})
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddWebDirectoryUsersTemplate(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplate{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
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

func DelWebDirectoryUsersTemplate(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err := intermediateDB.DeleteById(&altStruct.WebDirectoryUsersTemplate{Id: data.Id})
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func SwitchWebDirectoryUsersTemplate(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplate{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err = intermediateDB.UpdateByIdByValuesMap(
		&altStruct.WebDirectoryUsersTemplate{Id: item.Id, Enabled: !item.Enabled},
		map[string]bool{"Enabled": true},
	)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	item.Enabled = !item.Enabled
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func UpdateWebDirectoryUsersTemplate(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplate{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err = intermediateDB.UpdateByIdAll(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// /////////////////////////////////////////////////////////////////////////////////////////
func GetWebDirectoryUsersTemplateParameters(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&altStruct.WebDirectoryUsersTemplateParameter{Parent: &altStruct.WebDirectoryUsersTemplate{Id: data.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddWebDirectoryUsersTemplateParameter(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplateParameter{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Name == "" || item.Parent.Id == 0 {
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

func DelWebDirectoryUsersTemplateParameter(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := altStruct.WebDirectoryUsersTemplateParameter{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &data.Id}
}

func SwitchWebDirectoryUsersTemplateParameter(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplateParameter{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err = intermediateDB.UpdateByIdByValuesMap(
		&altStruct.WebDirectoryUsersTemplateParameter{Id: item.Id, Enabled: !item.Enabled},
		map[string]bool{"Enabled": true},
	)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	item.Enabled = !item.Enabled
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func UpdateWebDirectoryUsersTemplateParameter(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplateParameter{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err = intermediateDB.UpdateByIdAll(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

// /////////////////////////////////////////////////////////////////////////////////////////
func GetWebDirectoryUsersTemplateVariables(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	res, err := intermediateDB.GetByValueAsMap(
		&altStruct.WebDirectoryUsersTemplateVariable{Parent: &altStruct.WebDirectoryUsersTemplate{Id: data.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func AddWebDirectoryUsersTemplateVariable(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplateVariable{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Name == "" || item.Parent.Id == 0 {
		return webStruct.UserResponse{Error: "no name or parent id", MessageType: data.Event}
	}

	res, err := intermediateDB.InsertItem(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	item.Id = res
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func DelWebDirectoryUsersTemplateVariable(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	item := altStruct.WebDirectoryUsersTemplateVariable{Id: data.Id}
	err := intermediateDB.DeleteById(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't del", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Id: &item.Id}
}

func SwitchWebDirectoryUsersTemplateVariable(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplateVariable{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err = intermediateDB.UpdateByIdByValuesMap(
		&altStruct.WebDirectoryUsersTemplateVariable{Id: item.Id, Enabled: !item.Enabled},
		map[string]bool{"Enabled": true},
	)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	item.Enabled = !item.Enabled
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func UpdateWebDirectoryUsersTemplateVariable(data *webStruct.MessageData) webStruct.UserResponse {
	item := altStruct.WebDirectoryUsersTemplateVariable{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}

	err = intermediateDB.UpdateByIdAll(&item)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	var result interface{}
	result = item

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func GetWebDirectoryUsersTemplatesList(data *webStruct.MessageData) webStruct.UserResponse {
	var err error
	var res interface{}
	if data.Context.User.GroupId == mainStruct.GetAdminId() {
		res, err = intermediateDB.GetAllFromDBAsMap(&altStruct.WebDirectoryUsersTemplate{})
	} else {
		var userI interface{}
		userI, err = intermediateDB.GetByIdArg(
			&altStruct.DirectoryDomainUser{},
			data.Context.User.SipId.Int64,
		)
		sipUser, ok := userI.(altStruct.DirectoryDomainUser)
		if !ok || sipUser.Id == 0 {
			return webStruct.UserResponse{Error: "no directory sip user", MessageType: data.Event}
		}
		res, err = intermediateDB.GetByValueAsMap(
			&altStruct.WebDirectoryUsersTemplate{Domain: &altStruct.DirectoryDomain{Id: sipUser.Parent.Id}},
			map[string]bool{"Domain": true},
		)
	}

	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	var result interface{}
	result = res

	return webStruct.UserResponse{MessageType: data.Event, Data: &result}
}

func GetWebDirectoryUsersTemplateForm(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	var err error
	var res map[int64]interface{}
	var resStruct = templateObj{Id: data.Id}

	if data.Context.User.GroupId != mainStruct.GetAdminId() {
		var userI interface{}
		userI, err = intermediateDB.GetByIdArg(
			&altStruct.DirectoryDomainUser{},
			data.Context.User.SipId.Int64,
		)
		sipUser, ok := userI.(altStruct.DirectoryDomainUser)
		if !ok || sipUser.Id == 0 {
			return webStruct.UserResponse{Error: "no directory sip user", MessageType: data.Event}
		}
		res, err = intermediateDB.GetByValueAsMap(
			&altStruct.WebDirectoryUsersTemplate{Domain: &altStruct.DirectoryDomain{Id: sipUser.Parent.Id}},
			map[string]bool{"Domain": true},
		)
		if err != nil || res == nil {
			log.Println(err)
			return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
		}
		_, ok = res[data.Id]
		if !ok {
			return webStruct.UserResponse{Error: "no access", MessageType: data.Event}
		}
	}

	res, err = intermediateDB.GetByValueAsMap(
		&altStruct.WebDirectoryUsersTemplateParameter{Show: true, Parent: &altStruct.WebDirectoryUsersTemplate{Id: data.Id}},
		map[string]bool{"Parent": true, "Show": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	for _, it := range res {
		item, ok := it.(altStruct.WebDirectoryUsersTemplateParameter)
		if !ok {
			continue
		}
		name := item.Name
		if item.Placeholder != "" {
			name = item.Placeholder
		}
		resStruct.Parameters = append(resStruct.Parameters, templateItem{
			Id:          item.Id,
			Name:        name,
			Value:       item.Value,
			Description: item.Description,
			Editable:    item.Editable,
		})
	}

	res, err = intermediateDB.GetByValueAsMap(
		&altStruct.WebDirectoryUsersTemplateVariable{Show: true, Parent: &altStruct.WebDirectoryUsersTemplate{Id: data.Id}},
		map[string]bool{"Parent": true, "Show": true},
	)
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	for _, it := range res {
		item, ok := it.(altStruct.WebDirectoryUsersTemplateVariable)
		if !ok {
			continue
		}
		name := item.Name
		if item.Placeholder != "" {
			name = item.Placeholder
		}
		resStruct.Variables = append(resStruct.Variables, templateItem{
			Id:          item.Id,
			Name:        name,
			Value:       item.Value,
			Description: item.Description,
			Editable:    item.Editable,
		})
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &resStruct}
}

// ¯\_(ツ)_/¯
func CreateWebDirectoryUsersByTemplate(data *webStruct.MessageData) webStruct.UserResponse {
	item := templateObj{}
	err := json.Unmarshal(data.Data, &item)
	if err != nil {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	if item.Id == 0 {
		return webStruct.UserResponse{Error: "no id", MessageType: data.Event}
	}
	if item.Name == "" {
		return webStruct.UserResponse{Error: "no name", MessageType: data.Event}
	}

	var res map[int64]interface{}
	var template altStruct.WebDirectoryUsersTemplate
	var searchStruct = &altStruct.WebDirectoryUsersTemplate{Id: item.Id}
	var searchFields = map[string]bool{"Id": true}
	if data.Context.User.GroupId != mainStruct.GetAdminId() {
		sipUserI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: data.Context.User.SipId.Int64})
		if err != nil || sipUserI == nil {
			return webStruct.UserResponse{Error: "no directory sip user", MessageType: data.Event}
		}
		sipUser, ok := sipUserI.(altStruct.DirectoryDomainUser)
		if !ok {
			return webStruct.UserResponse{Error: "no directory sip user", MessageType: data.Event}
		}
		searchStruct.Domain = &altStruct.DirectoryDomain{Id: sipUser.Parent.Id}
		searchFields["Domain"] = true
	}

	res, err = intermediateDB.GetByValueAsMap(
		searchStruct,
		searchFields,
	)
	if err != nil || res == nil || len(res) == 0 {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	templateIntf, ok := res[item.Id]
	if !ok {
		return webStruct.UserResponse{Error: "no access", MessageType: data.Event}
	}

	template, ok = templateIntf.(altStruct.WebDirectoryUsersTemplate)
	if !ok || template.Id == 0 {
		return webStruct.UserResponse{Error: "no access", MessageType: data.Event}
	}

	// USER
	newUserId, err := altData.SetDirectoryDomainUser(template.Domain.Id, item.Name, "", template.Cidr, "")
	if err != nil {
		log.Println(err)
		return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
	}
	_, err = intermediateDB.UpdateFunc(&altStruct.DirectoryDomainUser{Id: newUserId, Cache: uint(template.Cache)}, []string{"Cache"})
	if err != nil {
		log.Println(err)
	}
	// PARAMS
	tParams, err := intermediateDB.GetByValueAsMap(
		&altStruct.WebDirectoryUsersTemplateParameter{Parent: &altStruct.WebDirectoryUsersTemplate{Id: template.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
		return webStruct.UserResponse{Error: "error while adding", MessageType: data.Event}
	}

	for _, v := range tParams {
		param, ok := v.(altStruct.WebDirectoryUsersTemplateParameter)
		if !ok {
			intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
			return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
		}
		var value = param.Value
		if param.Editable {
			for _, vv := range item.Parameters {
				if vv.Id != param.Id {
					continue
				}
				value = vv.Value
			}
		}
		_, err = intermediateDB.InsertItem(&altStruct.DirectoryDomainUserParameter{Parent: &altStruct.DirectoryDomainUser{Id: newUserId}, Name: param.Name, Value: value})
		if err != nil {
			log.Println(err)
			intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
			return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
		}
	}
	param, err := intermediateDB.GetByValue(&altStruct.DirectoryDomainUserParameter{Name: "password", Parent: &altStruct.DirectoryDomainUser{Id: newUserId}}, map[string]bool{"Name": true, "Parent": true})
	if len(param) == 0 {
		_, err = intermediateDB.InsertItem(&altStruct.DirectoryDomainUserParameter{Parent: &altStruct.DirectoryDomainUser{Id: newUserId}, Name: "password", Value: RandStringBytes(10)})
		if err != nil {
			log.Println(err)
			intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
			return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
		}
	}

	// VARIABLES
	tVars, err := intermediateDB.GetByValueAsMap(
		&altStruct.WebDirectoryUsersTemplateVariable{Parent: &altStruct.WebDirectoryUsersTemplate{Id: template.Id}},
		map[string]bool{"Parent": true},
	)
	if err != nil {
		log.Println(err)
		intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
		return webStruct.UserResponse{Error: "error while adding", MessageType: data.Event}
	}
	for _, v := range tVars {
		variable, ok := v.(altStruct.WebDirectoryUsersTemplateVariable)
		if !ok {
			intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
			return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
		}

		var value = variable.Value
		if variable.Editable {
			for _, vv := range item.Variables {
				if vv.Id != variable.Id {
					continue
				}
				value = vv.Value
			}
		}
		_, err = intermediateDB.InsertItem(&altStruct.DirectoryDomainUserVariable{Parent: &altStruct.DirectoryDomainUser{Id: newUserId}, Name: variable.Name, Value: value})
		if err != nil {
			log.Println(err)
			intermediateDB.DeleteById(&altStruct.DirectoryDomainUser{Id: newUserId})
			return webStruct.UserResponse{Error: "can't add", MessageType: data.Event}
		}
	}
	sipUserI, err := intermediateDB.GetByIdFromDB(&altStruct.DirectoryDomainUser{Id: newUserId})
	if err != nil || sipUserI == nil {
		return webStruct.UserResponse{Error: "no directory sip user", MessageType: data.Event}
	}
	return webStruct.UserResponse{Data: sipUserI, MessageType: data.Event}
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=!@#$%^&*()_+"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func CheckPassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func GetConversationRoomMessages(data *webStruct.MessageData) webStruct.UserResponse {
	var err error
	var res interface{}
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	var filter = customorm.Filters{}
	filter.SetLimit(20).
		SetOrder([]string{"created_at"}, true).
		EqualToValue("id", data.Id)
	if !data.UpToTime.IsZero() {
		filter.LessToValue("created_at", data.UpToTime)
	}

	if filter.Error != nil {
		return webStruct.UserResponse{Error: "wrong search data", MessageType: data.Event}
	}
	res, err = intermediateDB.GetByFilteredValues(&altStruct.ConversationRoomMessage{}, filter)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func GetConversationPrivateMessages(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	var upToTime string
	var args = []interface{}{data.Context.User.Id, data.Id, data.Id, data.Context.User.Id}
	if !data.UpToTime.IsZero() {
		upToTime = " AND created_at < $5"
		args = append(args, data.UpToTime)
	}
	sqlLine := fmt.Sprintf(`SELECT id, created_at, text, sender_id, receiver_id 
FROM conversation_private_messages 
WHERE ((sender_id = $1 AND receiver_id = $2) OR (sender_id = $3 AND receiver_id = $4)) %s ORDER BY created_at DESC  LIMIT 20;`, upToTime)

	r, err := db.GetDB().Query(sqlLine, args...)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	defer r.Close()
	var res []altStruct.ConversationPrivateMessage
	for r.Next() {
		mes := altStruct.ConversationPrivateMessage{Sender: &mainStruct.WebUser{}, Receiver: &mainStruct.WebUser{}}
		err := r.Scan(&mes.Id, &mes.CreatedAt, &mes.Text, &mes.Sender.Id, &mes.Receiver.Id)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		res = append(res, mes)
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func SendConversationPrivateMessage(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 || data.Text == "" {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	receiver := webcache.GetWebUserById(data.Id)
	if receiver == nil {
		return webStruct.UserResponse{Error: "unknown receiver id", MessageType: data.Event}
	}
	if len(data.Text) > 7000 {
		return webStruct.UserResponse{Error: "too big message", MessageType: data.Event}
	}

	msg := altStruct.ConversationPrivateMessage{
		Sender:    data.Context.User,
		Receiver:  receiver,
		Text:      data.Text,
		CreatedAt: time.Now(),
	}
	res, err := intermediateDB.InsertItem(&msg)
	if err != nil {
		return webStruct.UserResponse{Error: "can't send", MessageType: data.Event}
	}
	mes, err := intermediateDB.GetByIdArg(msg, res)
	if err != nil {
		return webStruct.UserResponse{Error: "can't send", MessageType: data.Event}
	}
	sent := b.Unicast(webStruct.UserResponse{MessageType: "NewMessage", Data: mes}, []*mainStruct.WebUser{{Id: data.Context.User.Id}, {Id: data.Id}})

	return webStruct.UserResponse{MessageType: data.Event, Data: sent}
}

func GetConversationPrivateCalls(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	var upToTime string
	var args = []interface{}{data.Context.User.Id, data.Id, data.Id, data.Context.User.Id}
	if !data.UpToTime.IsZero() {
		upToTime = " AND created_at < $5"
		args = append(args, data.UpToTime)
	}
	sqlLine := fmt.Sprintf(`SELECT id, created_at, duration, sender_id, receiver_id 
FROM conversation_private_calls 
WHERE ((sender_id = $1 AND receiver_id = $2) OR (sender_id = $3 AND receiver_id = $4)) %s ORDER BY created_at DESC  LIMIT 20;`, upToTime)

	r, err := db.GetDB().Query(sqlLine, args...)
	if err != nil {
		return webStruct.UserResponse{Error: "can't get", MessageType: data.Event}
	}
	defer r.Close()
	var res []altStruct.ConversationPrivateCall
	for r.Next() {
		mes := altStruct.ConversationPrivateCall{Sender: &mainStruct.WebUser{}, Receiver: &mainStruct.WebUser{}}
		err := r.Scan(&mes.Id, &mes.CreatedAt, &mes.Duration, &mes.Sender.Id, &mes.Receiver.Id)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		res = append(res, mes)
	}

	return webStruct.UserResponse{MessageType: data.Event, Data: &res}
}

func SendConversationPrivateCall(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}
	receiver := webcache.GetWebUserById(data.Id)
	if receiver == nil {
		return webStruct.UserResponse{Error: "unknown receiver id", MessageType: data.Event}
	}

	msg := altStruct.ConversationPrivateCall{
		Sender:    data.Context.User,
		Receiver:  receiver,
		CreatedAt: time.Now(),
	}
	res, err := intermediateDB.InsertItem(&msg)
	if err != nil {
		return webStruct.UserResponse{Error: "can't send", MessageType: data.Event}
	}
	mes, err := intermediateDB.GetByIdArg(msg, res)
	if err != nil {
		return webStruct.UserResponse{Error: "can't send", MessageType: data.Event}
	}
	sent := b.Unicast(webStruct.UserResponse{MessageType: "NewMessage", Data: mes}, []*mainStruct.WebUser{{Id: data.Context.User.Id}, {Id: data.Id}})

	return webStruct.UserResponse{MessageType: data.Event, Data: sent}
}

func SendConversationRoomMessage(data *webStruct.MessageData) webStruct.UserResponse {
	if data.Id == 0 || data.Text == "" {
		return webStruct.UserResponse{Error: "wrong data", MessageType: data.Event}
	}

	receiver, err := intermediateDB.GetByIdArg(&altStruct.ConversationRoom{}, data.Id)
	if err != nil || receiver == nil {
		return webStruct.UserResponse{Error: "unknown room id", MessageType: data.Event}
	}
	if len(data.Text) > 7000 {
		return webStruct.UserResponse{Error: "too big message", MessageType: data.Event}
	}

	room, ok := receiver.(altStruct.ConversationRoom)
	if !ok || room.Id == 0 {
		return webStruct.UserResponse{Error: "unknown room", MessageType: data.Event}
	}

	roomParticipant, err := intermediateDB.GetByValueAsMap(
		&altStruct.ConversationRoomParticipant{
			Room: &room,
		},
		map[string]bool{"room_id": true},
	)
	var sendTo []*mainStruct.WebUser
	var author *altStruct.ConversationRoomParticipant
	for _, v := range roomParticipant {
		participant, ok := v.(altStruct.ConversationRoomParticipant)
		if !ok {
			continue
		}
		sendTo = append(sendTo, participant.User)
		if participant.User.Id == data.Context.User.Id {
			author = &participant
		}
	}
	if author == nil {
		return webStruct.UserResponse{Error: "author is not participant of the room", MessageType: data.Event}
	}

	msg := altStruct.ConversationRoomMessage{
		Room:        &room,
		Participant: author,
		Text:        data.Text,
		CreatedAt:   time.Now(),
	}
	res, err := intermediateDB.InsertItem(&msg)
	if err != nil {
		return webStruct.UserResponse{Error: "can't send", MessageType: data.Event}
	}
	mes, err := intermediateDB.GetByIdArg(msg, res)
	if err != nil {
		return webStruct.UserResponse{Error: "can't send", MessageType: data.Event}
	}

	sent := b.Unicast(webStruct.UserResponse{MessageType: "NewMessage", Data: mes}, sendTo)

	return webStruct.UserResponse{MessageType: data.Event, Data: sent}
}
