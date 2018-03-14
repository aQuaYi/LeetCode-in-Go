package problem0678

func checkValidString(s string) bool {
	l, r := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		// 从左侧开始的扫描
		// 并认为所有的 '*' 是 '('
		if s[i] == ')' {
			l--
		} else {
			l++
		}

		// 从右侧开始的扫描
		// 并认为所有的 '*' 是 ')'
		j := n - i - 1
		if s[j] == '(' {
			r--
		} else {
			r++
		}

		if l < 0 || r < 0 {
			// l < 0 意味着
			// 就算所有的 '*' 变成的 '('
			// s[:i] 中也没有能与 s[i] == ')' 匹配的 '('
			// l < 0 意味着
			// 就算所有的 '*' 变成的 ')'
			// s[j+1:] 中也没有能与 s[j] == '(' 匹配的 ')'
			return false
		}
	}

	return true
}
