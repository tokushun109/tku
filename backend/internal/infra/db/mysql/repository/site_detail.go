package repository

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/site_detail"
)

type SiteDetailRepository struct {
	db *sqlx.DB
}

func NewSiteDetailRepository(db *sqlx.DB) *SiteDetailRepository {
	return &SiteDetailRepository{db: db}
}

func (r *SiteDetailRepository) ReplaceByProductUUID(ctx context.Context, productUUID primitive.UUID, details []*domain.SiteDetail) error {
	executor := getExecutor(ctx, r.db)

	if _, err := executor.ExecContext(
		ctx,
		`DELETE FROM site_detail WHERE product_uuid = ?`,
		productUUID.Value(),
	); err != nil {
		return err
	}
	if len(details) == 0 {
		return nil
	}

	const insertQueryPrefix = `INSERT INTO site_detail (uuid, detail_url, product_uuid, sales_site_uuid, created_at, updated_at) VALUES `

	var queryBuilder strings.Builder
	queryBuilder.WriteString(insertQueryPrefix)

	args := make([]interface{}, 0, len(details)*4)
	for _, detail := range details {
		if len(args) > 0 {
			queryBuilder.WriteString(", ")
		}
		queryBuilder.WriteString("(?, ?, ?, ?, UTC_TIMESTAMP(6), UTC_TIMESTAMP(6))")
		args = append(
			args,
			detail.UUID().Value(),
			detail.DetailURL().Value(),
			detail.ProductUUID().Value(),
			detail.SalesSiteUUID().Value(),
		)
	}

	if _, err := executor.ExecContext(ctx, queryBuilder.String(), args...); err != nil {
		return err
	}
	return nil
}

func (r *SiteDetailRepository) DeleteByProductUUID(ctx context.Context, productUUID primitive.UUID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`DELETE FROM site_detail WHERE product_uuid = ?`,
		productUUID.Value(),
	)
	return err
}
