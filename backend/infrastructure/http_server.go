package infrastructure

import (
	"github.com/tokushun109/tku/backend/infrastructure/database"
	"github.com/tokushun109/tku/backend/infrastructure/router"

	"gorm.io/gorm"
)

type config struct {
	dbSQL         *gorm.DB
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

func (c *config) DbSQL(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		panic(err)
	}
	c.dbSQL = db
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(instance, c.dbSQL, c.webServerPort)
	if err != nil {
		panic(err)
	}
	c.webServer = s
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
