package router

import (
	"errors"

	"github.com/tokushun109/tku/backend/adapter/logger"
	"github.com/tokushun109/tku/backend/adapter/repository"
)

type Server interface {
	Listen()
}

type Port string

const (
	InstanceGorillaMux = iota
)

func NewWebServerFactory(instance int, log logger.Logger, db repository.SQLDB, port string) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return NewGorillaMuxServer(log, db, Port(port)), nil
	default:
		return nil, errors.New("invalid web server instance")
	}
}
