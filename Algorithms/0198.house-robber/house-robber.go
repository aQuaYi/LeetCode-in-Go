package problem0198

func rob(nums []int) int {
	// a 对于偶数位上的最大值的记录
	// b 对于奇数位上的最大值的记录
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
