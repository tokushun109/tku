package models

import (
	"api/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var DbConnection *sql.DB
var Db *sql.DB

var err error

func init() {
	DbConnection, err = sql.Open(config.Config.SQLDriver, "root:@/")
	if err != nil {
		log.Fatalln(err)
	}
	defer DbConnection.Close()

	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Config.DBName)
	_, err = DbConnection.Exec(cmdCreateDB)
	if err != nil {
		panic(err)
	}

	// DBの接続
	Db, err = sql.Open(config.Config.SQLDriver, fmt.Sprintf("root:@/%s?parseTime=true", config.Config.DBName))
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
