package infrastructure

import (
	"github.com/tokushun109/tku/backend/infrastructure/router"
)

type config struct {
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

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(instance, c.webServerPort)
	if err != nil {
		panic(err)
	}
	c.webServer = s
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
