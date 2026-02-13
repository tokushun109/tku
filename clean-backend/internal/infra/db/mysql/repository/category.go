package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, c *domain.Category) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO category (uuid, name, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`,
		c.ID, c.Name,
	)
	return err
}

func (r *CategoryRepository) FindAll(ctx context.Context) ([]*domain.Category, error) {
	type row struct {
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	if err := r.db.SelectContext(ctx, &rows, `SELECT uuid, name FROM category`); err != nil {
		return nil, err
	}
	res := make([]*domain.Category, 0, len(rows))
	for _, r := range rows {
		res = append(res, &domain.Category{ID: r.UUID, Name: r.Name})
	}
	return res, nil
}

func (r *CategoryRepository) FindUsed(ctx context.Context) ([]*domain.Category, error) {
	type row struct {
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	query := `
		SELECT c.uuid, c.name
		FROM category c
		INNER JOIN product p ON p.category_id = c.id
		WHERE p.deleted_at IS NULL
		GROUP BY c.id, c.uuid, c.name
	`
	if err := r.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, err
	}
	res := make([]*domain.Category, 0, len(rows))
	for _, r := range rows {
		res = append(res, &domain.Category{ID: r.UUID, Name: r.Name})
	}
	return res, nil
}

func (r *CategoryRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	if err := r.db.GetContext(ctx, &count, `SELECT COUNT(1) FROM category WHERE name = ?`, name); err != nil {
		return false, err
	}
	return count > 0, nil
}
