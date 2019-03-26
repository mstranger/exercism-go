package hamming

import "errors"

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Different length")
	}

	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
