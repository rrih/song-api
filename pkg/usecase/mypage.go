package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dgrijalva/jwt-go"
	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

func FindLoginUser(w http.ResponseWriter, r *http.Request) {
	// cors
	middleware.SetupHeader(w, r)
	// ログイン判定
	// 認証チェック
	// headerからjwtを確認
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := VerifyToken(tokenString)
	if err != nil {
		entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	email := claims["user"]
	body, err := repository.FindByEmail(email.(string))
	middleware.Response(w, err, map[string]interface{}{"data": body})
}

// ログイン時、ログインユーザーの情報を更新する機能
// 主に名前、メールアドレス
func UpdateLoginUser(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
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

// 退会処理 という名の論理削除
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	if r.Method == "DELETE" {
		len := utf8.RuneCountInString("/api/v1/users/delete/")
		userId := r.URL.Path[len:]
		id, _ := strconv.Atoi(userId)
		repository.LogicalDeleteUser(id)
	}
}
