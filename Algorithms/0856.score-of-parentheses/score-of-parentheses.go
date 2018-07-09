package problem0856

import "strings"

func scoreOfParentheses(s string) int {
	s = strings.Replace(s, "()", "1", -1)
	return helper(s)
}

func helper(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == '1' {
		return 1 + helper(s[1:])
	}
	i := match(s)
	return 2*helper(s[1:i]) + helper(s[i+1:])
}

// 寻找与 s[0] == '('
func match(s string) int {
	count := 1
	size := len(s)
	i := 1
	for ; i < size; i++ {
		switch s[i] {
		case '(':
			count++
		case ')':
			count--
		}
		if count == 0 {
			break
		}
	}
	return i
}
