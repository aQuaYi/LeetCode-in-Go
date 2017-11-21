package Problem0496

func nextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	var j int
	for i, n := range nums {
		for j = i + 1; j < len(nums) && n >= nums[j]; j++ {
		}

		if j == len(nums) {
			m[n] = -1
		} else {
			m[n] = nums[j]
		}
	}

	res := make([]int, len(findNums))
	for i, n := range findNums {
		res[i] = m[n]
	}

	return res
}
