package Problem0542

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = m + n
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			checkLU(matrix, m, n, i, j)
		}
	}

	for i := m - 1; 0 <= i; i-- {
		for j := n - 1; 0 <= j; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			checkRD(matrix, m, n, i, j)
		}
	}

	return matrix
}

var dx = []int{0, -1, 1, 0}
var dy = []int{-1, 0, 0, 1}

func checkLU(mat [][]int, m, n, i, j int) {
	var x, y int
	for k := 0; k < 2; k++ {
		x = i + dx[k]
		y = j + dy[k]
		if 0 <= x && x < m &&
			0 <= y && y < n {
			mat[i][j] = min(mat[i][j], mat[x][y]+1)
		}
	}
}

func checkRD(mat [][]int, m, n, i, j int) {
	var x, y int
	for k := 2; k < 4; k++ {
		x = i + dx[k]
		y = j + dy[k]
		if 0 <= x && x < m &&
			0 <= y && y < n {
			mat[i][j] = min(mat[i][j], mat[x][y]+1)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
