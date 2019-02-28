package problem0944

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	res := 0
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			if A[i-1][j] > A[i][j] {
				res++
				break
			}
		}
	}
	return res
}
