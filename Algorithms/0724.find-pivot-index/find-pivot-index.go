package problem0724

func pivotIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	left := 0
	for i := range nums {
		if left*2+nums[i] == sum {
			return i
		}
		left += nums[i]
	}

	return -1
}
