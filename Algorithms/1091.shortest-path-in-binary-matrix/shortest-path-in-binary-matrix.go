package problem1091

var dx = []int{-1, -1, -1, 0, 1, 1, 1, 0}
var dy = []int{-1, 0, 1, 1, 1, 0, -1, -1}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	N := len(grid)

	cells := make([][2]int, 1, N*N)
	grid[0][0] = 1
	length := 1

	for len(cells) > 0 {
		size := len(cells)
		for s := 0; s < size; s++ {
			c := cells[s]
			x, y := c[0], c[1]
			if x == N-1 && y == N-1 {
				return length
			}
			for k := 0; k < 8; k++ {
				i, j := x+dx[k], y+dy[k]
				if i < 0 || N <= i ||
					j < 0 || N <= j ||
					grid[i][j] == 1 {
					continue
				}
				cells = append(cells, [2]int{i, j})
				grid[i][j] = 1
			}
		}
		length++
		cells = cells[size:]
	}

	return -1
}
