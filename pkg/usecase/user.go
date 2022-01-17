package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
	if r.Method == "GET" {
		body := repository.FindAllUsers()
		// TODO: エラー時を想定していない
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	_, err := IsLogin(w, r)
	if err != nil {
		return
	}
	if r.Method == "GET" {
		// リクエストのURLからuser_idを取り出す
		len := utf8.RuneCountInString("/api/v1/users/view/")
		userId := r.URL.Path[len:]
		id, _ := strconv.Atoi(userId)
		body, err := repository.FindById(id)
		if err != nil {
			entity.ErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		entity.SuccessResponse(w, map[string]interface{}{"data": body})
	}
}

// ユーザー登録
// TODO: サインアップ系に命名修正する
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	// cors解決
	middleware.SetupHeader(w, r)
	if r.Method == "POST" {
		// TODO: 絶対ここでパスワードのハッシュ化するのおかしいからあとで直す
		var p entity.InsertedUser
		json.NewDecoder(r.Body).Decode(&p)
		// TODO: 直す
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

// header の jwt から jwt と user id を返す
func AuthJwt(w http.ResponseWriter, r *http.Request) {
	// cors
	middleware.SetupHeader(w, r)
	// ログイン判定
	IsLogin(w, r)

	token, err := IsLogin(w, r)
	if err != nil {
		entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		// エラーなら終了
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["user"]
	body, err := repository.FindByEmail(email.(string))
	if err != nil {
		entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := repository.FindByEmail(body.Email)
	entity.SuccessResponse(w, &entity.LoginResponse{
		Token:  r.Header.Get("Authorization"),
		UserID: user.ID,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	// cors解決
	err := middleware.SetupHeader(w, r)
	if err != nil {
		return
	}
	if r.Method == "POST" {
		// bodyの読み出し
		body, err := io.ReadAll(r.Body)
		if err != nil {
			// エラー処理
			panic(err)
		}
		var loginRequest entity.LoginRequest
		err = json.Unmarshal(body, &loginRequest)
		if err != nil {
			// TODO: エラー処理
			panic(err)
		}
		// ユーザ認証
		// TODO: あとで整理
		var hash string
		db := infrastructure.DbConn()
		row := db.QueryRow("SELECT password FROM users WHERE email=?", loginRequest.Email)
		if err = row.Scan(&hash); err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		}
		// パスワード検証
		err = PasswordVerify(hash, loginRequest.Password)
		if err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		}
		// tokenの発行
		token, err := CreateJwtToken(loginRequest.Email)
		if err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		}
		// response
		// ex: {"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzAwOTIwODUsInVzZXIiOiJyc2tsdnZAdGVzdC5kZGRkZGQifQ.5jo5phdc-WuVaeYEalz5qn0my3AJHHlv4wQwudBambY"}
		user, err := repository.FindByEmail(loginRequest.Email)
		entity.SuccessResponse(w, &entity.LoginResponse{
			Token:  token,
			UserID: user.ID,
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
	var secretKey = os.Getenv("SIGNING_KEY")
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
		return []byte(os.Getenv("SIGNING_KEY")), nil // CreateJwtToken で指定した任意の文字列
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

// LogoutResponse はログアウトレスポンス用
type LogoutResponse struct {
	Message string `json:"message"`
}
