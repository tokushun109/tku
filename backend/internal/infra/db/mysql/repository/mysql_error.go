package repository

import (
	"errors"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

const mySQLErrorDuplicateEntry uint16 = 1062

func isDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}

	var mysqlErr *mysqlDriver.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == mySQLErrorDuplicateEntry
}
