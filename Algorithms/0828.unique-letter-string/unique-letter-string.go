package problem0828

const module = 1e9 + 7

func uniqueLetterString(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	index := [26][]int{}
	for i := range index {
		index[i] = make([]int, 0, size)
	}

	for i, b := range s {
		index[b-'A'] = append(index[b-'A'], i)
	}

	res := 0
	for i := range index {
		if len(index[i]) == 0 {
			continue
		}
		index[i] = append(index[i], size)
		a, b, c := 0, -1, index[i][0]
		for j := 1; j < len(index[i]); j++ {
			a, b, c = b, c, index[i][j]
			res += ((b - a) * (c - b)) % module
			res %= module
		}
	}

	return res
}
