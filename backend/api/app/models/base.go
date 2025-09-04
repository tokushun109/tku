package models

import (
	"api/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// ID, CreatedAt, UpdatedAt, DeletedAtのfieldを持つDefaultModelを継承
type DefaultModel struct {
	ID        *uint          `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func gormConnect() *gorm.DB {
	// DB接続設定読み込み
	dBUser := config.Config.DBUser
	dBPass := config.Config.DBPass
	protocol := config.Config.Protocol
	dBName := config.Config.DBName
	sqlConnect := dBUser + ":" + dBPass + "@" + protocol + "/" + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"

	// SQLに接続
	dbConnection, err := gorm.Open(mysql.Open(sqlConnect), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	// DBの作成
	cmdCreateDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dBName)
	dbConnection.Exec(cmdCreateDB)

	// DBに接続
	dbConnect := dBUser + ":" + dBPass + "@" + protocol + "/" + dBName + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dbConnect), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	return db.Debug()
}

func GetDBConnection() *gorm.DB {
	return db
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
	db = gormConnect()
	creator := GetCreator()
	if creator.ID == nil {
		// 製作者の初期データが未作成の場合のみ作成する
		err := initialInsertCreator()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
