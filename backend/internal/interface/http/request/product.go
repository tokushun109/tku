package request

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

const (
	defaultCategoryProductLimit = 4
	defaultProductLimit         = 20
	maxCategoryProductLimit     = 20
	maxProductLimit             = 100
	maxProductKeywordLength     = 100
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
	Price         *int                         `json:"price"`
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
	Price         *int                         `json:"price"`
	IsRecommend   bool                         `json:"isRecommend"`
	IsActive      bool                         `json:"isActive"`
	Category      ProductClassificationRequest `json:"category"`
	Target        ProductClassificationRequest `json:"target"`
	Tags          []ProductTagRequest          `json:"tags"`
	ProductImages []ProductImageRequest        `json:"productImages"`
	SiteDetails   []ProductSiteDetailRequest   `json:"siteDetails"`
}

type ListProductQuery struct {
	Mode            string
	ActiveStatus    string
	Category        string
	Keyword         string
	Limit           int
	MaxPrice        *int
	MinPrice        *int
	Page            int
	RecommendStatus string
	TagUUIDs        []string
	Target          string
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
	activeStatus := strings.TrimSpace(q.Get("activeStatus"))
	category := strings.TrimSpace(q.Get("category"))
	keyword := strings.TrimSpace(q.Get("keyword"))
	limit := defaultProductLimit
	var maxPrice *int
	var minPrice *int
	page := 1
	recommendStatus := strings.TrimSpace(q.Get("recommendStatus"))
	tagUUIDs := parseCSVQuery(q.Get("tagUuids"))
	target := strings.TrimSpace(q.Get("target"))

	if utf8.RuneCountInString(keyword) > maxProductKeywordLength {
		return ListProductQuery{}, errors.New("invalid keyword")
	}

	if rawPage := strings.TrimSpace(q.Get("page")); rawPage != "" {
		parsedPage, err := strconv.Atoi(rawPage)
		if err != nil || parsedPage <= 0 {
			return ListProductQuery{}, errors.New("invalid page")
		}
		page = parsedPage
	}

	if rawLimit := strings.TrimSpace(q.Get("limit")); rawLimit != "" {
		parsedLimit, err := strconv.Atoi(rawLimit)
		if err != nil || parsedLimit <= 0 || parsedLimit > maxProductLimit {
			return ListProductQuery{}, errors.New("invalid limit")
		}
		limit = parsedLimit
	}

	if rawMinPrice := strings.TrimSpace(q.Get("minPrice")); rawMinPrice != "" {
		parsedMinPrice, err := strconv.Atoi(rawMinPrice)
		if err != nil || parsedMinPrice < 0 {
			return ListProductQuery{}, errors.New("invalid min price")
		}
		minPrice = &parsedMinPrice
	}

	if rawMaxPrice := strings.TrimSpace(q.Get("maxPrice")); rawMaxPrice != "" {
		parsedMaxPrice, err := strconv.Atoi(rawMaxPrice)
		if err != nil || parsedMaxPrice < 0 {
			return ListProductQuery{}, errors.New("invalid max price")
		}
		maxPrice = &parsedMaxPrice
	}

	if minPrice != nil && maxPrice != nil && *minPrice > *maxPrice {
		return ListProductQuery{}, errors.New("invalid price range")
	}

	switch mode {
	case usecaseProduct.ListModeAll, usecaseProduct.ListModeActive:
		if category == "" || target == "" {
			return ListProductQuery{}, errors.New("invalid query")
		}
		return ListProductQuery{
			Mode:            mode,
			ActiveStatus:    activeStatus,
			Category:        category,
			Keyword:         keyword,
			Limit:           limit,
			MaxPrice:        maxPrice,
			MinPrice:        minPrice,
			Page:            page,
			RecommendStatus: recommendStatus,
			TagUUIDs:        tagUUIDs,
			Target:          target,
		}, nil
	default:
		return ListProductQuery{}, errors.New("invalid mode")
	}
}

func parseCSVQuery(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return []string{}
	}

	parts := strings.Split(raw, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		value := strings.TrimSpace(part)
		if value == "" {
			continue
		}
		values = append(values, value)
	}
	return values
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
