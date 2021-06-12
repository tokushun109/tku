package models

import (
	"api/config"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

// ID, CreatedAt, UpdatedAt, DeletedAtのfieldを持つgorm.Modelを継承
type DefaultModel struct {
	ID        *uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func GenerateUuid() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err)
	}
	return uuidObj.String(), err
}

func gormConnect() *gorm.DB {
	// DB接続設定読み込み
	Sql := config.Config.Sql
	DBUser := config.Config.DBUser
	DBPass := config.Config.DBPass
	Protocol := config.Config.Protocol
	DBName := config.Config.DBName
	SQL_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + "?charset=utf8&parseTime=True&loc=Local"

	// SQLに接続
	DbConnection, err := gorm.Open(Sql, SQL_CONNECT)
	if err != nil {
		log.Fatalln(err)
	}
	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", DBName)
	DbConnection.Exec(cmdCreateDB)

	// DBに接続
	DB_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + DBName + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(Sql, DB_CONNECT)
	if err != nil {
		log.Fatalln(err)
	}
	// テーブル名を単数系で認識
	Db.SingularTable(true)
	return Db
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func init() {
	Db = gormConnect()

	if GetCreator().ID == nil {
		// 製作者の初期データが未作成の場合のみ作成する
		initialInsertCreator()
	}
}
