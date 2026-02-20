package contact

import (
	"context"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/contact"
)

type Notifier interface {
	NotifyContactCreated(ctx context.Context, contact *domain.Contact)
}

type nopNotifier struct{}

func (nopNotifier) NotifyContactCreated(context.Context, *domain.Contact) {}
