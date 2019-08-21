package problem1162

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func maxDistance(grid [][]int) int {
	n := len(grid)

	lands := make([][2]int, 0, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				lands = append(lands, [2]int{i, j})
			}
		}
	}

	if len(lands) == n*n {
		return -1
	}

	dist := -1
	// BFS
	for len(lands) > 0 {
		dist++
		size := len(lands)
		for s := 0; s < size; s++ {
			g := lands[s]
			x, y := g[0], g[1]
			for k := 0; k < 4; k++ {
				i, j := x+dx[k], y+dy[k]
				if i < 0 || n <= i ||
					j < 0 || n <= j ||
					grid[i][j] == 1 {
					continue
				}
				lands = append(lands, [2]int{i, j})
				grid[i][j] = 1
			}
		}
		lands = lands[size:]
	}

	return dist
}
