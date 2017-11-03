package Problem0434

import (
	"strings"
)

func countSegments(s string) int {
	ss := strings.Split(s, " ")
	res := len(ss)
	for i := range ss {
		if ss[i] == "" {
			res--
		}
	}
	return res
}
