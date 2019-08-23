package problem1163

import "strings"

func lastSubstring(s string) string {
	sub := make([]byte, 1, 128)
	sub[0] = 'z'

	reduce := func() {
		sub[len(sub)-1]--
	}

	extend := func() {
		sub = append(sub, sub[0])
	}

	// operating string is expensive
	// so I have to make this function a little more responsibility
	// check and change the sub
	containsDouble := func() bool {
		sub = append(sub, sub...)
		if strings.Contains(s, string(sub)) {
			return true
		}
		sub = sub[:len(sub)/2]
		return false
	}

	count := strings.Count(s, string(sub))
	for count != 1 {
		if count == 0 {
			reduce()
		} else if !containsDouble() {
			extend()
		}
		count = strings.Count(s, string(sub))
	}

	index := strings.Index(s, string(sub))

	return s[index:]
}
