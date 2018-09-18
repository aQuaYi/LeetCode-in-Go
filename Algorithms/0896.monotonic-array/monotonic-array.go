package problem0896

func isMonotonic(A []int) bool {
	size := len(A)
	inc, dec := true, true
	for i := 1; i < size && (inc || dec); i++ {
		inc = inc && A[i-1] <= A[i]
		dec = dec && A[i-1] >= A[i]
	}

	return inc || dec
}
