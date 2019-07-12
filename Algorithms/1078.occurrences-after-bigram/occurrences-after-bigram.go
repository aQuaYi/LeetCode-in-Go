package problem1078

import "strings"

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	n := len(words)
	res := make([]string, 0, n)
	for i := 0; i+2 < n; i++ {
		if words[i] == first &&
			words[i+1] == second {
			res = append(res, words[i+2])
		}
	}
	return res
}
