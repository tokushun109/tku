package response

import (
	"encoding/json"
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type SuccessResponse struct {
	Success bool `json:"success"`
}

type Success struct {
	status int
	log    logger.Logger
}

func NewSuccess(log logger.Logger, status int) Success {
	return Success{status: status, log: log}
}

func (s Success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.status)
	if err := json.NewEncoder(w).Encode(SuccessResponse{Success: true}); err != nil && s.log != nil {
		s.log.Errorf("failed to write success response: %v", err)
	}
}
