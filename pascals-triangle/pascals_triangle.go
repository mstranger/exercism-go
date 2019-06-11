package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	tr := make([][]int, 0)
	for i := 1; i <= n; i++ {
		tr = append(tr, nextLine(&tr))
	}

	return tr
}

// returns next line for the given triangle
func nextLine(triangle *[][]int) []int {
	n := len(*triangle)

	if n == 0 {
		return []int{1}
	}

	line := make([]int, n+1)
	line[0], line[n] = 1, 1

	for i := 1; i < n; i++ {
		line[i] = (*triangle)[n-1][i-1] + (*triangle)[n-1][i]
	}

	return line
}
