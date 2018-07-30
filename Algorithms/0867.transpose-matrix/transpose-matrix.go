package problem0867

func transpose(A [][]int) [][]int {
	m, n := len(A), len(A[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = A[i][j]
		}
	}

	return res
}
