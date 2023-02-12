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
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	"log"
	"net"
	"strings"
)

const (
	ConstEqual    = "="
	ConstMore     = ">"
	ConstLess     = "<"
	ConstLike     = "LIKE"
	ConstTotalSql = "COUNT(*) OVER() AS total"
	ConstTotal    = "total"
	ConstMoreOrEq = "=>"
	ConstLessOrEq = "=<"
	ConstNotEqual = "!="
	ConstNotLike  = "NOT LIKE"
)

func GetList(limit, offset int, filters []mainStruct.Filter, order mainStruct.Order) ([]map[string]interface{}, error) {
	if !daemonCache.State.CdrDatabaseConnection || db == nil {
		StartDB()
		if !daemonCache.State.CdrDatabaseConnection {
			return nil, errors.New("no db connection")
		}
	}
	moduleName := webcache.GetWebSetting(webcache.CdrModule)
	var mod *altStruct.ConfigurationsList
	table := ""
	var fieldsArr []string
	if moduleName != "odbc_cdr" {
		mod, _ = altData.GetModuleByName(mainStruct.ModCdrPgCsv)
		setI, err := intermediateDB.GetByValue(
			&altStruct.ConfigCdrPgCsvSetting{Parent: mod, Enabled: true, Name: "db-table"},
			map[string]bool{"Parent": true, "Enabled": true, "Name": true},
		)
		if err == nil && len(setI) > 0 {
			set, ok := setI[0].(altStruct.ConfigCdrPgCsvSetting)
			if ok {
				table = fsesl.ParseGlobalVars(set.Value)
			}
		}
	}
	if mod == nil {
		mod, _ = altData.GetModuleByName(mainStruct.ModOdbcCdr)
		table = webcache.GetWebSetting(webcache.CdrTable)
	}

	if mod == nil {
		return nil, errors.New("no config")
	}

	if table == "" {
		table = "cdr"
	}

	if mod.Module == mainStruct.ModCdrPgCsv {
		list, err := intermediateDB.GetByValue(
			&altStruct.ConfigCdrPgCsvSchema{Parent: mod, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		if err != nil {
			return nil, errors.New("no schema")
		}
		for _, fieldI := range list {
			field, ok := fieldI.(altStruct.ConfigCdrPgCsvSchema)
			if !ok {
				continue
			}
			if field.Column != "" {
				fieldsArr = append(fieldsArr, field.Column)
			} else if field.Var != "" {
				fieldsArr = append(fieldsArr, field.Var)
			}
		}
	} else {
		odbcTableI, err := intermediateDB.GetByValue(
			&altStruct.ConfigOdbcCdrTable{Parent: mod, Enabled: true, Name: table},
			map[string]bool{"Parent": true, "Enabled": true, "Name": true},
		)
		if err != nil || len(odbcTableI) == 0 {
			return nil, errors.New("table not found")
		}
		odbcTable, ok := odbcTableI[0].(altStruct.ConfigOdbcCdrTable)
		if !ok {
			return nil, errors.New("table not found")
		}

		list, err := intermediateDB.GetByValue(
			&altStruct.ConfigOdbcCdrTableField{Parent: &odbcTable, Enabled: true},
			map[string]bool{"Parent": true, "Enabled": true},
		)
		if list == nil {
			return nil, errors.New("no schema")
		}
		for _, fieldI := range list {
			field, ok := fieldI.(altStruct.ConfigOdbcCdrTableField)
			if !ok {
				continue
			}
			if field.Name != "" {
				fieldsArr = append(fieldsArr, field.Name)
			} else if field.ChanVarName != "" {
				fieldsArr = append(fieldsArr, field.ChanVarName)
			}
		}
	}

	if len(fieldsArr) == 0 {
		return nil, errors.New("no columns to select")
	}
	fieldsArr = append(fieldsArr, ConstTotalSql)

	sql := squirrel.Select(fieldsArr...).From(table).Limit(uint64(limit)).Offset(uint64(offset)).PlaceholderFormat(squirrel.Dollar)
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
			sql = sql.Where(squirrel.Eq{filter.Field: filter.FieldValue})
		case ConstMore:
			sql = sql.Where(squirrel.Gt{filter.Field: filter.FieldValue})
		case ConstLess:
			sql = sql.Where(squirrel.Lt{filter.Field: filter.FieldValue})
		case ConstLike:
			sql = sql.Where(squirrel.Like{filter.Field: filter.FieldValue})
		case ConstMoreOrEq:
			sql = sql.Where(squirrel.GtOrEq{filter.Field: filter.FieldValue})
		case ConstLessOrEq:
			sql = sql.Where(squirrel.LtOrEq{filter.Field: filter.FieldValue})
		case ConstNotEqual:
			sql = sql.Where(squirrel.NotEq{filter.Field: filter.FieldValue})
		case ConstNotLike:
			sql = sql.Where(squirrel.NotLike{filter.Field: filter.FieldValue})
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
		sql = sql.OrderBy(order.Fields...)
	}
	query, args, _ := sql.ToSql()
	log.Println(query)
	log.Println(args)
	log.Println(args)
	cdrs, err := db.Query(context.Background(), query, args...)
	if err != nil {
		log.Printf("%+v", err)
		return nil, errors.New("error on select")
	}
	defer cdrs.Close()
	/*colTypes, err := cdrs.ColumnTypes()
	if err!= nil {
		log.Printf("%+v", err)
		return nil, errors.New("error on select")
	}
	var uuidToConvert []string
	var inetToConvert []string
	var enumToConvert []string
	for i, s := range colTypes {
		columnType := s.DatabaseTypeName()
		switch columnType{
		case "UUID":
			uuidToConvert = append(uuidToConvert, fieldsArr[i])
		case "INET":
			inetToConvert = append(inetToConvert, fieldsArr[i])
		case "ENUM":
			enumToConvert = append(enumToConvert, fieldsArr[i])
		}
	}
	*/
	var total []map[string]interface{}

	cdrRecordColumn := webcache.GetWebSetting(webcache.CdrFileServeColumn)
	cdrRecordPath := webcache.GetWebSetting(webcache.CdrFileServerPath)
	for cdrs.Next() {
		//var rows []interface{}
		fieldDescriptions := cdrs.FieldDescriptions()
		result := make(map[string]interface{})
		res, err := cdrs.Values()
		if err != nil {
			log.Printf("%+v", err)
			continue
		}
		for i, desc := range fieldDescriptions {

			//var str = reflect.New(reflect.TypeOf(reflect.Interface))
			//rows = append(rows, &str)
			/*			if string(desc.Name) == ConstTotalSql {
						result[ConstTotal] = res[i]
						continue
					}*/

			switch desc.DataTypeOID {
			case pgtype.UUIDOID:
				re, ok := res[i].([16]uint8)
				if !ok {
					result[string(desc.Name)] = nil
					continue
				}
				ress, err := uuid.FromBytes(re[:])
				if err != nil {
					log.Println(err)
					result[string(desc.Name)] = nil
					continue
				}
				result[string(desc.Name)] = ress.String()
			case pgtype.InetOID:
				re, ok := res[i].(*net.IPNet)
				if !ok {
					result[string(desc.Name)] = nil
					continue
				}
				result[string(desc.Name)] = re.String()
			default:
				result[string(desc.Name)] = res[i]
			}

		}
		/*err := cdrs.Scan(rows...)
		if err != nil {
			log.Printf("%+v", err)
			continue
		}*/
		/*
			for _, s := range uuidToConvert {
				val, ok := result[s]
				if !ok {
					continue
				}
				var v interface{}
				var rawValue = *(val.(*interface{}))
				b, ok := rawValue.([]byte)
				log.Println(s, b)
				if ok {
					v = b
					log.Println(v, err)
				} else {
					v = val
				}
				result[s] = v
			}
		*/
		replaceField, ok := result[cdrRecordColumn]
		if ok {
			var intFace interface{}
			intFace = strings.Replace(fmt.Sprintf("%v", replaceField), cdrRecordPath, "./cdr/records/", 1)
			if replaceField == intFace {
				intFace = "./cdr/records/" + intFace.(string)
			}
			result[cdrRecordColumn] = &intFace
		}
		total = append(total, result)
	}
	return total, nil
}
