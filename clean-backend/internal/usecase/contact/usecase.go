package contact

import (
	"context"
	"errors"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/contact"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Usecase interface {
	List(ctx context.Context) ([]*domain.Contact, error)
	Create(ctx context.Context, name string, company string, phoneNumber string, email string, content string) error
}

type Service struct {
	repo     domain.Repository
	notifier Notifier
}

func New(repo domain.Repository, notifier Notifier) *Service {
	if notifier == nil {
		notifier = nopNotifier{}
	}
	return &Service{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *Service) List(ctx context.Context) ([]*domain.Contact, error) {
	contactList, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	return contactList, nil
}

func (s *Service) Create(ctx context.Context, name string, company string, phoneNumber string, email string, content string) error {
	contact, err := domain.New(name, company, phoneNumber, email, content)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) ||
			errors.Is(err, domain.ErrInvalidCompany) ||
			errors.Is(err, primitive.ErrInvalidPhoneNumber) ||
			errors.Is(err, primitive.ErrInvalidEmail) ||
			errors.Is(err, domain.ErrInvalidContent) {
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		}
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if err := s.repo.Create(ctx, contact); err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	s.notifier.NotifyContactCreated(ctx, contact)
	return nil
}
