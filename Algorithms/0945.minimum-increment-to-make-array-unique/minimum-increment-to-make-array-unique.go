package problem0945

import "sort"

func minIncrementForUnique(A []int) int {

	sort.Ints(A)

	size := len(A)
	res := 0
	for i := 1; i < size; i++ {
		for A[i-1] >= A[i] {
			A[i]++
			res++
		}
	}

	return res
}
