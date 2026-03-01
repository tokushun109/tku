package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/sales_site"
)

type SalesSiteRepository struct {
	db *sqlx.DB
}

func NewSalesSiteRepository(db *sqlx.DB) *SalesSiteRepository {
	return &SalesSiteRepository{db: db}
}

func (r *SalesSiteRepository) Create(ctx context.Context, s *domain.SalesSite) (*domain.SalesSite, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO sales_site (uuid, name, url, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`,
		s.UUID().Value(), s.Name().Value(), s.URL().Value(),
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	created, err := domain.Rebuild(uint(lastID), s.UUID().Value(), s.Name().Value(), s.URL().Value())
	if err != nil {
		return nil, fmt.Errorf("invalid sales site row: %w", err)
	}
	return created, nil
}

func (r *SalesSiteRepository) FindAll(ctx context.Context) ([]*domain.SalesSite, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT id, uuid, name, url FROM sales_site WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.SalesSite, 0, len(rows))
	for _, r := range rows {
		s, err := toDomainSalesSite(r.ID, r.UUID, r.Name, r.URL)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}

func (r *SalesSiteRepository) FindByName(ctx context.Context, name domain.SalesSiteName) (*domain.SalesSite, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name, url FROM sales_site WHERE name = ? AND deleted_at IS NULL`, name.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSalesSite(rrow.ID, rrow.UUID, rrow.Name, rrow.URL)
}

func (r *SalesSiteRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.SalesSite, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name, url FROM sales_site WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSalesSite(rrow.ID, rrow.UUID, rrow.Name, rrow.URL)
}

func (r *SalesSiteRepository) Update(ctx context.Context, s *domain.SalesSite) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE sales_site SET name = ?, url = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
		s.Name().Value(),
		s.URL().Value(),
		s.UUID().Value(),
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

func (r *SalesSiteRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.ExecContext(
		ctx,
		`DELETE FROM site_detail
		 WHERE sales_site_uuid = ?`,
		uuid.Value(),
	); err != nil {
		return false, err
	}

	res, err := tx.ExecContext(ctx, `UPDATE sales_site SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value())
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

func toDomainSalesSite(id uint, uuidStr, nameStr, urlStr string) (*domain.SalesSite, error) {
	salesSite, err := domain.Rebuild(id, uuidStr, nameStr, urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sales site row: %w", err)
	}
	return salesSite, nil
}
