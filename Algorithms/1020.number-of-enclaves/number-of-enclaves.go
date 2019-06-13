package problem1020

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func numEnclaves(A [][]int) int {
	m, n := len(A), len(A[0])

	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || m <= x ||
			y < 0 || n <= y ||
			A[x][y] == 0 {
			return
		}
		A[x][y] = 0
		for i := 0; i < 4; i++ {
			dfs(x+dx[i], y+dy[i])
		}
	}

	line := func(i, di, j, dj int) {
		for i < m && j < n {
			dfs(i, j)
			i += di
			j += dj
		}
	}

	// search from the boundary
	line(0, 0, 0, 1)   // top
	line(m-1, 0, 0, 1) // down
	line(0, 1, 0, 0)   // left
	line(0, 1, n-1, 0) // right

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += A[i][j]
		}
	}

	return res
}
