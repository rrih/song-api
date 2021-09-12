package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rrih/managedby/pkg/usecase"
)

// bootstrap
func main() {
	// router
	http.HandleFunc("/", Index)
	http.HandleFunc("/api/v1/users", usecase.FindUsers)
	http.HandleFunc("/api/v1/users/view/", usecase.FindUser)
	http.HandleFunc("/api/v1/auth/signup/", usecase.CreateUsers)
	http.HandleFunc("/api/v1/auth/login/", usecase.Login)
	http.HandleFunc("/api/v1/auth/logout/", usecase.Logout)
	http.HandleFunc("/api/v1/mypage/", usecase.FindLoginUser)
	http.HandleFunc("/api/v1/mypage/update/", usecase.UpdateLoginUser)
	http.HandleFunc("/api/v1/mypage/delete/", usecase.DeleteAccount) // ここ論理削除

	// TODO: 404、5XX 系のルーティング

	// TODO: dev -> "localhost:8080", prod -> ":8080"
	http.ListenAndServe("localhost:8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("managedby api"))
}
