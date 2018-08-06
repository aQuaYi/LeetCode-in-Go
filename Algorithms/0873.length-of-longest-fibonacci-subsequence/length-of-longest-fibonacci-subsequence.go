package problem0873

func lenLongestFibSubseq(A []int) int {
	size := len(A)

	idxs := make(map[int]int, size)
	mp := make(map[int]bool, size)
	for i, a := range A {
		idxs[a] = i
		mp[a] = true
	}

	res := 0

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			tmp := 2
			i0, i1 := i, j
			for mp[A[i0]+A[i1]] {
				tmp++
				res = max(res, tmp)
				i0, i1 = i1, idxs[A[i0]+A[i1]]
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
