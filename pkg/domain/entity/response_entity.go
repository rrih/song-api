package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code    int    `json: "code"`
	Message string `json: "message"`
}

// TODO: 配置場所再検討
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)
	fmt.Println(data)
	w.WriteHeader(http.StatusOK)
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
