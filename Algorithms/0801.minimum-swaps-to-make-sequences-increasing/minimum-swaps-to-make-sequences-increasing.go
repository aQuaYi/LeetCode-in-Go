package problem0801

func minSwap(a []int, b []int) int {
	n := len(a)
	r0, r1 := 0, 1

	for i := 1; i < n; i++ {
		newR0, newR1 := 1<<63-1, 1<<63-1
		if a[i] > a[i-1] && b[i] > b[i-1] {
			newR0 = min(newR0, r0)
			newR1 = min(newR1, r1+1)
		}
		if a[i] > b[i-1] && b[i] > a[i-1] {
			newR0 = min(newR0, r1)
			newR1 = min(newR1, r0+1)
		}
		r0, r1 = newR0, newR1
	}
	return min(r0, r1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
