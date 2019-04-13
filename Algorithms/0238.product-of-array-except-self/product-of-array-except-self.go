package problem0238

func productExceptSelf(nums []int) []int {
	size := len(nums)

	res := make([]int, size)

	left := 1 // product of all left
	for i := 0; i < size; i++ {
		res[i] = left
		left *= nums[i]
	}

	right := 1 // product of all right
	for i := size - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
