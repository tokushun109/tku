package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseContact "github.com/tokushun109/tku/backend/internal/usecase/contact"
)

type ContactHandler struct {
	contactUC usecaseContact.Usecase
}

func NewContactHandler(contactUC usecaseContact.Usecase) *ContactHandler {
	return &ContactHandler{contactUC: contactUC}
}

func (h *ContactHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.contactUC.List(r.Context())
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToContactResponses(list)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ContactHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	company := ""
	if req.Company != nil {
		company = *req.Company
	}

	phoneNumber := ""
	if req.PhoneNumber != nil {
		phoneNumber = *req.PhoneNumber
	}

	if err := h.contactUC.Create(r.Context(), req.Name, company, phoneNumber, req.Email, req.Content); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
