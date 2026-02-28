package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseSkillMarket "github.com/tokushun109/tku/backend/internal/usecase/skill_market"
)

type SkillMarketHandler struct {
	skillMarketUC usecaseSkillMarket.Usecase
}

func NewSkillMarketHandler(skillMarketUC usecaseSkillMarket.Usecase) *SkillMarketHandler {
	return &SkillMarketHandler{skillMarketUC: skillMarketUC}
}

func (h *SkillMarketHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.skillMarketUC.List(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	res := presenter.ToSkillMarketResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *SkillMarketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateSkillMarketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.skillMarketUC.Create(r.Context(), req.Name, req.URL, req.Icon); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *SkillMarketHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["skill_market_uuid"]

	var req request.UpdateSkillMarketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.skillMarketUC.Update(r.Context(), uuid, req.Name, req.URL, req.Icon); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *SkillMarketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["skill_market_uuid"]

	if err := h.skillMarketUC.Delete(r.Context(), uuid); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
