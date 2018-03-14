package problem0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}
	cs2(candidates, solution, target, &res)

	return res
}

func cs2(candidates, solution []int, target int, result *[][]int) {
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

	// 去掉已使用了的candidates[0]
	cs2(candidates[1:], append(solution, candidates[0]), target-candidates[0], result)

	// 不使用candidates[0]的话，就要把所有和candidates[0]相等的元素都去掉。
	cs2(next(candidates), solution, target, result)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}
