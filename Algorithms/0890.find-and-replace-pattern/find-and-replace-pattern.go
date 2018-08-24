package problem0890

func findAndReplacePattern(words []string, pattern string) []string {
	np := normalize(pattern)
	res := make([]string, 0, len(words))
	for _, w := range words {
		if normalize(w) == np {
			res = append(res, w)
		}
	}
	return res
}

// 把 str 标准化
func normalize(str string) string {
	m := make(map[rune]byte, len(str))
	/**按照字符在 s 中第一次出现的顺序，依次映射到 a,b,c,... */
	for _, c := range str {
		if _, ok := m[c]; !ok {
			m[c] = byte(len(m)) + 'a'
		}
	}
	/**按照字符的映射关系，映射 s 到 res */
	res := make([]byte, len(str))
	for i, c := range str {
		res[i] = m[c]
	}
	return string(res)
}
