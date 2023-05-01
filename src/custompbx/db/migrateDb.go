package db

import (
	"context"
	"custompbx/mainStruct"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func InitCustomDB() {
	createCustomSettingsTable(db)
}

func createCustomSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS web_settings(
		param_name VARCHAR NOT NULL,
		param_value VARCHAR NOT NULL DEFAULT '',
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE,
		UNIQUE(param_name, instance_id)
	)
	WITH (OIDS=FALSE);`,
	)
	panicErr(err)
}

func UpdateVersionRequest(instanceId int64, tx *sql.Tx) error {
	var err error
	if instanceId == 0 {
		return errors.New("no instance id")
	}
	if tx == nil {
		return errors.New("no transaction")
	}
	_, err = tx.Exec("INSERT INTO web_settings(param_name, param_value, instance_id) VALUES($1, $2, $3) ON CONFLICT(param_name, instance_id) DO UPDATE SET param_value = $2", mainStruct.CustomPBXVersion, mainStruct.Version, instanceId)
	if err != nil {
		return err
	}

	return err
}

func UpdateVersion(instanceId int64) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = UpdateVersionRequest(instanceId, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetVersion(instanceId int64) string {
	var version string
	var instanceIdExists string
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return ""
	}
	err = tx.QueryRow("SELECT column_name FROM information_schema.columns WHERE table_name='web_settings' and column_name='instance_id'").Scan(&instanceIdExists)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		tx.Rollback()
		return ""
	}
	if instanceIdExists == "" {
		err = tx.QueryRow("SELECT param_value FROM web_settings WHERE param_name = $1", mainStruct.CustomPBXVersion).Scan(&version)
	} else {
		err = tx.QueryRow("SELECT param_value FROM web_settings WHERE param_name = $1 AND instance_id = $2", mainStruct.CustomPBXVersion, instanceId).Scan(&version)
	}
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return ""
	}
	tx.Commit()

	return version
}

func Migrate(switchName string) (bool, error) {
	var err error
	updated := false
	instanceId := getInstanceId(switchName)
	switch GetVersion(instanceId) {
	case "":
		//return updated, nil
		//fallthrough
	case "0.0.1":
		fallthrough
	case "0.0.101":
		log.Println("Updating schema from 0.0.101")
		err = migrateForV0v0v102(instanceId)
		if err != nil {
			return false, err
		}
		updated = true
		fallthrough
	case mainStruct.Version:
		return updated, nil
	}

	err = UpdateVersion(instanceId)
	return updated, err
}

func GetWebSettings(settings *mainStruct.WebSettings, instanceId int64) {
	params, err := db.Query(
		`SELECT param_name, param_value FROM web_settings WHERE instance_id = $1`, instanceId,
	)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer params.Close()
	for params.Next() {
		var name string
		var value string
		err := params.Scan(&name, &value)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		settings.Set(name, value)
	}
}

func AddWebSetting(name, value string, instanceId int64) error {
	_, err := db.Exec(
		`INSERT INTO web_settings(param_name, param_value, instance_id) VALUES($1, $2, $3) ON CONFLICT(param_name, instance_id) DO UPDATE SET param_value = $2`, name, value, instanceId)

	return err
}

func getInstanceId(switchName string) int64 {
	var instanceId int64
	err := db.QueryRow("SELECT id FROM fs_instances WHERE instance_name = $1", switchName).Scan(&instanceId)
	if err != nil {
		return 0
	}

	return instanceId
}

func migrateForV0v0v102(instanceId int64) error {
	if instanceId == 0 {
		return errors.New("no id")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "ALTER TABLE IF EXISTS web_users ADD COLUMN IF NOT EXISTS instance_id BIGINT NOT NULL DEFAULT 0")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, fmt.Sprintf(`
DO $$                  
    BEGIN 
        IF EXISTS
            ( SELECT 1
              FROM   information_schema.tables 
              WHERE  table_schema = 'public'
              AND    table_name = 'web_users'
            )
        THEN
            UPDATE web_users 
            SET instance_id = %d;
        END IF ;
    END
   $$ ;`, instanceId))
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "ALTER TABLE IF EXISTS web_users DROP CONSTRAINT IF EXISTS web_users_ipk")
	_, err = tx.ExecContext(ctx, "ALTER TABLE IF EXISTS web_users ADD CONSTRAINT web_users_ipk FOREIGN KEY (instance_id) REFERENCES fs_instances (id) ON DELETE CASCADE ")
	if err != nil {
		tx.Rollback()
		return err
	}
	err = UpdateVersionRequest(instanceId, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
