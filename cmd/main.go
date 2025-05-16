package main

import (
	"context"

	"github.com/nutochk/ef-test/internal/config"
	"github.com/nutochk/ef-test/pkg/logger"
	"github.com/nutochk/ef-test/pkg/postgres"
	"go.uber.org/zap"
)

func main() {
	logger, err := logger.New()
	if err != nil {
		panic(err)
	}
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to read config", zap.Error(err))
	}
	logger.Debug("config content", zap.Any("config", cfg))
	pgConn, err := postgres.New(cfg.Postgres)
	if err != nil {
		logger.Fatal("failed to connect to postgres", zap.Error(err))
	} else {
		logger.Debug("connected to postgres successfully")
	}
	defer pgConn.Close(context.Background())
}
