package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainSession "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()
var testUUIDVO = mustNewUUID(testUUID)

type stubUserRepo struct {
	userByEmail    *domainUser.User
	userByEmailErr error
	userByID       *domainUser.User
	userByIDErr    error
}

func (s *stubUserRepo) FindByEmail(ctx context.Context, email string) (*domainUser.User, error) {
	if s.userByEmailErr != nil {
		return nil, s.userByEmailErr
	}
	return s.userByEmail, nil
}

func (s *stubUserRepo) FindByID(ctx context.Context, id uint) (*domainUser.User, error) {
	if s.userByIDErr != nil {
		return nil, s.userByIDErr
	}
	return s.userByID, nil
}

type stubSessionRepo struct {
	createErr       error
	findByUUID      *domainSession.Session
	findByUUIDErr   error
	deleteByUUIDErr error
	deleteByUserErr error
	created         *domainSession.Session
	deletedUserID   uint
}

func (s *stubSessionRepo) Create(ctx context.Context, sess *domainSession.Session) error {
	if s.createErr != nil {
		return s.createErr
	}
	s.created = sess
	return nil
}

func (s *stubSessionRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainSession.Session, error) {
	if s.findByUUIDErr != nil {
		return nil, s.findByUUIDErr
	}
	return s.findByUUID, nil
}

func (s *stubSessionRepo) DeleteByUUID(ctx context.Context, uuid primitive.UUID) error {
	return s.deleteByUUIDErr
}

func (s *stubSessionRepo) DeleteByUserID(ctx context.Context, userID uint) error {
	s.deletedUserID = userID
	return s.deleteByUserErr
}

type stubSessionUC struct {
	resolveRes *domainSession.Session
	resolveErr error
	deleteErr  error
}

func (s *stubSessionUC) Validate(ctx context.Context, token string) error {
	return s.resolveErr
}

func (s *stubSessionUC) Resolve(ctx context.Context, token string) (*domainSession.Session, error) {
	if s.resolveErr != nil {
		return nil, s.resolveErr
	}
	return s.resolveRes, nil
}

func (s *stubSessionUC) Delete(ctx context.Context, token string) error {
	return s.deleteErr
}

type stubPasswordHasher struct {
	verifyOK  bool
	verifyErr error
}

func (s *stubPasswordHasher) Hash(plain string) (domainUser.UserPasswordHash, error) {
	return domainUser.NewUserPasswordHash("hash")
}

func (s *stubPasswordHasher) Verify(plain string, hash domainUser.UserPasswordHash) (bool, error) {
	if s.verifyErr != nil {
		return false, s.verifyErr
	}
	return s.verifyOK, nil
}

type stubUUIDGen struct {
	uuid primitive.UUID
	err  error
}

func (s *stubUUIDGen) New() (primitive.UUID, error) {
	if s.err != nil {
		return "", s.err
	}
	return s.uuid, nil
}

type stubClock struct {
	now time.Time
}

func (c *stubClock) Now() time.Time {
	return c.now
}

func TestLogin_OK(t *testing.T) {
	hash, err := domainUser.NewUserPasswordHash("hash")
	if err != nil {
		t.Fatalf("unexpected hash error: %v", err)
	}
	userRepo := &stubUserRepo{userByEmail: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
	sessionRepo := &stubSessionRepo{}
	sessionUC := &stubSessionUC{}
	hasher := &stubPasswordHasher{verifyOK: true}
	clock := &stubClock{now: time.Date(2026, 2, 17, 10, 0, 0, 0, time.UTC)}

	uc := New(userRepo, sessionRepo, sessionUC, hasher, &stubUUIDGen{uuid: testUUIDVO}, clock)

	sess, err := uc.Login(context.Background(), "mail@example.com", "password")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sess == nil || sess.UUID.String() != testUUID {
		t.Fatalf("unexpected session: %+v", sess)
	}
	if sessionRepo.deletedUserID != 1 {
		t.Fatalf("expected deleted user id=1, got %d", sessionRepo.deletedUserID)
	}
	if sessionRepo.created == nil {
		t.Fatalf("expected session create called")
	}
}

func TestLogin_UserNotFound_Unauthorized(t *testing.T) {
	userRepo := &stubUserRepo{userByEmail: nil}
	uc := New(userRepo, &stubSessionRepo{}, &stubSessionUC{}, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUIDVO}, &stubClock{now: time.Now()})

	_, err := uc.Login(context.Background(), "mail@example.com", "password")
	if err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestLogin_WrongPassword_Unauthorized(t *testing.T) {
	hash, err := domainUser.NewUserPasswordHash("hash")
	if err != nil {
		t.Fatalf("unexpected hash error: %v", err)
	}
	userRepo := &stubUserRepo{userByEmail: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
	uc := New(userRepo, &stubSessionRepo{}, &stubSessionUC{}, &stubPasswordHasher{verifyOK: false}, &stubUUIDGen{uuid: testUUIDVO}, &stubClock{now: time.Now()})

	_, err = uc.Login(context.Background(), "mail@example.com", "bad")
	if err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestGetBySessionToken_OK(t *testing.T) {
	hash, err := domainUser.NewUserPasswordHash("hash")
	if err != nil {
		t.Fatalf("unexpected hash error: %v", err)
	}
	userRepo := &stubUserRepo{userByID: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
	sessionUC := &stubSessionUC{resolveRes: &domainSession.Session{UserID: 1}}
	uc := New(userRepo, &stubSessionRepo{}, sessionUC, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUIDVO}, &stubClock{now: time.Now()})

	u, err := uc.GetBySessionToken(context.Background(), testUUID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u == nil || u.ID != 1 {
		t.Fatalf("unexpected user: %+v", u)
	}
}

func TestGetBySessionToken_Unauthorized(t *testing.T) {
	sessionUC := &stubSessionUC{resolveErr: usecase.NewAppError(usecase.ErrUnauthorized)}
	uc := New(&stubUserRepo{}, &stubSessionRepo{}, sessionUC, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUIDVO}, &stubClock{now: time.Now()})

	_, err := uc.GetBySessionToken(context.Background(), "bad")
	if err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestLogout_OK(t *testing.T) {
	sessionUC := &stubSessionUC{}
	uc := New(&stubUserRepo{}, &stubSessionRepo{}, sessionUC, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUIDVO}, &stubClock{now: time.Now()})

	if err := uc.Logout(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func mustUser(id uint, uuidStr string, name string, email string, hash domainUser.UserPasswordHash, isAdmin bool) *domainUser.User {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	return &domainUser.User{
		ID:           id,
		UUID:         uuid,
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		IsAdmin:      isAdmin,
	}
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}
