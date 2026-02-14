package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBHost    string
	DBPort    int
	DBName    string
	DBUser    string
	DBPass    string
	ClientURL string
}

func Load() (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	port := getEnv("PORT", "8081")
	clientURL := getEnv("CLIENT_URL", "")

	dbHost := getEnv("MYSQL_HOST", "127.0.0.1")
	dbPortStr := getEnv("MYSQL_PORT", "3306")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid MYSQL_PORT: %w", err)
	}
	cfg := &Config{
		Port:      port,
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBName:    getEnv("DB_NAME", ""),
		DBUser:    getEnv("DB_USER", ""),
		DBPass:    getEnv("DB_PASS", ""),
		ClientURL: clientURL,
	}
	return cfg, nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
