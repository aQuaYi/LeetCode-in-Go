package problem1054

func rearrangeBarcodes(A []int) []int {
	n := len(A)
	for i := 1; i < n; i++ {
		if A[i] != A[i-1] {
			continue
		}
		j := n - 1
		for A[j] == A[i] {
			j--
		}
		A[i], A[j] = A[j], A[i]
	}
	return A
}
