package Problem0486

// PredictTheWinner 在 play1 能赢的时候，返回 true
func PredictTheWinner(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	return search(nums, 0, len(nums)-1, 0, 0)
}

func search(nums []int, left, right, s1, s2 int) bool {
	if left > right {
		return s1 >= s2
	}

	return !search(nums, left+1, right, s2, s1+nums[left]) ||
		!search(nums, left, right-1, s2, s1+nums[right])
}
