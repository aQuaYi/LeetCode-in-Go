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
			res := make([]int, 2, len(s))
			res[0] = a
			res[1] = b
			if find(a, b, s[j:], &res) {
				return res
			}
		}
	}
	return nil
}

func find(a, b int, s string, res *[]int) bool {
	if s == "" {
		return true
	}

	c := a + b
	if c > 1<<31-1 {
		return false
	}
	cs := strconv.Itoa(c)
	csl := len(cs)

	if len(s) >= csl && s[:csl] == cs {
		*res = append(*res, c)
		return find(b, c, s[csl:], res)
	}

	return false
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
