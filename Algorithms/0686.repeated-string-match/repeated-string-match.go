package Problem0686

import (
	"strings"
)

func repeatedStringMatch(a, b string) int {

	c := a
	res := 1
	for len(c) < len(b) {
		c += a
		res++
	}

	for len(a)*2+len(b) > len(c) {
		if strings.Contains(c, b) {
			return res
		}
		c += a
		res++

	}

	return -1
}
