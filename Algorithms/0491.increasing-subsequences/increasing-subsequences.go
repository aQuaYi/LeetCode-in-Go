package Problem0491

func findSubsequences(nums []int) [][]int {
	size := len(nums)
	res := make([][]int, 0, (size-1)*(size-1))
	if size < 2 {
		return res
	}

	var dfs func([]int, int)
	dfs = func(ints []int, idx int) {
		if len(ints) > 1 {
			tmp := make([]int, len(ints))
			copy(tmp, ints)
			res = append(res, tmp)
		}

		t := nums[0] - 1
		for i := idx; i < size; i++ {
			if t == nums[i] {
				continue
			}
			dfs(append(ints, nums[i]), i+1)
			t = nums[i]
		}
	}

	dfs([]int{}, 0)
	return res
}
