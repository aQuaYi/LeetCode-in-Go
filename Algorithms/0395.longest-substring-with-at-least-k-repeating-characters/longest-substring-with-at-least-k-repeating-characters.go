package Problem0395

func longestSubstring(s string, k int) int {
	count := make(map[byte]int, len(s))
	maxCount := 0
	for i := range s {
		count[s[i]]++
		maxCount = max(maxCount, count[s[i]])
	}
	if maxCount < k {
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

	res := 0
	for _, b := range lessLetters {
		for i, bs := range s {
			if b == bs {
				res = max(res, longestSubstring(s[:i], k))
				res = max(res, longestSubstring(s[i+1:], k))
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
