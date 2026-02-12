package logging

import (
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type Info struct {
	log        logger.Logger
	r          *http.Request
	httpStatus int
}

func NewInfo(log logger.Logger, r *http.Request, httpStatus int) Info {
	return Info{
		log:        log,
		r:          r,
		httpStatus: httpStatus,
	}
}

func (i Info) Log(msg string) {
	method, path := getRequestMeta(i.r)
	i.log.Infof("method=%s path=%s status=%d msg=%s", method, path, i.httpStatus, msg)
}
