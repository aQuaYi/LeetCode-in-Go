package problem0035

func searchInsert(nums []int, target int) int {
	// 没有把i放入for语句中
	// 是为了兼容，len(nums) == 0 和 target > nums[len(nums)-1]两种情况
	i := 0

	for i < len(nums) && nums[i] <= target {
		// 相等的时候，直接返回
		if nums[i] == target {
			return i
		}

		// 否则，就去检查下一个
		i++
	}

	return i
}
