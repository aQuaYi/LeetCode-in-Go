package problem1144

import "math"

func movesToMakeZigzag(A []int) int {
	return min(
		move(math.MaxInt64, 0, A),
		move(A[0], 1, A),
	)
}

// decrease A[i] to satisfy A[i-1] > A[i] < A[i+1]
func move(left, i int, A []int) int {
	n := len(A)
	res := 0
	for i+1 < n {
		right := A[i+1]
		minVal := min(left, right)
		if A[i] >= minVal {
			res += A[i] - minVal + 1
		}
		left = right
		i += 2
	}
	if i+1 == n &&
		A[i] >= left {
		res += A[i] - left + 1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
