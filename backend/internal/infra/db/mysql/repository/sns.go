package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/sns"
)

type SnsRepository struct {
	db *sqlx.DB
}

func NewSnsRepository(db *sqlx.DB) *SnsRepository {
	return &SnsRepository{db: db}
}

func (r *SnsRepository) Create(ctx context.Context, s *domain.Sns) (*domain.Sns, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO sns (uuid, name, url, created_at, updated_at) VALUES (?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`,
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
		return nil, fmt.Errorf("invalid sns row: %w", err)
	}
	return created, nil
}

func (r *SnsRepository) FindAll(ctx context.Context) ([]*domain.Sns, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT id, uuid, name, url FROM sns WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.Sns, 0, len(rows))
	for _, r := range rows {
		s, err := toDomainSns(r.ID, r.UUID, r.Name, r.URL)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}

func (r *SnsRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Sns, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name, url FROM sns WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSns(rrow.ID, rrow.UUID, rrow.Name, rrow.URL)
}

func (r *SnsRepository) Update(ctx context.Context, s *domain.Sns) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE sns SET name = ?, url = ?, updated_at = UTC_TIMESTAMP() WHERE uuid = ? AND deleted_at IS NULL`,
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

func (r *SnsRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE sns SET deleted_at = UTC_TIMESTAMP(), updated_at = UTC_TIMESTAMP() WHERE uuid = ? AND deleted_at IS NULL`,
		uuid.Value(),
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

func toDomainSns(id uint, uuidStr, nameStr, urlStr string) (*domain.Sns, error) {
	sns, err := domain.Rebuild(id, uuidStr, nameStr, urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sns row: %w", err)
	}
	return sns, nil
}
