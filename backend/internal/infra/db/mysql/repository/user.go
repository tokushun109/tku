package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/user"
)

type UserRepository struct {
	db *sqlx.DB
}

type userRow struct {
	ID       uint   `db:"id"`
	UUID     string `db:"uuid"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsAdmin  bool   `db:"is_admin"`
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email primitive.Email) (*domain.User, error) {
	var rrow userRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name, email, password, is_admin FROM user WHERE email = ? AND deleted_at IS NULL LIMIT 1`,
		email.Value(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainUser(rrow)
}

func (r *UserRepository) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.User, error) {
	var rrow userRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name, email, password, is_admin FROM user WHERE uuid = ? AND deleted_at IS NULL LIMIT 1`,
		uuid.Value(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainUser(rrow)
}

func (r *UserRepository) FindContactNotificationUsers(ctx context.Context) ([]*domain.ContactNotificationUser, error) {
	type row struct {
		ID    uint   `db:"id"`
		Name  string `db:"name"`
		Email string `db:"email"`
	}

	var rows []row
	err := getExecutor(ctx, r.db).SelectContext(
		ctx,
		&rows,
		`SELECT id, name, email FROM user WHERE is_admin = TRUE AND deleted_at IS NULL`,
	)
	if err != nil {
		return nil, err
	}

	res := make([]*domain.ContactNotificationUser, 0, len(rows))
	for _, row := range rows {
		user, err := domain.RebuildContactNotificationUser(row.ID, row.Name, row.Email)
		if err != nil {
			return nil, fmt.Errorf("invalid admin user row: %w", err)
		}
		res = append(res, user)
	}

	return res, nil
}

func toDomainUser(rrow userRow) (*domain.User, error) {
	user, err := domain.Rebuild(
		rrow.ID,
		rrow.UUID,
		rrow.Name,
		rrow.Email,
		rrow.Password,
		rrow.IsAdmin,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid user row: %w", err)
	}
	return user, nil
}
