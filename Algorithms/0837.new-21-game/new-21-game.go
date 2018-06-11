package problem0837

func new21Game(N int, K int, W int) float64 {
	if K == 0 ||
		N >= K+W {
		return 1.0
	}

	// dp[i]: 获取分数 i 的概率
	// 假设 p 为抽取 [1,W] 中任意一个数的概率，即 p = 1/W
	//
	// dp[i] = dp[i-1]*p + dp[i-2]*p + ... + dp[i-W]*p
	// 		 = (dp[i-1] + dp[i-2] + ... + dp[i-W]) * p
	// 		 = (dp[i-1] + dp[i-2] + ... + dp[i-W]) / W
	// 令 last= dp[i-1] + dp[i-2] + ... + dp[i-W]
	// dp[i] = last/ W
	//
	// 由于 i 的取值范围为 [1,N]
	// 当 W	<= i < K 时
	// 		last= dp[i-1] + dp[i-2] + ... + dp[i-W]
	// 当 i < W 时
	// 		last应为 dp[i-1] + dp[i-2] + ... + dp[0]
	// 当 K <= i 时
	// 		last应为 dp[k-1] + dp[k-2] + ... + dp[i-W]
	dp := make([]float64, N+1)
	dp[0] = 1.0
	last := 1.0 // 初始概率为 1
	res := 0.0
	for i := 1; i <= N; i++ {
		dp[i] = last / float64(W)

		if W <= i {
			last -= dp[i-W]
		}

		if i < K {
			last += dp[i]
		} else {
			res += dp[i]
		}

	}

	return res
}
