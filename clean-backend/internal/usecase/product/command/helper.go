package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainProduct "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	domainSiteDetail "github.com/tokushun109/tku/clean-backend/internal/domain/site_detail"
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
