package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/product"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
	"github.com/tokushun109/tku/backend/internal/infra/db/mysql/mysqlutil"
)

type ProductRepository struct {
	db *sqlx.DB
}

type productRow struct {
	ID           uint           `db:"id"`
	UUID         string         `db:"uuid"`
	Name         string         `db:"name"`
	Description  sql.NullString `db:"description"`
	Price        int            `db:"price"`
	IsActive     bool           `db:"is_active"`
	IsRecommend  bool           `db:"is_recommend"`
	CategoryUUID sql.NullString `db:"category_uuid"`
	TargetUUID   sql.NullString `db:"target_uuid"`
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *domain.Product) (primitive.ID, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO product (uuid, name, description, price, is_active, is_recommend, category_uuid, target_uuid, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`,
		p.UUID().Value(),
		p.Name().Value(),
		domainVO.ToValuePtr(p.Description()),
		p.Price().Value(),
		p.IsActive(),
		p.IsRecommend(),
		domainVO.ToValuePtr(p.CategoryUUID()),
		domainVO.ToValuePtr(p.TargetUUID()),
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
		`SELECT
			p.id,
			p.uuid,
			p.name,
			p.description,
			p.price,
			p.is_active,
			p.is_recommend,
			p.category_uuid,
			p.target_uuid
		 FROM product p
		 WHERE p.uuid = ? AND p.deleted_at IS NULL`,
		uuid.Value(),
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
		 SET name = ?, description = ?, price = ?, is_active = ?, is_recommend = ?, category_uuid = ?, target_uuid = ?, category_id = NULL, target_id = NULL, updated_at = NOW()
		 WHERE uuid = ? AND deleted_at IS NULL`,
		p.Name().Value(),
		domainVO.ToValuePtr(p.Description()),
		p.Price().Value(),
		p.IsActive(),
		p.IsRecommend(),
		domainVO.ToValuePtr(p.CategoryUUID()),
		domainVO.ToValuePtr(p.TargetUUID()),
		p.UUID().Value(),
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

func (r *ProductRepository) ReplaceTags(ctx context.Context, productUUID primitive.UUID, tagUUIDs []primitive.UUID) error {
	if _, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`DELETE FROM product_to_tag WHERE product_uuid = ?`,
		productUUID.Value(),
	); err != nil {
		return err
	}
	if len(tagUUIDs) == 0 {
		return nil
	}

	for _, tagUUID := range tagUUIDs {
		if _, err := getExecutor(ctx, r.db).ExecContext(
			ctx,
			`INSERT INTO product_to_tag (product_uuid, tag_uuid, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`,
			productUUID.Value(),
			tagUUID.Value(),
		); err != nil {
			return err
		}
	}
	return nil
}

func toDomainProduct(row productRow) (*domain.Product, error) {
	description := mysqlutil.NullStringOrEmpty(row.Description)

	var categoryUUID *string
	if row.CategoryUUID.Valid {
		v := row.CategoryUUID.String
		categoryUUID = &v
	}

	var targetUUID *string
	if row.TargetUUID.Valid {
		v := row.TargetUUID.String
		targetUUID = &v
	}

	product, err := domain.Rebuild(
		row.ID,
		row.UUID,
		row.Name,
		description,
		row.Price,
		row.IsActive,
		row.IsRecommend,
		categoryUUID,
		targetUUID,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid product row: %w", err)
	}
	return product, nil
}
