package middleware

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, err error, body map[string]interface{}) error {
	if err == nil {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(body)
		w.Write(data)
	}
	return err
}
