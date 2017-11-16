package Problem0493

func reversePairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				res++
			}
		}
	}

	return res
}
