package Problem0080

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	i, res := 2, 2

	for ; i < length; i++ {
		if nums[i] == nums[res-1] && nums[i] == nums[res-2] {
			continue
		}

		if res != i {
			nums[res] = nums[i]
		}
		res++
	}

	return res
}
