package problem0448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := make([]int, 0, len(nums))

	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
