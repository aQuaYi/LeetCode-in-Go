package problem1144

import "math"

func movesToMakeZigzag(nums []int) int {
	return min(
		search(0, nums),
		search(1, nums),
	)
}

// decrease A[i] to satisfy A[i-1] > A[i] < A[i+1]
func search(i int, A []int) int {
	var prev, next int
	if i == 0 {
		prev = math.MaxInt64
	} else {
		prev = A[i-1]
	}
	n := len(A)
	res := 0
	for i < n {
		if i+1 == n {
			next = math.MaxInt64
		} else {
			next = A[i+1]
		}
		minVal := min(prev, next)
		if A[i] >= minVal {
			res += A[i] - minVal + 1
		}
		prev = next
		i += 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
