package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	Env                 string
	APIBaseURL          string
	APIBucketName       string
	S3UsePathStyle      bool
	SendGridAPIKey      string
	ContactSupportEmail string
	DBHost              string
	DBPort              int
	DBName              string
	DBUser              string
	DBPass              string
	ClientURL           string
	SessionTTL          time.Duration
}

func Load() (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	port := getEnv("PORT", "8081")
	env := getEnv("ENV", "local")
	clientURL := getEnv("CLIENT_URL", "")
	apiBaseURL := getEnv("API_BASE_URL", "")
	apiBucketName := getEnv("API_BUCKET_NAME", "")
	// AWS_ENDPOINT_URL_S3が設定されている場合はS3互換ストレージを使用するため、PathStyleを有効にする
	// (主にローカル開発環境でMinIOを使用するため)
	s3UsePathStyle := os.Getenv("AWS_ENDPOINT_URL_S3") != ""

	dbHost := getEnv("MYSQL_HOST", "127.0.0.1")
	dbPortStr := getEnv("MYSQL_PORT", "3306")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid MYSQL_PORT: %w", err)
	}

	ttlStr := getEnv("SESSION_TTL", "24h")
	ttl, err := time.ParseDuration(ttlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SESSION_TTL: %w", err)
	}
	sendGridAPIKey := getEnv("SEND_GRID_API_KEY", "")
	contactSupportEmail := getEnv("CONTACT_SUPPORT_EMAIL", "no-reply@tocoriri.com")
	cfg := &Config{
		Port:                port,
		Env:                 env,
		APIBaseURL:          apiBaseURL,
		APIBucketName:       apiBucketName,
		S3UsePathStyle:      s3UsePathStyle,
		SendGridAPIKey:      sendGridAPIKey,
		ContactSupportEmail: contactSupportEmail,
		DBHost:              dbHost,
		DBPort:              dbPort,
		DBName:              getEnv("DB_NAME", ""),
		DBUser:              getEnv("DB_USER", ""),
		DBPass:              getEnv("DB_PASS", ""),
		ClientURL:           clientURL,
		SessionTTL:          ttl,
	}
	return cfg, nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
