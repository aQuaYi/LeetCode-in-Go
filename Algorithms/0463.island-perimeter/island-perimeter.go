package Problem0463

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += check(grid, m, n, i, j)
		}
	}

	return res
}

var ds = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

// check 返回 (i, j) 周围的周长
func check(grid [][]int, m, n, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	res := 4
	for _, d := range ds {
		x, y := i+d[0], j+d[1]
		if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
			res--
		}
	}

	return res
}
