package problem0213

func rob(nums []int) int {
	size := len(nums)
	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	// 把环斩断，就变成了 198 题
	// 斩断的方式分为，在 0 位和不在 0 位 两种。
	return max(robbing(nums[1:]), robbing(nums[:size-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func robbing(nums []int) int {
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
