package problem0791

import (
	"strings"
)

func customSortString(S string, T string) string {
	res := ""
	for i := 0; i < len(S); i++ {
		count := strings.Count(T, S[i:i+1])
		res += strings.Repeat(S[i:i+1], count)
		T = strings.Replace(T, S[i:i+1], "", -1)
	}

	return res + T
}
