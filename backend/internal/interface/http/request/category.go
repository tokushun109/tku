package request

import (
	"errors"
	"net/http"

	usecaseCategory "github.com/tokushun109/tku/backend/internal/usecase/category"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

type ListCategoryQuery struct {
	Mode string
}

func ParseListCategoryQuery(r *http.Request) (ListCategoryQuery, error) {
	q := r.URL.Query()
	mode := q.Get("mode")
	switch mode {
	case usecaseCategory.ListModeAll, usecaseCategory.ListModeUsed:
		return ListCategoryQuery{Mode: mode}, nil
	default:
		return ListCategoryQuery{}, errors.New("invalid mode")
	}
}
