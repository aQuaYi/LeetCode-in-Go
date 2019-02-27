package problem0942

func diStringMatch(S string) []int {
	size := len(S)
	l, r := 0, size
	a := make([]int, size+1)
	for i, b := range S {
		if b == 'I' {
			a[i] = l
			l++
		} else {
			a[i] = r
			r--
		}
		i++
	}

	a[size] = l

	return a
}
