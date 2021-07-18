package usecase

import (
	"net/http"

	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: add cors
	body := repository.FindAll()
	middleware.Response(w, nil, map[string]interface{}{"data": body})
}
