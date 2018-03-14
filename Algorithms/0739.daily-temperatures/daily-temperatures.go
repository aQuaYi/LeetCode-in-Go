package problem0739

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)

	stack := make([]int, n)

	top := -1
	for i := 0; i < n; i++ {
		for top >= 0 && temperatures[stack[top]] < temperatures[i] {
			res[stack[top]] = i - stack[top]
			top--
		}

		top++
		stack[top] = i
	}

	return res
}
