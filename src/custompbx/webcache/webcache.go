package webcache

import (
	"custompbx/db"
	"custompbx/mainStruct"
	"database/sql"
	"errors"
	"github.com/custompbx/customorm"
)

var users *mainStruct.WebUsers
var dashboard *mainStruct.Dashboard
var webMetaData *mainStruct.WebMetaData
var webSettings *mainStruct.WebSettings

const (
	CdrModule          = "cdr_module"
	CdrTable           = "cdr_table"
	CdrFileServeColumn = "cdr_file_serve_column"
	CdrFileServerPath  = "cdr_file_server_path"
)

func InitCacheObjects() {
	users = mainStruct.NewUsersCache()
	dashboard = mainStruct.NewDashboard()
	webMetaData = mainStruct.NewWebMetaData()
	webSettings = mainStruct.NewWebSettings()
}

func InitUsersCache() {
	db.GetWebUsers(users)
}

func InitWebSettings(instanceId int64) {
	db.GetWebSettings(webSettings, instanceId)
}

func InitWebData() {
	corm := customorm.Init(db.GetDB())

	corm.CreateTable(&mainStruct.WebDirectoryUsersTemplate{})
	corm.CreateTable(&mainStruct.WebDirectoryUsersTemplateParameter{})
	corm.CreateTable(&mainStruct.WebDirectoryUsersTemplateVariable{})
}

func GetWebSetting(key string) string {
	return webSettings.Get(key)
}

func AddWebSetting(key, value string, instanceId int64) error {
	err := db.AddWebSetting(key, value, instanceId)
	if err != nil {
		return err
	}
	webSettings.Set(key, value)
	return nil
}

func GetWebUser(login string) *mainStruct.WebUser {
	return users.GetByLogin(login)
}

func GetWebUserById(id int64) *mainStruct.WebUser {
	return users.GetById(id)
}

func GetWebUserByLogin(login string) *mainStruct.WebUser {
	return users.GetByLogin(login)
}

func CountWebUsers() int {
	return users.Count()
}

func DelWebUser(user *mainStruct.WebUser) bool {
	if user == nil {
		return false
	}
	ok := db.DeleteWebUser(user.Id)
	if !ok {
		return false
	}

	users.Remove(user)
	return true
}

func SwitchWebUser(user *mainStruct.WebUser, switcher bool) bool {
	if user == nil {
		return false
	}
	ok := db.SwitchWebUser(user.Id, switcher)
	if !ok {
		return false
	}
	user.Enabled = switcher
	return true
}

func RenameWebUser(user *mainStruct.WebUser, login string) bool {
	if user == nil || login == "" {
		return false
	}
	ok := db.RenameWebUser(user.Id, login)
	if !ok {
		return false
	}

	users.Rename(user.Login, login)
	return true
}

func UpdateWebUserPassword(user *mainStruct.WebUser, key string) bool {
	if user == nil || key == "" {
		return false
	}
	ok := db.UpdateWebUserPassword(user.Id, key)
	if !ok {
		return false
	}

	user.Key = key
	return true
}

func UpdateWebUserLangId(user *mainStruct.WebUser, id int64) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserLangId(user.Id, id)
	if !ok {
		return false
	}

	user.Lang = uint(id)
	return true
}

func UpdateWebUserSipId(user *mainStruct.WebUser, id int64) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserSipId(user.Id, id)
	if !ok {
		return false
	}

	user.SipId = sql.NullInt64{Int64: id, Valid: true}
	return true
}

func UpdateWebUserWs(user *mainStruct.WebUser, ws string) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserWs(user.Id, ws)
	if !ok {
		return false
	}

	user.Ws = ws
	return true
}

func UpdateWebUserVertoWs(user *mainStruct.WebUser, ws string) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserVertoWs(user.Id, ws)
	if !ok {
		return false
	}

	user.VertoWs = ws
	return true
}

func UpdateWebUserWebRTCLib(user *mainStruct.WebUser, lib string) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserWebRTCLib(user.Id, lib)
	if !ok {
		return false
	}

	user.WebRTCLib = lib
	return true
}

func UpdateWebUserStun(user *mainStruct.WebUser, stun string) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserStun(user.Id, stun)
	if !ok {
		return false
	}

	user.Stun = stun
	return true
}

func UpdateWebUserAvatar(user *mainStruct.WebUser, avatar, avatarFormat string) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserAvatar(user.Id, avatar, avatarFormat)
	if !ok {
		return false
	}

	user.Avatar = avatar
	user.AvatarFormat = avatarFormat
	return true
}

