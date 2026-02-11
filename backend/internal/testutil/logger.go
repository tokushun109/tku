package testutil

import "fmt"

type Logger struct {
	Infos  []string
	Warns  []string
	Errors []string
	Debugs []string
}

func (l *Logger) Infof(format string, args ...any) {
	l.Infos = append(l.Infos, fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...any) {
	l.Warns = append(l.Warns, fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...any) {
	l.Errors = append(l.Errors, fmt.Sprintf(format, args...))
}

func (l *Logger) Debugf(format string, args ...any) {
	l.Debugs = append(l.Debugs, fmt.Sprintf(format, args...))
}
