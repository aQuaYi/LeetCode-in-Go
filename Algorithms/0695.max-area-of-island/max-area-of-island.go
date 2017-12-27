package Problem0695

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	isChecked := make([]bool, m*n)

	maxArea := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !isChecked[i*n+j] {
				search(grid, isChecked, i, j, &maxArea)
			}
		}
	}

	return maxArea
}

func search(grid [][]int, isChecked []bool, x, y int, maxArea *int) {
	m, n := len(grid), len(grid[0])
	area := 0
	xq := make([]int, 1, 50*50)
	yq := make([]int, 1, 50*50)
	xq[0] = x
	yq[0] = y
	// 添加新点到 queue 后，马上更新 isCehcked 和 area
	isChecked[x*n+y] = true
	area++

	for len(xq) > 0 {
		x = xq[0]
		y = yq[0]
		xq = xq[1:]
		yq = yq[1:]

		// 添加 [x,y] 周围还没有检查的点
		for k := 0; k < 4; k++ {
			i := x + dx[k]
			j := y + dy[k]
			if 0 <= i && i < m && 0 <= j && j < n &&
				grid[i][j] == 1 &&
				!isChecked[i*n+j] {

				xq = append(xq, i)
				yq = append(yq, j)
				// 添加新点到 queue 后，马上更新 isCehcked 和 area
				isChecked[i*n+j] = true
				area++
			}
		}

	}

	*maxArea = max(*maxArea, area)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
