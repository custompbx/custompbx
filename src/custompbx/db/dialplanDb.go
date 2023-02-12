package db

import (
	"custompbx/mainStruct"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

func InitDialplanDB() {
	createDialplanContextsTable(db)
	createDialplanExtensionsTable(db)
	createDialplanConditionsTable(db)
	createDialplanRegexexTable(db)
	createDialplanActionsTable(db)
	createDialplanAntiActionsTable(db)
	createDialplanSettingsTable(db)

}

func createDialplanSettingsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_settings(
		id serial NOT NULL PRIMARY KEY,
		name VARCHAR,
		value VARCHAR,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE(name, instance_id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDialplanContextsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_contexts(
		id serial NOT NULL PRIMARY KEY,
		name VARCHAR,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE(name, instance_id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDialplanExtensionsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_extensions(
		id serial NOT NULL PRIMARY KEY,
		position integer NOT NULL,
		context_id bigint NOT NULL REFERENCES dialplan_contexts (id) ON DELETE CASCADE,
		name VARCHAR,
		continue VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (context_id, name),
		UNIQUE (context_id, position)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDialplanConditionsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_conditions(
		id serial NOT NULL PRIMARY KEY,
		position integer NOT NULL,
		extension_id bigint NOT NULL REFERENCES dialplan_extensions (id) ON DELETE CASCADE,
		break VARCHAR,
		field VARCHAR,
		expression VARCHAR,
		hour VARCHAR,
		mday VARCHAR,
		mon VARCHAR,
		mweek VARCHAR,
		wday VARCHAR,
		date_time  VARCHAR,
		time_of_day VARCHAR,
		year VARCHAR,
		minute VARCHAR,
		week VARCHAR,
		yday VARCHAR,
		minday VARCHAR,
		tz_offset VARCHAR,
		dst VARCHAR,
		regex VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDialplanActionsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_actions(
		id serial NOT NULL PRIMARY KEY,
		position integer NOT NULL,
		condition_id bigint NOT NULL REFERENCES dialplan_conditions (id) ON DELETE CASCADE,
		application VARCHAR,
		data VARCHAR,
		inline BOOLEAN DEFAULT FALSE,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (id, position)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDialplanRegexexTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_regexes(
		id serial NOT NULL PRIMARY KEY,
		condition_id bigint NOT NULL REFERENCES dialplan_conditions (id) ON DELETE CASCADE,
		field VARCHAR,
		expression VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (id)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func createDialplanAntiActionsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dialplan_antiactions(
		id serial NOT NULL PRIMARY KEY,
		position integer NOT NULL,
		condition_id bigint NOT NULL REFERENCES dialplan_conditions (id) ON DELETE CASCADE,
		application VARCHAR,
		data VARCHAR,
		enabled BOOLEAN NOT NULL DEFAULT TRUE,
		UNIQUE (condition_id, position)
	)
	WITH (OIDS=FALSE);
	`)
	panicErr(err)
}

func SetContext(name string, instanceId int64) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO dialplan_contexts(name, instance_id) values($1, $2) returning id;", name, instanceId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetContextExtension(contextId int64, name, extContinue string) (int64, int64, error) {
	var id int64
	var position int64
	err := db.QueryRow(`INSERT INTO dialplan_extensions(context_id, name, continue, position)
							values($1, $2, $3, (SELECT COALESCE((SELECT position + 1 FROM dialplan_extensions WHERE context_id = $1 ORDER BY position DESC LIMIT 1), 1))) returning id, position;`,
		contextId, name, extContinue).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

func SetExtensionCondition(
	extensionId int64,
	conditionBreak,
	field,
	expression,
	hour,
	mday,
	mon,
	mweek,
	wday,
	dateTime,
	timeOfDay,
	year,
	minute,
	week,
	yday,
	minday,
	tzOffset,
	dst,
	regex string,
) (int64, int64, error) {
	var id int64
	var position int64
	err := db.QueryRow(`INSERT INTO dialplan_conditions(extension_id, break, field, expression, hour, mday, mon, mweek, wday,
		date_time, time_of_day, year, minute, week, yday, minday, tz_offset, dst, regex, position)
							values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
							       (SELECT COALESCE((SELECT position + 1 FROM dialplan_conditions WHERE extension_id = $1 ORDER BY position DESC LIMIT 1), 1))) returning id, position;`,
		extensionId, conditionBreak, field, expression, hour, mday, mon, mweek, wday, dateTime,
		timeOfDay, year, minute, week, yday, minday, tzOffset, dst, regex).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

func SetConditionRegex(extensionId int64, field, expr string) (int64, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO dialplan_regexes(condition_id, field, expression)
							values($1, $2, $3) returning id;`,
		extensionId, field, expr).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func SetConditionAction(extensionId int64, application, data string, inline bool) (int64, int64, error) {
	var id int64
	var position int64
	err := db.QueryRow(`INSERT INTO dialplan_actions(condition_id, application, data, inline, position)
							values($1, $2, $3, $4, (SELECT COALESCE((SELECT position + 1 FROM dialplan_actions WHERE condition_id = $1 ORDER BY position DESC LIMIT 1), 1))) returning id, position;`,
		extensionId, application, data, inline).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

func SetConditionAntiAction(extensionId int64, application, data string) (int64, int64, error) {
	var id int64
	var position int64
	err := db.QueryRow(`INSERT INTO dialplan_antiactions(condition_id, application, data, position)
							values($1, $2, $3, (SELECT COALESCE((SELECT position + 1 FROM dialplan_antiactions WHERE condition_id = $1 ORDER BY position DESC LIMIT 1), 1))) returning id, position;`,
		extensionId, application, data).Scan(&id, &position)
	if err != nil {
		return 0, 0, err
	}
	return id, position, err
}

// GET
func GetDialplanSettings(dialplan *mainStruct.Dialplans, instanceId int64) {
	contexts, err := db.Query(`SELECT name, value, enabled FROM dialplan_settings WHERE instance_id = $1;`, instanceId)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer contexts.Close()
	for contexts.Next() {
		var name string
		var value string
		var enabled bool
		err := contexts.Scan(&name, &value, &enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		if !enabled {
			continue
		}
		switch name {
		case mainStruct.NoProceedName:
			dialplan.NoProceed = false
			if value == "true" {
				dialplan.NoProceed = true
			}
		}
	}
}

func GetContexts(dialplan *mainStruct.Dialplans, instanceId int64) {
	contexts, err := db.Query(`SELECT id, name, enabled FROM dialplan_contexts WHERE instance_id = $1;`, instanceId)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer contexts.Close()
	for contexts.Next() {
		var context mainStruct.Context
		err := contexts.Scan(&context.Id, &context.Name, &context.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		context.Dialplan = dialplan
		dialplan.Contexts.Set(&context)
	}
}

func GetContextExtensions(context *mainStruct.Context, dialplan *mainStruct.Dialplans) {
	extens, err := db.Query(`SELECT id, name, continue, position, enabled FROM dialplan_extensions WHERE context_id = $1;`, context.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer extens.Close()
	for extens.Next() {
		var extension mainStruct.Extension
		err := extens.Scan(&extension.Id, &extension.Name, &extension.Continue, &extension.Position, &extension.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		extension.Context = context
		context.Extensions.Set(&extension)
		dialplan.Extensions.Set(&extension)
	}
}

func GetExtensionConditions(extension *mainStruct.Extension, dialplan *mainStruct.Dialplans) {
	conditions, err := db.Query(`SELECT 
       id, break, field, expression, hour, mday, mon, mweek, wday, position, date_time, time_of_day, year, minute, week, yday, minday, tz_offset, dst, regex, enabled
       FROM dialplan_conditions WHERE extension_id = $1;`, extension.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	defer conditions.Close()
	for conditions.Next() {
		var condition mainStruct.Condition
		err := conditions.Scan(
			&condition.Id,
			&condition.Break,
			&condition.Field,
			&condition.Expression,
			&condition.Hour,
			&condition.Mday,
			&condition.Mon,
			&condition.Mweek,
			&condition.Wday,
			&condition.Position,
			&condition.DateTime,
			&condition.TimeOfDay,
			&condition.Year,
			&condition.Minute,
			&condition.Week,
			&condition.Yday,
			&condition.Minday,
			&condition.TzOffset,
			&condition.Dst,
			&condition.Regex,
			&condition.Enabled,
		)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		condition.Extension = extension
		extension.Conditions.Set(&condition)
		dialplan.Conditions.Set(&condition)
	}
}

func GetConditionRegexes(condition *mainStruct.Condition, dialplan *mainStruct.Dialplans) {
	regexes, err := db.Query(`SELECT id, field, expression, enabled FROM dialplan_regexes WHERE condition_id = $1;`, condition.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer regexes.Close()
	for regexes.Next() {
		var regex mainStruct.Regex
		err := regexes.Scan(&regex.Id, &regex.Field, &regex.Expression, &regex.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		regex.Condition = condition
		condition.Regexes.Set(&regex)
		dialplan.Regexes.Set(&regex)
	}
}

func GetConditionActions(condition *mainStruct.Condition, dialplan *mainStruct.Dialplans) {
	actions, err := db.Query(`SELECT id, application, data, inline, position, enabled FROM dialplan_actions WHERE condition_id = $1;`, condition.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer actions.Close()
	for actions.Next() {
		var action mainStruct.Action
		err := actions.Scan(&action.Id, &action.Application, &action.Data, &action.Inline, &action.Position, &action.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		action.Condition = condition
		condition.Actions.Set(&action)
		dialplan.Actions.Set(&action)
	}
}

func GetConditionAntiActions(condition *mainStruct.Condition, dialplan *mainStruct.Dialplans) {
	antiActions, err := db.Query(`SELECT id, application, data, position, enabled FROM dialplan_antiactions WHERE condition_id = $1;`, condition.Id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer antiActions.Close()
	for antiActions.Next() {
		var antiAction mainStruct.AntiAction
		err := antiActions.Scan(&antiAction.Id, &antiAction.Application, &antiAction.Data, &antiAction.Position, &antiAction.Enabled)
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		antiAction.Condition = condition
		condition.AntiActions.Set(&antiAction)
		dialplan.AntiActions.Set(&antiAction)
	}
}

func MoveContextExtension(extension *mainStruct.Extension, newPosition int64) error {
	pos1 := newPosition
	pos2 := extension.Position
	pos3 := newPosition + 1

	if extension.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE dialplan_extensions SET position = (position + 1)*-1 WHERE context_id = $1 AND position > $2`,
		extension.Context.Id, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_extensions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_extensions SET position = $2 WHERE id = $1`,
		extension.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_extensions SET position = (position - 1)*-1 WHERE context_id = $1 AND position > $2`,
		extension.Context.Id, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_extensions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	extensions := extension.Context.Extensions.Props()
	switch extension.Position > newPosition {
	case true:
		for _, v := range extensions {
			if v.Position >= newPosition && v.Position < extension.Position {
				v.Position = v.Position + 1
			}
		}
	case false:
		for _, v := range extensions {
			if v.Position > extension.Position && v.Position <= newPosition {
				v.Position = v.Position - 1
			}
		}
	}
	extension.Position = newPosition

	return err
}

func MoveContextCondition(condition *mainStruct.Condition, newPosition int64) error {
	pos1 := newPosition
	pos2 := condition.Position
	pos3 := newPosition + 1

	if condition.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE dialplan_conditions SET position = (position + 1)*-1 WHERE extension_id = $1 AND position > $2`,
		condition.Extension.Id, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_conditions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_conditions SET position = $2 WHERE id = $1`,
		condition.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_conditions SET position = (position - 1)*-1 WHERE extension_id = $1 AND position > $2`,
		condition.Extension.Id, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_conditions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	conditions := condition.Extension.Conditions.Props()
	switch condition.Position > newPosition {
	case true:
		for _, v := range conditions {
			if v.Position >= newPosition && v.Position < condition.Position {
				v.Position = v.Position + 1
			}
		}
	case false:
		for _, v := range conditions {
			if v.Position > condition.Position && v.Position <= newPosition {
				v.Position = v.Position - 1
			}
		}
	}
	condition.Position = newPosition

	return err
}

func MoveContextAction(action *mainStruct.Action, newPosition int64) error {
	pos1 := newPosition
	pos2 := action.Position
	pos3 := newPosition + 1

	if action.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE dialplan_actions SET position = (position + 1)*-1 WHERE condition_id = $1 AND position > $2`,
		action.Condition.Id, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_actions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_actions SET position = $2 WHERE id = $1`,
		action.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_actions SET position = (position - 1)*-1 WHERE condition_id = $1 AND position > $2`,
		action.Condition.Id, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_actions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	actions := action.Condition.Actions.Props()
	switch action.Position > newPosition {
	case true:
		for _, v := range actions {
			if v.Position >= newPosition && v.Position < action.Position {
				v.Position = v.Position + 1
			}
		}
	case false:
		for _, v := range actions {
			if v.Position > action.Position && v.Position <= newPosition {
				v.Position = v.Position - 1
			}
		}
	}
	action.Position = newPosition

	return err
}

func MoveContextAntiAction(antiAction *mainStruct.AntiAction, newPosition int64) error {
	pos1 := newPosition
	pos2 := antiAction.Position
	pos3 := newPosition + 1

	if antiAction.Position > newPosition {
		pos1 = newPosition - 1
		pos3 = newPosition
	}
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	_, err = tr.Exec(`UPDATE dialplan_antiactions SET position = (position + 1)*-1 WHERE condition_id = $1 AND position > $2`,
		antiAction.Condition.Id, pos1)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_antiactions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_antiactions SET position = $2 WHERE id = $1`,
		antiAction.Id, pos3)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_antiactions SET position = (position - 1)*-1 WHERE condition_id = $1 AND position > $2`,
		antiAction.Condition.Id, pos2)
	if err != nil {
		return err
	}
	_, err = tr.Exec(`UPDATE dialplan_antiactions SET position = (position)*-1 WHERE position < 0`)
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	antiActions := antiAction.Condition.AntiActions.Props()
	switch antiAction.Position > newPosition {
	case true:
		for _, v := range antiActions {
			if v.Position >= newPosition && v.Position < antiAction.Position {
				v.Position = v.Position + 1
			}
		}
	case false:
		for _, v := range antiActions {
			if v.Position > antiAction.Position && v.Position <= newPosition {
				v.Position = v.Position - 1
			}
		}
	}
	antiAction.Position = newPosition

	return err
}

func DelContextRegex(id int64) bool {
	_, err := db.Exec(`DELETE FROM dialplan_regexes WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func DelContextAction(id int64) bool {
	_, err := db.Exec(`DELETE FROM dialplan_actions WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func DelContextAntiAction(id int64) bool {
	_, err := db.Exec(`DELETE FROM dialplan_antiactions WHERE id = $1;`, id)
	if err != nil {
		return false
	}
	return true
}

func UpdateContextRegex(id int64, field, expression string) bool {
	_, err := db.Exec(`UPDATE dialplan_regexes SET field = $2, expression = $3 WHERE id = $1;`, id, field, expression)
	if err != nil {
		return false
	}
	return true
}

func UpdateContextAction(id int64, application, data string, inline bool) bool {
	_, err := db.Exec(`UPDATE dialplan_actions SET application = $2, data = $3, inline = $4 WHERE id = $1;`, id, application, data, inline)
	if err != nil {
		return false
	}
	return true
}

func UpdateContextAntiAction(id int64, application, data string) bool {
	_, err := db.Exec(`UPDATE dialplan_antiactions SET application = $2, data = $3 WHERE id = $1;`, id, application, data)
	if err != nil {
		return false
	}
	return true
}

func SwitchContextRegex(id int64, switcher bool) bool {
	_, err := db.Exec(`UPDATE dialplan_regexes SET field = $2 WHERE id = $1;`, id, switcher)
	if err != nil {
		return false
	}
	return true
}

func SwitchContextAction(id int64, switcher bool) bool {
	_, err := db.Exec(`UPDATE dialplan_actions SET enabled = $2 WHERE id = $1;`, id, switcher)
	if err != nil {
		return false
	}
	return true
}

func SwitchContextAntiAction(id int64, switcher bool) bool {
	_, err := db.Exec(`UPDATE dialplan_antiactions SET enabled = $2 WHERE id = $1;`, id, switcher)
	if err != nil {
		return false
	}
	return true
}

func RenameContext(id int64, name string) error {
	_, err := db.Exec("UPDATE dialplan_contexts SET name = $2 WHERE id = $1;", id, name)
	return err
}

func RenameExtension(id int64, name string) error {
	_, err := db.Exec("UPDATE dialplan_extensions SET name = $2 WHERE id = $1;", id, name)
	return err
}

func DeleteContext(id int64) error {
	_, err := db.Exec("DELETE FROM dialplan_contexts WHERE id = $1;", id)
	return err
}

func DeleteExtension(id int64) error {
	_, err := db.Exec("DELETE FROM dialplan_extensions WHERE id = $1;", id)
	return err
}

func SwitchExtensionContinue(id int64, extensionContinue string) error {
	if extensionContinue != "" && extensionContinue != "true" {
		return errors.New("wrong continue value")
	}

	_, err := db.Exec("UPDATE dialplan_extensions SET continue = $2 WHERE id = $1;", id, extensionContinue)
	return err
}

func UpdateExtensionCondition(
	conditionId int64, conditionBreak, field, expression, hour, mday, mon, mweek, wday, dateTime, timeOfDay, year, minute,
	week, yday, minday, tzOffset, dst, regex string) error {
	_, err := db.Exec(`
UPDATE dialplan_conditions SET break = $2, field = $3, expression = $4, hour = $5, mday = $6, mon = $7, mweek = $8, 
						   wday = $9, date_time = $10, time_of_day = $11, year = $12, minute = $13, week = $14, 
						   yday = $15, minday = $16, tz_offset = $17, dst = $18, regex = $19
WHERE id = $1;
`,
		conditionId, conditionBreak, field, expression, hour, mday, mon, mweek, wday, dateTime,
		timeOfDay, year, minute, week, yday, minday, tzOffset, dst, regex)
	return err
}

func SwitchContextConditon(id int64, switcher bool) error {
	_, err := db.Exec(`UPDATE dialplan_conditions SET enabled = $2 WHERE id = $1;`, id, switcher)
	return err
}

func DeleteContextConditon(id int64) error {
	_, err := db.Exec(`DELETE FROM dialplan_conditions WHERE id = $1;`, id)
	return err
}

func SetDialplanSettings(dialplan *mainStruct.Dialplans, name, value string) error {
	if name == "" {
		return errors.New("no param name")
	}
	var enabled bool
	err := db.QueryRow(`INSERT INTO dialplan_settings(name, value) VALUES($1, $2) ON CONFLICT(name) DO UPDATE SET value = $2 RETURNING enabled;`,
		name, value,
	).Scan(&enabled)
	if err != nil {
		log.Printf("%+v", err)
		return errors.New("cant' update")
	}
	if !enabled {
		return nil
	}
	switch name {
	case mainStruct.NoProceedName:
		dialplan.NoProceed = false
		if value == "true" {
			dialplan.NoProceed = true
		}
	}
	return nil
}
