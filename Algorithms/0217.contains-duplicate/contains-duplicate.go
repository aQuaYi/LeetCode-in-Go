package problem0217

func containsDuplicate(nums []int) bool {
	// 利用 m 记录 nums 中的元素是否出现过
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}

	return false
}
