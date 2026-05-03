// Package internal
package internal

import (
	"log"
	"os"
)

type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewLogger() *Logger {
	logFlags := log.Lshortfile | log.Ldate

	return &Logger{
		debugLogger:   log.New(os.Stdout, "DEBUG: ", logFlags),
		infoLogger:    log.New(os.Stdout, "INFO: ", logFlags),
		warningLogger: log.New(os.Stdout, "WARNING: ", logFlags),
		errorLogger:   log.New(os.Stderr, "ERROR: ", logFlags),
	}
}

func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}
