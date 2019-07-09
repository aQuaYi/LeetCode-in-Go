package problem1054

func rearrangeBarcodes(A []int) []int {
	n := len(A)

	// odds in left, evens in right
	i, j := 0, n-1
	for i < j {
		if A[i]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			j--
		} else {
			i++
		}
	}
	i, j = 0, n/2*2-1
	if A[n/2]%2 == 1 {
		i, j = 1, n-1
	}
	for i < j {
		A[i], A[j] = A[j], A[i]
		i += 2
		j -= 2
	}

	return A
}
