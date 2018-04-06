package problem0784

import (
	"strings"
)

func letterCasePermutation(S string) []string {
	s := strings.ToLower(S)

	bytess := make([][]byte, 1<<countLetters(s))
	for i := range bytess {
		bytess[i] = []byte(s)
	}

	base := 2
	half := base / 2
	d := byte('a' - 'A')
	for i := len(s) - 1; i >= 0; i-- {
		if !isLetter(s[i]) {
			continue
		}
		for j := 0; j < len(bytess); j++ {
			if j%base < half {
				continue
			}
			bytess[j][i] -= d
		}

		base *= 2
		half *= 2
	}

	res := make([]string, len(bytess))
	for i := range res {
		res[i] = string(bytess[i])
	}

	return res
}

// 统计 s 中的字母个数
func countLetters(s string) uint {
	res := uint(0)
	for i := range s {
		if isLetter(s[i]) {
			res++
		}
	}
	return res
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z'
}
