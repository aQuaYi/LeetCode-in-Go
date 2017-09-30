package Problem0198

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		return max(nums[0]+rob(nums[2:]), nums[1]+rob(nums[3:]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
