package usecase

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

func FindLoginUser(w http.ResponseWriter, r *http.Request) {
	// cors
	middleware.SetupHeader(w, r)
	// TODO: この分岐、いけてない気がするので要検討
	// if r.Method == "GET" {
	// 認証チェック
	// headerからjwtを確認
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println("tokenString")
	fmt.Println(tokenString)
	token, err := VerifyToken(tokenString)
	if err != nil {
		entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
	}
	fmt.Println("token")
	fmt.Println(token)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("claims")
	fmt.Println(claims)
	email := claims["user"]
	fmt.Println("email")
	// fmt.Println(email.(string))
	body, err := repository.FindByEmail(email.(string))
	entity.SuccessResponse(w, email)
	// body, err := repository.FindById(userID)
	// if err != nil {
	middleware.Response(w, err, map[string]interface{}{"data": body})
	// }
	// middleware.Response(w, err, map[string]interface{}{"data": body})
	// } else {
	// 	middleware.Response(w, nil, r.Header.Get())
	// }
	// ログイン判定
	// ログインユーザー取得
	// ユーザー情報返却
}
