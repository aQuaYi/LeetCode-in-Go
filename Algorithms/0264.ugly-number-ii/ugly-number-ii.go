package problem0264

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	pos := []int{0, 0, 0}
	factors := []int{2, 3, 5}
	candidates := []int{2, 3, 5}

	res := make([]int, n)

	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = min(candidates)
		for j := 0; j < 3; j++ {
			if res[i] == candidates[j] {
				pos[j]++
				candidates[j] = res[pos[j]] * factors[j]
			}
		}
	}

	return res[n-1]
}

func min(candidates []int) int {
	return lesser(
		candidates[0],
		lesser(
			candidates[1],
			candidates[2],
		),
	)
}

func lesser(a, b int) int {
	if a < b {
		return a
	}
	return b
}
