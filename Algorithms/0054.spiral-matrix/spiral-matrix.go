package problem0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])

	next := nextFunc(m, n)

	res := make([]int, m*n)
	for i := range res {
		x, y := next()
		res[i] = matrix[x][y]
	}

	return res
}

func nextFunc(m, n int) func() (int, int) {
	top, down := 0, m-1
	left, right := 0, n-1
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch { // 如果撞墙了，需要修改 dx, dy 和相应的边界值
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > down:
			right--
			dx, dy = 0, -1
		case y+dy < left:
			down--
			dx, dy = -1, 0
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		return x, y
	}
}
