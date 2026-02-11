package database

import (
	"errors"

	"gorm.io/gorm"
)

const (
	InstanceMySQL = iota
)

func NewDatabaseSQLFactory(instance int) (*gorm.DB, error) {
	switch instance {
	case InstanceMySQL:
		return GetMySQL()
	default:
		return nil, errors.New("invalid sql database instance")
	}
}
