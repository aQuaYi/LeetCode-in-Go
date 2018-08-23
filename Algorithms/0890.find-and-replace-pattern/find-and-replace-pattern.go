package problem0890

func findAndReplacePattern(words []string, pattern string) []string {
	np := normal(pattern)
	res := make([]string, 0, len(words))
	for _, w := range words {
		if normal(w) == np {
			res = append(res, w)
		}
	}
	return res
}

// 把 s 标准化
func normal(s string) string {
	m := make(map[rune]byte, len(s))
	for _, c := range s {
		/**按照字符在 s 中第一次出现的顺序，依次映射到 a,b,c,... */
		if _, ok := m[c]; !ok {
			m[c] = byte(len(m)) + 'a'
		}
	}

	res := make([]byte, len(s))
	for i, c := range s {
		/**按照字符的映射关系，映射 s 到 res */
		res[i] = m[c]
	}

	return string(res)
}
