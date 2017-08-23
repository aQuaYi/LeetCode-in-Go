package Problem0073

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	temp := make([][]int, m)
	for i := range temp {
		temp[i] = append([]int{}, matrix[i]...)
	}

	hasZero := func(r, c int) bool {
		for i := 0; i < m; i++ {
			if temp[i][c] == 0 {
				return true
			}
		}

		for j := 0; j < n; j++ {
			if temp[r][j] == 0 {
				return true
			}
		}

		return false
	}

	for i := range matrix {
		for j := range matrix[0] {
			if temp[i][j] != 0 && hasZero(i, j) {
				matrix[i][j] = 0
			}
		}
	}
}
