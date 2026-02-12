package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/tokushun109/tku/backend/infrastructure"
	"github.com/tokushun109/tku/backend/infrastructure/database"
	"github.com/tokushun109/tku/backend/infrastructure/log"
	"github.com/tokushun109/tku/backend/infrastructure/router"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	infrastructure.NewConfig().
		Logger(log.InstanceStdLogger).
		DbSQL(database.InstanceMySQL).
		WebServerPort(port).
		WebServer(router.InstanceGorillaMux).
		Start()
}
