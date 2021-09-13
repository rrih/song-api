package usecase

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/rrih/managedby/pkg/domain/entity"
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
	if r.Method == "GET" {
		// リクエストのURLからuser_idを取り出す
		len := utf8.RuneCountInString("/api/v1/songs/view/")
		songID := r.URL.Path[len:]
		id, _ := strconv.Atoi(songID)
		body, err := repository.FindSongByID(id)
		if err != nil {
			entity.ErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		entity.SuccessResponse(w, map[string]interface{}{"data": body})
	}
}

// CreateSong 曲追加
func CreateSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
	if r.Method == "POST" {
		// TODO: 認証判定切り出す
		// header から読み出し
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		// token の認証
		_, err := VerifyToken(tokenString)
		if err != nil {
			entity.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}
		// ログイン中でないとsongを作成することはできない
		var song entity.Song
		json.NewDecoder(r.Body).Decode(&song)
		repository.SaveSong(song)
	}
}

// UpdateSong 曲更新
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
	if r.Method == "PUT" {
		var s entity.Song
		json.NewDecoder(r.Body).Decode(&s)
		repository.UpdateSong(s)
	}
}

// DeleteSong 曲削除
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
	if r.Method == "DELETE" {
		var s entity.Song
		json.NewDecoder(r.Body).Decode(&s)
		repository.DeleteSong(s)
	}
}
