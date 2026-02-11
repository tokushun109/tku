package main

import (
	"os"

	"github.com/tokushun109/tku/backend/infrastructure"
	"github.com/tokushun109/tku/backend/infrastructure/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	infrastructure.NewConfig().
		WebServerPort(port).
		WebServer(router.InstanceGorillaMux).
		Start()
}
