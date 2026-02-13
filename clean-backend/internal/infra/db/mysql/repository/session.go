package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	"github.com/tokushun109/tku/clean-backend/internal/domain/user"
)

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) FindByUUID(ctx context.Context, uuid domain.SessionUUID) (*domain.Session, error) {
	type row struct {
		UUID      string    `db:"uuid"`
		UserID    uint      `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
	}
	var rrow row
	err := r.db.GetContext(ctx, &rrow, `SELECT uuid, user_id, created_at FROM session WHERE uuid = ?`, uuid.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	parsed, err := domain.ParseSessionUUID(rrow.UUID)
	if err != nil {
		return nil, err
	}
	return &domain.Session{UUID: parsed, UserID: user.NewUserID(rrow.UserID), CreatedAt: rrow.CreatedAt}, nil
}
