package Problem0080

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	res := 2
	for i := 2; i < length; i++ {
		if nums[i] != nums[res-2] {
			nums[res] = nums[i]
			res++
		}
	}

	return res
}
