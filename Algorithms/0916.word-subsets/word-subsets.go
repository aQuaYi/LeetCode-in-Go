package problem0916

func wordSubsets(A, B []string) []string {
	mix := make([]int, 26)
	for i := range B {
		maxSlice(mix, count(B[i]))
	}

	res := make([]string, 0, len(A))
	for _, a := range A {
		if isSubset(count(a), mix) {
			res = append(res, a)
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

func isSubset(s, mix []int) bool {
	res := true
	i := 0
	for res && i < 26 {
		res = res && (s[i] >= mix[i])
		i++
	}
	return res
}

// maxSlice to a
func maxSlice(a, b []int) {
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
