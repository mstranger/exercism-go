package perfect

import "fmt"

// Classification contains the classification values for numbers.
type Classification int

// Classification constants names
const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

// ErrOnlyPositive contains error message if number will be negative.
var ErrOnlyPositive = fmt.Errorf("only positive numbers")

// Classify determines if a number is perfect, abundant or deficient.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return -1, ErrOnlyPositive
	}

	sum := aliquotSum(n)

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}

func aliquotSum(n int64) int64 {
	var result int64
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			result += i
		}
	}
	return result
}
