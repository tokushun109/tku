package category

import (
	"context"
	"errors"

	domain "github.com/tokushun109/tku/backend/internal/domain/category"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context, mode string) ([]*domain.Category, error)
	Create(ctx context.Context, name string) error
	Update(ctx context.Context, uuid string, name string) error
	Delete(ctx context.Context, uuid string) error
}

type Service struct {
	repo      domain.Repository
	uuidGen   usecase.UUIDGenerator
	txManager usecase.TxManager
}

const (
	ListModeAll  = "all"
	ListModeUsed = "used"
)

func New(repo domain.Repository, uuidGen usecase.UUIDGenerator, txManager usecase.TxManager) *Service {
	return &Service{repo: repo, uuidGen: uuidGen, txManager: txManager}
}

func (s *Service) List(ctx context.Context, mode string) ([]*domain.Category, error) {
	switch mode {
	case ListModeAll:
		categories, err := s.repo.FindAll(ctx)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return categories, nil
	case ListModeUsed:
		categories, err := s.repo.FindUsed(ctx)
		if err != nil {
			return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
		return categories, nil
	default:
		return nil, usecase.NewAppError(usecase.ErrInvalidInput)
	}
}

func (s *Service) Create(ctx context.Context, name string) error {
	newUUID := s.uuidGen.New()

	c, err := domain.New(newUUID, name)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	exists, err := s.repo.ExistsByName(ctx, c.Name())
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if exists {
		return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
	}

	if _, err := s.repo.Create(ctx, c); err != nil {
		if errors.Is(err, domain.ErrNameDuplicated) {
			return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return nil
}

func (s *Service) Update(ctx context.Context, uuidStr string, name string) error {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return usecase.NewAppError(usecase.ErrInvalidInput)
	}
	newName, err := domain.NewCategoryName(name)
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
		if errors.Is(err, domain.ErrNameDuplicated) {
			return usecase.NewAppErrorWithMessage(usecase.ErrConflict, domain.ErrNameDuplicated.Error())
		}
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

	var deleted bool
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		deleted, err = s.repo.Delete(txCtx, uuid)
		return err
	}); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !deleted {
		return usecase.NewAppError(usecase.ErrNotFound)
	}
	return nil
}
