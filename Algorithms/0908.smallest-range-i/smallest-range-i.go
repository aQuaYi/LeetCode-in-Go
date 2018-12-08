package problem0908

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for i := range A {
		if A[i] < min {
			min = A[i]
		} else if max < A[i] {
			max = A[i]
		}
	}

	if min+K >= max-K {
		return 0
	}

	return max - min - K*2
}
