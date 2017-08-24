package Problem0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var cur func(int, []int, *[][]int)

	cur = func(idx int, temp []int, res *[][]int) {
		*res = append(*res, temp)

		if idx == len(nums) {
			return
		}

		cur(idx+1, append(temp[:len(temp):len(temp)], nums[idx]), res)

		for i := idx + 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				cur(i+1, append(temp[:len(temp):len(temp)], nums[i]), res)
			}
		}
	}

	cur(0, []int{}, &res)

	return res
}
