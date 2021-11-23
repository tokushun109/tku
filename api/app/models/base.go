package models

import (
	"api/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

// ID, CreatedAt, UpdatedAt, DeletedAtのfieldを持つDefaultModelを継承
type DefaultModel struct {
	ID        *uint          `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GenerateUuid() (uuidString string, err error) {
	uuidObj, err := uuid.NewRandom()
	uuidString = uuidObj.String()
	return uuidString, err
}

func gormConnect() *gorm.DB {
	// DB接続設定読み込み
	DBUser := config.Config.DBUser
	DBPass := config.Config.DBPass
	Protocol := config.Config.Protocol
	DBName := config.Config.DBName
	SQL_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + "?charset=utf8&parseTime=True&loc=Local"

	// SQLに接続
	DbConnection, err := gorm.Open(mysql.Open(SQL_CONNECT), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", DBName)
	DbConnection.Exec(cmdCreateDB)

	// DBに接続
	DB_CONNECT := DBUser + ":" + DBPass + "@" + Protocol + "/" + DBName + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(DB_CONNECT), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	return Db.Debug()
}

func removeFile(path string) (err error) {
	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return err
}

func init() {
	Db = gormConnect()

	creator := GetCreator()
	if creator.ID == nil {
		// 製作者の初期データが未作成の場合のみ作成する
		err := initialInsertCreator()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
