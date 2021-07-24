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
	http.HandleFunc("/api/v1/users/post/", usecase.CreateUsers)
	http.HandleFunc("/api/v1/users/update/", usecase.UpdateUser)
	http.HandleFunc("/api/v1/users/delete/", usecase.DeleteUser)

	// TODO: 404、5XX 系のルーティング

	// TODO: dev -> "localhost:8080", prod -> ":8080"
	http.ListenAndServe("localhost:8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("managedby api"))
}
