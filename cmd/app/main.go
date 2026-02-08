package main

import (
	"context"
	"log"

	"github.com/williepotgieter/order-packs-calculator/config"
	"github.com/williepotgieter/order-packs-calculator/internal/adapters/http"
	"go.uber.org/zap"
)

var (
	cfg    config.AppConfig
	ctx    context.Context
	logger *zap.Logger
)

func init() {
	var err error

	ctx = context.Background()

	if cfg, err = config.Load(ctx); err != nil {
		log.Fatalln(err)
	}

	if cfg.Prod {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
}

func main() {
	server, err := http.NewAdapter(cfg, logger)
	if err != nil {
		logger.Fatal("unable to setup server", zap.Error(err))
	}

	if err := server.Run(); err != nil {
		logger.Fatal("server experienced an error", zap.Error(err))
	}
}
