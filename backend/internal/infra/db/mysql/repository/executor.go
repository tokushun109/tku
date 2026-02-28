package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/infra/db/mysql/txctx"
)

type queryExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func getExecutor(ctx context.Context, db *sqlx.DB) queryExecutor {
	if tx, ok := txctx.TxFromContext(ctx); ok {
		return tx
	}
	return db
}
