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

func (r *UserRepository) FindByEmail(ctx context.Context, email primitive.Email) (*domain.User, error) {
	var rrow userRow
	err := getExecutor(ctx, r.db).GetContext(
		ctx,
		&rrow,
		`SELECT id, uuid, name, email, password, is_admin FROM user WHERE email = ? AND deleted_at IS NULL LIMIT 1`,
		email.String(),
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

func (r *UserRepository) FindContactNotificationUsers(ctx context.Context) ([]*domain.ContactNotificationUser, error) {
	type row struct {
		Name  string `db:"name"`
		Email string `db:"email"`
	}

	var rows []row
	err := getExecutor(ctx, r.db).SelectContext(
		ctx,
		&rows,
		`SELECT name, email FROM user WHERE is_admin = TRUE AND deleted_at IS NULL`,
	)
	if err != nil {
		return nil, err
	}

	res := make([]*domain.ContactNotificationUser, 0, len(rows))
	for _, row := range rows {
		name, err := domain.NewUserName(row.Name)
		if err != nil {
			return nil, fmt.Errorf("invalid admin name: %w", err)
		}
		email, err := primitive.NewEmail(row.Email)
		if err != nil {
			return nil, fmt.Errorf("invalid admin email: %w", err)
		}
		res = append(res, &domain.ContactNotificationUser{
			Name:  name,
			Email: email,
		})
	}

	return res, nil
}

func toDomainUser(rrow userRow) (*domain.User, error) {
	uuid, err := primitive.NewUUID(rrow.UUID)
	if err != nil {
		return nil, fmt.Errorf("invalid user uuid: %w", err)
	}
	name, err := domain.NewUserName(rrow.Name)
	if err != nil {
		return nil, fmt.Errorf("invalid user name: %w", err)
	}
	passwordHash, err := domain.NewUserPasswordHash(rrow.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid user password hash: %w", err)
	}
	email, err := primitive.NewEmail(rrow.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid user email: %w", err)
	}

	return &domain.User{
		ID:           rrow.ID,
		UUID:         uuid,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		IsAdmin:      rrow.IsAdmin,
	}, nil
}
