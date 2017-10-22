package Problem0329

func longestIncreasingPath(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	m, n := len(mat), len(mat[0])

	// path[i][j] 从 (i,j) 点出发，能到达的最长路径
	path := make([][]int, m)
	visited := make([][]bool, m)
	for i := range path {
		path[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	res := 0
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if visited[x][y] {
			return path[x][y]
		}

		path[x][y] = 1

		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && mat[x][y] < mat[nx][ny] {
				path[x][y] = max(path[x][y], dfs(nx, ny)+1)
			}
		}

		visited[x][y] = true
		return path[x][y]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j))
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
