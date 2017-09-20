package Problem0131

func partition(s string) [][]string {
	res := [][]string{}

	if len(s) <= 1 {
		res = append(res, []string{s})
		return res
	}

	basic := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		basic[i] = s[i : i+1]
	}
	res = append(res, basic)

	search := func(ss []string, k int) {
		for i := 0; i+k < len(ss); i++ {
			s := sum(ss, i, i+k)
			if isPalindrome(s) {
				temp := make([]string, len(ss)-k)
				copy(temp[:i], ss)
				temp[i] = s
				copy(temp[i+1:], ss[i+k+1:])
				res = append(res, temp)
			}
		}
	}

	for i := 1; i < len(basic); i++ {
		search(basic, i)
	}

	return res
}

func sum(ss []string, first, last int) string {
	res := ""
	for i := first; i <= last; i++ {
		res += ss[i]
	}
	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
