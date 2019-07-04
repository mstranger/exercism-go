package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a 2-dim matrix of numbers.
type Matrix [][]int

// New creates a new matrix from a given string.
func New(s string) (Matrix, error) {
	var m Matrix

	for _, row := range strings.Split(s, "\n") {
		r := []int{}
		for _, val := range strings.Fields(row) {
			v, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}

			r = append(r, v)
		}

		m = append(m, r)
	}

	if err := checkErrors(m); err != nil {
		return nil, err
	}

	return m, nil
}

// Rows return matrix rows (independent copy).
func (m Matrix) Rows() [][]int {
	mc := make(Matrix, 0)
	for i := 0; i < len(m); i++ {
		r := make([]int, len(m[i]))
		copy(r, m[i])
		mc = append(mc, r)
	}
	return mc
}

// Cols return matrix columns (independent copy).
func (m Matrix) Cols() [][]int {
	l := len(m[0])
	t := make([][]int, l)

	for i := 0; i < l; i++ {
		for j := 0; j < len(m); j++ {
			t[i] = append(t[i], m[j][i])
		}
	}

	return t
}

// Set sets the new value for the given position in the matrix.
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(*m) || col >= len((*m)[0]) {
		return false
	}

	(*m)[row][col] = val

	return true
}

func checkErrors(m Matrix) error {
	for i := 0; i < len(m)-1; i++ {
		if len(m[i]) != len(m[i+1]) {
			return fmt.Errorf("invalid rows")
		}
	}

	return nil
}
