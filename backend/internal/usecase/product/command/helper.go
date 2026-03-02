package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	domainCategory "github.com/tokushun109/tku/backend/internal/domain/category"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/backend/internal/domain/product"
	domainSalesSite "github.com/tokushun109/tku/backend/internal/domain/sales_site"
	domainSiteDetail "github.com/tokushun109/tku/backend/internal/domain/site_detail"
	domainTarget "github.com/tokushun109/tku/backend/internal/domain/target"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

type normalizedProductCSVRow struct {
	uuid         string
	name         string
	price        int
	categoryName *domainCategory.CategoryName
	targetName   *domainTarget.TargetName
}

func normalizeProductCSVRows(rows []usecaseProduct.ProductCSVInputRow) ([]normalizedProductCSVRow, error) {
	result := make([]normalizedProductCSVRow, 0, len(rows))
	for i, row := range rows {
		productUUID, err := primitive.NewUUID(row.UUID)
		if err != nil {
			return nil, fmt.Errorf("row %d: product uuid is invalid", i+1)
		}

		productName, err := domainProduct.NewProductName(row.Name)
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", i+1, err)
		}

		productPrice, err := domainProduct.NewProductPrice(row.Price)
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", i+1, err)
		}

		var categoryName *domainCategory.CategoryName
		if strings.TrimSpace(row.CategoryName) != "" {
			parsedCategoryName, err := domainCategory.NewCategoryName(row.CategoryName)
			if err != nil {
				return nil, fmt.Errorf("row %d: %w", i+1, err)
			}
			categoryName = &parsedCategoryName
		}

		var targetName *domainTarget.TargetName
		if strings.TrimSpace(row.TargetName) != "" {
			parsedTargetName, err := domainTarget.NewTargetName(row.TargetName)
			if err != nil {
				return nil, fmt.Errorf("row %d: %w", i+1, err)
			}
			targetName = &parsedTargetName
		}

		result = append(result, normalizedProductCSVRow{
			uuid:         productUUID.Value(),
			name:         productName.Value(),
			price:        productPrice.Value(),
			categoryName: categoryName,
			targetName:   targetName,
		})
	}

	return result, nil
}

func (s *Service) findOrCreateCategoryByName(
	ctx context.Context,
	name domainCategory.CategoryName,
	cache map[string]*domainCategory.Category,
) (*domainCategory.Category, error) {
	key := name.Value()

	if cached, ok := cache[key]; ok {
		return cached, nil
	}

	found, err := s.categoryRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if found != nil {
		cache[key] = found
		return found, nil
	}

	newCategory, err := domainCategory.New(s.uuidGen.New(), key)
	if err != nil {
		return nil, err
	}
	created, err := s.categoryRepo.Create(ctx, newCategory)
	if err != nil {
		if errors.Is(err, domainCategory.ErrNameDuplicated) {
			found, findErr := s.categoryRepo.FindByName(ctx, name)
			if findErr != nil {
				return nil, findErr
			}
			if found != nil {
				cache[key] = found
				return found, nil
			}
		}
		return nil, err
	}
	if created == nil {
		return nil, fmt.Errorf("created category was not found")
	}

	cache[key] = created
	return created, nil
}

func (s *Service) findOrCreateTargetByName(
	ctx context.Context,
	name domainTarget.TargetName,
	cache map[string]*domainTarget.Target,
) (*domainTarget.Target, error) {
	key := name.Value()

	if cached, ok := cache[key]; ok {
		return cached, nil
	}

	found, err := s.targetRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if found != nil {
		cache[key] = found
		return found, nil
	}

	newTarget, err := domainTarget.New(s.uuidGen.New(), key)
	if err != nil {
		return nil, err
	}
	created, err := s.targetRepo.Create(ctx, newTarget)
	if err != nil {
		if errors.Is(err, domainTarget.ErrNameDuplicated) {
			found, findErr := s.targetRepo.FindByName(ctx, name)
			if findErr != nil {
				return nil, findErr
			}
			if found != nil {
				cache[key] = found
				return found, nil
			}
		}
		return nil, err
	}
	if created == nil {
		return nil, fmt.Errorf("created target was not found")
	}

	cache[key] = created
	return created, nil
}

func (s *Service) findSalesSiteByName(ctx context.Context, rawName string) (*domainSalesSite.SalesSite, error) {
	name, err := domainSalesSite.NewSalesSiteName(rawName)
	if err != nil {
		return nil, err
	}
	return s.salesSiteRepo.FindByName(ctx, name)
}

func (s *Service) buildSiteDetails(ctx context.Context, productUUID primitive.UUID, inputs []usecaseProduct.SiteDetailInput) ([]*domainSiteDetail.SiteDetail, error) {
	details := make([]*domainSiteDetail.SiteDetail, 0, len(inputs))
	for _, input := range inputs {
		salesSiteUUID := strings.TrimSpace(input.SalesSiteUUID)
		if salesSiteUUID == "" {
			return nil, usecase.ErrInvalidInput
		}
		uuid, err := primitive.NewUUID(salesSiteUUID)
		if err != nil {
			return nil, usecase.ErrInvalidInput
		}

		salesSite, err := s.salesSiteRepo.FindByUUID(ctx, uuid)
		if err != nil {
			return nil, err
		}
		if salesSite == nil {
			return nil, usecase.ErrInvalidInput
		}

		siteDetail, err := domainSiteDetail.New(
			s.uuidGen.New(),
			input.DetailURL,
			productUUID.Value(),
			salesSiteUUID,
		)
		if err != nil {
			return nil, err
		}
		details = append(details, siteDetail)
	}
	return details, nil
}

func normalizeDuplicateProductURL(rawURL string) (string, error) {
	trimmed := strings.TrimSpace(rawURL)
	if _, err := primitive.NewURL(trimmed); err != nil {
		return "", err
	}

	return trimmed, nil
}

func buildProductImagePath(uuidStr string, mimeType domainProduct.ProductImageMimeType) (domainProduct.ProductImagePath, error) {
	if len(uuidStr) < 2 {
		return "", fmt.Errorf("invalid uuid length: %d", len(uuidStr))
	}

	rawPath := fmt.Sprintf(
		"img/product/%s/%s/%s%s",
		uuidStr[0:1],
		uuidStr[1:2],
		uuidStr,
		mimeType.Extension(),
	)

	return domainProduct.NewProductImagePath(rawPath)
}

func isInvalidProductInputError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, primitive.ErrInvalidUUID) ||
		errors.Is(err, primitive.ErrInvalidURL) ||
		errors.Is(err, domainProduct.ErrInvalidName) ||
		errors.Is(err, domainProduct.ErrInvalidPrice) ||
		errors.Is(err, domainProduct.ErrInvalidCategoryUUID) ||
		errors.Is(err, domainProduct.ErrInvalidTargetUUID) ||
		errors.Is(err, domainProduct.ErrInvalidImageName) ||
		errors.Is(err, domainProduct.ErrInvalidImageMimeType) ||
		errors.Is(err, domainProduct.ErrInvalidImagePath) ||
		errors.Is(err, domainProduct.ErrInvalidImageDisplayOrder) ||
		errors.Is(err, domainProduct.ErrInvalidImageProductUUID) ||
		errors.Is(err, domainSiteDetail.ErrInvalidDetailURL) ||
		errors.Is(err, domainSiteDetail.ErrInvalidProductUUID) ||
		errors.Is(err, domainSiteDetail.ErrInvalidSalesSiteUUID)
}
