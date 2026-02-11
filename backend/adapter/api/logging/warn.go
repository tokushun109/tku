package logging

import (
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type Warn struct {
	log        logger.Logger
	r          *http.Request
	err        error
	httpStatus int
}

func NewWarn(log logger.Logger, r *http.Request, httpStatus int, err error) Warn {
	return Warn{
		log:        log,
		r:          r,
		err:        err,
		httpStatus: httpStatus,
	}
}

func (w Warn) Log(msg string) {
	errMsg := ""
	if w.err != nil {
		errMsg = w.err.Error()
	}
	method := ""
	path := ""
	if w.r != nil {
		method = w.r.Method
		if w.r.URL != nil {
			path = w.r.URL.String()
		}
	}
	w.log.Warnf("method=%s path=%s http_status=%d error=%s msg=%s", method, path, w.httpStatus, errMsg, msg)
}
