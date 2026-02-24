package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/creator"
)

type CreatorRepository struct {
	db *sqlx.DB
}

type creatorRow struct {
	ID           uint           `db:"id"`
	Name         string         `db:"name"`
	Introduction sql.NullString `db:"introduction"`
	MimeType     sql.NullString `db:"mime_type"`
	Logo         sql.NullString `db:"logo"`
}

func NewCreatorRepository(db *sqlx.DB) *CreatorRepository {
	return &CreatorRepository{db: db}
}

func (r *CreatorRepository) Find(ctx context.Context) (*domain.Creator, error) {
	var row creatorRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&row,
		`SELECT id, name, introduction, mime_type, logo FROM creator WHERE deleted_at IS NULL ORDER BY id ASC LIMIT 1`,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainCreator(row)
}

func (r *CreatorRepository) UpdateProfile(ctx context.Context, c *domain.Creator) (bool, error) {
	result, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE creator SET name = ?, introduction = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`,
		c.Name.String(),
		c.Introduction.String(),
		c.ID,
	)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affectedRows > 0, nil
}

func (r *CreatorRepository) UpdateLogo(
	ctx context.Context,
	creatorID uint,
	mimeType domain.CreatorLogoMimeType,
	logoPath domain.CreatorLogoPath,
) (bool, error) {
	result, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE creator SET mime_type = ?, logo = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`,
		mimeType.String(),
		logoPath.String(),
		creatorID,
	)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affectedRows > 0, nil
}

func toDomainCreator(row creatorRow) (*domain.Creator, error) {
	name, err := domain.NewCreatorName(row.Name)
	if err != nil {
		return nil, fmt.Errorf("invalid creator name: %w", err)
	}

	introductionRaw := ""
	if row.Introduction.Valid {
		introductionRaw = row.Introduction.String
	}
	introduction, err := domain.NewCreatorIntroductionForRead(introductionRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid creator introduction: %w", err)
	}

	var mimeType *domain.CreatorLogoMimeType
	if row.MimeType.Valid && strings.TrimSpace(row.MimeType.String) != "" {
		v, err := domain.NewCreatorLogoMimeType(row.MimeType.String)
		if err != nil {
			return nil, fmt.Errorf("invalid creator logo mime type: %w", err)
		}
		mimeType = &v
	}

	var logoPath *domain.CreatorLogoPath
	if row.Logo.Valid && strings.TrimSpace(row.Logo.String) != "" {
		v, err := domain.NewCreatorLogoPath(row.Logo.String)
		if err != nil {
			return nil, fmt.Errorf("invalid creator logo path: %w", err)
		}
		logoPath = &v
	}

	return &domain.Creator{
		ID:           row.ID,
		Name:         name,
		Introduction: introduction,
		LogoMimeType: mimeType,
		LogoPath:     logoPath,
	}, nil
}
