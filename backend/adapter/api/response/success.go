package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool `json:"success"`
}

type Success struct {
	status int
}

func NewSuccess(status int) Success {
	return Success{status: status}
}

func (s Success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.status)
	_ = json.NewEncoder(w).Encode(SuccessResponse{Success: true})
}
