package problem0942

func diStringMatch(S string) []int {
	size := len(S)
	i, l, r := 0, 0, size
	a := make([]int, size+1)
	bytes := []byte(S)
	for i < size {
		if bytes[i] == 'I' {
			a[i] = l
			l++
		} else {
			a[i] = r
			r--
		}
		i++
	}

	a[i] = l

	return a
}
