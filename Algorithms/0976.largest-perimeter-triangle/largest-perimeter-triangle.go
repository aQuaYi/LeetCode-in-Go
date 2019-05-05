package problem0976

import "sort"

func largestPerimeter(A []int) int {
	size := len(A)
	sort.Ints(A)
	res := 0
	for i := 0; i < size; i++ {
		a := A[i]
		for j := i + 1; j < size; j++ {
			b := A[j]
			for k := j + 1; k < size; k++ {
				c := A[k]
				if isOK(a, b, c) {
					res = max(res, a+b+c)
				}
			}
		}
	}
	return res
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
