package problem0885

func spiralMatrixIII(R int, C int, r int, c int) [][]int {
	countDown := R * C

	res := make([][]int, 1, countDown)
	res[0] = []int{r, c}
	countDown--

	steps := 2
	x, y, tmp := 0, 1, 0

	for countDown > 0 {
		for s := steps / 2; s > 0; s-- {
			r += x
			c += y
			if 0 <= r && r < R && 0 <= c && c < C {
				res = append(res, []int{r, c})
				countDown--
			}
		}

		// 方向变化这一段，很巧妙，可以参考
		// https://leetcode.com/problems/spiral-matrix-iii/discuss/158970/C++JavaPython-112233-Steps

		tmp = x
		x = y
		y = -tmp
		steps++
	}

	return res
}
