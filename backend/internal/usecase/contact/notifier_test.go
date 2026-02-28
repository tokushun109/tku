package contact

import (
	"context"
	"errors"
	"testing"

	domainContact "github.com/tokushun109/tku/backend/internal/domain/contact"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
	usecase "github.com/tokushun109/tku/backend/internal/usecase"
)

type stubMailer struct {
	sent    []*usecase.MailMessage
	sendErr error
}

func (s *stubMailer) Send(ctx context.Context, message *usecase.MailMessage) error {
	if s.sendErr != nil {
		return s.sendErr
	}
	s.sent = append(s.sent, message)
	return nil
}

type stubNotificationUserRepo struct {
	users []*domainUser.ContactNotificationUser
	err   error
}

func (s *stubNotificationUserRepo) FindByEmail(ctx context.Context, email primitive.Email) (*domainUser.User, error) {
	return nil, nil
}

func (s *stubNotificationUserRepo) FindByID(ctx context.Context, id primitive.ID) (*domainUser.User, error) {
	return nil, nil
}

func (s *stubNotificationUserRepo) FindContactNotificationUsers(ctx context.Context) ([]*domainUser.ContactNotificationUser, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.users, nil
}

func TestContactNotifierNotifyContactCreated(t *testing.T) {
	t.Run("送信先管理者が存在するなら自動返信と管理者通知を送信する", func(t *testing.T) {
		notificationUser, err := domainUser.RebuildContactNotificationUser(1, "管理者", "admin@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		mailer := &stubMailer{}
		repo := &stubNotificationUserRepo{
			users: []*domainUser.ContactNotificationUser{
				notificationUser,
			},
		}
		notifier := NewContactNotifier(mailer, repo, "")

		contact, err := domainContact.New("山田太郎", "株式会社サンプル", "09012345678", "user@example.com", "お問い合わせです")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		notifier.NotifyContactCreated(context.Background(), contact)

		if len(mailer.sent) != 2 {
			t.Fatalf("expected 2 mails sent, got %d", len(mailer.sent))
		}
		if mailer.sent[0].Subject != autoReplySubject {
			t.Fatalf("expected first subject %q, got %q", autoReplySubject, mailer.sent[0].Subject)
		}
		if len(mailer.sent[0].To) != 1 || mailer.sent[0].To[0].Email != "user@example.com" {
			t.Fatalf("expected first mail to user@example.com, got %#v", mailer.sent[0].To)
		}
		if mailer.sent[1].Subject != adminMailSubject {
			t.Fatalf("expected second subject %q, got %q", adminMailSubject, mailer.sent[1].Subject)
		}
		if len(mailer.sent[1].To) != 1 || mailer.sent[1].To[0].Email != "admin@example.com" {
			t.Fatalf("expected second mail to admin@example.com, got %#v", mailer.sent[1].To)
		}
	})

	t.Run("管理者宛先が取得できないなら自動返信のみ送信する", func(t *testing.T) {
		mailer := &stubMailer{}
		repo := &stubNotificationUserRepo{err: errors.New("db error")}
		notifier := NewContactNotifier(mailer, repo, "")

		contact, err := domainContact.New("山田太郎", "", "", "user@example.com", "お問い合わせです")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		notifier.NotifyContactCreated(context.Background(), contact)

		if len(mailer.sent) != 1 {
			t.Fatalf("expected 1 mail sent, got %d", len(mailer.sent))
		}
		if mailer.sent[0].Subject != autoReplySubject {
			t.Fatalf("expected subject %q, got %q", autoReplySubject, mailer.sent[0].Subject)
		}
	})
}
