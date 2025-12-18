package logger

import (
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func New() *Logger {
	return &Logger{
		debug: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
		info:  log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime),
		warn:  log.New(os.Stdout, "[WARN]  ", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Debug(msg string) {
	l.debug.Println(msg)
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Warn(msg string) {
	l.warn.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.error.Println(msg)
}
