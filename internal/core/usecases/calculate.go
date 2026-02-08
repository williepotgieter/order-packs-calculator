package usecases

import (
	"math"
	"slices"

	"github.com/williepotgieter/order-packs-calculator/internal/core/entities"
)

// CalculateOrderPacks is an adapted version of the "coin change" algorithm applied
// in Dynamic Programming (see https://en.wikipedia.org/wiki/Dynamic_programming).
// Instead of only finding the optimal number of packs required for shipping the
// required items, the algoritm also keeps track of the packs that make up the set/order
func CalculateOrderPacks(items int, packSizes []int) entities.Order {
	// packs is a utility struct that is only used within
	// the scope of the CalculateOrderPacks function
	type pack struct {
		count    int // Number of packs needed to fill the order
		previous int // Previously used pack
	}

	// Sort pack sizes in ascending order
	slices.Sort(packSizes)

	// Determine the "worst case" scenario for the end
	// of the set if the lasgest provided pack is used.
	limit := items + packSizes[len(packSizes)-1]

	// Create a slice for storing pack objects that will
	// be used to re-generate the order set
	dp := make([]pack, limit+1)
	for i := range dp {
		dp[i] = pack{
			count:    math.MaxInt32,
			previous: -1,
		}
	}

	// Dynamic Programming concept where the first item is set to zero.
	dp[0] = pack{
		count: 0,
	}

	// Build all reachable totals for finding the number of packages
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

	// Pick the best total number of paks ≥ order items
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

	// Reconstruct optimal packs set using the pack sructuer
	result := make(entities.Order)
	for curr := bestTotal; curr > 0; {
		p := dp[curr].previous
		result[p]++
		curr -= p
	}

	return result
}
