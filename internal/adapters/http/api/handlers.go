package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williepotgieter/order-packs-calculator/internal/core/usecases"
)

type calculatePacksRequest struct {
	Items uint   `json:"items"`
	Packs []uint `json:"packs"`
}

type packDetails struct {
	Size     uint `json:"size"`
	Quantity uint `json:"quantity"`
}

type calculatePacksResponse []packDetails

func handleCalculatePacks(c *gin.Context) {
	orderMap, err := usecases.CalculateOrderPacks(25850, 500, 250, 1000, 0)
	if err != nil {
		httpError(c, http.StatusInternalServerError, "error while calculating order packs", err)
		return
	}

	// Transform domain struct to response
	order := calculatePacksResponse{}
	for size, qty := range orderMap {
		order = append(order, packDetails{
			Size:     size,
			Quantity: qty,
		})
	}

	c.JSON(http.StatusOK, order)
}
