package Problem0394

func decodeString(s string) string {
	var i, beg, end int
	for i = 0; i < len(s); i++ {
		if s[i] == '[' {
			beg = i
			break
		}
	}

	if i == len(s) {
		return s
	}

	count := 1

	for end = i + 1; end < len(s); end++ {
		if s[end] == '[' {
			count++
		} else if s[end] == ']' {
			count--
		}

		if count == 0 {
			break
		}
	}

	return ""
}

func times(n int, s string) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
