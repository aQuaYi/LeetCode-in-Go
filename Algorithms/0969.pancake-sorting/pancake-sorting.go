package problem0969

func pancakeSort(A []int) []int {
	size := len(A)
	res := make([]int, 0, size*2)
	for Ai := size; Ai > 1; Ai-- {
		// every round makes A[i-1] is i
		if A[Ai-1] == Ai {
			continue
		}
		i := Ai - 2
		for A[i] != Ai {
			i--
		}
		if i != 0 {
			reverse(A, i+1) // move Ai to A[0]
			res = append(res, i+1)
		}
		reverse(A, Ai) // move Ai to A[Ai-1]
		res = append(res, Ai)
	}
	return res
}

func reverse(A []int, k int) {
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}
