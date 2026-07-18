package db

import (
	"crypto/sha256"
	"custompbx/mainStruct"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func InitWebDB(instanceId int64) {
	createWebUsersTable(db, instanceId)
	migrateWebUserLocales(db)
	createWebUsersTokensTable(db)
	migrateWebUserTokens(db)

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
		locale VARCHAR(16) NOT NULL DEFAULT 'en',
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

func migrateWebUserLocales(db *sql.DB) {
	_, err := db.Exec(`ALTER TABLE web_users ADD COLUMN IF NOT EXISTS locale VARCHAR(16) NOT NULL DEFAULT 'en'`)
	panicErr(err)
	_, err = db.Exec(`UPDATE web_users SET locale = CASE WHEN lang = 1 THEN 'ru' ELSE 'en' END WHERE locale IS NULL OR locale = '' OR (locale = 'en' AND lang = 1)`)
	panicErr(err)
	_, err = db.Exec(`UPDATE web_users SET locale = 'en' WHERE locale NOT IN ('en','fr','de','es','pt-BR','it','tr','ru','ar','fa','hi','zh-Hans','ja','ko')`)
	panicErr(err)
}

func createWebUsersTokensTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS web_users_tokens(
		id serial NOT NULL PRIMARY KEY,
		user_id BIGINT REFERENCES web_users (id) ON DELETE CASCADE,
		token VARCHAR,
		token_hash CHAR(64),
		created TIMESTAMP DEFAULT now(),
		expires_at TIMESTAMP,
		last_used TIMESTAMP,
		purpose VARCHAR DEFAULT 'gui',
		UNIQUE (user_id, token)
	)
	WITH (OIDS=FALSE);`,
	)
	panicErr(err)
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func migrateWebUserTokens(db *sql.DB) {
	_, err := db.Exec(`ALTER TABLE web_users_tokens ADD COLUMN IF NOT EXISTS token_hash CHAR(64), ADD COLUMN IF NOT EXISTS expires_at TIMESTAMP, ADD COLUMN IF NOT EXISTS last_used TIMESTAMP`)
	panicErr(err)
	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS web_users_tokens_hash_idx ON web_users_tokens(token_hash)`)
	panicErr(err)
	rows, err := db.Query(`SELECT id, token, purpose, created FROM web_users_tokens WHERE token_hash IS NULL AND token IS NOT NULL`)
	panicErr(err)
	type legacyToken struct {
		id             int64
		token, purpose string
		created        time.Time
	}
	var legacy []legacyToken
	for rows.Next() {
		var item legacyToken
		panicErr(rows.Scan(&item.id, &item.token, &item.purpose, &item.created))
		legacy = append(legacy, item)
	}
	panicErr(rows.Close())
	for _, item := range legacy {
		ttl := 24 * time.Hour
		if item.purpose == "api" {
			ttl = 90 * 24 * time.Hour
		}
		_, err = db.Exec(`UPDATE web_users_tokens SET token_hash=$1, token=NULL, expires_at=COALESCE(expires_at,$2) WHERE id=$3`, HashToken(item.token), time.Now().Add(ttl), item.id)
		panicErr(err)
	}
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
					wu.lang as lang,
					wu.locale as locale,
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
		err := user.Scan(&wUser.Id, &wUser.Login, &wUser.SipId, &wUser.WebRTCLib, &wUser.Ws, &wUser.VertoWs, &wUser.Stun, &wUser.Key, &wUser.Lang, &wUser.Locale, &wUser.Enabled)
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
					wu.locale as locale,
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
		err := user.Scan(&wUser.Id, &wUser.Login, &wUser.GroupId, &wUser.SipId, &wUser.WebRTCLib, &wUser.Ws, &wUser.VertoWs, &wUser.Stun, &wUser.Key, &wUser.Lang, &wUser.Locale, &wUser.Avatar, &wUser.AvatarFormat, &wUser.Enabled)
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
	ttl := 24 * time.Hour
	if purpose == "api" {
		ttl = 90 * 24 * time.Hour
	}
	err := db.QueryRow("INSERT INTO web_users_tokens(user_id, token_hash, purpose, expires_at) values($1, $2, $3, $4) returning id, to_char(created, 'YYYY-MM-DD HH24:MI:SS.MS'), purpose, to_char(expires_at, 'YYYY-MM-DD HH24:MI:SS.MS')", userId, HashToken(tokenValue), purpose, time.Now().Add(ttl)).Scan(&token.Id, &token.Created, &token.Purpose, &token.Expires)
	if err != nil {
		return mainStruct.WebUserToken{}, err
	}

	token.Token = tokenValue
	return token, err
}

func DelWebUserToken(userId int64, token string) error {
	_, err := db.Exec("DELETE FROM web_users_tokens WHERE user_id = $1 AND token_hash = $2", userId, HashToken(token))
	if err != nil {
		return err
	}
	return nil
}

func DelWebUserTokenById(id int64) (string, int64) {
	var userId int64
	err := db.QueryRow("DELETE FROM web_users_tokens WHERE id = $1 returning user_id", id).Scan(&userId)
	if err != nil {
		log.Printf("%+v", err)
		return "", 0
	}
	return "deleted", userId
}

func GetWebUserByToken(token string) (*mainStruct.WebUser, error) {
	user, err := db.Query(
		`WITH valid_token AS (
			UPDATE web_users_tokens SET last_used=now()
			WHERE token_hash=$1 AND (expires_at IS NULL OR expires_at > now())
			RETURNING user_id
		) SELECT
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
					wu.locale as locale,
					wu.avatar as avatar,
					wu.avatar_format as avatar_format,
					wu.enabled as enabled
				FROM web_users wu
				JOIN valid_token vt ON wu.id = vt.user_id
					LIMIT 1`,
		HashToken(token),
	)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	defer user.Close()
	var wUser mainStruct.WebUser
	wUser.Tokens = mainStruct.NewWebUserTokens()
	for user.Next() {
		err := user.Scan(&wUser.Id, &wUser.Login, &wUser.GroupId, &wUser.SipId, &wUser.WebRTCLib, &wUser.Ws, &wUser.VertoWs, &wUser.Stun, &wUser.Key, &wUser.Lang, &wUser.Locale, &wUser.Avatar, &wUser.AvatarFormat, &wUser.Enabled)
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

func UpdateWebUserLocale(id int64, locale string) bool {
	_, err := db.Exec(`UPDATE web_users SET locale = $2 WHERE id = $1`, id, locale)
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
					to_char(created, 'YYYY-MM-DD HH24:MI:SS.MS'),
					purpose,
					to_char(expires_at, 'YYYY-MM-DD HH24:MI:SS.MS'),
					COALESCE(to_char(last_used, 'YYYY-MM-DD HH24:MI:SS.MS'), '')
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
		err := user.Scan(&token.Id, &token.Created, &token.Purpose, &token.Expires, &token.LastUsed)
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
