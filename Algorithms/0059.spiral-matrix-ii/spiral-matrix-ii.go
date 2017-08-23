package Problem0059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	// 4 条边界，依照题意，沿着边界填写
	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1
	for top <= bottom && left <= right {
		// →
		for j := left; j <= right; j++ {
			res[top][j] = num
			num++
		}
		top++
		// ↓
		for j := top; j <= bottom; j++ {
			res[j][right] = num
			num++
		}
		right--
		// ←
		for j := right; j >= left; j-- {
			res[bottom][j] = num
			num++
		}
		bottom--
		// ↑
		for j := bottom; j >= top; j-- {
			res[j][left] = num
			num++
		}
		left++
	}

	return res
}
