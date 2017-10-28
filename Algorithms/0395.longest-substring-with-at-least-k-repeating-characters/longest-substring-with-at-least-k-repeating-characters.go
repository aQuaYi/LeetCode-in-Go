package Problem0395

import (
	"strings"
)

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	// count 中，记录了每个字母出现的次数
	count := make(map[byte]int, len(s))
	// maxCount 出现最多字母的出现次数
	maxCount := 0
	for i := range s {
		count[s[i]]++
		maxCount = max(maxCount, count[s[i]])
	}
	if maxCount < k {
		// 没有字母达到 k 次
		return 0
	}

	lessLetters := ""
	for b, c := range count {
		if c < k {
			lessLetters += string(b)
		}
	}

	if len(lessLetters) == 0 {
		return len(s)
	}

	i := 0
	for !strings.Contains(lessLetters, string(s[i])) {
		i++
	}
	j := i + 1
	for j < len(s) && strings.Contains(lessLetters, string(s[j])) {
		j++
	}

	// s[i:j] 中所有的字母的个数都 <k
	// 所以，以 s[i:j] 分割 s，在两边寻找 longestSubstring
	return max(longestSubstring(s[:i], k), longestSubstring(s[j:], k))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
