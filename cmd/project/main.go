package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rrih/managedby/pkg/usecase"
)

// bootstrap
func main() {
	// router
	Router("/", Index)
	Router("/api/v1/users", usecase.FindUsers)
	Router("/api/v1/users/view/", usecase.FindUser)
	Router("/api/v1/users/signup/", usecase.CreateUsers)
	Router("/api/v1/users/update/", usecase.UpdateUser)
	Router("/api/v1/users/delete/", usecase.DeleteUser)
	// Router("/api/v1/users/signin/", usecase.SignIn)
	// Router("/api/v1/users/signout", usercase.SignOut)

	// TODO: 404、5XX 系のルーティング

	// TODO: dev -> "localhost:8080", prod -> ":8080"
	http.ListenAndServe("localhost:8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("managedby api"))
}

// Router is wrapper
func Router(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}
