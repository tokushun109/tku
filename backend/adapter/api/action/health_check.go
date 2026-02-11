package action

import (
	"encoding/json"
	"net/http"
)

type healthCheckResponse struct {
	Success bool `json:"success"`
}

type HealthCheckAction struct{}

func NewHealthCheckAction() HealthCheckAction {
	return HealthCheckAction{}
}

func (a HealthCheckAction) Execute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthCheckResponse{Success: true})
}
