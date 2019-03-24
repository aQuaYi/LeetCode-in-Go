package problem0962

func maxWidthRamp(A []int) int {
	size := len(A)
	res := 0
	stack := make([]int, size)
	top := 0
	for i := 1; i < size; i++ {
		if A[stack[top]] > A[i] {
			top++
			stack[top] = i
			continue
		}

		for j := top; j >= 0; j-- {
			if A[stack[j]] > A[i] {
				break
			}
			res = i - stack[j]
		}
	}
	return res
}
