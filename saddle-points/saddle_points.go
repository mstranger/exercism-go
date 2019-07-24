package matrix

// Pair is a pair of ints
type Pair struct {
	r, c int
}

// Saddle finds the saddle points in a matrix
func (m Matrix) Saddle() []Pair {
	res := make([]Pair, 0)
	cols := m.Cols()

	for i, row := range m {
		for j, cell := range row {
			if is(cell, row, max) && is(cell, cols[j], min) {
				res = append(res, Pair{i, j})
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func is(n int, arr []int, f func(int, int) int) bool {
	m := arr[0]
	for _, v := range arr {
		m = f(m, v)
	}
	return n == m
}

const testVersion = 2
