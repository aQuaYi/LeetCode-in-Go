package problem1002

func commonChars(A []string) []string {
	minCount := count(A[0])
	for i := 1; i < len(A); i++ {
		c := count(A[i])
		for b, v := range c {
			minCount[b] = min(minCount[b], v)
		}
	}
	res := make([]string, 0, 128)

	for b, v := range minCount {
		for v > 0 {
			res = append(res, string(b+'a'))
			v--
		}
	}

	return res
}

func count(s string) []int {
	res := make([]int, 26)
	for _, b := range s {
		res[b-'a']++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
