package Problem0767

func reorganizeString(s string) string {
	n := len(s)
	bs := []byte(s)

	count := [26]int{}
	maxCount := 0
	for _, b := range bs {
		count[b-'a']++
		maxCount = max(maxCount, count[b-'a'])
	}

	if (n%2 == 0 && maxCount > n/2) ||
		(n%2 == 1 && maxCount > (n/2+1)) {
		return ""
	}

	res := make([]byte, len(bs))

	idx := 0

	for i := 0; i < 26; i++ {
		if count[i] == 0 {
			continue
		}
		b := byte('a' + i)

		for count[i] > 0 {
			res[idx] = b
			count[i]--
			idx = (idx + 2) % n
		}

	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
