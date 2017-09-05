package Problem0065

func isNumber(s string) bool {
	s = trim(s)
	if len(s) == 0 {
		return false
	}

	if yes, idx := hasE(s); yes {
		return isReal(s[:idx]) && isInteger(s[idx+1:])
	}

	return isReal(s)
}

func hasE(s string) (bool, int) {
	for i, c := range s {
		if c == 'e' {
			return true, i
		}
	}

	return false, 0
}

func isReal(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		return isNonnegReal(s[1:])
	}

	return isNonnegReal(s)
}

func isNonnegReal(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i, c := range s {
		switch {
		case '0' <= c && c <= '9':
			continue
		case c == '.':
			if i == len(s)-1 && i != 0 {
				return true
			}
			return isNonnegativeInteger(s[i+1:])
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
