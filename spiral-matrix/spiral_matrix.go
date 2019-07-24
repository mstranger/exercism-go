package main

import (
	"fmt"
)

func SpiralMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	rots := 0
	current := 1

	for rots < 4 {
		for i, v := range res[0] {
			if v != 0 {
				continue
			}

			res[0][i] = current
			current++
		}
		rotLeft(res)
		rots++
	}

	return res
}

func rotLeft(m [][]int) {
	l := len(m)

	for i := 0; i < l; i++ {
		for j := i; j >= 0; j-- {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}

	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
}

func main() {
	m := SpiralMatrix(4)
	// m := [][]int{{1, 2, 3}, {8, 9, 4}, {}}
	// fmt.Println(m)
	// rotLeft(m)
	fmt.Println(m)
}
