package problem0851

func loudAndRich(richer [][]int, quiet []int) []int {
	size := len(quiet)

	rs := make([][]int, size)
	for i := range rs {
		rs[i] = make([]int, 0, size)
	}
	for _, r := range richer {
		x, y := r[0], r[1]
		rs[y] = append(rs[y], x)
	}

	res := make([]int, size)
	for i := range res {
		res[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if res[i] >= 0 {
			return res[i]
		}
		res[i] = i
		for _, j := range rs[i] {
			if quiet[res[i]] > quiet[dfs(j)] {
				res[i] = res[j]
			}
		}
		return res[i]
	}

	for i := 0; i < size; i++ {
		dfs(i)
	}

	return res
}
