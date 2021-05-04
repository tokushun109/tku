package models

import (
	"api/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbConnection *sql.DB
var Db *sql.DB

var err error

func gormConnect() *gorm.DB {
	// DBMS := "mysql"
	// USER := "root"
	// PASS := "####"
	// PROTOCOL := "tcp(##.###.##.###:3306)"
	// DBNAME := "##"

	CONNECT := "root:@/"
	DbConnection, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		log.Fatalln(err)
	}
	return DbConnection
}

func init() {
	// DbConnection, err = sql.Open(config.Config.SQLDriver, "root:@/")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer DbConnection.Close()
	Db := gormConnect()
	defer Db.Close()

	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Config.DBName)
	Db.Exec(cmdCreateDB)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
