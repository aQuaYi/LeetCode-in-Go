package problem1034

func colorBorder(grid [][]int, x0 int, y0 int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	c := grid[x0][y0]
	isBorder := func(x, y int) bool {
		return x == 0 || x == m-1 ||
			y == 0 || y == n-1 ||
			grid[x][y-1] != c ||
			grid[x][y+1] != c ||
			grid[x+1][y] != c ||
			grid[x-1][y] != c
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == c && isBorder(i, j) {
				res[i][j] = color
			} else {
				res[i][j] = grid[i][j]
			}
		}
	}

	return res
}
