package Problem0279

func numSquares(n int) int {
	perfects := make([]int, intSqrt(n))
	for i := 0; (i+1)*(i+1) <= n; i++ {
		perfects[i] = (i + 1) * (i + 1)
	}

	// dp[i] 表示 the least number of perfect square numbers which sum to i
	var dp = make([]int, n+1)
	maxInt := 1<<63 - 1

	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		dp[i] = maxInt
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i]-1 > dp[i-p] {
				// 因为 i = ( i - p ) + p，p 是 平方数
				// 所以 dp[i] = dp[i-p] + 1
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}

// 返回 x 的平方根的整数部分
// 这个函数比 int(math.Sqrt(float64(x))) 快的多
// 详见 benchmark 的结果
func intSqrt(x int) int {
	res := x

	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}
