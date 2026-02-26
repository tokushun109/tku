package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/creator"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/mysqlutil"
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
	introduction := domainVO.ToValuePtr(c.Introduction())

	result, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE creator SET name = ?, introduction = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`,
		c.Name().Value(),
		introduction,
		c.ID().Value(),
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
	creatorID primitive.ID,
	mimeType domain.CreatorLogoMimeType,
	logoPath domain.CreatorLogoPath,
) (bool, error) {
	result, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE creator SET mime_type = ?, logo = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`,
		mimeType.Value(),
		logoPath.Value(),
		creatorID.Value(),
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
	introductionRaw := mysqlutil.NullStringOrEmpty(row.Introduction)
	mimeTypeRaw := mysqlutil.NullStringOrEmpty(row.MimeType)
	logoPathRaw := mysqlutil.NullStringOrEmpty(row.Logo)

	creator, err := domain.Rebuild(
		row.ID,
		row.Name,
		introductionRaw,
		mimeTypeRaw,
		logoPathRaw,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid creator row: %w", err)
	}
	return creator, nil
}
