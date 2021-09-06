package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, err error, body map[string]interface{}) error {
	// 一旦全部許可する
	// TODO: あとでCORS再検討
	fmt.Println(w)
	fmt.Println("eeee")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(body)
		w.Write(data)
	}
	return err
}

func SetupHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "https://managedby-next.vercel.app")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
