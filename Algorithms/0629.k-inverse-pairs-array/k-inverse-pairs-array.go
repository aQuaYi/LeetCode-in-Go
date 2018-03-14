package problem0629

const m = 1000000007

func kInversePairs(n int, k int) int {
	// dp[n][k] 表示 1～n 可以表示的有 k 个逆序对的个数
	// 想要得到 dp[n][k]，我们可以
	//     把 n 放在 dp[n-1][k]   的 n-1 位上
	//     把 n 放在 dp[n-1][k-1] 的 n-2 位上
	//     把 n 放在 dp[n-1][k-2] 的 n-3 位上
	//     .....
	//     把 n 放在 dp[n-1][k-(n-1)] 的 0 位上
	// 可得
	// dp[n][k] = dp[n-1][k]+dp[n-1][k-1]+dp[n-1][k-2]+...+dp[n-1][k+1-n+1]+dp[n-1][k-n+1]
	// 还可得
	// dp[n][k+1] = dp[n-1][k+1]+dp[n-1][k]+dp[n-1][k-1]+dp[n-1][k-2]+...+dp[n-1][k+1-n+1]
	// 以上两式相减，可得
	// dp[n][k+1] = dp[n][k]+dp[n-1][k+1]-dp[n-1][k+1-n]
	// 最终得到，
	// dp[n][k] = dp[n][k-1]+dp[n-1][k]-dp[n-1][k-n]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		// 1~i 最多只有 i*(i-1)/2 个逆序对
		maxJ := min(k, i*(i-1)/2)
		for j := 1; j <= maxJ; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % m
			if j >= i {
				dp[i][j] -= dp[i-1][j-i]
				if dp[i][j] < 0 {
					dp[i][j] += m
				}
			}
		}
	}

	return dp[n][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
