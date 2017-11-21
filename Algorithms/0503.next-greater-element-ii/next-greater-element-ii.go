package Problem0503

func nextGreaterElements(nums []int) []int {
	size := len(nums)

	res := make([]int, size)
	for i := range res {
		res[i] = -1
	}

	stack := make([]int, 0, size)

	for i := 0; i < size*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}
