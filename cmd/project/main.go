package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rrih/managedby/pkg/usecase"
)

// bootstrap
func main() {
	// router
	http.HandleFunc("/api/v1/users", usecase.FindUsers)
	http.HandleFunc("/api/v1/users/view/", usecase.FindUser)
	http.HandleFunc("/api/v1/auth/object/", usecase.AuthJwt)
	http.HandleFunc("/api/v1/auth/signup/", usecase.CreateUsers)
	http.HandleFunc("/api/v1/auth/login/", usecase.Login)
	http.HandleFunc("/api/v1/auth/logout/", usecase.Logout)
	http.HandleFunc("/api/v1/mypage/", usecase.FindLoginUser) // 自身の情報を取得するAPI
	http.HandleFunc("/api/v1/mypage/update/", usecase.UpdateLoginUser)
	http.HandleFunc("/api/v1/mypage/delete/", usecase.DeleteAccount) // ここ論理削除
	// songs
	http.HandleFunc("/api/v1/songs", usecase.FindSongs)
	http.HandleFunc("/api/v1/songs/view/", usecase.FindSong)
	http.HandleFunc("/api/v1/songs/add/", usecase.CreateSong)
	http.HandleFunc("/api/v1/songs/update/", usecase.UpdateSong)
	http.HandleFunc("/api/v1/songs/delete/", usecase.DeleteSong)
	// category
	http.HandleFunc("/api/v1/categories", usecase.FindCategories)
	http.HandleFunc("/api/v1/categories/view/", usecase.FindCategory)
	http.HandleFunc("/api/v1/categories/delete/", usecase.DeleteCategory)

	// 404
	http.HandleFunc("/", usecase.NotFound)

	port := "localhost:8080"
	isProd := os.Getenv("PORT") != ""
	if isProd {
		port = ":" + os.Getenv("PORT")
	}
	http.ListenAndServe(port, nil)
}
