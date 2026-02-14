package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/request"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseCategory "github.com/tokushun109/tku/clean-backend/internal/usecase/category"
)

type CategoryHandler struct {
	categoryUC usecaseCategory.Usecase
}

func NewCategoryHandler(categoryUC usecaseCategory.Usecase) *CategoryHandler {
	return &CategoryHandler{categoryUC: categoryUC}
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	q, err := request.ParseListCategoryQuery(r)
	if err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}
	list, err := h.categoryUC.List(r.Context(), q.Mode)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	res := presenter.ToCategoryResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.categoryUC.Create(r.Context(), req.Name); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["category_uuid"]

	var req request.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.categoryUC.Update(r.Context(), uuid, req.Name); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
