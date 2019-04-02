package problem0969

func pancakeSort(A []int) []int {
	size := len(A)
	res := make([]int, 0, size*2)
	for Ai := size; 0 < Ai; Ai-- {
		i := Ai - 1
		if A[i] == Ai {
			continue
		}
		for A[i] != Ai {
			i--
		}
		swap(A, i+1)
		swap(A, Ai)
		res = append(res, i+1, Ai)
	}
	return res
}

func swap(A []int, k int) {
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}
