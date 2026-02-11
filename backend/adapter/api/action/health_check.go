package action

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type healthCheckResponse struct {
	Success bool `json:"success"`
}

type HealthCheckAction struct {
	db *gorm.DB
}

func NewHealthCheckAction(db *gorm.DB) HealthCheckAction {
	return HealthCheckAction{db: db}
}

func (a HealthCheckAction) Execute(w http.ResponseWriter, r *http.Request) {
	if a.db == nil {
		http.Error(w, "db connection error", http.StatusInternalServerError)
		return
	}

	sqlDB, err := a.db.DB()
	if err != nil {
		log.Println(err)
		http.Error(w, "db connection error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		log.Println(err)
		http.Error(w, "db ping error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthCheckResponse{Success: true})
}
