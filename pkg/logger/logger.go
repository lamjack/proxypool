package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field
type ObjectEncoder = zapcore.ObjectEncoder

type Logger interface {
	Debug(string, ...Field)
	Error(string, ...Field)
	Info(string, ...Field)
	Panic(string, ...Field)
	Sync() error
	With(fields ...Field) *zap.Logger
}

func New(cfg Config) Logger {
	var logger Logger
	if cfg.GetLevel() == DevelopmentLevel {
		logger, _ = NewDevelopmentLogger()
	} else {
		logger, _ = NewProductionLogger()
	}
	return logger
}

func String(k, v string) Field {
	return zap.String(k, v)
}

func NewProductionLogger() (Logger, error) {
	return zap.NewProduction()
}

func NewDevelopmentLogger() (Logger, error) {
	return zap.NewDevelopment()
}
