package problem1073

func addNegabinary(A, B []int) []int {
	i, j := len(A)-1, len(B)-1
	res := make([]int, 0, i+j)
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += A[i]
			i--
		}
		if j >= 0 {
			carry += B[j]
			j--
		}
		res = append(res, carry&1)
		carry = -(carry >> 1)
	}

	res = reverse(res)

	// cut leading zero
	i, end := 0, len(res)-1
	for i < end && res[i] == 0 {
		i++
	}
	return res[i:]
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
