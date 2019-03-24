package problem0962

import "sort"

func maxWidthRamp(A []int) int {
	size := len(A)
	res := 0
	stack := make([]int, 1, size)

	for i := 1; i < size; i++ {
		x := A[i]
		top := len(stack) - 1
		if A[stack[top]] > x {
			// keep stack decrease
			stack = append(stack, i)
			continue
		}

		j := sort.Search(len(stack), func(i int) bool {
			return A[stack[i]] <= x
		})
		res = max(res, i-stack[j])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
