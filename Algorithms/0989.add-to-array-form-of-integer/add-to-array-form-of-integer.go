package problem0989

func addToArrayForm(A []int, K int) []int {
	size := len(A)
	A[size-1] += K
	for i := size - 1; i > 0 && A[i] > 9; i-- {
		A[i-1] += A[i] / 10
		A[i] %= 10
	}

	if A[0] < 10 {
		return A
	}

	A0 := num2ints(A[0])
	return append(A0, A[1:]...)
}

func num2ints(n int) []int {
	res := make([]int, 0, 8)
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	reverse(res)
	return res
}

func reverse(A []int) {
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}
