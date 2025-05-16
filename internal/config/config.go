package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/nutochk/ef-test/pkg/postgres"
)

type Config struct {
	Port     int `yaml:"PORT" env:"PORT"`
	Postgres postgres.Config
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	return &cfg, nil
}
