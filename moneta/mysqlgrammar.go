package moneta

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/cosmeoes/goroboro/config"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlGrammar struct {
    
} 

func (mg MysqlGrammar) execSelect(qb QueryBuilder) (*sql.Rows, error) {
    db := dbConn()
    defer db.Close()

    var queryColumns string;
    if len(qb.columns) == 0 {
        queryColumns = "*"
    } else {
        queryColumns = strings.Join(qb.columns, ", ")
    }

    var wheresQuery string
    if len(qb.wheres) == 0 {
        wheresQuery = ""
    } else {
        wheresQuery = "where " + strings.Join(qb.wheres, " and ")
    }

    tableName := findTableName(qb.model)
    query := fmt.Sprintf("SELECT %s from %s %s", queryColumns, tableName, wheresQuery)
    return db.Query(query)
}


func (mg MysqlGrammar) execSave(dml DMLBuilder) (sql.Result, error) {
    db := dbConn()
    defer db.Close()
    tableName := findTableName(dml.model)

    columns := "("
    valuesPlaceHolder := "("
    var values []interface{}
    var updateValues []interface{}
    updateFields := ""
    for _, name := range dml.columnOrder {
        columns += name + ", "
        valuesPlaceHolder += "?, "
        values = append(values, dml.columnValues[name])
        updateValues = append(updateValues, dml.columnValues[name])
        updateFields += name + " = ?, "
    }
    columns = strings.Trim(columns, ", ") + ")"
    updateFields = strings.Trim(updateFields, ", ")
    valuesPlaceHolder = strings.Trim(valuesPlaceHolder, ", ") + ")"
    insert := "INSERT INTO " + tableName + columns + " VALUES " + valuesPlaceHolder +
    " ON DUPLICATE KEY UPDATE " + updateFields 
    insForm, err := db.Prepare(insert)
    if err != nil {
        return nil, err
    }

    values = append(values, updateValues...)
    result, err := insForm.Exec(values...)
    return result, err
}


func (mg MysqlGrammar) execUpdate(dml DMLBuilder) (sql.Result, error) {
    db := dbConn()
    defer db.Close()
    // tableName := findTableName(dml.model)
    columns := "("
    var values []interface{}
    var first = true 
    for name, value := range dml.columnValues {
        if !first {
            columns += " and "
        }
        columns += name + " = ?"
        values = append(values, value)
        first = false
    }

    columns = strings.Trim(columns, "and ") + ")"
    // delete := "DELETE FROM " + tableName + " where " + columns
    // deleteSmt, err := db.Prepare(delete)
    return nil, nil
}

func (mg MysqlGrammar) execDelete(dml DMLBuilder) (sql.Result, error) {
    db := dbConn()
    defer db.Close()
    tableName := findTableName(dml.model)

    columns := "("
    var values []interface{}
    var first = true 
    for name, value := range dml.columnValues {
        if !first {
            columns += " and "
        }
        columns += name + " = ?"
        values = append(values, value)
        first = false
    }

    columns = strings.Trim(columns, "and ") + ")"
    delete := "DELETE FROM " + tableName + " where " + columns
    deleteSmt, err := db.Prepare(delete)

    if err != nil {
        return nil, err 
    }

    return deleteSmt.Exec(values...)
}

func dbConn() (db *sql.DB) {
    dbConfig := config.GetDBConfig()
    db, err := sql.Open(dbConfig.Driver, dbConfig.ToDataSource())
    if err != nil {
        panic(err.Error())
    }
    return db
}
