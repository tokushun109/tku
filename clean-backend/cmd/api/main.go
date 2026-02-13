package main

import (
	"log"
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/app/di"
)

func main() {
	cfg, handler, err := di.BuildServer()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + cfg.Port
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
