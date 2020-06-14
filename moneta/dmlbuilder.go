package moneta

import (
	"database/sql"
)

type DMLBuilder struct {
    model interface{}
    gr grammar
    columnValues map[string]string
    columnOrder []string
}

func (d DMLBuilder) Save() (sql.Result, error) {
    d.columnValues, d.columnOrder = getColumnsValuesMap(d.model) 
    return MysqlGrammar{}.execSave(d)
}

func (d DMLBuilder) Update() error {
    d.columnValues, d.columnOrder = getColumnsValuesMap(d.model) 
    return d.gr.execUpdate(d)
}

func (d DMLBuilder) Delete() (sql.Result, error) {
    d.columnValues, d.columnOrder = getColumnsValuesMap(d.model) 
    return MysqlGrammar{}.execDelete(d)
}

