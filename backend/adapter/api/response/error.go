package response

import (
	"encoding/json"
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/api/logging"
	"github.com/tokushun109/tku/backend/adapter/logger"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Message string `json:"message"`
}

type Error struct {
	message string
	status  int
}

func newError(message string, status int) Error {
	return Error{message: message, status: status}
}

func (e Error) send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error: ErrorBody{Message: e.message},
	})
}

func LogAndSendError(
	w http.ResponseWriter,
	r *http.Request,
	log logger.Logger,
	status int,
	err error,
	msg string,
) {
	logging.NewError(log, r, status, err).Log(msg)
	newError(msg, status).send(w)
}
