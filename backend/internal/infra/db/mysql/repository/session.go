package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/session"
)

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(ctx context.Context, s *domain.Session) (*domain.Session, error) {
	executor := getExecutor(ctx, r.db)

	_, err := executor.ExecContext(
		ctx,
		`INSERT INTO session (uuid, user_uuid, created_at)
		 VALUES (?, ?, UTC_TIMESTAMP())
		 ON DUPLICATE KEY UPDATE
		     uuid = VALUES(uuid),
		     created_at = VALUES(created_at)`,
		s.UUID().Value(),
		s.UserUUID().Value(),
	)
	if err != nil {
		return nil, err
	}

	var row struct {
		ID        uint      `db:"id"`
		UUID      string    `db:"uuid"`
		UserUUID  string    `db:"user_uuid"`
		CreatedAt time.Time `db:"created_at"`
	}
	if err := executor.GetContext(
		ctx,
		&row,
		`SELECT id, uuid, user_uuid, created_at FROM session WHERE user_uuid = ?`,
		s.UserUUID().Value(),
	); err != nil {
		return nil, err
	}

	created, err := domain.Rebuild(row.ID, row.UUID, row.UserUUID, row.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("invalid session row: %w", err)
	}
	return created, nil
}

func (r *SessionRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Session, error) {
	type row struct {
		ID        uint      `db:"id"`
		UUID      string    `db:"uuid"`
		UserUUID  string    `db:"user_uuid"`
		CreatedAt time.Time `db:"created_at"`
	}
	var rrow row
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT s.id, s.uuid, s.user_uuid, s.created_at
		 FROM session s
		 WHERE s.uuid = ?`,
		uuid.Value(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	sess, err := domain.Rebuild(rrow.ID, rrow.UUID, rrow.UserUUID, rrow.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("invalid session row: %w", err)
	}
	return sess, nil
}

func (r *SessionRepository) DeleteByUUID(ctx context.Context, uuid primitive.UUID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(ctx, `DELETE FROM session WHERE uuid = ?`, uuid.Value())
	return err
}

func (r *SessionRepository) DeleteByUserUUID(ctx context.Context, userUUID primitive.UUID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`DELETE FROM session WHERE user_uuid = ?`,
		userUUID.Value(),
	)
	return err
}
