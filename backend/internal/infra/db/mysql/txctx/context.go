package txctx

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type contextKey struct{}

func WithTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contextKey{}, tx)
}

func TxFromContext(ctx context.Context) (*sqlx.Tx, bool) {
	if ctx == nil {
		return nil, false
	}
	tx, ok := ctx.Value(contextKey{}).(*sqlx.Tx)
	return tx, ok
}
