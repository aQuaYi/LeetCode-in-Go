package problem1014

func maxScoreSightseeingPair(A []int) int {
	Ai, i := A[0], 0
	res := 0
	for j := 1; j < len(A); j++ {
		res = max(res, Ai+A[j]+i-j)
		if Ai+i < A[j]+j {
			Ai, i = A[j], j
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
