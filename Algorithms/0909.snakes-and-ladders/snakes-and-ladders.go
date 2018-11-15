package problem0909

func snakesAndLadders(p [][]int) int {
	m, n := len(p), len(p[0])
	mn := m * n
	dp := [401]int{}
	for i := 2; i <= mn; i++ {
		dp[i] = mn
	}
	isJumped := [401]bool{}

	for i := 2; i <= mn; i++ {
		x, y := location(m, n, i)
		pxy := p[x][y]
		if -1 < pxy && pxy < i {
			continue
		}
		for j := i - 1; j >= i-6 && j >= 0; j-- {
			dp[i] = min(dp[i], dp[j]+1)
		}
		if pxy > i && !isJumped[i] {
			dp[pxy] = min(dp[pxy], dp[i])
			isJumped[pxy] = true
		}
	}
	if dp[mn] == mn {
		return -1
	}
	return dp[mn]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func location(m, n, square int) (x, y int) {
	square--
	// 首先计算正常的位置
	x, y = square/n, square%n
	if x%2 == 1 {
		// 行号为奇数的行，需要左右翻转
		y = n - 1 - y
	}
	// 整体上下翻转
	x = m - 1 - x
	return
}
