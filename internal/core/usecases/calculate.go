package usecases

import (
	"math"
	"slices"

	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
)

type pack struct {
	count    int // how many packs needed to reach i
	previous int // which pack was used last
}

func CalculateOrderPacks(items int, packSizes []int) entities.Order {
	slices.Sort(packSizes)

	limit := items + packSizes[len(packSizes)-1]

	// dp[i] = best way to reach exactly i items
	dp := make([]pack, limit+1)
	for i := range dp {
		dp[i] = pack{
			count:    math.MaxInt32,
			previous: -1,
		}
	}

	dp[0] = pack{
		count: 0,
	}

	// Build all reachable totals
	for i := 0; i <= limit; i++ {
		if dp[i].count == math.MaxInt32 {
			continue // unreachable
		}

		for _, p := range packSizes {
			next := i + p
			if next <= limit && dp[i].count+1 < dp[next].count {
				dp[next].count = dp[i].count + 1
				dp[next].previous = p
			}
		}
	}

	// Pick the best total ≥ order
	bestTotal := -1
	for i := items; i <= limit; i++ {
		if dp[i].count == math.MaxInt32 {
			continue
		}
		if bestTotal == -1 ||
			i < bestTotal ||
			(i == bestTotal && dp[i].count < dp[bestTotal].count) {
			bestTotal = i
		}
	}

	// Reconstruct packs
	result := make(entities.Order)
	for curr := bestTotal; curr > 0; {
		p := dp[curr].previous
		result[p]++
		curr -= p
	}

	return result
}
