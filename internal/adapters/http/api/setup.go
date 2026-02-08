package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(router *gin.Engine, logger *zap.Logger) {
	v1 := router.Group("/api/v1")

	v1.POST("/calculate-packs", handleCalculatePacks(logger))
}
