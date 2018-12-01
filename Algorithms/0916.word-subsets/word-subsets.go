package problem0916

func wordSubsets(A, B []string) []string {
	mix := new([26]int)
	for _, b := range B {
		collect(mix, count(b))
	}

	res := make([]string, 0, len(A))
	for _, a := range A {
		if isSubset(count(a), mix) {
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

func isSubset(s, mix *[26]int) bool {
	isSubset := true
	i := 0
	for isSubset && i < 26 {
		isSubset = isSubset && (s[i] >= mix[i])
		i++
	}
	return isSubset
}

// collect to a
func collect(a, b *[26]int) {
	for i := range a {
		a[i] = max(a[i], b[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
