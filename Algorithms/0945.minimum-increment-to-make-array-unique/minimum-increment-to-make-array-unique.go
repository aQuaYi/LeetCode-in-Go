package problem0945

import "sort"

func minIncrementForUnique(A []int) int {

	sort.Ints(A)

	size := len(A)
	res := 0
	for i := 1; i < size; i++ {
		d := A[i-1] - A[i] + 1
		if d > 0 {
			res += d
			A[i] += d
		}
	}

	return res
}
