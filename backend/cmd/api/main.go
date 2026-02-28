package main

import (
	"net/http"

	"github.com/tokushun109/tku/backend/internal/app/di"
	"github.com/tokushun109/tku/backend/internal/shared/logger"
)

func main() {
	cfg, handler, err := di.BuildServer()
	if err != nil {
		logger.Fatalf("%v", err)
	}

	addr := ":" + cfg.Port
	logger.Infof("server starting on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Fatalf("%v", err)
	}
}
