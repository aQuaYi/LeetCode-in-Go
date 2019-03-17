package problem0955

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])

	for j := 0; j < n; j++ {
		i := 1
		for ; i < m; i++ {
			if A[i-1][j] > A[i][j] {
				break
			}
		}
		if i == m {
			return j
		}
	}

	return n
}
