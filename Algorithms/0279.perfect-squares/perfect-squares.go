package problem0279

import (
	"math"
)

func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	// dp[i] 表示 the least number of perfect square numbers which sum to i
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				// 因为 i = ( i - p ) + p，p 是 平方数
				// 所以 dp[i] = dp[i-p] + 1
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
