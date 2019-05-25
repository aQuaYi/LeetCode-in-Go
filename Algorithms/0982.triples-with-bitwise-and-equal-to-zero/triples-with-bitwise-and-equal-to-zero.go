package problem0982

func countTriplets(A []int) int {
	n := len(A)
	res := 0
	for i := 0; i < n; i++ {
		Ai := A[i]
		for j := 0; j < n; j++ {
			Aj := A[j]
			for k := 0; k < n; k++ {
				Ak := A[k]
				if Ai&Aj&Ak == 0 {
					res++
				}
			}
		}
	}
	return res
}
