package problem1140

func stoneGameII(A []int) int {
	n := len(A)

	for i := n - 2; i >= 0; i-- {
		A[i] += A[i+1]
	}

	// new A[i] = sum of old A[i:]

	mem := [101][33]int{}

	var dp func(int, int) int
	dp = func(i, m int) int {
		if i+2*m >= n {
			return A[i]
		}
		if mem[i][m] > 0 {
			return mem[i][m]
		}
		res := 0
		for x := 1; x <= 2*m; x++ {
			res = max(
				res,
				A[i]-dp(i+x, max(m, x)),
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
