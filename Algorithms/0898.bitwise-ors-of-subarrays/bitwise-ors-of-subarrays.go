package problem0898

func subarrayBitwiseORs(A []int) int {
	size := len(A)
	tmp := make(map[int]bool, size*2)
	for i := 0; i < size; i++ {
		for j := i + 1; j <= size; j++ {
			tmp[bor(A[i:j])] = true
		}
	}
	return len(tmp)
}

func bor(a []int) int {
	res := 0
	for i := range a {
		res |= a[i]
	}
	return res
}
