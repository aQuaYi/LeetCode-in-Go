package problem0885

func spiralMatrixIII(R, C, x, y int) [][]int {
	moves := R * C

	res := make([][]int, 1, moves)
	res[0] = []int{x, y}
	moves--

	round := 2
	dx, dy := 0, 1

	for moves > 0 {
		for m := round / 2; m > 0; m-- {
			x += dx
			y += dy
			if 0 <= x && x < R && 0 <= y && y < C {
				res = append(res, []int{x, y})
				moves--
			}
		}
		round++
		/**
		 * 从向西开始，(dx,dy) 的变化如下
		 * (0, 1)
		 * (1, 0)
		 * (0,-1)
		 * (-1,0)
		 * dx, dy = dy, -dx 可以实现以上变化
		 */
		dx, dy = dy, -dx
	}

	return res
}
