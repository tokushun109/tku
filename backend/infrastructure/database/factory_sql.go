package database

import (
	"errors"

	"github.com/tokushun109/tku/backend/adapter/repository"
)

const (
	InstanceMySQL = iota
)

func NewDatabaseSQLFactory(instance int) (repository.SQLDB, error) {
	switch instance {
	case InstanceMySQL:
		return GetMySQL()
	default:
		return nil, errors.New("invalid sql database instance")
	}
}
