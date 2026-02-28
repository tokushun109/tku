package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	domainCategory "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	domainSalesSite "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
	domainSiteDetail "github.com/tokushun109/tku/clean-backend/internal/domain/site_detail"
	domainTag "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	domainTarget "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/clean-backend/internal/usecase/product"
)

func (s *Service) resolveCategoryID(ctx context.Context, rawUUID string) (*uint, error) {
	trimmed := strings.TrimSpace(rawUUID)
	if trimmed == "" {
		return nil, nil
	}

	uuid, err := primitive.NewUUID(trimmed)
	if err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	category, err := s.categoryRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if category == nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	id := category.ID().Value()
	return &id, nil
}

func (s *Service) resolveTargetID(ctx context.Context, rawUUID string) (*uint, error) {
	trimmed := strings.TrimSpace(rawUUID)
	if trimmed == "" {
		return nil, nil
	}

	uuid, err := primitive.NewUUID(trimmed)
	if err != nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	target, err := s.targetRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if target == nil {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	id := target.ID().Value()
	return &id, nil
}

func (s *Service) resolveTagIDs(ctx context.Context, rawUUIDs []string) ([]primitive.ID, error) {
	if len(rawUUIDs) == 0 {
		return []primitive.ID{}, nil
	}

	seen := map[primitive.ID]struct{}{}
	tagIDs := make([]primitive.ID, 0, len(rawUUIDs))
	for _, rawUUID := range rawUUIDs {
		uuid, err := primitive.NewUUID(strings.TrimSpace(rawUUID))
		if err != nil {
			return nil, usecase.NewAppError(usecase.ErrInvalidInput)
		}

		tag, err := s.tagRepo.FindByUUID(ctx, uuid)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		if tag == nil {
			return nil, usecase.NewAppError(usecase.ErrInvalidInput)
		}
		tagID := tag.ID()
		if _, exists := seen[tagID]; exists {
			continue
		}
		seen[tagID] = struct{}{}
		tagIDs = append(tagIDs, tagID)
	}
	return tagIDs, nil
}

func (s *Service) resolveOrCreateTagIDsByNames(ctx context.Context, rawNames []string) ([]primitive.ID, error) {
	if len(rawNames) == 0 {
		return []primitive.ID{}, nil
	}

	seen := map[string]struct{}{}
	tagIDs := make([]primitive.ID, 0, len(rawNames))
	for _, rawName := range rawNames {
		tagName, err := domainTag.NewTagName(rawName)
		if err != nil {
			return nil, err
		}

		key := strings.ToLower(tagName.Value())
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}

		tag, err := s.tagRepo.FindByName(ctx, tagName)
		if err != nil {
			return nil, err
		}
		if tag == nil {
			newTag, err := domainTag.New(s.uuidGen.New(), tagName.Value())
			if err != nil {
				return nil, err
			}
			if err := s.tagRepo.Create(ctx, newTag); err != nil {
				return nil, err
			}

			tag, err = s.tagRepo.FindByName(ctx, tagName)
			if err != nil {
				return nil, err
			}
			if tag == nil {
				return nil, fmt.Errorf("created tag was not found: %s", tagName.Value())
			}
		}

		tagIDs = append(tagIDs, tag.ID())
	}

	return tagIDs, nil
}

type normalizedProductCSVRow struct {
	id           uint
	name         string
	price        int
	categoryName *domainCategory.CategoryName
	targetName   *domainTarget.TargetName
}

func normalizeProductCSVRows(rows []usecaseProduct.ProductCSVInputRow) ([]normalizedProductCSVRow, error) {
	result := make([]normalizedProductCSVRow, 0, len(rows))
	for i, row := range rows {
		if _, err := primitive.NewID(row.ID); err != nil {
			return nil, fmt.Errorf("row %d: product id is invalid", i+1)
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
			id:           row.ID,
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

func (s *Service) buildSiteDetails(ctx context.Context, productID primitive.ID, inputs []usecaseProduct.SiteDetailInput) ([]*domainSiteDetail.SiteDetail, error) {
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
			productID.Value(),
			salesSite.ID().Value(),
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
		errors.Is(err, domainProduct.ErrInvalidCategoryID) ||
		errors.Is(err, domainProduct.ErrInvalidTargetID) ||
		errors.Is(err, domainProduct.ErrInvalidImageName) ||
		errors.Is(err, domainProduct.ErrInvalidImageMimeType) ||
		errors.Is(err, domainProduct.ErrInvalidImagePath) ||
		errors.Is(err, domainProduct.ErrInvalidImageOrder) ||
		errors.Is(err, domainProduct.ErrInvalidImageProductID) ||
		errors.Is(err, domainSiteDetail.ErrInvalidDetailURL) ||
		errors.Is(err, domainSiteDetail.ErrInvalidProductID) ||
		errors.Is(err, domainSiteDetail.ErrInvalidSalesSiteID)
}
