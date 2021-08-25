package entity

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json: "code"`
	Message string `json: "message"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	jsonData, err := json.Marshal(&Response{
		Code:    code,
		Message: message,
	})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
