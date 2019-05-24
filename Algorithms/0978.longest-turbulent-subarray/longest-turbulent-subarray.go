package problem0978

func maxTurbulenceSize(A []int) int {
	size := len(A)
	greater, less := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		greater[i] = 1
		less[i] = 1
	}

	for i := 1; i < size; i++ {
		switch {
		case A[i] > A[i-1]:
			greater[i] = less[i-1] + 1
		case A[i] < A[i-1]:
			less[i] = greater[i-1] + 1
		}
	}

	return max(maxOf(greater), maxOf(less))
}

func maxOf(A []int) int {
	res := A[0]
	for i := 1; i < len(A); i++ {
		res = max(res, A[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
