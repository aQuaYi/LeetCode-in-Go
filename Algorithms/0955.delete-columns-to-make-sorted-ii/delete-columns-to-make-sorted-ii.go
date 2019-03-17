package problem0955

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])

	// isBigger[i]=true means A[i]>A[i-1]
	isBigger := make([]bool, m)
	res := 0

	for j := 0; j < n; j++ {
		t := make([]bool, m)
		i, count := 1, 1
		for ; i < m; i++ {
			if isBigger[i] ||
				A[i][j] > A[i-1][j] {
				t[i] = true
				count++
			} else if A[i-1][j] > A[i][j] {
				res++
				break
			}
		}
		if count == m { // for now, all is sorted.
			break
		}
		if i == m { // every string is checked
			isBigger = t
		}
	}

	return res
}
