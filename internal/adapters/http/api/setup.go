package api

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.POST("/calculate-packs", handleCalculatePacks)
}
