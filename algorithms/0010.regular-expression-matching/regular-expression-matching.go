package Problem0010

// 程序中存在以下假设
// "*" 不会出现在p的首位
// "**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

func isMatch(s, p string) bool {
	if p == "" {
		return s == ""
	}

	if len(p) == 1 || (len(p) > 1 && p[1] != '*') {
		if s == "" || (p[0] != '.' && s[0] != p[0]) {
			return false
		}

		return isMatch(s[1:], p[1:])
	}

	slen := len(s)
	i := -1

	for i < slen && (i < 0 || p[0] == '.' || p[0] == s[i]) {
		if len(p) >= 2 {
			if isMatch(s[i+1:], p[2:]) {
				return true
			}
		} else {
			if isMatch(s[i+1:], p[1:]) {
				return true
			}
		}
		i++
	}

	return false
}
