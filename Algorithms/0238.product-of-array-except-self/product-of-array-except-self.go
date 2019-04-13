package problem0238

func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)

	// 从左到右, res[i] 是 nums[i] 左侧所有元素的乘积
	res[0], res[1] = 1, nums[0]

	// 题目给出 size > 1
	for i := 2; i <= size-1; i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	// 从右到左, 算出 res[i] 两侧的乘积
	right := 1
	for i := size - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right *= nums[i]
	}

	return res
}
