package db

import (
	"custompbx/mainStruct"
	"database/sql"
	"log"
)

func InitGlobalVariablesDB() {
	createGlobalVariablesTable(db)
}

func createGlobalVariablesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS global_variables(
		id serial NOT NULL PRIMARY KEY,
		variable_name VARCHAR,
		variable_value VARCHAR,
		variable_type VARCHAR,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE,
		dynamic BOOLEAN NOT NULL DEFAULT FALSE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		position integer NOT NULL,
		UNIQUE(variable_name, instance_id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func SetGlobalVariable(name string, value string, varType string, instanceId int64, dynamic bool) (int64, int64, error) {
	switch varType {
	case "env-set":
	case "exec-set":
	case "stun-set":
	default:
		varType = "set"
	}
	var id int64
	var position int64
	err := db.QueryRow("INSERT INTO global_variables(variable_name, variable_value, variable_type, instance_id, dynamic, position) values($1, $2, $3, $4, $5, (SELECT COALESCE((SELECT position + 1 FROM global_variables WHERE instance_id = $4 ORDER BY position DESC LIMIT 1), 1))) returning id, position;", name, value, varType, instanceId, dynamic).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

func GetGlobalVariables(globalVariables *mainStruct.GlobalVariables, instanceId int64) {
	items, err := db.Query(`SELECT id, variable_name, variable_value, variable_type, enabled, dynamic, position FROM global_variables WHERE instance_id = $1;`, instanceId)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer items.Close()
	for items.Next() {
		var item mainStruct.GlobalVariable
		err := items.Scan(&item.Id, &item.Name, &item.Value, &item.Type, &item.Enabled, &item.Dynamic, &item.Position)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		globalVariables.Set(&item)
	}
}

func UpdateGlobalVariable(domainId int64, newName string, newValue string, varType string) error {
	switch varType {
	case "env-set":
	case "exec-set":
	case "stun-set":
	default:
		varType = "set"
	}
	_, err := db.Exec("UPDATE global_variables SET variable_name = $1, variable_value = $2, variable_type = $3 WHERE id = $4;", newName, newValue, varType, domainId)
	if err != nil {
		return err
	}
	return err
}

func DelGlobalVariable(Id int64) bool {
	_, err := db.Exec(`DELETE FROM global_variables WHERE id = $1;`, Id)
	if err != nil {
		return false
	}
	return true
}

func SwitchGlobalVariable(paramId int64, enabled bool) error {
	_, err := db.Exec("UPDATE global_variables SET enabled = $1 WHERE id = $2;", enabled, paramId)
	return err
}

func MoveGlobalVariable(variables *mainStruct.GlobalVariables, variable *mainStruct.GlobalVariable, newPosition int64, instanceId int64) error {
	pos1 := newPosition
	pos2 := variable.Position
	pos3 := newPosition + 1

	if variable.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE global_variables SET position = (position + 1)*-1 WHERE instance_id = $1 AND position > $2`,
		instanceId, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE global_variables SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE global_variables SET position = $2 WHERE id = $1`,
		variable.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE global_variables SET position = (position - 1)*-1 WHERE instance_id = $1 AND position > $2`,
		instanceId, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE global_variables SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	/*	vars := variables.Props()
		switch variable.Position > newPosition {
		case true:
			for _, v := range vars {
				if v.Position >= newPosition && v.Position < variable.Position {
					v.Position = v.Position + 1
				}
			}
		case false:
			for _, v := range vars {
				if v.Position > variable.Position && v.Position <= newPosition {
					v.Position = v.Position - 1
				}
			}
		}
		variable.Position = newPosition
	*/
	GetGlobalVariables(variables, instanceId)
	return err
}
