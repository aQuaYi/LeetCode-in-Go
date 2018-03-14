package problem0454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	n := len(A)
	sum := make(map[int]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum[C[i]+D[j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += sum[-(A[i] + B[j])]
		}
	}

	return res
}
