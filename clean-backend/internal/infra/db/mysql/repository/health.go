package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type HealthRepository struct {
	db *sqlx.DB
}

func NewHealthRepository(db *sqlx.DB) *HealthRepository {
	return &HealthRepository{db: db}
}

func (r *HealthRepository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}
