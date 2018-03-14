package problem0526

func countArrangement(N int) int {
	count := 0
	a := make([]int, N+1)
	for i := 0; i <= N; i++ {
		a[i] = i
	}

	var dfs func(int)
	dfs = func(idx int) {
		if idx == 0 {
			count++
			return
		}

		for i := idx; i > 0; i-- {
			a[idx], a[i] = a[i], a[idx]
			if isBeautiful(a[idx], idx) {
				dfs(idx - 1)
			}
			a[idx], a[i] = a[i], a[idx]
		}
	}

	dfs(N)
	return count
}

func isBeautiful(i, j int) bool {
	return i%j == 0 || j%i == 0
}
