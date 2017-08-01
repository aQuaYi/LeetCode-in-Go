package Problem0026

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if res != i {
			nums[res] = nums[i]
		}

		res++
	}

	return res
}
