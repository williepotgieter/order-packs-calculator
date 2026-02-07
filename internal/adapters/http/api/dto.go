package api

import (
	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
)

func calculatePacksResponseFromDto(dto entities.Order) calculatePacksResponse {
	resp := calculatePacksResponse{}

	getTotalItems := func(packs []*entities.Pack) uint {
		if len(packs) == 0 {
			return 0
		}

		var totalItems uint = 0
		for _, pack := range packs {
			totalItems += pack.Items
		}

		return totalItems
	}

	for size, packs := range dto {
		resp = append(resp, packDetails{
			Size:       size,
			Quantity:   uint(len(dto[size])),
			TotalItems: getTotalItems(packs),
		})
	}

	return resp
}
