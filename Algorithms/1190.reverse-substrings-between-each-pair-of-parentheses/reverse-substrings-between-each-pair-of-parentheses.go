package problem1190

import "strings"

// ref: https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/discuss/383670/JavaC%2B%2BPython-Why-not-O(N)

func reverseParentheses(s string) string {
	n := len(s)
	stack, top := make([]int, n), -1
	pair := make([]int, n)

	for i := 0; i < n; i++ {
		switch s[i] {
		case '(':
			top++
			stack[top] = i
		case ')':
			j := stack[top]
			top--
			pair[i], pair[j] = j, i
		}
	}
	var sb strings.Builder

	for i, d := 0, 1; i < n; i += d {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			d = -d
		} else {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}
