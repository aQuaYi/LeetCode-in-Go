package problem0787

const maxPrice = 1<<63 - 1

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	nextGroup := make([][][]int, n)
	for i := range flights {
		src := flights[i][0]
		nextGroup[src] = append(nextGroup[src], flights[i][1:])
	}

	isChecked := make([]bool, n)
	res := maxPrice
	dfs(src, dst, k, 0, &res, isChecked, nextGroup)

	if res == maxPrice {
		return -1
	}
	return res
}

func dfs(src, dst, k, spent int, res *int, isChecked []bool, nextGroup [][][]int) {
	if src == dst {
		*res = min(*res, spent)
		return
	}

	if k < 0 || isChecked[src] {
		return
	}

	isChecked[src] = true
	for _, next := range nextGroup[src] {
		dfs(next[0], dst, k-1, spent+next[1], res, isChecked, nextGroup)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
