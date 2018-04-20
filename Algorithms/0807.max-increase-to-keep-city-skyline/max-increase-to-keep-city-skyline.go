package problem0807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)

	maxRow := make([]int, n)
	maxCol := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maxRow[i] = max(maxRow[i], grid[i][j])
			maxCol[j] = max(maxCol[j], grid[i][j])
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g := grid[i][j]
			res += max(g, min(maxRow[i], maxCol[j])) - g
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
