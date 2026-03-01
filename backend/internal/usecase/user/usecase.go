package user

import (
	"context"
	"strings"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainSession "github.com/tokushun109/tku/backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseSession "github.com/tokushun109/tku/backend/internal/usecase/session"
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

	ok, err := s.passwordHasher.Verify(password, u.PasswordHash())
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if !ok {
		return nil, usecase.NewAppError(usecase.ErrUnauthorized)
	}

	sessionUUID := s.uuidGen.New()

	sess, err := domainSession.New(sessionUUID, u.UUID().Value(), s.clock.Now())
	if err != nil {
		return nil, usecase.NewAppErrorWithMessage(usecase.ErrInternal, err.Error())
	}
	if err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.sessionRepo.DeleteByUserUUID(txCtx, u.UUID()); err != nil {
			return err
		}
		createdSession, err := s.sessionRepo.Create(txCtx, sess)
		if err != nil {
			return err
		}
		sess = createdSession
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

	sessionUserUUID := sess.UserUUID()
	u, err := s.userRepo.FindByUUID(ctx, sessionUserUUID)
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
