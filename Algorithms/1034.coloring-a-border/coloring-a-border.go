package problem1034

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func colorBorder(grid [][]int, x0 int, y0 int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	c := grid[x0][y0]

	points := make([][2]int, 1, 1024)
	points[0] = [2]int{x0, y0}
	for len(points) > 0 {
		size := len(points)
		for i := 0; i < size; i++ {
			p := points[i]
			grid[p[0]][p[1]] = -c
			for i := 0; i < 4; i++ {
				x, y := p[0]+dx[i], p[1]+dy[i]
				if 0 <= x && x < m &&
					0 <= y && y < n &&
					grid[x][y] == c {
					points = append(points, [2]int{x, y})
				}
			}
		}
		points = points[size:]
	}

	isBorder := func(x, y int) bool {
		return x == 0 || x == m-1 ||
			y == 0 || y == n-1 ||
			grid[x][y-1] != -c ||
			grid[x][y+1] != -c ||
			grid[x+1][y] != -c ||
			grid[x-1][y] != -c
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				res[i][j] = grid[i][j]
				continue
			}
			if isBorder(i, j) {
				res[i][j] = color
			} else {
				res[i][j] = c
			}
		}
	}

	return res
}
