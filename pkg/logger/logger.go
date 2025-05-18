package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func New() (*Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // Включаем Debug

	l, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}
	return &Logger{logger: l}, nil
}
