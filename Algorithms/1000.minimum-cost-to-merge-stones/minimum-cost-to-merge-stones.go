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

	for width := K; width <= size; width++ {
		for l := 0; l+width <= size; l++ {
			r := l + width - 1
			dp[l][r] = 1 << 32
			for m := l; m < r; m += K - 1 {
				dp[l][r] = min(dp[l][r], dp[l][m]+dp[m+1][r])
			}
			if (r-l)%(K-1) == 0 { // 如果 stones[l:r+1] 可以合并成一块石头
				dp[l][r] += sum[r+1] - sum[l] // 还要加上最后一次合并的费用
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
