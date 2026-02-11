package action

import (
	"context"
	"net/http"
	"time"

	"github.com/tokushun109/tku/backend/adapter/api/response"
	"github.com/tokushun109/tku/backend/adapter/logger"
	"github.com/tokushun109/tku/backend/adapter/repository"
)

type HealthCheckAction struct {
	db  repository.SQLDB
	log logger.Logger
}

func NewHealthCheckAction(db repository.SQLDB, log logger.Logger) HealthCheckAction {
	return HealthCheckAction{db: db, log: log}
}

func (a HealthCheckAction) Execute(w http.ResponseWriter, r *http.Request) {

	if a.db == nil {
		response.LogAndSendError(
			w,
			r,
			a.log,
			http.StatusInternalServerError,
			nil,
			"db connection error",
		)
		return
	}

	sqlDB, err := a.db.DB()
	if err != nil {
		response.LogAndSendError(
			w,
			r,
			a.log,
			http.StatusInternalServerError,
			err,
			"db connection error",
		)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		response.LogAndSendError(
			w,
			r,
			a.log,
			http.StatusInternalServerError,
			err,
			"db ping error",
		)
		return
	}

	response.NewSuccess(http.StatusOK).Send(w)
}
