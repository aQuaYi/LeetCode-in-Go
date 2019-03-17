package problem0955

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	res := 0
	for j := 0; j < n; j++ {
		for i := 0; i < m-1; i++ {
			if A[i][j] > A[i+1][j] {
				res++
				break
			}
		}
	}
	return res
}
