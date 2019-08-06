package problem1131

// ref: https://leetcode.com/problems/maximum-of-absolute-value-expression/discuss/340070/topic

func maxAbsValExpr(A, B []int) int {
	Q := [2]int{-1, 1}

	n := len(A)
	res := 0
	for _, a := range Q {
		for _, b := range Q {
			low := a*A[0] + b*B[0]
			for i := 1; i < n; i++ {
				cur := a*A[i] + b*B[i] + i
				res = max(res, cur-low)
				low = min(low, cur)
			}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
