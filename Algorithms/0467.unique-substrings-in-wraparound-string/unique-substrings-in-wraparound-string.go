package Problem0467

func findSubstringInWraproundString(p string) int {
	count := [26]int{}

	length := 0

	for i := 0; i < len(p); i++ {
		if 0 < i &&
			(p[i-1]+1 == p[i] || p[i-1] == p[i]+25) {
			length++
		} else {
			length = 1
		}

		b := p[i] - 'a'
		count[b] = max(count[b], length)
	}

	res := 0
	for i := 0; i < 26; i++ {
		res += count[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
