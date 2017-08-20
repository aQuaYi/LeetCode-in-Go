package Problem0448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := []int{}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			res = append(res, i+1)
		}
	}

	return res
}
