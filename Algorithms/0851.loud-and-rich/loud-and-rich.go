package problem0851

func loudAndRich(richer [][]int, quiet []int) []int {
	size := len(quiet)

	// rs[y] 中保存了所有比 y 有钱的人
	rs := make([][]int, size)
	for _, r := range richer {
		x, y := r[0], r[1]
		rs[y] = append(rs[y], x)
	}

	res := make([]int, size)
	// res 全部设置为 -1
	// 作为是否设置过的标记
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
