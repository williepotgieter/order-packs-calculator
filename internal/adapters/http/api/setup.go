package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Setup assigns HTTP handler functions to the appropriate API route
func Setup(router *gin.Engine, logger *zap.Logger) {
	v1 := router.Group("/api/v1")

	v1.POST("/calculate-packs", handleCalculatePacks(logger))
}
