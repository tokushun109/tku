package repository

import "database/sql"

type SQLDB interface {
	DB() (*sql.DB, error)
}
