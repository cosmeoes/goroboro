package moneta

// import (
// 	"database/sql"
// 	"fmt"
// 	"strings"
// )

// type MysqlGrammar struct {
    
// } 

// func (mg MysqlGrammar) execSelect(qb QueryBuilder) (*sql.Rows, error) {
//     db := dbConn()
//     defer db.Close()

//     var queryColumns string;
//     if len(qb.columns) == 0 {
//         queryColumns = "*"
//     } else {
//         queryColumns = strings.Join(qb.columns, ", ")
//     }

//     var wheresQuery string
//     if len(qb.wheres) == 0 {
//         wheresQuery = ""
//     } else {
//         wheresQuery = "where " + strings.Join(qb.wheres, " and ")
//     }

//     tableName := qb.findTableName()
//     query := fmt.Sprintf("SELECT %s from %s %s", queryColumns, tableName, wheresQuery)
//     return db.Query(query)
// }


// func (mg MysqlGrammar) execDML(dmlBuilder DMLBuilder) {

// }


// func dbConn() (db *sql.DB) {
//     dbDriver := "mysql"
//     dbUser := "root"
//     dbPass := ""
//     dbName := "IMM"
//     db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
//     if err != nil {
//         panic(err.Error())
//     }
//     return db
// }
