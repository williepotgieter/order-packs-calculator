package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williepotgieter/order-packs-calculator/internal/core/usecases"
)

func handleCalculatePacks(c *gin.Context) {
	order, err := usecases.CalculateOrderPacks(500000, 23, 31, 53)
	if err != nil {
		httpError(c, http.StatusInternalServerError, "error while calculating order packs", err)
		return
	}

	c.JSON(http.StatusOK, calculatePacksResponseFromDto(order))
}
