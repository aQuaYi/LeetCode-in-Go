package problem1140

func stoneGameII(A []int) int {
	n := len(A)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + A[i]
	}

	mem := [101][65]int{}

	var dp func(int, int) int
	dp = func(i, m int) int {
		if i >= n {
			return 0
		}
		if mem[i][m] > 0 {
			return mem[i][m]
		}
		res := 0
		for x := 1; x <= 2*m; x++ {
			res = max(
				res,
				sum[n]-sum[i]-dp(i+x, max(m, x)),
			)
		}
		mem[i][m] = res
		return res
	}

	return dp(0, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
