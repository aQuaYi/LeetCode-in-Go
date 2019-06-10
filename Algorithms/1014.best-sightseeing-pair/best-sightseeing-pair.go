package problem1014

func maxScoreSightseeingPair(A []int) int {
	a := A[0] // a is A[i]+i
	res := 0
	for j := 1; j < len(A); j++ {
		res = max(res, a+A[j]-j)
		a = max(a, A[j]+j)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
