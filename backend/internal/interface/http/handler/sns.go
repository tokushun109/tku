package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseSns "github.com/tokushun109/tku/backend/internal/usecase/sns"
)

type SnsHandler struct {
	snsUC usecaseSns.Usecase
}

func NewSnsHandler(snsUC usecaseSns.Usecase) *SnsHandler {
	return &SnsHandler{snsUC: snsUC}
}

func (h *SnsHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.snsUC.List(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	res := presenter.ToSnsResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *SnsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateSnsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.snsUC.Create(r.Context(), req.Name, req.URL); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *SnsHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["sns_uuid"]

	var req request.UpdateSnsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.snsUC.Update(r.Context(), uuid, req.Name, req.URL); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *SnsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["sns_uuid"]

	if err := h.snsUC.Delete(r.Context(), uuid); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
