package user

import (
	"context"
	"strings"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainSession "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseSession "github.com/tokushun109/tku/clean-backend/internal/usecase/session"
)

type Usecase interface {
	GetBySessionToken(ctx context.Context, token string) (*domainUser.User, error)
	Login(ctx context.Context, email string, password string) (*domainSession.Session, error)
	Logout(ctx context.Context, token string) error
}

type Service struct {
	userRepo       domainUser.Repository
	sessionRepo    domainSession.Repository
	sessionUC      usecaseSession.Usecase
	passwordHasher PasswordHasher
	uuidGen        usecase.UUIDGenerator
	clock          usecase.Clock
	txManager      usecase.TxManager
}

func New(
	userRepo domainUser.Repository,
	sessionRepo domainSession.Repository,
	sessionUC usecaseSession.Usecase,
	passwordHasher PasswordHasher,
	uuidGen usecase.UUIDGenerator,
	clock usecase.Clock,
	txManager usecase.TxManager,
) *Service {
	return &Service{
		userRepo:       userRepo,
		sessionRepo:    sessionRepo,
		sessionUC:      sessionUC,
		passwordHasher: passwordHasher,
		uuidGen:        uuidGen,
		clock:          clock,
		txManager:      txManager,
	}
}

func (s *Service) Login(ctx context.Context, email string, password string) (*domainSession.Session, error) {
	trimmedEmail := strings.TrimSpace(email)
	if trimmedEmail == "" || password == "" {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}
	emailVO, err := primitive.NewEmail(trimmedEmail)
	if err != nil {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}

	u, err := s.userRepo.FindByEmail(ctx, emailVO)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if u == nil {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}

	ok, err := s.passwordHasher.Verify(password, u.PasswordHash)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !ok {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}

	sessionUUID, err := s.uuidGen.New()
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	sess := &domainSession.Session{
		UUID:      sessionUUID,
		UserID:    u.ID,
		CreatedAt: s.clock.Now(),
	}
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.sessionRepo.DeleteByUserID(txCtx, u.ID); err != nil {
			return err
		}
		if err := s.sessionRepo.Create(txCtx, sess); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}

	return sess, nil
}

func (s *Service) GetBySessionToken(ctx context.Context, token string) (*domainUser.User, error) {
	sess, err := s.sessionUC.Resolve(ctx, token)
	if err != nil {
		return nil, err
	}

	u, err := s.userRepo.FindByID(ctx, sess.UserID)
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if u == nil {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}

	return u, nil
}

func (s *Service) Logout(ctx context.Context, token string) error {
	return s.sessionUC.Delete(ctx, token)
}
