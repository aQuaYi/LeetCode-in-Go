package Problem0301

func removeInvalidParentheses(s string) []string {
	res := []string{}
	var dfs func()
	dfs = func() {}

	return res
}

func count(s string) (left, right int) {
	for i := range s {
		switch s[i] {
		case '(':
			left++
		case ')':
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}

	return
}

func isOK(s string) bool {
	l, r := count(s)
	return l+r == 0
}
