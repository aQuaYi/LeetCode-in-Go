package problem0921

import "strings"

func minAddToMakeValid(S string) int {
	for strings.Contains(S, "()") {
		S = strings.Replace(S, "()", "", -1)
	}
	return len(S)
}
