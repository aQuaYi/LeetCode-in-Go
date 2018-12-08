package problem0903

const mod = 1e9 + 7

// ref: https://leetcode.com/problems/valid-permutations-for-di-sequence/discuss/168278/C++JavaPython-DP-Solution-O(N2)
func numPermsDISequence(S string) int {
	n := len(S)

	dp := [201][201]int{}

	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}

	for i := 0; i < n; i++ {
		count := 0
		if S[i] == 'I' {
			for j := 0; j < n-i; j++ {
				count += dp[i][j]
				dp[i+1][j] = count % mod

			}
		} else {
			for j := n - i - 1; j >= 0; j-- {
				count += dp[i][j+1]
				dp[i+1][j] = count % mod
			}
		}
	}

	return dp[n][0]
}
