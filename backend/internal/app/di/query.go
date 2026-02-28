package di

import (
	"github.com/jmoiron/sqlx"
	mysqlQuery "github.com/tokushun109/tku/backend/internal/infra/db/mysql/query"
	usecaseProductQuery "github.com/tokushun109/tku/backend/internal/usecase/product/query"
)

type queries struct {
	product usecaseProductQuery.Reader
}

func newQueries(db *sqlx.DB) (*queries, error) {
	// 入力側の依存関係のチェック
	if err := requireNonNil("db", db); err != nil {
		return nil, err
	}

	qrs := &queries{
		product: mysqlQuery.NewProductQueryReader(db),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("queries", qrs); err != nil {
		return nil, err
	}

	return qrs, nil
}
