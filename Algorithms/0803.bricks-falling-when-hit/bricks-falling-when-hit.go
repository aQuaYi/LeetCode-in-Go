package problem0803

func hitBricks(grid [][]int, hits [][]int) []int {
	res := make([]int, len(hits))
	for _, hit := range hits {
		i, j := hit[0], hit[1]
		for k := 0; k < 4; k++ {
			res[i] += check(i+dx[k], j+dy[k], grid)
		}
	}
	return res
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

			if x == 0 {
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

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}
