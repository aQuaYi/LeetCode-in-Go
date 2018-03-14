package problem0749

// 贪心算法
// 每次把会导致最多感染的区域围起来

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func containVirus(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	bfs := func(i, j, regionID int, isChecked []bool) {
		queue := make([][2]int, 1, m*n)
		queue[0] = [2]int{i, j}
		isChecked[i*n+j] = true

		for len(queue) > 0 {
			i, j = queue[0][0], queue[0][1]
			queue = queue[1:]
			grid[i][j] = regionID

			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < m &&
					0 <= y && y < n &&
					!isChecked[x*n+y] &&
					grid[x][y] > 0 {
					queue = append(queue, [2]int{x, y})
					isChecked[x*n+y] = true
				}
			}
		}
	}

	// 统计区域数量，并对其进行编号
	checkRegion := func() int {
		regionID := 1
		isChecked := make([]bool, m*n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] > 0 && !isChecked[i*n+j] {
					bfs(i, j, regionID, isChecked)
					regionID++
				}
			}
		}
		return regionID
	}

	// 返回，如果不修墙的话，下次会导致最多感染的区域 ID
	// 如果没有新感染的话，返回 -1
	maxRegion := func(size int) int {
		// 统计各个区域会导致感染的最大数量
		// 统计的时候，
		// 	0 如果可以被同一个区域的多个细胞感染的话，只能算作这个区域的一次感染
		//	0 如果可以被多个区域感染的话，每个区域分别多一个
		count := make([]int, size)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 0 {
					isCounted := make(map[int]bool, 4)
					for k := 0; k < 4; k++ {
						x := i + dx[k]
						y := j + dy[k]
						if 0 <= x && x < m &&
							0 <= y && y < n &&
							grid[x][y] > 0 &&
							!isCounted[grid[x][y]] {
							count[grid[x][y]]++
							isCounted[grid[x][y]] = true
						}
					}
				}
			}
		}
		// 寻找会导致最多感染的区域
		max := 0
		id := -1
		for i := 1; i < size; i++ {
			if max < count[i] {
				max = count[i]
				id = i
			}
		}

		return id
	}

	// 给指定区域安装围墙
	// 同时把区域内的所有值，标记为 -1，就能不在统计此区域
	installWalls := func(regionID int) int {
		res := 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == regionID {
					grid[i][j] = -1
					for k := 0; k < 4; k++ {
						x := i + dx[k]
						y := j + dy[k]
						if 0 <= x && x < m &&
							0 <= y && y < n &&
							grid[x][y] == 0 {
							res++
						}
					}
				}
			}
		}
		return res
	}

	// 夜间的病毒传播
	nightInfect := func() {
		isNewInfected := make([]bool, m*n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] > 0 && !isNewInfected[i*n+j] {
					for k := 0; k < 4; k++ {
						x := i + dx[k]
						y := j + dy[k]
						if 0 <= x && x < m &&
							0 <= y && y < n &&
							grid[x][y] == 0 {
							grid[x][y] = 1
							isNewInfected[x*n+y] = true
						}
					}
				}
			}
		}
	}

	res := 0
	for {
		size := checkRegion()
		id := maxRegion(size)
		if id == -1 {
			break
		}
		res += installWalls(id)
		nightInfect()
	}

	return res
}
