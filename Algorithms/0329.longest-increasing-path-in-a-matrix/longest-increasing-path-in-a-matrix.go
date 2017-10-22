package Problem0329

func longestIncreasingPath(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	m, n := len(mat), len(mat[0])

	path := make([][]bool, m)
	for i := range path {
		path[i] = make([]bool, n)
	}

	res := 0
	var dfs func(int, int, int)
	dfs = func(i, j, count int) {
		if res < count {
			res = count
		}

		path[i][j] = true

		if 0 <= i-1 && !path[i-1][j] && mat[i-1][j] > mat[i][j] {
			dfs(i-1, j, count+1)
		}

		if 0 <= j-1 && !path[i][j-1] && mat[i][j-1] > mat[i][j] {
			dfs(i, j-1, count+1)
		}

		if i+1 < m && !path[i+1][j] && mat[i+1][j] > mat[i][j] {
			dfs(i+1, j, count+1)
		}

		if j+1 < n && !path[i][j+1] && mat[i][j+1] > mat[i][j] {
			dfs(i, j+1, count+1)
		}

		path[i][j] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 1)
		}
	}

	return res
}
