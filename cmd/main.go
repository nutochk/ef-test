package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/nutochk/ef-test/docs"
	"github.com/nutochk/ef-test/internal/config"
	"github.com/nutochk/ef-test/internal/repository"
	"github.com/nutochk/ef-test/internal/server"
	"github.com/nutochk/ef-test/internal/service"
	"github.com/nutochk/ef-test/pkg/logger"
	"github.com/nutochk/ef-test/pkg/postgres"
	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

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
	}
	logger.Debug("connected to postgres successfully")

	repo := repository.NewRepo(pgConn)

	apiService := service.New(repo, *logger)
	apiServer := server.New(apiService)

	go func() {
		logger.Info("Server is listening on port:" + strconv.Itoa(cfg.Port))
		if err := apiServer.Run(cfg.Port); err != nil {
			stop()
		}
	}()

	<-ctx.Done()
	logger.Info("Shutting down server...")
	apiServer.Shutdown(ctx)
	logger.Info("Server shut down")
	defer pgConn.Close(context.Background())
}
