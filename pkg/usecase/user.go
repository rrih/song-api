package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: add cors
	if r.Method == "GET" {
		body := repository.FindAll()
		// TODO: エラー時を想定していない
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		userId := strings.Trim(r.URL.Path, "/api/v1/users/view/")
		id, _ := strconv.Atoi(userId)
		body := repository.FindById(id)
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var p entity.InsertedUser
		json.NewDecoder(r.Body).Decode(&p)
		repository.Insert(p)
		// TODO: 保存されたデータを返す
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		// ユーザーIDを取得
		userId := strings.Trim(r.URL.Path, "/api/v1/users/update/")
		id, _ := strconv.Atoi(userId)
		// TODO: http メソッドが put であるかチェックする
		// TODO: id 存在するユーザIDか存在しないユーザIDかでエラーハンドリングする
		// TODO: 異常系
		var p entity.InsertedUser
		json.NewDecoder(r.Body).Decode(&p)
		repository.Update(p, id)
		// TODO: repository.Update の結果を response として返す
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		userId := strings.Trim(r.URL.Path, "/api/v1/users/delete/")
		id, _ := strconv.Atoi(userId)
		repository.LogicalDeleteUser(id)
	}
}
