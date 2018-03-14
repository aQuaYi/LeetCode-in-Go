package problem0070

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	rec := make([]int, n+1)
	rec[0], rec[1] = 1, 1

	for i := 2; i <= n; i++ {
		rec[i] = rec[i-1] + rec[i-2]
	}

	return rec[n]
}
