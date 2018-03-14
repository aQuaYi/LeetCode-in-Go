package problem0467

func findSubstringInWraproundString(p string) int {
	// count[0] = 4 表示，以 'a' 结尾的连续字符串的最大长度为 4
	// 那么，在符合题意的 subString 中以 'a' 结尾的个数为 4
	// 这样统计起来，既不会遗漏也不会重复
	//
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
