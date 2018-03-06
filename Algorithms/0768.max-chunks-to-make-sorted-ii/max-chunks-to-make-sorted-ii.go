package Problem0768

import (
	"sort"
)

func maxChunksToSorted(arr []int) int {
	return helper(convert(arr))
}

// 把 768 题的参数，转换成 769 题的参数
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
	// 充分利用好 arr[i] will be a permutation of [0, 1, ..., arr.length - 1]
	// 假设某个 chunk 为 arr[j:k]，则对于
	// j <=     i  < k ，必有
	// j <= arr[i] < k
	// 因为排序后，arr[x]=x
	//
	// lastIdx 是 下一个要切下来的 chunk 中最后一个元素的索引号
	lastIdx, res := 0, 0
	for i := 0; i < len(arr); i++ {
		if lastIdx < arr[i] {
			lastIdx = arr[i]
			// lastIdx == arr[i]，才能满足前面提到的要求
			continue
		}

		if i == lastIdx {
			// 此时可以切下一刀
			res++
			lastIdx++
		}

	}

	return res
}
