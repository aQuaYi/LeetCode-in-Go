package problem0873

func lenLongestFibSubseq(A []int) int {
	size := len(A)

	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	res := 0

	index := make(map[int]int, size)
	for k := 0; k < size; k++ {
		index[A[k]] = k
		for j := 0; j < k; j++ {
			dp[j][k] = 2
			i, ok := index[A[k]-A[j]]
			if ok && A[k]-A[j] < A[j] {
				dp[j][k] = dp[i][j] + 1
			}
			res = max(res, dp[j][k])
		}
	}

	if res <= 2 {
		return 0
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
