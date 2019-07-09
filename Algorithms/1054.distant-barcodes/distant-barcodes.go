package problem1054

import "sort"

func rearrangeBarcodes(A []int) []int {
	n := len(A)
	sort.Ints(A)
	for i, j := 1, (n-1)/2*2; i < j; i, j = i+2, j-2 {
		A[i], A[j] = A[j], A[i]
	}
	return A
}
