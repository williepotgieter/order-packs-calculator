package http

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/williepotgieter/order-packs-calculator/config"
	"github.com/williepotgieter/order-packs-calculator/internal/adapters/http/api"
	"github.com/williepotgieter/order-packs-calculator/internal/adapters/http/client"
	"github.com/williepotgieter/order-packs-calculator/internal/core/ports"
	"go.uber.org/zap"
)

type adapter struct {
	router *gin.Engine
	cfg    config.AppConfig
}

func NewAdapter(cfg config.AppConfig, logger *zap.Logger) (ports.Server, error) {
	if cfg.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Middleware
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// Setup client
	if err := client.Setup(r); err != nil {
		return nil, err
	}

	// Setup API
	api.Setup(r, logger)

	return &adapter{
		router: r,
		cfg:    cfg,
	}, nil
}

func (a *adapter) Run() error {
	return a.router.Run(fmt.Sprintf(":%d", a.cfg.Api.Port))
}
