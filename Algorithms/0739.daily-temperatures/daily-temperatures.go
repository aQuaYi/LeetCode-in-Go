package Problem0739

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)

	stack := make([]record, 0, n)

	last := -1
	for i := 0; i < n; i++ {
		for last >= 0 && stack[last].t < temperatures[i] {
			res[stack[last].i] = i - stack[last].i
			stack = stack[:last]
			last--
		}

		stack = append(stack, record{t: temperatures[i], i: i})
		last++
	}

	return res
}

type record struct {
	t, i int
}
