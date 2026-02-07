package usecases

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"

	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
)

func CalculateOrderPacks(items uint, packSizes ...uint) (entities.Order, error) {

	if items == 0 {
		return nil, errors.New("number of items must be greater than zero")
	}

	if len(packSizes) == 0 {
		return nil, errors.New("at least one pack size must be specified")
	}

	// Remove duplicate and zero packs
	rawPacks := make(map[uint]bool)
	uniquePacks := []uint{}

	for _, size := range packSizes {
		if !rawPacks[size] && size > 0 {
			rawPacks[size] = true
			uniquePacks = append(uniquePacks, size)
		}
	}

	if len(uniquePacks) == 0 {
		return nil, errors.New("pack sizes must contain at least one non-zero pack size")
	}

	// Sort unique pack sizes in descending order
	slices.SortFunc(uniquePacks, func(a, b uint) int {
		return cmp.Compare(b, a)
	})

	fmt.Println("SORTED:", uniquePacks)

	// Calculate packs
	countPacks := func(i *uint, s uint) uint {
		cnt := uint(math.Floor(float64(*i) / float64(s)))
		*i = *i - (cnt * s)

		return cnt
	}

	order := make(entities.Order)
	tmpItems := items
	for _, pack := range uniquePacks {
		order[pack] = countPacks(&tmpItems, pack)
	}

	return order, nil
}
