package problem0065

func isNumber(s string) bool {
	// 去掉首位的空格
	s = trim(s)

	// 判断是否是实数
	return isReal(s)
}

func isReal(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		return isNonnegReal(s[1:], false)
	}

	return isNonnegReal(s, false)
}

func isNonnegReal(s string, hasDot bool) bool {
	if len(s) == 0 {
		return false
	}

	for i, c := range s {
		switch {
		case '0' <= c && c <= '9':
			continue
		case c == '.':
			if hasDot {
				// 前面已经有了一个 '.' 了
				return false
			}

			if i == len(s)-1 && i != 0 {
				// 以 '.' 结尾的情况
				return true
			}

			if i+1 < len(s) && s[i+1] == 'e' {
				// "2.e3" 是正确的数字表
				// 但 ".e1" 不是
				return i != 0 && isInteger(s[i+2:])
			}

			// 继续判断是否为 非负实数
			return isNonnegReal(s[i+1:], true)
		case c == 'e':
			if i == 0 {
				// 'e' 不能开头
				return false
			}
			return isInteger(s[i+1:])
		default:
			return false
		}
	}

	return true
}

func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		return isNonnegativeInteger(s[1:])
	}

	return isNonnegativeInteger(s)
}

func isNonnegativeInteger(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, c := range s {
		if c < '0' || '9' < c {
			return false
		}
	}
	return true
}

func trim(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	j := len(s) - 1
	for i <= j && s[j] == ' ' {
		j--
	}

	return s[i : j+1]
}
