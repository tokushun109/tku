package log

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	l.info.Print(sanitizeForLog(fmt.Sprintf(format, args...)))
}

func (l *stdLogger) Warnf(format string, args ...any) {
	l.warn.Print(sanitizeForLog(fmt.Sprintf(format, args...)))
}

func (l *stdLogger) Errorf(format string, args ...any) {
	l.error.Print(sanitizeForLog(fmt.Sprintf(format, args...)))
}

func (l *stdLogger) Debugf(format string, args ...any) {
	l.debug.Print(sanitizeForLog(fmt.Sprintf(format, args...)))
}

func sanitizeForLog(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' {
			return ' '
		}
		// ログ注入対策としてASCII制御文字(0x00-0x1F, 0x7F)を空白に置換する。
		if r < 32 || r == 127 {
			return ' '
		}
		return r
	}, s)
}
