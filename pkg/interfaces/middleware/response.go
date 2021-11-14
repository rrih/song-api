package middleware

import (
	"encoding/json"
	"net/http"
	"os"
)

func Response(w http.ResponseWriter, err error, body map[string]interface{}) {
	if err == nil {
		data, _ := json.Marshal(body)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}
}

func SetupHeader(w http.ResponseWriter, r *http.Request) {
	origin := ""
	isProd := os.Getenv("PORT") != ""
	if isProd {
		origin = "https://songscoreonline.vercel.app"
	} else {
		origin = "http://localhost:3000"
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
