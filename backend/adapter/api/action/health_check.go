package action

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/tokushun109/tku/backend/adapter/api/logging"
	"github.com/tokushun109/tku/backend/adapter/logger"
	"gorm.io/gorm"
)

type healthCheckResponse struct {
	Success bool `json:"success"`
}

type HealthCheckAction struct {
	db  *gorm.DB
	log logger.Logger
}

func NewHealthCheckAction(db *gorm.DB, log logger.Logger) HealthCheckAction {
	return HealthCheckAction{db: db, log: log}
}

func (a HealthCheckAction) Execute(w http.ResponseWriter, r *http.Request) {

	if a.db == nil {
		logging.NewError(a.log, r, http.StatusInternalServerError, nil).Log("db connection is nil")
		http.Error(w, "db connection error", http.StatusInternalServerError)
		return
	}

	sqlDB, err := a.db.DB()
	if err != nil {
		logging.NewError(a.log, r, http.StatusInternalServerError, err).Log("db connection error")
		http.Error(w, "db connection error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		logging.NewError(a.log, r, http.StatusInternalServerError, err).Log("db ping error")
		http.Error(w, "db ping error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(healthCheckResponse{Success: true}); err != nil {
		logging.NewError(a.log, r, http.StatusInternalServerError, err).Log("response encode error")
		return
	}
}
