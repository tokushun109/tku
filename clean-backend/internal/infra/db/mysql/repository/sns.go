package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sns"
)

type SnsRepository struct {
	db *sqlx.DB
}

func NewSnsRepository(db *sqlx.DB) *SnsRepository {
	return &SnsRepository{db: db}
}

func (r *SnsRepository) Create(ctx context.Context, s *domain.Sns) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO sns (uuid, name, url, icon, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`,
		s.UUID.String(), s.Name.String(), s.URL.String(), s.Icon,
	)
	return err
}

func (r *SnsRepository) FindAll(ctx context.Context) ([]*domain.Sns, error) {
	type row struct {
		UUID string         `db:"uuid"`
		Name string         `db:"name"`
		URL  string         `db:"url"`
		Icon sql.NullString `db:"icon"`
	}
	var rows []row
	if err := r.db.SelectContext(ctx, &rows, `SELECT uuid, name, url, icon FROM sns WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.Sns, 0, len(rows))
	for _, r := range rows {
		s, err := toDomainSns(r.UUID, r.Name, r.URL, r.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}

func (r *SnsRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Sns, error) {
	type row struct {
		UUID string         `db:"uuid"`
		Name string         `db:"name"`
		URL  string         `db:"url"`
		Icon sql.NullString `db:"icon"`
	}
	var rrow row
	if err := r.db.GetContext(ctx, &rrow, `SELECT uuid, name, url, icon FROM sns WHERE uuid = ? AND deleted_at IS NULL`, uuid.String()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSns(rrow.UUID, rrow.Name, rrow.URL, rrow.Icon)
}

func (r *SnsRepository) Update(ctx context.Context, s *domain.Sns) (bool, error) {
	res, err := r.db.ExecContext(
		ctx,
		`UPDATE sns SET name = ?, url = ?, icon = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
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

func (r *SnsRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	res, err := r.db.ExecContext(
		ctx,
		`UPDATE sns SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
		uuid.String(),
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

func toDomainSns(uuidStr, nameStr, urlStr string, icon sql.NullString) (*domain.Sns, error) {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sns uuid: %w", err)
	}
	name, err := domain.NewSnsName(nameStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sns name: %w", err)
	}
	snsURL, err := primitive.NewURL(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sns url: %w", err)
	}

	iconValue := ""
	if icon.Valid {
		iconValue = icon.String
	}

	return &domain.Sns{UUID: uuid, Name: name, URL: snsURL, Icon: iconValue}, nil
}
