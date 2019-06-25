package problem1043

// ref: https://leetcode.com/problems/partition-array-for-maximum-sum/discuss/290863/JavaC%2B%2BPython-DP
func maxSumAfterPartitioning(A []int, K int) int {
	n := len(A)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		curMax := 0
		for k := 1; k <= K && 0 <= i-k+1; k++ {
			curMax = max(curMax, A[i-k+1])
			cur := curMax * k
			if i >= k {
				cur += dp[i-k]
			}
			dp[i] = max(dp[i], cur)
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
