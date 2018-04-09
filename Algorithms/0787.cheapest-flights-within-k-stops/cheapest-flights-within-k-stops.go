package problem0787

import "fmt"

const maxPrice = 1<<63 - 1

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	nextGroup := make([][][]int, n)
	for i := range flights {
		nextGroup[flights[i][0]] = append(nextGroup[flights[i][0]], flights[i][1:])
	}

	isChecked := make([]bool, n)
	res := maxPrice
	path := make([]int, k+5)
	path[0] = src
	dfs(src, dst, k, 0, &res, isChecked, nextGroup, path)

	if res == maxPrice {
		return -1
	}
	return res
}

func dfs(src, dst, k, spent int, res *int, isChecked []bool, nextGroup [][][]int, path []int) {
	if src == dst {
		*res = min(*res, spent)

		// TODO: 删除此处输出}
		fmt.Println(path)

		return
	}

	if k < 0 {
		return
	}

	isChecked[src] = true

	for _, next := range nextGroup[src] {
		path[len(path)-k-4] = next[0]
		dfs(next[0], dst, k-1, spent+next[1], res, isChecked, nextGroup, path)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
