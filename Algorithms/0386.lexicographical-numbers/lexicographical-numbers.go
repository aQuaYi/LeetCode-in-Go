package problem0386

func lexicalOrder(max int) []int {
	res := make([]int, 0, max)

	var dfs func(int)
	dfs = func(begin int) {
		limit := (begin + 10) / 10 * 10
		for begin <= max && begin < limit {
			res = append(res, begin)
			dfs(begin * 10)
			begin++
		}
	}

	dfs(1)
	return res
}
