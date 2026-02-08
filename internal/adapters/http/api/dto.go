package api

import (
	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
)

func calculatePacksResponseFromDto(dto entities.Order) calculatePacksResponse {
	resp := calculatePacksResponse{}

	for size, count := range dto {
		resp = append(resp, packDetails{
			Size:       int(size),
			Quantity:   count,
			TotalItems: size * count,
		})
	}

	return resp
}
