package problem0132

func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i] == minCut(s[:i])
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = i - 1
	}

	for i := 0; i < n+1; i++ {
		// With char at i as the center to check palindrome and set the new min cut number for
		// the up-boundary of the palindrome to be minimum of its current number and the min of
		// the low-boundary plus 1.
		// Need to check char at i to be the center for odd and event length of palindrome.
		for j := 0; 0 <= i-j && i+j < n && s[i-j] == s[i+j]; j++ {
			dp[i+j+1] = min(dp[i-j]+1, dp[i+j+1])
		}

		for j := 1; 0 <= i-j+1 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			dp[i+j+1] = min(dp[i-j+1]+1, dp[i+j+1])
		}
	}

	return dp[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
