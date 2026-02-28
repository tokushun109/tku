package handler

import (
	"net/http"

	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	usecaseHealth "github.com/tokushun109/tku/backend/internal/usecase/health"
)

type HealthHandler struct {
	healthUC usecaseHealth.Usecase
}

func NewHealthHandler(healthUC usecaseHealth.Usecase) *HealthHandler {
	return &HealthHandler{healthUC: healthUC}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	if err := h.healthUC.Check(r.Context()); err != nil {
		response.WriteAppError(w, err)
		return
	}
	response.WriteSuccess(w)
}
