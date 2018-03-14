package problem0189

func rotate(nums []int, k int) {
	// 我已经假定 k >= 0

	n := len(nums)

	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
