package Problem0217

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = true
	}

	return false
}
