package change

import (
	"errors"
	"sort"
)

var (
	errNegativeChange = errors.New("change is negative")
	errNoCombinations = errors.New("no combinations")
)

// Change determines the fewest number of coins changes.
func Change(coins []int, amount int) ([]int, error) {
	if amount == 0 {
		return []int{}, nil
	}
	if amount < 0 {
		return nil, errNegativeChange
	}

	// all changes
	results := make([][]int, 0)

	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] <= amount {
			result := currentChange(coins[:i+1], amount)

			if result != nil {
				results = append(results, result)
			}
		}
	}

	// find shortest
	if len(results) > 0 {
		sort.Slice(
			results,
			func(i, j int) bool { return len(results[i]) < len(results[j]) },
		)
		return results[0], nil
	}

	// else no combinations
	return nil, errNoCombinations
}

// helper function
func currentChange(coins []int, amount int) []int {
	num := amount / coins[len(coins)-1]
	rem := amount % coins[len(coins)-1]
	crem := make([]int, 0)

	for rem != 0 {
		if len(coins) == 1 {
			return nil
		}

		crem = currentChange(coins[:len(coins)-1], rem)

		if len(crem) == 0 {
			if num <= 1 {
				return nil
			}

			num--
			rem += coins[len(coins)-1]
		} else {
			rem = 0
		}
	}

	cnum := make([]int, num)
	// append last coin `num` times
	for i := 0; i < num; i++ {
		cnum[i] = coins[len(coins)-1]
	}

	return append(crem, cnum...)
}
