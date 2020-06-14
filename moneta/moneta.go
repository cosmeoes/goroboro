package moneta

import (
	"database/sql"
	"reflect"
	"strconv"
)

func Find(model interface{}) QueryBuilder {
    qb := QueryBuilder{}
    qb.Model(model)
    return qb
}

func Model(model interface{}) DMLBuilder {
    dmlb := DMLBuilder{}
    dmlb.model = model
    return dmlb
}


type grammar interface {
    execSelect(QueryBuilder) *sql.Rows
    execSave(DMLBuilder) error
    execUpdate(DMLBuilder) error
    execDelete(DMLBuilder) error
}

func findTableName(model interface{}) string {
    tableName := reflect.TypeOf(model).Name()

    v := reflect.New(reflect.TypeOf(model))
    nameMethod := v.MethodByName("TableName")
    if nameMethod.Kind() != reflect.Invalid {
        if f, ok := nameMethod.Interface().(func() string); ok {
            tableName = f()
        }
    }

    return tableName
}

func getColumnsValuesMap(model interface{}) (columnValues map[string]string, columnOrder []string){
    v := reflect.ValueOf(model)

    columnValues = make(map[string]string, v.NumField())
    for i := 0; i < v.NumField(); i++ { 
        field := v.Type().Field(i)
        name, _ := field.Tag.Lookup("moneta")

        switch v.Field(i).Type().Kind() {
        case reflect.Int:
            columnOrder = append(columnOrder, name)
            columnValues[name] = strconv.FormatInt(v.Field(i).Int(), 10)
        case reflect.String:
            columnOrder = append(columnOrder, name)
            columnValues[name] = v.Field(i).String()
        }
    }

    return columnValues, columnOrder
}

