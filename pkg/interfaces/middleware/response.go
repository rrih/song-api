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

func SetupHeader(w http.ResponseWriter, r *http.Request) error {
	origin := ""
	isProd := os.Getenv("PORT") != ""
	if isProd {
		origin = "https://sso-front.vercel.app"
	} else {
		origin = "http://localhost:3000"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Content-Language, Accept-Language, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// preflight用
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}
