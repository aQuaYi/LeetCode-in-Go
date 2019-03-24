package problem0962

func maxWidthRamp(A []int) int {
	size := len(A)

	stack := make([]int, 1, size)
	top := 0
	for i := 1; i < size; i++ {
		if A[stack[top]] > A[i] {
			stack = append(stack, i)
			top++
		}
	}

	res := 0
	for j := size - 1; j >= 0 && top >= 0; j-- {
		width := 0
		for top >= 0 && A[stack[top]] <= A[j] {
			width = j - stack[top]
			stack = stack[:top]
			top--
		}
		res = max(res, width)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
