package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	resp := Response{
		Message: message,
		Code:    code,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}
