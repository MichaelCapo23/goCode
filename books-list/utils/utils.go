package utils

import (
	"books-list/model"
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, status int, err model.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func sendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
