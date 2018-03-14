package problem0039

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}
	cs(candidates, solution, target, &res)

	return res
}

func cs(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		// target < candidates[0] 因为candidates是排序好的
		return
	}

	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// 可以注释掉以下语句，运行单元测试，查看错误发生。
	solution = solution[:len(solution):len(solution)]

	cs(candidates, append(solution, candidates[0]), target-candidates[0], result)

	cs(candidates[1:], solution, target, result)
}
