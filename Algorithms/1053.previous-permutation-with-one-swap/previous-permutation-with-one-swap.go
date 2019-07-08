package problem1053

func prevPermOpt1(A []int) []int {
	n := len(A)
	stack, top := make([]int, n), 0
	stack[top] = 0

	for i := 1; i < n; i++ {
		if A[i] <= A[stack[top]] {
			top++
			stack[top] = i
			continue
		}
		for top-1 >= 0 && A[i] > A[stack[top-1]] {
			top--
		}
		j := stack[top]
		A[i], A[j] = A[j], A[i]
		break
	}

	return A
}
