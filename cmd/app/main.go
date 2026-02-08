package main

import (
	"context"
	"log"

	"github.com/williepotgieter/order-packs-calculator/config"
	"github.com/williepotgieter/order-packs-calculator/internal/adapters/http"
)

var (
	cfg config.AppConfig
	ctx context.Context
)

func init() {
	var err error

	ctx = context.Background()

	if cfg, err = config.Load(ctx); err != nil {
		log.Fatal(err)
	}
}

func main() {
	server, err := http.NewAdapter(cfg)
	if err != nil {
		log.Fatalln("unable to setup server:", err.Error())
	}

	if err := server.Run(); err != nil {
		log.Fatalln("server experienced an error:", err.Error())
	}
}
