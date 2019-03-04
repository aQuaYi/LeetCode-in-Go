package problem0946

func validateStackSequences(pushed, popped []int) bool {
	size := len(pushed)
	in, out, s := 0, 0, 0
	stack := make([]int, size)

	for in = 0; in < size && out < size; in++ {
		if pushed[in] == popped[out] {
			out++
			for s > 0 && stack[s-1] == popped[out] {
				s--
				out++
			}
			continue
		}
		stack[s] = pushed[in]
		s++
	}

	return s == 0 && out == size
}
