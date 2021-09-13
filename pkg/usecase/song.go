package usecase

import (
	"net/http"

	"github.com/rrih/managedby/pkg/domain/repository"
	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

// FindSongs 全曲を取得
func FindSongs(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
	if r.Method == "GET" {
		body := repository.FindAllSongs()
		middleware.Response(w, nil, map[string]interface{}{"data": body})
	}
}

// FindSong idから曲を取得
func FindSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
}

// CreateSong 曲追加
func CreateSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
}

// UpdateSong 曲更新
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
}

// DeleteSong 曲削除
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
}
