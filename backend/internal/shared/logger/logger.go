package logger

import (
	"io"
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "", log.LstdFlags)
	warnLogger  = log.New(os.Stdout, "", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "", log.LstdFlags)
)

func Infof(format string, args ...any) {
	infoLogger.Printf("[INFO] "+format, args...)
}

func Warnf(format string, args ...any) {
	warnLogger.Printf("[WARN] "+format, args...)
}

func Errorf(format string, args ...any) {
	errorLogger.Printf("[ERROR] "+format, args...)
}

func Fatalf(format string, args ...any) {
	errorLogger.Fatalf("[ERROR] "+format, args...)
}

func SetOutputs(stdout, stderr io.Writer) {
	if stdout == nil {
		stdout = os.Stdout
	}
	if stderr == nil {
		stderr = os.Stderr
	}

	infoLogger.SetOutput(stdout)
	warnLogger.SetOutput(stdout)
	errorLogger.SetOutput(stderr)
}

func SetFlags(flags int) {
	infoLogger.SetFlags(flags)
	warnLogger.SetFlags(flags)
	errorLogger.SetFlags(flags)
}

func Reset() {
	SetOutputs(os.Stdout, os.Stderr)
	SetFlags(log.LstdFlags)
}
