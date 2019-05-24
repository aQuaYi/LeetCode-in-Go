package problem0978

func maxTurbulenceSize(A []int) int {
	size := len(A)
	greater, less := 1, 1

	res := 1
	for i := 1; i < size; i++ {
		switch {
		case A[i] > A[i-1]:
			greater, less = less+1, 1
			res = max(res, greater)
		case A[i] < A[i-1]:
			greater, less = 1, greater+1
			res = max(res, less)
		default:
			greater, less = 1, 1
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
