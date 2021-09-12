package infrastructure

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
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

func DbConn() (db *sql.DB) {
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}
	// err := godotenv.Load(fmt.Sprintf("./.env"))
	// see: https://staticcheck.io/docs/checks#S1039
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	var (
		user           = os.Getenv("MYSQL_USER")
		password       = os.Getenv("MYSQL_PASSWORD")
		database       = os.Getenv("DB_NAME")
		connectionName = os.Getenv("DB_PROTOCOL")
	)
	dbConf := user + ":" + password + "@" + connectionName + "/" + database + "?charset=utf8mb4"
	db, openErr := sql.Open("mysql", dbConf)
	if openErr != nil {
		panic(openErr.Error())
	}
	return db
}
