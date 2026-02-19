package mysql

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/txctx"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var (
	ErrNilDBInTxManager = errors.New("nil db")
	ErrNilTxFunc        = errors.New("nil transaction function")
)

type TxManager struct {
	db *sqlx.DB
}

var _ usecase.TxManager = (*TxManager)(nil)

func NewTxManager(db *sqlx.DB) (*TxManager, error) {
	if db == nil {
		return nil, ErrNilDBInTxManager
	}
	return &TxManager{db: db}, nil
}

func (m *TxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	if fn == nil {
		return ErrNilTxFunc
	}
	if ctx == nil {
		ctx = context.Background()
	}

	// Reuse current tx when the caller is already in a transaction.
	if _, ok := txctx.TxFromContext(ctx); ok {
		return fn(ctx)
	}

	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
		if !committed {
			_ = tx.Rollback()
		}
	}()

	txCtx := txctx.WithTx(ctx, tx)
	if err := fn(txCtx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}
