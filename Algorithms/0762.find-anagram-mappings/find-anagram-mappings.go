package Problem0762

func anagramMappings(A, B []int) []int {
	indexs := make(map[int]int, len(A))
	values := make(map[int]int, len(B))
	res := make([]int, len(B))

	for i := range res {
		if A[i] == B[i] {
			res[i] = i
			continue
		}

		if v, ok := values[A[i]]; ok {
			res[i] = v
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
