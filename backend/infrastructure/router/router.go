package router

import "errors"

type Server interface {
	Listen()
}

type Port string

const (
	InstanceGorillaMux = iota
)

func NewWebServerFactory(instance int, port string) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return NewGorillaMuxServer(Port(port)), nil
	default:
		return nil, errors.New("invalid web server instance")
	}
}
