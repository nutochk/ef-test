package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Config struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" `
	Port     int    `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" `
	User     string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" `
	Password string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB"`
}

func New(cfg Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	migrationsDir := "./migrations"
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	sqlDB := stdlib.OpenDB(*connConfig)
	if err = goose.Up(sqlDB, migrationsDir); err != nil {
		if !errors.Is(err, goose.ErrNoMigrations) {
			return nil, fmt.Errorf("failed to apply migrations: %w", err)
		}
	}
	return conn, nil
}
