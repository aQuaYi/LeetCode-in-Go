package Problem0316

import (
	"strings"
)

func removeDuplicateLetters(s string) string {
	lens := len(s)
	if lens == 0 {
		return ""
	}

	count := [26]int{}
	var ch rune
	for _, ch = range s {
		count[ch-'a']++
	}

	pos := 0
	var i int
	for i = range s {
		if s[i] < s[pos] {
			pos = i
		}
		count[s[i]-'a']--
		if count[s[i]-'a'] == 0 {
			break
		}
	}

	return string(s[pos]) + removeDuplicateLetters(strings.Replace(s[pos+1:], string(s[pos]), "", -1))
}
