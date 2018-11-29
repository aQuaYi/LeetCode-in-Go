package problem0915

func partitionDisjoint(A []int) int {
	size := len(A)
	inStack := make([]int, 1, size)
	deStack := make([]int, 1, size)
	inStack[0] = 0
	deStack[0] = size - 1

	push := func(stack []int, i int) {
		stack = append(stack, i)
	}

	top := func(stack []int) int {
		i := stack[len(stack)-1]
		return A[i]
	}

	minRight := A[0]

	for i := 1; i < size; i++ {
		if top(inStack) < A[i] {
			push(inStack, i)
		}
	}

	for i := size - 2; i >= 0; i-- {
		if top(deStack) > A[i] {
			push(deStack, i)
		}
	}

	for i, j := 0, len(deStack)-1; i < j; i, j = i+1, j-1 {
		deStack[i], deStack[j] = deStack[i], deStack[i]
	}

	return i
}
