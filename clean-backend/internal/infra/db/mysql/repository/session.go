package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/session"
)

type SessionModel struct {
	ID        uint      `db:"id"`
	UUID      string    `db:"uuid"`
	UserID    uint      `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (SessionModel) TableName() string { return "session" }

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) FindByUUID(ctx context.Context, uuid string) (*domain.Session, error) {
	var model SessionModel
	err := r.db.GetContext(ctx, &model, `SELECT uuid, user_id, created_at FROM session WHERE uuid = ?`, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &domain.Session{UUID: model.UUID, UserID: model.UserID, CreatedAt: model.CreatedAt}, nil
}
