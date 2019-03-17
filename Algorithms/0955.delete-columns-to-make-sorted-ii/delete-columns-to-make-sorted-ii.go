package problem0955

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	// isBigger[i]=true means A[i]>A[i-1]
	isBigger := make([]bool, m)

	res := 0

	for j := 0; j < n; j++ {
		t := make([]bool, m)
		i := 1
		count := 1
		for ; i < m; i++ {
			if isBigger[i] ||
				A[i-1][j] < A[i][j] {
				t[i] = true
				count++
			} else if A[i-1][j] > A[i][j] {
				break
			}
		}
		if i == m {
			if count == m {
				return res
			}
			isBigger = t
		} else {
			res++
		}
	}
	return n
}
