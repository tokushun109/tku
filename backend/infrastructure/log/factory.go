package log

import (
	"errors"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

const (
	InstanceStdLogger = iota
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceStdLogger:
		return NewStdLogger(), nil
	default:
		return nil, errors.New("invalid logger instance")
	}
}
