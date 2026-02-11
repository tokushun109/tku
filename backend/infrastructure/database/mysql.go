package database

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
	dbErr  error
)

func GetMySQL() (*gorm.DB, error) {
	dbOnce.Do(func() {
		db, dbErr = newMySQLFromEnv()
	})
	return db, dbErr
}

func newMySQLFromEnv() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("MYSQL_HOST")
	if user == "" || pass == "" || name == "" || host == "" {
		return nil, errors.New("DB_USER, DB_PASS, DB_NAME, and MYSQL_HOST are required")
	}

	protocol := fmt.Sprintf("tcp(%s)", host)
	dsn := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo&timeout=5s&readTimeout=5s&writeTimeout=5s",
		user,
		pass,
		protocol,
		name,
	)

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	return gdb, nil
}
