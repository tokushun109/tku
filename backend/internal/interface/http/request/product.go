package request

import (
	"errors"
	"net/http"
	"strings"

	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

type ProductClassificationRequest struct {
	UUID string `json:"uuid"`
}

type ProductTagRequest struct {
	UUID string `json:"uuid"`
}

type ProductSalesSiteRequest struct {
	UUID string `json:"uuid"`
}

type ProductSiteDetailRequest struct {
	UUID      string                  `json:"uuid"`
	DetailURL string                  `json:"detailUrl"`
	SalesSite ProductSalesSiteRequest `json:"salesSite"`
}

type ProductImageRequest struct {
	UUID  string `json:"uuid"`
	Order int    `json:"order"`
}

type CreateProductRequest struct {
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	Price         int                          `json:"price"`
	IsRecommend   bool                         `json:"isRecommend"`
	IsActive      bool                         `json:"isActive"`
	Category      ProductClassificationRequest `json:"category"`
	Target        ProductClassificationRequest `json:"target"`
	Tags          []ProductTagRequest          `json:"tags"`
	ProductImages []ProductImageRequest        `json:"productImages"`
	SiteDetails   []ProductSiteDetailRequest   `json:"siteDetails"`
}

type UpdateProductRequest struct {
	UUID          string                       `json:"uuid"`
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	Price         int                          `json:"price"`
	IsRecommend   bool                         `json:"isRecommend"`
	IsActive      bool                         `json:"isActive"`
	Category      ProductClassificationRequest `json:"category"`
	Target        ProductClassificationRequest `json:"target"`
	Tags          []ProductTagRequest          `json:"tags"`
	ProductImages []ProductImageRequest        `json:"productImages"`
	SiteDetails   []ProductSiteDetailRequest   `json:"siteDetails"`
}

type ListProductQuery struct {
	Mode     string
	Category string
	Target   string
}

type ListCategoryProductQuery struct {
	Mode     string
	Category string
	Target   string
}

func ParseListProductQuery(r *http.Request) (ListProductQuery, error) {
	q := r.URL.Query()
	mode := strings.TrimSpace(q.Get("mode"))
	category := strings.TrimSpace(q.Get("category"))
	target := strings.TrimSpace(q.Get("target"))

	switch mode {
	case usecaseProduct.ListModeAll, usecaseProduct.ListModeActive:
		if category == "" || target == "" {
			return ListProductQuery{}, errors.New("invalid query")
		}
		return ListProductQuery{Mode: mode, Category: category, Target: target}, nil
	default:
		return ListProductQuery{}, errors.New("invalid mode")
	}
}

func ParseListCategoryProductQuery(r *http.Request) (ListCategoryProductQuery, error) {
	q := r.URL.Query()
	mode := strings.TrimSpace(q.Get("mode"))
	category := strings.TrimSpace(q.Get("category"))
	target := strings.TrimSpace(q.Get("target"))

	switch mode {
	case usecaseProduct.ListModeAll, usecaseProduct.ListModeActive:
		if category == "" || target == "" {
			return ListCategoryProductQuery{}, errors.New("invalid query")
		}
		return ListCategoryProductQuery{Mode: mode, Category: category, Target: target}, nil
	default:
		return ListCategoryProductQuery{}, errors.New("invalid mode")
	}
}