func AddWebUser(login, key string, groupId int) *mainStruct.WebUser {
	if login == "" {
		return nil
	}
	id := db.AddWebUser(login, key, groupId)
	if id == 0 {
		return nil
	}
	user := &mainStruct.WebUser{}
	user.Id = id
	user.Login = login
	user.Key = key
	user.Enabled = true
	user.Tokens = mainStruct.NewWebUserTokens()
	user.GroupId = mainStruct.GetWebUserGroup(groupId).Id
	users.Set(user)

	return user
}

func GetWebUsers() map[int64]*mainStruct.WebUser {
	return users.GetList()
}

func GetWebUsersByDirectory() map[int64]*interface{} {
	return users.GetListByDirectory()
}

func SaveWebUserToken(user *mainStruct.WebUser, token, purpose string) (mainStruct.WebUserToken, error) {
	if user == nil {
		return mainStruct.WebUserToken{}, errors.New("user not found")
	}

	tok, err := db.SaveWebUserToken(user.Id, token, purpose)
	if err != nil {
		return mainStruct.WebUserToken{}, err
	}

	user.Tokens.Set(token)

	return tok, err
}

func DelWebUserToken(user *mainStruct.WebUser, token string) error {
	if user == nil {
		return errors.New("user not found")
	}

	err := db.DelWebUserToken(user.Id, token)
	if err != nil {
		return err
	}

	user.Tokens.Delete(token)

	return err
}

func DelWebUserTokenById(id int64) (int64, error) {
	token, userId := db.DelWebUserTokenById(id)
	if token == "" || userId == 0 {
		return 0, errors.New("token not found")
	}

	user := users.GetById(userId)
	if user == nil {
		return 0, errors.New("user not found")
	}
	user.Tokens.Delete(token)

	return user.Id, nil
}

func GetWebUserByToken(token string) (*mainStruct.WebUser, error) {
	var err error
	user, find := users.GetByToken(token)
	if !find || user == nil {
		user, err = db.GetWebUserByToken(token)
	}
	var userRes *mainStruct.WebUser
	if user != nil && user.Id != 0 {
		oldUser := users.GetById(user.Id)
		if oldUser != nil {
			oldUser.Tokens.Set(token)
			userRes = oldUser
		} else {
			user.Tokens.Set(token)
			users.Set(user)
			userRes = user
		}
	}

	return userRes, err
}

func CheckUser() {

}

func ClearUser() {

}

func DashBoardSetStaticData(data mainStruct.DashboardData) {
	dashboard.Timestamp = data.Timestamp
	dashboard.Hostname = data.Hostname
	dashboard.OS = data.OS
	dashboard.Platform = data.Platform
	dashboard.CPUModel = data.CPUModel
	dashboard.CPUFrequency = data.CPUFrequency
	dashboard.DynamicMetrics.PercentageUsedMemory = data.DynamicMetrics.PercentageUsedMemory
	dashboard.DynamicMetrics.TotalMemory = data.DynamicMetrics.TotalMemory
	dashboard.DynamicMetrics.FreeMemory = data.DynamicMetrics.FreeMemory
	dashboard.DynamicMetrics.TotalDiscSpace = data.DynamicMetrics.TotalDiscSpace
	dashboard.DynamicMetrics.FreeDiskSpace = data.DynamicMetrics.FreeDiskSpace
	dashboard.DynamicMetrics.PercentageDiskUsage = data.DynamicMetrics.PercentageDiskUsage
	dashboard.DynamicMetrics.CoreUtilization = data.DynamicMetrics.CoreUtilization
}

func DashBoardSetSipRegs(regs map[string]int) {
	dashboard.DomainSipRegs = regs
}

func DashBoardSetSofiaData(profiles []interface{}, gateways []interface{}) {
	dashboard.SofiaProfiles = profiles
	dashboard.SofiaGateways = gateways
}

func DashBoardSetCallsCounter(total int, answered int) {
	dashboard.CallsCounter = map[string]int{"total": total, "answered": answered}
}

func GetDashboardData() *mainStruct.DashboardData {
	return dashboard.DashboardData
}
func GetDashboardFSMetrics() *mainStruct.FSMetrics {
	return dashboard.FSMetrics
}

func GetWebMetaData() *mainStruct.WebMetaData {
	return webMetaData
}

func GetWebUserTokens(user *mainStruct.WebUser) []mainStruct.WebUserToken {
	if user == nil {
		return []mainStruct.WebUserToken{}
	}

	tokens, err := db.GetWebUserTokens(user.Id)
	if err != nil {
		return []mainStruct.WebUserToken{}
	}

	return tokens
}

func ChangeWebUserGroup(user *mainStruct.WebUser, groupId int) bool {
	if user == nil {
		return false
	}
	ok := db.UpdateWebUserGroup(user.Id, groupId)
	if !ok {
		return false
	}

	user.GroupId = groupId
	return true
}
