package problem1092

func shortestCommonSupersequence(A string, B string) string {
	m, n := len(A), len(B)

	dp := [1001][1001]int{}

	if A[0] == B[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 2
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := i + j + 2
			if i > 0 && cur > dp[i-1][j]+1 {
				cur = dp[i-1][j] + 1
			}
			if j > 0 && cur > dp[i][j-1]+1 {
				cur = dp[i][j-1] + 1
			}

			if A[i] == B[j] {
				if i > 0 && j > 0 && cur > dp[i-1][j-1]+1 {
					cur = dp[i-1][j-1] + 1
				} else if i == 0 {
					cur = 1 + j
				} else if j == 0 {
					cur = 1 + i
				}
			}

			dp[i][j] = cur
		}
	}

	m, n = m-1, n-1
	ans := make([]byte, 0, dp[m][n])
	for m >= 0 && n >= 0 {
		if m > 0 && dp[m][n] == dp[m-1][n]+1 {
			ans = append(ans, A[m])
			m--
		} else if n > 0 && dp[m][n] == dp[m][n-1]+1 {
			ans = append(ans, B[n])
			n--
		} else if m > 0 && n > 0 && dp[m][n] == dp[m-1][n-1]+1 {
			ans = append(ans, A[m])
			m--
			n--
		} else if m == 0 && n+2 == dp[m][n] {
			ans = append(ans, A[m])
			m--
		} else if n == 0 && m+2 == dp[m][n] {
			ans = append(ans, B[n])
			n--
		} else if m+n+1 == dp[m][n] {
			ans = append(ans, A[m])
			m--
			n--
		}
	}
	for ; m >= 0; m-- {
		ans = append(ans, A[m])
	}
	for ; n >= 0; n-- {
		ans = append(ans, B[n])
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
