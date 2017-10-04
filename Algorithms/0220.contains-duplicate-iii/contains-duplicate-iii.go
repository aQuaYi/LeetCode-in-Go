package Problem0220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i+k < len(nums); i++ {
		for j := i + 1; j <= i+k; j++ {
			if abs(nums[i]-nums[j]) > t {
				return false
			}
		}
	}

	return true
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
