package problem0874

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func robotSim(commands []int, obstacles [][]int) int {
	isBlocked := make(map[int]bool, 10000)
	for _, o := range obstacles {
		x, y := o[0], o[1]
		isBlocked[encode(x, y)] = true
	}
	i, j, res := 0, 0, 0
	index := 100000

	for _, c := range commands {
		switch {
		case c == -2:
			index--
		case c == -1:
			index++
		case 1 <= c && c <= 9:
			idx := index % 4
			x := dx[idx]
			y := dy[idx]
			for c > 0 && !isBlocked[encode(i+x, j+y)] {
				c--
				i += x
				j += y
			}
			res = max(res, i*i+j*j)
		}
	}
	return res
}

func encode(x, y int) int {
	x += 30000
	y += 30000
	return x<<16 | y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
