package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/site_detail"
)

type SiteDetailRepository struct {
	db *sqlx.DB
}

func NewSiteDetailRepository(db *sqlx.DB) *SiteDetailRepository {
	return &SiteDetailRepository{db: db}
}

func (r *SiteDetailRepository) ReplaceByProductID(ctx context.Context, productID primitive.ID, details []*domain.SiteDetail) error {
	if _, err := getExecutor(ctx, r.db).ExecContext(ctx, `DELETE FROM site_detail WHERE product_id = ?`, productID.Uint()); err != nil {
		return err
	}
	if len(details) == 0 {
		return nil
	}

	for _, detail := range details {
		if _, err := getExecutor(ctx, r.db).ExecContext(
			ctx,
			`INSERT INTO site_detail (uuid, detail_url, product_id, sales_site_id, created_at, updated_at)
			 VALUES (?, ?, ?, ?, NOW(), NOW())`,
			detail.UUID().String(),
			detail.DetailURL().String(),
			detail.ProductID(),
			detail.SalesSiteID(),
		); err != nil {
			return err
		}
	}
	return nil
}

func (r *SiteDetailRepository) DeleteByProductID(ctx context.Context, productID primitive.ID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(ctx, `DELETE FROM site_detail WHERE product_id = ?`, productID.Uint())
	return err
}
