package problem0827

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 收集 grid 中的 0
	zeros := collectZero(grid)
	if len(zeros) <= 1 {
		return m * n
	}

	colors := addColor(grid)

	res := 0

	for _, z := range zeros {
		i, j := z[0], z[1]
		isConnected := make([]bool, len(colors))
		temp := 1
		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]
			if 0 <= x && x < m &&
				0 <= y && y < n &&
				grid[x][y] > 1 && !isConnected[grid[x][y]] {
				temp += colors[grid[x][y]]
				isConnected[grid[x][y]] = true
			}
		}
		res = max(res, temp)
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

// 返回每种色彩的数量
func addColor(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	res := make([]int, 2, m)
	color := 2

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = append(res, bfs(i, j, color, grid))
				color++
			}
		}
	}

	return res
}

func bfs(i, j, color int, grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{[]int{i, j}}
	grid[i][j] = color
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
				grid[x][y] == 1 {
				queue = append(queue, []int{x, y})
				grid[x][y] = color
				res++
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
