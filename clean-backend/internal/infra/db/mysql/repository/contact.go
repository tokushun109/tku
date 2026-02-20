package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/contact"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
)

type ContactRepository struct {
	db *sqlx.DB
}

type contactRow struct {
	ID          uint           `db:"id"`
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
		`SELECT id, name, company, phone_number, email, content, created_at FROM contact WHERE deleted_at IS NULL ORDER BY created_at DESC`,
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

func (r *ContactRepository) Create(ctx context.Context, contact *domain.Contact) error {
	company := optional.ToStringPtr(contact.Company)
	phoneNumber := optional.ToStringPtr(contact.PhoneNumber)

	_, err := getExecutor(ctx, r.db).ExecContext(
		ctx,
		`INSERT INTO contact (name, company, phone_number, email, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())`,
		contact.Name.String(),
		company,
		phoneNumber,
		contact.Email.String(),
		contact.Content.String(),
	)
	return err
}

func toDomainContact(r contactRow) (*domain.Contact, error) {
	name, err := domain.NewContactName(r.Name)
	if err != nil {
		return nil, fmt.Errorf("invalid contact name: %w", err)
	}

	var company *domain.ContactCompany
	if r.Company.Valid && strings.TrimSpace(r.Company.String) != "" {
		companyValue, err := domain.NewContactCompany(r.Company.String)
		if err != nil {
			return nil, fmt.Errorf("invalid contact company: %w", err)
		}
		company = &companyValue
	}

	var phoneNumber *primitive.PhoneNumber
	if r.PhoneNumber.Valid && strings.TrimSpace(r.PhoneNumber.String) != "" {
		phoneNumberValue, err := primitive.NewPhoneNumber(r.PhoneNumber.String)
		if err != nil {
			return nil, fmt.Errorf("invalid contact phone number: %w", err)
		}
		phoneNumber = &phoneNumberValue
	}

	email, err := primitive.NewEmail(r.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid contact email: %w", err)
	}

	content, err := domain.NewContactContent(r.Content)
	if err != nil {
		return nil, fmt.Errorf("invalid contact content: %w", err)
	}

	createdAt := time.Time{}
	if r.CreatedAt.Valid {
		createdAt = r.CreatedAt.Time
	}

	return &domain.Contact{
		ID:          r.ID,
		Name:        name,
		Company:     company,
		PhoneNumber: phoneNumber,
		Email:       email,
		Content:     content,
		CreatedAt:   createdAt,
	}, nil
}
