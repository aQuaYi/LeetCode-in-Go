package Problem0268

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] < len(nums) && nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	for i, n := range nums {
		if i != n {
			return i
		}
	}

	return len(nums)
}
