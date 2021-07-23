package repository

import (
	"log"
	"time"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/infrastructure"
)

func FindAll() []entity.User {
	db := infrastructure.DbConn()
	rows, err := db.Query(
		"select id, name, email, password, is_admin, deleted, created, modified from users",
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	user := entity.User{}
	users := []entity.User{}
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		var isAdmin bool
		var deleted *string
		var created string
		var modified string
		err := rows.Scan(&id, &name, &email, &password, &isAdmin, &deleted, &created, &modified)
		if err != nil {
			panic(err)
		}
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.IsAdmin = isAdmin
		user.Deleted = deleted
		user.Created = created
		user.Modified = modified
		users = append(users, user)
	}
	defer db.Close()
	return users
}

func Insert(u entity.InsertedUser) {
	db := infrastructure.DbConn()
	// TODO: 日本時間にする
	created, modified := time.Now(), time.Now()
	_, err := db.Exec(
		"insert into users (name, email, password, is_admin, deleted, created, modified) values (?, ?, ?, ?, ?, ?, ?)", u.Name, u.Email, u.Password, u.IsAdmin, nil, created, modified,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Update(u entity.InsertedUser, id int) {
	db := infrastructure.DbConn()
	modified := time.Now()
	_, err := db.Exec(
		"update users set name = ?, email = ?, password = ?, modified = ? where id = ?",
		u.Name, u.Email, u.Password, modified, id,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
