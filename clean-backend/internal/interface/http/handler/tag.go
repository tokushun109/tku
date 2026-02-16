package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/request"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseTag "github.com/tokushun109/tku/clean-backend/internal/usecase/tag"
)

type TagHandler struct {
	tagUC usecaseTag.Usecase
}

func NewTagHandler(tagUC usecaseTag.Usecase) *TagHandler {
	return &TagHandler{tagUC: tagUC}
}

func (h *TagHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.tagUC.List(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	res := presenter.ToTagResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateTagRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.tagUC.Create(r.Context(), req.Name); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *TagHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["tag_uuid"]

	var req request.UpdateTagRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.tagUC.Update(r.Context(), uuid, req.Name); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}

func (h *TagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["tag_uuid"]

	if err := h.tagUC.Delete(r.Context(), uuid); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
