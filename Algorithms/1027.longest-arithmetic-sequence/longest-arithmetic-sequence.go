package problem1027

func longestArithSeqLength(A []int) int {
	n := len(A)
	res := 0
	var dfs func(int, int, int, int)
	dfs = func(i, diff, next, count int) {
		if i == n {
			res = max(res, count)
			return
		}
		if A[i] == next {
			next += diff
			count++
		}
		i++
		dfs(i, diff, next, count)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := A[j] - A[i]
			dfs(j+1, diff, A[j]+diff, 2)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
