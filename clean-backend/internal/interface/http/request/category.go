package request

import (
	"errors"
	"net/http"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type ListCategoryQuery struct {
	Mode string
}

const (
	CategoryModeAll  = "all"
	CategoryModeUsed = "used"
)

func ParseListCategoryQuery(r *http.Request) (ListCategoryQuery, error) {
	q := r.URL.Query()
	mode := q.Get("mode")
	switch mode {
	case CategoryModeAll, CategoryModeUsed:
		return ListCategoryQuery{Mode: mode}, nil
	default:
		return ListCategoryQuery{}, errors.New("invalid mode")
	}
}
