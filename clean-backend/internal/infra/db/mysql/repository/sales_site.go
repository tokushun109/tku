package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sales_site"
)

type SalesSiteRepository struct {
	db *sqlx.DB
}

func NewSalesSiteRepository(db *sqlx.DB) *SalesSiteRepository {
	return &SalesSiteRepository{db: db}
}

func (r *SalesSiteRepository) Create(ctx context.Context, s *domain.SalesSite) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO sales_site (uuid, name, url, icon, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`,
		s.UUID.String(), s.Name.String(), s.URL.String(), s.Icon,
	)
	return err
}

func (r *SalesSiteRepository) FindAll(ctx context.Context) ([]*domain.SalesSite, error) {
	type row struct {
		UUID string         `db:"uuid"`
		Name string         `db:"name"`
		URL  string         `db:"url"`
		Icon sql.NullString `db:"icon"`
	}
	var rows []row
	if err := r.db.SelectContext(ctx, &rows, `SELECT uuid, name, url, icon FROM sales_site WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.SalesSite, 0, len(rows))
	for _, r := range rows {
		s, err := toDomainSalesSite(r.UUID, r.Name, r.URL, r.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}

func (r *SalesSiteRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.SalesSite, error) {
	type row struct {
		UUID string         `db:"uuid"`
		Name string         `db:"name"`
		URL  string         `db:"url"`
		Icon sql.NullString `db:"icon"`
	}
	var rrow row
	if err := r.db.GetContext(ctx, &rrow, `SELECT uuid, name, url, icon FROM sales_site WHERE uuid = ? AND deleted_at IS NULL`, uuid.String()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSalesSite(rrow.UUID, rrow.Name, rrow.URL, rrow.Icon)
}

func (r *SalesSiteRepository) Update(ctx context.Context, s *domain.SalesSite) (bool, error) {
	res, err := r.db.ExecContext(
		ctx,
		`UPDATE sales_site SET name = ?, url = ?, icon = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
		s.Name.String(),
		s.URL.String(),
		s.Icon,
		s.UUID.String(),
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

	var salesSiteID int64
	if err := tx.GetContext(ctx, &salesSiteID, `SELECT id FROM sales_site WHERE uuid = ? AND deleted_at IS NULL`, uuid.String()); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if _, err := tx.ExecContext(ctx, `DELETE FROM site_detail WHERE sales_site_id = ?`, salesSiteID); err != nil {
		return false, err
	}

	res, err := tx.ExecContext(ctx, `UPDATE sales_site SET deleted_at = NOW(), updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`, salesSiteID)
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

func toDomainSalesSite(uuidStr, nameStr, urlStr string, icon sql.NullString) (*domain.SalesSite, error) {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sales site uuid: %w", err)
	}
	name, err := domain.NewSalesSiteName(nameStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sales site name: %w", err)
	}
	salesSiteURL, err := primitive.NewURL(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sales site url: %w", err)
	}

	iconValue := ""
	if icon.Valid {
		iconValue = icon.String
	}

	return &domain.SalesSite{UUID: uuid, Name: name, URL: salesSiteURL, Icon: iconValue}, nil
}
