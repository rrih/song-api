package infrastructure

import (
	"database/sql"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlRows struct {
	Rows *sql.Rows
}

type SqlResult struct {
	Result *sql.Result
}

// func InitSqlHandler() database.SqlHandler {
// 	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s%s/%s", configs.MysqlUser, configs.MysqlPassword, configs.MysqlHost, configs.MysqlDataBase))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	sqlHandler := new(SqlHandler)
// 	sqlHandler.Conn = conn
// 	return sqlHandler
// }
