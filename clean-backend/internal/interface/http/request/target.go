package request

import (
	"errors"
	"net/http"

	usecaseTarget "github.com/tokushun109/tku/clean-backend/internal/usecase/target"
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

func ParseListTargetQuery(r *http.Request) (ListTargetQuery, error) {
	q := r.URL.Query()
	mode := q.Get("mode")
	switch mode {
	case usecaseTarget.ListModeAll, usecaseTarget.ListModeUsed:
		return ListTargetQuery{Mode: mode}, nil
	default:
		return ListTargetQuery{}, errors.New("invalid mode")
	}
}
