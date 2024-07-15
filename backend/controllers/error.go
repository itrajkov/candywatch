package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/itrajkov/candywatch/backend/dtos"
)

func errorHandler(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(dtos.NewErrorResponse(msg))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
