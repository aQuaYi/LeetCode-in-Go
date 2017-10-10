package Problem0300

func lengthOfLIS(nums []int) int {
	size := len(nums)

	max := 0
	var dfs func(int, int, int)
	dfs = func(idx, n, count int) {
		if max < count {
			max = count
		}

		for i := idx + 1; i < size; i++ {
			// i < size-max+count 用于提前结束
			if nums[i] > n && i < size-max+count {
				dfs(i, nums[i], count+1)
			}
		}
	}

	dfs(-1, -1<<63, 0)

	return max
}
