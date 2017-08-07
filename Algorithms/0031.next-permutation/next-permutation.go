package Problem0031

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i != -1 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		if j >= 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
