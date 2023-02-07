package db

import (
	"custompbx/mainStruct"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
	"strconv"
	"time"
)

const (
	ConstEqual    = "="
	ConstMore     = ">"
	ConstLess     = "<"
	ConstLike     = "LIKE"
	ConstTotalSql = "COUNT(*) OVER() AS total"
	ConstTotal    = "total"
	ConstMoreOrEq = ">="
	ConstLessOrEq = "<="
	ConstNotEqual = "!="
	ConstNotLike  = "NOT LIKE"
)

func InitLogDB() {
	createLogTable(db)
	createLogTableIndex(db)
	go dbCron()
}

func createLogTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS freeswitch_logs(
		created TIMESTAMP DEFAULT now(),
		log_file VARCHAR,
		log_func VARCHAR,
		log_line INTEGER,
		log_level INTEGER,
		text_channel INTEGER,
		user_data VARCHAR,
		body VARCHAR,
		instance_id bigint NOT NULL REFERENCES fs_instances (id) ON DELETE CASCADE
	)
	WITH (OIDS=FALSE);`,
	)
	panicErr(err)
}

func createLogTableIndex(db *sql.DB) {
	_, err := db.Exec(`
	CREATE INDEX IF NOT EXISTS x_log_created ON freeswitch_logs(
		created
	);`,
	)
	panicErr(err)
}

func SetLogLine(logLine mainStruct.LogType) error {
	_, err := db.Exec(
		`INSERT INTO freeswitch_logs(log_file, log_func, log_line, log_level, text_channel, user_data, body) VALUES($1, $2, $3, $4, $5, $6, $7);`,
		logLine.LogFile, logLine.LogFunc, logLine.LogLine, logLine.LogLevel, logLine.TextChannel, logLine.UserData, logLine.Body)
	if err != nil {
		log.Printf("%+v", err.Error())
		return err
	}
	return nil
}

func SetLogLines(logLines []mainStruct.LogType, instanceId int64) error {
	if len(logLines) == 0 {
		return nil
	}
	query := "INSERT INTO freeswitch_logs(log_file, log_func, log_line, log_level, text_channel, user_data, body, instance_id) VALUES"
	var values []interface{}
	for i, logLine := range logLines {
		values = append(values, logLine.LogFile, logLine.LogFunc, logLine.LogLine, logLine.LogLevel, logLine.TextChannel, logLine.UserData, logLine.Body, instanceId)

		numFields := 8
		n := i * numFields

		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1]
	_, err := db.Exec(query, values...)
	if err != nil {
		log.Printf("%+v", err.Error())
		return err
	}
	return nil
}

func GetLogs() ([]mainStruct.LogType, error) {
	logs, err := db.Query(
		`SELECT
					created,
					log_file,
					log_func,
					log_line,
					log_level,
					text_channel,
					user_data,
					body
				FROM freeswitch_logs
				ORDER BY created DESC
				LIMIT 1000
					`,
	)

	defer logs.Close()
	if err != nil {
		log.Printf("%+v", err.Error())
		return nil, err
	}
	var logLines []mainStruct.LogType
	for logs.Next() {
		logLine := mainStruct.LogType{}
		err := logs.Scan(&logLine.Created, &logLine.LogFile, &logLine.LogFunc, &logLine.LogLine, &logLine.LogLevel, &logLine.TextChannel, &logLine.UserData, &logLine.Body)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		logLines = append(logLines, logLine)
	}
	return logLines, nil
}

func clearOldLogs() {
	_, err := db.Exec(
		`DELETE from freeswitch_logs
               WHERE created < now() - INTERVAL '1 day'`)
	if err != nil {
		log.Printf("clearOldLogs ERROR: %+v", err.Error())
	}
}

func dbCron() {
	everyHourTick := time.Tick(1 * time.Hour)

	for {
		select {
		case <-everyHourTick:
			clearOldLogs()
			// at hepDb.go
			clearOldHEPs()
		}
	}
}

func GetList(limit, offset int, filters []mainStruct.Filter, order mainStruct.Order, instanceId int64) ([]map[string]*interface{}, error) {
	table := "freeswitch_logs"

	var fieldsArr = []string{
		"created",
		"log_file",
		"log_func",
		"log_line",
		"log_level",
		"text_channel",
		"user_data",
		"body",
	}

	fieldsArr = append(fieldsArr, ConstTotalSql)
	fieldsArrModified := make([]string, len(fieldsArr))
	copy(fieldsArrModified, fieldsArr)
	fieldsArrModified[0] = fmt.Sprintf("to_char(%s, 'YYYY-MM-DD HH24:MI:SS.MS')", fieldsArrModified[0])

	queryBuilder := squirrel.Select(fieldsArrModified...).From(table).Limit(uint64(limit)).Offset(uint64(offset)).PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{"instance_id": instanceId})
	operandsMap := map[string]bool{ConstEqual: true, ConstMore: true, ConstLess: true, ConstLike: true, ConstMoreOrEq: true, ConstLessOrEq: true, ConstNotEqual: true, ConstNotLike: true}
	for _, filter := range filters {
		if !operandsMap[filter.Operand] {
			log.Println(filter.Operand)
			continue
		}
		var found bool
		for _, field := range fieldsArr {
			if filter.Field == field {
				found = true
				if filter.Field == fieldsArrModified[0] {
					filter.Field = fmt.Sprintf("%s::TEXT')", filter.Field)
				}
				break
			} else if filter.Field == ConstTotal {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		switch filter.Operand {
		case ConstEqual:
			queryBuilder = queryBuilder.Where(squirrel.Eq{filter.Field: filter.FieldValue})
		case ConstMore:
			queryBuilder = queryBuilder.Where(squirrel.Gt{filter.Field: filter.FieldValue})
		case ConstLess:
			queryBuilder = queryBuilder.Where(squirrel.Lt{filter.Field: filter.FieldValue})
		case ConstLike:
			queryBuilder = queryBuilder.Where(squirrel.Like{filter.Field: filter.FieldValue})
		case ConstMoreOrEq:
			queryBuilder = queryBuilder.Where(squirrel.GtOrEq{filter.Field: filter.FieldValue})
		case ConstLessOrEq:
			queryBuilder = queryBuilder.Where(squirrel.LtOrEq{filter.Field: filter.FieldValue})
		case ConstNotEqual:
			queryBuilder = queryBuilder.Where(squirrel.NotEq{filter.Field: filter.FieldValue})
		case ConstNotLike:
			queryBuilder = queryBuilder.Where(squirrel.NotLike{filter.Field: filter.FieldValue})
		default:
			continue
		}
	}
	if len(order.Fields) > 0 {
		if order.Desc {
			lastOne := order.Fields[len(order.Fields)-1]
			lastOne = lastOne + " DESC"
			order.Fields[len(order.Fields)-1] = lastOne
		}
		queryBuilder = queryBuilder.OrderBy(order.Fields...)
	} else {
		queryBuilder = queryBuilder.OrderBy("created ASC")
	}
	query, args, _ := queryBuilder.ToSql()
	log.Println(query)
	log.Println(args)
	log.Println(args)
	logs, err := db.Query(query, args...)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error on select")
	}
	defer logs.Close()

	var logLines []map[string]*interface{}

	for logs.Next() {
		var rows []interface{}
		result := make(map[string]*interface{})
		for _, name := range fieldsArr {
			var str interface{}
			rows = append(rows, &str)
			if name == ConstTotalSql {
				result[ConstTotal] = &str
				continue
			}
			result[name] = &str
		}
		err := logs.Scan(rows...)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		logLines = append(logLines, result)
	}
	return logLines, nil
}

func GetHEPList(limit, offset int, filters []mainStruct.Filter, order mainStruct.Order) ([]map[string]*interface{}, error) {
	table := "hep_packets"

	var fieldsArr = []string{
		"hep_timestamp",
		"sip_first_method",
		"sip_call_id",
		//"hep_sid",
		//"hep_cid",
		"hep_dst_ip",
		"hep_dst_port",
		"hep_src_ip",
		"hep_src_port",
		"sip_from_host",
		"sip_from_user",
		//"sip_to_user",
		//"sip_to_host",
		"sip_uri_user",
		"sip_uri_host",
		"sip_user_agent",
		"hep_node_id",
	}

	//fieldsArr = append(fieldsArr, ConstTotalSql)
	//fieldsArr = append(fieldsArr, "(select count(sip_first_method) from hep_packets  WHERE sip_first_method IN ('INVITE', 'ACK', 'BYE', 'CANCEL', 'REGISTER', 'OPTIONS', 'PRACK', 'SUBSCRIBE', 'NOTIFY', 'PUBLISH', 'INFO', 'REFER', 'MESSAGE', 'UPDATE')) AS total")

	fieldsArrModified := make([]string, len(fieldsArr))
	copy(fieldsArrModified, fieldsArr)
	fieldsArrModified[0] = fmt.Sprintf("to_char(%s, 'YYYY-MM-DD HH24:MI:SS.MS')", fieldsArrModified[0])

	queryBuilder := squirrel.Select(fieldsArrModified...).From(table).Limit(uint64(limit)).Offset(uint64(offset)).PlaceholderFormat(squirrel.Dollar)
	countQueryBuilder := squirrel.Select("count(sip_first_method)").From(table).PlaceholderFormat(squirrel.Dollar)

	//queryBuilder.Column("(select count(sip_first_method) from hep_packets  WHERE sip_first_method IN ('INVITE', 'ACK', 'BYE', 'CANCEL', 'REGISTER', 'OPTIONS', 'PRACK', 'SUBSCRIBE', 'NOTIFY', 'PUBLISH', 'INFO', 'REFER', 'MESSAGE', 'UPDATE')) AS total")
	operandsMap := map[string]bool{ConstEqual: true, ConstMore: true, ConstLess: true, ConstLike: true, ConstMoreOrEq: true, ConstLessOrEq: true, ConstNotEqual: true, ConstNotLike: true}
	for _, filter := range filters {
		if !operandsMap[filter.Operand] {
			log.Println(filter.Operand)
			continue
		}
		var found bool
		for _, field := range fieldsArr {
			if filter.Field == field {
				found = true
				if filter.Field == fieldsArrModified[0] {
					filter.Field = fmt.Sprintf("%s::TEXT')", filter.Field)
				}
				break
			} else if filter.Field == ConstTotal {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		switch filter.Operand {
		case ConstEqual:
			queryBuilder = queryBuilder.Where(squirrel.Eq{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.Eq{filter.Field: filter.FieldValue})
		case ConstMore:
			queryBuilder = queryBuilder.Where(squirrel.Gt{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.Gt{filter.Field: filter.FieldValue})
		case ConstLess:
			queryBuilder = queryBuilder.Where(squirrel.Lt{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.Lt{filter.Field: filter.FieldValue})
		case ConstLike:
			queryBuilder = queryBuilder.Where(squirrel.Like{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.Like{filter.Field: filter.FieldValue})
		case ConstMoreOrEq:
			queryBuilder = queryBuilder.Where(squirrel.GtOrEq{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.GtOrEq{filter.Field: filter.FieldValue})
		case ConstLessOrEq:
			queryBuilder = queryBuilder.Where(squirrel.LtOrEq{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.LtOrEq{filter.Field: filter.FieldValue})
		case ConstNotEqual:
			queryBuilder = queryBuilder.Where(squirrel.NotEq{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.NotEq{filter.Field: filter.FieldValue})
		case ConstNotLike:
			queryBuilder = queryBuilder.Where(squirrel.NotLike{filter.Field: filter.FieldValue})
			countQueryBuilder = countQueryBuilder.Where(squirrel.NotLike{filter.Field: filter.FieldValue})
		default:
			continue
		}
	}

	//queryBuilder = queryBuilder.Where("sip_first_method !~ '^\\d+$'")
	queryBuilder = queryBuilder.Where("sip_first_method IN ('INVITE', 'ACK', 'BYE', 'CANCEL', 'REGISTER', 'OPTIONS', 'PRACK', 'SUBSCRIBE', 'NOTIFY', 'PUBLISH', 'INFO', 'REFER', 'MESSAGE', 'UPDATE')")
	countQueryBuilder = countQueryBuilder.Where("sip_first_method IN ('INVITE', 'ACK', 'BYE', 'CANCEL', 'REGISTER', 'OPTIONS', 'PRACK', 'SUBSCRIBE', 'NOTIFY', 'PUBLISH', 'INFO', 'REFER', 'MESSAGE', 'UPDATE')")

	countQuery, _, _ := countQueryBuilder.ToSql()
	queryBuilder = queryBuilder.Column("(" + countQuery + ") AS total")

	if len(order.Fields) > 0 {
		if order.Desc {
			lastOne := order.Fields[len(order.Fields)-1]
			lastOne = lastOne + " DESC"
			order.Fields[len(order.Fields)-1] = lastOne
		}
		queryBuilder = queryBuilder.OrderBy(order.Fields...)
	} else {
		queryBuilder = queryBuilder.OrderBy("hep_timestamp ASC")
	}
	query, args, _ := queryBuilder.ToSql()
	log.Println(query)
	log.Println(args)
	logs, err := db.Query(query, args...)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error on select")
	}
	defer logs.Close()

	var logLines []map[string]*interface{}

	for logs.Next() {
		var rows []interface{}
		result := make(map[string]*interface{})
		for _, name := range fieldsArr {
			var str interface{}
			rows = append(rows, &str)
			/*			if name == ConstTotalSql {
						result[ConstTotal] = &str
						continue
					}*/
			result[name] = &str
		}
		var str interface{}
		result[ConstTotal] = &str
		rows = append(rows, &str)
		err := logs.Scan(rows...)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		logLines = append(logLines, result)
	}
	return logLines, nil
}
