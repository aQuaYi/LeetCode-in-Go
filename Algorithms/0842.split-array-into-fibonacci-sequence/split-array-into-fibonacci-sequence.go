package problem0842

import (
	"strconv"
)

func splitIntoFibonacci(s string) []int {
	for i := 1; i <= len(s)/2; i++ {
		if isLeadingZero(s[:i]) {
			break
		}
		for j := i + 1; max(i, j-i) <= len(s)-j; j++ {
			if isLeadingZero(s[i:j]) {
				break
			}
			a, _ := strconv.Atoi(s[:i])
			b, _ := strconv.Atoi(s[i:j])
			res := make([]int, 0, len(s))
			res = append(res, a, b)
			if find(a, b, s[j:], &res) {
				return res
			}
		}
	}
	return nil
}

func find(a, b int, s string, res *[]int) bool {
	for len(s) > 0 {
		c := a + b
		if c > 1<<31-1 {
			return false
		}
		cs := strconv.Itoa(c)
		csl := len(cs)

		if len(s) < csl || s[:csl] != cs {
			return false
		}

		*res = append(*res, c)
		a, b = b, c
		s = s[csl:]
	}

	return true
}

func isLeadingZero(s string) bool {
	return len(s) > 1 && s[0] == '0'
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
