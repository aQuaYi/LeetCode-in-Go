package problem1034

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func colorBorder(grid [][]int, x0 int, y0 int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	c := grid[x0][y0]

	isBorder := func(x, y int) bool {
		return x == 0 || x == m-1 ||
			y == 0 || y == n-1 ||
			grid[x][y-1] != c ||
			grid[x][y+1] != c ||
			grid[x-1][y] != c ||
			grid[x+1][y] != c
	}

	connected := make([][2]int, 1, 1024)
	hasSeen := [51][51]bool{}

	connected[0] = [2]int{x0, y0}
	hasSeen[x0][y0] = true

	borders := make([][2]int, 0, 1024)
	for len(connected) > 0 {
		size := len(connected)
		for i := 0; i < size; i++ {
			p := connected[i]
			if isBorder(p[0], p[1]) {
				borders = append(borders, p)
			}
			for k := 0; k < 4; k++ {
				x, y := p[0]+dx[k], p[1]+dy[k]
				if 0 <= x && x < m &&
					0 <= y && y < n &&
					grid[x][y] == c &&
					!hasSeen[x][y] {
					connected = append(connected, [2]int{x, y})
					hasSeen[x][y] = true
				}
			}
		}
		connected = connected[size:]
	}

	for _, b := range borders {
		grid[b[0]][b[1]] = color
	}

	return grid
}
