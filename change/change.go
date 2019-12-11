package change

import "errors"

// Change determines the fewest number of coins changes.
func Change(coins []int, amount int) ([]int, error) {
	minLen := amount
	result := []int{}

	if amount == 0 {
		return result, nil
	}

	sets := allSubsets(coins)

	for _, set := range sets {
		// change for the current set
		ch, err := currentChange(set, amount)
		if err != nil {
			continue
		}
		// check shortest
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
		// amount less than the smallest coin
		// or coins are absent
		if len(coins) == 0 || amount < coins[0] {
			return nil, errors.New("can't be changed")
		}
		for i := range coins {
			// select coin
			current := coins[len(coins)-i-1]
			rest := amount % current

			// if the rest less than smallest coin but more than 0
			// try find other combinations
			if len(coins) > 1 && rest < coins[0] && rest > 0 {
				out = append(out, current)
				f, err := tryFind(coins[:2], amount-current)
				if err != nil {
					return nil, errors.New("not find")
				}

				return append(f, out...), nil
			}

			if amount/current > 0 {
				for t := 0; t < amount/current; t++ {
					out = append(out, current)
				}
				amount = rest
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

// only for the two elements
// try find (i, j) in the equation a*i + b*j = value
func tryFind(coins []int, value int) ([]int, error) {
	for i := 0; ; i++ {
		if i*coins[0] > value {
			break
		}
		for j := 0; ; j++ {
			if i*coins[0]+j*coins[1] == value {
				return write([2]int{i, coins[0]}, [2]int{j, coins[1]}), nil
			}
			if j*coins[1] > value {
				break
			}
		}
	}

	return nil, errors.New("not find")
}

// return slice like [2 2 2 5] from input data [[3, 2], [1, 5]]
func write(a [2]int, b [2]int) []int {
	out := []int{}

	for i := 0; i < a[0]; i++ {
		out = append(out, a[1])
	}
	for j := 0; j < b[0]; j++ {
		out = append(out, b[1])
	}

	return out
}
