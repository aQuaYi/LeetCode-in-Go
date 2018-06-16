package problem0844

func backspaceCompare(S string, T string) bool {
	i := len(S)
	j := len(T)

	for i >= 0 || j >= 0 {
		i = nextIndex(&S, i)
		j = nextIndex(&T, j)

		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}

	}

	return i == j
}

// 返回 s[:i] 中，不是 '#' 的字符的最大的索引号
func nextIndex(s *string, i int) int {
	i--
	count := 0
	for i >= 0 && ((*s)[i] == '#' || count > 0) {
		if (*s)[i] == '#' {
			count++
		} else {
			count--
		}
		i--
	}
	return i
}
