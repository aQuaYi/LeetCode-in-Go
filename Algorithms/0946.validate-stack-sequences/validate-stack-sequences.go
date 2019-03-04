package problem0946

func validateStackSequences(pushed, popped []int) bool {
	size := len(pushed)
	top, in, out := -1, 0, 0

	for in < size {
		if pushed[in] == popped[out] {
			out++
			for top >= 0 && pushed[top] == popped[out] {
				top--
				out++
			}
		} else {
			top++
		}
		in++
	}

	return top == -1 && out == size
}
