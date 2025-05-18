package logger

import (
	"fmt"
	"net/http"
	"time"

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

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
	body   []byte
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	if r.status >= 400 { // Сохраняем тело только для ошибок
		r.body = make([]byte, len(b))
		copy(r.body, b)
	}
	size, err := r.ResponseWriter.Write(b)
	r.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.status = statusCode
}

func LogMiddleware(l Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			lrw := loggingResponseWriter{
				ResponseWriter: w,
				size:           0,
				status:         200,
			}
			next.ServeHTTP(&lrw, r)
			duration := time.Since(start)

			if lrw.status >= 400 {
				l.Error("Request failed:",
					zap.String("method", r.Method),
					zap.String("url", r.URL.Path),
					zap.Int("status", lrw.status),
					zap.Duration("duration", duration),
					zap.ByteString("response_body", lrw.body),
				)
			} else {
				l.Info("Request completed:", zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", lrw.status), zap.Duration("duration", duration))
			}
		}
		return http.HandlerFunc(fn)
	}
}
