package problem0915

func partitionDisjoint(A []int) int {
	leftMax, nextMax, rightBegin := A[0], A[0], 1

	for i := 1; i < len(A); i++ {
		if leftMax > A[i] {
			// 此时，A[i] 比在 left 中，
			leftMax = nextMax
			// right 的起点，至少是 i+1
			rightBegin = i + 1
		} else {
			nextMax = max(nextMax, A[i])
		}
	}

	return rightBegin
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
