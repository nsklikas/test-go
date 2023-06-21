package logger

import (
	"log"
	"time"
)

type logger struct{}

// Should implement the go Logger interface?
type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
}

func (l logger) log(level string, msg string) {
	log.Println("[", level, "][", time.Now(), "]: ", msg)
}

func (l logger) Debug(msg string) {
	l.log("INFO", msg)
}

func (l logger) Info(msg string) {
	l.log("INFO", msg)
}

func (l logger) Warn(msg string) {
	l.log("WARN", msg)
}

func (l logger) Error(msg string) {
	l.log("ERROR", msg)
}

func NewLogger() logger {
	return logger{}
}
