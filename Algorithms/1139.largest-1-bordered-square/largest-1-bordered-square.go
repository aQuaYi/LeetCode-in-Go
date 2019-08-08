package problem1139

func largest1BorderedSquare(A [][]int) int {
	m, n := len(A), len(A[0])

	isBorderOk := func(x, y, dx, dy, w int) bool {
		xMax, yMax := x+w, y+w
		for x < xMax && y < yMax {
			if A[x][y] == 0 {
				return false
			}
			x += dx
			y += dy
		}
		return true
	}

	isOk := func(x, y, w int) bool {
		return isBorderOk(x, y, 0, 1, w) &&
			isBorderOk(x, y, 1, 0, w) &&
			isBorderOk(x+w-1, y, 0, 1, w) &&
			isBorderOk(x, y+w-1, 1, 0, w)
	}

	res := 0
	for i := 0; i+res <= m; i++ {
		for j := 0; j+res <= n; j++ {
			w := min(m-i, n-j)
			for w > res {
				if isOk(i, j, w) {
					res = w
					break
				}
				w--
			}
		}
	}

	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
