package problem0864

var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	startX, startY := m, n
	start := 1 << 6
	mask := start
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r := grid[i][j]
			if r == '@' {
				startX, startY = i, j
			} else if r >= 'a' {
				mask |= 1 << uint(r-'a')
			}
		}
	}

	masks := make([]int, m*n)
	masks[startX*n+startY] = start

	queue := make([][3]int, 1, m*n)
	queue[0] = [3]int{start, startX, startY}
	steps := 0

	for len(queue) > 0 {
		steps++
		size := len(queue)
		for i := 0; i < size; i++ {
			qi := queue[i]
			mk, x0, y0 := qi[0], qi[1], qi[2]
			for j := 0; j < 4; j++ {
				x, y := x0+dx[j], y0+dy[j]
				if 0 <= x && x < m &&
					0 <= y && y < n &&
					mk != masks[x*n+y] {

					bt := grid[x][y]
					if bt == '#' ||
						'A' <= bt && bt <= 'F' &&
							mk&(1<<uint(bt-'A')) == 0 {
						continue
					}

					newmk := mk
					if 'a' <= bt {
						newmk = mk | 1<<uint(bt-'a')
					}

					if newmk == mask {
						return steps
					}

					masks[x*n+y] |= mk
					queue = append(queue, [3]int{newmk, x, y})
				}
			}
		}

		queue = queue[size:]
	}

	return -1
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}
