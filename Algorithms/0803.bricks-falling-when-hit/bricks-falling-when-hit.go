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
			res[idx] += check(x, y, grid)
		}
	}
	return res
}

func isHanging(i, j int, isVisited []bool, grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	if i < 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] == 0 {
		return false
	}

	isVisited[i*n+j] = true

	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x == 0 && grid[x][y] == 1 {
			return true
		}
		isVisited[x*n+y] = true
		if isHanging(x, y, isVisited, grid) {
			return true
		}
	}

	return false
}

func check(i, j int, grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if i <= 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] == 0 {
		return 0
	}

	isVisited := make([]bool, m*n)
	isVisited[i*n+j] = true
	locations := make([][]int, 1, m*n)
	locations[0] = []int{i, j}
	cur := 0

	for cur < len(locations) {
		i = locations[cur][0]
		j = locations[cur][1]

		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]

			if x == 0 && grid[x][y] == 1 {
				return 0
			}

			if 0 < x && x < m && 0 <= y && y < n && !isVisited[x*n+y] && grid[x][y] == 1 {
				isVisited[x*n+y] = true
				locations = append(locations, []int{x, y})
			}

		}

		cur++
	}

	for _, loc := range locations {
		grid[loc[0]][loc[1]] = 0
	}

	return len(locations)
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}
