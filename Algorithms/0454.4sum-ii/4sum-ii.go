package Problem0454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	dHas := make(map[int]int, len(D))
	for i := 0; i < len(D); i++ {
		dHas[D[i]]++
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			for k := 0; k < len(C); k++ {
				res += dHas[0-(A[i]+B[j]+C[k])]
			}
		}
	}

	return res
}
