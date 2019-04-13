package lsproduct

import (
	"fmt"
	"strconv"
)

// LargestSeriesProduct calculates the largest product for a substring
// of digits of given length.
func LargestSeriesProduct(digist string, span int) (int64, error) {
	var product int64

	if len(digist) < span {
		return 0, fmt.Errorf("span must be smaller than string length")
	}

	if span < 0 {
		return 0, fmt.Errorf("span must be greater than zero")
	}

	for i := 0; i <= len(digist)-span; i++ {
		t, err := productString(digist[i : i+span])
		if err != nil {
			return 0, err
		}
		if product < t {
			product = t
		}
	}

	return product, nil
}

func productString(s string) (int64, error) {
	var product int64 = 1
	for _, v := range s {
		n, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0, err
		}
		product *= n
	}

	return product, nil
}
