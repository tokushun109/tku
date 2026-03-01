package mysql

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/backend/internal/infra/config"
)

var ErrNilConfig = errors.New("nil config")

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	if cfg == nil {
		return nil, ErrNilConfig
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC&time_zone='%%2B00%%3A00'",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}
