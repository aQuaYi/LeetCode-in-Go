package Problem0198

func rob(nums []int) int {
	return max(robbing(nums), robbing(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robbing(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}
