package Problem0567

import (
	"strings"
)

func checkInclusion(s1, s2 string) bool {
	s1 = reverse(s1)
	return strings.Contains(s2, s1)
}

func reverse(s string) string {
	bs := []byte(s)
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
