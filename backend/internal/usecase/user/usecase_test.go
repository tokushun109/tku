package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainSession "github.com/tokushun109/tku/backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

var testUUID = id.GenerateUUID()

type stubUserRepo struct {
	userByEmail    *domainUser.User
	userByEmailErr error
	userByID       *domainUser.User
	userByIDErr    error
}

func (s *stubUserRepo) FindByEmail(ctx context.Context, email primitive.Email) (*domainUser.User, error) {
	if s.userByEmailErr != nil {
		return nil, s.userByEmailErr
	}
	return s.userByEmail, nil
}

func (s *stubUserRepo) FindByID(ctx context.Context, id primitive.ID) (*domainUser.User, error) {
	if s.userByIDErr != nil {
		return nil, s.userByIDErr
	}
	return s.userByID, nil
}

func (s *stubUserRepo) FindContactNotificationUsers(ctx context.Context) ([]*domainUser.ContactNotificationUser, error) {
	return nil, nil
}

type stubSessionRepo struct {
	createErr       error
	findByUUID      *domainSession.Session
	findByUUIDErr   error
	deleteByUUIDErr error
	deleteByUserErr error
	created         *domainSession.Session
	deletedUserID   primitive.ID
}

func (s *stubSessionRepo) Create(ctx context.Context, sess *domainSession.Session) (*domainSession.Session, error) {
	if s.createErr != nil {
		return nil, s.createErr
	}
	s.created = sess
	return sess, nil
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

func (s *stubSessionRepo) DeleteByUserID(ctx context.Context, userID primitive.ID) error {
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
	uuid string
}

func (s *stubUUIDGen) New() string {
	return s.uuid
}

type stubClock struct {
	now time.Time
}

func (c *stubClock) Now() time.Time {
	return c.now
}

type stubTxManager struct {
	err       error
	called    int
	calledCtx context.Context
}

func (s *stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	s.called++
	s.calledCtx = ctx
	if s.err != nil {
		return s.err
	}
	return fn(ctx)
}

func TestLogin(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		hash, err := domainUser.NewUserPasswordHash("hash")
		if err != nil {
			t.Fatalf("unexpected hash error: %v", err)
		}
		userRepo := &stubUserRepo{userByEmail: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
		sessionRepo := &stubSessionRepo{}
		sessionUC := &stubSessionUC{}
		hasher := &stubPasswordHasher{verifyOK: true}
		clock := &stubClock{now: time.Date(2026, 2, 17, 10, 0, 0, 0, time.UTC)}
		txManager := &stubTxManager{}

		uc := New(userRepo, sessionRepo, sessionUC, hasher, &stubUUIDGen{uuid: testUUID}, clock, txManager)

		sess, err := uc.Login(context.Background(), "mail@example.com", "password")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if sess == nil || sess.UUID().Value() != testUUID {
			t.Fatalf("unexpected session: %+v", sess)
		}
		if sessionRepo.deletedUserID.Value() != 1 {
			t.Fatalf("expected deleted user id=1, got %d", sessionRepo.deletedUserID.Value())
		}
		if sessionRepo.created == nil {
			t.Fatalf("expected session create called")
		}
		if txManager.called != 1 {
			t.Fatalf("expected tx manager called once, got %d", txManager.called)
		}
	})
	t.Run("ユーザーが見つからないなら未認証エラーを返す", func(t *testing.T) {

		userRepo := &stubUserRepo{userByEmail: nil}
		uc := New(userRepo, &stubSessionRepo{}, &stubSessionUC{}, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUID}, &stubClock{now: time.Now()}, &stubTxManager{})

		_, err := uc.Login(context.Background(), "mail@example.com", "password")
		if err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
			t.Fatalf("expected ErrUnauthorized, got %v", err)
		}
	})
	t.Run("パスワードが一致しないなら未認証エラーを返す", func(t *testing.T) {

		hash, err := domainUser.NewUserPasswordHash("hash")
		if err != nil {
			t.Fatalf("unexpected hash error: %v", err)
		}
		userRepo := &stubUserRepo{userByEmail: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
		uc := New(userRepo, &stubSessionRepo{}, &stubSessionUC{}, &stubPasswordHasher{verifyOK: false}, &stubUUIDGen{uuid: testUUID}, &stubClock{now: time.Now()}, &stubTxManager{})

		_, err = uc.Login(context.Background(), "mail@example.com", "bad")
		if err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
			t.Fatalf("expected ErrUnauthorized, got %v", err)
		}
	})
	t.Run("TxManagerでエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		hash, err := domainUser.NewUserPasswordHash("hash")
		if err != nil {
			t.Fatalf("unexpected hash error: %v", err)
		}
		userRepo := &stubUserRepo{userByEmail: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
		uc := New(
			userRepo,
			&stubSessionRepo{},
			&stubSessionUC{},
			&stubPasswordHasher{verifyOK: true},
			&stubUUIDGen{uuid: testUUID},
			&stubClock{now: time.Now()},
			&stubTxManager{err: errors.New("tx failed")},
		)

		_, err = uc.Login(context.Background(), "mail@example.com", "password")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}
func TestGetBySessionToken(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		hash, err := domainUser.NewUserPasswordHash("hash")
		if err != nil {
			t.Fatalf("unexpected hash error: %v", err)
		}
		userRepo := &stubUserRepo{userByID: mustUser(1, testUUID, "name", "mail@example.com", hash, true)}
		sessionUC := &stubSessionUC{resolveRes: mustSession(1, testUUID, 1, time.Now())}
		uc := New(userRepo, &stubSessionRepo{}, sessionUC, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUID}, &stubClock{now: time.Now()}, &stubTxManager{})

		u, err := uc.GetBySessionToken(context.Background(), testUUID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u == nil || u.ID().Value() != 1 {
			t.Fatalf("unexpected user: %+v", u)
		}
	})
	t.Run("認証情報がないなら未認証エラーを返す", func(t *testing.T) {

		sessionUC := &stubSessionUC{resolveErr: usecase.NewAppError(usecase.ErrUnauthorized)}
		uc := New(&stubUserRepo{}, &stubSessionRepo{}, sessionUC, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUID}, &stubClock{now: time.Now()}, &stubTxManager{})

		_, err := uc.GetBySessionToken(context.Background(), "bad")
		if err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
			t.Fatalf("expected ErrUnauthorized, got %v", err)
		}
	})
}

func TestLogout(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		sessionUC := &stubSessionUC{}
		uc := New(&stubUserRepo{}, &stubSessionRepo{}, sessionUC, &stubPasswordHasher{verifyOK: true}, &stubUUIDGen{uuid: testUUID}, &stubClock{now: time.Now()}, &stubTxManager{})

		if err := uc.Logout(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

}

func mustUser(id uint, uuidStr string, name string, email string, hash domainUser.UserPasswordHash, isAdmin bool) *domainUser.User {
	user, err := domainUser.Rebuild(
		id,
		uuidStr,
		name,
		email,
		hash.Value(),
		isAdmin,
	)
	if err != nil {
		panic(err)
	}
	return user
}

func mustSession(id uint, uuidStr string, userID uint, createdAt time.Time) *domainSession.Session {
	sess, err := domainSession.Rebuild(id, uuidStr, userID, createdAt)
	if err != nil {
		panic(err)
	}
	return sess
}
