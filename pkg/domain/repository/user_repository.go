package repository

import (
	"log"

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

// func Add() entity.User {
// 	db := infrastructure.DbConn()
// 	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
// 	created := time.Now()
// 	modified := time.Now()
// 	rows, err := db.Query(
// 		"insert into users (id, name, email, password, created, modified) values (?, ?, ?, ?)",
// 	)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// }
