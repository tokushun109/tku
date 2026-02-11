package router

import (
	"errors"

	"gorm.io/gorm"
)

type Server interface {
	Listen()
}

type Port string

const (
	InstanceGorillaMux = iota
)

func NewWebServerFactory(instance int, db *gorm.DB, port string) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return NewGorillaMuxServer(db, Port(port)), nil
	default:
		return nil, errors.New("invalid web server instance")
	}
}
