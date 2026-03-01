package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/product"
)

type ProductImageRepository struct {
	db *sqlx.DB
}

type productImageRow struct {
	ID           uint   `db:"id"`
	UUID         string `db:"uuid"`
	Name         string `db:"name"`
	MimeType     string `db:"mime_type"`
	Path         string `db:"path"`
	DisplayOrder int    `db:"display_order"`
	ProductUUID  string `db:"product_uuid"`
}

func NewProductImageRepository(db *sqlx.DB) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

func (r *ProductImageRepository) Create(ctx context.Context, image *domain.ProductImage) (*domain.ProductImage, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`
			INSERT INTO product_image (uuid, name, mime_type, path, display_order, product_uuid, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())
			`,
		image.UUID().Value(),
		image.Name().Value(),
		image.MimeType().Value(),
		image.Path().Value(),
		image.DisplayOrder().Value(),
		image.ProductUUID().Value(),
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	created, err := domain.RebuildProductImage(
		uint(lastID),
		image.UUID().Value(),
		image.Name().Value(),
		image.MimeType().Value(),
		image.Path().Value(),
		image.DisplayOrder().Value(),
		image.ProductUUID().Value(),
	)
	if err != nil {
		return nil, fmt.Errorf("invalid product image row: %w", err)
	}
	return created, nil
}

func (r *ProductImageRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.ProductImage, error) {
	var row productImageRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&row,
		`
			SELECT
				pi.id,
				pi.uuid,
				pi.name,
				pi.mime_type,
				pi.path,
				pi.display_order,
				pi.product_uuid
			FROM product_image pi
			WHERE pi.uuid = ? AND pi.deleted_at IS NULL
			`,
		uuid.Value(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainProductImage(row)
}

func (r *ProductImageRepository) FindByProductUUID(ctx context.Context, productUUID primitive.UUID) ([]*domain.ProductImage, error) {
	rows := []productImageRow{}
	err := getExecutor(ctx, r.db).SelectContext(
		ctx,
		&rows,
		`
			SELECT
				pi.id,
				pi.uuid,
				pi.name,
				pi.mime_type,
				pi.path,
				pi.display_order,
				pi.product_uuid
			FROM product_image pi
			WHERE pi.product_uuid = ? AND pi.deleted_at IS NULL
			ORDER BY pi.display_order DESC, pi.id ASC
			`,
		productUUID.Value(),
	)
	if err != nil {
		return nil, err
	}

	images := make([]*domain.ProductImage, 0, len(rows))
	for _, row := range rows {
		image, err := toDomainProductImage(row)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (r *ProductImageRepository) UpdateDisplayOrder(ctx context.Context, uuid primitive.UUID, displayOrder int) (bool, error) {
	validatedDisplayOrder, err := domain.NewProductImageDisplayOrder(displayOrder)
	if err != nil {
		return false, err
	}

	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`
		UPDATE product_image
		SET display_order = ?, updated_at = NOW()
		WHERE uuid = ? AND deleted_at IS NULL
		`,
		validatedDisplayOrder.Value(),
		uuid.Value(),
	)
	if err != nil {
		return false, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return affectedRows > 0, nil
}

func (r *ProductImageRepository) DeleteByUUID(ctx context.Context, uuid primitive.UUID) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE product_image SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
		uuid.Value(),
	)
	if err != nil {
		return false, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return affectedRows > 0, nil
}

func (r *ProductImageRepository) DeleteByProductUUID(ctx context.Context, productUUID primitive.UUID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE product_image
		 SET deleted_at = NOW(), updated_at = NOW()
		 WHERE product_uuid = ? AND deleted_at IS NULL`,
		productUUID.Value(),
	)
	return err
}

func toDomainProductImage(row productImageRow) (*domain.ProductImage, error) {
	image, err := domain.RebuildProductImage(
		row.ID,
		row.UUID,
		row.Name,
		row.MimeType,
		row.Path,
		row.DisplayOrder,
		row.ProductUUID,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid product image row: %w", err)
	}
	return image, nil
}
