package db

import (
	"custompbx/mainStruct"
	"database/sql"
	"fmt"
	"log"
)

func InitWebDB(instanceId int64) {
	createWebUsersTable(db, instanceId)
	createWebUsersTokensTable(db)

	//corm := customOrm.Init(db)
	//corm.CreateTable(mainStruct.WebDirectoryUsersTemplate{})
	//CreateTableByStruct(&mainStruct.WebDirectoryUsersTemplate{})
}

func createWebUsersTable(db *sql.DB, instanceId int64) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS web_users(
		id serial NOT NULL PRIMARY KEY,
		login VARCHAR UNIQUE,
		group_id INTEGER NOT NULL DEFAULT 0,
		sip_id BIGINT REFERENCES directory_domain_users (id) ON DELETE SET NULL,
		webrtc_lib VARCHAR DEFAULT '` + mainStruct.WebRTCLibSipJs + `',
		ws VARCHAR DEFAULT '',
		verto_ws VARCHAR DEFAULT '',
		stun VARCHAR DEFAULT '',
		key VARCHAR,
		lang INTEGER DEFAULT 0,
		avatar TEXT DEFAULT '',
		avatar_format VARCHAR DEFAULT '',
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE
	)
	WITH (OIDS=FALSE);`,
	)
	panicErr(err)

	sqlReq := fmt.Sprintf(
		"INSERT INTO web_users(login, key, group_id, instance_id) SELECT '%s', '%s', %d, %d WHERE NOT EXISTS (SELECT login FROM web_users WHERE login = '%s');",
		"admin",
		"$2a$10$kq/GYf1EVEm7GKks6VbD6.ghCwDNDlucW/rPs8pDeolY23kX2XieW",
		1,
		instanceId,
		"admin",
	)
	db.QueryRow(sqlReq)
}

func createWebUsersTokensTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS web_users_tokens(
		id serial NOT NULL PRIMARY KEY,
		user_id BIGINT REFERENCES web_users (id) ON DELETE CASCADE,
		token VARCHAR,
		created TIMESTAMP DEFAULT now(),
		purpose VARCHAR DEFAULT 'gui',
		UNIQUE (user_id, token)
	)
	WITH (OIDS=FALSE);`,
	)
	panicErr(err)
}

