package problem0058

func lengthOfLastWord(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	res := 0
	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}

	return res
}
