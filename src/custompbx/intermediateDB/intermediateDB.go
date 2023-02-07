package intermediateDB

import (
	"custompbx/db"
	"errors"
	"github.com/custompbx/customorm"
	"log"
	"strings"
)

func GetCORM() *customorm.CORM {
	return customorm.Init(db.GetDB())
}

// new staff

func GetAllFromDBAsMap(item interface{}) (interface{}, error) {
	corm := GetCORM()
	return corm.GetDataAll(item, true)
}

func GetAllFromDBAsSlice(item interface{}) ([]interface{}, error) {
	corm := GetCORM()
	res, err := corm.GetDataAll(item, false)
	if err != nil {
		return nil, err
	}
	result, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("nothing to return")
	}
	return result, nil
}

func GetByIdFromDB(item interface{}) (interface{}, error) {
	corm := GetCORM()
	return corm.GetDataById(item, 0)
}

func GetByIdArg(item interface{}, id int64) (interface{}, error) {
	corm := GetCORM()
	return corm.GetDataById(item, id)
}

func GetByValue(item interface{}, fieldNames map[string]bool) ([]interface{}, error) {
	filter := map[string]customorm.FilterFields{}
	for k, v := range fieldNames {
		filter[k] = customorm.FilterFields{Flag: v}
	}

	return GetByValues(item, filter)
}

func GetByValues(item interface{}, filter map[string]customorm.FilterFields) ([]interface{}, error) {
	corm := GetCORM()
	filterStr := customorm.Filters{Fields: filter}
	res, err := corm.GetDataByValue(item, filterStr, false)
	if err != nil {
		return nil, err
	}
	result, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("nothing to return")
	}
	return result, nil
}

func GetByFilteredValues(item interface{}, filter customorm.Filters) ([]interface{}, error) {
	corm := GetCORM()
	res, err := corm.GetDataByValue(item, filter, false)
	if err != nil {
		return nil, err
	}
	result, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("nothing to return")
	}
	return result, nil
}

func GetByValueAsMap(item interface{}, fieldNames map[string]bool) (map[int64]interface{}, error) {
	filter := map[string]customorm.FilterFields{}
	for k, v := range fieldNames {
		filter[k] = customorm.FilterFields{Flag: v}
	}
	return GetByValuesAsMap(item, filter)
}

func GetByValuesAsMap(item interface{}, filter map[string]customorm.FilterFields) (map[int64]interface{}, error) {
	corm := GetCORM()
	filterStr := customorm.Filters{Fields: filter}
	res, err := corm.GetDataByValue(item, filterStr, true)
	if err != nil {
		return nil, err
	}
	mapItem, ok := res.(map[int64]interface{})
	if !ok {
		//return nil, errors.New("no items")
		return nil, nil
	}
	return mapItem, err
}

func DeleteById(item interface{}) error {
	corm := GetCORM()
	return corm.DeleteRowById(item)
}

func DeleteByIdParam(item interface{}, id int64) error {
	corm := GetCORM()
	return corm.DeleteRowByArgId(item, id)
}

func DeleteRows(item interface{}, fields map[string]bool) error {
	corm := GetCORM()
	return corm.DeleteRows(item, fields)
}

func UpdateByIdAll(item interface{}) error {
	corm := GetCORM()
	return corm.UpdateRow(item, false, nil)
}

func UpdateByIdByValuesMap(item interface{}, fields map[string]bool) error {
	corm := GetCORM()
	return corm.UpdateRow(item, true, fields)
}

func InsertItem(item interface{}) (int64, error) {
	corm := GetCORM()
	return corm.InsertRow(item)
}

func UpdateFunc(item interface{}, fields []string) (interface{}, error) {
	var result interface{}
	var err error
	if len(fields) == 0 {
		err = UpdateByIdAll(item)
	} else {
		theMap := make(map[string]bool)
		for _, v := range fields {
			theMap[strings.Title(strings.Trim(v, " "))] = true
		}
		err = UpdateByIdByValuesMap(
			item,
			theMap,
		)
	}
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("can't update")
	}
	result, err = GetByIdFromDB(item)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("can't get updated item")
	}

	return result, nil
}
