package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainTag "github.com/tokushun109/tku/backend/internal/domain/tag"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

func (s *Service) resolveCategoryUUID(ctx context.Context, rawUUID string) (*string, error) {
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

	return &trimmed, nil
}

func (s *Service) resolveTargetUUID(ctx context.Context, rawUUID string) (*string, error) {
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

	return &trimmed, nil
}

func (s *Service) resolveTagUUIDs(ctx context.Context, rawUUIDs []string) ([]primitive.UUID, error) {
	if len(rawUUIDs) == 0 {
		return []primitive.UUID{}, nil
	}

	seen := map[primitive.UUID]struct{}{}
	tagUUIDs := make([]primitive.UUID, 0, len(rawUUIDs))
	for _, rawUUID := range rawUUIDs {
		uuid, err := primitive.NewUUID(strings.TrimSpace(rawUUID))
		if err != nil {
			return nil, usecase.NewAppError(usecase.ErrInvalidInput)
		}
		if _, exists := seen[uuid]; exists {
			continue
		}
		seen[uuid] = struct{}{}
		tagUUIDs = append(tagUUIDs, uuid)
	}

	tags, err := s.tagRepo.FindByUUIDs(ctx, tagUUIDs)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if len(tags) != len(tagUUIDs) {
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}

	return tagUUIDs, nil
}

func (s *Service) resolveOrCreateTagUUIDsByNames(ctx context.Context, rawNames []string) ([]primitive.UUID, error) {
	if len(rawNames) == 0 {
		return []primitive.UUID{}, nil
	}

	seen := map[string]struct{}{}
	tagUUIDs := make([]primitive.UUID, 0, len(rawNames))
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
			tag, err = s.tagRepo.Create(ctx, newTag)
			if err != nil {
				return nil, err
			}
			if tag == nil {
				return nil, fmt.Errorf("created tag was not returned: %s", tagName.Value())
			}
		}

		tagUUIDs = append(tagUUIDs, tag.UUID())
	}

	return tagUUIDs, nil
}
