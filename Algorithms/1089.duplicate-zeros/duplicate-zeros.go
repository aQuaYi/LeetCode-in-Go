package problem1089

func duplicateZeros(A []int) {
	n := len(A)
	//
	i, count := 0, 0
	for i+count < n {
		if A[i] == 0 {
			count++
		}
		i++
	}
	i--
	//
	j := n - 1
	for count > 0 {
		if A[i] == 0 {
			A[j-1], A[j] = 0, 0
			j -= 2
			count--
		} else {
			A[j] = A[i]
			j--
		}
		i--
	}
}
