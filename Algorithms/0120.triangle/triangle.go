package Problem0120

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}

	for i := 1; i < m; i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][i] += triangle[i-1][i-1]
	}

	for i := 2; i < m; i++ {
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}

	min := triangle[m-1][0]

	for j := 1; j < m; j++ {
		if min > triangle[m-1][j] {
			min = triangle[m-1][j]
		}
	}

	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
