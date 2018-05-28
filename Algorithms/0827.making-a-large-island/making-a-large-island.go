package problem0827

func largestIsland(grid [][]int) int {
	zeros := collectZero(grid)
	if len(zeros) <= 1 {
		return len(grid) * len(grid[0])
	}

	res := 0
	for i := range zeros {
		x, y := zeros[i][0], zeros[i][1]
		grid[x][y] = 1
		res = max(res, sizeOfLargestIsland(grid))
		grid[x][y] = 0
	}

	return res
}

func collectZero(grid [][]int) [][]int {
	res := make([][]int, 0, len(grid))
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func sizeOfLargestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	isChecked := make([]bool, m*n)

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !isChecked[i*n+j] {
				isChecked[i*n+j] = true
				res = max(res, bfs(i, j, grid, isChecked))
			}
		}
	}

	return res
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bfs(i, j int, grid [][]int, isChecked []bool) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{[]int{i, j}}
	res := 1

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		i, j := next[0], next[1]
		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]
			if 0 <= x && x < m &&
				0 <= y && y < n &&
				grid[x][y] == 1 && !isChecked[x*n+y] {
				queue = append(queue, []int{x, y})
				res++
				isChecked[x*n+y] = true
			}
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
