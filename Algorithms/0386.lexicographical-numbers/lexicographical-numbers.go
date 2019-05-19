package problem0386

func lexicalOrder(max int) []int {
	res := make([]int, 0, max)

	var dfs func(int)
	dfs = func(x int) {
		limit := (x + 10) / 10 * 10
		for x <= max && x < limit {
			res = append(res, x)
			if x*10 <= max {
				dfs(x * 10)
			}
			x++
		}
	}

	dfs(1)
	return res
}
