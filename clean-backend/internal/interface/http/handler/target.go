package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/request"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseTarget "github.com/tokushun109/tku/clean-backend/internal/usecase/target"
)

type TargetHandler struct {
	targetUC usecaseTarget.Usecase
}

func NewTargetHandler(targetUC usecaseTarget.Usecase) *TargetHandler {
	return &TargetHandler{targetUC: targetUC}
}

func (h *TargetHandler) List(w http.ResponseWriter, r *http.Request) {
	q, err := request.ParseListTargetQuery(r)
	if err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}
	list, err := h.targetUC.List(r.Context(), q.Mode)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	res := presenter.ToTargetResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *TargetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateTargetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.targetUC.Create(r.Context(), req.Name); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *TargetHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["target_uuid"]

	var req request.UpdateTargetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.targetUC.Update(r.Context(), uuid, req.Name); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *TargetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["target_uuid"]

	if err := h.targetUC.Delete(r.Context(), uuid); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
