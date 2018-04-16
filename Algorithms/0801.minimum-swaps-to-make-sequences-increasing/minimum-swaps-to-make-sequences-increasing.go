package problem0801

func minSwap(a []int, b []int) int {
	res := 0
	size := len(a)

	for i := 0; i < size-1; i++ {
		if a[i] < a[i+1] {
			continue
		}
		a[i], b[i] = b[i], a[i]
		res++
	}

	for i := 0; i < size-1; i++ {
		if b[i] < b[i+1] {
			continue
		}
		a[i], b[i] = b[i], a[i]
		res++
	}

	return res
}
