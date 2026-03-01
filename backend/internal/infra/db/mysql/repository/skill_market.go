package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/skill_market"
)

type SkillMarketRepository struct {
	db *sqlx.DB
}

func NewSkillMarketRepository(db *sqlx.DB) *SkillMarketRepository {
	return &SkillMarketRepository{db: db}
}

func (r *SkillMarketRepository) Create(ctx context.Context, s *domain.SkillMarket) (*domain.SkillMarket, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO skill_market (uuid, name, url, created_at, updated_at) VALUES (?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`,
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
		return nil, fmt.Errorf("invalid skill market row: %w", err)
	}
	return created, nil
}

func (r *SkillMarketRepository) FindAll(ctx context.Context) ([]*domain.SkillMarket, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT id, uuid, name, url FROM skill_market WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.SkillMarket, 0, len(rows))
	for _, r := range rows {
		s, err := toDomainSkillMarket(r.ID, r.UUID, r.Name, r.URL)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}

func (r *SkillMarketRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.SkillMarket, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
		URL  string `db:"url"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name, url FROM skill_market WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSkillMarket(rrow.ID, rrow.UUID, rrow.Name, rrow.URL)
}

func (r *SkillMarketRepository) Update(ctx context.Context, s *domain.SkillMarket) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE skill_market SET name = ?, url = ?, updated_at = UTC_TIMESTAMP() WHERE uuid = ? AND deleted_at IS NULL`,
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

func (r *SkillMarketRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE skill_market SET deleted_at = UTC_TIMESTAMP(), updated_at = UTC_TIMESTAMP() WHERE uuid = ? AND deleted_at IS NULL`,
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

func toDomainSkillMarket(id uint, uuidStr, nameStr, urlStr string) (*domain.SkillMarket, error) {
	skillMarket, err := domain.Rebuild(id, uuidStr, nameStr, urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid skill market row: %w", err)
	}
	return skillMarket, nil
}
