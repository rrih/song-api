package usecase

import (
	"net/http"

	"github.com/rrih/managedby/pkg/interfaces/middleware"
)

// NotFound ルーティングで path 一致がなかった場合返す
func NotFound(w http.ResponseWriter, r *http.Request) {
	middleware.SetupHeader(w, r)
	body := make(map[string]string, 2)
	body["code"] = "404"
	body["message"] = "Not Found"
	middleware.Response(w, nil, map[string]interface{}{"data": body})
}
