package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/dgrijalva/jwt-go"
	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/infrastructure"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
	"golang.org/x/crypto/bcrypt"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	// TODO: add cors
	if r.Method == "GET" {
		body := repository.FindAll()
		// TODO: エラー時を想定していない
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	if r.Method == "GET" {
		// リクエストのURLからuser_idを取り出す
		len := utf8.RuneCountInString("/api/v1/users/view/")
		userId := r.URL.Path[len:]
		id, _ := strconv.Atoi(userId)
		body, err := repository.FindById(id)
		if err != nil {
			middleware.Response(w, err, map[string]interface{}{"data": body})
		}
		middleware.Response(w, err, map[string]interface{}{"data": body})
	}
}

// ユーザー登録
// TODO: サインアップ系に命名修正する
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	if r.Method == "POST" {
		// TODO: 絶対ここでパスワードのハッシュ化するのおかしいからあとで直す
		// hash, err = PasswordHash()
		var p entity.InsertedUser
		json.NewDecoder(r.Body).Decode(&p)
		// TODO: てきとうすぎるので直す
		password, err := PasswordHash(p.Password)
		if err != nil {
			panic(err)
		}
		p.Password = password
		repository.Insert(p)
		// TODO: response用の構造体定義してポインタ返す
		// middleware.Response(w, nil, map[string]interface{}{"data": p})
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	if r.Method == "DELETE" {
		len := utf8.RuneCountInString("/api/v1/users/delete/")
		userId := r.URL.Path[len:]
		id, _ := strconv.Atoi(userId)
		repository.LogicalDeleteUser(id)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	if r.Method == "POST" {
		// bodyの読み出し
		body, err := io.ReadAll(r.Body)
		if err != nil {
			// エラー処理
			panic(err)
		}
		var req entity.LoginRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			panic(err)
		}

		// ユーザ認証
		// TODO: やっつけ実装すぎるのであとで整理
		var hash string
		db := infrastructure.DbConn()
		row := db.QueryRow("SELECT password FROM users WHERE email=?", req.Email)
		if err = row.Scan(&hash); err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		}
		// パスワード検証
		err = PasswordVerify(hash, req.Password)
		if err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		}

		// tokenの発行
		token, err := CreateJwtToken(req.Email)
		if err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		}
		// response
		// ex: {"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzAwOTIwODUsInVzZXIiOiJyc2tsdnZAdGVzdC5kZGRkZGQifQ.5jo5phdc-WuVaeYEalz5qn0my3AJHHlv4wQwudBambY"}
		entity.SuccessResponse(w, &entity.LoginResponse{
			Token: token,
		})
	}
}

func CreateJwtToken(userID string) (string, error) {
	// tokenの生成
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// クレームの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定(とりあえず1時間としている)
	}
	// 署名
	var secretKey = "secret" // 任意の文字列
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// token の検証
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // CreateJwtToken で指定した任意の文字列
	})
	return token, err
}

// リクエストの読み出し→jwtの検証→レスポンスの作成

// パスワード検証
func PasswordVerify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// パスワードのハッシュ化
func PasswordHash(password string) (string, error) {
	if len(password) > 70 {
		return "", errors.New("password is too long.")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("logout")
	// header から読み出し
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	// token の認証
	token, err := VerifyToken(tokenString)
	if err != nil {
		entity.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	// response
	claims := token.Claims.(jwt.MapClaims)
	entity.SuccessResponse(w, &LogoutResponse{
		Message: fmt.Sprintf("bye %s !", claims["user"]),
	})
}

// ログインのために users.name と users.email を取得する
func getLoginUserByAuthHeaderToken(token string) {

}

type LogoutResponse struct {
	Message string `json: "message"`
}
