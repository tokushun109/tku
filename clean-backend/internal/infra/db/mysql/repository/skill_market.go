package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/skill_market"
)

type SkillMarketRepository struct {
	db *sqlx.DB
}

func NewSkillMarketRepository(db *sqlx.DB) *SkillMarketRepository {
	return &SkillMarketRepository{db: db}
}

func (r *SkillMarketRepository) Create(ctx context.Context, s *domain.SkillMarket) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO skill_market (uuid, name, url, icon, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`,
		s.UUID.String(), s.Name.String(), s.URL.String(), s.Icon,
	)
	return err
}

func (r *SkillMarketRepository) FindAll(ctx context.Context) ([]*domain.SkillMarket, error) {
	type row struct {
		UUID string         `db:"uuid"`
		Name string         `db:"name"`
		URL  string         `db:"url"`
		Icon sql.NullString `db:"icon"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT uuid, name, url, icon FROM skill_market WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.SkillMarket, 0, len(rows))
	for _, r := range rows {
		s, err := toDomainSkillMarket(r.UUID, r.Name, r.URL, r.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	return res, nil
}

func (r *SkillMarketRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.SkillMarket, error) {
	type row struct {
		UUID string         `db:"uuid"`
		Name string         `db:"name"`
		URL  string         `db:"url"`
		Icon sql.NullString `db:"icon"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT uuid, name, url, icon FROM skill_market WHERE uuid = ? AND deleted_at IS NULL`, uuid.String()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainSkillMarket(rrow.UUID, rrow.Name, rrow.URL, rrow.Icon)
}

func (r *SkillMarketRepository) Update(ctx context.Context, s *domain.SkillMarket) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE skill_market SET name = ?, url = ?, icon = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
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

func (r *SkillMarketRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE skill_market SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
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

func toDomainSkillMarket(uuidStr, nameStr, urlStr string, icon sql.NullString) (*domain.SkillMarket, error) {
	uuid, err := primitive.NewUUID(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("invalid skill market uuid: %w", err)
	}
	name, err := domain.NewSkillMarketName(nameStr)
	if err != nil {
		return nil, fmt.Errorf("invalid skill market name: %w", err)
	}
	skillMarketURL, err := primitive.NewURL(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid skill market url: %w", err)
	}

	iconValue := ""
	if icon.Valid {
		iconValue = icon.String
	}

	return &domain.SkillMarket{UUID: uuid, Name: name, URL: skillMarketURL, Icon: iconValue}, nil
}
