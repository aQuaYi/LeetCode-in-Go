package Problem0424

func characterReplacement(s string, k int) int {
	count := [128]int{}
	var maxLen, start int

	for end := 0; end < len(s); end++ {
		count[s[end]]++
		maxLen = max(maxLen, count[s[end]])
		if maxLen+k <= end-start {
			count[s[start]]--
			start++
		}
	}

	return len(s) - start
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
