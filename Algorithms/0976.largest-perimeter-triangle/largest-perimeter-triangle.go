package problem0976

import "sort"

func largestPerimeter(A []int) int {
	size := len(A)
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i := 0; i < size; i++ {
		a := A[i]
		for j := i + 1; j < size; j++ {
			b := A[j]
			for k := j + 1; k < size; k++ {
				c := A[k]
				if a < b+c {
					return a + b + c
				}
			}
		}
	}
	return 0
}

func isOK(a, b, c int) bool {
	return a+b > c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
