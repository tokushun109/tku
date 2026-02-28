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
	createdAt := s.CreatedAt()
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	res, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO session (uuid, user_id, created_at) VALUES (?, ?, ?)`,
		s.UUID().Value(),
		s.UserID(),
		createdAt,
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	created, err := domain.Rebuild(uint(lastID), s.UUID().Value(), s.UserID(), createdAt)
	if err != nil {
		return nil, fmt.Errorf("invalid session row: %w", err)
	}
	return created, nil
}

func (r *SessionRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Session, error) {
	type row struct {
		ID        uint      `db:"id"`
		UUID      string    `db:"uuid"`
		UserID    uint      `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
	}
	var rrow row
	err := getExecutor(ctx, r.db).GetContext(ctx, &rrow, `SELECT id, uuid, user_id, created_at FROM session WHERE uuid = ?`, uuid.Value())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	sess, err := domain.Rebuild(rrow.ID, rrow.UUID, rrow.UserID, rrow.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("invalid session row: %w", err)
	}
	return sess, nil
}

func (r *SessionRepository) DeleteByUUID(ctx context.Context, uuid primitive.UUID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(ctx, `DELETE FROM session WHERE uuid = ?`, uuid.Value())
	return err
}

func (r *SessionRepository) DeleteByUserID(ctx context.Context, userID primitive.ID) error {
	_, err := getExecutor(ctx, r.db).ExecContext(ctx, `DELETE FROM session WHERE user_id = ?`, userID.Value())
	return err
}
