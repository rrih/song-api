package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/rrih/managedby/pkg/domain/entity"
	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: add cors
	body := repository.FindAll()
	// TODO: エラー時を想定していない
	middleware.Response(w, nil, map[string]interface{}{"data": body})
}

func AddUsers(w http.ResponseWriter, r *http.Request) {
	var p entity.InsertedUser
	json.NewDecoder(r.Body).Decode(&p)
	repository.Insert(p)
	// TODO: 保存されたデータを返す
}
