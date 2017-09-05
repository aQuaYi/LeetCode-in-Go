package Problem0070

func climbStairs(n int) int {
	res := 0

	var dfs func(n int)
	dfs = func(n int) {
		switch {
		case n < 0:
			return
		case n == 0:
			res++
			return
		default:
			dfs(n - 1)
			dfs(n - 2)
		}
	}

	dfs(n)

	return res
}
