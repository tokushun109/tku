package request

import (
	"errors"
	"net/http"
)

type CreateTargetRequest struct {
	Name string `json:"name"`
}

type UpdateTargetRequest struct {
	Name string `json:"name"`
}

type ListTargetQuery struct {
	Mode string
}

const (
	TargetModeAll  = "all"
	TargetModeUsed = "used"
)

func ParseListTargetQuery(r *http.Request) (ListTargetQuery, error) {
	q := r.URL.Query()
	mode := q.Get("mode")
	switch mode {
	case TargetModeAll, TargetModeUsed:
		return ListTargetQuery{Mode: mode}, nil
	default:
		return ListTargetQuery{}, errors.New("invalid mode")
	}
}
