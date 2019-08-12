package problem1143

func longestCommonSubsequence(s1, s2 string) int {
	m, n := len(s1), len(s2)
	A, B := []byte(s1), []byte(s2)

	cur := make([]int, m+1)
	prev := make([]int, m+1)

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			rec := max(prev[i+1], cur[i])
			if A[i] == B[j] {
				cur[i+1] = max(rec, prev[i]+1)
			} else {
				cur[i+1] = rec
			}
		}
		cur, prev = prev, cur
	}
	return prev[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
