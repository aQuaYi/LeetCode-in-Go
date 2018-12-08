package problem0916

func wordSubsets(A, B []string) []string {
	clt := new([26]int)
	for _, b := range B {
		collect(clt, count(b))
	}

	res := make([]string, 0, len(A))
	for _, a := range A {
		if isSubset(count(a), clt) {
			res = append(res, a)
		}
	}

	return res
}

func count(s string) *[26]int {
	res := [26]int{}
	for _, b := range s {
		res[b-'a']++
	}
	return &res
}

func isSubset(s, clt *[26]int) bool {
	isSubset := true
	i := 0
	for isSubset && i < 26 {
		isSubset = isSubset && (s[i] >= clt[i])
		i++
	}
	return isSubset
}

// collect to clt
func collect(clt, b *[26]int) {
	for i := range clt {
		clt[i] = max(clt[i], b[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
