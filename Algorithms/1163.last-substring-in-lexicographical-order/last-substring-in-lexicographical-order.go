package problem1163

import "strings"

func lastSubstring(s string) string {
	sub := make([]byte, 1, 64)
	sub[0] = 'z'

	reduce := func() {
		sub[len(sub)-1]--
	}

	extend := func() {
		sub = append(sub, sub[len(sub)-1])
	}

	double := func() string {
		sub = append(sub, sub...)
		return string(sub)
	}

	half := func() {
		sub = sub[:len(sub)/2]
	}

	count := strings.Count(s, string(sub))
	for count != 1 {
		if count == 0 {
			reduce()
		} else {
			for strings.Contains(s, double()) {
			}
			half()
			if strings.HasSuffix(s, string(sub)) {
				break
			}
			extend()
		}
		count = strings.Count(s, string(sub))
	}

	index := strings.Index(s, string(sub))

	return s[index:]
}
