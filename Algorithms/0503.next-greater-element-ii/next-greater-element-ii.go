package Problem0503

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	res := make([]int, size)

	for i, n := range nums {
		res[i] = -1

		j := i + 1
		for {
			if n > nums[j%size] {
				j++
				continue
			}

			if n == nums[j%size] {
				res[i] = res[j%size]
			} else {
				res[i] = nums[j%size]
			}
			break
		}
	}

	return res
}
