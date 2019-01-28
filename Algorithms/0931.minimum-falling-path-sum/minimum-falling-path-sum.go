package problem0931

func minFallingPathSum(A [][]int) int {
	size := len(A)
	res := 1 << 20

	var falling func(int, int, int)
	falling = func(i, j, sum int) {
		if i == size {
			res = min(res, sum)
			return
		}

		sum += A[i][j]

		x, y := i+1, max(j-1, 0)

		for ; y < size && y <= j+1; y++ {
			falling(x, y, sum)
		}
	}

	for j := 0; j < size; j++ {
		falling(0, j, 0)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
