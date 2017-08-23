package Problem0059

func generateMatrix(n int) [][]int {
	total := n * n
	num := 1
	res := [][]int{}

	// 准备矩阵
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}

	// i 代表了从外往里的第 i 个圈
	for i := 0; num <= total; i++ {
		// →
		for col := i; col < n-i; col++ {
			res[i][col] = num
			num++
		}
		// ↓
		for row := i + 1; row < n-i; row++ {
			res[row][n-i-1] = num
			num++
		}
		// ←
		for col := n - i - 2; col >= i; col-- {
			res[n-i-1][col] = num
			num++
		}
		// ↑
		for row := n - i - 2; row >= i+1; row-- {
			res[row][i] = num
			num++
		}
	}

	return res
}
