package problem1071

import "strings"

func gcdOfStrings(s1, s2 string) string {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	l1, l2 := len(s1), len(s2)
	max := l2 / 2
	res := ""
	for i := 1; i <= max; i++ {
		if l2%i != 0 || l1%i != 0 {
			continue
		}
		tmp := s2[:l2/i]
		if s2 == strings.Repeat(tmp, i) &&
			s1 == strings.Repeat(tmp, l1/i) {
			res = tmp
			break
		}
	}

	return res
}
