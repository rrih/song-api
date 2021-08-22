package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"
	"unicode/utf8"

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
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var p entity.InsertedUser
		json.NewDecoder(r.Body).Decode(&p)
		repository.Insert(p)
		// TODO: 保存されたデータを返す
		// TODO: クッキーをセットする
		// &http.Cookie{}
		// type Cookie struct {
		// 	Name  string
		// 	Value string
		// 	Path       string    // optional
		// 	Domain     string    // optional
		// 	Expires    time.Time // optional
		// 	RawExpires string    // for reading cookies only
		// 	// MaxAge=0 means no 'Max-Age' attribute specified.
		// 	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
		// 	// MaxAge>0 means Max-Age attribute present and given in seconds
		// 	MaxAge   int
		// 	Secure   bool
		// 	HttpOnly bool
		// 	SameSite SameSite
		// 	Raw      string
		// 	Unparsed []string // Raw text of unparsed attribute-value pairs
		// }
		// TODO: SetCookie みたいな関数に切り出す
		// http.SetCookie(w, cookie)
		// TODO: DeleteCookie みたいな関数も作る
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
