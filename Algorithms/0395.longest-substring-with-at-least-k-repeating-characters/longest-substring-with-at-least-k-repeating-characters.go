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

	var b byte
	var c int

	// useless 收集了所有没有达到 k 次的字母
	useless := make([]string, 0, len(count))
	for b, c = range count {
		if c < k {
			useless = append(useless, string(b))
		}
	}

	if len(useless) == 0 {
		// 所有的字母都达到了 k 次
		return len(s)
	}

	var strB string
	for _, strB = range useless {
		s = strings.Replace(s, strB, ",", -1)
	}

	ss := strings.Split(s, ",")
// ss 中
	maxLen := 0
	for _, s = range ss {
		if len(s) >= k {
			maxLen = max(maxLen, len(s))
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
