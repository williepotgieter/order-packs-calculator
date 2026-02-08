package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
)

func TestCalculatePacksResponseFromDto(t *testing.T) {
	cases := []struct {
		name string
		dto  entities.Order
		want calculatePacksResponse
	}{
		{
			name: "happy path",
			dto: entities.Order{
				23: 2,
				31: 7,
				53: 9429,
			},
			want: calculatePacksResponse{
				{
					Size:       23,
					Quantity:   2,
					TotalItems: 46,
				},
				{
					Size:       31,
					Quantity:   7,
					TotalItems: 217,
				},
				{
					Size:       53,
					Quantity:   9429,
					TotalItems: 499737,
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePacksResponseFromDto(tt.dto)

			assert.Equal(t, tt.want, got)
		})
	}
}
