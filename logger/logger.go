package logger

import (
	azer_logger "github.com/azer/logger"
)

type Logger interface {
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type logger struct{}

func CreateLogger(key string) Logger {
	return azer_logger.New(key)
}

func (logger *logger) Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

func (logger *logger) Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}
