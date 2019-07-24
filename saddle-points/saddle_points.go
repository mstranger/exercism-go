package matrix

import "sort"

type Pair struct {
	r, c int
}

func (m Matrix) Saddle() []Pair {
	res := make([]Pair, 0)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			_, maxRow := minmaxIntSlice(m.Rows()[i])
			minCol, _ := minmaxIntSlice(m.Cols()[j])

			if m[i][j] == maxRow && m[i][j] == minCol {
				res = append(res, Pair{i, j})
			}
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}

func minmaxIntSlice(v []int) (int, int) {
	t := make([]int, len(v))
	copy(t, v)
	sort.Ints(t)
	return t[0], t[len(t)-1]
}

const testVersion = 1
