package Problem0421

func findMaximumXOR(nums []int) int {
	maxXOR := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			maxXOR = max(maxXOR, nums[i]^nums[j])
		}
	}

	return maxXOR
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
