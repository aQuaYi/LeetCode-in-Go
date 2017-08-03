package Problem0030

import "sort"

func findSubstring(s string, words []string) []int {
	// 题目已经确保了 len(words) > 0
	lenWord := len(words[0])

	sort.Strings(words)
	interval := lenWord * len(words)

	res := []int{}

	for i := 0; i < lenWord; i++ {
		j := i
		for j <= len(s)-interval {
			if isComposable(words, s[j:j+interval]) {
				res = append(res, j)
			}
			j += lenWord
		}
	}

	return res
}

// 如果words中的单词经过排列组合后，可以构成s，返回 true
func isComposable(words []string, s string) bool {
	temp := split(s, len(words[0]))
	sort.Strings(temp)

	for i, w := range words {
		if w != temp[i] {
			return false
		}
	}
	return true
}

func split(s string, length int) []string {
	res := make([]string, len(s)/length)

	for i := range res {
		res[i] = s[i*length : (i+1)*length]
	}

	return res
}
