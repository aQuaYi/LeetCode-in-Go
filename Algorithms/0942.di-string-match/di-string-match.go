package problem0942

func diStringMatch(S string) []int {
	size := len(S)
	i, l, r := 0, 0, size
	a := make([]int, size+1)
	for i < size {
		switch S[i] {
		case 'I':
			a[i] = l
			l++
		case 'D':
			a[i] = r
			r--
		}
		i++
	}

	a[i] = l

	return a
}
