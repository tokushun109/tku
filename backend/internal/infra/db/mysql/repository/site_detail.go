package repository

import (
	"context"

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
	if _, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`DELETE sd
		 FROM site_detail sd
		 LEFT JOIN product p ON sd.product_uuid IS NULL AND p.id = sd.product_id
		 WHERE COALESCE(sd.product_uuid, p.uuid) = ?`,
		productUUID.Value(),
	); err != nil {
		return err
	}
	if len(details) == 0 {
		return nil
	}

	for _, detail := range details {
		if _, err := getExecutor(ctx, r.db).ExecContext(
			ctx,
			`INSERT INTO site_detail (uuid, detail_url, product_uuid, sales_site_uuid, created_at, updated_at)
			 VALUES (?, ?, ?, ?, NOW(), NOW())`,
			detail.UUID().Value(),
			detail.DetailURL().Value(),
			detail.ProductUUID().Value(),
			detail.SalesSiteUUID().Value(),
		); err != nil {
			return err
		}
	}
	return nil
}

func (r *SiteDetailRepository) DeleteByProductUUID(ctx context.Context, productUUID primitive.UUID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`DELETE sd
		 FROM site_detail sd
		 LEFT JOIN product p ON sd.product_uuid IS NULL AND p.id = sd.product_id
		 WHERE COALESCE(sd.product_uuid, p.uuid) = ?`,
		productUUID.Value(),
	)
	return err
}
