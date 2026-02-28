package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseCreator "github.com/tokushun109/tku/backend/internal/usecase/creator"
)

const maxCreatorLogoSize = 20 << 20 // 20MB

type CreatorHandler struct {
	creatorUC usecaseCreator.Usecase
}

func NewCreatorHandler(creatorUC usecaseCreator.Usecase) *CreatorHandler {
	return &CreatorHandler{creatorUC: creatorUC}
}

func (h *CreatorHandler) Get(w http.ResponseWriter, r *http.Request) {
	detail, err := h.creatorUC.Get(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToCreatorResponse(detail)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *CreatorHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateCreatorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.creatorUC.Update(r.Context(), req.Name, req.Introduction); err != nil {
		response.WriteAppError(w, err)
		return
	}
	response.WriteSuccess(w)
}

func (h *CreatorHandler) UpdateLogo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("logo")
	if err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}
	defer func() {
		_ = file.Close()
	}()

	logoBytes, err := io.ReadAll(io.LimitReader(file, maxCreatorLogoSize+1))
	if err != nil {
		response.WriteAppError(w, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error()))
		return
	}
	if len(logoBytes) == 0 || len(logoBytes) > maxCreatorLogoSize {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	if err := h.creatorUC.UpdateLogo(r.Context(), logoBytes); err != nil {
		response.WriteAppError(w, err)
		return
	}
	response.WriteSuccess(w)
}

func (h *CreatorHandler) GetLogoBlob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logoFile := vars["logo_file"]

	logoBlob, err := h.creatorUC.GetLogoBlob(r.Context(), logoFile)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}
	defer func() {
		_ = logoBlob.Body.Close()
	}()

	w.Header().Set("Content-Type", logoBlob.ContentType)
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, logoBlob.Body)
}
