package logging

import (
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type Error struct {
	log        logger.Logger
	r          *http.Request
	httpStatus int
	err        error
}

func NewError(log logger.Logger, r *http.Request, httpStatus int, err error) Error {
	return Error{
		log:        log,
		r:          r,
		httpStatus: httpStatus,
		err:        err,
	}
}

func (e Error) Log(msg string) {
	errMsg := ""
	if e.err != nil {
		errMsg = e.err.Error()
	}
	method := ""
	path := ""
	if e.r != nil {
		method = e.r.Method
		if e.r.URL != nil {
			path = e.r.URL.String()
		}
	}
	e.log.Errorf("method=%s path=%s status=%d error=%s msg=%s", method, path, e.httpStatus, errMsg, msg)
}
