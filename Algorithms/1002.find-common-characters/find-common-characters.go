package problem1002

import "strings"

func commonChars(A []string) []string {
	res := make([]string, 0, 128)
	for i := 'a'; i <= 'z'; i++ {
		sub := string(i)
		count := strings.Count(A[0], sub)
		for j := 1; j < len(A) && count > 0; j++ {
			count = min(count, strings.Count(A[j], sub))
		}
		for count > 0 {
			res = append(res, sub)
			count--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
