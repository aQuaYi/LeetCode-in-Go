package problem0803

func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	res := make([]int, len(hits))

	for idx, hit := range hits {
		i, j := hit[0], hit[1]
		grid[i][j] = 0

		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]
			isVisited := make([]bool, m*n)
			if isHanging(x, y, isVisited, grid) {
				continue
			}

			res[idx] += drop(x, y, grid)
		}
	}

	return res
}

func isHanging(i, j int, isVisited []bool, grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	if i == 0 && 0 <= j && j < n && grid[i][j] == 1 {
		return true
	}

	if i < 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] == 0 ||
		isVisited[i*n+j] {
		return false
	}

	isVisited[i*n+j] = true

	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if isHanging(x, y, isVisited, grid) {
			return true
		}
	}

	return false
}

func drop(i, j int, grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if i < 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	res := 1

	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		res += drop(x, y, grid)
	}

	return res
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}
