package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/backend/internal/domain/contact"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
	"github.com/tokushun109/tku/backend/internal/infra/db/mysql/mysqlutil"
)

type ContactRepository struct {
	db *sqlx.DB
}

type contactRow struct {
	ID          uint           `db:"id"`
	UUID        string         `db:"uuid"`
	Name        string         `db:"name"`
	Company     sql.NullString `db:"company"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Email       string         `db:"email"`
	Content     string         `db:"content"`
	CreatedAt   sql.NullTime   `db:"created_at"`
}

func NewContactRepository(db *sqlx.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) FindAll(ctx context.Context) ([]*domain.Contact, error) {
	var rows []contactRow
	if err := getExecutor(ctx, r.db).SelectContext(
		ctx,
		&rows,
		`SELECT id, uuid, name, company, phone_number, email, content, created_at FROM contact WHERE deleted_at IS NULL ORDER BY created_at DESC`,
	); err != nil {
		return nil, err
	}

	res := make([]*domain.Contact, 0, len(rows))
	for _, row := range rows {
		contact, err := toDomainContact(row)
		if err != nil {
			return nil, err
		}
		res = append(res, contact)
	}

	return res, nil
}

func (r *ContactRepository) Create(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {
	company := domainVO.ToValuePtr(contact.Company())
	phoneNumber := domainVO.ToValuePtr(contact.PhoneNumber())
	executor := getExecutor(ctx, r.db)

	res, err := executor.ExecContext(
		ctx,
		`INSERT INTO contact (uuid, name, company, phone_number, email, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`,
		contact.UUID().Value(),
		contact.Name().Value(),
		company,
		phoneNumber,
		contact.Email().Value(),
		contact.Content().Value(),
	)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var row struct {
		CreatedAt sql.NullTime `db:"created_at"`
	}
	if err := executor.GetContext(
		ctx,
		&row,
		`SELECT created_at FROM contact WHERE id = ?`,
		lastID,
	); err != nil {
		return nil, err
	}
	if !row.CreatedAt.Valid {
		return nil, fmt.Errorf("invalid contact row: created_at is null")
	}

	created, err := domain.Rebuild(
		uint(lastID),
		contact.UUID().Value(),
		contact.Name().Value(),
		domainVO.ToValueOrEmpty(contact.Company()),
		domainVO.ToValueOrEmpty(contact.PhoneNumber()),
		contact.Email().Value(),
		contact.Content().Value(),
		row.CreatedAt.Time,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid contact row: %w", err)
	}
	return created, nil
}

func toDomainContact(r contactRow) (*domain.Contact, error) {
	createdAt := time.Time{}
	if r.CreatedAt.Valid {
		createdAt = r.CreatedAt.Time
	}

	company := mysqlutil.NullStringOrEmpty(r.Company)
	phoneNumber := mysqlutil.NullStringOrEmpty(r.PhoneNumber)

	contact, err := domain.Rebuild(
		r.ID,
		r.UUID,
		r.Name,
		company,
		phoneNumber,
		r.Email,
		r.Content,
		createdAt,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid contact row: %w", err)
	}
	return contact, nil
}
