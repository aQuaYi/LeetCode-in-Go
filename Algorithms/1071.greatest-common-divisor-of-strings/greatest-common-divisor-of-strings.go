package problem1071

import "strings"

func gcdOfStrings(s1, s2 string) string {
	if len(s1) > len(s2) {
		return gcd(s1, s2)
	}
	return gcd(s2, s1)
}

func gcd(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	if l1%l2 == 0 {
		if s1 == strings.Repeat(s2, l1/l2) {
			return s2
		}
		return ""
	}
	return gcd(s2, s1[l1/l2*l2:])
}
