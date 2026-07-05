//go:build integration

package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestHEPDetailsCallIDsAreParameterized(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_DSN")
	if dsn == "" {
		t.Skip("TEST_DATABASE_DSN is not configured")
	}
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if err = conn.Ping(); err != nil {
		t.Fatal(err)
	}
	db = conn
	_, err = db.Exec(`DROP TABLE IF EXISTS hep_packets; CREATE TABLE hep_packets (
		hep_timestamp timestamp, hep_dst_ip text, hep_src_ip text, hep_dst_port int,
		hep_src_port int, hep_payload text, sip_first_method text, sip_cseq_method text,
		sip_call_id text, instance_id bigint
	)`)
	if err != nil {
		t.Fatal(err)
	}
	malicious := `call-id'); SELECT pg_sleep(10); --`
	_, err = db.Exec(`INSERT INTO hep_packets VALUES (now(),'dst','src',1,2,'payload','INVITE','INVITE',$1,7)`, malicious)
	if err != nil {
		t.Fatal(err)
	}
	rows, err := GetHEPDetailsList([]string{malicious}, 7)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("got %d rows, want 1", len(rows))
	}
}

func TestLegacyTokensAreHashedAndRemainValid(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_DSN")
	if dsn == "" {
		t.Skip("TEST_DATABASE_DSN is not configured")
	}
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if err = conn.Ping(); err != nil {
		t.Fatal(err)
	}
	db = conn
	_, err = db.Exec(`DROP TABLE IF EXISTS web_users_tokens; DROP TABLE IF EXISTS web_users; CREATE TABLE web_users (
		id bigserial PRIMARY KEY, login text, group_id int, sip_id bigint, webrtc_lib text, ws text,
		verto_ws text, stun text, key text, lang int, avatar text, avatar_format text, enabled boolean
	); CREATE TABLE web_users_tokens (
		id bigserial PRIMARY KEY, user_id bigint REFERENCES web_users(id), token varchar,
		created timestamp DEFAULT now(), purpose varchar DEFAULT 'gui', UNIQUE(user_id, token)
	)`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO web_users(id,login,group_id,webrtc_lib,ws,verto_ws,stun,key,lang,avatar,avatar_format,enabled) VALUES (1,'admin',1,'sipjs','','','','hash',0,'','',true); INSERT INTO web_users_tokens(user_id,token,purpose) VALUES (1,'legacy-secret','gui')`)
	if err != nil {
		t.Fatal(err)
	}
	migrateWebUserTokens(db)
	user, err := GetWebUserByToken("legacy-secret")
	if err != nil {
		t.Fatal(err)
	}
	if user == nil || user.Login != "admin" {
		t.Fatalf("legacy token did not authenticate: %+v", user)
	}
	var raw *string
	var hash string
	if err = db.QueryRow(`SELECT token, token_hash FROM web_users_tokens WHERE user_id=1`).Scan(&raw, &hash); err != nil {
		t.Fatal(err)
	}
	if raw != nil || hash != HashToken("legacy-secret") {
		t.Fatal("legacy token was not safely migrated")
	}
}
