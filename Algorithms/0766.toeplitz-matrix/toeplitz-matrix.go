package problem0766

func isToeplitzMatrix(mat [][]int) bool {
	m, n := len(mat), len(mat[0])

	for i := 0; i+1 < m; i++ {
		for j := 0; j+1 < n; j++ {
			if mat[i][j] != mat[i+1][j+1] {
				return false
			}
		}
	}

	return true
}
