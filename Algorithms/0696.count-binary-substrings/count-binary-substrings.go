package Problem0696

import (
	"strings"
)

func countBinarySubstrings(s string) int {
	s = strings.Replace(s, "01", "0,1", -1)
	s = strings.Replace(s, "10", "1,0", -1)
	ss := strings.Split(s, ",")

	res := 0
	for i := 1; i < len(ss); i++ {
		res += min(len(ss[i-1]), len(ss[i]))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
