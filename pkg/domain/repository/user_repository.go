package repository

import (
	"log"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/infrastructure"
)

// TODO: repository から呼び出せるようにする
func FindAll() []entity.User {
	db := infrastructure.DbConn()
	rows, err := db.Query(
		"select id, name, email, password, created, modified from users",
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
