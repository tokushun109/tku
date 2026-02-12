package logging

import (
	"net/http"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type Debug struct {
	log logger.Logger
	r   *http.Request
}

func NewDebug(log logger.Logger, r *http.Request) Debug {
	return Debug{
		log: log,
		r:   r,
	}
}

func (d Debug) Log(msg string) {
	method, path := getRequestMeta(d.r)
	d.log.Debugf("method=%s path=%s msg=%s", method, path, msg)
}
