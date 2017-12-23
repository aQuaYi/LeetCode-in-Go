package Problem0665

func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}

	return true
}
