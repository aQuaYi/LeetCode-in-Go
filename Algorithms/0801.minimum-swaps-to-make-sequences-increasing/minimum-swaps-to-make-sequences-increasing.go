package problem0801

func minSwap(a []int, b []int) int {
	res := 0
	size := len(a)
	if size%2 == 0 {
		a = append(a, 3000)
		b = append(b, 3000)
		size++
	}

	for i := 1; i < size; i += 2 {
		if a[i-1] < a[i] && a[i] < a[i+1] &&
			b[i-1] < b[i] && b[i] < b[i+1] {
			continue
		}
		a[i], b[i] = b[i], a[i]
		res++
	}

	return res
}
