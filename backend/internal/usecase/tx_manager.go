package usecase

import "context"

type TxManager interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
