package Problem0456

func find132pattern(nums []int) bool {
	size := len(nums)
	if size < 3 {
		return false
	}

	Min := min(nums[0], nums[1])
	Max := max(nums[0], nums[1])

	for i := 2; i < size; i++ {
		if Min < nums[i] && nums[i] < Max {
			return true
		}

		Min = min(Min, nums[i])
		Max = max(Max, nums[i])
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
