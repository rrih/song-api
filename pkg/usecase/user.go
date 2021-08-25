package usecase

import (
	"encoding/json"
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
		// エラー処理
		panic(err)
		return
	}
	// パスワード検証
	err = PasswordVerify(hash, req.Password)
	if err != nil {
		panic(err)
	}
	log.Println("login success: email = ", req.Email)

	// tokenの発行
	token, err := CreateJwtToken(req.Email)
	if err != nil {
		panic(err)
	}
	// response
	entity.SuccessResponse(w, &entity.LoginResponse{
		Token: token,
	})
}

func CreateJwtToken(userID string) (string, error) {
	// tokenの生成
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// クレームの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定
	}
	// 署名
	var secretKey = "secret" // 任意の文字列
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // CreateTokenにて指定した文字列を使います
	})
	// これエラー用の分岐いらなそう
	if err != nil {
		return token, err
	}
	return token, nil
}

// リクエストの読み出し→jwtの検証→レスポンスの作成

// パスワード検証
func PasswordVerify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func Logout(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
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

type LogoutResponse struct {
	Message string `json: "message"`
}
