package logs

import "log"

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type AppLogger struct{}

func (l *AppLogger) Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func (l *AppLogger) Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}
