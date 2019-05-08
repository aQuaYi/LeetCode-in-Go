package problem0989

func addToArrayForm(A []int, K int) []int {
	reverse(A)
	A[0] += K
	size := len(A)
	for i := 0; i+1 < size && A[i] > 9; i++ {
		A[i+1] += A[i] / 10
		A[i] %= 10
	}
	var tail int
	for A[size-1] > 9 {
		A[size-1], tail = A[size-1]%10, A[size-1]/10
		A = append(A, tail)
		size++
	}
	reverse(A)
	return A
}

func reverse(A []int) {
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}
