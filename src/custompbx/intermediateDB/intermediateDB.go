package intermediateDB

import (
	"custompbx/db"
	"errors"
	"github.com/custompbx/customorm"
	"log"
	"strings"
)

// GetCORM initializes a custom ORM instance using the database connection from custompbx.
func GetCORM() *customorm.CORM {
	return customorm.Init(db.GetDB())
}

// GetAllFromDBAsMap retrieves all records of the specified item type from the database as a map.
func GetAllFromDBAsMap(item interface{}) (map[int64]interface{}, error) {
	corm := GetCORM()
	res, err := corm.GetDataAll(item, true)
	if err != nil {
		return nil, err
	}
	result, ok := res.(map[int64]interface{})
	if !ok {
		return nil, errors.New("expected map[int64]interface{}")
	}
	return result, nil
}

// GetAllFromDBAsSlice retrieves all records of the specified item type from the database as a slice.
func GetAllFromDBAsSlice(item interface{}) ([]interface{}, error) {
	corm := GetCORM()
	res, err := corm.GetDataAll(item, false)
	if err != nil {
		return nil, err
	}
	result, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("expected []interface{}")
	}
	return result, nil
}

// GetByIdFromDB retrieves a single record from the database by its ID, using a default ID of 0.
func GetByIdFromDB(item interface{}) (interface{}, error) {
	return GetByIdArg(item, 0)
}

// GetByIdArg retrieves a single record from the database by the specified ID.
func GetByIdArg(item interface{}, id int64) (interface{}, error) {
	corm := GetCORM()
	return corm.GetDataById(item, id)
}

// GetByValue retrieves records from the database that match the given field names and values.
func GetByValue(item interface{}, fieldNames map[string]bool) ([]interface{}, error) {
	filter := createFilterFields(fieldNames)
	return GetByValues(item, filter)
}

// GetByStructValue retrieves records from the database that match the given field names and values.
func GetByStructValue(item interface{}, fieldNames []string) ([]interface{}, error) {
	theMap := createFieldMap(fieldNames)
	filter := createFilterFields(theMap)
	return GetByValues(item, filter)
}

// GetByValues retrieves records from the database that match the given filter criteria.
func GetByValues(item interface{}, filter map[string]customorm.FilterFields) ([]interface{}, error) {
	corm := GetCORM()
	filterStr := customorm.Filters{Fields: filter}
	res, err := corm.GetDataByValue(item, filterStr, false)
	if err != nil {
		return nil, err
	}
	result, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("expected []interface{}")
	}
	return result, nil
}

// GetByFilteredValues retrieves records from the database using a pre-defined filter structure.
func GetByFilteredValues(item interface{}, filter customorm.Filters) ([]interface{}, error) {
	corm := GetCORM()
	res, err := corm.GetDataByValue(item, filter, false)
	if err != nil {
		return nil, err
	}
	result, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("expected []interface{}")
	}
	return result, nil
}

// GetByValueAsMap retrieves records from the database that match the given field names and values, and returns them as a map.
func GetByValueAsMap(item interface{}, fieldNames map[string]bool) (map[int64]interface{}, error) {
	filter := createFilterFields(fieldNames)
	return GetByValuesAsMap(item, filter)
}

// GetByValuesAsMap retrieves records from the database that match the given filter criteria, and returns them as a map.
func GetByValuesAsMap(item interface{}, filter map[string]customorm.FilterFields) (map[int64]interface{}, error) {
	corm := GetCORM()
	filterStr := customorm.Filters{Fields: filter}
	res, err := corm.GetDataByValue(item, filterStr, true)
	if err != nil {
		return nil, err
	}
	mapItem, ok := res.(map[int64]interface{})
	if !ok {
		return nil, errors.New("expected map[int64]interface{}")
	}
	return mapItem, err
}

// DeleteById deletes a record from the database by its ID.
func DeleteById(item interface{}) error {
	corm := GetCORM()
	return corm.DeleteRowById(item)
}

// DeleteByIdParam deletes a record from the database by a specified ID.
func DeleteByIdParam(item interface{}, id int64) error {
	corm := GetCORM()
	return corm.DeleteRowByArgId(item, id)
}

// DeleteRows deletes records from the database that match the specified field names and values.
func DeleteRows(item interface{}, fields map[string]bool) error {
	corm := GetCORM()
	return corm.DeleteRows(item, fields)
}

// UpdateByIdAll updates all fields of a record in the database.
func UpdateByIdAll(item interface{}) error {
	corm := GetCORM()
	return corm.UpdateRow(item, false, nil)
}

// UpdateByIdByValuesMap updates specified fields of a record in the database.
func UpdateByIdByValuesMap(item interface{}, fields map[string]bool) error {
	corm := GetCORM()
	return corm.UpdateRow(item, true, fields)
}

// InsertItem inserts a new record into the database and returns its ID.
func InsertItem(item interface{}) (int64, error) {
	corm := GetCORM()
	return corm.InsertRow(item)
}

// UpdateFunc updates a record in the database based on specified fields or all fields if none are specified, then retrieves the updated record.
func UpdateFunc(item interface{}, fields []string) (interface{}, error) {
	var err error
	if len(fields) == 0 {
		err = UpdateByIdAll(item)
	} else {
		theMap := createFieldMap(fields)
		err = UpdateByIdByValuesMap(item, theMap)
	}
	if err != nil {
		log.Println("Update error:", err)
		return nil, errors.New("can't update")
	}
	result, err := GetByIdFromDB(item)
	if err != nil {
		log.Println("Get updated item error:", err)
		return nil, errors.New("can't get updated item")
	}
	return result, nil
}

// createFilterFields converts a map of field names to a map of customorm.FilterFields.
func createFilterFields(fieldNames map[string]bool) map[string]customorm.FilterFields {
	filter := make(map[string]customorm.FilterFields)
	for k, v := range fieldNames {
		filter[k] = customorm.FilterFields{Flag: v}
	}
	return filter
}

// createFieldMap converts a slice of field names to a map with boolean values.
func createFieldMap(fields []string) map[string]bool {
	theMap := make(map[string]bool)
	for _, v := range fields {
		theMap[strings.Title(strings.TrimSpace(v))] = true
	}
	return theMap
}
