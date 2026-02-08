package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
	"github.com/williepotgieter/order-packs-calculator/internal/core/usecases"
)

func TestCalculateOrderPacks(t *testing.T) {
	cases := []struct {
		name      string
		items     int
		packSizes []int
		want      entities.Order
	}{
		{
			name:      "happy path 1",
			items:     500000,
			packSizes: []int{23, 31, 53},
			want: entities.Order{
				23: 2,
				31: 7,
				53: 9429,
			},
		},
		{
			name:      "happy path 2",
			items:     1,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			want:      entities.Order{250: 1},
		},
		{
			name:      "happy path 3",
			items:     250,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			want:      entities.Order{250: 1},
		},
		{
			name:      "happy path 4",
			items:     251,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			want:      entities.Order{500: 1},
		},
		{
			name:      "happy path 5",
			items:     501,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			want: entities.Order{
				500: 1,
				250: 1,
			},
		},
		{
			name:      "happy path 6",
			items:     12001,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			want: entities.Order{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := usecases.CalculateOrderPacks(tt.items, tt.packSizes)

			assert.Equal(t, tt.want, got)
		})
	}
}
