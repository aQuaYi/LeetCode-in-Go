package problem0219

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}
	
	// 利用 m 记录 a[i]中的值，每次出现的(索引号+1)
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if m[n] != 0 && i-(m[n]-1) <= k {
			return true
		}
		m[n] = i + 1
	}

	return false
}
