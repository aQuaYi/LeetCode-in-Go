package problem0290

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(str, " ")

	if len(ps) != len(ss) {
		return false
	}

	return isMatch(ps, ss) && isMatch(ss, ps)
}

func isMatch(s1, s2 []string) bool {
	size := len(s1)

	m := make(map[string]string, size)

	var i int
	var w string
	var ok bool

	for i = 0; i < size; i++ {
		if w, ok = m[s1[i]]; ok {
			if w != s2[i] {
				return false
			}
		} else {
			m[s1[i]] = s2[i]
		}
	}
	return true
}
