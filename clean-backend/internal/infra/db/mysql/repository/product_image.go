package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/product"
)

type ProductImageRepository struct {
	db *sqlx.DB
}

type productImageRow struct {
	ID        uint   `db:"id"`
	UUID      string `db:"uuid"`
	Name      string `db:"name"`
	MimeType  string `db:"mime_type"`
	Path      string `db:"path"`
	Order     int    `db:"order"`
	ProductID uint   `db:"product_id"`
}

func NewProductImageRepository(db *sqlx.DB) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

func (r *ProductImageRepository) Create(ctx context.Context, image *domain.ProductImage) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		"INSERT INTO product_image (uuid, name, mime_type, path, `order`, product_id, created_at, updated_at) "+
			"VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())",
		image.UUID().String(),
		image.Name().String(),
		image.MimeType().String(),
		image.Path().String(),
		image.Order().Int(),
		image.ProductID(),
	)
	return err
}

func (r *ProductImageRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.ProductImage, error) {
	var row productImageRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&row,
		"SELECT id, uuid, name, mime_type, path, `order`, product_id "+
			"FROM product_image "+
			"WHERE uuid = ? AND deleted_at IS NULL",
		uuid.String(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainProductImage(row)
}

func (r *ProductImageRepository) FindByProductID(ctx context.Context, productID primitive.ID) ([]*domain.ProductImage, error) {
	rows := []productImageRow{}
	err := getExecutor(ctx, r.db).SelectContext(
		ctx,
		&rows,
		"SELECT id, uuid, name, mime_type, path, `order`, product_id "+
			"FROM product_image "+
			"WHERE product_id = ? AND deleted_at IS NULL "+
			"ORDER BY `order` DESC, id ASC",
		productID.Uint(),
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

func (r *ProductImageRepository) UpdateOrder(ctx context.Context, uuid primitive.UUID, order int) (bool, error) {
	validatedOrder, err := domain.NewProductImageOrder(order)
	if err != nil {
		return false, err
	}

	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		"UPDATE product_image SET `order` = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL",
		validatedOrder.Int(),
		uuid.String(),
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
		uuid.String(),
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

func (r *ProductImageRepository) DeleteByProductID(ctx context.Context, productID primitive.ID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE product_image SET deleted_at = NOW(), updated_at = NOW() WHERE product_id = ? AND deleted_at IS NULL`,
		productID.Uint(),
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
		row.Order,
		row.ProductID,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid product image row: %w", err)
	}
	return image, nil
}
