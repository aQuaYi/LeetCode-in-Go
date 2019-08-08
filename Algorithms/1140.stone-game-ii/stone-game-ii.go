package problem1140

func stoneGameII(A []int) int {
	n := len(A)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + A[i]
	}

	var cur func(i, M int) int
	cur = func(i, M int) int {
		if i >= n {
			return 0
		}
		res := 0
		for x := 1; x <= 2*M; x++ {
			res = max(
				res,
				sum[n]-sum[i]-cur(i+x, max(M, x)),
			)
		}
		return res
	}

	return cur(0, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
