package db

import (
	"custompbx/cfg"
	"custompbx/daemonCache"
	"custompbx/mainStruct"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

var db *sql.DB

func DBConnect() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s fallback_application_name=%s sslmode=disable",
		cfg.CustomPbx.Db.User, cfg.CustomPbx.Db.Pass, cfg.CustomPbx.Db.Name, cfg.CustomPbx.Db.Host, strconv.Itoa(cfg.CustomPbx.Db.Port), cfg.AppName)
	db, err = sql.Open("postgres", dbinfo)
	if err != nil || db == nil || db.Ping() != nil {
		daemonCache.State.DatabaseConnection = false
		daemonCache.State.DataBaseError = err
		return
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	daemonCache.State.DatabaseConnection = true
}

func StartDB() {
	DBConnect()
	for {
		if daemonCache.State.DatabaseConnection {
			return
		}
		log.Println("Failed to connect to the database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		DBConnect()
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitRootDB() {
	createInstancesTable(db)
}

func createInstancesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS fs_instances(
		id serial NOT NULL PRIMARY KEY,
		instance_name VARCHAR UNIQUE,
		host VARCHAR,
		port INTEGER,
		auth VARCHAR,
		token VARCHAR,
		description VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE 
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func GetFsInstances(fsInstances *mainStruct.FsInstances) {
	instances, err := db.Query(`SELECT id, instance_name, host, port, auth, token, description, enabled FROM fs_instances;`)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer instances.Close()
	for instances.Next() {
		var instance mainStruct.FsInstance
		err := instances.Scan(&instance.Id, &instance.Name, &instance.Host, &instance.Port, &instance.Auth, &instance.Token, &instance.Description, &instance.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		fsInstances.Set(&instance)
	}
}

func SetFSInstance(switchName string) (int64, error) {
	var instanceId int64
	err := db.QueryRow("INSERT INTO fs_instances(instance_name, host, port, auth, token, description) VALUES($1, $2, $3, 'blank', 'token', 'new instance') RETURNING id",
		switchName,
		cfg.CustomPbx.Web.Host,
		cfg.CustomPbx.Web.Port,
	).Scan(&instanceId)
	if err != nil {
		return instanceId, err
	}
	return instanceId, nil
}

func UpdateFSInstanceDescription(id int64, description string) error {
	_, err := db.Exec("UPDATE fs_instances SET description = $1 WHERE id = $2",
		description,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
