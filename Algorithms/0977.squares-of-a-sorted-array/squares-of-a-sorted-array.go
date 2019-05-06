package problem0977

import "sort"

func sortedSquares(A []int) []int {
	r := sort.SearchInts(A, 0)
	l := r - 1
	size := len(A)
	res := make([]int, 0, size)
	for 0 <= l || r < size {
		if r == size ||
			(0 <= l && A[l]+A[r] > 0) {
			res = append(res, A[l]*A[l])
			l--
		} else {
			res = append(res, A[r]*A[r])
			r++
		}
	}
	return res
}
