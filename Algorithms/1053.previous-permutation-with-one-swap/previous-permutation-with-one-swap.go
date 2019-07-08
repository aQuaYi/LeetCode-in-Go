package problem1053

func prevPermOpt1(A []int) []int {
	n := len(A)
	stack, top := make([]int, n), 0
	stack[top] = n - 1

	for i := n - 2; i >= 0; i-- {
		if A[i] <= A[stack[top]] {
			top++
			stack[top] = i
			continue
		}
		for top-1 >= 0 &&
			A[stack[top]] < A[stack[top-1]] &&
			A[stack[top-1]] < A[i] {
			top--
		}
		j := stack[top]
		A[i], A[j] = A[j], A[i]
		break
	}

	return A
}
