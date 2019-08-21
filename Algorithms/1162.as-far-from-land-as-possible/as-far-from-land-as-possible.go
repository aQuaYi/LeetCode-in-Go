package problem1162

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func maxDistance(grid [][]int) int {
	n := len(grid)
	hasSeen := [128][128]bool{}

	queue := make([][2]int, 0, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
				hasSeen[i][j] = true
			}
		}
	}

	if len(queue) == n*n {
		return -1
	}

	dist := -1
	for len(queue) > 0 {
		dist++
		size := len(queue)
		for s := 0; s < size; s++ {
			g := queue[s]
			x, y := g[0], g[1]
			for k := 0; k < 4; k++ {
				i, j := x+dx[k], y+dy[k]
				if i < 0 || n <= i ||
					j < 0 || n <= j ||
					hasSeen[i][j] {
					continue
				}
				queue = append(queue, [2]int{i, j})
				hasSeen[i][j] = true
			}
		}
		queue = queue[size:]
	}

	return dist
}
