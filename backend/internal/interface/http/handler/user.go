package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tokushun109/tku/backend/internal/interface/http/middleware"
	"github.com/tokushun109/tku/backend/internal/interface/http/presenter"
	"github.com/tokushun109/tku/backend/internal/interface/http/request"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseUser "github.com/tokushun109/tku/backend/internal/usecase/user"
)

type UserHandler struct {
	userUC usecaseUser.Usecase
}

func NewUserHandler(userUC usecaseUser.Usecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrInvalidInput))
		return
	}

	sess, err := h.userUC.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		response.WriteAppError(w, err)
		return
	}

	res := presenter.ToLoginSessionResponse(sess)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	authUser, ok := middleware.AuthenticatedUserFromContext(r.Context())
	if !ok {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrUnauthorized))
		return
	}

	res := &response.LoginUserResponse{
		UUID:    authUser.UUID,
		Name:    authUser.Name,
		Email:   authUser.Email,
		IsAdmin: authUser.IsAdmin,
	}
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authUser, ok := middleware.AuthenticatedUserFromContext(r.Context())
	if !ok {
		response.WriteAppError(w, usecase.NewAppError(usecase.ErrUnauthorized))
		return
	}

	if err := h.userUC.Logout(r.Context(), authUser.SessionToken); err != nil {
		response.WriteAppError(w, err)
		return
	}

	response.WriteSuccess(w)
}
