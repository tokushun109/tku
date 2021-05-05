package models

import (
	"api/config"
	"crypto/sha1"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func gormConnect() *gorm.DB {
	// DB接続設定読み込み
	Sql := config.Config.Sql
	DBUser := config.Config.DBUser
	DBPass := config.Config.DBPass
	Protocol := config.Config.Protocol
	DBName := config.Config.DBName
	SQL_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/"

	// SQLに接続
	DbConnection, err := gorm.Open(Sql, SQL_CONNECT)
	if err != nil {
		log.Fatalln(err)
	}
	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", DBName)
	DbConnection.Exec(cmdCreateDB)

	// DBに接続
	DB_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + DBName
	Db, err := gorm.Open(Sql, DB_CONNECT)
	if err != nil {
		log.Fatalln(err)
	}
	return Db
}

func init() {
	Db := gormConnect()
	defer Db.Close()
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
