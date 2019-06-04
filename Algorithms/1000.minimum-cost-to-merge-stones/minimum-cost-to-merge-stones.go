package problem1000

// ref: https://leetcode.com/problems/minimum-cost-to-merge-stones/discuss/247567/JavaC%2B%2BPython-DP
func mergeStones(stones []int, K int) int {
	size := len(stones)
	if (size-1)%(K-1) != 0 {
		return -1
	}

	sum := [31]int{}
	for i := 0; i < size; i++ {
		sum[i+1] = sum[i] + stones[i]
	}

	dp := [31][31]int{}
	// dp[l][r] 的含义是，
	//   stones[l:r+1] 完成**所有**合并后，所花费的最小费用。
	// 我着重强调了 **所有** 二字
	// dp[l][r] 在完成所有合并后，有两种可能
	//    1.     合并成一块石头
	//    2. 没有合并成一块石头

	for width := K; width <= size; width++ {
		for l := 0; l+width <= size; l++ {
			r := l + width - 1
			dp[l][r] = 1 << 32
			for m := l; m < r; m += K - 1 {
				dp[l][r] = min(dp[l][r], dp[l][m]+dp[m+1][r])
				// 把 dp[l][r] 分解成两个部分
				//     dp[l][m]   代表了，能合并成一块石头的左边，
				//     dp[m+1][r] 代表了，剩下的右边。
			}
			if (r-l)%(K-1) == 0 {
				// 如果 dp[l][r] 的左边和右边还能够进行合并
				// 还要加上最后一次合并的费用
				dp[l][r] += sum[r+1] - sum[l]
			}
		}
	}

	return dp[0][size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
