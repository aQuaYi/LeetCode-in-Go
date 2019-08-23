package problem1163

import "strings"

func lastSubstring(s string) string {
	var index int
	for i := 'z'; i >= 'a'; i-- {
		index = strings.Index(s, string(i))
		if index != -1 {
			break
		}
	}
	return s[index:]
}
