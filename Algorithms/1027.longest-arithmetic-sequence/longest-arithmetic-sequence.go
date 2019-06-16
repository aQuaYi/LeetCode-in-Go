package problem1027

// ref: https://leetcode.com/problems/longest-arithmetic-sequence/discuss/274611/JavaC%2B%2BPython-DP
func longestArithSeqLength(A []int) int {
	n := len(A)
	res := 2
	dp := [2001][20001]int{}
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			d := A[j] - A[i] + 10000 // delta
			dp[j][d] = max(1, dp[i][d]) + 1
			res = max(res, dp[j][d])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
