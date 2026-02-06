package http

import (
	"github.com/gin-gonic/gin"
	"github.com/williepotgieter/order-packs-calculator/internal/adapters/http/client"
	"github.com/williepotgieter/order-packs-calculator/internal/core/ports"
)

type adapter struct {
	router *gin.Engine
}

func NewAdapter() (ports.Server, error) {
	r := gin.Default()

	// Setup client
	if err := client.Setup(r); err != nil {
		return nil, err
	}

	return &adapter{
		router: r,
	}, nil
}

func (a *adapter) Run() error {
	return a.router.Run(":3000")
}
