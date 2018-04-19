package problem0803

func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	res := make([]int, len(hits))

	for i := range hits {
		// 先把所有的点都消去
		// 使用 -- 而不是 =0
		// 是为了让原先是 0 的点变成 -1
		// 这很重要
		grid[hits[i][0]][hits[i][1]]--
	}

	// 全部点消去后，仍然挂着的点，全部变成 2
	for j := 0; j < len(grid[0]); j++ {
		if grid[0][j] == 1 {
			check(0, j, grid)
		}
	}

	// 倒着往上面添加点
	for idx := len(hits) - 1; 0 <= idx; idx-- {
		i, j := hits[idx][0], hits[idx][1]
		// 先把 (i,j) 还原
		grid[i][j]++
		if grid[i][j] == 0 {
			// 原先就是 0 的点，不会引起别的点下落
			continue
		}
		// 然后，检查 (i,j) 是否和 2 相连
		// 或者
		// i == 0，表示 (i,j) 本身，就在顶部
		if (0 <= i-1 && grid[i-1][j] == 2) ||
			(0 <= j-1 && grid[i][j-1] == 2) ||
			(i+1 < m && grid[i+1][j] == 2) ||
			(j+1 < n && grid[i][j+1] == 2) ||
			i == 0 {
			// 所有和 (i,j) 相连的值为 1 的点，都变成 2
			// 这些点在 (i,j) 被擦除后，都会掉落
			// 但是这些点中也包含了 (i,j) ，所以结果要 -1
			res[idx] = add(i, j, grid) - 1
		}
	}

	return res
}

func check(i, j int, grid [][]int) {
	grid[i][j] = 2
	m, n := len(grid), len(grid[0])
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
			check(x, y, grid)
		}
	}
}

func add(i, j int, grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if i < 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 2
	res := 1

	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		res += add(x, y, grid)
	}

	return res
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}
