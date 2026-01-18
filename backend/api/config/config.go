package config

import (
	"api/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port           string
	Sql            string
	LogFile        string
	Env            string
	DBUser         string
	DBName         string
	DBPass         string
	Protocol       string
	ClientUrl      string
	ApiBaseUrl     string
	CreatorName    string
	ApiBucketName  string
	SendGridApiKey string
}

var Config ConfigList

func init() {
	LoadConfig()
	if Config.Env == "local" {
		if _, err := os.Stat(Config.LogFile); err != nil {
			os.Create(Config.LogFile)
		}
		utils.LoggingSettings(Config.LogFile)
	}
}

func LoadConfig() {
	// config.iniからの設定読み込み
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	// localの場合.envから設定読み込み
	if os.Getenv("ENV") == "local" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalln(err)
		}
	}

	protocol := fmt.Sprintf("tcp(%s)", os.Getenv("MYSQL_HOST"))
	Config = ConfigList{
		Port:           cfg.Section("web").Key("port").MustString("8000"),
		LogFile:        cfg.Section("web").Key("logfile").String(),
		Sql:            cfg.Section("db").Key("sql").String(),
		Env:            os.Getenv("ENV"),
		DBUser:         os.Getenv("DB_USER"),
		DBName:         os.Getenv("DB_NAME"),
		DBPass:         os.Getenv("DB_PASS"),
		Protocol:       protocol,
		ClientUrl:      os.Getenv("CLIENT_URL"),
		ApiBaseUrl:     os.Getenv("API_BASE_URL"),
		CreatorName:    os.Getenv("CREATOR_NAME"),
		ApiBucketName:  os.Getenv("API_BUCKET_NAME"),
		SendGridApiKey: os.Getenv("SEND_GRID_API_KEY"),
	}
}
