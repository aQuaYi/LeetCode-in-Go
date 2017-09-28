package Problem0200

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	x := make([]int, 0, m*n)
	y := make([]int, 0, m*n)

	var bfs = func(col, row int) int {
		if grid[col][row] == '0' {
			return 0
		}
		x = append(x, col)
		y = append(y, row)
		grid[col][row] = '0'

		for len(x) > 0 {
			i := x[0]
			x = x[1:]
			j := y[0]
			y = y[1:]

			if 0 <= i-1 && grid[i-1][j] == '1' {
				x = append(x, i-1)
				y = append(y, j)
				grid[i-1][j] = '0'
			}

			if i+1 < m && grid[i+1][j] == '1' {
				x = append(x, i+1)
				y = append(y, j)
				grid[i+1][j] = '0'
			}

			if 0 <= j-1 && grid[i][j-1] == '1' {
				x = append(x, i)
				y = append(y, j-1)
				grid[i][j-1] = '0'
			}

			if j+1 < n && grid[i][j+1] == '1' {
				x = append(x, i)
				y = append(y, j+1)
				grid[i][j+1] = '0'
			}
		}

		return 1
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += bfs(i, j)
		}
	}

	return res
}
