package problem0765

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
		// 新产生的 pair 有可能会打乱 pairs 的顺序
		// 重新排序是为了保证每次的 pairs[0][0] 和 pairs[1][0] 都是一对
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i][0] < pairs[j][0]
		})

		for len(pairs) > 1 && isCouple(pairs[0][1], pairs[1][1]) {
			pairs = pairs[2:]
			// 此时相当于完成了一次交换。
			// 但是促成了两个 couple
			res++
			// 而且由于没有产生新的 pair，
			// 所以，pairs 的顺序没有被打乱，可以重复使用。
		}

		if len(pairs) == 0 {
			break
		}

		// 因为 pairs[0][0] 和 pairs[1][0] 肯定就是一个 couple
		// 要把 pairs[0][1] 和 pairs[1][1] 重新凑成一个 pair
		pairs[1] = makePair(pairs[0][1], pairs[1][1])
		pairs = pairs[1:]
		// 此时相当于完成了一次交换。
		// 只促成了一个 couple
		res++
	}

	return res
}

// 把较小的数放在前面
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
