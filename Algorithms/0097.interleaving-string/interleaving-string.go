package Problem0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([]bool, m+n+1)
	dp[0] = true
	i, j, k := 0, 0, 0
	for i <= m && j <= n && k < m+n+1 {
		switch {
		case i < m && s1[i] == s3[k]:
			dp[k] = true
			i++
			k++
		case j < n && s2[j] == s3[k]:
			dp[k] = true
			j++
			k++
		default:
			return false
		}
		if i == m && j == n {
			return true
		}
	}
	return false
}
