package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rrih/managedby/pkg/usecase"
)

// bootstrap
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("managedby api"))
	})

	// router
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		usecase.FindUsers(w, r)
	})
	http.HandleFunc("/users/signup", func(w http.ResponseWriter, r *http.Request) {
		usecase.AddUsers(w, r)
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		usecase.UpdateUser(w, r)
	})

	// TODO: dev -> "localhost:8080", prod -> ":8080"
	http.ListenAndServe("localhost:8080", nil)
}
