package usecases

import (
	"cmp"
	"errors"
	"fmt"
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
	uniquePackSizes := []uint{}

	for _, size := range packSizes {
		if !rawPacks[size] && size > 0 {
			rawPacks[size] = true
			uniquePackSizes = append(uniquePackSizes, size)
		}
	}

	if len(uniquePackSizes) == 0 {
		return nil, errors.New("pack sizes must contain at least one non-zero pack size")
	}

	// Sort unique pack sizes in descending order
	slices.SortFunc(uniquePackSizes, func(a, b uint) int {
		return cmp.Compare(b, a)
	})

	fmt.Println("UNIQUE SIZES:", uniquePackSizes)

	order := make(entities.Order)
	unpackedItems := items
	for _, packSize := range uniquePackSizes {
		packs, overflow, err := createPacks(packSize, unpackedItems)
		if err != nil {
			return nil, err
		}

		order[packSize] = packs

		if overflow == 0 {
			break
		}

		unpackedItems = overflow
	}

	return order, nil
}

func createPacks(size, items uint) (packs []*entities.Pack, overflow uint, err error) {
	if size == 0 || items == 0 {
		return nil, 0, errors.New("size and items args must both be non-zero")
	}

	overflow = items
	packs = []*entities.Pack{}
	for overflow >= size {
		p, err := entities.NewPack(size)
		if err != nil {
			return nil, 0, err
		}

		overflow = p.Fill(overflow)

		packs = append(packs, p)
	}

	return packs, overflow, nil
}
