package usecase

import (
	"net/http"
)

// FindCategories カテゴリー一覧を取得
func FindCategories(w http.ResponseWriter, r *http.Request) {
}

// FindSong idから特定のカテゴリーを取得
func FindCategory(w http.ResponseWriter, r *http.Request) {
}

// DeleteCategory 曲削除
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
}
