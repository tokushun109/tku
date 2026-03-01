package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/tag"
)

type TagRepository struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Create(ctx context.Context, t *domain.Tag) (*domain.Tag, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO tag (uuid, name, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`,
		t.UUID().Value(), t.Name().Value(),
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	created, err := domain.Rebuild(uint(lastID), t.UUID().Value(), t.Name().Value())
	if err != nil {
		return nil, fmt.Errorf("invalid tag row: %w", err)
	}
	return created, nil
}

func (r *TagRepository) FindAll(ctx context.Context) ([]*domain.Tag, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT id, uuid, name FROM tag WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.Tag, 0, len(rows))
	for _, r := range rows {
		t, err := toDomainTag(r.ID, r.UUID, r.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *TagRepository) FindByName(ctx context.Context, name domain.TagName) (*domain.Tag, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name FROM tag WHERE name = ? AND deleted_at IS NULL`, name.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainTag(rrow.ID, rrow.UUID, rrow.Name)
}

func (r *TagRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Tag, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name FROM tag WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainTag(rrow.ID, rrow.UUID, rrow.Name)
}

func (r *TagRepository) ExistsByName(ctx context.Context, name domain.TagName) (bool, error) {
	var count int64
	if err := getExecutor(ctx, r.db).GetContext(ctx, &count, `SELECT COUNT(1) FROM tag WHERE name = ? AND deleted_at IS NULL`, name.Value()); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TagRepository) Update(ctx context.Context, t *domain.Tag) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE tag SET name = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
		t.Name().Value(),
		t.UUID().Value(),
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

func (r *TagRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.ExecContext(
		ctx,
		`DELETE FROM product_to_tag
		 WHERE tag_uuid = ? OR (tag_uuid IS NULL AND tag_id = (SELECT id FROM tag WHERE uuid = ? LIMIT 1))`,
		uuid.Value(),
		uuid.Value(),
	); err != nil {
		return false, err
	}

	res, err := tx.ExecContext(ctx, `UPDATE tag SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value())
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

func toDomainTag(id uint, uuidStr, nameStr string) (*domain.Tag, error) {
	tag, err := domain.Rebuild(id, uuidStr, nameStr)
	if err != nil {
		return nil, fmt.Errorf("invalid tag row: %w", err)
	}
	return tag, nil
}
