package problem1131

func maxAbsValExpr(A, B []int) int {
	n := len(A)
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			res = max(
				res,
				abs(A[i]-A[j])+abs(B[i]-B[j])+i-j,
			)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
