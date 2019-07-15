package problem1089

func duplicateZeros(A []int) {
	n := len(A)
	//
	count := 0
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			count++
		}
	}
	// copy A[i] to A[j]
	copy := func(i, j int) {
		if j < n {
			A[j] = A[i]
		}
	}
	//
	i, j := n-1, n+count-1
	for i < j {
		copy(i, j)
		if A[i] == 0 {
			j--
			copy(i, j)
		}
		i--
		j--
	}
}
