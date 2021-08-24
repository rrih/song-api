package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/dgrijalva/jwt-go"
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
		len := utf8.RuneCountInString("/api/v1/users/view/")
		userId := r.URL.Path[len:]
		id, _ := strconv.Atoi(userId)
		body := repository.FindById(id)
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

// ユーザー登録
// TODO: サインアップ系に命名修正する
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var p entity.InsertedUser
		json.NewDecoder(r.Body).Decode(&p)
		repository.Insert(p)
		// TODO: response用の構造体定義してポインタ返す
		// middleware.Response(w, nil, map[string]interface{}{"data": p})
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		// ユーザーIDを取得
		len := utf8.RuneCountInString("/api/v1/users/update/")
		userId := r.URL.Path[len:]
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
		len := utf8.RuneCountInString("/api/v1/users/delete/")
		userId := r.URL.Path[len:]
		id, _ := strconv.Atoi(userId)
		repository.LogicalDeleteUser(id)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// bodyの読み出し
	// ユーザ認証
	// tokenの発行
	// response
}

func CreateJwtToken(userID string) (string, error) {
	// tokenの生成
	// クレームの設定
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	// jwt token を返す
}
