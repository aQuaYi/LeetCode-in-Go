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

	var dfs func([]string)
	dfs = func(ss []string) {
		for i := 0; i+1 < len(ss); i++ {
			if isPalindrome(ss[i] + ss[i+1]) {
				temp := make([]string, len(ss)-1)
				copy(temp[:i], ss)
				temp[i] = ss[i] + ss[i+1]
				copy(temp[i+1:], ss[i+2:])
				res = append(res, temp)
				dfs(temp)
			}
		}
	}

	dfs(basic)

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
