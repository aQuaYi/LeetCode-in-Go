package problem1071

import "strings"

func gcdOfStrings(s1, s2 string) string {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	len1, len2 := len(s1), len(s2)
	res := ""
	for p := 1; p <= len2; p++ {
		if len2%p != 0 || len1%(len2/p) != 0 {
			continue
		}
		tmp := s2[:len2/p]
		if s2 == strings.Repeat(tmp, p) &&
			s1 == strings.Repeat(tmp, len1*p/len2) {
			res = tmp
			break
		}
	}

	return res
}
