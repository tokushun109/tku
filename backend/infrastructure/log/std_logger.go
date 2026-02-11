package log

import (
	"log"
	"os"

	"github.com/tokushun109/tku/backend/adapter/logger"
)

type stdLogger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	debug *log.Logger
}

func NewStdLogger() logger.Logger {
	return &stdLogger{
		info:  log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		warn:  log.New(os.Stdout, "[WARN] ", log.LstdFlags),
		error: log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debug: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
	}
}

func (l *stdLogger) Infof(format string, args ...any) {
	l.info.Printf(format, args...)
}

func (l *stdLogger) Warnf(format string, args ...any) {
	l.warn.Printf(format, args...)
}

func (l *stdLogger) Errorf(format string, args ...any) {
	l.error.Printf(format, args...)
}

func (l *stdLogger) Debugf(format string, args ...any) {
	l.debug.Printf(format, args...)
}
