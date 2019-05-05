package problem0976

import "sort"

func largestPerimeter(A []int) int {
	size := len(A)
	sort.Ints(A)
	a, b := A[size-1], A[size-2]
	for i := size - 3; i >= 0; i-- {
		c := A[i]
		if a < b+c {
			return a + b + c
		}
		a, b = b, c
	}
	return 0
}
