package infrastructure

import (
	"github.com/tokushun109/tku/backend/adapter/logger"
	"github.com/tokushun109/tku/backend/adapter/repository"
	"github.com/tokushun109/tku/backend/infrastructure/database"
	"github.com/tokushun109/tku/backend/infrastructure/log"
	"github.com/tokushun109/tku/backend/infrastructure/router"
)

type config struct {
	dbSQL         repository.SQLDB
	logger        logger.Logger
	webServerPort string
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WebServerPort(port string) *config {
	c.webServerPort = port
	return c
}

func (c *config) Logger(instance int) *config {
	l, err := log.NewLoggerFactory(instance)
	if err != nil {
		panic(err)
	}
	c.logger = l
	return c
}

func (c *config) DbSQL(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		panic(err)
	}
	c.dbSQL = db
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(instance, c.logger, c.dbSQL, c.webServerPort)
	if err != nil {
		panic(err)
	}
	c.webServer = s
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
