package contact

import (
	"context"
	"errors"

	domain "github.com/tokushun109/tku/backend/internal/domain/contact"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	"github.com/tokushun109/tku/backend/internal/shared/logger"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type Notifier interface {
	NotifyContactCreated(ctx context.Context, content string) error
}

type Usecase interface {
	List(ctx context.Context) ([]*domain.Contact, error)
	Create(ctx context.Context, name string, company string, phoneNumber string, email string, content string) error
}

type Service struct {
	repo     domain.Repository
	notifier Notifier
	uuidGen  usecase.UUIDGenerator
}

func New(repo domain.Repository, notifier Notifier, uuidGen usecase.UUIDGenerator) *Service {
	return &Service{
		repo:     repo,
		notifier: notifier,
		uuidGen:  uuidGen,
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
	contact, err := domain.New(s.uuidGen.New(), name, company, phoneNumber, email, content)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidName),
			errors.Is(err, domain.ErrInvalidCompany),
			errors.Is(err, primitive.ErrInvalidPhoneNumber),
			errors.Is(err, primitive.ErrInvalidEmail),
			errors.Is(err, domain.ErrInvalidContent):
			return usecase.NewAppErrorWithMessage(usecase.ErrInvalidInput, err.Error())
		default:
			return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
		}
	}

	createdContact, err := s.repo.Create(ctx, contact)
	if err != nil {
		return usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	if s.notifier != nil {
		// 保存済みの問い合わせは、送信元が切断した場合も通知する。
		if err := s.notifier.NotifyContactCreated(context.WithoutCancel(ctx), createdContact.Content().Value()); err != nil {
			logger.Warnf("contact Slack notification failed: %v", err)
		}
	}
	return nil
}
