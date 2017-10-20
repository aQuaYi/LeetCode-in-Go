package Problem0316

func removeDuplicateLetters(s string) string {
	cnt := make([]int, 26)
	pos := 0
	lens := len(s)
	var i int
	for i = 0; i < lens; i++ {
		cnt[s[i]-'a']++
	}
	for i = 0; i < lens; i++ {
		if s[i] < s[pos] {
			pos = i
		}
		cnt[s[i]-'a']--
		if cnt[s[i]-'a'] == 0 {
			break
		}
	}
	if lens == 0 {
		return ""
	}

	newS := ""
	for i = pos + 1; i < lens; i++ {
		if s[i] != s[pos] {
			newS += string(s[i])
		}
	}

	return string(s[pos]) + removeDuplicateLetters(newS)
}
