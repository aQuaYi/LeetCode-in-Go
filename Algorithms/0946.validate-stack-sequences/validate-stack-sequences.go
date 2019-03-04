package problem0946

func validateStackSequences(pushed, popped []int) bool {
	size := len(pushed)

	stack := make([]int, size)
	top := -1

	for in, out := 0, 0; in < size; in++ {
		if pushed[in] != popped[out] {
			top++
			stack[top] = pushed[in]
		} else {
			out++
			for top >= 0 && stack[top] == popped[out] {
				top--
				out++
			}
		}
	}

	return top == -1
}
