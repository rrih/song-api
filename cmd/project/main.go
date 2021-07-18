package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// TODO: あとで /pkg/domain 配下に移動する
type User struct {
	ID       int    `json: "id"`
	Name     string `json: "string"`
	Email    string `json: "string"`
	Password string `json: "string"`
	Created  string `json: "string"`
	Modified string `json: "string"`
}

// bootstrap
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		w.Write([]byte(host + "hogeho77777ge"))
	})

	// コントローラー割当
	// http.Handle の第二引数に HandlerFunc 型の handler を持たせる
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// TODO: add cors
		body := FindAll()
		data, _ := json.Marshal(body)
		w.Write(data)
	})

	// dev環境では"localhost:8080"を指定、本番では":8080"を指定
	http.ListenAndServe("localhost:8080", nil)
}

func dbConn() (db *sql.DB) {
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}
	err := godotenv.Load(fmt.Sprintf("./.env"))
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
	fmt.Println(dbConf)
	db, openErr := sql.Open("mysql", dbConf)
	if openErr != nil {
		panic(openErr.Error())
	}
	return db
}

// TODO: repository から呼び出せるようにする
func FindAll() []User {
	db := dbConn()
	rows, err := db.Query("select id, name, email, password, created, modified from users")
	if err != nil {
		log.Fatal(err.Error())
	}
	user := User{}
	users := []User{}
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		var created string
		var modified string
		err := rows.Scan(&id, &name, &email, &password, &created, &modified)
		if err != nil {
			panic(err)
		}
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.Created = created
		user.Modified = modified
		users = append(users, user)
	}
	defer db.Close()
	return users
}
