package problem0885

var xs = []int{0, 1, 0, -1}
var ys = []int{1, 0, -1, 0}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	countDown := R * C

	res := make([][]int, 1, countDown)
	res[0] = []int{r0, c0}
	countDown--

	x, y := r0, c0

	steps := 2
	direct := 0

	for countDown > 0 {
		dx, dy := xs[direct], ys[direct]
		s := steps / 2
		for s > 0 {
			s--
			x += dx
			y += dy
			if 0 <= x && x < R && 0 <= y && y < C {
				res = append(res, []int{x, y})
				countDown--
			}
		}
		direct = (direct + 1) % 4
		steps++
	}

	return res
}
