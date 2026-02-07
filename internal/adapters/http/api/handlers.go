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

	order, err := usecases.CalculateOrderPacks(payload.Items, payload.Packs...)
	if err != nil {
		httpError(c, http.StatusInternalServerError, "error while calculating order packs", err)
		return
	}

	c.JSON(http.StatusOK, calculatePacksResponseFromDto(order))
}
