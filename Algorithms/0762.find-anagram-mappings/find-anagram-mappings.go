package problem0762

func anagramMappings(A, B []int) []int {
	n := len(A)
	values := make(map[int]int, n)
	res := make([]int, n)

	for i := range B {
		values[B[i]] = i
	}

	for i := range res {
		res[i] = values[A[i]]
	}

	return res
}
