package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williepotgieter/order-packs-calculator/internal/core/usecases"
)

func handleCalculatePacks(c *gin.Context) {
	payload := new(calculatePacksRequest)

	if err := c.Bind(payload); err != nil {
		httpError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	order := usecases.CalculateOrderPacks(payload.Items, payload.Packs)

	c.JSON(http.StatusOK, calculatePacksResponseFromDto(order))
}
