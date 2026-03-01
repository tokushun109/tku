package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/backend/internal/domain/category"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO category (uuid, name, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`,
		c.UUID().Value(), c.Name().Value(),
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	created, err := domain.Rebuild(uint(lastID), c.UUID().Value(), c.Name().Value())
	if err != nil {
		return nil, fmt.Errorf("invalid category row: %w", err)
	}
	return created, nil
}

func (r *CategoryRepository) FindAll(ctx context.Context) ([]*domain.Category, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT id, uuid, name FROM category WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.Category, 0, len(rows))
	for _, r := range rows {
		c, err := toDomainCategory(r.ID, r.UUID, r.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, nil
}

func (r *CategoryRepository) FindUsed(ctx context.Context) ([]*domain.Category, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	query := `
		SELECT DISTINCT c.id, c.uuid, c.name
		FROM category c
		INNER JOIN product p ON (p.category_uuid = c.uuid OR (p.category_uuid IS NULL AND p.category_id = c.id))
		WHERE c.deleted_at IS NULL AND p.deleted_at IS NULL
	`
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, query); err != nil {
		return nil, err
	}
	res := make([]*domain.Category, 0, len(rows))
	for _, r := range rows {
		c, err := toDomainCategory(r.ID, r.UUID, r.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, nil
}

func (r *CategoryRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Category, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name FROM category WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainCategory(rrow.ID, rrow.UUID, rrow.Name)
}

func (r *CategoryRepository) FindByName(ctx context.Context, name domain.CategoryName) (*domain.Category, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name FROM category WHERE name = ? AND deleted_at IS NULL LIMIT 1`,
		name.Value(),
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainCategory(rrow.ID, rrow.UUID, rrow.Name)
}

func (r *CategoryRepository) ExistsByName(ctx context.Context, name domain.CategoryName) (bool, error) {
	var count int64
	if err := getExecutor(ctx, r.db).GetContext(ctx, &count, `SELECT COUNT(1) FROM category WHERE name = ? AND deleted_at IS NULL`, name.Value()); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CategoryRepository) Update(ctx context.Context, c *domain.Category) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE category SET name = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
		c.Name().Value(),
		c.UUID().Value(),
	)
	if err != nil {
		return false, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.ExecContext(
		ctx,
		`UPDATE product
		 SET category_uuid = NULL, category_id = NULL
		 WHERE category_uuid = ? OR (category_uuid IS NULL AND category_id = (SELECT id FROM category WHERE uuid = ? LIMIT 1))`,
		uuid.Value(),
		uuid.Value(),
	); err != nil {
		return false, err
	}

	res, err := tx.ExecContext(ctx, `UPDATE category SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value())
	if err != nil {
		return false, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if err := tx.Commit(); err != nil {
		return false, err
	}
	return affected > 0, nil
}

func toDomainCategory(id uint, uuidStr, nameStr string) (*domain.Category, error) {
	category, err := domain.Rebuild(id, uuidStr, nameStr)
	if err != nil {
		return nil, fmt.Errorf("invalid category row: %w", err)
	}
	return category, nil
}
