package problem1004

func longestOnes(A []int, K int) int {
	A = append(A, 0)
	size := len(A)
	indexs := make([]int, 0, size)
	if A[0] == 1 {
		indexs = append(indexs, 0)
	}
	for i := 1; i < size; i++ {
		if (A[i-1] == 0 && A[i] == 1) ||
			(A[i-1] == 1 && A[i] == 0) {
			indexs = append(indexs, i)
		}
	}

	res := 0
	for i := 0; i < len(indexs); i += 2 {
		k := K
		for j := 2; j < size && k > 0; j += 2 {

		}
	}
	return min(res, size-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
