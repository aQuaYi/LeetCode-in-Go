package problem1021

import "strings"

func removeOuterParentheses(S string) string {
	var sb strings.Builder
	i, count := 0, 0
	for j := 0; j < len(S); j++ {
		if S[j] == '(' {
			count++
		} else {
			count--
		}
		if count == 0 {
			sb.WriteString(S[i+1 : j])
			i = j + 1
		}
	}
	return sb.String()
}
