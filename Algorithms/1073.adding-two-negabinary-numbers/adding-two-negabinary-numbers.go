package problem1073

func addNegabinary(A, B []int) []int {
	A, B = reverse(A), reverse(B)
	if len(A) < len(B) {
		A, B = B, A
	}
	res := make([]int, 0, len(A)+2)
	carry := 0
	n := len(B)
	for i, a := range A {
		res = append(res, a)
		if i < n {
			res[i] += B[i]
		}
		res[i] += carry
		carry = 0
		switch res[i] {
		case -1:
			res[i], carry = 1, 1
		case 2:
			res[i], carry = 0, -1
		case 3:
			res[i], carry = 1, -1
		}
	}

	switch carry {
	case 1:
		res = append(res, 1)
	case -1:
		res = append(res, 1, 1)
	}

	return reverse(res)
}

func reverse(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
	return A
}
