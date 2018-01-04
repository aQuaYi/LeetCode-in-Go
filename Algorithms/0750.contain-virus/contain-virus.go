package Problem0750

// 贪心算法
// 每次把感染最多的区域围起来

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

	// 如果不修墙的话，下次会导致最多感染的区域 ID
	// 如果没有新感染的话，返回 -1
	mostRegionID := func(size int) int {
		count := make([]int, size)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 0 {
					for k := 0; k < 4; k++ {
						x := i + dx[k]
						y := j + dy[k]
						if 0 <= x && x < m &&
							0 <= y && y < n &&
							grid[x][y] > 0 {
							count[grid[x][y]]++
						}
					}
				}
			}
		}
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
		id := mostRegionID(size)
		if id == -1 {
			break
		}
		res += installWalls(id)
		nightInfect()
	}

	return res
}
