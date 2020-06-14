package moneta

import (
	"fmt"
	"reflect"
	"strconv"
)

type QueryBuilder struct {
    model interface{}
    wheres []string
    columns []string
}


func (q *QueryBuilder) Model(model interface{}) {
    q.model = model
}

func (q QueryBuilder) Where(column string, condition string, value string) QueryBuilder {
    q.wheres = append(q.wheres, fmt.Sprintf("%s %s %s", column, condition, value))
    return q
}

func (q QueryBuilder) WhereNotNull(column string) QueryBuilder {
    return q.Where(column, "not null", "")
}

func (q QueryBuilder) Get(columns ...string) []interface{} {
    q.columns = columns
    rows, err := MysqlGrammar{}.execSelect(q)
    if err != nil {
        fmt.Println("Error executing select", err)
    }

    tableColums, err := rows.Columns()
    if err != nil {
        //TODO: handle error 
        fmt.Println(err)
    }
    
    var resultRows []map[string]interface{}
    var objects []interface{}

    for rows.Next() {
        columns := make([]interface{}, len(tableColums))
        columnPointers := make([]interface{}, len(tableColums))
        for i := range tableColums {
            columnPointers[i]  = &columns[i]
        }

        rows.Scan(columnPointers...)

        colMap := make(map[string]interface{}, len(tableColums))
        for i, name := range tableColums {
            colMap[name] = columns[i]
        }

        v := reflect.New(reflect.TypeOf(q.model)).Elem()
        for i := 0; i < v.NumField(); i++ {
            field := v.Type().Field(i)
            name, _ := field.Tag.Lookup("moneta")

            if (colMap[name] == nil) {
                continue
            }

            switch field.Type.Kind() {
            case reflect.String:
                v.Field(i).SetString(string(colMap[name].([]byte)))
            case reflect.Int:
                stringValue := string(colMap["id"].([]byte))
                value, _ := strconv.ParseInt(stringValue, 10, 64)
                v.Field(i).SetInt(int64(value))
            }
        }

        objects = append(objects, v.Interface())

        resultRows = append(resultRows, colMap)
    }

    return objects
}


func (q QueryBuilder) First() interface{} {
    return q.Get()[0]
}

