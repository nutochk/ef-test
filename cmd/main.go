package main

import (
	"context"
	"log"

	"github.com/nutochk/ef-test/internal/config"
	"github.com/nutochk/ef-test/pkg/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
	pgConn, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to postgres")
	defer pgConn.Close(context.Background())
}
