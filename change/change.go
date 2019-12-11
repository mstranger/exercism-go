package change

import "errors"

// Change determines the fewest number of coins changes.
func Change(coins []int, amount int) ([]int, error) {
	out := make([]int, 0)

	for amount > 0 {
		// amount less than the smalles coin
		if amount < coins[0] {
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

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// func main() {
// 	d := []int{5, 10, 25, 50}
// 	input := 4
// 	fmt.Println(Change(d, input))
// }