func GetWebUser(login string, instanceId int64) (*mainStruct.WebUser, error) {
	user, err := db.Query(
		`SELECT
					wu.id as id,
					wu.login as login,
					wu.sip_id as sip_id,
					wu.webrtc_lib as webrtc_lib,
					wu.ws as ws,
					wu.verto_ws as verto_ws,
					wu.stun as stun,
					wu.key as key,
					wu.enabled as enabled
				FROM web_users wu
				WHERE
					wu.login = $1 AND instance_id = $2`,
		login,
		instanceId,
	)
	defer user.Close()
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	var wUser mainStruct.WebUser
	for user.Next() {
		err := user.Scan(&wUser.Id, &wUser.Login, &wUser.SipId, &wUser.WebRTCLib, &wUser.Ws, &wUser.VertoWs, &wUser.Stun, &wUser.Key, &wUser.Lang, &wUser.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
	}
	return &wUser, nil
}

func GetWebUsers(users *mainStruct.WebUsers, instanceId int64) {
	user, err := db.Query(
		`SELECT
					wu.id as id,
					wu.login as login,
					wu.group_id as group_id,
					wu.sip_id as sip_id,
					wu.webrtc_lib as webrtc_lib,
					wu.ws as ws,
					wu.verto_ws as verto_ws,
					wu.stun as stun,
					wu.key as key,
					wu.lang as lang,
					wu.avatar as avatar,
					wu.avatar_format as avatar_format,
					wu.enabled as enabled
				FROM web_users wu
				WHERE instance_id = $1`,
		instanceId,
	)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer user.Close()
	for user.Next() {
		var wUser mainStruct.WebUser
		err := user.Scan(&wUser.Id, &wUser.Login, &wUser.GroupId, &wUser.SipId, &wUser.WebRTCLib, &wUser.Ws, &wUser.VertoWs, &wUser.Stun, &wUser.Key, &wUser.Lang, &wUser.Avatar, &wUser.AvatarFormat, &wUser.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		wUser.Tokens = mainStruct.NewWebUserTokens()
		users.Set(&wUser)
	}
}

func SaveWebUserToken(userId int64, tokenValue, purpose string) (mainStruct.WebUserToken, error) {
	if purpose != "gui" && purpose != "api" {
		purpose = "gui"
	}
	var token mainStruct.WebUserToken
	err := db.QueryRow("INSERT INTO web_users_tokens(user_id, token, purpose) values($1, $2, $3) returning id, token, to_char(created, 'YYYY-MM-DD HH24:MI:SS.MS'), purpose", userId, tokenValue, purpose).Scan(&token.Id, &token.Token, &token.Created, &token.Purpose)
	if err != nil {
		return mainStruct.WebUserToken{}, err
	}

	return token, err
}

func DelWebUserToken(userId int64, token string) error {
	_, err := db.Exec("DELETE FROM web_users_tokens WHERE user_id = $1 AND token = $2", userId, token)
	if err != nil {
		return err
	}
	return nil
}

func DelWebUserTokenById(id int64) (string, int64) {
	var token string
	var userId int64
	err := db.QueryRow("DELETE FROM web_users_tokens WHERE id = $1 returning token, user_id", id).Scan(&token, &userId)
	if err != nil {
		log.Printf("%+v", err)
		return "", 0
	}
	return token, userId
}

func GetWebUserByToken(token string) (*mainStruct.WebUser, error) {
	user, err := db.Query(
		`SELECT
					wu.id as id,
					wu.login as login,
					wu.group_id as group_id,
					wu.sip_id as sip_id,
					wu.webrtc_lib as webrtc_lib,
					wu.ws as ws,
					wu.verto_ws as verto_ws,
					wu.stun as stun,
					wu.key as key,
					wu.lang as lang,
					wu.avatar as avatar,
					wu.avatar_format as avatar_format,
					wu.enabled as enabled
				FROM web_users wu
				LEFT JOIN web_users_tokens wut ON wu.id = wut.user_id
				WHERE
					wut.token = $1
					LIMIT 1`,
		token,
	)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	defer user.Close()
	var wUser mainStruct.WebUser
	wUser.Tokens = mainStruct.NewWebUserTokens()
	for user.Next() {
		err := user.Scan(&wUser.Id, &wUser.Login, &wUser.GroupId, &wUser.SipId, &wUser.WebRTCLib, &wUser.Ws, &wUser.VertoWs, &wUser.Stun, &wUser.Key, &wUser.Lang, &wUser.Avatar, &wUser.AvatarFormat, &wUser.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
	}
	return &wUser, nil
}

func DeleteWebUser(id int64) bool {
	_, err := db.Exec(
		`DELETE FROM web_users WHERE id = $1`, id)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func RenameWebUser(id int64, login string) bool {
	_, err := db.Exec(
		`UPDATE web_users SET login = $2 WHERE id = $1`, id, login)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func AddWebUser(login, key string, groupId int, instanceId int64) int64 {
	var id int64
	err := db.QueryRow(
		`INSERT INTO web_users(login, key, group_id, instance_id) VALUES($1, $2, $3, $4) RETURNING id;`, login, key, groupId, instanceId).Scan(&id)
	if err != nil {
		log.Printf("%+v", err)
		return id
	}
	return id
}

func SwitchWebUser(id int64, switcher bool) bool {
	_, err := db.Exec(
		`UPDATE web_users SET enabled = $1 WHERE id = $2`, switcher, id)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserPassword(id int64, key string) bool {
	_, err := db.Exec(
		`UPDATE web_users SET key = $2 WHERE id = $1`, id, key)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserLangId(id, langId int64) bool {
	_, err := db.Exec(
		`UPDATE web_users SET lang = $2 WHERE id = $1`, id, langId)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserSipId(id, sipId int64) bool {
	_, err := db.Exec(
		`UPDATE web_users SET sip_id = $2 WHERE id = $1`, id, sipId)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserWs(id int64, ws string) bool {
	_, err := db.Exec(
		`UPDATE web_users SET ws = $2 WHERE id = $1`, id, ws)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserVertoWs(id int64, ws string) bool {
	_, err := db.Exec(
		`UPDATE web_users SET verto_ws = $2 WHERE id = $1`, id, ws)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserWebRTCLib(id int64, lib string) bool {
	if lib != mainStruct.WebRTCLibVerto {
		lib = mainStruct.WebRTCLibSipJs
	}
	_, err := db.Exec(
		`UPDATE web_users SET webrtc_lib  = $2 WHERE id = $1`, id, lib)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserStun(id int64, stun string) bool {
	_, err := db.Exec(
		`UPDATE web_users SET stun = $2 WHERE id = $1`, id, stun)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func UpdateWebUserAvatar(id int64, avatar, avatarFormat string) bool {
	_, err := db.Exec(
		`UPDATE web_users SET avatar = $2, avatar_format = $3 WHERE id = $1`, id, avatar, avatarFormat)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}

func GetWebUserTokens(userId int64) ([]mainStruct.WebUserToken, error) {
	user, err := db.Query(
		`SELECT
					id,
					token,
					to_char(created, 'YYYY-MM-DD HH24:MI:SS.MS'),
                    purpose
				FROM web_users_tokens
				WHERE
					user_id = $1
					ORDER BY id`,
		userId,
	)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	defer user.Close()
	var tokens []mainStruct.WebUserToken
	for user.Next() {
		var token mainStruct.WebUserToken
		err := user.Scan(&token.Id, &token.Token, &token.Created, &token.Purpose)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func UpdateWebUserGroup(id int64, groupId int) bool {
	_, err := db.Exec(
		`UPDATE web_users SET group_id = $2 WHERE id = $1`, id, groupId)
	if err != nil {
		log.Printf("%+v", err)
		return false
	}
	return true
}
