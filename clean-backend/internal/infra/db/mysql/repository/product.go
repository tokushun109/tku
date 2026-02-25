package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/product"
	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/mysqlutil"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
)

type ProductRepository struct {
	db *sqlx.DB
}

type productRow struct {
	ID          uint           `db:"id"`
	UUID        string         `db:"uuid"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	Price       int            `db:"price"`
	IsActive    bool           `db:"is_active"`
	IsRecommend bool           `db:"is_recommend"`
	CategoryID  sql.NullInt64  `db:"category_id"`
	TargetID    sql.NullInt64  `db:"target_id"`
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *domain.Product) (primitive.ID, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO product (uuid, name, description, price, is_active, is_recommend, category_id, target_id, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`,
		p.UUID().String(),
		p.Name().String(),
		optional.ToStringPtr(p.Description()),
		p.Price().Int(),
		p.IsActive(),
		p.IsRecommend(),
		toUintPtrFromPrimitiveID(p.CategoryID()),
		toUintPtrFromPrimitiveID(p.TargetID()),
	)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if lastID <= 0 {
		return 0, fmt.Errorf("invalid inserted product id: %d", lastID)
	}
	id, err := primitive.NewID(uint(lastID))
	if err != nil {
		return 0, fmt.Errorf("invalid inserted product id: %d", lastID)
	}
	return id, nil
}

func (r *ProductRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Product, error) {
	var row productRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&row,
		`SELECT id, uuid, name, description, price, is_active, is_recommend, category_id, target_id
		 FROM product
		 WHERE uuid = ? AND deleted_at IS NULL`,
		uuid.String(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainProduct(row)
}

func (r *ProductRepository) Update(ctx context.Context, p *domain.Product) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE product
		 SET name = ?, description = ?, price = ?, is_active = ?, is_recommend = ?, category_id = ?, target_id = ?, updated_at = NOW()
		 WHERE uuid = ? AND deleted_at IS NULL`,
		p.Name().String(),
		optional.ToStringPtr(p.Description()),
		p.Price().Int(),
		p.IsActive(),
		p.IsRecommend(),
		toUintPtrFromPrimitiveID(p.CategoryID()),
		toUintPtrFromPrimitiveID(p.TargetID()),
		p.UUID().String(),
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

func (r *ProductRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE product SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
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

func (r *ProductRepository) ReplaceTags(ctx context.Context, productID primitive.ID, tagIDs []primitive.ID) error {
	if _, err := getExecutor(ctx, r.db).ExecContext(ctx, `DELETE FROM product_to_tag WHERE product_id = ?`, productID.Uint()); err != nil {
		return err
	}
	if len(tagIDs) == 0 {
		return nil
	}

	for _, tagID := range tagIDs {
		if _, err := getExecutor(ctx, r.db).ExecContext(
			ctx,
			`INSERT INTO product_to_tag (product_id, tag_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`,
			productID.Uint(),
			tagID.Uint(),
		); err != nil {
			return err
		}
	}
	return nil
}

func toDomainProduct(row productRow) (*domain.Product, error) {
	description := mysqlutil.NullStringOrEmpty(row.Description)

	var categoryID *uint
	if row.CategoryID.Valid && row.CategoryID.Int64 > 0 {
		v := uint(row.CategoryID.Int64)
		categoryID = &v
	}

	var targetID *uint
	if row.TargetID.Valid && row.TargetID.Int64 > 0 {
		v := uint(row.TargetID.Int64)
		targetID = &v
	}

	product, err := domain.Rebuild(
		row.ID,
		row.UUID,
		row.Name,
		description,
		row.Price,
		row.IsActive,
		row.IsRecommend,
		categoryID,
		targetID,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid product row: %w", err)
	}
	return product, nil
}

func toUintPtrFromPrimitiveID(v *primitive.ID) *uint {
	if v == nil {
		return nil
	}
	id := v.Uint()
	return &id
}
