package Problem0219

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, v := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if v == nums[j] {
				return true
			}
		}
	}

	return false
}
