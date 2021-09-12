package repository

import (
	"errors"
	"log"
	"time"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/infrastructure"
)

func FindAll() []entity.User {
	db := infrastructure.DbConn()
	rows, err := db.Query(
		"select id, name, email, password, is_admin, deleted, created, modified from users where deleted is null",
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

// id からユーザを取得する
func FindById(userId int) (entity.User, error) {
	db := infrastructure.DbConn()
	row, err := db.Query(
		"select id, name, email, password, is_admin, deleted, created, modified from users where id = ?", userId,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	var u entity.User
	for row.Next() {
		err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.Deleted, &u.Created, &u.Modified)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// TODO: 要検討。仮で該当データが存在しなかった場合を u.ID == 0 としてる
	if u.ID == 0 {
		err := errors.New("該当データが存在しません")
		return u, err
	}
	// 一旦以下コメントアウト
	// see: https://yaruki-strong-zero.hatenablog.jp/entry/go_sql_open_close
	defer row.Close()
	return u, nil
}

// email からユーザを取得する
// とりあえず FindByID のコピペ
func FindByEmail(email string) (entity.User, error) {
	db := infrastructure.DbConn()
	row, err := db.Query(
		"select id, name, email, password, is_admin, deleted, created, modified from users where email = ?", email,
	)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	var u entity.User
	for row.Next() {
		err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.Deleted, &u.Created, &u.Modified)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// TODO: 要検討。仮で該当データが存在しなかった場合を u.ID == 0 としてる
	if u.ID == 0 {
		err := errors.New("該当データが存在しません")
		if err != nil {
			return u, err
		}
	}
	return u, nil
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

// 論理削除
func LogicalDeleteUser(id int) {
	db := infrastructure.DbConn()
	modified := time.Now()
	deleted := time.Now()
	_, err := db.Exec(
		"update users set deleted = ?, modified = ? where id = ?",
		deleted, modified, id,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// 物理削除
func PhysicalDeleteUser(id int) {
	db := infrastructure.DbConn()
	_, err := db.Exec(
		"delete from users where id = ?", id,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
