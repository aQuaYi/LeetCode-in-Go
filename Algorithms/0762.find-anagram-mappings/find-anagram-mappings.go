package Problem0762

func anagramMappings(A, B []int) []int {
	n := len(A)
	indexs := make(map[int]int, n)
	values := make(map[int]int, n)
	res := make([]int, n)

	for i := range res {
		if A[i] == B[i] {
			res[i] = i
			continue
		}

		if val, ok := values[A[i]]; ok {
			res[i] = val
			delete(values, A[i])
		} else {
			indexs[A[i]] = i
		}

		if idx, ok := indexs[B[i]]; ok {
			res[idx] = i
			delete(indexs, B[i])
		} else {
			values[B[i]] = i
		}
	}

	return res
}
