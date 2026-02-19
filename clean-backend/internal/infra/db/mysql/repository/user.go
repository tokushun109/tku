package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/user"
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

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var rrow userRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name, email, password, is_admin FROM user WHERE email = ? AND deleted_at IS NULL LIMIT 1`,
		email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainUser(rrow)
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var rrow userRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name, email, password, is_admin FROM user WHERE id = ? AND deleted_at IS NULL LIMIT 1`,
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return toDomainUser(rrow)
}

func toDomainUser(rrow userRow) (*domain.User, error) {
	uuid, err := primitive.NewUUID(rrow.UUID)
	if err != nil {
		return nil, fmt.Errorf("invalid user uuid: %w", err)
	}
	passwordHash, err := domain.NewUserPasswordHash(rrow.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid user password hash: %w", err)
	}

	return &domain.User{
		ID:           rrow.ID,
		UUID:         uuid,
		Name:         rrow.Name,
		Email:        rrow.Email,
		PasswordHash: passwordHash,
		IsAdmin:      rrow.IsAdmin,
	}, nil
}
