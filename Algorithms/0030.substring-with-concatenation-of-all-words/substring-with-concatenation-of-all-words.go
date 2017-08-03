package Problem0030

import "sort"

func findSubstring(s string, words []string) []int {
	// 题目已经确保了 len(words) > 0
	lenWord := len(words[0])
	numWords := len(words)
	interval := lenWord * numWords

	if lenWord == 0 {
		res := make([]int, len(s)+1)
		for i := range res {
			res[i] = i
		}
		return res
	}

	sort.Strings(words)

	res := []int{}

	for i := 0; i < lenWord; i++ {
		j := i
		for j <= len(s)-interval {
			if equal(words, s[j:j+interval]) {
				res = append(res, j)
			}
			j += lenWord
		}
	}

	return res
}

func equal(words []string, s string) bool {
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
