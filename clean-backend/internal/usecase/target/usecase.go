package target

import (
	"context"
	"errors"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context, mode string) ([]*domain.Target, error)
	Create(ctx context.Context, name string) error
	Update(ctx context.Context, uuid string, name string) error
	Delete(ctx context.Context, uuid string) error
}

type Service struct {
	repo    domain.Repository
	uuidGen usecase.UUIDGenerator
}

const (
	ListModeAll  = "all"
	ListModeUsed = "used"
)

func New(repo domain.Repository, uuidGen usecase.UUIDGenerator) *Service {
	return &Service{repo: repo, uuidGen: uuidGen}
}

func (s *Service) List(ctx context.Context, mode string) ([]*domain.Target, error) {
	switch mode {
	case ListModeAll:
		targets, err := s.repo.FindAll(ctx)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return targets, nil
	case ListModeUsed:
		targets, err := s.repo.FindUsed(ctx)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return targets, nil
	default:
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}
}

func (s *Service) Create(ctx context.Context, name string) error {
	newUUID := s.uuidGen.New()

	t, err := domain.New(newUUID, name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	exists, err := s.repo.ExistsByName(ctx, t.Name())
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if exists {
		return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
	}

	if _, err := s.repo.Create(ctx, t); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}

func (s *Service) Update(ctx context.Context, uuidStr string, name string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	newName, err := domain.NewTargetName(name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	current, err := s.repo.FindByUUID(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if current == nil {
		return usecase.NewAppError(usecase.ErrNotFound)
	}

	if current.Name() != newName {
		exists, err := s.repo.ExistsByName(ctx, newName)
		if err != nil {
			return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		if exists {
			return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
		}
	}

	if err := current.ChangeName(name); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	updated, err := s.repo.Update(ctx, current)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !updated {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, uuidStr string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}

	deleted, err := s.repo.Delete(ctx, uuid)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !deleted {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	return nil
}
