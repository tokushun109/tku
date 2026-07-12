package query

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

const (
	listModeAll    = "all"
	listModeActive = "active"

	defaultProductImagePresignTTL = 30 * time.Minute
	defaultCarouselLimit          = 5
	maxProductKeywordLength       = 100
)

type Usecase interface {
	List(ctx context.Context, q ListProductsQuery) (*ProductPage, error)
	ListByCategory(ctx context.Context, q ListCategoryProductsQuery) ([]*CategoryProducts, error)
	ListCarousel(ctx context.Context) ([]*CarouselItem, error)
	Get(ctx context.Context, productUUID string) (*Product, error)
	ExportCSV(ctx context.Context) ([]*ProductCSVRow, error)
}

type Service struct {
	queryReader Reader
	storage     usecase.Storage
}

var _ Usecase = (*Service)(nil)

func New(queryReader Reader, storage usecase.Storage) *Service {
	return &Service{
		queryReader: queryReader,
		storage:     storage,
	}
}

func (s *Service) List(ctx context.Context, q ListProductsQuery) (*ProductPage, error) {
	trimmedCategory := strings.TrimSpace(q.Category)
	trimmedKeyword := strings.TrimSpace(q.Keyword)
	trimmedTarget := strings.TrimSpace(q.Target)

	if !isValidListMode(q.Mode) || trimmedCategory == "" || trimmedTarget == "" || q.Page <= 0 || q.Limit <= 0 {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}
	if len([]rune(trimmedKeyword)) > maxProductKeywordLength {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	productPage, err := s.queryReader.ListProducts(ctx, ListProductsQuery{
		Mode:     q.Mode,
		Category: trimmedCategory,
		Keyword:  trimmedKeyword,
		Limit:    q.Limit,
		Page:     q.Page,
		Target:   trimmedTarget,
	})
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if productPage == nil {
		return &ProductPage{}, nil
	}

	if err := s.attachPresignedImageURLs(ctx, productPage.Products); err != nil {
		return nil, err
	}

	return productPage, nil
}

func (s *Service) ListByCategory(ctx context.Context, q ListCategoryProductsQuery) ([]*CategoryProducts, error) {
	trimmedCategory := strings.TrimSpace(q.Category)
	trimmedCursor := strings.TrimSpace(q.Cursor)
	trimmedTarget := strings.TrimSpace(q.Target)

	if trimmedCategory == "" || trimmedTarget == "" || q.Limit <= 0 {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}
	if trimmedCategory == "all" && trimmedCursor != "" {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	query := ListCategoryProductsQuery{
		Category: trimmedCategory,
		Cursor:   trimmedCursor,
		Limit:    q.Limit,
		Target:   trimmedTarget,
	}

	categoryProducts, err := s.queryReader.ListCategoryProducts(ctx, query)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			return nil, usecase.NewAppError(usecase.ErrNotFound)
		}
		if errors.Is(err, ErrInvalidCursor) {
			return nil, usecase.NewAppError(usecase.ErrInvalidInput)
		}
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	for _, group := range categoryProducts {
		if group == nil {
			continue
		}
		if err := s.attachPresignedImageURLs(ctx, group.Products); err != nil {
			return nil, err
		}
	}

	return categoryProducts, nil
}

func (s *Service) ListCarousel(ctx context.Context) ([]*CarouselItem, error) {
	items, err := s.queryReader.ListCarouselItems(ctx, ListCarouselQuery{
		Limit: defaultCarouselLimit,
	})
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	products := make([]*Product, 0, len(items))
	for _, item := range items {
		if item == nil || item.Product == nil {
			continue
		}
		products = append(products, item.Product)
	}
	if err := s.attachPresignedImageURLs(ctx, products); err != nil {
		return nil, err
	}

	carouselItems := make([]*CarouselItem, 0, len(items))
	for _, item := range items {
		if item == nil || item.Product == nil || len(item.Product.ProductImages) == 0 {
			continue
		}

		item.APIPath = item.Product.ProductImages[0].APIPath
		carouselItems = append(carouselItems, item)
	}

	return carouselItems, nil
}

func (s *Service) Get(ctx context.Context, productUUID string) (*Product, error) {
	if _, err := primitive.NewUUID(productUUID); err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	product, err := s.queryReader.GetProductByUUID(ctx, productUUID)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if product == nil {
		return nil, usecase.NewAppError(usecase.ErrNotFound)
	}

	if err := s.attachPresignedImageURLs(ctx, []*Product{product}); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) ExportCSV(ctx context.Context) ([]*ProductCSVRow, error) {
	rows, err := s.queryReader.ExportProductsCSV(ctx, ExportProductsCSVQuery{})
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return rows, nil
}

func (s *Service) attachPresignedImageURLs(ctx context.Context, products []*Product) error {
	for _, product := range products {
		for i := range product.ProductImages {
			if strings.TrimSpace(product.ProductImages[i].Path) == "" {
				continue
			}
			url, err := s.storage.PresignGet(ctx, product.ProductImages[i].Path, defaultProductImagePresignTTL)
			if err != nil {
				return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
			}
			product.ProductImages[i].APIPath = url
		}
	}
	return nil
}

func isValidListMode(mode string) bool {
	switch mode {
	case listModeAll, listModeActive:
		return true
	default:
		return false
	}
}
