package problem0820

import (
	"sort"
)

func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		curi := len(words[i])
		curj := len(words[j])
		for curi > 0 && curj > 0 {
			curi--
			curj--
			if words[i][curi] == words[j][curj] {
				continue
			}
			if words[i][curi] > words[j][curj] {
				return true
			}
			return false
		}
		// 短的排在前面
		return curi == 0
	})

	words = append(words, "")
	res, i := 0, 1
	for ; i < len(words); i++ {
		if !endWith(words[i], words[i-1]) {
			res += len(words[i-1]) + 1
		}
	}

	return res
}

// if s end with post, return true
func endWith(s, post string) bool {
	if len(s) < len(post) {
		return false
	}
	return s[len(s)-len(post):] == post
}
