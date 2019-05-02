package allyourbase

import (
	"fmt"
	"math"
)

// ConvertToBase converts a number, represented as a sequence of digits
// in one base, to any other base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, fmt.Errorf("input base must be >= 2")
	}

	if outputBase < 2 {
		return []int{}, fmt.Errorf("output base must be >= 2")
	}

	n, err := convertToDecimal(inputBase, inputDigits)
	if err != nil {
		return []int{}, err
	}

	output, err := convertFromDecimal(outputBase, n)
	if err != nil {
		return []int{}, err
	}

	if len(output) == 0 {
		output = []int{0}
	}

	return output, nil
}

// converts from inputBase to decimal number.
func convertToDecimal(inputBase int, inputDigits []int) (int, error) {
	decimal := 0
	for i, v := range inputDigits {
		if v < 0 || v >= inputBase {
			return 0, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		decimal += v * power(inputBase, len(inputDigits)-i-1)
	}

	return decimal, nil
}

// converts the decimal number to some other base
func convertFromDecimal(outputBase int, decimal int) ([]int, error) {
	output := make([]int, 0)

	for decimal > 0 {
		d, r := decimal/outputBase, decimal%outputBase
		output = append(output, r)
		decimal = d
	}

	reverse(output)
	return output, nil
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func power(a, b int) int { return int(math.Pow(float64(a), float64(b))) }
