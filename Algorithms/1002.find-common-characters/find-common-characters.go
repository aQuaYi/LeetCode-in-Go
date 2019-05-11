package problem1002

func commonChars(A []string) []string {
	minCount := count(A[0])
	for i := 1; i < len(A); i++ {
		c := count(A[i])
		for j, n := range c {
			minCount[j] = min(minCount[j], n)
		}
	}

	res := make([]string, 0, 128)
	for i, n := range minCount {
		for n > 0 {
			res = append(res, string('a'+i))
			n--
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
