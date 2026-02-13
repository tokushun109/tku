package response

import (
	"errors"
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

func MapError(err error) (status int, msg string) {
	switch {
	case errors.Is(err, usecase.ErrInvalidInput):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, usecase.ErrNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, usecase.ErrConflict):
		return http.StatusConflict, err.Error()
	case errors.Is(err, usecase.ErrUnauthorized):
		return http.StatusUnauthorized, err.Error()
	default:
		return http.StatusInternalServerError, usecase.ErrInternal.Error()
	}
}
