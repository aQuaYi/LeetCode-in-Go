package Problem0765

import (
	"sort"
)

func minSwapsCouples(row []int) int {
	n := len(row)

	// pairs 把不成对的数收集起来
	pairs := make([][]int, 0, 30)
	for i := 0; i < n; i += 2 {
		if !isCouple(row[i], row[i+1]) {
			pairs = append(pairs, makePair(row[i], row[i+1]))
		}
	}

	res := 0

	for len(pairs) > 0 {
		// 每次都要重新排序是为了保证每次的 pairs[0][0] 和 pairs[1][0] 都是一对
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i][0] < pairs[j][0]
		})

		// 因为 pairs[0][0] 和 pairs[1][0] 是一对
		// 把 pairs[0][1] 和 pairs[1][1] 重新凑成一个 pair
		pairs[1] = makePair(pairs[0][1], pairs[1][1])
		// pairs 可以删除 pairs[0]
		pairs = pairs[1:]
		res++

		if isCouple(pairs[0][0], pairs[0][1]) {
			pairs = pairs[1:]
		}

	}

	return res
}

func makePair(a, b int) []int {
	if a > b {
		a, b = b, a
	}
	return []int{a, b}
}

func isCouple(a, b int) bool {
	if a > b {
		a, b = b, a
	}
	return a+1 == b && a%2 == 0
}
