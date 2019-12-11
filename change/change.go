package main

import (
	"errors"
	"fmt"
)

// Change determines the fewest number of coins changes.
func Change(coins []int, amount int) ([]int, error) {
	minLen := coins[len(coins)-1]
	result := []int{}

	sets := allSubsets(coins)

	for _, set := range sets {
		ch, err := currentChange(set, amount)
		if err != nil {
			continue
		}
		if len(ch) < minLen {
			minLen, result = len(ch), ch
		}
	}

	if len(result) == 0 {
		return result, errors.New("can't be changed")
	}

	return result, nil
}

// find change variant for the given coins
func currentChange(coins []int, amount int) ([]int, error) {
	out := make([]int, 0)

	for amount > 0 {
		// amount less than the smalles coin
		// or coins are absent
		if len(coins) == 0 || amount < coins[0] {
			return nil, errors.New("can't be changed")
		}
		for i := range coins {
			current := coins[len(coins)-i-1]
			if amount/current > 0 {
				for t := 0; t < amount/current; t++ {
					out = append(out, current)
				}
				amount = amount % current
				break
			}
		}
	}

	reverse(out)
	return out, nil
}

// find all subsets of coins
func allSubsets(coins []int) [][]int {
	n := len(coins)
	sets := make([][]int, 0)

	for i := 0; i < (1 << n); i++ {
		set := []int{}
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				set = append(set, coins[j])
			}
		}
		sets = append(sets, set)
	}

	return sets
}

// reverse slice in place
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	d := []int{2, 5, 10, 20, 50}
	i := 21

	fmt.Println(currentChange(d, i))

	// fmt.Println(Change(d, i))
}
