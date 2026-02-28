package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseSalesSite "github.com/tokushun109/tku/backend/internal/usecase/sales_site"
)

type SalesSiteHandler struct {
	salesSiteUC usecaseSalesSite.Usecase
}

func NewSalesSiteHandler(salesSiteUC usecaseSalesSite.Usecase) *SalesSiteHandler {
	return &SalesSiteHandler{salesSiteUC: salesSiteUC}
}

func (h *SalesSiteHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.salesSiteUC.List(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	res := presenter.ToSalesSiteResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *SalesSiteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateSalesSiteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.salesSiteUC.Create(r.Context(), req.Name, req.URL, req.Icon); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *SalesSiteHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["sales_site_uuid"]

	var req request.UpdateSalesSiteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.salesSiteUC.Update(r.Context(), uuid, req.Name, req.URL, req.Icon); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *SalesSiteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["sales_site_uuid"]

	if err := h.salesSiteUC.Delete(r.Context(), uuid); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
