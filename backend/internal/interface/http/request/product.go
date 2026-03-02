package request

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

const (
	defaultCategoryProductLimit = 4
	maxCategoryProductLimit     = 20
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
	UUID         string `json:"uuid"`
	DisplayOrder int    `json:"displayOrder"`
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
	Category string
	Cursor   string
	Limit    int
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
	category := strings.TrimSpace(q.Get("category"))
	cursor := strings.TrimSpace(q.Get("cursor"))
	limit := defaultCategoryProductLimit
	target := strings.TrimSpace(q.Get("target"))

	if category == "" || target == "" {
		return ListCategoryProductQuery{}, errors.New("invalid query")
	}

	if rawLimit := strings.TrimSpace(q.Get("limit")); rawLimit != "" {
		parsedLimit, err := strconv.Atoi(rawLimit)
		if err != nil || parsedLimit <= 0 || parsedLimit > maxCategoryProductLimit {
			return ListCategoryProductQuery{}, errors.New("invalid limit")
		}
		limit = parsedLimit
	}

	return ListCategoryProductQuery{
		Category: category,
		Cursor:   cursor,
		Limit:    limit,
		Target:   target,
	}, nil
}
