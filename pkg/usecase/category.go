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

// FindCategories カテゴリー一覧を取得
func FindCategories(w http.ResponseWriter, r *http.Request) {
	err := middleware.SetupHeader(w, r)
	if err != nil {
		return
	}
	if r.Method == "GET" {
		body := repository.FindAllCategories()
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

// FindCategory idから特定のカテゴリーを取得
// 使用するか要検討
func FindCategory(w http.ResponseWriter, r *http.Request) {
	err := middleware.SetupHeader(w, r)
	if err != nil {
		return
	}
	if r.Method == "GET" {
		// リクエストのURLからuser_idを取り出す
		len := utf8.RuneCountInString("/api/v1/songs/view/")
		categoryID := r.URL.Path[len:]
		id, _ := strconv.Atoi(categoryID)
		body, err := repository.FindCategoryByID(id)
		if err != nil {
			entity.ErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		entity.SuccessResponse(w, map[string]interface{}{"data": body})
	}
}

// DeleteCategory 曲削除
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	err := middleware.SetupHeader(w, r)
	if err != nil {
		return
	}
	if r.Method == "DELETE" {
		var c entity.Category
		json.NewDecoder(r.Body).Decode(&c)
		repository.DeleteCategory(c)
	}
}
