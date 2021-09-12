package entity

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// TODO: 配置場所再検討
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	var res Response
	res.Code = code
	res.Message = message
	jsonData, err := json.Marshal(map[string]interface{}{"data": res})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
