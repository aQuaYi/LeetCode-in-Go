package problem0960

// ref: https://leetcode.com/problems/delete-columns-to-make-sorted-iii/discuss/205679/C%2B%2BJavaPython-Maximum-Increasing-Subsequence

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	res := n - 1

	dp := make([]int, n)
	// dp[i] means the longest subsequence ends with i-th element.
	// For all i < j, if A[][i] <= A[][j], then dp[j] = max(dp[j], dp[i] + 1)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			k := 0
			for ; k < m; k++ {
				if A[k][i] > A[k][j] {
					break
				}
			}
			if k == m && // compared all strings
				dp[j] < dp[i]+1 { // find longer sub string
				dp[j] = dp[i] + 1
			}
		}
		res = min(res, n-dp[j])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
