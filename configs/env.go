package configs

import "os"

func SetEnv() {
	// mysql
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "root")
	os.Setenv("MYSQL_DB", "managedby_db")
	os.Setenv("MYSQL_HOST", "")
}
