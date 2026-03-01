package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/target"
)

type TargetRepository struct {
	db *sqlx.DB
}

func NewTargetRepository(db *sqlx.DB) *TargetRepository {
	return &TargetRepository{db: db}
}

func (r *TargetRepository) Create(ctx context.Context, t *domain.Target) (*domain.Target, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO target (uuid, name, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`,
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
		return nil, fmt.Errorf("invalid target row: %w", err)
	}
	return created, nil
}

func (r *TargetRepository) FindAll(ctx context.Context) ([]*domain.Target, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, `SELECT id, uuid, name FROM target WHERE deleted_at IS NULL`); err != nil {
		return nil, err
	}
	res := make([]*domain.Target, 0, len(rows))
	for _, r := range rows {
		t, err := toDomainTarget(r.ID, r.UUID, r.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *TargetRepository) FindUsed(ctx context.Context) ([]*domain.Target, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rows []row
	query := `
		SELECT DISTINCT t.id, t.uuid, t.name
		FROM target t
		INNER JOIN product p ON (p.target_uuid = t.uuid OR (p.target_uuid IS NULL AND p.target_id = t.id))
		WHERE t.deleted_at IS NULL AND p.deleted_at IS NULL
	`
	if err := getExecutor(ctx, r.db).SelectContext(ctx, &rows, query); err != nil {
		return nil, err
	}
	res := make([]*domain.Target, 0, len(rows))
	for _, r := range rows {
		t, err := toDomainTarget(r.ID, r.UUID, r.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *TargetRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Target, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, name FROM target WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return toDomainTarget(rrow.ID, rrow.UUID, rrow.Name)
}

func (r *TargetRepository) FindByName(ctx context.Context, name domain.TargetName) (*domain.Target, error) {
	type row struct {
		ID   uint   `db:"id"`
		UUID string `db:"uuid"`
		Name string `db:"name"`
	}
	var rrow row
	if err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name FROM target WHERE name = ? AND deleted_at IS NULL LIMIT 1`,
		name.Value(),
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainTarget(rrow.ID, rrow.UUID, rrow.Name)
}

func (r *TargetRepository) ExistsByName(ctx context.Context, name domain.TargetName) (bool, error) {
	var count int64
	if err := getExecutor(ctx, r.db).GetContext(ctx, &count, `SELECT COUNT(1) FROM target WHERE name = ? AND deleted_at IS NULL`, name.Value()); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TargetRepository) Update(ctx context.Context, t *domain.Target) (bool, error) {
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`UPDATE target SET name = ?, updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`,
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

func (r *TargetRepository) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.ExecContext(
		ctx,
		`UPDATE product
		 SET target_uuid = NULL, target_id = NULL
		 WHERE target_uuid = ? OR (target_uuid IS NULL AND target_id = (SELECT id FROM target WHERE uuid = ? LIMIT 1))`,
		uuid.Value(),
		uuid.Value(),
	); err != nil {
		return false, err
	}

	res, err := tx.ExecContext(ctx, `UPDATE target SET deleted_at = NOW(), updated_at = NOW() WHERE uuid = ? AND deleted_at IS NULL`, uuid.Value())
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

func toDomainTarget(id uint, uuidStr, nameStr string) (*domain.Target, error) {
	target, err := domain.Rebuild(id, uuidStr, nameStr)
	if err != nil {
		return nil, fmt.Errorf("invalid target row: %w", err)
	}
	return target, nil
}
