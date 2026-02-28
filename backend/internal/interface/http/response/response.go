package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func WriteJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, ErrorResponse{Message: msg})
}

func WriteAppError(w http.ResponseWriter, err error) {
	status, msg := MapError(err)
	if status >= 500 {
		log.Printf("internal error: %v", err)
	}
	WriteError(w, status, msg)
}

func WriteSuccess(w http.ResponseWriter) {
	WriteJSON(w, http.StatusOK, SuccessResponse{Success: true})
}
