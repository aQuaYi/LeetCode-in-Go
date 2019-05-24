package problem0978

func maxTurbulenceSize(A []int) int {
	size := len(A)
	greater, less := make([]int, size), make([]int, size)

	res := 0
	for i := 1; i < size; i++ {
		switch {
		case A[i] > A[i-1]:
			greater[i] = less[i-1] + 1
			res = max(res, greater[i])
		case A[i] < A[i-1]:
			less[i] = greater[i-1] + 1
			res = max(res, less[i])
		}
	}

	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
