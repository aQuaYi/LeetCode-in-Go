package problem1163

import "strings"

func lastSubstring(s string) string {
	m, k := "", 'z'
	mark := m + string(k)
	count := strings.Count(s, mark)
	for count != 1 {
		if count == 0 {
			k--
		} else {
			double := mark + mark
			for strings.Contains(s, double) {
				mark = double
				double = mark + mark
			}
			if strings.HasSuffix(s, mark) {
				break
			}
			m, k = mark, 'z'
		}
		mark = m + string(k)
		count = strings.Count(s, mark)
	}

	index := strings.Index(s, mark)

	return s[index:]
}
