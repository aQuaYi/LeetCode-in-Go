package problem1021

import "strings"

func removeOuterParentheses(S string) string {
	var sb strings.Builder
	i, count, size := 0, 0, len(S)
	for j := 0; j < size; j++ {
		if S[j] == '(' {
			count++
			continue
		}
		count--
		if count == 0 { // S[i] and S[j] are outer parentheses
			sb.WriteString(S[i+1 : j])
			i = j + 1
		}
	}
	return sb.String()
}
