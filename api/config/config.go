package config

import (
	"api/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port        string
	Sql         string
	LogFile     string
	Env         string
	DBUser      string
	DBName      string
	DBPass      string
	Protocol    string
	ApiBaseUrl  string
	CreatorName string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	// config.iniからの設定読み込み
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	// .envからの設定読み込み
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:        cfg.Section("web").Key("port").MustString("8000"),
		LogFile:     cfg.Section("web").Key("logfile").String(),
		Sql:         cfg.Section("db").Key("sql").String(),
		Env:         os.Getenv("ENV"),
		DBUser:      os.Getenv("DB_USER"),
		DBName:      os.Getenv("DB_NAME"),
		DBPass:      os.Getenv("DB_PASS"),
		Protocol:    os.Getenv("PROTOCOL"),
		ApiBaseUrl:  os.Getenv("API_BASE_URL"),
		CreatorName: os.Getenv("CREATOR_NAME"),
	}
}
