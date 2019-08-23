package problem1163

import "strings"

// replace mark
const (
	MARK   = "@"
	DOUBLE = "@@"
)

func lastSubstring(s string) string {
	old := s
	m, k := "", 'z'
	c := 0
	mark := m + string(k)
	count := strings.Count(s, mark)
	for count != 1 {
		if count == 0 {
			k--
		} else {
			s = strings.Replace(s, mark, MARK, -1)
			c++
			for strings.Contains(s, DOUBLE) {
				s = strings.Replace(s, DOUBLE, MARK, -1)
				c *= 2
			}
			m, k = MARK, 'z'
		}
		mark = m + string(k)
		count = strings.Count(s, mark)
	}

	index := strings.Index(s, mark)
	s = s[:index]
	count = strings.Count(s, MARK)
	index += (c - 2) * count

	return old[index:]
}
