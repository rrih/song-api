package usecase

import (
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
