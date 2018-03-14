package problem0768

import (
	"sort"
)

func maxChunksToSorted(arr []int) int {
	return helper(convert(arr))
}

// 把
// 	arr will have length in range [1, 2000].
// 	arr[i] will be an integer in range [0, 10**8]
// 转换成
// 	arr will have length in range [1, 10].
// 	arr[i] will be a permutation of [0, 1, ..., arr.length - 1]
// 就可以利用 769 题的答案来解答了
func convert(arr []int) []int {
	a := make([][]int, len(arr))
	for i := range a {
		a[i] = []int{arr[i], i}
	}

	sort.Slice(a, func(i int, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})

	res := make([]int, len(arr))

	for i := range a {
		res[a[i][1]] = i
	}

	return res
}

//  直接复制了 769 的答案
func helper(arr []int) int {
	lastIdx, res := 0, 0
	for i := 0; i < len(arr); i++ {
		if lastIdx < arr[i] {
			lastIdx = arr[i]
			continue
		}

		if i == lastIdx {
			res++
			lastIdx++
		}

	}

	return res
}
