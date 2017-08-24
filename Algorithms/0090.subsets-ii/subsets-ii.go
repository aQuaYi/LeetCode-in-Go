package Problem0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var dfs func(int, []int, *[][]int)

	dfs = func(idx int, temp []int, res *[][]int) {
		t := make([]int, len(temp))
		copy(t, temp)
		// 没有以上两行，答案就是错的
		// 因为temp的底层数组在递归过程中，不停地修改
		// 程序结束时，temp的底层数组的值，全部是 nums 的最大值。
		*res = append(*res, t)

		for i := idx; i < len(nums); i++ {
			if i == idx || nums[i] != nums[i-1] {
				dfs(i+1, append(temp, nums[i]), res)
			}
		}
	}

	temp := make([]int, 0, len(nums))
	dfs(0, temp, &res)

	return res
}
