package cdrDb

import (
	"context"
	"custompbx/altData"
	"custompbx/altStruct"
	"custompbx/daemonCache"
	"custompbx/fsesl"
	"custompbx/intermediateDB"
	"custompbx/mainStruct"
	"custompbx/webcache"
	"errors"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
)

var db *pgx.Conn

func StartDB() {
	dbInfo := ""
	if daemonCache.State.CdrDatabaseConnection {
		return
	}
	moduleName := webcache.GetWebSetting(webcache.CdrModule)

	var mod *altStruct.ConfigurationsList
	if moduleName != "odbc_cdr" {
		moduleName = "cdr_pg_csv"
		mod, _ = altData.GetModuleByName(mainStruct.ModCdrPgCsv)
		setI, err := intermediateDB.GetByValue(
			&altStruct.ConfigCdrPgCsvSetting{Parent: mod, Enabled: true, Name: "db-info"},
			map[string]bool{"Parent": true, "Enabled": true, "Name": true},
		)
		if err == nil && len(setI) > 0 {
			set, ok := setI[0].(altStruct.ConfigCdrPgCsvSetting)
			if ok {
				dbInfo = fsesl.ParseGlobalVars(set.Value)
			}
		}
	}
	if mod == nil {
		moduleName = "odbc_cdr"
		mod, _ = altData.GetModuleByName(mainStruct.ModOdbcCdr)
		setI, err := intermediateDB.GetByValue(
			&altStruct.ConfigOdbcCdrSetting{Parent: mod, Enabled: true, Name: "odbc-dsn"},
			map[string]bool{"Parent": true, "Enabled": true, "Name": true},
		)
		if err == nil && len(setI) > 0 {
			set, ok := setI[0].(altStruct.ConfigOdbcCdrSetting)
			if ok {
				dbInfo = fsesl.ParseGlobalVars(set.Value)
			}
		}
	}

	if dbInfo == "" {
		daemonCache.State.CdrDatabaseConnection = false
		daemonCache.State.CdrDataBaseError = errors.New("no config for cdr module " + moduleName)
		return
	}
	dbInfo = strings.Replace(dbInfo, "hostaddr=", "host=", 1)
	dbInfo = strings.Replace(dbInfo, "pgsql://", "", 1)

	conn, err := pgx.Connect(context.Background(), dbInfo)
	// if err != nil || conn == nil || conn.Ping() != nil {
	if err != nil || conn == nil {
		daemonCache.State.CdrDatabaseConnection = false
		daemonCache.State.CdrDataBaseError = err
		log.Println(err)

		return
	}
	/*
		conn.SetMaxOpenConns(20)
		conn.SetMaxIdleConns(3)
		conn.SetConnMaxLifetime(time.Hour)*/

	db = conn
	daemonCache.State.CdrDatabaseConnection = true
}
